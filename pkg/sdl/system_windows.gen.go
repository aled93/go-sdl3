package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&SetWindowsMessageHook, lib, "SDL_SetWindowsMessageHook")
        
    })
}
