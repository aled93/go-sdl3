package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &OpenTitleStorage, "SDL_OpenTitleStorage" },
            { &OpenUserStorage, "SDL_OpenUserStorage" },
            { &OpenFileStorage, "SDL_OpenFileStorage" },
            { &OpenStorage, "SDL_OpenStorage" },
            { &CloseStorage, "SDL_CloseStorage" },
            { &StorageReady, "SDL_StorageReady" },
            { &ReadStorageFile, "SDL_ReadStorageFile" },
            { &WriteStorageFile, "SDL_WriteStorageFile" },
            { &CreateStorageDirectory, "SDL_CreateStorageDirectory" },
            { &EnumerateStorageDirectory, "SDL_EnumerateStorageDirectory" },
            { &RemoveStoragePath, "SDL_RemoveStoragePath" },
            { &RenameStoragePath, "SDL_RenameStoragePath" },
            { &CopyStorageFile, "SDL_CopyStorageFile" },
            { &GetStoragePathInfo, "SDL_GetStoragePathInfo" },
            { &GetStorageSpaceRemaining, "SDL_GetStorageSpaceRemaining" },
            { &GlobStorageDirectory, "SDL_GlobStorageDirectory" },
        }
    })
}
