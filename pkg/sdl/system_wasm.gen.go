//go:build wasm

package sdl



//go:wasmimport sdl3 SDL_IsTablet
func __SDL_IsTablet() int32

//go:wasmimport sdl3 SDL_IsTV
func __SDL_IsTV() int32

//go:wasmimport sdl3 SDL_GetSandbox
func __SDL_GetSandbox() int32

//go:wasmimport sdl3 SDL_OnApplicationWillTerminate
func __SDL_OnApplicationWillTerminate()

//go:wasmimport sdl3 SDL_OnApplicationDidReceiveMemoryWarning
func __SDL_OnApplicationDidReceiveMemoryWarning()

//go:wasmimport sdl3 SDL_OnApplicationWillEnterBackground
func __SDL_OnApplicationWillEnterBackground()

//go:wasmimport sdl3 SDL_OnApplicationDidEnterBackground
func __SDL_OnApplicationDidEnterBackground()

//go:wasmimport sdl3 SDL_OnApplicationWillEnterForeground
func __SDL_OnApplicationWillEnterForeground()

//go:wasmimport sdl3 SDL_OnApplicationDidEnterForeground
func __SDL_OnApplicationDidEnterForeground()



func __gowrap__SDL_IsTablet() (__result bool) {
	__result = (bool)(__SDL_IsTablet() != 0)
	return
}

func __gowrap__SDL_IsTV() (__result bool) {
	__result = (bool)(__SDL_IsTV() != 0)
	return
}

func __gowrap__SDL_GetSandbox() (__result Sandbox) {
	__result = (Sandbox)(__SDL_GetSandbox())
	return
}

func __gowrap__SDL_OnApplicationWillTerminate() {
	__SDL_OnApplicationWillTerminate()
}

func __gowrap__SDL_OnApplicationDidReceiveMemoryWarning() {
	__SDL_OnApplicationDidReceiveMemoryWarning()
}

func __gowrap__SDL_OnApplicationWillEnterBackground() {
	__SDL_OnApplicationWillEnterBackground()
}

func __gowrap__SDL_OnApplicationDidEnterBackground() {
	__SDL_OnApplicationDidEnterBackground()
}

func __gowrap__SDL_OnApplicationWillEnterForeground() {
	__SDL_OnApplicationWillEnterForeground()
}

func __gowrap__SDL_OnApplicationDidEnterForeground() {
	__SDL_OnApplicationDidEnterForeground()
}

 
func init() {
	IsTablet = __gowrap__SDL_IsTablet
	IsTV = __gowrap__SDL_IsTV
	GetSandbox = __gowrap__SDL_GetSandbox
	OnApplicationWillTerminate = __gowrap__SDL_OnApplicationWillTerminate
	OnApplicationDidReceiveMemoryWarning = __gowrap__SDL_OnApplicationDidReceiveMemoryWarning
	OnApplicationWillEnterBackground = __gowrap__SDL_OnApplicationWillEnterBackground
	OnApplicationDidEnterBackground = __gowrap__SDL_OnApplicationDidEnterBackground
	OnApplicationWillEnterForeground = __gowrap__SDL_OnApplicationWillEnterForeground
	OnApplicationDidEnterForeground = __gowrap__SDL_OnApplicationDidEnterForeground
}
