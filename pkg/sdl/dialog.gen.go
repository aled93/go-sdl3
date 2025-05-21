package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &ShowOpenFileDialog, "SDL_ShowOpenFileDialog" },
            { &ShowSaveFileDialog, "SDL_ShowSaveFileDialog" },
            { &ShowOpenFolderDialog, "SDL_ShowOpenFolderDialog" },
            { &ShowFileDialogWithProperties, "SDL_ShowFileDialogWithProperties" },
        }
    })
}
