package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&ShowOpenFileDialog, lib, "SDL_ShowOpenFileDialog")
        purego.RegisterLibFunc(&ShowSaveFileDialog, lib, "SDL_ShowSaveFileDialog")
        purego.RegisterLibFunc(&ShowOpenFolderDialog, lib, "SDL_ShowOpenFolderDialog")
        purego.RegisterLibFunc(&ShowFileDialogWithProperties, lib, "SDL_ShowFileDialogWithProperties")
        
    })
}
