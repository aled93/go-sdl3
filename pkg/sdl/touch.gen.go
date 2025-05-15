package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetTouchDevices, lib, "SDL_GetTouchDevices")
        purego.RegisterLibFunc(&GetTouchDeviceName, lib, "SDL_GetTouchDeviceName")
        purego.RegisterLibFunc(&GetTouchDeviceType, lib, "SDL_GetTouchDeviceType")
        purego.RegisterLibFunc(&GetTouchFingers, lib, "SDL_GetTouchFingers")
        
    })
}
