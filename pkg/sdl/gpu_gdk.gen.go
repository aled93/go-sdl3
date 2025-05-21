//go:build GDK
package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GDKSuspendGPU, "SDL_GDKSuspendGPU" },
            { &GDKResumeGPU, "SDL_GDKResumeGPU" },
        }
    })
}
