package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetHaptics, lib, "SDL_GetHaptics")
        purego.RegisterLibFunc(&GetHapticNameForID, lib, "SDL_GetHapticNameForID")
        purego.RegisterLibFunc(&OpenHaptic, lib, "SDL_OpenHaptic")
        purego.RegisterLibFunc(&GetHapticFromID, lib, "SDL_GetHapticFromID")
        purego.RegisterLibFunc(&GetHapticID, lib, "SDL_GetHapticID")
        purego.RegisterLibFunc(&GetHapticName, lib, "SDL_GetHapticName")
        purego.RegisterLibFunc(&IsMouseHaptic, lib, "SDL_IsMouseHaptic")
        purego.RegisterLibFunc(&OpenHapticFromMouse, lib, "SDL_OpenHapticFromMouse")
        purego.RegisterLibFunc(&IsJoystickHaptic, lib, "SDL_IsJoystickHaptic")
        purego.RegisterLibFunc(&OpenHapticFromJoystick, lib, "SDL_OpenHapticFromJoystick")
        purego.RegisterLibFunc(&CloseHaptic, lib, "SDL_CloseHaptic")
        purego.RegisterLibFunc(&GetMaxHapticEffects, lib, "SDL_GetMaxHapticEffects")
        purego.RegisterLibFunc(&GetMaxHapticEffectsPlaying, lib, "SDL_GetMaxHapticEffectsPlaying")
        purego.RegisterLibFunc(&GetHapticFeatures, lib, "SDL_GetHapticFeatures")
        purego.RegisterLibFunc(&GetNumHapticAxes, lib, "SDL_GetNumHapticAxes")
        purego.RegisterLibFunc(&HapticEffectSupported, lib, "SDL_HapticEffectSupported")
        purego.RegisterLibFunc(&CreateHapticEffect, lib, "SDL_CreateHapticEffect")
        purego.RegisterLibFunc(&UpdateHapticEffect, lib, "SDL_UpdateHapticEffect")
        purego.RegisterLibFunc(&RunHapticEffect, lib, "SDL_RunHapticEffect")
        purego.RegisterLibFunc(&StopHapticEffect, lib, "SDL_StopHapticEffect")
        purego.RegisterLibFunc(&DestroyHapticEffect, lib, "SDL_DestroyHapticEffect")
        purego.RegisterLibFunc(&GetHapticEffectStatus, lib, "SDL_GetHapticEffectStatus")
        purego.RegisterLibFunc(&SetHapticGain, lib, "SDL_SetHapticGain")
        purego.RegisterLibFunc(&SetHapticAutocenter, lib, "SDL_SetHapticAutocenter")
        purego.RegisterLibFunc(&PauseHaptic, lib, "SDL_PauseHaptic")
        purego.RegisterLibFunc(&ResumeHaptic, lib, "SDL_ResumeHaptic")
        purego.RegisterLibFunc(&StopHapticEffects, lib, "SDL_StopHapticEffects")
        purego.RegisterLibFunc(&HapticRumbleSupported, lib, "SDL_HapticRumbleSupported")
        purego.RegisterLibFunc(&InitHapticRumble, lib, "SDL_InitHapticRumble")
        purego.RegisterLibFunc(&PlayHapticRumble, lib, "SDL_PlayHapticRumble")
        purego.RegisterLibFunc(&StopHapticRumble, lib, "SDL_StopHapticRumble")
        
    })
}
