package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetNumCameraDrivers, lib, "SDL_GetNumCameraDrivers")
        purego.RegisterLibFunc(&GetCameraDriver, lib, "SDL_GetCameraDriver")
        purego.RegisterLibFunc(&GetCurrentCameraDriver, lib, "SDL_GetCurrentCameraDriver")
        purego.RegisterLibFunc(&GetCameras, lib, "SDL_GetCameras")
        purego.RegisterLibFunc(&GetCameraSupportedFormats, lib, "SDL_GetCameraSupportedFormats")
        purego.RegisterLibFunc(&GetCameraName, lib, "SDL_GetCameraName")
        purego.RegisterLibFunc(&GetCameraPosition, lib, "SDL_GetCameraPosition")
        purego.RegisterLibFunc(&OpenCamera, lib, "SDL_OpenCamera")
        purego.RegisterLibFunc(&GetCameraPermissionState, lib, "SDL_GetCameraPermissionState")
        purego.RegisterLibFunc(&GetCameraID, lib, "SDL_GetCameraID")
        purego.RegisterLibFunc(&GetCameraProperties, lib, "SDL_GetCameraProperties")
        purego.RegisterLibFunc(&GetCameraFormat, lib, "SDL_GetCameraFormat")
        purego.RegisterLibFunc(&AcquireCameraFrame, lib, "SDL_AcquireCameraFrame")
        purego.RegisterLibFunc(&ReleaseCameraFrame, lib, "SDL_ReleaseCameraFrame")
        purego.RegisterLibFunc(&CloseCamera, lib, "SDL_CloseCamera")
        
    })
}
