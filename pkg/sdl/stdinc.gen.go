package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &Malloc, "SDL_malloc" },
            { &Calloc, "SDL_calloc" },
            { &Realloc, "SDL_realloc" },
            { &Free, "SDL_free" },
            { &GetOriginalMemoryFunctions, "SDL_GetOriginalMemoryFunctions" },
            { &GetMemoryFunctions, "SDL_GetMemoryFunctions" },
            { &SetMemoryFunctions, "SDL_SetMemoryFunctions" },
            { &AlignedAlloc, "SDL_aligned_alloc" },
            { &AlignedFree, "SDL_aligned_free" },
            { &GetNumAllocations, "SDL_GetNumAllocations" },
            { &GetEnvironment, "SDL_GetEnvironment" },
            { &CreateEnvironment, "SDL_CreateEnvironment" },
            { &GetEnvironmentVariable, "SDL_GetEnvironmentVariable" },
            { &SetEnvironmentVariable, "SDL_SetEnvironmentVariable" },
            { &UnsetEnvironmentVariable, "SDL_UnsetEnvironmentVariable" },
            { &DestroyEnvironment, "SDL_DestroyEnvironment" },
            { &Getenv, "SDL_getenv" },
            { &GetenvUnsafe, "SDL_getenv_unsafe" },
            { &SetenvUnsafe, "SDL_setenv_unsafe" },
            { &UnsetenvUnsafe, "SDL_unsetenv_unsafe" },
        }
    })
}
