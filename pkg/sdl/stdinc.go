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

/**
* # CategoryStdinc
*
* SDL provides its own implementation of some of the most important C runtime
* functions.
*
* Using these functions allows an app to have access to common C
* functionality without depending on a specific C runtime (or a C runtime at
* all). More importantly, the SDL implementations work identically across
* platforms, so apps can avoid surprises like snprintf() behaving differently
* between Windows and Linux builds, or itoa() only existing on some
* platforms.
*
* For many of the most common functions, like SDL_memcpy, SDL might just call
* through to the usual C runtime behind the scenes, if it makes sense to do
* so (if it's faster and always available/reliable on a given platform),
* reducing library size and offering the most optimized option.
*
* SDL also offers other C-runtime-adjacent functionality in this header that
* either isn't, strictly speaking, part of any C runtime standards, like
* SDL_crc32() and SDL_reinterpret_cast, etc. It also offers a few better
* options, like SDL_strlcpy(), which functions as a safer form of strcpy().
 */
package sdl

import "math"

type size_t = uint

/**
* The largest value that a `size_t` can hold for the target platform.
*
* `size_t` is generally the same size as a pointer in modern times, but this
* can get weird on very old and very esoteric machines. For example, on a
* 16-bit Intel 286, you might have a 32-bit "far" pointer (16-bit segment
* plus 16-bit offset), but `size_t` is 16 bits, because it can only deal with
* the offset into an individual segment.
*
* In modern times, it's generally expected to cover an entire linear address
* space. But be careful!
*
* \since This macro is available since SDL 3.2.0.
 */
const SIZE_MAX = math.MaxUint

/**
* The number of elements in a static array.
*
* This will compile but return incorrect results for a pointer to an array;
* it has to be an array the compiler knows the size of.
*
* This macro looks like it double-evaluates the argument, but it does so
* inside of `sizeof`, so there are no side-effects here, as expressions do
* not actually run any code in these cases.
*
* \since This macro is available since SDL 3.2.0.
 */
// #define SDL_arraysize(array) (sizeof(array)/sizeof(array[0]))

/**
* Macro useful for building other macros with strings in them.
*
* For example:
*
* ```c
* #define LOG_ERROR(X) OutputDebugString(SDL_STRINGIFY_ARG(__FUNCTION__) ": " X "\n")`
* ```
*
* \param arg the text to turn into a string literal.
*
* \since This macro is available since SDL 3.2.0.
 */
// #define SDL_STRINGIFY_ARG(arg)  #arg

/**
*  \name Cast operators
*
*  Use proper C++ casts when compiled as C++ to be compatible with the option
*  -Wold-style-cast of GCC (and -Werror=old-style-cast in GCC 4.2 and above).
 */
/* @{ */

/* @} */ /* Cast operators */

/**
* Define a four character code as a Uint32.
*
* \param A the first ASCII character.
* \param B the second ASCII character.
* \param C the third ASCII character.
* \param D the fourth ASCII character.
* \returns the four characters converted into a Uint32, one character
*          per-byte.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func FOURCC(A, B, C, D uint8) uint32 {
	return (uint32(A) << 0) |
		(uint32(B) << 8) |
		(uint32(C) << 16) |
		(uint32(D) << 24)
}

/**
*  \name Basic data types
 */
/* @{ */

/**
* SDL times are signed, 64-bit integers representing nanoseconds since the
* Unix epoch (Jan 1, 1970).
*
* They can be converted between POSIX time_t values with SDL_NS_TO_SECONDS()
* and SDL_SECONDS_TO_NS(), and between Windows FILETIME values with
* SDL_TimeToWindows() and SDL_TimeFromWindows().
*
* \since This macro is available since SDL 3.2.0.
*
* \sa SDL_MAX_SINT64
* \sa SDL_MIN_SINT64
 */
type Time int64

const (
	MAX_TIME Time = math.MaxInt64
	MIN_TIME Time = math.MinInt64
)

