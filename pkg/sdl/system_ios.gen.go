package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&SetiOSAnimationCallback, lib, "SDL_SetiOSAnimationCallback")
        purego.RegisterLibFunc(&SetiOSEventPump, lib, "SDL_SetiOSEventPump")
        purego.RegisterLibFunc(&OnApplicationDidChangeStatusBarOrientation, lib, "SDL_OnApplicationDidChangeStatusBarOrientation")
        
    })
}
