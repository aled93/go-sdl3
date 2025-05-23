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
 * # CategoryTimer
 *
 * SDL provides time management functionality. It is useful for dealing with
 * (usually) small durations of time.
 *
 * This is not to be confused with _calendar time_ management, which is
 * provided by [CategoryTime](CategoryTime).
 *
 * This category covers measuring time elapsed (SDL_GetTicks(),
 * SDL_GetPerformanceCounter()), putting a thread to sleep for a certain
 * amount of time (SDL_Delay(), SDL_DelayNS(), SDL_DelayPrecise()), and firing
 * a callback function after a certain amount of time has elasped
 * (SDL_AddTimer(), etc).
 *
 * There are also useful macros to convert between time units, like
 * SDL_SECONDS_TO_NS() and such.
 */
package sdl

/* SDL time constants */

/**
 * Number of milliseconds in a second.
 *
 * This is always 1000.
 *
 * \since This macro is available since SDL 3.2.0.
 */
const MS_PER_SECOND = 1000

/**
 * Number of microseconds in a second.
 *
 * This is always 1000000.
 *
 * \since This macro is available since SDL 3.2.0.
 */
const US_PER_SECOND = 1000000

/**
 * Number of nanoseconds in a second.
 *
 * This is always 1000000000.
 *
 * \since This macro is available since SDL 3.2.0.
 */
const NS_PER_SECOND = 1000000000

/**
 * Number of nanoseconds in a millisecond.
 *
 * This is always 1000000.
 *
 * \since This macro is available since SDL 3.2.0.
 */
const NS_PER_MS = 1000000

/**
 * Number of nanoseconds in a microsecond.
 *
 * This is always 1000.
 *
 * \since This macro is available since SDL 3.2.0.
 */
const NS_PER_US = 1000

/**
 * Convert seconds to nanoseconds.
 *
 * This only converts whole numbers, not fractional seconds.
 *
 * \param S the number of seconds to convert.
 * \returns S, expressed in nanoseconds.
 *
 * \threadsafety It is safe to call this macro from any thread.
 *
 * \since This macro is available since SDL 3.2.0.
 */
func SECONDS_TO_NS(S uint32) uint64 {
	return uint64(S) * NS_PER_SECOND
}

/**
 * Convert nanoseconds to seconds.
 *
 * This performs a division, so the results can be dramatically different if
 * `NS` is an integer or floating point value.
 *
 * \param NS the number of nanoseconds to convert.
 * \returns NS, expressed in seconds.
 *
 * \threadsafety It is safe to call this macro from any thread.
 *
 * \since This macro is available since SDL 3.2.0.
 */
func NS_TO_SECONDS(NS uint64) uint32 {
	return uint32(NS / NS_PER_SECOND)
}

/**
 * Convert milliseconds to nanoseconds.
 *
 * This only converts whole numbers, not fractional milliseconds.
 *
 * \param MS the number of milliseconds to convert.
 * \returns MS, expressed in nanoseconds.
 *
 * \threadsafety It is safe to call this macro from any thread.
 *
 * \since This macro is available since SDL 3.2.0.
 */
func MS_TO_NS(MS uint32) uint64 {
	return uint64(MS) * NS_PER_MS
}

/**
 * Convert nanoseconds to milliseconds.
 *
 * This performs a division, so the results can be dramatically different if
 * `NS` is an integer or floating point value.
 *
 * \param NS the number of nanoseconds to convert.
 * \returns NS, expressed in milliseconds.
 *
 * \threadsafety It is safe to call this macro from any thread.
 *
 * \since This macro is available since SDL 3.2.0.
 */
func NS_TO_MS(NS uint64) uint32 {
	return uint32(NS) / NS_PER_MS
}

/**
 * Convert microseconds to nanoseconds.
 *
 * This only converts whole numbers, not fractional microseconds.
 *
 * \param US the number of microseconds to convert.
 * \returns US, expressed in nanoseconds.
 *
 * \threadsafety It is safe to call this macro from any thread.
 *
 * \since This macro is available since SDL 3.2.0.
 */