/* @} */ /* Basic data types */

/**
*  \name Floating-point constants
 */
/* @{ */

/**
* Epsilon constant, used for comparing floating-point numbers.
*
* Equals by default to platform-defined `FLT_EPSILON`, or
* `1.1920928955078125e-07F` if that's not available.
*
* \since This macro is available since SDL 3.2.0.
 */
const FLT_EPSILON = 1.1920928955078125e-07 /* 0x0.000002p0 */

/* @} */ /* Floating-point constants */

/**
* Allocate uninitialized memory.
*
* The allocated memory returned by this function must be freed with
* SDL_free().
*
* If `size` is 0, it will be set to 1.
*
* If the allocation is successful, the returned pointer is guaranteed to be
* aligned to either the *fundamental alignment* (`alignof(max_align_t)` in
* C11 and later) or `2 * sizeof(void *)`, whichever is smaller. Use
* SDL_aligned_alloc() if you need to allocate memory aligned to an alignment
* greater than this guarantee.
*
* \param size the size to allocate.
* \returns a pointer to the allocated memory, or NULL if allocation failed.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_free
* \sa SDL_calloc
* \sa SDL_realloc
* \sa SDL_aligned_alloc
 */
//go:sdl3extern(malloc)
var Malloc func(size uint) uintptr

/**
* Allocate a zero-initialized array.
*
* The memory returned by this function must be freed with SDL_free().
*
* If either of `nmemb` or `size` is 0, they will both be set to 1.
*
* If the allocation is successful, the returned pointer is guaranteed to be
* aligned to either the *fundamental alignment* (`alignof(max_align_t)` in
* C11 and later) or `2 * sizeof(void *)`, whichever is smaller.
*
* \param nmemb the number of elements in the array.
* \param size the size of each element of the array.
* \returns a pointer to the allocated array, or NULL if allocation failed.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_free
* \sa SDL_malloc
* \sa SDL_realloc
 */
//go:sdl3extern(calloc)
var Calloc func(nmemb, size uint) uintptr

/**
* Change the size of allocated memory.
*
* The memory returned by this function must be freed with SDL_free().
*
* If `size` is 0, it will be set to 1. Note that this is unlike some other C
* runtime `realloc` implementations, which may treat `realloc(mem, 0)` the
* same way as `free(mem)`.
*
* If `mem` is NULL, the behavior of this function is equivalent to
* SDL_malloc(). Otherwise, the function can have one of three possible
* outcomes:
*
* - If it returns the same pointer as `mem`, it means that `mem` was resized
*   in place without freeing.
* - If it returns a different non-NULL pointer, it means that `mem` was freed
*   and cannot be dereferenced anymore.
* - If it returns NULL (indicating failure), then `mem` will remain valid and
*   must still be freed with SDL_free().
*
* If the allocation is successfully resized, the returned pointer is
* guaranteed to be aligned to either the *fundamental alignment*
* (`alignof(max_align_t)` in C11 and later) or `2 * sizeof(void *)`,
* whichever is smaller.
*
* \param mem a pointer to allocated memory to reallocate, or NULL.
* \param size the new size of the memory.
* \returns a pointer to the newly allocated memory, or NULL if allocation
*          failed.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_free
* \sa SDL_malloc
* \sa SDL_calloc
 */
//go:sdl3extern(realloc)
var Realloc func(mem uintptr, size uint) uintptr

/**
* Free allocated memory.
*
* The pointer is no longer valid after this call and cannot be dereferenced
* anymore.
*
* If `mem` is NULL, this function does nothing.
*
* \param mem a pointer to allocated memory, or NULL.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_malloc
* \sa SDL_calloc
* \sa SDL_realloc
 */
//go:sdl3extern(free)
var Free func(mem uintptr)

