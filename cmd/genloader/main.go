package main

import (
	_ "embed"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/fs"
	"os"
	"path"
	"regexp"
	"slices"
	"strings"
	"text/template"
)

var (
	reSdl3Extern = regexp.MustCompile(`//go:sdl3extern(?:\((\w+)\)|)`)
	//go:embed loader.tmpl
	loaderTmplSrc string
	loaderTmpl    = template.Must(template.New("loader").Parse(loaderTmplSrc))

	//go:embed wasm_import.tmpl
	wasmImportTmplSrc string
	wasmImportTmpl    = template.Must(template.New("wasm_import").Parse(wasmImportTmplSrc))

	optVerbose = flag.Bool("verbose", false, "Print to stdout progress")
)

func main() {
	flag.Parse()

	// pkgs, err := packages.Load(&packages.Config{
	// 	Mode: packages.LoadAllSyntax,
	// 	Dir:  "./pkg/sdl",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// pkg := pkgs[0]

	fset := token.NewFileSet()
	rootDir := os.DirFS("./pkg/sdl")
	goFiles, err := fs.Glob(rootDir, "**.go")
	if err != nil {
		panic(err)
	}

	var filesPath []string
	var files []*ast.File
	for _, fname := range goFiles {
		if strings.HasSuffix(strings.ToLower(fname), ".gen.go") {
			if *optVerbose {
				fmt.Printf("Ignored generated file \"%s\"\n", path.Join("./pkg/sdl", fname))
			}

			continue
		}

		fpath := path.Join("./pkg/sdl", fname)
		f, err := parser.ParseFile(fset, fpath, nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}

		files = append(files, f)
		filesPath = append(filesPath, fpath)
	}

	slices.SortFunc(files, func(a, b *ast.File) int {
		return strings.Compare(a.Name.Name, b.Name.Name)
	})

	for i, file := range files {
		origFilePath := filesPath[i]
		origFileName := path.Base(origFilePath)
		origBaseName := strings.TrimSuffix(origFileName, path.Ext(origFileName))

		if *optVerbose {
			fmt.Printf("Processing \"%s\"\n", origFilePath)
		}

		tmplData := templateData{
			Pkg:      file.Name.Name,
			Filename: origBaseName,
		}

		useRuntime := false
		useUnsafe := false

		for _, cmtGrp := range file.Comments {
			for _, cmt := range cmtGrp.List {
				if buildTag, ok := strings.CutPrefix(cmt.Text, "//go:build "); ok {
					tmplData.BuildTags = buildTag
				} else if buildTag, ok := strings.CutPrefix(cmt.Text, "// +build "); ok {
					tmplData.BuildTags = buildTag
				}
			}
		}

		ast.Inspect(file, func(n ast.Node) bool {
			decl, ok := n.(*ast.GenDecl)
			if !ok || decl.Tok != token.VAR || len(decl.Specs) != 1 {
				return true
			}

			spec, ok := decl.Specs[0].(*ast.ValueSpec)
			if !ok {
				return true
			}

			if decl.Doc != nil {
				for _, cmt := range decl.Doc.List {
					if matches := reSdl3Extern.FindStringSubmatch(cmt.Text); matches != nil {
						funcType, ok := spec.Type.(*ast.FuncType)
						if !ok {
							panic("only variables with func type must be marked with //go:sdl3extern")
						}
						if funcType.TypeParams != nil {
							panic("function with type parameters not allowed")
						}

						var gowrapParams strings.Builder
						var wasmParamTypes []string
						var wasmArgs []string
						var preCallLines []string
						var postCallLines []string

						// map go param types to compatible with wasm import
						for i, param := range funcType.Params.List {
							if len(param.Names) == 0 {
								panic("please provide names for all parameters (" + fset.Position(funcType.Func).String() + ")")
							}

							for j, paramIdent := range param.Names {
								paramName := paramIdent.Name
								wasmArg := paramName

								var typeText strings.Builder
								if err := printer.Fprint(&typeText, fset, param.Type); err != nil {
									panic(err)
								}

								if i+j > 0 {
									gowrapParams.WriteString(", ")
								}
								gowrapParams.WriteString(paramName)
								gowrapParams.WriteRune(' ')
								gowrapParams.WriteString(typeText.String())

								var typeName string
								var isSlice, isPtr, _ bool
								typeName, isSlice = strings.CutPrefix(typeText.String(), "[]")
								typeName, isPtr = strings.CutPrefix(typeName, "*")
								typeName, _ = strings.CutPrefix(typeName, "*")

								if isSlice {
									wasmParamTypes = append(wasmParamTypes, "uintptr")
									wasmArg = fmt.Sprintf(`uintptr(unsafe.Pointer(&%s[0]))`, wasmArg)
									postCallLines = append(postCallLines, "runtime.KeepAlive("+paramName+")")
									useRuntime = true
									useUnsafe = true
								} else if isPtr {
									wasmParamTypes = append(wasmParamTypes, "uintptr")
									wasmArg = fmt.Sprintf(`uintptr(unsafe.Pointer(%s))`, wasmArg)
									useUnsafe = true
								} else {
									knownMappedTyp, ok := knownTypeMap[typeName]
									if !ok {
										panic("unknown type " + typeName + " (" + fset.Position(paramIdent.Pos()).String() + ")")
									} else {
										switch knownMappedTyp {
										case wasmType_Callback:
											wasmParamTypes = append(wasmParamTypes, "uintptr")
											wasmArg = "0 /* TODO: callbacks */"

										case wasmType_Ptr:
											wasmParamTypes = append(wasmParamTypes, "uintptr")
											wasmArg = fmt.Sprintf(`uintptr(unsafe.Pointer(%s))`, wasmArg)
											useUnsafe = true

										case wasmType_String:
											wasmParamTypes = append(wasmParamTypes, "uintptr")
											wasmArg = fmt.Sprintf(`((*[2]uintptr)(unsafe.Pointer(&%s)))[0]`, wasmArg)
											postCallLines = append(postCallLines, "runtime.KeepAlive("+paramName+")")
											useRuntime = true
											useUnsafe = true

										case wasmType_Struct:
											panic("passing struct arguments not supported")

										case wasmType_f32:
											wasmParamTypes = append(wasmParamTypes, "float32")
											wasmArg = fmt.Sprintf(`*(*float32)(unsafe.Pointer(&%s))`, wasmArg)
											useUnsafe = true

										case wasmType_f64:
											wasmParamTypes = append(wasmParamTypes, "float64")
											wasmArg = fmt.Sprintf(`*(*float64)(unsafe.Pointer(&%s))`, wasmArg)
											useUnsafe = true

										case wasmType_i32:
											wasmParamTypes = append(wasmParamTypes, "int32")
											wasmArg = fmt.Sprintf(`*(*int32)(unsafe.Pointer(&%s))`, wasmArg)
											useUnsafe = true

										case wasmType_i64:
											wasmParamTypes = append(wasmParamTypes, "int64")
											wasmArg = fmt.Sprintf(`*(*int64)(unsafe.Pointer(&%s))`, wasmArg)
											useUnsafe = true

										default:
											panic(fmt.Sprintf("unexpected main.wasmType: %#v", knownMappedTyp))
										}
									}
								}

								wasmArgs = append(wasmArgs, wasmArg)
							}
						}

						if funcType.Results != nil && len(funcType.Results.List) > 1 {
							panic("multiple return values not supported (" + fset.Position(funcType.Results.Pos()).String() + ")")
						}

						hasResult := funcType.Results != nil
						hasMappedResult := hasResult
						var resTypeText strings.Builder
						var resTypeMapped string
						var origCallPrefix, origCallPostfix string

						if hasResult {
							printer.Fprint(&resTypeText, fset, funcType.Results.List[0].Type)

							var resTypeName string
							var isSlice, isPtr, _ bool
							resTypeName, isSlice = strings.CutPrefix(resTypeText.String(), "[]")
							sliceElemType := resTypeName
							resTypeName, isPtr = strings.CutPrefix(resTypeName, "*")
							resTypeName, _ = strings.CutPrefix(resTypeName, "*")

							if isSlice {
								resTypeMapped = "uintptr"
								origCallPrefix = "newCSliceRaw[" + sliceElemType + "](" + origCallPrefix
								origCallPostfix = origCallPostfix + ").Collect()"
							} else if isPtr {
								resTypeMapped = "uintptr"
								origCallPrefix = "unsafe.Pointer(" + origCallPrefix
								origCallPostfix = origCallPostfix + ")"
							} else {
								knownMappedTyp, ok := knownTypeMap[resTypeName]
								if !ok {
									panic("unknown type " + resTypeName + " (" + fset.Position(funcType.Results.Pos()).String() + ")")
								} else {
									switch knownMappedTyp {
									case wasmType_Callback:
										panic("return callback not supported (" + fset.Position(funcType.End()).String() + ")")

									case wasmType_Ptr:
										resTypeMapped = "uintptr"
										origCallPrefix = "unsafe.Pointer(" + origCallPrefix
										origCallPostfix = origCallPostfix + ")"

									case wasmType_String:
										resTypeMapped = "uintptr"
										origCallPrefix = "string(newCSliceRaw[byte](" + origCallPrefix
										origCallPostfix = origCallPostfix + ").Collect())"

									case wasmType_Struct:
										hasMappedResult = false
										wasmParamTypes = slices.Insert(wasmParamTypes, 0, "uintptr")
										wasmArgs = slices.Insert(wasmArgs, 0, "uintptr(unsafe.Pointer(&__result))")
										useUnsafe = true

									case wasmType_f32:
										resTypeMapped = "float32"

									case wasmType_f64:
										resTypeMapped = "float64"

									case wasmType_i32:
										if resTypeName == "bool" {
											origCallPostfix = " != 0" + origCallPostfix
										}

										resTypeMapped = "int32"

									case wasmType_i64:
										resTypeMapped = "int64"

									default:
										panic(fmt.Sprintf("unexpected main.wasmType: %#v", knownMappedTyp))
									}
								}
							}
						}

						var wrapperArgs strings.Builder
						var funcSign strings.Builder
						funcSign.WriteRune('(')
						for i, p := range wasmParamTypes {
							if i > 0 {
								wrapperArgs.WriteString(", ")
								funcSign.WriteString(", ")
							}

							wrapperArgs.WriteString(wasmArgs[i])

							funcSign.WriteString(p)
						}
						if !hasMappedResult {
							funcSign.WriteRune(')')
						} else {
							funcSign.WriteString(") ")
							funcSign.WriteString(resTypeMapped)
						}

						origFuncName := ""
						if len(matches[1]) > 0 {
							origFuncName = matches[1]
						} else {
							origFuncName = spec.Names[0].Name
						}
						origFuncName = "SDL_" + origFuncName

						var wrapper strings.Builder
						fmt.Fprintf(&wrapper, "func __gowrap__%s(%s) ", origFuncName, gowrapParams.String())
						if hasResult {
							wrapper.WriteString("(__result ")
							wrapper.WriteString(resTypeText.String())
							wrapper.WriteString(") ")
						}
						wrapper.WriteString("{\n")
						for _, ln := range preCallLines {
							wrapper.WriteRune('\t')
							wrapper.WriteString(ln)
							wrapper.WriteRune('\n')
						}
						wrapper.WriteRune('\t')
						if hasResult && hasMappedResult {
							wrapper.WriteString("__result = (")
							wrapper.WriteString(resTypeText.String())
							wrapper.WriteString(")(")
							wrapper.WriteString(origCallPrefix)
						}
						fmt.Fprintf(&wrapper, "__%s(%s)", origFuncName, wrapperArgs.String())
						if hasResult && hasMappedResult {
							wrapper.WriteString(origCallPostfix)
							wrapper.WriteRune(')')
						}
						wrapper.WriteRune('\n')
						for _, ln := range postCallLines {
							wrapper.WriteRune('\t')
							wrapper.WriteString(ln)
							wrapper.WriteRune('\n')
						}
						if hasResult {
							wrapper.WriteString("\treturn\n")
						}
						wrapper.WriteRune('}')

						tmplData.Sdl3Externs = append(tmplData.Sdl3Externs, externInfo{
							FuncPath:      spec.Names[0].Name,
							FuncOrigName:  origFuncName,
							FuncSignature: funcSign.String(),
						})

						tmplData.Wrappers = append(tmplData.Wrappers, wrapper.String())
					}
				}
			}
			return false
		})

		if useRuntime {
			tmplData.Uses = append(tmplData.Uses, "runtime")
		}
		if useUnsafe {
			tmplData.Uses = append(tmplData.Uses, "unsafe")
		}

		if len(tmplData.Sdl3Externs) == 0 {
			continue
		}

		outputFilePath := strings.TrimSuffix(origFilePath, ".go") + ".gen.go"
		output, err := os.Create(outputFilePath)
		if err != nil {
			panic(err)
		}

		if err := loaderTmpl.Execute(output, &tmplData); err != nil {
			panic(err)
		}
		output.Close()

		wasmTargetOutputFilePath := strings.TrimSuffix(origFilePath, ".go") + "_wasm.gen.go"
		wasmTargetOutput, err := os.Create(wasmTargetOutputFilePath)
		if err != nil {
			panic(err)
		}

		if err := wasmImportTmpl.Execute(wasmTargetOutput, &tmplData); err != nil {
			panic(err)
		}
		wasmTargetOutput.Close()
	}
}

