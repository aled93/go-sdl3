package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetHaptics, "SDL_GetHaptics" },
            { &GetHapticNameForID, "SDL_GetHapticNameForID" },
            { &OpenHaptic, "SDL_OpenHaptic" },
            { &GetHapticFromID, "SDL_GetHapticFromID" },
            { &GetHapticID, "SDL_GetHapticID" },
            { &GetHapticName, "SDL_GetHapticName" },
            { &IsMouseHaptic, "SDL_IsMouseHaptic" },
            { &OpenHapticFromMouse, "SDL_OpenHapticFromMouse" },
            { &IsJoystickHaptic, "SDL_IsJoystickHaptic" },
            { &OpenHapticFromJoystick, "SDL_OpenHapticFromJoystick" },
            { &CloseHaptic, "SDL_CloseHaptic" },
            { &GetMaxHapticEffects, "SDL_GetMaxHapticEffects" },
            { &GetMaxHapticEffectsPlaying, "SDL_GetMaxHapticEffectsPlaying" },
            { &GetHapticFeatures, "SDL_GetHapticFeatures" },
            { &GetNumHapticAxes, "SDL_GetNumHapticAxes" },
            { &HapticEffectSupported, "SDL_HapticEffectSupported" },
            { &CreateHapticEffect, "SDL_CreateHapticEffect" },
            { &UpdateHapticEffect, "SDL_UpdateHapticEffect" },
            { &RunHapticEffect, "SDL_RunHapticEffect" },
            { &StopHapticEffect, "SDL_StopHapticEffect" },
            { &DestroyHapticEffect, "SDL_DestroyHapticEffect" },
            { &GetHapticEffectStatus, "SDL_GetHapticEffectStatus" },
            { &SetHapticGain, "SDL_SetHapticGain" },
            { &SetHapticAutocenter, "SDL_SetHapticAutocenter" },
            { &PauseHaptic, "SDL_PauseHaptic" },
            { &ResumeHaptic, "SDL_ResumeHaptic" },
            { &StopHapticEffects, "SDL_StopHapticEffects" },
            { &HapticRumbleSupported, "SDL_HapticRumbleSupported" },
            { &InitHapticRumble, "SDL_InitHapticRumble" },
            { &PlayHapticRumble, "SDL_PlayHapticRumble" },
            { &StopHapticRumble, "SDL_StopHapticRumble" },
        }
    })
}
