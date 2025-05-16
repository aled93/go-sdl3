package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&Malloc, lib, "SDL_malloc")
        purego.RegisterLibFunc(&Calloc, lib, "SDL_calloc")
        purego.RegisterLibFunc(&Realloc, lib, "SDL_realloc")
        purego.RegisterLibFunc(&Free, lib, "SDL_free")
        purego.RegisterLibFunc(&GetOriginalMemoryFunctions, lib, "SDL_GetOriginalMemoryFunctions")
        purego.RegisterLibFunc(&GetMemoryFunctions, lib, "SDL_GetMemoryFunctions")
        purego.RegisterLibFunc(&SetMemoryFunctions, lib, "SDL_SetMemoryFunctions")
        purego.RegisterLibFunc(&AlignedAlloc, lib, "SDL_aligned_alloc")
        purego.RegisterLibFunc(&AlignedFree, lib, "SDL_aligned_free")
        purego.RegisterLibFunc(&GetNumAllocations, lib, "SDL_GetNumAllocations")
        purego.RegisterLibFunc(&GetEnvironment, lib, "SDL_GetEnvironment")
        purego.RegisterLibFunc(&CreateEnvironment, lib, "SDL_CreateEnvironment")
        purego.RegisterLibFunc(&GetEnvironmentVariable, lib, "SDL_GetEnvironmentVariable")
        purego.RegisterLibFunc(&SetEnvironmentVariable, lib, "SDL_SetEnvironmentVariable")
        purego.RegisterLibFunc(&UnsetEnvironmentVariable, lib, "SDL_UnsetEnvironmentVariable")
        purego.RegisterLibFunc(&DestroyEnvironment, lib, "SDL_DestroyEnvironment")
        purego.RegisterLibFunc(&Getenv, lib, "SDL_getenv")
        purego.RegisterLibFunc(&GetenvUnsafe, lib, "SDL_getenv_unsafe")
        purego.RegisterLibFunc(&SetenvUnsafe, lib, "SDL_setenv_unsafe")
        purego.RegisterLibFunc(&UnsetenvUnsafe, lib, "SDL_unsetenv_unsafe")
        
    })
}
