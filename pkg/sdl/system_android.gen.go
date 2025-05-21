package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetAndroidJNIEnv, "SDL_GetAndroidJNIEnv" },
            { &GetAndroidActivity, "SDL_GetAndroidActivity" },
            { &GetAndroidSDKVersion, "SDL_GetAndroidSDKVersion" },
            { &IsChromebook, "SDL_IsChromebook" },
            { &IsDeXMode, "SDL_IsDeXMode" },
            { &SendAndroidBackButton, "SDL_SendAndroidBackButton" },
            { &GetAndroidInternalStoragePath, "SDL_GetAndroidInternalStoragePath" },
            { &GetAndroidExternalStorageState, "SDL_GetAndroidExternalStorageState" },
            { &GetAndroidExternalStoragePath, "SDL_GetAndroidExternalStoragePath" },
            { &GetAndroidCachePath, "SDL_GetAndroidCachePath" },
            { &RequestAndroidPermission, "SDL_RequestAndroidPermission" },
            { &ShowAndroidToast, "SDL_ShowAndroidToast" },
            { &SendAndroidMessage, "SDL_SendAndroidMessage" },
        }
    })
}
