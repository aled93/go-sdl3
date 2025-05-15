package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&Init, lib, "SDL_Init")
        purego.RegisterLibFunc(&InitSubSystem, lib, "SDL_InitSubSystem")
        purego.RegisterLibFunc(&QuitSubSystem, lib, "SDL_QuitSubSystem")
        purego.RegisterLibFunc(&WasInit, lib, "SDL_WasInit")
        purego.RegisterLibFunc(&Quit, lib, "SDL_Quit")
        purego.RegisterLibFunc(&IsMainThread, lib, "SDL_IsMainThread")
        purego.RegisterLibFunc(&RunOnMainThread, lib, "SDL_RunOnMainThread")
        purego.RegisterLibFunc(&SetAppMetadata, lib, "SDL_SetAppMetadata")
        purego.RegisterLibFunc(&SetAppMetadataProperty, lib, "SDL_SetAppMetadataProperty")
        purego.RegisterLibFunc(&GetAppMetadataProperty, lib, "SDL_GetAppMetadataProperty")
        
    })
}
