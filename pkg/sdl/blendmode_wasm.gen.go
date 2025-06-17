//go:build wasm

package sdl

import (
  "unsafe"
)



//go:wasmimport sdl3 SDL_ComposeCustomBlendMode
func __SDL_ComposeCustomBlendMode(int32, int32, int32, int32, int32, int32) int32



func __gowrap__SDL_ComposeCustomBlendMode(srcColorFactor BlendFactor, dstColorFactor BlendFactor, colorOperation BlendOperation, srcAlphaFactor BlendFactor, dstAlphaFactor BlendFactor, alphaOperation BlendOperation) (__result BlendMode) {
	__result = (BlendMode)(__SDL_ComposeCustomBlendMode(*(*int32)(unsafe.Pointer(&srcColorFactor)), *(*int32)(unsafe.Pointer(&dstColorFactor)), *(*int32)(unsafe.Pointer(&colorOperation)), *(*int32)(unsafe.Pointer(&srcAlphaFactor)), *(*int32)(unsafe.Pointer(&dstAlphaFactor)), *(*int32)(unsafe.Pointer(&alphaOperation))))
	return
}

 
func init() {
	ComposeCustomBlendMode = __gowrap__SDL_ComposeCustomBlendMode
}
