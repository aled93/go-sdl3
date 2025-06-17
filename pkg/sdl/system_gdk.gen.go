//go:build GDK

package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetGDKTaskQueue, "SDL_GetGDKTaskQueue" },
            { &GetGDKDefaultUser, "SDL_GetGDKDefaultUser" },
        }
    })
}
