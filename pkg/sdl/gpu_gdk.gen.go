//go:build GDK
package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GDKSuspendGPU, lib, "SDL_GDKSuspendGPU")
        purego.RegisterLibFunc(&GDKResumeGPU, lib, "SDL_GDKResumeGPU")
        
    })
}
