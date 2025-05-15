package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetSensors, lib, "SDL_GetSensors")
        purego.RegisterLibFunc(&GetSensorNameForID, lib, "SDL_GetSensorNameForID")
        purego.RegisterLibFunc(&GetSensorTypeForID, lib, "SDL_GetSensorTypeForID")
        purego.RegisterLibFunc(&GetSensorNonPortableTypeForID, lib, "SDL_GetSensorNonPortableTypeForID")
        purego.RegisterLibFunc(&OpenSensor, lib, "SDL_OpenSensor")
        purego.RegisterLibFunc(&GetSensorFromID, lib, "SDL_GetSensorFromID")
        purego.RegisterLibFunc(&GetSensorProperties, lib, "SDL_GetSensorProperties")
        purego.RegisterLibFunc(&GetSensorName, lib, "SDL_GetSensorName")
        purego.RegisterLibFunc(&GetSensorType, lib, "SDL_GetSensorType")
        purego.RegisterLibFunc(&GetSensorNonPortableType, lib, "SDL_GetSensorNonPortableType")
        purego.RegisterLibFunc(&GetSensorID, lib, "SDL_GetSensorID")
        purego.RegisterLibFunc(&GetSensorData, lib, "SDL_GetSensorData")
        purego.RegisterLibFunc(&CloseSensor, lib, "SDL_CloseSensor")
        purego.RegisterLibFunc(&UpdateSensors, lib, "SDL_UpdateSensors")
        
    })
}
