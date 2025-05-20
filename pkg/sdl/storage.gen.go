package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&OpenTitleStorage, lib, "SDL_OpenTitleStorage")
        purego.RegisterLibFunc(&OpenUserStorage, lib, "SDL_OpenUserStorage")
        purego.RegisterLibFunc(&OpenFileStorage, lib, "SDL_OpenFileStorage")
        purego.RegisterLibFunc(&OpenStorage, lib, "SDL_OpenStorage")
        purego.RegisterLibFunc(&CloseStorage, lib, "SDL_CloseStorage")
        purego.RegisterLibFunc(&StorageReady, lib, "SDL_StorageReady")
        purego.RegisterLibFunc(&ReadStorageFile, lib, "SDL_ReadStorageFile")
        purego.RegisterLibFunc(&WriteStorageFile, lib, "SDL_WriteStorageFile")
        purego.RegisterLibFunc(&CreateStorageDirectory, lib, "SDL_CreateStorageDirectory")
        purego.RegisterLibFunc(&EnumerateStorageDirectory, lib, "SDL_EnumerateStorageDirectory")
        purego.RegisterLibFunc(&RemoveStoragePath, lib, "SDL_RemoveStoragePath")
        purego.RegisterLibFunc(&RenameStoragePath, lib, "SDL_RenameStoragePath")
        purego.RegisterLibFunc(&CopyStorageFile, lib, "SDL_CopyStorageFile")
        purego.RegisterLibFunc(&GetStoragePathInfo, lib, "SDL_GetStoragePathInfo")
        purego.RegisterLibFunc(&GetStorageSpaceRemaining, lib, "SDL_GetStorageSpaceRemaining")
        purego.RegisterLibFunc(&GlobStorageDirectory, lib, "SDL_GlobStorageDirectory")
        
    })
}
