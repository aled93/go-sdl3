package sdl

/*
 * Platform specific functions for UNIX
 */

/* this is defined in Xlib's headers, just need a simple declaration here. */
type XEvent struct{}

/**
 * A callback to be used with SDL_SetX11EventHook.
 *
 * This callback may modify the event, and should return true if the event
 * should continue to be processed, or false to prevent further processing.
 *
 * As this is processing an event directly from the X11 event loop, this
 * callback should do the minimum required work and return quickly.
 *
 * \param userdata the app-defined pointer provided to SDL_SetX11EventHook.
 * \param xevent a pointer to an Xlib XEvent union to process.
 * \returns true to let event continue on, false to drop it.
 *
 * \threadsafety This may only be called (by SDL) from the thread handling the
 *               X11 event loop.
 *
 * \since This datatype is available since SDL 3.2.0.
 *
 * \sa SDL_SetX11EventHook
 */
type X11EventHook func(userdata uintptr, xevent *XEvent) bool

/**
 * Set a callback for every X11 event.
 *
 * The callback may modify the event, and should return true if the event
 * should continue to be processed, or false to prevent further processing.
 *
 * \param callback the SDL_X11EventHook function to call.
 * \param userdata a pointer to pass to every iteration of `callback`.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetX11EventHook func(callback X11EventHook, userdata uintptr)

/* Platform specific functions for Linux*/

/**
 * Sets the UNIX nice value for a thread.
 *
 * This uses setpriority() if possible, and RealtimeKit if available.
 *
 * \param threadID the Unix thread ID to change priority of.
 * \param priority the new, Unix-specific, priority value.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetLinuxThreadPriority func(threadID int64, priority int32) bool

/**
 * Sets the priority (not nice level) and scheduling policy for a thread.
 *
 * This uses setpriority() if possible, and RealtimeKit if available.
 *
 * \param threadID the Unix thread ID to change priority of.
 * \param sdlPriority the new SDL_ThreadPriority value.
 * \param schedPolicy the new scheduling policy (SCHED_FIFO, SCHED_RR,
 *                    SCHED_OTHER, etc...).
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetLinuxThreadPriorityAndPolicy func(threadID int64, sdlPriority int32, schedPolicy int32) bool
