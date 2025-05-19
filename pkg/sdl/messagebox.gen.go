package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&ShowMessageBox, lib, "SDL_ShowMessageBox")
        purego.RegisterLibFunc(&ShowSimpleMessageBox, lib, "SDL_ShowSimpleMessageBox")
        
    })
}
