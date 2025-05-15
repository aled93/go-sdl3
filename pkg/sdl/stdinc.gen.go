package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&Malloc, lib, "SDL_Malloc")
        purego.RegisterLibFunc(&Calloc, lib, "SDL_Calloc")
        purego.RegisterLibFunc(&Realloc, lib, "SDL_Realloc")
        purego.RegisterLibFunc(&Free, lib, "SDL_Free")
        purego.RegisterLibFunc(&GetOriginalMemoryFunctions, lib, "SDL_GetOriginalMemoryFunctions")
        purego.RegisterLibFunc(&GetMemoryFunctions, lib, "SDL_GetMemoryFunctions")
        purego.RegisterLibFunc(&SetMemoryFunctions, lib, "SDL_SetMemoryFunctions")
        purego.RegisterLibFunc(&AlignedAlloc, lib, "SDL_AlignedAlloc")
        purego.RegisterLibFunc(&AlignedFree, lib, "SDL_AlignedFree")
        purego.RegisterLibFunc(&GetNumAllocations, lib, "SDL_GetNumAllocations")
        purego.RegisterLibFunc(&GetEnvironment, lib, "SDL_GetEnvironment")
        purego.RegisterLibFunc(&CreateEnvironment, lib, "SDL_CreateEnvironment")
        purego.RegisterLibFunc(&GetEnvironmentVariable, lib, "SDL_GetEnvironmentVariable")
        purego.RegisterLibFunc(&SetEnvironmentVariable, lib, "SDL_SetEnvironmentVariable")
        purego.RegisterLibFunc(&UnsetEnvironmentVariable, lib, "SDL_UnsetEnvironmentVariable")
        purego.RegisterLibFunc(&DestroyEnvironment, lib, "SDL_DestroyEnvironment")
        purego.RegisterLibFunc(&Getenv, lib, "SDL_Getenv")
        purego.RegisterLibFunc(&GetenvUnsafe, lib, "SDL_GetenvUnsafe")
        purego.RegisterLibFunc(&SetenvUnsafe, lib, "SDL_SetenvUnsafe")
        purego.RegisterLibFunc(&UnsetenvUnsafe, lib, "SDL_UnsetenvUnsafe")
        
    })
}
