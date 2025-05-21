package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetGlobalProperties, "SDL_GetGlobalProperties" },
            { &CreateProperties, "SDL_CreateProperties" },
            { &CopyProperties, "SDL_CopyProperties" },
            { &LockProperties, "SDL_LockProperties" },
            { &UnlockProperties, "SDL_UnlockProperties" },
            { &SetPointerPropertyWithCleanup, "SDL_SetPointerPropertyWithCleanup" },
            { &SetPointerProperty, "SDL_SetPointerProperty" },
            { &SetStringProperty, "SDL_SetStringProperty" },
            { &SetNumberProperty, "SDL_SetNumberProperty" },
            { &SetFloatProperty, "SDL_SetFloatProperty" },
            { &SetBooleanProperty, "SDL_SetBooleanProperty" },
            { &HasProperty, "SDL_HasProperty" },
            { &GetPropertyType, "SDL_GetPropertyType" },
            { &GetPointerProperty, "SDL_GetPointerProperty" },
            { &GetStringProperty, "SDL_GetStringProperty" },
            { &GetNumberProperty, "SDL_GetNumberProperty" },
            { &GetFloatProperty, "SDL_GetFloatProperty" },
            { &GetBooleanProperty, "SDL_GetBooleanProperty" },
            { &ClearProperty, "SDL_ClearProperty" },
            { &EnumerateProperties, "SDL_EnumerateProperties" },
            { &DestroyProperties, "SDL_DestroyProperties" },
        }
    })
}