/**
* A callback used to implement SDL_malloc().
*
* SDL will always ensure that the passed `size` is greater than 0.
*
* \param size the size to allocate.
* \returns a pointer to the allocated memory, or NULL if allocation failed.
*
* \threadsafety It should be safe to call this callback from any thread.
*
* \since This datatype is available since SDL 3.2.0.
*
* \sa SDL_malloc
* \sa SDL_GetOriginalMemoryFunctions
* \sa SDL_GetMemoryFunctions
* \sa SDL_SetMemoryFunctions
 */
type MallocFunc func(size uint) uintptr

/**
* A callback used to implement SDL_calloc().
*
* SDL will always ensure that the passed `nmemb` and `size` are both greater
* than 0.
*
* \param nmemb the number of elements in the array.
* \param size the size of each element of the array.
* \returns a pointer to the allocated array, or NULL if allocation failed.
*
* \threadsafety It should be safe to call this callback from any thread.
*
* \since This datatype is available since SDL 3.2.0.
*
* \sa SDL_calloc
* \sa SDL_GetOriginalMemoryFunctions
* \sa SDL_GetMemoryFunctions
* \sa SDL_SetMemoryFunctions
 */
type CallocFunc func(nmemb, size uint) uintptr

/**
* A callback used to implement SDL_realloc().
*
* SDL will always ensure that the passed `size` is greater than 0.
*
* \param mem a pointer to allocated memory to reallocate, or NULL.
* \param size the new size of the memory.
* \returns a pointer to the newly allocated memory, or NULL if allocation
*          failed.
*
* \threadsafety It should be safe to call this callback from any thread.
*
* \since This datatype is available since SDL 3.2.0.
*
* \sa SDL_realloc
* \sa SDL_GetOriginalMemoryFunctions
* \sa SDL_GetMemoryFunctions
* \sa SDL_SetMemoryFunctions
 */
type ReallocFunc func(mem uintptr, size uint) uintptr

/**
* A callback used to implement SDL_free().
*
* SDL will always ensure that the passed `mem` is a non-NULL pointer.
*
* \param mem a pointer to allocated memory.
*
* \threadsafety It should be safe to call this callback from any thread.
*
* \since This datatype is available since SDL 3.2.0.
*
* \sa SDL_free
* \sa SDL_GetOriginalMemoryFunctions
* \sa SDL_GetMemoryFunctions
* \sa SDL_SetMemoryFunctions
 */
type FreeFunc func(mem uintptr)

/**
* Get the original set of SDL memory functions.
*
* This is what SDL_malloc and friends will use by default, if there has been
* no call to SDL_SetMemoryFunctions. This is not necessarily using the C
* runtime's `malloc` functions behind the scenes! Different platforms and
* build configurations might do any number of unexpected things.
*
* \param malloc_func filled with malloc function.
* \param calloc_func filled with calloc function.
* \param realloc_func filled with realloc function.
* \param free_func filled with free function.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetOriginalMemoryFunctions func(malloc_func MallocFunc,
	calloc_func CallocFunc,
	realloc_func ReallocFunc,
	free_func FreeFunc)

/**
* Get the current set of SDL memory functions.
*
* \param malloc_func filled with malloc function.
* \param calloc_func filled with calloc function.
* \param realloc_func filled with realloc function.
* \param free_func filled with free function.
*
* \threadsafety This does not hold a lock, so do not call this in the
*               unlikely event of a background thread calling
*               SDL_SetMemoryFunctions simultaneously.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetMemoryFunctions
* \sa SDL_GetOriginalMemoryFunctions
 */
//go:sdl3extern
var GetMemoryFunctions func(malloc_func *MallocFunc,
	calloc_func *CallocFunc,
	realloc_func *ReallocFunc,
	free_func *FreeFunc)

