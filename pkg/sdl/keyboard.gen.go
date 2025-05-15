package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&HasKeyboard, lib, "SDL_HasKeyboard")
        purego.RegisterLibFunc(&GetKeyboards, lib, "SDL_GetKeyboards")
        purego.RegisterLibFunc(&GetKeyboardNameForID, lib, "SDL_GetKeyboardNameForID")
        purego.RegisterLibFunc(&GetKeyboardFocus, lib, "SDL_GetKeyboardFocus")
        purego.RegisterLibFunc(&GetKeyboardState, lib, "SDL_GetKeyboardState")
        purego.RegisterLibFunc(&ResetKeyboard, lib, "SDL_ResetKeyboard")
        purego.RegisterLibFunc(&GetModState, lib, "SDL_GetModState")
        purego.RegisterLibFunc(&SetModState, lib, "SDL_SetModState")
        purego.RegisterLibFunc(&GetKeyFromScancode, lib, "SDL_GetKeyFromScancode")
        purego.RegisterLibFunc(&GetScancodeFromKey, lib, "SDL_GetScancodeFromKey")
        purego.RegisterLibFunc(&SetScancodeName, lib, "SDL_SetScancodeName")
        purego.RegisterLibFunc(&GetScancodeName, lib, "SDL_GetScancodeName")
        purego.RegisterLibFunc(&GetScancodeFromName, lib, "SDL_GetScancodeFromName")
        purego.RegisterLibFunc(&GetKeyName, lib, "SDL_GetKeyName")
        purego.RegisterLibFunc(&GetKeyFromName, lib, "SDL_GetKeyFromName")
        purego.RegisterLibFunc(&StartTextInput, lib, "SDL_StartTextInput")
        purego.RegisterLibFunc(&StartTextInputWithProperties, lib, "SDL_StartTextInputWithProperties")
        purego.RegisterLibFunc(&TextInputActive, lib, "SDL_TextInputActive")
        purego.RegisterLibFunc(&StopTextInput, lib, "SDL_StopTextInput")
        purego.RegisterLibFunc(&ClearComposition, lib, "SDL_ClearComposition")
        purego.RegisterLibFunc(&SetTextInputArea, lib, "SDL_SetTextInputArea")
        purego.RegisterLibFunc(&GetTextInputArea, lib, "SDL_GetTextInputArea")
        purego.RegisterLibFunc(&HasScreenKeyboardSupport, lib, "SDL_HasScreenKeyboardSupport")
        purego.RegisterLibFunc(&ScreenKeyboardShown, lib, "SDL_ScreenKeyboardShown")
        
    })
}
