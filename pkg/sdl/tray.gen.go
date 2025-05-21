package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &CreateTray, "SDL_CreateTray" },
            { &SetTrayIcon, "SDL_SetTrayIcon" },
            { &SetTrayTooltip, "SDL_SetTrayTooltip" },
            { &CreateTrayMenu, "SDL_CreateTrayMenu" },
            { &CreateTraySubmenu, "SDL_CreateTraySubmenu" },
            { &GetTrayMenu, "SDL_GetTrayMenu" },
            { &GetTraySubmenu, "SDL_GetTraySubmenu" },
            { &GetTrayEntries, "SDL_GetTrayEntries" },
            { &RemoveTrayEntry, "SDL_RemoveTrayEntry" },
            { &InsertTrayEntryAt, "SDL_InsertTrayEntryAt" },
            { &SetTrayEntryLabel, "SDL_SetTrayEntryLabel" },
            { &GetTrayEntryLabel, "SDL_GetTrayEntryLabel" },
            { &SetTrayEntryChecked, "SDL_SetTrayEntryChecked" },
            { &GetTrayEntryChecked, "SDL_GetTrayEntryChecked" },
            { &SetTrayEntryEnabled, "SDL_SetTrayEntryEnabled" },
            { &GetTrayEntryEnabled, "SDL_GetTrayEntryEnabled" },
            { &SetTrayEntryCallback, "SDL_SetTrayEntryCallback" },
            { &ClickTrayEntry, "SDL_ClickTrayEntry" },
            { &DestroyTray, "SDL_DestroyTray" },
            { &GetTrayEntryParent, "SDL_GetTrayEntryParent" },
            { &GetTrayMenuParentEntry, "SDL_GetTrayMenuParentEntry" },
            { &GetTrayMenuParentTray, "SDL_GetTrayMenuParentTray" },
            { &UpdateTrays, "SDL_UpdateTrays" },
        }
    })
}
