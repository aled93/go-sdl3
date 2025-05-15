//go:build windows || WINGDK
package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetDirect3D9AdapterIndex, lib, "SDL_GetDirect3D9AdapterIndex")
        purego.RegisterLibFunc(&GetDXGIOutputInfo, lib, "SDL_GetDXGIOutputInfo")
        
    })
}
