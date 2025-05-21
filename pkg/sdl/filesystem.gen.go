package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetBasePath, "SDL_GetBasePath" },
            { &GetPrefPath, "SDL_GetPrefPath" },
            { &GetUserFolder, "SDL_GetUserFolder" },
            { &CreateDirectory, "SDL_CreateDirectory" },
            { &EnumerateDirectory, "SDL_EnumerateDirectory" },
            { &RemovePath, "SDL_RemovePath" },
            { &RenamePath, "SDL_RenamePath" },
            { &CopyFile, "SDL_CopyFile" },
            { &GetPathInfo, "SDL_GetPathInfo" },
            { &GlobDirectory, "SDL_GlobDirectory" },
            { &GetCurrentDirectory, "SDL_GetCurrentDirectory" },
        }
    })
}
