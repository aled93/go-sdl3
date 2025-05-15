//go:build GDK
package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetGDKTaskQueue, lib, "SDL_GetGDKTaskQueue")
        purego.RegisterLibFunc(&GetGDKDefaultUser, lib, "SDL_GetGDKDefaultUser")
        
    })
}