/**
* Replace SDL's memory allocation functions with a custom set.
*
* It is not safe to call this function once any allocations have been made,
* as future calls to SDL_free will use the new allocator, even if they came
* from an SDL_malloc made with the old one!
*
* If used, usually this needs to be the first call made into the SDL library,
* if not the very first thing done at program startup time.
*
* \param malloc_func custom malloc function.
* \param calloc_func custom calloc function.
* \param realloc_func custom realloc function.
* \param free_func custom free function.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety It is safe to call this function from any thread, but one
*               should not replace the memory functions once any allocations
*               are made!
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetMemoryFunctions
* \sa SDL_GetOriginalMemoryFunctions
 */
//go:sdl3extern
var SetMemoryFunctions func(malloc_func MallocFunc,
	calloc_func CallocFunc,
	realloc_func ReallocFunc,
	free_func FreeFunc) bool

/**
* Allocate memory aligned to a specific alignment.
*
* The memory returned by this function must be freed with SDL_aligned_free(),
* _not_ SDL_free().
*
* If `alignment` is less than the size of `void *`, it will be increased to
* match that.
*
* The returned memory address will be a multiple of the alignment value, and
* the size of the memory allocated will be a multiple of the alignment value.
*
* \param alignment the alignment of the memory.
* \param size the size to allocate.
* \returns a pointer to the aligned memory, or NULL if allocation failed.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_aligned_free
 */
//go:sdl3extern(aligned_alloc)
var AlignedAlloc func(alignment, size uint) uintptr

/**
* Free memory allocated by SDL_aligned_alloc().
*
* The pointer is no longer valid after this call and cannot be dereferenced
* anymore.
*
* If `mem` is NULL, this function does nothing.
*
* \param mem a pointer previously returned by SDL_aligned_alloc(), or NULL.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_aligned_alloc
 */
//go:sdl3extern(aligned_free)
var AlignedFree func(mem uintptr)

/**
* Get the number of outstanding (unfreed) allocations.
*
* \returns the number of allocations or -1 if allocation counting is
*          disabled.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetNumAllocations func() int32

/**
* A thread-safe set of environment variables
*
* \since This struct is available since SDL 3.2.0.
*
* \sa SDL_GetEnvironment
* \sa SDL_CreateEnvironment
* \sa SDL_GetEnvironmentVariable
* \sa SDL_GetEnvironmentVariables
* \sa SDL_SetEnvironmentVariable
* \sa SDL_UnsetEnvironmentVariable
* \sa SDL_DestroyEnvironment
 */
type Environment uintptr

/**
* Get the process environment.
*
* This is initialized at application start and is not affected by setenv()
* and unsetenv() calls after that point. Use SDL_SetEnvironmentVariable() and
* SDL_UnsetEnvironmentVariable() if you want to modify this environment, or
* SDL_setenv_unsafe() or SDL_unsetenv_unsafe() if you want changes to persist
* in the C runtime environment after SDL_Quit().
*
* \returns a pointer to the environment for the process or NULL on failure;
*          call SDL_GetError() for more information.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetEnvironmentVariable
* \sa SDL_GetEnvironmentVariables
* \sa SDL_SetEnvironmentVariable
* \sa SDL_UnsetEnvironmentVariable
 */
//go:sdl3extern
var GetEnvironment func() Environment

/**
* Create a set of environment variables
*
* \param populated true to initialize it from the C runtime environment,
*                  false to create an empty environment.
* \returns a pointer to the new environment or NULL on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety If `populated` is false, it is safe to call this function
*               from any thread, otherwise it is safe if no other threads are
*               calling setenv() or unsetenv()
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetEnvironmentVariable
* \sa SDL_GetEnvironmentVariables
* \sa SDL_SetEnvironmentVariable
* \sa SDL_UnsetEnvironmentVariable
* \sa SDL_DestroyEnvironment
 */
//go:sdl3extern
var CreateEnvironment func(populated bool) Environment

/**
* Get the value of a variable in the environment.
*
* \param env the environment to query.
* \param name the name of the variable to get.
* \returns a pointer to the value of the variable or NULL if it can't be
*          found.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetEnvironment
* \sa SDL_CreateEnvironment
* \sa SDL_GetEnvironmentVariables
* \sa SDL_SetEnvironmentVariable
* \sa SDL_UnsetEnvironmentVariable
 */
