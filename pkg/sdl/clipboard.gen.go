package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&SetClipboardText, lib, "SDL_SetClipboardText")
        purego.RegisterLibFunc(&GetClipboardText, lib, "SDL_GetClipboardText")
        purego.RegisterLibFunc(&HasClipboardText, lib, "SDL_HasClipboardText")
        purego.RegisterLibFunc(&SetPrimarySelectionText, lib, "SDL_SetPrimarySelectionText")
        purego.RegisterLibFunc(&GetPrimarySelectionText, lib, "SDL_GetPrimarySelectionText")
        purego.RegisterLibFunc(&HasPrimarySelectionText, lib, "SDL_HasPrimarySelectionText")
        purego.RegisterLibFunc(&SetClipboardData, lib, "SDL_SetClipboardData")
        purego.RegisterLibFunc(&ClearClipboardData, lib, "SDL_ClearClipboardData")
        purego.RegisterLibFunc(&GetClipboardData, lib, "SDL_GetClipboardData")
        purego.RegisterLibFunc(&HasClipboardData, lib, "SDL_HasClipboardData")
        purego.RegisterLibFunc(&GetClipboardMimeTypes, lib, "SDL_GetClipboardMimeTypes")
        
    })
}