func US_TO_NS(US uint64) uint64 {
	return uint64(US) * NS_PER_US
}

/**
 * Convert nanoseconds to microseconds.
 *
 * This performs a division, so the results can be dramatically different if
 * `NS` is an integer or floating point value.
 *
 * \param NS the number of nanoseconds to convert.
 * \returns NS, expressed in microseconds.
 *
 * \threadsafety It is safe to call this macro from any thread.
 *
 * \since This macro is available since SDL 3.2.0.
 */
func NS_TO_US(NS uint64) uint64 {
	return NS / NS_PER_US
}

/**
 * Get the number of milliseconds that have elapsed since the SDL library
 * initialization.
 *
 * \returns an unsigned 64‑bit integer that represents the number of
 *          milliseconds that have elapsed since the SDL library was
 *          initialized (typically via a call to SDL_Init).
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetTicks func() uint64

/**
 * Get the number of nanoseconds since SDL library initialization.
 *
 * \returns an unsigned 64-bit value representing the number of nanoseconds
 *          since the SDL library initialized.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetTicksNS func() uint64

/**
 * Get the current value of the high resolution counter.
 *
 * This function is typically used for profiling.
 *
 * The counter values are only meaningful relative to each other. Differences
 * between values can be converted to times by using
 * SDL_GetPerformanceFrequency().
 *
 * \returns the current counter value.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetPerformanceFrequency
 */
//go:sdl3extern
var GetPerformanceCounter func() uint64

/**
 * Get the count per second of the high resolution counter.
 *
 * \returns a platform-specific count per second.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetPerformanceCounter
 */
//go:sdl3extern
var GetPerformanceFrequency func() uint64

/**
 * Wait a specified number of milliseconds before returning.
 *
 * This function waits a specified number of milliseconds before returning. It
 * waits at least the specified time, but possibly longer due to OS
 * scheduling.
 *
 * \param ms the number of milliseconds to delay.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_DelayNS
 * \sa SDL_DelayPrecise
 */
//go:sdl3extern
var Delay func(ms uint32)

/**
 * Wait a specified number of nanoseconds before returning.
 *
 * This function waits a specified number of nanoseconds before returning. It
 * waits at least the specified time, but possibly longer due to OS
 * scheduling.
 *
 * \param ns the number of nanoseconds to delay.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_Delay
 * \sa SDL_DelayPrecise
 */
//go:sdl3extern
var DelayNS func(ns uint64)

/**
 * Wait a specified number of nanoseconds before returning.
 *
 * This function waits a specified number of nanoseconds before returning. It
 * will attempt to wait as close to the requested time as possible, busy
 * waiting if necessary, but could return later due to OS scheduling.
 *
 * \param ns the number of nanoseconds to delay.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_Delay
 * \sa SDL_DelayNS
 */
//go:sdl3extern
var DelayPrecise func(ns uint64)

/**
 * Definition of the timer ID type.
 *
 * \since This datatype is available since SDL 3.2.0.
 */
type TimerID uint32

/**
 * Function prototype for the millisecond timer callback function.
 *
 * The callback function is passed the current timer interval and returns the
 * next timer interval, in milliseconds. If the returned value is the same as
 * the one passed in, the periodic alarm continues, otherwise a new alarm is
 * scheduled. If the callback returns 0, the periodic alarm is canceled and
 * will be removed.
 *
 * \param userdata an arbitrary pointer provided by the app through
 *                 SDL_AddTimer, for its own use.
 * \param timerID the current timer being processed.
 * \param interval the current callback time interval.
 * \returns the new callback time interval, or 0 to disable further runs of
 *          the callback.
 *
 * \threadsafety SDL may call this callback at any time from a background
 *               thread; the application is responsible for locking resources
 *               the callback touches that need to be protected.
 *
 * \since This datatype is available since SDL 3.2.0.
 *
 * \sa SDL_AddTimer
 */
