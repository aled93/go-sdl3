package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&HasMouse, lib, "SDL_HasMouse")
        purego.RegisterLibFunc(&GetMice, lib, "SDL_GetMice")
        purego.RegisterLibFunc(&GetMouseNameForID, lib, "SDL_GetMouseNameForID")
        purego.RegisterLibFunc(&GetMouseFocus, lib, "SDL_GetMouseFocus")
        purego.RegisterLibFunc(&GetMouseState, lib, "SDL_GetMouseState")
        purego.RegisterLibFunc(&GetGlobalMouseState, lib, "SDL_GetGlobalMouseState")
        purego.RegisterLibFunc(&GetRelativeMouseState, lib, "SDL_GetRelativeMouseState")
        purego.RegisterLibFunc(&WarpMouseInWindow, lib, "SDL_WarpMouseInWindow")
        purego.RegisterLibFunc(&WarpMouseGlobal, lib, "SDL_WarpMouseGlobal")
        purego.RegisterLibFunc(&SetRelativeMouseTransform, lib, "SDL_SetRelativeMouseTransform")
        purego.RegisterLibFunc(&SetWindowRelativeMouseMode, lib, "SDL_SetWindowRelativeMouseMode")
        purego.RegisterLibFunc(&GetWindowRelativeMouseMode, lib, "SDL_GetWindowRelativeMouseMode")
        purego.RegisterLibFunc(&CaptureMouse, lib, "SDL_CaptureMouse")
        purego.RegisterLibFunc(&CreateCursor, lib, "SDL_CreateCursor")
        purego.RegisterLibFunc(&CreateColorCursor, lib, "SDL_CreateColorCursor")
        purego.RegisterLibFunc(&CreateSystemCursor, lib, "SDL_CreateSystemCursor")
        purego.RegisterLibFunc(&SetCursor, lib, "SDL_SetCursor")
        purego.RegisterLibFunc(&GetCursor, lib, "SDL_GetCursor")
        purego.RegisterLibFunc(&GetDefaultCursor, lib, "SDL_GetDefaultCursor")
        purego.RegisterLibFunc(&DestroyCursor, lib, "SDL_DestroyCursor")
        purego.RegisterLibFunc(&ShowCursor, lib, "SDL_ShowCursor")
        purego.RegisterLibFunc(&HideCursor, lib, "SDL_HideCursor")
        purego.RegisterLibFunc(&CursorVisible, lib, "SDL_CursorVisible")
        
    })
}