//go:sdl3extern
var GetEnvironmentVariable func(env Environment, name uintptr) uintptr

/**
* Get all variables in the environment.
*
* \param env the environment to query.
* \returns a NULL terminated array of pointers to environment variables in
*          the form "variable=value" or NULL on failure; call SDL_GetError()
*          for more information. This is a single allocation that should be
*          freed with SDL_free() when it is no longer needed.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetEnvironment
* \sa SDL_CreateEnvironment
* \sa SDL_GetEnvironmentVariables
* \sa SDL_SetEnvironmentVariable
* \sa SDL_UnsetEnvironmentVariable
 */
var GetEnvironmentVariables func(env Environment) *uintptr

/**
* Set the value of a variable in the environment.
*
* \param env the environment to modify.
* \param name the name of the variable to set.
* \param value the value of the variable to set.
* \param overwrite true to overwrite the variable if it exists, false to
*                  return success without setting the variable if it already
*                  exists.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetEnvironment
* \sa SDL_CreateEnvironment
* \sa SDL_GetEnvironmentVariable
* \sa SDL_GetEnvironmentVariables
* \sa SDL_UnsetEnvironmentVariable
 */
//go:sdl3extern
var SetEnvironmentVariable func(env Environment, name, value uintptr, overwrite bool) bool

/**
* Clear a variable from the environment.
*
* \param env the environment to modify.
* \param name the name of the variable to unset.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetEnvironment
* \sa SDL_CreateEnvironment
* \sa SDL_GetEnvironmentVariable
* \sa SDL_GetEnvironmentVariables
* \sa SDL_SetEnvironmentVariable
* \sa SDL_UnsetEnvironmentVariable
 */
//go:sdl3extern
var UnsetEnvironmentVariable func(env Environment, name uintptr) bool

/**
* Destroy a set of environment variables.
*
* \param env the environment to destroy.
*
* \threadsafety It is safe to call this function from any thread, as long as
*               the environment is no longer in use.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_CreateEnvironment
 */
//go:sdl3extern
var DestroyEnvironment func(env Environment)

/**
* Get the value of a variable in the environment.
*
* This function uses SDL's cached copy of the environment and is thread-safe.
*
* \param name the name of the variable to get.
* \returns a pointer to the value of the variable or NULL if it can't be
*          found.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern(getenv)
var Getenv func(name uintptr) uintptr

/**
* Get the value of a variable in the environment.
*
* This function bypasses SDL's cached copy of the environment and is not
* thread-safe.
*
* \param name the name of the variable to get.
* \returns a pointer to the value of the variable or NULL if it can't be
*          found.
*
* \threadsafety This function is not thread safe, consider using SDL_getenv()
*               instead.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_getenv
 */
//go:sdl3extern(getenv_unsafe)
var GetenvUnsafe func(name uintptr) uintptr

/**
* Set the value of a variable in the environment.
*
* \param name the name of the variable to set.
* \param value the value of the variable to set.
* \param overwrite 1 to overwrite the variable if it exists, 0 to return
*                  success without setting the variable if it already exists.
* \returns 0 on success, -1 on error.
*
* \threadsafety This function is not thread safe, consider using
*               SDL_SetEnvironmentVariable() instead.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetEnvironmentVariable
 */
//go:sdl3extern(setenv_unsafe)
var SetenvUnsafe func(name, value uintptr, overwrite int32) int32

/**
* Clear a variable from the environment.
*
* \param name the name of the variable to unset.
* \returns 0 on success, -1 on error.
*
* \threadsafety This function is not thread safe, consider using
*               SDL_UnsetEnvironmentVariable() instead.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_UnsetEnvironmentVariable
 */
//go:sdl3extern(unsetenv_unsafe)
var UnsetenvUnsafe func(name uintptr) int32
