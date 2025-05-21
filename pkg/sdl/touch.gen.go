package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetTouchDevices, "SDL_GetTouchDevices" },
            { &GetTouchDeviceName, "SDL_GetTouchDeviceName" },
            { &GetTouchDeviceType, "SDL_GetTouchDeviceType" },
            { &GetTouchFingers, "SDL_GetTouchFingers" },
        }
    })
}
