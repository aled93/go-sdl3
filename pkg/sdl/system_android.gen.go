package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetAndroidJNIEnv, lib, "SDL_GetAndroidJNIEnv")
        purego.RegisterLibFunc(&GetAndroidActivity, lib, "SDL_GetAndroidActivity")
        purego.RegisterLibFunc(&GetAndroidSDKVersion, lib, "SDL_GetAndroidSDKVersion")
        purego.RegisterLibFunc(&IsChromebook, lib, "SDL_IsChromebook")
        purego.RegisterLibFunc(&IsDeXMode, lib, "SDL_IsDeXMode")
        purego.RegisterLibFunc(&SendAndroidBackButton, lib, "SDL_SendAndroidBackButton")
        purego.RegisterLibFunc(&GetAndroidInternalStoragePath, lib, "SDL_GetAndroidInternalStoragePath")
        purego.RegisterLibFunc(&GetAndroidExternalStorageState, lib, "SDL_GetAndroidExternalStorageState")
        purego.RegisterLibFunc(&GetAndroidExternalStoragePath, lib, "SDL_GetAndroidExternalStoragePath")
        purego.RegisterLibFunc(&GetAndroidCachePath, lib, "SDL_GetAndroidCachePath")
        purego.RegisterLibFunc(&RequestAndroidPermission, lib, "SDL_RequestAndroidPermission")
        purego.RegisterLibFunc(&ShowAndroidToast, lib, "SDL_ShowAndroidToast")
        purego.RegisterLibFunc(&SendAndroidMessage, lib, "SDL_SendAndroidMessage")
        
    })
}
