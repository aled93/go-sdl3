package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetGlobalProperties, lib, "SDL_GetGlobalProperties")
        purego.RegisterLibFunc(&CreateProperties, lib, "SDL_CreateProperties")
        purego.RegisterLibFunc(&CopyProperties, lib, "SDL_CopyProperties")
        purego.RegisterLibFunc(&LockProperties, lib, "SDL_LockProperties")
        purego.RegisterLibFunc(&UnlockProperties, lib, "SDL_UnlockProperties")
        purego.RegisterLibFunc(&SetPointerPropertyWithCleanup, lib, "SDL_SetPointerPropertyWithCleanup")
        purego.RegisterLibFunc(&SetPointerProperty, lib, "SDL_SetPointerProperty")
        purego.RegisterLibFunc(&SetStringProperty, lib, "SDL_SetStringProperty")
        purego.RegisterLibFunc(&SetNumberProperty, lib, "SDL_SetNumberProperty")
        purego.RegisterLibFunc(&SetFloatProperty, lib, "SDL_SetFloatProperty")
        purego.RegisterLibFunc(&SetBooleanProperty, lib, "SDL_SetBooleanProperty")
        purego.RegisterLibFunc(&HasProperty, lib, "SDL_HasProperty")
        purego.RegisterLibFunc(&GetPropertyType, lib, "SDL_GetPropertyType")
        purego.RegisterLibFunc(&GetPointerProperty, lib, "SDL_GetPointerProperty")
        purego.RegisterLibFunc(&GetStringProperty, lib, "SDL_GetStringProperty")
        purego.RegisterLibFunc(&GetNumberProperty, lib, "SDL_GetNumberProperty")
        purego.RegisterLibFunc(&GetFloatProperty, lib, "SDL_GetFloatProperty")
        purego.RegisterLibFunc(&GetBooleanProperty, lib, "SDL_GetBooleanProperty")
        purego.RegisterLibFunc(&ClearProperty, lib, "SDL_ClearProperty")
        purego.RegisterLibFunc(&EnumerateProperties, lib, "SDL_EnumerateProperties")
        purego.RegisterLibFunc(&DestroyProperties, lib, "SDL_DestroyProperties")
        
    })
}
