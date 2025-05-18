package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&OpenURL, lib, "SDL_OpenURL")
        
    })
}
