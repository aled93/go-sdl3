package main

import (
	_ "embed"
	"go/ast"
	"go/parser"
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
	reSdl3Extern = regexp.MustCompile(`//go:sdl3extern(\(\w+\))?`)
	//go:embed loader.tmpl
	loaderTmplSrc string
	loaderTmpl    = template.Must(template.New("loader").Parse(loaderTmplSrc))
)

func main() {
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

		tmplData := templateData{
			Pkg:      file.Name.Name,
			Filename: origBaseName,
		}

		for _, cmtGrp := range file.Comments {
			for _, cmt := range cmtGrp.List {
				if strings.HasPrefix(cmt.Text, "//go:build") || strings.HasPrefix(cmt.Text, "// +build ") {
					tmplData.BuildTags = append(tmplData.BuildTags, cmt.Text)
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
					if matches := reSdl3Extern.FindAllString(cmt.Text, 4); matches != nil {
						orig := ""
						if len(matches) >= 2 {
							orig = matches[1]
						} else {
							orig = spec.Names[0].Name
						}
						orig = "SDL_" + orig
						tmplData.Sdl3Externs = append(tmplData.Sdl3Externs, externInfo{
							FuncPath:     spec.Names[0].Name,
							FuncOrigName: orig,
						})
					}
				}
			}
			return false
		})

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
	}
}

type templateData struct {
	Pkg         string
	Filename    string
	BuildTags   []string
	Sdl3Externs []externInfo
}

type externInfo struct {
	FuncPath     string
	FuncOrigName string
}
