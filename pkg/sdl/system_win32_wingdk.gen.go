//go:build windows || WINGDK
package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetDirect3D9AdapterIndex, "SDL_GetDirect3D9AdapterIndex" },
            { &GetDXGIOutputInfo, "SDL_GetDXGIOutputInfo" },
        }
    })
}
