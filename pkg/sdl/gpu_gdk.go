//go:build GDK

package sdl

/**
 * Call this to suspend GPU operation on Xbox when you receive the
 * SDL_EVENT_DID_ENTER_BACKGROUND event.
 *
 * Do NOT call any SDL_GPU functions after calling this function! This must
 * also be called before calling SDL_GDKSuspendComplete.
 *
 * \param device a GPU context.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddEventWatch
 */
//go:sdl3extern
var GDKSuspendGPU func(device GPUDevice)

/**
 * Call this to resume GPU operation on Xbox when you receive the
 * SDL_EVENT_WILL_ENTER_FOREGROUND event.
 *
 * When resuming, this function MUST be called before calling any other
 * SDL_GPU functions.
 *
 * \param device a GPU context.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddEventWatch
 */
//go:sdl3extern
var GDKResumeGPU func(device GPUDevice)
