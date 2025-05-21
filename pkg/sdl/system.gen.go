package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &IsTablet, "SDL_IsTablet" },
            { &IsTV, "SDL_IsTV" },
            { &GetSandbox, "SDL_GetSandbox" },
            { &OnApplicationWillTerminate, "SDL_OnApplicationWillTerminate" },
            { &OnApplicationDidReceiveMemoryWarning, "SDL_OnApplicationDidReceiveMemoryWarning" },
            { &OnApplicationWillEnterBackground, "SDL_OnApplicationWillEnterBackground" },
            { &OnApplicationDidEnterBackground, "SDL_OnApplicationDidEnterBackground" },
            { &OnApplicationWillEnterForeground, "SDL_OnApplicationWillEnterForeground" },
            { &OnApplicationDidEnterForeground, "SDL_OnApplicationDidEnterForeground" },
        }
    })
}
