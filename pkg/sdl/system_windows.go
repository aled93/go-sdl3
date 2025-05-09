package sdl

type MSG struct{}

/**
 * A callback to be used with SDL_SetWindowsMessageHook.
 *
 * This callback may modify the message, and should return true if the message
 * should continue to be processed, or false to prevent further processing.
 *
 * As this is processing a message directly from the Windows event loop, this
 * callback should do the minimum required work and return quickly.
 *
 * \param userdata the app-defined pointer provided to
 *                 SDL_SetWindowsMessageHook.
 * \param msg a pointer to a Win32 event structure to process.
 * \returns true to let event continue on, false to drop it.
 *
 * \threadsafety This may only be called (by SDL) from the thread handling the
 *               Windows event loop.
 *
 * \since This datatype is available since SDL 3.2.0.
 *
 * \sa SDL_SetWindowsMessageHook
 * \sa SDL_HINT_WINDOWS_ENABLE_MESSAGELOOP
 */
type WindowsMessageHook func(userdata uintptr, msg *MSG) bool

/**
 * Set a callback for every Windows message, run before TranslateMessage().
 *
 * The callback may modify the message, and should return true if the message
 * should continue to be processed, or false to prevent further processing.
 *
 * \param callback the SDL_WindowsMessageHook function to call.
 * \param userdata a pointer to pass to every iteration of `callback`.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_WindowsMessageHook
 * \sa SDL_HINT_WINDOWS_ENABLE_MESSAGELOOP
 */
//go:sdl3extern
var SetWindowsMessageHook func(callback WindowsMessageHook, userdata uintptr)
