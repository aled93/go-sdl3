package sdl

/*
 * Platform specific functions for iOS
 */

/**
 * The prototype for an Apple iOS animation callback.
 *
 * This datatype is only useful on Apple iOS.
 *
 * After passing a function pointer of this type to
 * SDL_SetiOSAnimationCallback, the system will call that function pointer at
 * a regular interval.
 *
 * \param userdata what was passed as `callbackParam` to
 *                 SDL_SetiOSAnimationCallback as `callbackParam`.
 *
 * \since This datatype is available since SDL 3.2.0.
 *
 * \sa SDL_SetiOSAnimationCallback
 */
type IOSAnimationCallback func(userdata uintptr)

/**
 * Use this function to set the animation callback on Apple iOS.
 *
 * The function prototype for `callback` is:
 *
 * ```c
 * void callback(void *callbackParam);
 * ```
 *
 * Where its parameter, `callbackParam`, is what was passed as `callbackParam`
 * to SDL_SetiOSAnimationCallback().
 *
 * This function is only available on Apple iOS.
 *
 * For more information see:
 *
 * https://wiki.libsdl.org/SDL3/README-ios
 *
 * Note that if you use the "main callbacks" instead of a standard C `main`
 * function, you don't have to use this API, as SDL will manage this for you.
 *
 * Details on main callbacks are here:
 *
 * https://wiki.libsdl.org/SDL3/README-main-functions
 *
 * \param window the window for which the animation callback should be set.
 * \param interval the number of frames after which **callback** will be
 *                 called.
 * \param callback the function to call for every frame.
 * \param callbackParam a pointer that is passed to `callback`.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SetiOSEventPump
 */
//go:sdl3extern
var SetiOSAnimationCallback func(window *Window, interval int, callback IOSAnimationCallback, callbackParam uintptr) bool

/**
 * Use this function to enable or disable the SDL event pump on Apple iOS.
 *
 * This function is only available on Apple iOS.
 *
 * \param enabled true to enable the event pump, false to disable it.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SetiOSAnimationCallback
 */
//go:sdl3extern
var SetiOSEventPump func(enabled bool)

/**
* Let iOS apps with external event handling report
* onApplicationDidChangeStatusBarOrientation.
*
* This functions allows iOS apps that have their own event handling to hook
* into SDL to generate SDL events. This maps directly to an iOS-specific
* event, but since it doesn't do anything iOS-specific internally, it is
* available on all platforms, in case it might be useful for some specific
* paradigm. Most apps do not need to use this directly; SDL's internal event
* code will handle all this for windows created by SDL_CreateWindow!
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var OnApplicationDidChangeStatusBarOrientation func()
