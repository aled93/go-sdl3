package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetNumCameraDrivers, "SDL_GetNumCameraDrivers" },
            { &GetCameraDriver, "SDL_GetCameraDriver" },
            { &GetCurrentCameraDriver, "SDL_GetCurrentCameraDriver" },
            { &GetCameras, "SDL_GetCameras" },
            { &GetCameraSupportedFormats, "SDL_GetCameraSupportedFormats" },
            { &GetCameraName, "SDL_GetCameraName" },
            { &GetCameraPosition, "SDL_GetCameraPosition" },
            { &OpenCamera, "SDL_OpenCamera" },
            { &GetCameraPermissionState, "SDL_GetCameraPermissionState" },
            { &GetCameraID, "SDL_GetCameraID" },
            { &GetCameraProperties, "SDL_GetCameraProperties" },
            { &GetCameraFormat, "SDL_GetCameraFormat" },
            { &AcquireCameraFrame, "SDL_AcquireCameraFrame" },
            { &ReleaseCameraFrame, "SDL_ReleaseCameraFrame" },
            { &CloseCamera, "SDL_CloseCamera" },
        }
    })
}
