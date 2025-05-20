package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&Metal_CreateView, lib, "SDL_Metal_CreateView")
        purego.RegisterLibFunc(&Metal_DestroyView, lib, "SDL_Metal_DestroyView")
        purego.RegisterLibFunc(&Metal_GetLayer, lib, "SDL_Metal_GetLayer")
        
    })
}
