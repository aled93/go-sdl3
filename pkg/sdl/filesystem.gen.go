package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetBasePath, lib, "SDL_GetBasePath")
        purego.RegisterLibFunc(&GetPrefPath, lib, "SDL_GetPrefPath")
        purego.RegisterLibFunc(&GetUserFolder, lib, "SDL_GetUserFolder")
        purego.RegisterLibFunc(&CreateDirectory, lib, "SDL_CreateDirectory")
        purego.RegisterLibFunc(&EnumerateDirectory, lib, "SDL_EnumerateDirectory")
        purego.RegisterLibFunc(&RemovePath, lib, "SDL_RemovePath")
        purego.RegisterLibFunc(&RenamePath, lib, "SDL_RenamePath")
        purego.RegisterLibFunc(&CopyFile, lib, "SDL_CopyFile")
        purego.RegisterLibFunc(&GetPathInfo, lib, "SDL_GetPathInfo")
        purego.RegisterLibFunc(&GlobDirectory, lib, "SDL_GlobDirectory")
        purego.RegisterLibFunc(&GetCurrentDirectory, lib, "SDL_GetCurrentDirectory")
        
    })
}