type templateData struct {
	Pkg         string
	Uses        []string
	Filename    string
	BuildTags   string
	Sdl3Externs []externInfo
	Wrappers    []string
}

type externInfo struct {
	FuncPath      string
	FuncOrigName  string
	FuncSignature string
}

// type wrapperInfo struct {
// 	WrappedSignature  string
// }

// func mapGoTypeToWasmCompatible(fset *token.FileSet, typ ast.Expr) (newType string, mapped bool) {
// 	var typeText strings.Builder
// 	if err := printer.Fprint(&typeText, fset, typ); err != nil {
// 		panic(err)
// 	}

// 	var isSlice, isPtr, isPtr2Ptr bool
// 	newType, isSlice = strings.CutPrefix(typeText.String(), "[]")
// 	newType, isPtr = strings.CutPrefix(newType, "*")
// 	newType, isPtr2Ptr = strings.CutPrefix(newType, "*")

// 	if isPtr {
// 		return "i32", true
// 	}

// 	knownMappedTyp, ok := knownTypeMap[newType]

// 	panic("")
// }

type wasmType int

const (
	wasmType_i32 wasmType = iota
	wasmType_i64
	wasmType_f32
	wasmType_f64
	wasmType_Ptr
	wasmType_Callback
	wasmType_String
	wasmType_Struct
)

