/*
  Simple DirectMedia Layer
  Copyright (C) 1997-2025 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/

/*
  Language bindings and additional adaptations by AleD, 2025
  This file contains modified code and/or translated interfaces.
  Original SDL3 source: https://github.com/libsdl-org/SDL
*/

/**
 * # CategoryMutex
 *
 * SDL offers several thread synchronization primitives. This document can't
 * cover the complicated topic of thread safety, but reading up on what each
 * of these primitives are, why they are useful, and how to correctly use them
 * is vital to writing correct and safe multithreaded programs.
 *
 * - Mutexes: SDL_CreateMutex()
 * - Read/Write locks: SDL_CreateRWLock()
 * - Semaphores: SDL_CreateSemaphore()
 * - Condition variables: SDL_CreateCondition()
 *
 * SDL also offers a datatype, SDL_InitState, which can be used to make sure
 * only one thread initializes/deinitializes some resource that several
 * threads might try to use for the first time simultaneously.
 */
package sdl

/**
 *  \name Thread-safe initialization state functions
 */
/* @{ */

/**
 * The current status of an SDL_InitState structure.
 *
 * \since This enum is available since SDL 3.2.0.
 */
type InitStatus int32

const (
	InitStatusUninitialized InitStatus = iota
	InitStatusInitializing
	InitStatusInitialized
	InitStatusUninitializing
)

/**
 * A structure used for thread-safe initialization and shutdown.
 *
 * Here is an example of using this:
 *
 * ```c
 *    static SDL_InitState init;
 *
 *    bool InitSystem(void)
 *    {
 *        if (!SDL_ShouldInit(&init)) {
 *            // The system is initialized
 *            return true;
 *        }
 *
 *        // At this point, you should not leave this function without calling SDL_SetInitialized()
 *
 *        bool initialized = DoInitTasks();
 *        SDL_SetInitialized(&init, initialized);
 *        return initialized;
 *    }
 *
 *    bool UseSubsystem(void)
 *    {
 *        if (SDL_ShouldInit(&init)) {
 *            // Error, the subsystem isn't initialized
 *            SDL_SetInitialized(&init, false);
 *            return false;
 *        }
 *
 *        // Do work using the initialized subsystem
 *
 *        return true;
 *    }
 *
 *    void QuitSystem(void)
 *    {
 *        if (!SDL_ShouldQuit(&init)) {
 *            // The system is not initialized
 *            return;
 *        }
 *
 *        // At this point, you should not leave this function without calling SDL_SetInitialized()
 *
 *        DoQuitTasks();
 *        SDL_SetInitialized(&init, false);
 *    }
 * ```
 *
 * Note that this doesn't protect any resources created during initialization,
 * or guarantee that nobody is using those resources during cleanup. You
 * should use other mechanisms to protect those, if that's a concern for your
 * code.
 *
 * \since This struct is available since SDL 3.2.0.
 */
type InitState struct {
	Status   int32  // AtomicInt
	Thread   uint64 // ThreadID
	Reserved uintptr
}

/**
 * Return whether initialization should be done.
 *
 * This function checks the passed in state and if initialization should be
 * done, sets the status to `SDL_INIT_STATUS_INITIALIZING` and returns true.
 * If another thread is already modifying this state, it will wait until
 * that's done before returning.
 *
 * If this function returns true, the calling code must call
 * SDL_SetInitialized() to complete the initialization.
 *
 * \param state the initialization state to check.
 * \returns true if initialization needs to be done, false otherwise.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SetInitialized
 * \sa SDL_ShouldQuit
 */
//go:sdl3extern
var ShouldInit func(state *InitState) bool

/**
 * Return whether cleanup should be done.
 *
 * This function checks the passed in state and if cleanup should be done,
 * sets the status to `SDL_INIT_STATUS_UNINITIALIZING` and returns true.
 *
 * If this function returns true, the calling code must call
 * SDL_SetInitialized() to complete the cleanup.
 *
 * \param state the initialization state to check.
 * \returns true if cleanup needs to be done, false otherwise.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SetInitialized
 * \sa SDL_ShouldInit
 */
//go:sdl3extern
var ShouldQuit func(state *InitState) bool

/**
 * Finish an initialization state transition.
 *
 * This function sets the status of the passed in state to
 * `SDL_INIT_STATUS_INITIALIZED` or `SDL_INIT_STATUS_UNINITIALIZED` and allows
 * any threads waiting for the status to proceed.
 *
 * \param state the initialization state to check.
 * \param initialized the new initialization state.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_ShouldInit
 * \sa SDL_ShouldQuit
 */
//go:sdl3extern
var SetInitialized func(state *InitState, initialized bool)

/* @} */ /* Thread-safe initialization state functions */