type TimerCallback func(userdata uintptr, timerID TimerID, interval uint32) uint32

/**
 * Call a callback function at a future time.
 *
 * The callback function is passed the current timer interval and the user
 * supplied parameter from the SDL_AddTimer() call and should return the next
 * timer interval. If the value returned from the callback is 0, the timer is
 * canceled and will be removed.
 *
 * The callback is run on a separate thread, and for short timeouts can
 * potentially be called before this function returns.
 *
 * Timers take into account the amount of time it took to execute the
 * callback. For example, if the callback took 250 ms to execute and returned
 * 1000 (ms), the timer would only wait another 750 ms before its next
 * iteration.
 *
 * Timing may be inexact due to OS scheduling. Be sure to note the current
 * time with SDL_GetTicksNS() or SDL_GetPerformanceCounter() in case your
 * callback needs to adjust for variances.
 *
 * \param interval the timer delay, in milliseconds, passed to `callback`.
 * \param callback the SDL_TimerCallback function to call when the specified
 *                 `interval` elapses.
 * \param userdata a pointer that is passed to `callback`.
 * \returns a timer ID or 0 on failure; call SDL_GetError() for more
 *          information.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddTimerNS
 * \sa SDL_RemoveTimer
 */
//go:sdl3extern
var AddTimer func(interval uint32, callback TimerCallback, userdata uintptr) TimerID

/**
 * Function prototype for the nanosecond timer callback function.
 *
 * The callback function is passed the current timer interval and returns the
 * next timer interval, in nanoseconds. If the returned value is the same as
 * the one passed in, the periodic alarm continues, otherwise a new alarm is
 * scheduled. If the callback returns 0, the periodic alarm is canceled and
 * will be removed.
 *
 * \param userdata an arbitrary pointer provided by the app through
 *                 SDL_AddTimer, for its own use.
 * \param timerID the current timer being processed.
 * \param interval the current callback time interval.
 * \returns the new callback time interval, or 0 to disable further runs of
 *          the callback.
 *
 * \threadsafety SDL may call this callback at any time from a background
 *               thread; the application is responsible for locking resources
 *               the callback touches that need to be protected.
 *
 * \since This datatype is available since SDL 3.2.0.
 *
 * \sa SDL_AddTimerNS
 */
type NSTimerCallback func(userdata uintptr, timerID TimerID, interval uint64) uint64

/**
 * Call a callback function at a future time.
 *
 * The callback function is passed the current timer interval and the user
 * supplied parameter from the SDL_AddTimerNS() call and should return the
 * next timer interval. If the value returned from the callback is 0, the
 * timer is canceled and will be removed.
 *
 * The callback is run on a separate thread, and for short timeouts can
 * potentially be called before this function returns.
 *
 * Timers take into account the amount of time it took to execute the
 * callback. For example, if the callback took 250 ns to execute and returned
 * 1000 (ns), the timer would only wait another 750 ns before its next
 * iteration.
 *
 * Timing may be inexact due to OS scheduling. Be sure to note the current
 * time with SDL_GetTicksNS() or SDL_GetPerformanceCounter() in case your
 * callback needs to adjust for variances.
 *
 * \param interval the timer delay, in nanoseconds, passed to `callback`.
 * \param callback the SDL_TimerCallback function to call when the specified
 *                 `interval` elapses.
 * \param userdata a pointer that is passed to `callback`.
 * \returns a timer ID or 0 on failure; call SDL_GetError() for more
 *          information.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddTimer
 * \sa SDL_RemoveTimer
 */
//go:sdl3extern
var AddTimerNS func(interval uint64, callback NSTimerCallback, userdata uintptr) TimerID

/**
 * Remove a timer created with SDL_AddTimer().
 *
 * \param id the ID of the timer to remove.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddTimer
 */
//go:sdl3extern
var RemoveTimer func(id TimerID) bool
