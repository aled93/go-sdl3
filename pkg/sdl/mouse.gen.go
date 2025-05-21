package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &HasMouse, "SDL_HasMouse" },
            { &GetMice, "SDL_GetMice" },
            { &GetMouseNameForID, "SDL_GetMouseNameForID" },
            { &GetMouseFocus, "SDL_GetMouseFocus" },
            { &GetMouseState, "SDL_GetMouseState" },
            { &GetGlobalMouseState, "SDL_GetGlobalMouseState" },
            { &GetRelativeMouseState, "SDL_GetRelativeMouseState" },
            { &WarpMouseInWindow, "SDL_WarpMouseInWindow" },
            { &WarpMouseGlobal, "SDL_WarpMouseGlobal" },
            { &SetRelativeMouseTransform, "SDL_SetRelativeMouseTransform" },
            { &SetWindowRelativeMouseMode, "SDL_SetWindowRelativeMouseMode" },
            { &GetWindowRelativeMouseMode, "SDL_GetWindowRelativeMouseMode" },
            { &CaptureMouse, "SDL_CaptureMouse" },
            { &CreateCursor, "SDL_CreateCursor" },
            { &CreateColorCursor, "SDL_CreateColorCursor" },
            { &CreateSystemCursor, "SDL_CreateSystemCursor" },
            { &SetCursor, "SDL_SetCursor" },
            { &GetCursor, "SDL_GetCursor" },
            { &GetDefaultCursor, "SDL_GetDefaultCursor" },
            { &DestroyCursor, "SDL_DestroyCursor" },
            { &ShowCursor, "SDL_ShowCursor" },
            { &HideCursor, "SDL_HideCursor" },
            { &CursorVisible, "SDL_CursorVisible" },
        }
    })
}
