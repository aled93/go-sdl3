package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetSensors, "SDL_GetSensors" },
            { &GetSensorNameForID, "SDL_GetSensorNameForID" },
            { &GetSensorTypeForID, "SDL_GetSensorTypeForID" },
            { &GetSensorNonPortableTypeForID, "SDL_GetSensorNonPortableTypeForID" },
            { &OpenSensor, "SDL_OpenSensor" },
            { &GetSensorFromID, "SDL_GetSensorFromID" },
            { &GetSensorProperties, "SDL_GetSensorProperties" },
            { &GetSensorName, "SDL_GetSensorName" },
            { &GetSensorType, "SDL_GetSensorType" },
            { &GetSensorNonPortableType, "SDL_GetSensorNonPortableType" },
            { &GetSensorID, "SDL_GetSensorID" },
            { &GetSensorData, "SDL_GetSensorData" },
            { &CloseSensor, "SDL_CloseSensor" },
            { &UpdateSensors, "SDL_UpdateSensors" },
        }
    })
}
