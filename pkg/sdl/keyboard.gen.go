package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &HasKeyboard, "SDL_HasKeyboard" },
            { &GetKeyboards, "SDL_GetKeyboards" },
            { &GetKeyboardNameForID, "SDL_GetKeyboardNameForID" },
            { &GetKeyboardFocus, "SDL_GetKeyboardFocus" },
            { &GetKeyboardState, "SDL_GetKeyboardState" },
            { &ResetKeyboard, "SDL_ResetKeyboard" },
            { &GetModState, "SDL_GetModState" },
            { &SetModState, "SDL_SetModState" },
            { &GetKeyFromScancode, "SDL_GetKeyFromScancode" },
            { &GetScancodeFromKey, "SDL_GetScancodeFromKey" },
            { &SetScancodeName, "SDL_SetScancodeName" },
            { &GetScancodeName, "SDL_GetScancodeName" },
            { &GetScancodeFromName, "SDL_GetScancodeFromName" },
            { &GetKeyName, "SDL_GetKeyName" },
            { &GetKeyFromName, "SDL_GetKeyFromName" },
            { &StartTextInput, "SDL_StartTextInput" },
            { &StartTextInputWithProperties, "SDL_StartTextInputWithProperties" },
            { &TextInputActive, "SDL_TextInputActive" },
            { &StopTextInput, "SDL_StopTextInput" },
            { &ClearComposition, "SDL_ClearComposition" },
            { &SetTextInputArea, "SDL_SetTextInputArea" },
            { &GetTextInputArea, "SDL_GetTextInputArea" },
            { &HasScreenKeyboardSupport, "SDL_HasScreenKeyboardSupport" },
            { &ScreenKeyboardShown, "SDL_ScreenKeyboardShown" },
        }
    })
}
