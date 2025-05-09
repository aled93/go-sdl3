//go:build GDK

package sdl

/*
* Functions used only by GDK
 */
type XTaskQueueObject struct{}
type XTaskQueueHandle *XTaskQueueObject
type XUser struct{}
type XUserHandle *XUser

/**
* Gets a reference to the global async task queue handle for GDK,
* initializing if needed.
*
* Once you are done with the task queue, you should call
* XTaskQueueCloseHandle to reduce the reference count to avoid a resource
* leak.
*
* \param outTaskQueue a pointer to be filled in with task queue handle.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGDKTaskQueue func(outTaskQueue *XTaskQueueHandle) bool

/**
* Gets a reference to the default user handle for GDK.
*
* This is effectively a synchronous version of XUserAddAsync, which always
* prefers the default user and allows a sign-in UI.
*
* \param outUserHandle a pointer to be filled in with the default user
*                      handle.
* \returns true if success or false on failure; call SDL_GetError() for more
*          information.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGDKDefaultUser func(outUserHandle *XUserHandle) bool
