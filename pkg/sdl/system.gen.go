package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&IsTablet, lib, "SDL_IsTablet")
        purego.RegisterLibFunc(&IsTV, lib, "SDL_IsTV")
        purego.RegisterLibFunc(&GetSandbox, lib, "SDL_GetSandbox")
        purego.RegisterLibFunc(&OnApplicationWillTerminate, lib, "SDL_OnApplicationWillTerminate")
        purego.RegisterLibFunc(&OnApplicationDidReceiveMemoryWarning, lib, "SDL_OnApplicationDidReceiveMemoryWarning")
        purego.RegisterLibFunc(&OnApplicationWillEnterBackground, lib, "SDL_OnApplicationWillEnterBackground")
        purego.RegisterLibFunc(&OnApplicationDidEnterBackground, lib, "SDL_OnApplicationDidEnterBackground")
        purego.RegisterLibFunc(&OnApplicationWillEnterForeground, lib, "SDL_OnApplicationWillEnterForeground")
        purego.RegisterLibFunc(&OnApplicationDidEnterForeground, lib, "SDL_OnApplicationDidEnterForeground")
        
    })
}
