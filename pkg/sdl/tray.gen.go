package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&CreateTray, lib, "SDL_CreateTray")
        purego.RegisterLibFunc(&SetTrayIcon, lib, "SDL_SetTrayIcon")
        purego.RegisterLibFunc(&SetTrayTooltip, lib, "SDL_SetTrayTooltip")
        purego.RegisterLibFunc(&CreateTrayMenu, lib, "SDL_CreateTrayMenu")
        purego.RegisterLibFunc(&CreateTraySubmenu, lib, "SDL_CreateTraySubmenu")
        purego.RegisterLibFunc(&GetTrayMenu, lib, "SDL_GetTrayMenu")
        purego.RegisterLibFunc(&GetTraySubmenu, lib, "SDL_GetTraySubmenu")
        purego.RegisterLibFunc(&GetTrayEntries, lib, "SDL_GetTrayEntries")
        purego.RegisterLibFunc(&RemoveTrayEntry, lib, "SDL_RemoveTrayEntry")
        purego.RegisterLibFunc(&InsertTrayEntryAt, lib, "SDL_InsertTrayEntryAt")
        purego.RegisterLibFunc(&SetTrayEntryLabel, lib, "SDL_SetTrayEntryLabel")
        purego.RegisterLibFunc(&GetTrayEntryLabel, lib, "SDL_GetTrayEntryLabel")
        purego.RegisterLibFunc(&SetTrayEntryChecked, lib, "SDL_SetTrayEntryChecked")
        purego.RegisterLibFunc(&GetTrayEntryChecked, lib, "SDL_GetTrayEntryChecked")
        purego.RegisterLibFunc(&SetTrayEntryEnabled, lib, "SDL_SetTrayEntryEnabled")
        purego.RegisterLibFunc(&GetTrayEntryEnabled, lib, "SDL_GetTrayEntryEnabled")
        purego.RegisterLibFunc(&SetTrayEntryCallback, lib, "SDL_SetTrayEntryCallback")
        purego.RegisterLibFunc(&ClickTrayEntry, lib, "SDL_ClickTrayEntry")
        purego.RegisterLibFunc(&DestroyTray, lib, "SDL_DestroyTray")
        purego.RegisterLibFunc(&GetTrayEntryParent, lib, "SDL_GetTrayEntryParent")
        purego.RegisterLibFunc(&GetTrayMenuParentEntry, lib, "SDL_GetTrayMenuParentEntry")
        purego.RegisterLibFunc(&GetTrayMenuParentTray, lib, "SDL_GetTrayMenuParentTray")
        purego.RegisterLibFunc(&UpdateTrays, lib, "SDL_UpdateTrays")
        
    })
}