var knownTypeMap = map[string]wasmType{
	"AudioDeviceID":                     wasmType_i32,
	"AudioFormat":                       wasmType_i32,
	"AudioPostmixCallback":              wasmType_Callback,
	"AudioSpec":                         wasmType_Struct,
	"AudioStream":                       wasmType_Ptr,
	"AudioStreamCallback":               wasmType_Callback,
	"BlendFactor":                       wasmType_i32,
	"BlendMode":                         wasmType_i32,
	"BlendOperation":                    wasmType_i32,
	"bool":                              wasmType_i32,
	"byte":                              wasmType_i32,
	"CallocFunc":                        wasmType_Callback,
	"Camera":                            wasmType_Ptr,
	"CameraID":                          wasmType_i32,
	"CameraSpec":                        wasmType_Struct,
	"CameraPosition":                    wasmType_i32,
	"CleanupPropertyCallback":           wasmType_Callback,
	"ClipboardCleanupCallback":          wasmType_Callback,
	"ClipboardDataCallback":             wasmType_Callback,
	"Color":                             wasmType_i32,
	"Colorspace":                        wasmType_i32,
	"Cursor":                            wasmType_Ptr,
	"DateFormat":                        wasmType_i32,
	"DateTime":                          wasmType_Struct,
	"DialogFileCallback":                wasmType_Callback,
	"DialogFileFilter":                  wasmType_Struct,
	"DisplayID":                         wasmType_i32,
	"DisplayMode":                       wasmType_Struct,
	"EGLAttribArrayCallback":            wasmType_Callback,
	"EGLIntArrayCallback":               wasmType_Callback,
	"EnumerateDirectoryCallback":        wasmType_Callback,
	"EnumeratePropertiesCallback":       wasmType_Callback,
	"Environment":                       wasmType_Ptr,
	"Event":                             wasmType_Struct,
	"EventAction":                       wasmType_i32,
	"EventFilter":                       wasmType_Callback,
	"EventType":                         wasmType_i32,
	"FColor":                            wasmType_Struct,
	"FileDialogType":                    wasmType_i32,
	"FlashOperation":                    wasmType_i32,
	"FlipMode":                          wasmType_i32,
	"float32":                           wasmType_f32,
	"float64":                           wasmType_f64,
	"Folder":                            wasmType_i32,
	"FPoint":                            wasmType_Struct,
	"FRect":                             wasmType_Struct,
	"FreeFunc":                          wasmType_Callback,
	"Gamepad":                           wasmType_Ptr,
	"GamepadAxis":                       wasmType_i32,
	"GamepadButton":                     wasmType_i32,
	"GamepadButtonLabel":                wasmType_i32,
	"GamepadType":                       wasmType_i32,
	"GLAttr":                            wasmType_i32,
	"GLContext":                         wasmType_Ptr,
	"GlobFlags":                         wasmType_i32,
	"GPUBlitInfo":                       wasmType_Struct,
	"GPUBuffer":                         wasmType_Ptr,
	"GPUBufferBinding":                  wasmType_Struct,
	"GPUBufferCreateInfo":               wasmType_Struct,
	"GPUBufferLocation":                 wasmType_Struct,
	"GPUBufferRegion":                   wasmType_Struct,
	"GPUColorTargetInfo":                wasmType_Struct,
	"GPUCommandBuffer":                  wasmType_Ptr,
	"GPUComputePass":                    wasmType_Ptr,
	"GPUComputePipeline":                wasmType_Ptr,
	"GPUComputePipelineCreateInfo":      wasmType_Struct,
	"GPUCopyPass":                       wasmType_Ptr,
	"GPUDepthStencilTargetInfo":         wasmType_Struct,
	"GPUDevice":                         wasmType_Ptr,
	"GPUFence":                          wasmType_Ptr,
	"GPUGraphicsPipeline":               wasmType_Ptr,
	"GPUGraphicsPipelineCreateInfo":     wasmType_Struct,
	"GPUIndexElementSize":               wasmType_i32,
	"GPUPresentMode":                    wasmType_i32,
	"GPURenderPass":                     wasmType_Ptr,
	"GPURenderState":                    wasmType_Struct,
	"GPURenderStateDesc":                wasmType_Struct,
	"GPUSampleCount":                    wasmType_i32,
	"GPUSampler":                        wasmType_Ptr,
	"GPUSamplerCreateInfo":              wasmType_Struct,
	"GPUShader":                         wasmType_Ptr,
	"GPUShaderCreateInfo":               wasmType_Struct,
	"GPUShaderFormat":                   wasmType_i32,
	"GPUStorageBufferReadWriteBinding":  wasmType_Struct,
	"GPUStorageTextureReadWriteBinding": wasmType_Struct,
	"GPUSwapchainComposition":           wasmType_i32,
	"GPUTexture":                        wasmType_Ptr,
	"GPUTextureCreateInfo":              wasmType_Struct,
	"GPUTextureFormat":                  wasmType_i32,
	"GPUTextureLocation":                wasmType_Struct,
	"GPUTextureRegion":                  wasmType_Struct,
	"GPUTextureSamplerBinding":          wasmType_Struct,
	"GPUTextureTransferInfo":            wasmType_Struct,
	"GPUTextureType":                    wasmType_i32,
	"GPUTextureUsageFlags":              wasmType_i32,
	"GPUTransferBuffer":                 wasmType_Ptr,
	"GPUTransferBufferCreateInfo":       wasmType_Struct,
	"GPUTransferBufferLocation":         wasmType_Struct,
	"GPUViewport":                       wasmType_Struct,
	"GUID":                              wasmType_Struct,
	"Haptic":                            wasmType_Ptr,
	"HapticEffect":                      wasmType_Struct,
	"HapticEffectID":                    wasmType_i32,
	"HapticID":                          wasmType_i32,
	"Hat":                               wasmType_i32,
	"HIDDevice":                         wasmType_Ptr,
	"HIDDeviceInfo":                     wasmType_Struct,
	"HintCallback":                      wasmType_Callback,
	"HintPriority":                      wasmType_i32,
	"HitTest":                           wasmType_i32,
	"InitFlags":                         wasmType_i32,
	"InitState":                         wasmType_Struct,
	"int16":                             wasmType_i32,
	"int32":                             wasmType_i32,
	"int64":                             wasmType_i64,
	"int8":                              wasmType_i32,
	"IOSAnimationCallback":              wasmType_Callback,
	"IOStatus":                          wasmType_i32,
	"IOStream":                          wasmType_Ptr,
	"IOStreamInterface":                 wasmType_Struct,
	"IOWhence":                          wasmType_i32,
	"Joystick":                          wasmType_Ptr,
	"JoystickConnectionState":           wasmType_i32,
	"JoystickID":                        wasmType_i32,
	"KeyboardID":                        wasmType_i32,
	"Keycode":                           wasmType_i32,
	"Keymod":                            wasmType_i32,
	"MainThreadCallback":                wasmType_Callback,
	"MallocFunc":                        wasmType_Callback,
	"MessageBoxData":                    wasmType_Struct,
	"MessageBoxFlags":                   wasmType_i32,
	"MetalView":                         wasmType_Ptr,
	"MouseID":                           wasmType_i32,
	"MouseMotionTransformCallback":      wasmType_Callback,
	"NSTimerCallback":                   wasmType_Callback,
	"Palette":                           wasmType_Struct,
	"PathInfo":                          wasmType_Struct,
	"PixelFormat":                       wasmType_i32,
	"PixelFormatDetails":                wasmType_Struct,
	"Point":                             wasmType_Struct,
	"PowerState":                        wasmType_i32,
	"ProgressState":                     wasmType_i32,
	"PropertiesID":                      wasmType_i32,
	"PropertyName":                      wasmType_String,
	"ReallocFunc":                       wasmType_Callback,
	"Rect":                              wasmType_Struct,
	"Renderer":                          wasmType_Ptr,
	"RendererLogicalPresentation":       wasmType_i32,
	"RequestAndroidPermissionCallback":  wasmType_Callback,
	"ScaleMode":                         wasmType_i32,
	"Scancode":                          wasmType_i32,
	"Sensor":                            wasmType_Ptr,
	"SensorID":                          wasmType_i32,
	"SensorType":                        wasmType_i32,
	"size_t":                            wasmType_i32,
	"Storage":                           wasmType_Ptr,
	"StorageInterface":                  wasmType_Struct,
	"string":                            wasmType_String,
	"Surface":                           wasmType_i32,
	"SystemCursor":                      wasmType_i32,
	"Texture":                           wasmType_Struct,
	"TextureAccess":                     wasmType_i32,
	"TextureAddressMode":                wasmType_i32,
	"Time":                              wasmType_i64,
	"TimeFormat":                        wasmType_i32,
	"TimerCallback":                     wasmType_Callback,
	"TimerID":                           wasmType_i32,
	"TouchID":                           wasmType_i64,
	"Tray":                              wasmType_Ptr,
	"TrayCallback":                      wasmType_Callback,
	"TrayEntry":                         wasmType_Ptr,
	"TrayEntryFlags":                    wasmType_i32,
	"TrayMenu":                          wasmType_Ptr,
	"uint":                              wasmType_i64,
	"uint16":                            wasmType_i32,
	"uint32":                            wasmType_i32,
	"uint64":                            wasmType_i64,
	"uint8":                             wasmType_i32,
	"uintptr":                           wasmType_Ptr,
	"Vertex":                            wasmType_Struct,
	"VirtualJoystickDesc":               wasmType_Struct,
	"VkAllocationCallbacks":             wasmType_Callback,
	"VkInstance":                        wasmType_Ptr,
	"VkPhysicalDevice":                  wasmType_Ptr,
	"VkSurfaceKHR":                      wasmType_Ptr,
	"VSync":                             wasmType_i32,
	"Window":                            wasmType_Ptr,
	"WindowFlags":                       wasmType_i64,
	"WindowID":                          wasmType_i32,
	"WindowsMessageHook":                wasmType_Callback,
	"X11EventHook":                      wasmType_Callback,
	"XTaskQueueHandle":                  wasmType_Ptr,
	"XUserHandle":                       wasmType_Ptr,
	"JoystickType":                      wasmType_i32,
	"MouseButtonFlags":                  wasmType_i32,
	"PropertyType":                      wasmType_i32,
	"Sandbox":                           wasmType_i32,
	"TouchDeviceType":                   wasmType_i32,
	"SystemTheme":                       wasmType_i32,
	"DisplayOrientation":                wasmType_i32,
	"EGLDisplay":                        wasmType_Ptr,
	"EGLConfig":                         wasmType_Ptr,
	"EGLSurface":                        wasmType_Ptr,
	"EGLAttrib":                         wasmType_Ptr,
	"EGLint":                            wasmType_i32,
}
