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
 * # CategoryVideo
 *
 * SDL's video subsystem is largely interested in abstracting window
 * management from the underlying operating system. You can create windows,
 * manage them in various ways, set them fullscreen, and get events when
 * interesting things happen with them, such as the mouse or keyboard
 * interacting with a window.
 *
 * The video subsystem is also interested in abstracting away some
 * platform-specific differences in OpenGL: context creation, swapping
 * buffers, etc. This may be crucial to your app, but also you are not
 * required to use OpenGL at all. In fact, SDL can provide rendering to those
 * windows as well, either with an easy-to-use
 * [2D API](https://wiki.libsdl.org/SDL3/CategoryRender)
 * or with a more-powerful
 * [GPU API](https://wiki.libsdl.org/SDL3/CategoryGPU)
 * . Of course, it can simply get out of your way and give you the window
 * handles you need to use Vulkan, Direct3D, Metal, or whatever else you like
 * directly, too.
 *
 * The video subsystem covers a lot of functionality, out of necessity, so it
 * is worth perusing the list of functions just to see what's available, but
 * most apps can get by with simply creating a window and listening for
 * events, so start with SDL_CreateWindow() and SDL_PollEvent().
 */
package sdl

/**
* This is a unique ID for a display for the time it is connected to the
* system, and is never reused for the lifetime of the application.
*
* If the display is disconnected and reconnected, it will get a new ID.
*
* The value 0 is an invalid ID.
*
* \since This datatype is available since SDL 3.2.0.
 */
type DisplayID uint32

/**
* This is a unique ID for a window.
*
* The value 0 is an invalid ID.
*
* \since This datatype is available since SDL 3.2.0.
 */
type WindowID uint32

/* Global video properties... */

/**
* The pointer to the global `wl_display` object used by the Wayland video
* backend.
*
* Can be set before the video subsystem is initialized to import an external
* `wl_display` object from an application or toolkit for use in SDL, or read
* after initialization to export the `wl_display` used by the Wayland video
* backend. Setting this property after the video subsystem has been
* initialized has no effect, and reading it when the video subsystem is
* uninitialized will either return the user provided value, if one was set
* prior to initialization, or NULL. See docs/README-wayland.md for more
* information.
 */
const PROP_GLOBAL_VIDEO_WAYLAND_WL_DISPLAY_POINTER = "SDL.video.wayland.wl_display"

/**
* System theme.
*
* \since This enum is available since SDL 3.2.0.
 */
type SystemTheme int32

const (
	SystenThemeUnknown SystemTheme = iota /**< Unknown system theme */
	SystenThemeLight                      /**< Light colored system theme */
	SystenThemeDark                       /**< Dark colored system theme */
)

/**
* Internal display mode data.
*
* This lives as a field in SDL_DisplayMode, as opaque data.
*
* \since This struct is available since SDL 3.2.0.
*
* \sa SDL_DisplayMode
 */
type DisplayModeData struct{}

/**
* The structure that defines a display mode.
*
* \since This struct is available since SDL 3.2.0.
*
* \sa SDL_GetFullscreenDisplayModes
* \sa SDL_GetDesktopDisplayMode
* \sa SDL_GetCurrentDisplayMode
* \sa SDL_SetWindowFullscreenMode
* \sa SDL_GetWindowFullscreenMode
 */
type DisplayMode struct {
	DisplayID              DisplayID   /**< the display this mode is associated with */
	Format                 PixelFormat /**< pixel format */
	W                      int32       /**< width */
	H                      int32       /**< height */
	PixelDensity           float32     /**< scale converting size to pixels (e.g. a 1920x1080 mode with 2.0 scale would have 3840x2160 pixels) */
	RefreshRate            float32     /**< refresh rate (or 0.0f for unspecified) */
	RefreshRateNumerator   int32       /**< precise refresh rate numerator (or 0 for unspecified) */
	RefreshRateDenominator int32       /**< precise refresh rate denominator */

	Internal *DisplayModeData /**< Private */
}

/**
* Display orientation values; the way a display is rotated.
*
* \since This enum is available since SDL 3.2.0.
 */
type DisplayOrientation int32

const (
	OrientationUnknown          DisplayOrientation = iota /**< The display orientation can't be determined */
	OrientationLandscape                                  /**< The display is in landscape mode, with the right side up, relative to portrait mode */
	OrientationLandscapeFlipped                           /**< The display is in landscape mode, with the left side up, relative to portrait mode */
	OrientationPortrait                                   /**< The display is in portrait mode */
	OrientationPortraitFlipped                            /**< The display is in portrait mode, upside down */
)

/**
* The struct used as an opaque handle to a window.
*
* \since This struct is available since SDL 3.2.0.
*
* \sa SDL_CreateWindow
 */
type Window uintptr

/**
* The flags on a window.
*
* These cover a lot of true/false, or on/off, window state. Some of it is
* immutable after being set through SDL_CreateWindow(), some of it can be
* changed on existing windows by the app, and some of it might be altered by
* the user or system outside of the app's control.
*
* \since This datatype is available since SDL 3.2.0.
*
* \sa SDL_GetWindowFlags
 */
type WindowFlags uint64

const (
	Fullscreen        WindowFlags = 0x0000000000000001 /**< window is in fullscreen mode */
	OpenGL            WindowFlags = 0x0000000000000002 /**< window usable with OpenGL context */
	Occluded          WindowFlags = 0x0000000000000004 /**< window is occluded */
	Hidden            WindowFlags = 0x0000000000000008 /**< window is neither mapped onto the desktop nor shown in the taskbar/dock/window list; SDL_ShowWindow() is required for it to become visible */
	Borderless        WindowFlags = 0x0000000000000010 /**< no window decoration */
	Resizable         WindowFlags = 0x0000000000000020 /**< window can be resized */
	Minimized         WindowFlags = 0x0000000000000040 /**< window is minimized */
	Maximized         WindowFlags = 0x0000000000000080 /**< window is maximized */
	MouseGrabbed      WindowFlags = 0x0000000000000100 /**< window has grabbed mouse input */
	InputFocus        WindowFlags = 0x0000000000000200 /**< window has input focus */
	MouseFocus        WindowFlags = 0x0000000000000400 /**< window has mouse focus */
	External          WindowFlags = 0x0000000000000800 /**< window not created by SDL */
	Modal             WindowFlags = 0x0000000000001000 /**< window is modal */
	HighPixelDensity  WindowFlags = 0x0000000000002000 /**< window uses high pixel density back buffer if possible */
	Mouse_capture     WindowFlags = 0x0000000000004000 /**< window has mouse captured (unrelated to MOUSE_GRABBED) */
	MouseRelativeMode WindowFlags = 0x0000000000008000 /**< window has relative mode enabled */
	AlwaysOnTop       WindowFlags = 0x0000000000010000 /**< window should always be above others */
	Utility           WindowFlags = 0x0000000000020000 /**< window should be treated as a utility window, not showing in the task bar and window list */
	Tooltip           WindowFlags = 0x0000000000040000 /**< window should be treated as a tooltip and does not get mouse or keyboard focus, requires a parent window */
	Popup_menu        WindowFlags = 0x0000000000080000 /**< window should be treated as a popup menu, requires a parent window */
	KeyboardGrabbed   WindowFlags = 0x0000000000100000 /**< window has grabbed keyboard input */
	Vulkan            WindowFlags = 0x0000000010000000 /**< window usable for Vulkan surface */
	Metal             WindowFlags = 0x0000000020000000 /**< window usable for Metal view */
	Transparent       WindowFlags = 0x0000000040000000 /**< window with transparent buffer */
	NotFocusable      WindowFlags = 0x0000000080000000 /**< window should not be focusable */
)

/**
* A magic value used with SDL_WINDOWPOS_UNDEFINED.
*
* Generally this macro isn't used directly, but rather through
* SDL_WINDOWPOS_UNDEFINED or SDL_WINDOWPOS_UNDEFINED_DISPLAY.
*
* \since This macro is available since SDL 3.2.0.
 */
const WINDOWPOS_UNDEFINED_MASK = 0x1FFF0000

/**
* Used to indicate that you don't care what the window position is.
*
* If you _really_ don't care, SDL_WINDOWPOS_UNDEFINED is the same, but always
* uses the primary display instead of specifying one.
*
* \param X the SDL_DisplayID of the display to use.
*
* \since This macro is available since SDL 3.2.0.
 */
//#define SDL_WINDOWPOS_UNDEFINED_DISPLAY(X)  (SDL_WINDOWPOS_UNDEFINED_MASK|(X))

/**
* Used to indicate that you don't care what the window position/display is.
*
* This always uses the primary display.
*
* \since This macro is available since SDL 3.2.0.
 */
//#define SDL_WINDOWPOS_UNDEFINED         SDL_WINDOWPOS_UNDEFINED_DISPLAY(0)

/**
* A macro to test if the window position is marked as "undefined."
*
* \param X the window position value.
*
* \since This macro is available since SDL 3.2.0.
 */
//#define SDL_WINDOWPOS_ISUNDEFINED(X)    (((X)&0xFFFF0000) == SDL_WINDOWPOS_UNDEFINED_MASK)

/**
* A magic value used with SDL_WINDOWPOS_CENTERED.
*
* Generally this macro isn't used directly, but rather through
* SDL_WINDOWPOS_CENTERED or SDL_WINDOWPOS_CENTERED_DISPLAY.
*
* \since This macro is available since SDL 3.2.0.
 */
const WINDOWPOS_CENTERED_MASK = 0x2FFF0000

/**
* Used to indicate that the window position should be centered.
*
* SDL_WINDOWPOS_CENTERED is the same, but always uses the primary display
* instead of specifying one.
*
* \param X the SDL_DisplayID of the display to use.
*
* \since This macro is available since SDL 3.2.0.
 */
// #define SDL_WINDOWPOS_CENTERED_DISPLAY(X)  (SDL_WINDOWPOS_CENTERED_MASK|(X))

/**
* Used to indicate that the window position should be centered.
*
* This always uses the primary display.
*
* \since This macro is available since SDL 3.2.0.
 */
// #define SDL_WINDOWPOS_CENTERED         SDL_WINDOWPOS_CENTERED_DISPLAY(0)

/**
* A macro to test if the window position is marked as "centered."
*
* \param X the window position value.
*
* \since This macro is available since SDL 3.2.0.
 */
// #define SDL_WINDOWPOS_ISCENTERED(X)    \
// 			(((X)&0xFFFF0000) == SDL_WINDOWPOS_CENTERED_MASK)

/**
* Window flash operation.
*
* \since This enum is available since SDL 3.2.0.
 */
type FlashOperation int32

const (
	FlashCancel       FlashOperation = iota /**< Cancel any window flash state */
	FlashBriefly                            /**< Flash the window briefly to get attention */
	FlashUntilFocused                       /**< Flash the window until it gets focus */
)

/**
* Window progress state
*
* \since This enum is available since SDL 3.2.8.
 */
type ProgressState int32

const (
	ProgressStateInvalid       ProgressState = iota - 1 /**< An invalid progress state indicating an error; check SDL_GetError() */
	ProgressStateNone                                   /**< No progress bar is shown */
	ProgressStateIndeterminate                          /**< The progress bar is shown in a indeterminate state */
	ProgressStateNormal                                 /**< The progress bar is shown in a normal state */
	ProgressStatePaused                                 /**< The progress bar is shown in a paused state */
	ProgressStateError                                  /**< The progress bar is shown in a state indicating the application had an error */
)

/**
* An opaque handle to an OpenGL context.
*
* \since This datatype is available since SDL 3.2.0.
*
* \sa SDL_GL_CreateContext
 */
type GLContext uintptr

/**
* Opaque type for an EGL display.
*
* \since This datatype is available since SDL 3.2.0.
 */
type EGLDisplay uintptr

/**
* Opaque type for an EGL config.
*
* \since This datatype is available since SDL 3.2.0.
 */
type EGLConfig uintptr

/**
* Opaque type for an EGL surface.
*
* \since This datatype is available since SDL 3.2.0.
 */
type EGLSurface uintptr

/**
* An EGL attribute, used when creating an EGL context.
*
* \since This datatype is available since SDL 3.2.0.
 */
type EGLAttrib uintptr

/**
* An EGL integer attribute, used when creating an EGL surface.
*
* \since This datatype is available since SDL 3.2.0.
 */
type EGLint int32

/**
* EGL platform attribute initialization callback.
*
* This is called when SDL is attempting to create an EGL context, to let the
* app add extra attributes to its eglGetPlatformDisplay() call.
*
* The callback should return a pointer to an EGL attribute array terminated
* with `EGL_NONE`. If this function returns NULL, the SDL_CreateWindow
* process will fail gracefully.
*
* The returned pointer should be allocated with SDL_malloc() and will be
* passed to SDL_free().
*
* The arrays returned by each callback will be appended to the existing
* attribute arrays defined by SDL.
*
* \param userdata an app-controlled pointer that is passed to the callback.
* \returns a newly-allocated array of attributes, terminated with `EGL_NONE`.
*
* \since This datatype is available since SDL 3.2.0.
*
* \sa SDL_EGL_SetAttributeCallbacks
 */
type EGLAttribArrayCallback func(userdata uintptr) *EGLAttrib

/**
* EGL surface/context attribute initialization callback types.
*
* This is called when SDL is attempting to create an EGL surface, to let the
* app add extra attributes to its eglCreateWindowSurface() or
* eglCreateContext calls.
*
* For convenience, the EGLDisplay and EGLConfig to use are provided to the
* callback.
*
* The callback should return a pointer to an EGL attribute array terminated
* with `EGL_NONE`. If this function returns NULL, the SDL_CreateWindow
* process will fail gracefully.
*
* The returned pointer should be allocated with SDL_malloc() and will be
* passed to SDL_free().
*
* The arrays returned by each callback will be appended to the existing
* attribute arrays defined by SDL.
*
* \param userdata an app-controlled pointer that is passed to the callback.
* \param display the EGL display to be used.
* \param config the EGL config to be used.
* \returns a newly-allocated array of attributes, terminated with `EGL_NONE`.
*
* \since This datatype is available since SDL 3.2.0.
*
* \sa SDL_EGL_SetAttributeCallbacks
 */
// typedef SDL_EGLint *(SDLCALL *SDL_EGLIntArrayCallback)(void *userdata, SDL_EGLDisplay display, SDL_EGLConfig config);
type EGLIntArrayCallback func(userdata uintptr, display EGLDisplay, config EGLConfig) EGLint

/**
* An enumeration of OpenGL configuration attributes.
*
* While you can set most OpenGL attributes normally, the attributes listed
* above must be known before SDL creates the window that will be used with
* the OpenGL context. These attributes are set and read with
* SDL_GL_SetAttribute() and SDL_GL_GetAttribute().
*
* In some cases, these attributes are minimum requests; the GL does not
* promise to give you exactly what you asked for. It's possible to ask for a
* 16-bit depth buffer and get a 24-bit one instead, for example, or to ask
* for no stencil buffer and still have one available. Context creation should
* fail if the GL can't provide your requested attributes at a minimum, but
* you should check to see exactly what you got.
*
* \since This enum is available since SDL 3.2.0.
 */
type GLAttr int32

const (
	GLA_RedSize                  GLAttr = iota /**< the minimum number of bits for the red channel of the color buffer; defaults to 8. */
	GLA_GreenSize                              /**< the minimum number of bits for the green channel of the color buffer; defaults to 8. */
	GLA_BlueSize                               /**< the minimum number of bits for the blue channel of the color buffer; defaults to 8. */
	GLA_AlphaSize                              /**< the minimum number of bits for the alpha channel of the color buffer; defaults to 8. */
	GLA_BufferSize                             /**< the minimum number of bits for frame buffer size; defaults to 0. */
	GLA_Doublebuffer                           /**< whether the output is single or double buffered; defaults to double buffering on. */
	GLA_DepthSize                              /**< the minimum number of bits in the depth buffer; defaults to 16. */
	GLA_StencilSize                            /**< the minimum number of bits in the stencil buffer; defaults to 0. */
	GLA_AccumRedSize                           /**< the minimum number of bits for the red channel of the accumulation buffer; defaults to 0. */
	GLA_AccumGreenSize                         /**< the minimum number of bits for the green channel of the accumulation buffer; defaults to 0. */
	GLA_AccumBlueSize                          /**< the minimum number of bits for the blue channel of the accumulation buffer; defaults to 0. */
	GLA_AccumAlphaSize                         /**< the minimum number of bits for the alpha channel of the accumulation buffer; defaults to 0. */
	GLA_Stereo                                 /**< whether the output is stereo 3D; defaults to off. */
	GLA_Multisamplebuffers                     /**< the number of buffers used for multisample anti-aliasing; defaults to 0. */
	GLA_Multisamplesamples                     /**< the number of samples used around the current pixel used for multisample anti-aliasing. */
	GLA_AcceleratedVisual                      /**< set to 1 to require hardware acceleration, set to 0 to force software rendering; defaults to allow either. */
	GLA_RetainedBacking                        /**< not used (deprecated). */
	GLA_ContextMajorVersion                    /**< OpenGL context major version. */
	GLA_ContextMinorVersion                    /**< OpenGL context minor version. */
	GLA_ContextFlags                           /**< some combination of 0 or more of elements of the SDL_GLContextFlag enumeration; defaults to 0. */
	GLA_ContextProfileMask                     /**< type of GL context (Core, Compatibility, ES). See SDL_GLProfile; default value depends on platform. */
	GLA_ShareWithCurrentContext                /**< OpenGL context sharing; defaults to 0. */
	GLA_FramebufferSrgbCapable                 /**< requests sRGB capable visual; defaults to 0. */
	GLA_ContextReleaseBehavior                 /**< sets context the release behavior. See SDL_GLContextReleaseFlag; defaults to FLUSH. */
	GLA_ContextResetNotification               /**< set context reset notification. See SDL_GLContextResetNotification; defaults to NO_NOTIFICATION. */
	GLA_ContextNoError
	GLA_Floatbuffers
	GLA_EglPlatform
)

/**
* Possible values to be set for the SDL_GL_CONTEXT_PROFILE_MASK attribute.
*
* \since This datatype is available since SDL 3.2.0.
 */
type GLProfile uint32

const (
	GL_CONTEXT_PROFILE_CORE          GLProfile = 0x0001 /**< OpenGL Core Profile context */
	GL_CONTEXT_PROFILE_COMPATIBILITY GLProfile = 0x0002 /**< OpenGL Compatibility Profile context */
	GL_CONTEXT_PROFILE_ES            GLProfile = 0x0004 /**< GLX_CONTEXT_ES2_PROFILE_BIT_EXT */
)

/**
* Possible flags to be set for the SDL_GL_CONTEXT_FLAGS attribute.
*
* \since This datatype is available since SDL 3.2.0.
 */
type GLContextFlag uint32

const (
	GL_CONTEXT_DEBUG_FLAG              GLContextFlag = 0x0001
	GL_CONTEXT_FORWARD_COMPATIBLE_FLAG GLContextFlag = 0x0002
	GL_CONTEXT_ROBUST_ACCESS_FLAG      GLContextFlag = 0x0004
	GL_CONTEXT_RESET_ISOLATION_FLAG    GLContextFlag = 0x0008
)

/**
* Possible values to be set for the SDL_GL_CONTEXT_RELEASE_BEHAVIOR
* attribute.
*
* \since This datatype is available since SDL 3.2.0.
 */
type GLContextReleaseFlag uint32

const (
	GL_CONTEXT_RELEASE_BEHAVIOR_NONE  GLContextReleaseFlag = 0x0000
	GL_CONTEXT_RELEASE_BEHAVIOR_FLUSH GLContextReleaseFlag = 0x0001
)

/**
* Possible values to be set SDL_GL_CONTEXT_RESET_NOTIFICATION attribute.
*
* \since This datatype is available since SDL 3.2.0.
 */
type GLContextResetNotification uint32

const (
	GL_CONTEXT_RESET_NO_NOTIFICATION GLContextResetNotification = 0x0000
	GL_CONTEXT_RESET_LOSE_CONTEXT    GLContextResetNotification = 0x0001
)

/* Function prototypes */

/**
* Get the number of video drivers compiled into SDL.
*
* \returns the number of built in video drivers.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetVideoDriver
 */
//go:sdl3extern
var GetNumVideoDrivers func() int32

/**
* Get the name of a built in video driver.
*
* The video drivers are presented in the order in which they are normally
* checked during initialization.
*
* The names of drivers are all simple, low-ASCII identifiers, like "cocoa",
* "x11" or "windows". These never have Unicode characters, and are not meant
* to be proper names.
*
* \param index the index of a video driver.
* \returns the name of the video driver with the given **index**.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetNumVideoDrivers
 */
//go:sdl3extern
var GetVideoDriver func(index int32) *byte

/**
* Get the name of the currently initialized video driver.
*
* The names of drivers are all simple, low-ASCII identifiers, like "cocoa",
* "x11" or "windows". These never have Unicode characters, and are not meant
* to be proper names.
*
* \returns the name of the current video driver or NULL if no driver has been
*          initialized.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetNumVideoDrivers
* \sa SDL_GetVideoDriver
 */
//go:sdl3extern
var GetCurrentVideoDriver func() *byte

/**
* Get the current system theme.
*
* \returns the current system theme, light, dark, or unknown.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetSystemTheme func() SystemTheme

/**
* Get a list of currently connected displays.
*
* \param count a pointer filled in with the number of displays returned, may
*              be NULL.
* \returns a 0 terminated array of display instance IDs or NULL on failure;
*          call SDL_GetError() for more information. This should be freed
*          with SDL_free() when it is no longer needed.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetDisplays func(count *int32) DisplayID

/**
* Return the primary display.
*
* \returns the instance ID of the primary display on success or 0 on failure;
*          call SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetPrimaryDisplay func() DisplayID

/**
* Get the properties associated with a display.
*
* The following read-only properties are provided by SDL:
*
* - `SDL_PROP_DISPLAY_HDR_ENABLED_BOOLEAN`: true if the display has HDR
*   headroom above the SDR white point. This is for informational and
*   diagnostic purposes only, as not all platforms provide this information
*   at the display level.
*
* On KMS/DRM:
*
* - `SDL_PROP_DISPLAY_KMSDRM_PANEL_ORIENTATION_NUMBER`: the "panel
*   orientation" property for the display in degrees of clockwise rotation.
*   Note that this is provided only as a hint, and the application is
*   responsible for any coordinate transformations needed to conform to the
*   requested display orientation.
*
* On Wayland:
*
* - `SDL_PROP_DISPLAY_WAYLAND_WL_OUTPUT_POINTER`: the wl_output associated
*   with the display
*
* \param displayID the instance ID of the display to query.
* \returns a valid property ID on success or 0 on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetDisplayProperties func(displayID DisplayID) PropertiesID

const (
	PROP_DISPLAY_HDR_ENABLED_BOOLEAN             PropertyName = "SDL.display.HDR_enabled"
	PROP_DISPLAY_KMSDRM_PANEL_ORIENTATION_NUMBER PropertyName = "SDL.display.KMSDRM.panel_orientation"
	PROP_DISPLAY_WAYLAND_WL_OUTPUT_POINTER       PropertyName = "SDL.display.wayland.wl_output"
)

/**
* Get the name of a display in UTF-8 encoding.
*
* \param displayID the instance ID of the display to query.
* \returns the name of a display or NULL on failure; call SDL_GetError() for
*          more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetDisplayName func(displayID DisplayID) *byte

/**
* Get the desktop area represented by a display.
*
* The primary display is often located at (0,0), but may be placed at a
* different location depending on monitor layout.
*
* \param displayID the instance ID of the display to query.
* \param rect the SDL_Rect structure filled in with the display bounds.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDisplayUsableBounds
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetDisplayBounds func(displayID DisplayID, rect *Rect) bool

/**
* Get the usable desktop area represented by a display, in screen
* coordinates.
*
* This is the same area as SDL_GetDisplayBounds() reports, but with portions
* reserved by the system removed. For example, on Apple's macOS, this
* subtracts the area occupied by the menu bar and dock.
*
* Setting a window to be fullscreen generally bypasses these unusable areas,
* so these are good guidelines for the maximum space available to a
* non-fullscreen window.
*
* \param displayID the instance ID of the display to query.
* \param rect the SDL_Rect structure filled in with the display bounds.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDisplayBounds
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetDisplayUsableBounds func(displayID DisplayID, rect *Rect) bool

/**
* Get the orientation of a display when it is unrotated.
*
* \param displayID the instance ID of the display to query.
* \returns the SDL_DisplayOrientation enum value of the display, or
*          `SDL_ORIENTATION_UNKNOWN` if it isn't available.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetNaturalDisplayOrientation func(displayID DisplayID) DisplayOrientation

/**
* Get the orientation of a display.
*
* \param displayID the instance ID of the display to query.
* \returns the SDL_DisplayOrientation enum value of the display, or
*          `SDL_ORIENTATION_UNKNOWN` if it isn't available.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetCurrentDisplayOrientation func(displayID DisplayID) DisplayOrientation

/**
* Get the content scale of a display.
*
* The content scale is the expected scale for content based on the DPI
* settings of the display. For example, a 4K display might have a 2.0 (200%)
* display scale, which means that the user expects UI elements to be twice as
* big on this display, to aid in readability.
*
* After window creation, SDL_GetWindowDisplayScale() should be used to query
* the content scale factor for individual windows instead of querying the
* display for a window and calling this function, as the per-window content
* scale factor may differ from the base value of the display it is on,
* particularly on high-DPI and/or multi-monitor desktop configurations.
*
* \param displayID the instance ID of the display to query.
* \returns the content scale of the display, or 0.0f on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowDisplayScale
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetDisplayContentScale func(displayID DisplayID) float32

/**
* Get a list of fullscreen display modes available on a display.
*
* The display modes are sorted in this priority:
*
* - w -> largest to smallest
* - h -> largest to smallest
* - bits per pixel -> more colors to fewer colors
* - packed pixel layout -> largest to smallest
* - refresh rate -> highest to lowest
* - pixel density -> lowest to highest
*
* \param displayID the instance ID of the display to query.
* \param count a pointer filled in with the number of display modes returned,
*              may be NULL.
* \returns a NULL terminated array of display mode pointers or NULL on
*          failure; call SDL_GetError() for more information. This is a
*          single allocation that should be freed with SDL_free() when it is
*          no longer needed.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetFullscreenDisplayModes func(displayID DisplayID, count *int32) **DisplayMode

/**
* Get the closest match to the requested display mode.
*
* The available display modes are scanned and `closest` is filled in with the
* closest mode matching the requested mode and returned. The mode format and
* refresh rate default to the desktop mode if they are set to 0. The modes
* are scanned with size being first priority, format being second priority,
* and finally checking the refresh rate. If all the available modes are too
* small, then false is returned.
*
* \param displayID the instance ID of the display to query.
* \param w the width in pixels of the desired display mode.
* \param h the height in pixels of the desired display mode.
* \param refresh_rate the refresh rate of the desired display mode, or 0.0f
*                     for the desktop refresh rate.
* \param include_high_density_modes boolean to include high density modes in
*                                   the search.
* \param closest a pointer filled in with the closest display mode equal to
*                or larger than the desired mode.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDisplays
* \sa SDL_GetFullscreenDisplayModes
 */
//go:sdl3extern
var GetClosestFullscreenDisplayMode func(displayID DisplayID, w, h int32, refresh_rate float32, include_high_density_modes bool, closest *DisplayMode) bool

/**
* Get information about the desktop's display mode.
*
* There's a difference between this function and SDL_GetCurrentDisplayMode()
* when SDL runs fullscreen and has changed the resolution. In that case this
* function will return the previous native display mode, and not the current
* display mode.
*
* \param displayID the instance ID of the display to query.
* \returns a pointer to the desktop display mode or NULL on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetCurrentDisplayMode
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetDesktopDisplayMode func(displayID DisplayID) *DisplayMode

/**
* Get information about the current display mode.
*
* There's a difference between this function and SDL_GetDesktopDisplayMode()
* when SDL runs fullscreen and has changed the resolution. In that case this
* function will return the current display mode, and not the previous native
* display mode.
*
* \param displayID the instance ID of the display to query.
* \returns a pointer to the desktop display mode or NULL on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDesktopDisplayMode
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetCurrentDisplayMode func(displayID DisplayID) *DisplayMode

/**
* Get the display containing a point.
*
* \param point the point to query.
* \returns the instance ID of the display containing the point or 0 on
*          failure; call SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDisplayBounds
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetDisplayForPoint func(point *Point) DisplayID

/**
* Get the display primarily containing a rect.
*
* \param rect the rect to query.
* \returns the instance ID of the display entirely containing the rect or
*          closest to the center of the rect on success or 0 on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDisplayBounds
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetDisplayForRect func(rect *Rect) DisplayID

/**
* Get the display associated with a window.
*
* \param window the window to query.
* \returns the instance ID of the display containing the center of the window
*          on success or 0 on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetDisplayBounds
* \sa SDL_GetDisplays
 */
//go:sdl3extern
var GetDisplayForWindow func(window Window) DisplayID

/**
* Get the pixel density of a window.
*
* This is a ratio of pixel size to window size. For example, if the window is
* 1920x1080 and it has a high density back buffer of 3840x2160 pixels, it
* would have a pixel density of 2.0.
*
* \param window the window to query.
* \returns the pixel density or 0.0f on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowDisplayScale
 */
//go:sdl3extern
var GetWindowPixelDensity func(window Window) float32

/**
* Get the content display scale relative to a window's pixel size.
*
* This is a combination of the window pixel density and the display content
* scale, and is the expected scale for displaying content in this window. For
* example, if a 3840x2160 window had a display scale of 2.0, the user expects
* the content to take twice as many pixels and be the same physical size as
* if it were being displayed in a 1920x1080 window with a display scale of
* 1.0.
*
* Conceptually this value corresponds to the scale display setting, and is
* updated when that setting is changed, or the window moves to a display with
* a different scale setting.
*
* \param window the window to query.
* \returns the display scale, or 0.0f on failure; call SDL_GetError() for
*          more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetWindowDisplayScale func(window Window) float32

/**
* Set the display mode to use when a window is visible and fullscreen.
*
* This only affects the display mode used when the window is fullscreen. To
* change the window size when the window is not fullscreen, use
* SDL_SetWindowSize().
*
* If the window is currently in the fullscreen state, this request is
* asynchronous on some windowing systems and the new mode dimensions may not
* be applied immediately upon the return of this function. If an immediate
* change is required, call SDL_SyncWindow() to block until the changes have
* taken effect.
*
* When the new mode takes effect, an SDL_EVENT_WINDOW_RESIZED and/or an
* SDL_EVENT_WINDOW_PIXEL_SIZE_CHANGED event will be emitted with the new mode
* dimensions.
*
* \param window the window to affect.
* \param mode a pointer to the display mode to use, which can be NULL for
*             borderless fullscreen desktop mode, or one of the fullscreen
*             modes returned by SDL_GetFullscreenDisplayModes() to set an
*             exclusive fullscreen mode.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowFullscreenMode
* \sa SDL_SetWindowFullscreen
* \sa SDL_SyncWindow
 */
//go:sdl3extern
var SetWindowFullscreenMode func(window Window, mode *DisplayMode) bool

/**
* Query the display mode to use when a window is visible at fullscreen.
*
* \param window the window to query.
* \returns a pointer to the exclusive fullscreen mode to use or NULL for
*          borderless fullscreen desktop mode.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowFullscreenMode
* \sa SDL_SetWindowFullscreen
 */
//go:sdl3extern
var GetWindowFullscreenMode func(window Window) *DisplayMode

/**
* Get the raw ICC profile data for the screen the window is currently on.
*
* \param window the window to query.
* \param size the size of the ICC profile.
* \returns the raw ICC profile data on success or NULL on failure; call
*          SDL_GetError() for more information. This should be freed with
*          SDL_free() when it is no longer needed.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetWindowICCProfile func(window Window, size *uint) uintptr

/**
* Get the pixel format associated with the window.
*
* \param window the window to query.
* \returns the pixel format of the window on success or
*          SDL_PIXELFORMAT_UNKNOWN on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetWindowPixelFormat func(window Window) PixelFormat

/**
* Get a list of valid windows.
*
* \param count a pointer filled in with the number of windows returned, may
*              be NULL.
* \returns a NULL terminated array of SDL_Window pointers or NULL on failure;
*          call SDL_GetError() for more information. This is a single
*          allocation that should be freed with SDL_free() when it is no
*          longer needed.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetWindows func(count *int32) *Window

/**
* Create a window with the specified dimensions and flags.
*
* The window size is a request and may be different than expected based on
* the desktop layout and window manager policies. Your application should be
* prepared to handle a window of any size.
*
* `flags` may be any of the following OR'd together:
*
* - `SDL_WINDOW_FULLSCREEN`: fullscreen window at desktop resolution
* - `SDL_WINDOW_OPENGL`: window usable with an OpenGL context
* - `SDL_WINDOW_OCCLUDED`: window partially or completely obscured by another
*   window
* - `SDL_WINDOW_HIDDEN`: window is not visible
* - `SDL_WINDOW_BORDERLESS`: no window decoration
* - `SDL_WINDOW_RESIZABLE`: window can be resized
* - `SDL_WINDOW_MINIMIZED`: window is minimized
* - `SDL_WINDOW_MAXIMIZED`: window is maximized
* - `SDL_WINDOW_MOUSE_GRABBED`: window has grabbed mouse focus
* - `SDL_WINDOW_INPUT_FOCUS`: window has input focus
* - `SDL_WINDOW_MOUSE_FOCUS`: window has mouse focus
* - `SDL_WINDOW_EXTERNAL`: window not created by SDL
* - `SDL_WINDOW_MODAL`: window is modal
* - `SDL_WINDOW_HIGH_PIXEL_DENSITY`: window uses high pixel density back
*   buffer if possible
* - `SDL_WINDOW_MOUSE_CAPTURE`: window has mouse captured (unrelated to
*   MOUSE_GRABBED)
* - `SDL_WINDOW_ALWAYS_ON_TOP`: window should always be above others
* - `SDL_WINDOW_UTILITY`: window should be treated as a utility window, not
*   showing in the task bar and window list
* - `SDL_WINDOW_TOOLTIP`: window should be treated as a tooltip and does not
*   get mouse or keyboard focus, requires a parent window
* - `SDL_WINDOW_POPUP_MENU`: window should be treated as a popup menu,
*   requires a parent window
* - `SDL_WINDOW_KEYBOARD_GRABBED`: window has grabbed keyboard input
* - `SDL_WINDOW_VULKAN`: window usable with a Vulkan instance
* - `SDL_WINDOW_METAL`: window usable with a Metal instance
* - `SDL_WINDOW_TRANSPARENT`: window with transparent buffer
* - `SDL_WINDOW_NOT_FOCUSABLE`: window should not be focusable
*
* The SDL_Window is implicitly shown if SDL_WINDOW_HIDDEN is not set.
*
* On Apple's macOS, you **must** set the NSHighResolutionCapable Info.plist
* property to YES, otherwise you will not receive a High-DPI OpenGL canvas.
*
* The window pixel size may differ from its window coordinate size if the
* window is on a high pixel density display. Use SDL_GetWindowSize() to query
* the client area's size in window coordinates, and
* SDL_GetWindowSizeInPixels() or SDL_GetRenderOutputSize() to query the
* drawable size in pixels. Note that the drawable size can vary after the
* window is created and should be queried again if you get an
* SDL_EVENT_WINDOW_PIXEL_SIZE_CHANGED event.
*
* If the window is created with any of the SDL_WINDOW_OPENGL or
* SDL_WINDOW_VULKAN flags, then the corresponding LoadLibrary function
* (SDL_GL_LoadLibrary or SDL_Vulkan_LoadLibrary) is called and the
* corresponding UnloadLibrary function is called by SDL_DestroyWindow().
*
* If SDL_WINDOW_VULKAN is specified and there isn't a working Vulkan driver,
* SDL_CreateWindow() will fail, because SDL_Vulkan_LoadLibrary() will fail.
*
* If SDL_WINDOW_METAL is specified on an OS that does not support Metal,
* SDL_CreateWindow() will fail.
*
* If you intend to use this window with an SDL_Renderer, you should use
* SDL_CreateWindowAndRenderer() instead of this function, to avoid window
* flicker.
*
* On non-Apple devices, SDL requires you to either not link to the Vulkan
* loader or link to a dynamic library version. This limitation may be removed
* in a future version of SDL.
*
* \param title the title of the window, in UTF-8 encoding.
* \param w the width of the window.
* \param h the height of the window.
* \param flags 0, or one or more SDL_WindowFlags OR'd together.
* \returns the window that was created or NULL on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_CreateWindowAndRenderer
* \sa SDL_CreatePopupWindow
* \sa SDL_CreateWindowWithProperties
* \sa SDL_DestroyWindow
 */
//go:sdl3extern
var CreateWindow func(title *byte, w, h int32, flags WindowFlags) Window

/**
* Create a child popup window of the specified parent window.
*
* The window size is a request and may be different than expected based on
* the desktop layout and window manager policies. Your application should be
* prepared to handle a window of any size.
*
* The flags parameter **must** contain at least one of the following:
*
* - `SDL_WINDOW_TOOLTIP`: The popup window is a tooltip and will not pass any
*   input events.
* - `SDL_WINDOW_POPUP_MENU`: The popup window is a popup menu. The topmost
*   popup menu will implicitly gain the keyboard focus.
*
* The following flags are not relevant to popup window creation and will be
* ignored:
*
* - `SDL_WINDOW_MINIMIZED`
* - `SDL_WINDOW_MAXIMIZED`
* - `SDL_WINDOW_FULLSCREEN`
* - `SDL_WINDOW_BORDERLESS`
*
* The following flags are incompatible with popup window creation and will
* cause it to fail:
*
* - `SDL_WINDOW_UTILITY`
* - `SDL_WINDOW_MODAL`
*
* The parent parameter **must** be non-null and a valid window. The parent of
* a popup window can be either a regular, toplevel window, or another popup
* window.
*
* Popup windows cannot be minimized, maximized, made fullscreen, raised,
* flash, be made a modal window, be the parent of a toplevel window, or grab
* the mouse and/or keyboard. Attempts to do so will fail.
*
* Popup windows implicitly do not have a border/decorations and do not appear
* on the taskbar/dock or in lists of windows such as alt-tab menus.
*
* By default, popup window positions will automatically be constrained to
* keep the entire window within display bounds. This can be overridden with
* the `SDL_PROP_WINDOW_CREATE_CONSTRAIN_POPUP_BOOLEAN` property.
*
* By default, popup menus will automatically grab keyboard focus from the
* parent when shown. This behavior can be overridden by setting the
* `SDL_WINDOW_NOT_FOCUSABLE` flag, setting the
* `SDL_PROP_WINDOW_CREATE_FOCUSABLE_BOOLEAN` property to false, or toggling
* it after creation via the `SDL_SetWindowFocusable()` function.
*
* If a parent window is hidden or destroyed, any child popup windows will be
* recursively hidden or destroyed as well. Child popup windows not explicitly
* hidden will be restored when the parent is shown.
*
* \param parent the parent of the window, must not be NULL.
* \param offset_x the x position of the popup window relative to the origin
*                 of the parent.
* \param offset_y the y position of the popup window relative to the origin
*                 of the parent window.
* \param w the width of the window.
* \param h the height of the window.
* \param flags SDL_WINDOW_TOOLTIP or SDL_WINDOW_POPUP_MENU, and zero or more
*              additional SDL_WindowFlags OR'd together.
* \returns the window that was created or NULL on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_CreateWindow
* \sa SDL_CreateWindowWithProperties
* \sa SDL_DestroyWindow
* \sa SDL_GetWindowParent
 */
//go:sdl3extern
var CreatePopupWindow func(parent Window, offset_x, offset_y, w, h int32, flags WindowFlags) Window

/**
* Create a window with the specified properties.
*
* The window size is a request and may be different than expected based on
* the desktop layout and window manager policies. Your application should be
* prepared to handle a window of any size.
*
* These are the supported properties:
*
* - `SDL_PROP_WINDOW_CREATE_ALWAYS_ON_TOP_BOOLEAN`: true if the window should
*   be always on top
* - `SDL_PROP_WINDOW_CREATE_BORDERLESS_BOOLEAN`: true if the window has no
*   window decoration
* - `SDL_PROP_WINDOW_CREATE_CONSTRAIN_POPUP_BOOLEAN`: true if the "tooltip"
*   and "menu" window types should be automatically constrained to be
*   entirely within display bounds (default), false if no constraints on the
*   position are desired.
* - `SDL_PROP_WINDOW_CREATE_EXTERNAL_GRAPHICS_CONTEXT_BOOLEAN`: true if the
*   window will be used with an externally managed graphics context.
* - `SDL_PROP_WINDOW_CREATE_FOCUSABLE_BOOLEAN`: true if the window should
*   accept keyboard input (defaults true)
* - `SDL_PROP_WINDOW_CREATE_FULLSCREEN_BOOLEAN`: true if the window should
*   start in fullscreen mode at desktop resolution
* - `SDL_PROP_WINDOW_CREATE_HEIGHT_NUMBER`: the height of the window
* - `SDL_PROP_WINDOW_CREATE_HIDDEN_BOOLEAN`: true if the window should start
*   hidden
* - `SDL_PROP_WINDOW_CREATE_HIGH_PIXEL_DENSITY_BOOLEAN`: true if the window
*   uses a high pixel density buffer if possible
* - `SDL_PROP_WINDOW_CREATE_MAXIMIZED_BOOLEAN`: true if the window should
*   start maximized
* - `SDL_PROP_WINDOW_CREATE_MENU_BOOLEAN`: true if the window is a popup menu
* - `SDL_PROP_WINDOW_CREATE_METAL_BOOLEAN`: true if the window will be used
*   with Metal rendering
* - `SDL_PROP_WINDOW_CREATE_MINIMIZED_BOOLEAN`: true if the window should
*   start minimized
* - `SDL_PROP_WINDOW_CREATE_MODAL_BOOLEAN`: true if the window is modal to
*   its parent
* - `SDL_PROP_WINDOW_CREATE_MOUSE_GRABBED_BOOLEAN`: true if the window starts
*   with grabbed mouse focus
* - `SDL_PROP_WINDOW_CREATE_OPENGL_BOOLEAN`: true if the window will be used
*   with OpenGL rendering
* - `SDL_PROP_WINDOW_CREATE_PARENT_POINTER`: an SDL_Window that will be the
*   parent of this window, required for windows with the "tooltip", "menu",
*   and "modal" properties
* - `SDL_PROP_WINDOW_CREATE_RESIZABLE_BOOLEAN`: true if the window should be
*   resizable
* - `SDL_PROP_WINDOW_CREATE_TITLE_STRING`: the title of the window, in UTF-8
*   encoding
* - `SDL_PROP_WINDOW_CREATE_TRANSPARENT_BOOLEAN`: true if the window show
*   transparent in the areas with alpha of 0
* - `SDL_PROP_WINDOW_CREATE_TOOLTIP_BOOLEAN`: true if the window is a tooltip
* - `SDL_PROP_WINDOW_CREATE_UTILITY_BOOLEAN`: true if the window is a utility
*   window, not showing in the task bar and window list
* - `SDL_PROP_WINDOW_CREATE_VULKAN_BOOLEAN`: true if the window will be used
*   with Vulkan rendering
* - `SDL_PROP_WINDOW_CREATE_WIDTH_NUMBER`: the width of the window
* - `SDL_PROP_WINDOW_CREATE_X_NUMBER`: the x position of the window, or
*   `SDL_WINDOWPOS_CENTERED`, defaults to `SDL_WINDOWPOS_UNDEFINED`. This is
*   relative to the parent for windows with the "tooltip" or "menu" property
*   set.
* - `SDL_PROP_WINDOW_CREATE_Y_NUMBER`: the y position of the window, or
*   `SDL_WINDOWPOS_CENTERED`, defaults to `SDL_WINDOWPOS_UNDEFINED`. This is
*   relative to the parent for windows with the "tooltip" or "menu" property
*   set.
*
* These are additional supported properties on macOS:
*
* - `SDL_PROP_WINDOW_CREATE_COCOA_WINDOW_POINTER`: the
*   `(__unsafe_unretained)` NSWindow associated with the window, if you want
*   to wrap an existing window.
* - `SDL_PROP_WINDOW_CREATE_COCOA_VIEW_POINTER`: the `(__unsafe_unretained)`
*   NSView associated with the window, defaults to `[window contentView]`
*
* These are additional supported properties on Wayland:
*
* - `SDL_PROP_WINDOW_CREATE_WAYLAND_SURFACE_ROLE_CUSTOM_BOOLEAN` - true if
*   the application wants to use the Wayland surface for a custom role and
*   does not want it attached to an XDG toplevel window. See
*   [README-wayland](README-wayland) for more information on using custom
*   surfaces.
* - `SDL_PROP_WINDOW_CREATE_WAYLAND_CREATE_EGL_WINDOW_BOOLEAN` - true if the
*   application wants an associated `wl_egl_window` object to be created and
*   attached to the window, even if the window does not have the OpenGL
*   property or `SDL_WINDOW_OPENGL` flag set.
* - `SDL_PROP_WINDOW_CREATE_WAYLAND_WL_SURFACE_POINTER` - the wl_surface
*   associated with the window, if you want to wrap an existing window. See
*   [README-wayland](README-wayland) for more information.
*
* These are additional supported properties on Windows:
*
* - `SDL_PROP_WINDOW_CREATE_WIN32_HWND_POINTER`: the HWND associated with the
*   window, if you want to wrap an existing window.
* - `SDL_PROP_WINDOW_CREATE_WIN32_PIXEL_FORMAT_HWND_POINTER`: optional,
*   another window to share pixel format with, useful for OpenGL windows
*
* These are additional supported properties with X11:
*
* - `SDL_PROP_WINDOW_CREATE_X11_WINDOW_NUMBER`: the X11 Window associated
*   with the window, if you want to wrap an existing window.
*
* The window is implicitly shown if the "hidden" property is not set.
*
* These are additional supported properties with Emscripten:
*
* - `SDL_PROP_WINDOW_CREATE_EMSCRIPTEN_CANVAS_ID_STRING`: the id given to the
*   canvas element. This should start with a '#' sign
* - `SDL_PROP_WINDOW_CREATE_EMSCRIPTEN_KEYBOARD_ELEMENT_STRING`: override the
*   binding element for keyboard inputs for this canvas. The variable can be
*   one of:
* - "#window": the javascript window object (default)
* - "#document": the javascript document object
* - "#screen": the javascript window.screen object
* - "#canvas": the WebGL canvas element
* - "#none": Don't bind anything at all
* - any other string without a leading # sign applies to the element on the
*   page with that ID. Windows with the "tooltip" and "menu" properties are
*   popup windows and have the behaviors and guidelines outlined in
*   SDL_CreatePopupWindow().
*
* If this window is being created to be used with an SDL_Renderer, you should
* not add a graphics API specific property
* (`SDL_PROP_WINDOW_CREATE_OPENGL_BOOLEAN`, etc), as SDL will handle that
* internally when it chooses a renderer. However, SDL might need to recreate
* your window at that point, which may cause the window to appear briefly,
* and then flicker as it is recreated. The correct approach to this is to
* create the window with the `SDL_PROP_WINDOW_CREATE_HIDDEN_BOOLEAN` property
* set to true, then create the renderer, then show the window with
* SDL_ShowWindow().
*
* \param props the properties to use.
* \returns the window that was created or NULL on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_CreateProperties
* \sa SDL_CreateWindow
* \sa SDL_DestroyWindow
 */
//go:sdl3extern
var CreateWindowWithProperties func(props PropertiesID) Window

const (
	PROP_WINDOW_CREATE_ALWAYS_ON_TOP_BOOLEAN               PropertyName = "SDL.window.create.always_on_top"
	PROP_WINDOW_CREATE_BORDERLESS_BOOLEAN                  PropertyName = "SDL.window.create.borderless"
	PROP_WINDOW_CREATE_CONSTRAIN_POPUP_BOOLEAN             PropertyName = "SDL.window.create.constrain_popup"
	PROP_WINDOW_CREATE_FOCUSABLE_BOOLEAN                   PropertyName = "SDL.window.create.focusable"
	PROP_WINDOW_CREATE_EXTERNAL_GRAPHICS_CONTEXT_BOOLEAN   PropertyName = "SDL.window.create.external_graphics_context"
	PROP_WINDOW_CREATE_FLAGS_NUMBER                        PropertyName = "SDL.window.create.flags"
	PROP_WINDOW_CREATE_FULLSCREEN_BOOLEAN                  PropertyName = "SDL.window.create.fullscreen"
	PROP_WINDOW_CREATE_HEIGHT_NUMBER                       PropertyName = "SDL.window.create.height"
	PROP_WINDOW_CREATE_HIDDEN_BOOLEAN                      PropertyName = "SDL.window.create.hidden"
	PROP_WINDOW_CREATE_HIGH_PIXEL_DENSITY_BOOLEAN          PropertyName = "SDL.window.create.high_pixel_density"
	PROP_WINDOW_CREATE_MAXIMIZED_BOOLEAN                   PropertyName = "SDL.window.create.maximized"
	PROP_WINDOW_CREATE_MENU_BOOLEAN                        PropertyName = "SDL.window.create.menu"
	PROP_WINDOW_CREATE_METAL_BOOLEAN                       PropertyName = "SDL.window.create.metal"
	PROP_WINDOW_CREATE_MINIMIZED_BOOLEAN                   PropertyName = "SDL.window.create.minimized"
	PROP_WINDOW_CREATE_MODAL_BOOLEAN                       PropertyName = "SDL.window.create.modal"
	PROP_WINDOW_CREATE_MOUSE_GRABBED_BOOLEAN               PropertyName = "SDL.window.create.mouse_grabbed"
	PROP_WINDOW_CREATE_OPENGL_BOOLEAN                      PropertyName = "SDL.window.create.opengl"
	PROP_WINDOW_CREATE_PARENT_POINTER                      PropertyName = "SDL.window.create.parent"
	PROP_WINDOW_CREATE_RESIZABLE_BOOLEAN                   PropertyName = "SDL.window.create.resizable"
	PROP_WINDOW_CREATE_TITLE_STRING                        PropertyName = "SDL.window.create.title"
	PROP_WINDOW_CREATE_TRANSPARENT_BOOLEAN                 PropertyName = "SDL.window.create.transparent"
	PROP_WINDOW_CREATE_TOOLTIP_BOOLEAN                     PropertyName = "SDL.window.create.tooltip"
	PROP_WINDOW_CREATE_UTILITY_BOOLEAN                     PropertyName = "SDL.window.create.utility"
	PROP_WINDOW_CREATE_VULKAN_BOOLEAN                      PropertyName = "SDL.window.create.vulkan"
	PROP_WINDOW_CREATE_WIDTH_NUMBER                        PropertyName = "SDL.window.create.width"
	PROP_WINDOW_CREATE_X_NUMBER                            PropertyName = "SDL.window.create.x"
	PROP_WINDOW_CREATE_Y_NUMBER                            PropertyName = "SDL.window.create.y"
	PROP_WINDOW_CREATE_COCOA_WINDOW_POINTER                PropertyName = "SDL.window.create.cocoa.window"
	PROP_WINDOW_CREATE_COCOA_VIEW_POINTER                  PropertyName = "SDL.window.create.cocoa.view"
	PROP_WINDOW_CREATE_WAYLAND_SURFACE_ROLE_CUSTOM_BOOLEAN PropertyName = "SDL.window.create.wayland.surface_role_custom"
	PROP_WINDOW_CREATE_WAYLAND_CREATE_EGL_WINDOW_BOOLEAN   PropertyName = "SDL.window.create.wayland.create_egl_window"
	PROP_WINDOW_CREATE_WAYLAND_WL_SURFACE_POINTER          PropertyName = "SDL.window.create.wayland.wl_surface"
	PROP_WINDOW_CREATE_WIN32_HWND_POINTER                  PropertyName = "SDL.window.create.win32.hwnd"
	PROP_WINDOW_CREATE_WIN32_PIXEL_FORMAT_HWND_POINTER     PropertyName = "SDL.window.create.win32.pixel_format_hwnd"
	PROP_WINDOW_CREATE_X11_WINDOW_NUMBER                   PropertyName = "SDL.window.create.x11.window"
	PROP_WINDOW_CREATE_EMSCRIPTEN_CANVAS_ID_STRING         PropertyName = "SDL.window.create.emscripten.canvas_id"
	PROP_WINDOW_CREATE_EMSCRIPTEN_KEYBOARD_ELEMENT_STRING  PropertyName = "SDL.window.create.emscripten.keyboard_element"
)

/**
* Get the numeric ID of a window.
*
* The numeric ID is what SDL_WindowEvent references, and is necessary to map
* these events to specific SDL_Window objects.
*
* \param window the window to query.
* \returns the ID of the window on success or 0 on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowFromID
 */
//go:sdl3extern
var GetWindowID func(window Window) WindowID

/**
* Get a window from a stored ID.
*
* The numeric ID is what SDL_WindowEvent references, and is necessary to map
* these events to specific SDL_Window objects.
*
* \param id the ID of the window.
* \returns the window associated with `id` or NULL if it doesn't exist; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowID
 */
//go:sdl3extern
var GetWindowFromID func(id WindowID) Window

/**
* Get parent of a window.
*
* \param window the window to query.
* \returns the parent of the window on success or NULL if the window has no
*          parent.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_CreatePopupWindow
 */
//go:sdl3extern
var GetWindowParent func(window Window) Window

/**
* Get the properties associated with a window.
*
* The following read-only properties are provided by SDL:
*
* - `SDL_PROP_WINDOW_SHAPE_POINTER`: the surface associated with a shaped
*   window
* - `SDL_PROP_WINDOW_HDR_ENABLED_BOOLEAN`: true if the window has HDR
*   headroom above the SDR white point. This property can change dynamically
*   when SDL_EVENT_WINDOW_HDR_STATE_CHANGED is sent.
* - `SDL_PROP_WINDOW_SDR_WHITE_LEVEL_FLOAT`: the value of SDR white in the
*   SDL_COLORSPACE_SRGB_LINEAR colorspace. On Windows this corresponds to the
*   SDR white level in scRGB colorspace, and on Apple platforms this is
*   always 1.0 for EDR content. This property can change dynamically when
*   SDL_EVENT_WINDOW_HDR_STATE_CHANGED is sent.
* - `SDL_PROP_WINDOW_HDR_HEADROOM_FLOAT`: the additional high dynamic range
*   that can be displayed, in terms of the SDR white point. When HDR is not
*   enabled, this will be 1.0. This property can change dynamically when
*   SDL_EVENT_WINDOW_HDR_STATE_CHANGED is sent.
*
* On Android:
*
* - `SDL_PROP_WINDOW_ANDROID_WINDOW_POINTER`: the ANativeWindow associated
*   with the window
* - `SDL_PROP_WINDOW_ANDROID_SURFACE_POINTER`: the EGLSurface associated with
*   the window
*
* On iOS:
*
* - `SDL_PROP_WINDOW_UIKIT_WINDOW_POINTER`: the `(__unsafe_unretained)`
*   UIWindow associated with the window
* - `SDL_PROP_WINDOW_UIKIT_METAL_VIEW_TAG_NUMBER`: the NSInteger tag
*   associated with metal views on the window
* - `SDL_PROP_WINDOW_UIKIT_OPENGL_FRAMEBUFFER_NUMBER`: the OpenGL view's
*   framebuffer object. It must be bound when rendering to the screen using
*   OpenGL.
* - `SDL_PROP_WINDOW_UIKIT_OPENGL_RENDERBUFFER_NUMBER`: the OpenGL view's
*   renderbuffer object. It must be bound when SDL_GL_SwapWindow is called.
* - `SDL_PROP_WINDOW_UIKIT_OPENGL_RESOLVE_FRAMEBUFFER_NUMBER`: the OpenGL
*   view's resolve framebuffer, when MSAA is used.
*
* On KMS/DRM:
*
* - `SDL_PROP_WINDOW_KMSDRM_DEVICE_INDEX_NUMBER`: the device index associated
*   with the window (e.g. the X in /dev/dri/cardX)
* - `SDL_PROP_WINDOW_KMSDRM_DRM_FD_NUMBER`: the DRM FD associated with the
*   window
* - `SDL_PROP_WINDOW_KMSDRM_GBM_DEVICE_POINTER`: the GBM device associated
*   with the window
*
* On macOS:
*
* - `SDL_PROP_WINDOW_COCOA_WINDOW_POINTER`: the `(__unsafe_unretained)`
*   NSWindow associated with the window
* - `SDL_PROP_WINDOW_COCOA_METAL_VIEW_TAG_NUMBER`: the NSInteger tag
*   assocated with metal views on the window
*
* On OpenVR:
*
* - `SDL_PROP_WINDOW_OPENVR_OVERLAY_ID_NUMBER`: the OpenVR Overlay Handle ID
*   for the associated overlay window.
*
* On Vivante:
*
* - `SDL_PROP_WINDOW_VIVANTE_DISPLAY_POINTER`: the EGLNativeDisplayType
*   associated with the window
* - `SDL_PROP_WINDOW_VIVANTE_WINDOW_POINTER`: the EGLNativeWindowType
*   associated with the window
* - `SDL_PROP_WINDOW_VIVANTE_SURFACE_POINTER`: the EGLSurface associated with
*   the window
*
* On Windows:
*
* - `SDL_PROP_WINDOW_WIN32_HWND_POINTER`: the HWND associated with the window
* - `SDL_PROP_WINDOW_WIN32_HDC_POINTER`: the HDC associated with the window
* - `SDL_PROP_WINDOW_WIN32_INSTANCE_POINTER`: the HINSTANCE associated with
*   the window
*
* On Wayland:
*
* Note: The `xdg_*` window objects do not internally persist across window
* show/hide calls. They will be null if the window is hidden and must be
* queried each time it is shown.
*
* - `SDL_PROP_WINDOW_WAYLAND_DISPLAY_POINTER`: the wl_display associated with
*   the window
* - `SDL_PROP_WINDOW_WAYLAND_SURFACE_POINTER`: the wl_surface associated with
*   the window
* - `SDL_PROP_WINDOW_WAYLAND_VIEWPORT_POINTER`: the wp_viewport associated
*   with the window
* - `SDL_PROP_WINDOW_WAYLAND_EGL_WINDOW_POINTER`: the wl_egl_window
*   associated with the window
* - `SDL_PROP_WINDOW_WAYLAND_XDG_SURFACE_POINTER`: the xdg_surface associated
*   with the window
* - `SDL_PROP_WINDOW_WAYLAND_XDG_TOPLEVEL_POINTER`: the xdg_toplevel role
*   associated with the window
* - 'SDL_PROP_WINDOW_WAYLAND_XDG_TOPLEVEL_EXPORT_HANDLE_STRING': the export
*   handle associated with the window
* - `SDL_PROP_WINDOW_WAYLAND_XDG_POPUP_POINTER`: the xdg_popup role
*   associated with the window
* - `SDL_PROP_WINDOW_WAYLAND_XDG_POSITIONER_POINTER`: the xdg_positioner
*   associated with the window, in popup mode
*
* On X11:
*
* - `SDL_PROP_WINDOW_X11_DISPLAY_POINTER`: the X11 Display associated with
*   the window
* - `SDL_PROP_WINDOW_X11_SCREEN_NUMBER`: the screen number associated with
*   the window
* - `SDL_PROP_WINDOW_X11_WINDOW_NUMBER`: the X11 Window associated with the
*   window
*
* On Emscripten:
*
* - `SDL_PROP_WINDOW_EMSCRIPTEN_CANVAS_ID_STRING`: the id the canvas element
*   will have
* - `SDL_PROP_WINDOW_EMSCRIPTEN_KEYBOARD_ELEMENT_STRING`: the keyboard
*   element that associates keyboard events to this window
*
* \param window the window to query.
* \returns a valid property ID on success or 0 on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetWindowProperties func(window Window) PropertiesID

const (
	PROP_WINDOW_SHAPE_POINTER                             PropertyName = "SDL.window.shape"
	PROP_WINDOW_HDR_ENABLED_BOOLEAN                       PropertyName = "SDL.window.HDR_enabled"
	PROP_WINDOW_SDR_WHITE_LEVEL_FLOAT                     PropertyName = "SDL.window.SDR_white_level"
	PROP_WINDOW_HDR_HEADROOM_FLOAT                        PropertyName = "SDL.window.HDR_headroom"
	PROP_WINDOW_ANDROID_WINDOW_POINTER                    PropertyName = "SDL.window.android.window"
	PROP_WINDOW_ANDROID_SURFACE_POINTER                   PropertyName = "SDL.window.android.surface"
	PROP_WINDOW_UIKIT_WINDOW_POINTER                      PropertyName = "SDL.window.uikit.window"
	PROP_WINDOW_UIKIT_METAL_VIEW_TAG_NUMBER               PropertyName = "SDL.window.uikit.metal_view_tag"
	PROP_WINDOW_UIKIT_OPENGL_FRAMEBUFFER_NUMBER           PropertyName = "SDL.window.uikit.opengl.framebuffer"
	PROP_WINDOW_UIKIT_OPENGL_RENDERBUFFER_NUMBER          PropertyName = "SDL.window.uikit.opengl.renderbuffer"
	PROP_WINDOW_UIKIT_OPENGL_RESOLVE_FRAMEBUFFER_NUMBER   PropertyName = "SDL.window.uikit.opengl.resolve_framebuffer"
	PROP_WINDOW_KMSDRM_DEVICE_INDEX_NUMBER                PropertyName = "SDL.window.kmsdrm.dev_index"
	PROP_WINDOW_KMSDRM_DRM_FD_NUMBER                      PropertyName = "SDL.window.kmsdrm.drm_fd"
	PROP_WINDOW_KMSDRM_GBM_DEVICE_POINTER                 PropertyName = "SDL.window.kmsdrm.gbm_dev"
	PROP_WINDOW_COCOA_WINDOW_POINTER                      PropertyName = "SDL.window.cocoa.window"
	PROP_WINDOW_COCOA_METAL_VIEW_TAG_NUMBER               PropertyName = "SDL.window.cocoa.metal_view_tag"
	PROP_WINDOW_OPENVR_OVERLAY_ID_NUMBER                  PropertyName = "SDL.window.openvr.overlay_id"
	PROP_WINDOW_VIVANTE_DISPLAY_POINTER                   PropertyName = "SDL.window.vivante.display"
	PROP_WINDOW_VIVANTE_WINDOW_POINTER                    PropertyName = "SDL.window.vivante.window"
	PROP_WINDOW_VIVANTE_SURFACE_POINTER                   PropertyName = "SDL.window.vivante.surface"
	PROP_WINDOW_WIN32_HWND_POINTER                        PropertyName = "SDL.window.win32.hwnd"
	PROP_WINDOW_WIN32_HDC_POINTER                         PropertyName = "SDL.window.win32.hdc"
	PROP_WINDOW_WIN32_INSTANCE_POINTER                    PropertyName = "SDL.window.win32.instance"
	PROP_WINDOW_WAYLAND_DISPLAY_POINTER                   PropertyName = "SDL.window.wayland.display"
	PROP_WINDOW_WAYLAND_SURFACE_POINTER                   PropertyName = "SDL.window.wayland.surface"
	PROP_WINDOW_WAYLAND_VIEWPORT_POINTER                  PropertyName = "SDL.window.wayland.viewport"
	PROP_WINDOW_WAYLAND_EGL_WINDOW_POINTER                PropertyName = "SDL.window.wayland.egl_window"
	PROP_WINDOW_WAYLAND_XDG_SURFACE_POINTER               PropertyName = "SDL.window.wayland.xdg_surface"
	PROP_WINDOW_WAYLAND_XDG_TOPLEVEL_POINTER              PropertyName = "SDL.window.wayland.xdg_toplevel"
	PROP_WINDOW_WAYLAND_XDG_TOPLEVEL_EXPORT_HANDLE_STRING PropertyName = "SDL.window.wayland.xdg_toplevel_export_handle"
	PROP_WINDOW_WAYLAND_XDG_POPUP_POINTER                 PropertyName = "SDL.window.wayland.xdg_popup"
	PROP_WINDOW_WAYLAND_XDG_POSITIONER_POINTER            PropertyName = "SDL.window.wayland.xdg_positioner"
	PROP_WINDOW_X11_DISPLAY_POINTER                       PropertyName = "SDL.window.x11.display"
	PROP_WINDOW_X11_SCREEN_NUMBER                         PropertyName = "SDL.window.x11.screen"
	PROP_WINDOW_X11_WINDOW_NUMBER                         PropertyName = "SDL.window.x11.window"
	PROP_WINDOW_EMSCRIPTEN_CANVAS_ID_STRING               PropertyName = "SDL.window.emscripten.canvas_id"
	PROP_WINDOW_EMSCRIPTEN_KEYBOARD_ELEMENT_STRING        PropertyName = "SDL.window.emscripten.keyboard_element"
)

/**
* Get the window flags.
*
* \param window the window to query.
* \returns a mask of the SDL_WindowFlags associated with `window`.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_CreateWindow
* \sa SDL_HideWindow
* \sa SDL_MaximizeWindow
* \sa SDL_MinimizeWindow
* \sa SDL_SetWindowFullscreen
* \sa SDL_SetWindowMouseGrab
* \sa SDL_ShowWindow
 */
//go:sdl3extern
var GetWindowFlags func(window Window) WindowFlags

/**
* Set the title of a window.
*
* This string is expected to be in UTF-8 encoding.
*
* \param window the window to change.
* \param title the desired window title in UTF-8 format.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowTitle
 */
//go:sdl3extern
var SetWindowTitle func(window Window, title *byte) bool

/**
* Get the title of a window.
*
* \param window the window to query.
* \returns the title of the window in UTF-8 format or "" if there is no
*          title.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowTitle
 */
//go:sdl3extern
var GetWindowTitle func(window Window) *byte

/**
* Set the icon for a window.
*
* If this function is passed a surface with alternate representations, the
* surface will be interpreted as the content to be used for 100% display
* scale, and the alternate representations will be used for high DPI
* situations. For example, if the original surface is 32x32, then on a 2x
* macOS display or 200% display scale on Windows, a 64x64 version of the
* image will be used, if available. If a matching version of the image isn't
* available, the closest larger size image will be downscaled to the
* appropriate size and be used instead, if available. Otherwise, the closest
* smaller image will be upscaled and be used instead.
*
* \param window the window to change.
* \param icon an SDL_Surface structure containing the icon for the window.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetWindowIcon func(window Window, icon *Surface) bool

/**
* Request that the window's position be set.
*
* If the window is in an exclusive fullscreen or maximized state, this
* request has no effect.
*
* This can be used to reposition fullscreen-desktop windows onto a different
* display, however, as exclusive fullscreen windows are locked to a specific
* display, they can only be repositioned programmatically via
* SDL_SetWindowFullscreenMode().
*
* On some windowing systems this request is asynchronous and the new
* coordinates may not have have been applied immediately upon the return of
* this function. If an immediate change is required, call SDL_SyncWindow() to
* block until the changes have taken effect.
*
* When the window position changes, an SDL_EVENT_WINDOW_MOVED event will be
* emitted with the window's new coordinates. Note that the new coordinates
* may not match the exact coordinates requested, as some windowing systems
* can restrict the position of the window in certain scenarios (e.g.
* constraining the position so the window is always within desktop bounds).
* Additionally, as this is just a request, it can be denied by the windowing
* system.
*
* \param window the window to reposition.
* \param x the x coordinate of the window, or `SDL_WINDOWPOS_CENTERED` or
*          `SDL_WINDOWPOS_UNDEFINED`.
* \param y the y coordinate of the window, or `SDL_WINDOWPOS_CENTERED` or
*          `SDL_WINDOWPOS_UNDEFINED`.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowPosition
* \sa SDL_SyncWindow
 */
//go:sdl3extern
var SetWindowPosition func(window Window, x, y int32) bool

/**
* Get the position of a window.
*
* This is the current position of the window as last reported by the
* windowing system.
*
* If you do not need the value for one of the positions a NULL may be passed
* in the `x` or `y` parameter.
*
* \param window the window to query.
* \param x a pointer filled in with the x position of the window, may be
*          NULL.
* \param y a pointer filled in with the y position of the window, may be
*          NULL.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowPosition
 */
//go:sdl3extern
var GetWindowPosition func(window Window, x, y *int32) bool

/**
* Request that the size of a window's client area be set.
*
* If the window is in a fullscreen or maximized state, this request has no
* effect.
*
* To change the exclusive fullscreen mode of a window, use
* SDL_SetWindowFullscreenMode().
*
* On some windowing systems, this request is asynchronous and the new window
* size may not have have been applied immediately upon the return of this
* function. If an immediate change is required, call SDL_SyncWindow() to
* block until the changes have taken effect.
*
* When the window size changes, an SDL_EVENT_WINDOW_RESIZED event will be
* emitted with the new window dimensions. Note that the new dimensions may
* not match the exact size requested, as some windowing systems can restrict
* the window size in certain scenarios (e.g. constraining the size of the
* content area to remain within the usable desktop bounds). Additionally, as
* this is just a request, it can be denied by the windowing system.
*
* \param window the window to change.
* \param w the width of the window, must be > 0.
* \param h the height of the window, must be > 0.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowSize
* \sa SDL_SetWindowFullscreenMode
* \sa SDL_SyncWindow
 */
//go:sdl3extern
var SetWindowSize func(window Window, w, h int32) bool

/**
* Get the size of a window's client area.
*
* The window pixel size may differ from its window coordinate size if the
* window is on a high pixel density display. Use SDL_GetWindowSizeInPixels()
* or SDL_GetRenderOutputSize() to get the real client area size in pixels.
*
* \param window the window to query the width and height from.
* \param w a pointer filled in with the width of the window, may be NULL.
* \param h a pointer filled in with the height of the window, may be NULL.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetRenderOutputSize
* \sa SDL_GetWindowSizeInPixels
* \sa SDL_SetWindowSize
 */
//go:sdl3extern
var GetWindowSize func(window Window, w, h *int32) bool

/**
* Get the safe area for this window.
*
* Some devices have portions of the screen which are partially obscured or
* not interactive, possibly due to on-screen controls, curved edges, camera
* notches, TV overscan, etc. This function provides the area of the window
* which is safe to have interactable content. You should continue rendering
* into the rest of the window, but it should not contain visually important
* or interactible content.
*
* \param window the window to query.
* \param rect a pointer filled in with the client area that is safe for
*             interactive content.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetWindowSafeArea func(window Window, SDL_Rect *Rect) bool

/**
* Request that the aspect ratio of a window's client area be set.
*
* The aspect ratio is the ratio of width divided by height, e.g. 2560x1600
* would be 1.6. Larger aspect ratios are wider and smaller aspect ratios are
* narrower.
*
* If, at the time of this request, the window in a fixed-size state, such as
* maximized or fullscreen, the request will be deferred until the window
* exits this state and becomes resizable again.
*
* On some windowing systems, this request is asynchronous and the new window
* aspect ratio may not have have been applied immediately upon the return of
* this function. If an immediate change is required, call SDL_SyncWindow() to
* block until the changes have taken effect.
*
* When the window size changes, an SDL_EVENT_WINDOW_RESIZED event will be
* emitted with the new window dimensions. Note that the new dimensions may
* not match the exact aspect ratio requested, as some windowing systems can
* restrict the window size in certain scenarios (e.g. constraining the size
* of the content area to remain within the usable desktop bounds).
* Additionally, as this is just a request, it can be denied by the windowing
* system.
*
* \param window the window to change.
* \param min_aspect the minimum aspect ratio of the window, or 0.0f for no
*                   limit.
* \param max_aspect the maximum aspect ratio of the window, or 0.0f for no
*                   limit.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowAspectRatio
* \sa SDL_SyncWindow
 */
//go:sdl3extern
var SetWindowAspectRatio func(window Window, min_aspect, max_aspect float32) bool

/**
* Get the size of a window's client area.
*
* \param window the window to query the width and height from.
* \param min_aspect a pointer filled in with the minimum aspect ratio of the
*                   window, may be NULL.
* \param max_aspect a pointer filled in with the maximum aspect ratio of the
*                   window, may be NULL.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowAspectRatio
 */
//go:sdl3extern
var GetWindowAspectRatio func(window Window, min_aspect, max_aspect *float32) bool

/**
* Get the size of a window's borders (decorations) around the client area.
*
* Note: If this function fails (returns false), the size values will be
* initialized to 0, 0, 0, 0 (if a non-NULL pointer is provided), as if the
* window in question was borderless.
*
* Note: This function may fail on systems where the window has not yet been
* decorated by the display server (for example, immediately after calling
* SDL_CreateWindow). It is recommended that you wait at least until the
* window has been presented and composited, so that the window system has a
* chance to decorate the window and provide the border dimensions to SDL.
*
* This function also returns false if getting the information is not
* supported.
*
* \param window the window to query the size values of the border
*               (decorations) from.
* \param top pointer to variable for storing the size of the top border; NULL
*            is permitted.
* \param left pointer to variable for storing the size of the left border;
*             NULL is permitted.
* \param bottom pointer to variable for storing the size of the bottom
*               border; NULL is permitted.
* \param right pointer to variable for storing the size of the right border;
*              NULL is permitted.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowSize
 */
//go:sdl3extern
var GetWindowBordersSize func(window Window, top, left, bottom, right *int32) bool

/**
* Get the size of a window's client area, in pixels.
*
* \param window the window from which the drawable size should be queried.
* \param w a pointer to variable for storing the width in pixels, may be
*          NULL.
* \param h a pointer to variable for storing the height in pixels, may be
*          NULL.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_CreateWindow
* \sa SDL_GetWindowSize
 */
//go:sdl3extern
var GetWindowSizeInPixels func(window Window, w, h *int32) bool

/**
* Set the minimum size of a window's client area.
*
* \param window the window to change.
* \param min_w the minimum width of the window, or 0 for no limit.
* \param min_h the minimum height of the window, or 0 for no limit.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowMinimumSize
* \sa SDL_SetWindowMaximumSize
 */
//go:sdl3extern
var SetWindowMinimumSize func(window Window, min_w, min_h int32) bool

/**
* Get the minimum size of a window's client area.
*
* \param window the window to query.
* \param w a pointer filled in with the minimum width of the window, may be
*          NULL.
* \param h a pointer filled in with the minimum height of the window, may be
*          NULL.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowMaximumSize
* \sa SDL_SetWindowMinimumSize
 */
//go:sdl3extern
var GetWindowMinimumSize func(window Window, w, h *int32) bool

/**
* Set the maximum size of a window's client area.
*
* \param window the window to change.
* \param max_w the maximum width of the window, or 0 for no limit.
* \param max_h the maximum height of the window, or 0 for no limit.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowMaximumSize
* \sa SDL_SetWindowMinimumSize
 */
//go:sdl3extern
var SetWindowMaximumSize func(window Window, max_w, max_h int32) bool

/**
* Get the maximum size of a window's client area.
*
* \param window the window to query.
* \param w a pointer filled in with the maximum width of the window, may be
*          NULL.
* \param h a pointer filled in with the maximum height of the window, may be
*          NULL.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowMinimumSize
* \sa SDL_SetWindowMaximumSize
 */
//go:sdl3extern
var GetWindowMaximumSize func(window Window, w, h *int32) bool

/**
* Set the border state of a window.
*
* This will add or remove the window's `SDL_WINDOW_BORDERLESS` flag and add
* or remove the border from the actual window. This is a no-op if the
* window's border already matches the requested state.
*
* You can't change the border state of a fullscreen window.
*
* \param window the window of which to change the border state.
* \param bordered false to remove border, true to add border.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowFlags
 */
//go:sdl3extern
var SetWindowBordered func(window Window, bordered bool) bool

/**
* Set the user-resizable state of a window.
*
* This will add or remove the window's `SDL_WINDOW_RESIZABLE` flag and
* allow/disallow user resizing of the window. This is a no-op if the window's
* resizable state already matches the requested state.
*
* You can't change the resizable state of a fullscreen window.
*
* \param window the window of which to change the resizable state.
* \param resizable true to allow resizing, false to disallow.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowFlags
 */
//go:sdl3extern
var SetWindowResizable func(window Window, resizable bool) bool

/**
* Set the window to always be above the others.
*
* This will add or remove the window's `SDL_WINDOW_ALWAYS_ON_TOP` flag. This
* will bring the window to the front and keep the window above the rest.
*
* \param window the window of which to change the always on top state.
* \param on_top true to set the window always on top, false to disable.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowFlags
 */
//go:sdl3extern
var SetWindowAlwaysOnTop func(window Window, on_top bool) bool

/**
* Show a window.
*
* \param window the window to show.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_HideWindow
* \sa SDL_RaiseWindow
 */
//go:sdl3extern
var ShowWindow func(window Window) bool

/**
* Hide a window.
*
* \param window the window to hide.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_ShowWindow
* \sa SDL_WINDOW_HIDDEN
 */
//go:sdl3extern
var HideWindow func(window Window) bool

/**
* Request that a window be raised above other windows and gain the input
* focus.
*
* The result of this request is subject to desktop window manager policy,
* particularly if raising the requested window would result in stealing focus
* from another application. If the window is successfully raised and gains
* input focus, an SDL_EVENT_WINDOW_FOCUS_GAINED event will be emitted, and
* the window will have the SDL_WINDOW_INPUT_FOCUS flag set.
*
* \param window the window to raise.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var RaiseWindow func(window Window) bool

/**
* Request that the window be made as large as possible.
*
* Non-resizable windows can't be maximized. The window must have the
* SDL_WINDOW_RESIZABLE flag set, or this will have no effect.
*
* On some windowing systems this request is asynchronous and the new window
* state may not have have been applied immediately upon the return of this
* function. If an immediate change is required, call SDL_SyncWindow() to
* block until the changes have taken effect.
*
* When the window state changes, an SDL_EVENT_WINDOW_MAXIMIZED event will be
* emitted. Note that, as this is just a request, the windowing system can
* deny the state change.
*
* When maximizing a window, whether the constraints set via
* SDL_SetWindowMaximumSize() are honored depends on the policy of the window
* manager. Win32 and macOS enforce the constraints when maximizing, while X11
* and Wayland window managers may vary.
*
* \param window the window to maximize.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_MinimizeWindow
* \sa SDL_RestoreWindow
* \sa SDL_SyncWindow
 */
//go:sdl3extern
var MaximizeWindow func(window Window) bool

/**
* Request that the window be minimized to an iconic representation.
*
* If the window is in a fullscreen state, this request has no direct effect.
* It may alter the state the window is returned to when leaving fullscreen.
*
* On some windowing systems this request is asynchronous and the new window
* state may not have been applied immediately upon the return of this
* function. If an immediate change is required, call SDL_SyncWindow() to
* block until the changes have taken effect.
*
* When the window state changes, an SDL_EVENT_WINDOW_MINIMIZED event will be
* emitted. Note that, as this is just a request, the windowing system can
* deny the state change.
*
* \param window the window to minimize.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_MaximizeWindow
* \sa SDL_RestoreWindow
* \sa SDL_SyncWindow
 */
//go:sdl3extern
var MinimizeWindow func(window Window) bool

/**
* Request that the size and position of a minimized or maximized window be
* restored.
*
* If the window is in a fullscreen state, this request has no direct effect.
* It may alter the state the window is returned to when leaving fullscreen.
*
* On some windowing systems this request is asynchronous and the new window
* state may not have have been applied immediately upon the return of this
* function. If an immediate change is required, call SDL_SyncWindow() to
* block until the changes have taken effect.
*
* When the window state changes, an SDL_EVENT_WINDOW_RESTORED event will be
* emitted. Note that, as this is just a request, the windowing system can
* deny the state change.
*
* \param window the window to restore.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_MaximizeWindow
* \sa SDL_MinimizeWindow
* \sa SDL_SyncWindow
 */
//go:sdl3extern
var RestoreWindow func(window Window) bool

/**
* Request that the window's fullscreen state be changed.
*
* By default a window in fullscreen state uses borderless fullscreen desktop
* mode, but a specific exclusive display mode can be set using
* SDL_SetWindowFullscreenMode().
*
* On some windowing systems this request is asynchronous and the new
* fullscreen state may not have have been applied immediately upon the return
* of this function. If an immediate change is required, call SDL_SyncWindow()
* to block until the changes have taken effect.
*
* When the window state changes, an SDL_EVENT_WINDOW_ENTER_FULLSCREEN or
* SDL_EVENT_WINDOW_LEAVE_FULLSCREEN event will be emitted. Note that, as this
* is just a request, it can be denied by the windowing system.
*
* \param window the window to change.
* \param fullscreen true for fullscreen mode, false for windowed mode.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowFullscreenMode
* \sa SDL_SetWindowFullscreenMode
* \sa SDL_SyncWindow
* \sa SDL_WINDOW_FULLSCREEN
 */
//go:sdl3extern
var SetWindowFullscreen func(window Window, fullscreen bool) bool

/**
* Block until any pending window state is finalized.
*
* On asynchronous windowing systems, this acts as a synchronization barrier
* for pending window state. It will attempt to wait until any pending window
* state has been applied and is guaranteed to return within finite time. Note
* that for how long it can potentially block depends on the underlying window
* system, as window state changes may involve somewhat lengthy animations
* that must complete before the window is in its final requested state.
*
* On windowing systems where changes are immediate, this does nothing.
*
* \param window the window for which to wait for the pending state to be
*               applied.
* \returns true on success or false if the operation timed out before the
*          window was in the requested state.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowSize
* \sa SDL_SetWindowPosition
* \sa SDL_SetWindowFullscreen
* \sa SDL_MinimizeWindow
* \sa SDL_MaximizeWindow
* \sa SDL_RestoreWindow
* \sa SDL_HINT_VIDEO_SYNC_WINDOW_OPERATIONS
 */
//go:sdl3extern
var SyncWindow func(window Window) bool

/**
* Return whether the window has a surface associated with it.
*
* \param window the window to query.
* \returns true if there is a surface associated with the window, or false
*          otherwise.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowSurface
 */
//go:sdl3extern
var WindowHasSurface func(window Window) bool

/**
* Get the SDL surface associated with the window.
*
* A new surface will be created with the optimal format for the window, if
* necessary. This surface will be freed when the window is destroyed. Do not
* free this surface.
*
* This surface will be invalidated if the window is resized. After resizing a
* window this function must be called again to return a valid surface.
*
* You may not combine this with 3D or the rendering API on this window.
*
* This function is affected by `SDL_HINT_FRAMEBUFFER_ACCELERATION`.
*
* \param window the window to query.
* \returns the surface associated with the window, or NULL on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_DestroyWindowSurface
* \sa SDL_WindowHasSurface
* \sa SDL_UpdateWindowSurface
* \sa SDL_UpdateWindowSurfaceRects
 */
//go:sdl3extern
var GetWindowSurface func(window Window) *Surface

/**
* Toggle VSync for the window surface.
*
* When a window surface is created, vsync defaults to
* SDL_WINDOW_SURFACE_VSYNC_DISABLED.
*
* The `vsync` parameter can be 1 to synchronize present with every vertical
* refresh, 2 to synchronize present with every second vertical refresh, etc.,
* SDL_WINDOW_SURFACE_VSYNC_ADAPTIVE for late swap tearing (adaptive vsync),
* or SDL_WINDOW_SURFACE_VSYNC_DISABLED to disable. Not every value is
* supported by every driver, so you should check the return value to see
* whether the requested setting is supported.
*
* \param window the window.
* \param vsync the vertical refresh sync interval.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowSurfaceVSync
 */
//go:sdl3extern
var SetWindowSurfaceVSync func(window Window, vsync VSync) bool

/**
* Get VSync for the window surface.
*
* \param window the window to query.
* \param vsync an int filled with the current vertical refresh sync interval.
*              See SDL_SetWindowSurfaceVSync() for the meaning of the value.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowSurfaceVSync
 */
//go:sdl3extern
var GetWindowSurfaceVSync func(window Window, vsync *VSync) bool

/**
* Copy the window surface to the screen.
*
* This is the function you use to reflect any changes to the surface on the
* screen.
*
* This function is equivalent to the SDL 1.2 API SDL_Flip().
*
* \param window the window to update.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowSurface
* \sa SDL_UpdateWindowSurfaceRects
 */
//go:sdl3extern
var UpdateWindowSurface func(window Window) bool

/**
* Copy areas of the window surface to the screen.
*
* This is the function you use to reflect changes to portions of the surface
* on the screen.
*
* This function is equivalent to the SDL 1.2 API SDL_UpdateRects().
*
* Note that this function will update _at least_ the rectangles specified,
* but this is only intended as an optimization; in practice, this might
* update more of the screen (or all of the screen!), depending on what method
* SDL uses to send pixels to the system.
*
* \param window the window to update.
* \param rects an array of SDL_Rect structures representing areas of the
*              surface to copy, in pixels.
* \param numrects the number of rectangles.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowSurface
* \sa SDL_UpdateWindowSurface
 */
//go:sdl3extern
var UpdateWindowSurfaceRects func(window Window, rects *Rect, numrects int32) bool

/**
* Destroy the surface associated with the window.
*
* \param window the window to update.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowSurface
* \sa SDL_WindowHasSurface
 */
//go:sdl3extern
var DestroyWindowSurface func(window Window) bool

/**
* Set a window's keyboard grab mode.
*
* Keyboard grab enables capture of system keyboard shortcuts like Alt+Tab or
* the Meta/Super key. Note that not all system keyboard shortcuts can be
* captured by applications (one example is Ctrl+Alt+Del on Windows).
*
* This is primarily intended for specialized applications such as VNC clients
* or VM frontends. Normal games should not use keyboard grab.
*
* When keyboard grab is enabled, SDL will continue to handle Alt+Tab when the
* window is full-screen to ensure the user is not trapped in your
* application. If you have a custom keyboard shortcut to exit fullscreen
* mode, you may suppress this behavior with
* `SDL_HINT_ALLOW_ALT_TAB_WHILE_GRABBED`.
*
* If the caller enables a grab while another window is currently grabbed, the
* other window loses its grab in favor of the caller's window.
*
* \param window the window for which the keyboard grab mode should be set.
* \param grabbed this is true to grab keyboard, and false to release.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowKeyboardGrab
* \sa SDL_SetWindowMouseGrab
 */
//go:sdl3extern
var SetWindowKeyboardGrab func(window Window, grabbed bool) bool

/**
* Set a window's mouse grab mode.
*
* Mouse grab confines the mouse cursor to the window.
*
* \param window the window for which the mouse grab mode should be set.
* \param grabbed this is true to grab mouse, and false to release.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowMouseRect
* \sa SDL_SetWindowMouseRect
* \sa SDL_SetWindowKeyboardGrab
 */
//go:sdl3extern
var SetWindowMouseGrab func(window Window, grabbed bool) bool

/**
* Get a window's keyboard grab mode.
*
* \param window the window to query.
* \returns true if keyboard is grabbed, and false otherwise.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowKeyboardGrab
 */
//go:sdl3extern
var GetWindowKeyboardGrab func(window Window) bool

/**
* Get a window's mouse grab mode.
*
* \param window the window to query.
* \returns true if mouse is grabbed, and false otherwise.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowMouseRect
* \sa SDL_SetWindowMouseRect
* \sa SDL_SetWindowMouseGrab
* \sa SDL_SetWindowKeyboardGrab
 */
//go:sdl3extern
var GetWindowMouseGrab func(window Window) bool

/**
* Get the window that currently has an input grab enabled.
*
* \returns the window if input is grabbed or NULL otherwise.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowMouseGrab
* \sa SDL_SetWindowKeyboardGrab
 */
//go:sdl3extern
var GetGrabbedWindow func() Window

/**
* Confines the cursor to the specified area of a window.
*
* Note that this does NOT grab the cursor, it only defines the area a cursor
* is restricted to when the window has mouse focus.
*
* \param window the window that will be associated with the barrier.
* \param rect a rectangle area in window-relative coordinates. If NULL the
*             barrier for the specified window will be destroyed.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowMouseRect
* \sa SDL_GetWindowMouseGrab
* \sa SDL_SetWindowMouseGrab
 */
//go:sdl3extern
var SetWindowMouseRect func(window Window, rect *Rect) bool

/**
* Get the mouse confinement rectangle of a window.
*
* \param window the window to query.
* \returns a pointer to the mouse confinement rectangle of a window, or NULL
*          if there isn't one.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowMouseRect
* \sa SDL_GetWindowMouseGrab
* \sa SDL_SetWindowMouseGrab
 */
//go:sdl3extern
var GetWindowMouseRect func(window Window) *Rect

/**
* Set the opacity for a window.
*
* The parameter `opacity` will be clamped internally between 0.0f
* (transparent) and 1.0f (opaque).
*
* This function also returns false if setting the opacity isn't supported.
*
* \param window the window which will be made transparent or opaque.
* \param opacity the opacity value (0.0f - transparent, 1.0f - opaque).
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetWindowOpacity
 */
//go:sdl3extern
var SetWindowOpacity func(window Window, opacity float32) bool

/**
* Get the opacity of a window.
*
* If transparency isn't supported on this platform, opacity will be returned
* as 1.0f without error.
*
* \param window the window to get the current opacity value from.
* \returns the opacity, (0.0f - transparent, 1.0f - opaque), or -1.0f on
*          failure; call SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowOpacity
 */
//go:sdl3extern
var GetWindowOpacity func(window Window) float32

/**
* Set the window as a child of a parent window.
*
* If the window is already the child of an existing window, it will be
* reparented to the new owner. Setting the parent window to NULL unparents
* the window and removes child window status.
*
* If a parent window is hidden or destroyed, the operation will be
* recursively applied to child windows. Child windows hidden with the parent
* that did not have their hidden status explicitly set will be restored when
* the parent is shown.
*
* Attempting to set the parent of a window that is currently in the modal
* state will fail. Use SDL_SetWindowModal() to cancel the modal status before
* attempting to change the parent.
*
* Popup windows cannot change parents and attempts to do so will fail.
*
* Setting a parent window that is currently the sibling or descendent of the
* child window results in undefined behavior.
*
* \param window the window that should become the child of a parent.
* \param parent the new parent window for the child window.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowModal
 */
//go:sdl3extern
var SetWindowParent func(window Window, parent Window) bool

/**
* Toggle the state of the window as modal.
*
* To enable modal status on a window, the window must currently be the child
* window of a parent, or toggling modal status on will fail.
*
* \param window the window on which to set the modal state.
* \param modal true to toggle modal status on, false to toggle it off.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetWindowParent
* \sa SDL_WINDOW_MODAL
 */
//go:sdl3extern
var SetWindowModal func(window Window, modal bool) bool

/**
* Set whether the window may have input focus.
*
* \param window the window to set focusable state.
* \param focusable true to allow input focus, false to not allow input focus.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetWindowFocusable func(window Window, focusable bool) bool

/**
* Display the system-level window menu.
*
* This default window menu is provided by the system and on some platforms
* provides functionality for setting or changing privileged state on the
* window, such as moving it between workspaces or displays, or toggling the
* always-on-top property.
*
* On platforms or desktops where this is unsupported, this function does
* nothing.
*
* \param window the window for which the menu will be displayed.
* \param x the x coordinate of the menu, relative to the origin (top-left) of
*          the client area.
* \param y the y coordinate of the menu, relative to the origin (top-left) of
*          the client area.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var ShowWindowSystemMenu func(window Window, x, y int32) bool

/**
* Possible return values from the SDL_HitTest callback.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This enum is available since SDL 3.2.0.
*
* \sa SDL_HitTest
 */
type HitTestResult int32

const (
	HTR_Normal            HitTestResult = iota /**< Region is normal. No special properties. */
	HTR_Draggable                              /**< Region can drag entire window. */
	HTR_ResizeTopleft                          /**< Region is the resizable top-left corner border. */
	HTR_ResizeTop                              /**< Region is the resizable top border. */
	HTR_ResizeTopright                         /**< Region is the resizable top-right corner border. */
	HTR_ResizeRight                            /**< Region is the resizable right border. */
	HTR_ResizeBottomright                      /**< Region is the resizable bottom-right corner border. */
	HTR_ResizeBottom                           /**< Region is the resizable bottom border. */
	HTR_ResizeBottomleft                       /**< Region is the resizable bottom-left corner border. */
	HTR_ResizeLeft                             /**< Region is the resizable left border. */
)

/**
* Callback used for hit-testing.
*
* \param win the SDL_Window where hit-testing was set on.
* \param area an SDL_Point which should be hit-tested.
* \param data what was passed as `callback_data` to SDL_SetWindowHitTest().
* \returns an SDL_HitTestResult value.
*
* \sa SDL_SetWindowHitTest
 */
type HitTest func(win Window, area *Point, data uintptr) HitTestResult

/**
* Provide a callback that decides if a window region has special properties.
*
* Normally windows are dragged and resized by decorations provided by the
* system window manager (a title bar, borders, etc), but for some apps, it
* makes sense to drag them from somewhere else inside the window itself; for
* example, one might have a borderless window that wants to be draggable from
* any part, or simulate its own title bar, etc.
*
* This function lets the app provide a callback that designates pieces of a
* given window as special. This callback is run during event processing if we
* need to tell the OS to treat a region of the window specially; the use of
* this callback is known as "hit testing."
*
* Mouse input may not be delivered to your application if it is within a
* special area; the OS will often apply that input to moving the window or
* resizing the window and not deliver it to the application.
*
* Specifying NULL for a callback disables hit-testing. Hit-testing is
* disabled by default.
*
* Platforms that don't support this functionality will return false
* unconditionally, even if you're attempting to disable hit-testing.
*
* Your callback may fire at any time, and its firing does not indicate any
* specific behavior (for example, on Windows, this certainly might fire when
* the OS is deciding whether to drag your window, but it fires for lots of
* other reasons, too, some unrelated to anything you probably care about _and
* when the mouse isn't actually at the location it is testing_). Since this
* can fire at any time, you should try to keep your callback efficient,
* devoid of allocations, etc.
*
* \param window the window to set hit-testing on.
* \param callback the function to call when doing a hit-test.
* \param callback_data an app-defined void pointer passed to **callback**.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetWindowHitTest func(window Window, callback HitTest, callback_data uintptr) bool

/**
* Set the shape of a transparent window.
*
* This sets the alpha channel of a transparent window and any fully
* transparent areas are also transparent to mouse clicks. If you are using
* something besides the SDL render API, then you are responsible for drawing
* the alpha channel of the window to match the shape alpha channel to get
* consistent cross-platform results.
*
* The shape is copied inside this function, so you can free it afterwards. If
* your shape surface changes, you should call SDL_SetWindowShape() again to
* update the window. This is an expensive operation, so should be done
* sparingly.
*
* The window must have been created with the SDL_WINDOW_TRANSPARENT flag.
*
* \param window the window.
* \param shape the surface representing the shape of the window, or NULL to
*              remove any current shape.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetWindowShape func(window Window, shape *Surface) bool

/**
* Request a window to demand attention from the user.
*
* \param window the window to be flashed.
* \param operation the operation to perform.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var FlashWindow func(window Window, operation FlashOperation) bool

/**
* Sets the state of the progress bar for the given window’s taskbar icon.
*
* \param window the window whose progress state is to be modified.
* \param state the progress state. `SDL_PROGRESS_STATE_NONE` stops displaying
*              the progress bar.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.4.0.
 */
//go:sdl3extern
var SetWindowProgressState func(window Window, state ProgressState) bool

/**
* Get the state of the progress bar for the given window’s taskbar icon.
*
* \param window the window to get the current progress state from.
* \returns the progress state, or `SDL_PROGRESS_STATE_INVALID` on failure;
*          call SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.4.0.
 */
//go:sdl3extern
var GetWindowProgressState func(window Window) ProgressState

/**
* Sets the value of the progress bar for the given window’s taskbar icon.
*
* \param window the window whose progress value is to be modified.
* \param value the progress value in the range of [0.0f - 1.0f]. If the value
*              is outside the valid range, it gets clamped.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.4.0.
 */
//go:sdl3extern
var SetWindowProgressValue func(window Window, value float32) bool

/**
* Get the value of the progress bar for the given window’s taskbar icon.
*
* \param window the window to get the current progress value from.
* \returns the progress value in the range of [0.0f - 1.0f], or -1.0f on
*          failure; call SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.4.0.
 */
//go:sdl3extern
var GetWindowProgressValue func(window Window) float32

/**
* Destroy a window.
*
* Any child windows owned by the window will be recursively destroyed as
* well.
*
* Note that on some platforms, the visible window may not actually be removed
* from the screen until the SDL event loop is pumped again, even though the
* SDL_Window is no longer valid after this call.
*
* \param window the window to destroy.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_CreatePopupWindow
* \sa SDL_CreateWindow
* \sa SDL_CreateWindowWithProperties
 */
//go:sdl3extern
var DestroyWindow func(window Window)

/**
* Check whether the screensaver is currently enabled.
*
* The screensaver is disabled by default.
*
* The default can also be changed using `SDL_HINT_VIDEO_ALLOW_SCREENSAVER`.
*
* \returns true if the screensaver is enabled, false if it is disabled.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_DisableScreenSaver
* \sa SDL_EnableScreenSaver
 */
//go:sdl3extern
var ScreenSaverEnabled func() bool

/**
* Allow the screen to be blanked by a screen saver.
*
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_DisableScreenSaver
* \sa SDL_ScreenSaverEnabled
 */
//go:sdl3extern
var EnableScreenSaver func() bool

/**
* Prevent the screen from being blanked by a screen saver.
*
* If you disable the screensaver, it is automatically re-enabled when SDL
* quits.
*
* The screensaver is disabled by default, but this may by changed by
* SDL_HINT_VIDEO_ALLOW_SCREENSAVER.
*
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_EnableScreenSaver
* \sa SDL_ScreenSaverEnabled
 */
//go:sdl3extern
var DisableScreenSaver func() bool

/**
*  \name OpenGL support functions
 */
/* @{ */

/**
* Dynamically load an OpenGL library.
*
* This should be done after initializing the video driver, but before
* creating any OpenGL windows. If no OpenGL library is loaded, the default
* library will be loaded upon creation of the first OpenGL window.
*
* If you do this, you need to retrieve all of the GL functions used in your
* program from the dynamic library using SDL_GL_GetProcAddress().
*
* \param path the platform dependent OpenGL library name, or NULL to open the
*             default OpenGL library.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_GetProcAddress
* \sa SDL_GL_UnloadLibrary
 */
//go:sdl3extern
var GL_LoadLibrary func(path *byte) bool

/**
* Get an OpenGL function by name.
*
* If the GL library is loaded at runtime with SDL_GL_LoadLibrary(), then all
* GL functions must be retrieved this way. Usually this is used to retrieve
* function pointers to OpenGL extensions.
*
* There are some quirks to looking up OpenGL functions that require some
* extra care from the application. If you code carefully, you can handle
* these quirks without any platform-specific code, though:
*
* - On Windows, function pointers are specific to the current GL context;
*   this means you need to have created a GL context and made it current
*   before calling SDL_GL_GetProcAddress(). If you recreate your context or
*   create a second context, you should assume that any existing function
*   pointers aren't valid to use with it. This is (currently) a
*   Windows-specific limitation, and in practice lots of drivers don't suffer
*   this limitation, but it is still the way the wgl API is documented to
*   work and you should expect crashes if you don't respect it. Store a copy
*   of the function pointers that comes and goes with context lifespan.
* - On X11, function pointers returned by this function are valid for any
*   context, and can even be looked up before a context is created at all.
*   This means that, for at least some common OpenGL implementations, if you
*   look up a function that doesn't exist, you'll get a non-NULL result that
*   is _NOT_ safe to call. You must always make sure the function is actually
*   available for a given GL context before calling it, by checking for the
*   existence of the appropriate extension with SDL_GL_ExtensionSupported(),
*   or verifying that the version of OpenGL you're using offers the function
*   as core functionality.
* - Some OpenGL drivers, on all platforms, *will* return NULL if a function
*   isn't supported, but you can't count on this behavior. Check for
*   extensions you use, and if you get a NULL anyway, act as if that
*   extension wasn't available. This is probably a bug in the driver, but you
*   can code defensively for this scenario anyhow.
* - Just because you're on Linux/Unix, don't assume you'll be using X11.
*   Next-gen display servers are waiting to replace it, and may or may not
*   make the same promises about function pointers.
* - OpenGL function pointers must be declared `APIENTRY` as in the example
*   code. This will ensure the proper calling convention is followed on
*   platforms where this matters (Win32) thereby avoiding stack corruption.
*
* \param proc the name of an OpenGL function.
* \returns a pointer to the named OpenGL function. The returned pointer
*          should be cast to the appropriate function signature.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_ExtensionSupported
* \sa SDL_GL_LoadLibrary
* \sa SDL_GL_UnloadLibrary
 */
//go:sdl3extern
var GL_GetProcAddress func(proc *byte) uintptr

/**
* Get an EGL library function by name.
*
* If an EGL library is loaded, this function allows applications to get entry
* points for EGL functions. This is useful to provide to an EGL API and
* extension loader.
*
* \param proc the name of an EGL function.
* \returns a pointer to the named EGL function. The returned pointer should
*          be cast to the appropriate function signature.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_EGL_GetCurrentDisplay
 */
//go:sdl3extern
var EGL_GetProcAddress func(proc *byte) uintptr

/**
* Unload the OpenGL library previously loaded by SDL_GL_LoadLibrary().
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_LoadLibrary
 */
//go:sdl3extern
var GL_UnloadLibrary func()

/**
* Check if an OpenGL extension is supported for the current context.
*
* This function operates on the current GL context; you must have created a
* context and it must be current before calling this function. Do not assume
* that all contexts you create will have the same set of extensions
* available, or that recreating an existing context will offer the same
* extensions again.
*
* While it's probably not a massive overhead, this function is not an O(1)
* operation. Check the extensions you care about after creating the GL
* context and save that information somewhere instead of calling the function
* every time you need to know.
*
* \param extension the name of the extension to check.
* \returns true if the extension is supported, false otherwise.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GL_ExtensionSupported func(extension *byte) bool

/**
* Reset all previously set OpenGL context attributes to their default values.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_GetAttribute
* \sa SDL_GL_SetAttribute
 */
//go:sdl3extern
var GL_ResetAttributes func()

/**
* Set an OpenGL window attribute before window creation.
*
* This function sets the OpenGL attribute `attr` to `value`. The requested
* attributes should be set before creating an OpenGL window. You should use
* SDL_GL_GetAttribute() to check the values after creating the OpenGL
* context, since the values obtained can differ from the requested ones.
*
* \param attr an SDL_GLAttr enum value specifying the OpenGL attribute to
*             set.
* \param value the desired value for the attribute.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_GetAttribute
* \sa SDL_GL_ResetAttributes
 */
//go:sdl3extern
var GL_SetAttribute func(attr GLAttr, value int32) bool

/**
* Get the actual value for an attribute from the current context.
*
* \param attr an SDL_GLAttr enum value specifying the OpenGL attribute to
*             get.
* \param value a pointer filled in with the current value of `attr`.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_ResetAttributes
* \sa SDL_GL_SetAttribute
 */
//go:sdl3extern
var GL_GetAttribute func(attr GLAttr, value *int32) bool

/**
* Create an OpenGL context for an OpenGL window, and make it current.
*
* Windows users new to OpenGL should note that, for historical reasons, GL
* functions added after OpenGL version 1.1 are not available by default.
* Those functions must be loaded at run-time, either with an OpenGL
* extension-handling library or with SDL_GL_GetProcAddress() and its related
* functions.
*
* SDL_GLContext is opaque to the application.
*
* \param window the window to associate with the context.
* \returns the OpenGL context associated with `window` or NULL on failure;
*          call SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_DestroyContext
* \sa SDL_GL_MakeCurrent
 */
//go:sdl3extern
var GL_CreateContext func(window Window) GLContext

/**
* Set up an OpenGL context for rendering into an OpenGL window.
*
* The context must have been created with a compatible window.
*
* \param window the window to associate with the context.
* \param context the OpenGL context to associate with the window.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_CreateContext
 */
//go:sdl3extern
var GL_MakeCurrent func(window Window, context GLContext) bool

/**
* Get the currently active OpenGL window.
*
* \returns the currently active OpenGL window on success or NULL on failure;
*          call SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GL_GetCurrentWindow func() Window

/**
* Get the currently active OpenGL context.
*
* \returns the currently active OpenGL context or NULL on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_MakeCurrent
 */
//go:sdl3extern
var GL_GetCurrentContext func() GLContext

/**
* Get the currently active EGL display.
*
* \returns the currently active EGL display or NULL on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var EGL_GetCurrentDisplay func() EGLDisplay

/**
* Get the currently active EGL config.
*
* \returns the currently active EGL config or NULL on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var EGL_GetCurrentConfig func() EGLConfig

/**
* Get the EGL surface associated with the window.
*
* \param window the window to query.
* \returns the EGLSurface pointer associated with the window, or NULL on
*          failure.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var EGL_GetWindowSurface func(window Window) EGLSurface

/**
* Sets the callbacks for defining custom EGLAttrib arrays for EGL
* initialization.
*
* Callbacks that aren't needed can be set to NULL.
*
* NOTE: These callback pointers will be reset after SDL_GL_ResetAttributes.
*
* \param platformAttribCallback callback for attributes to pass to
*                               eglGetPlatformDisplay. May be NULL.
* \param surfaceAttribCallback callback for attributes to pass to
*                              eglCreateSurface. May be NULL.
* \param contextAttribCallback callback for attributes to pass to
*                              eglCreateContext. May be NULL.
* \param userdata a pointer that is passed to the callbacks.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var EGL_SetAttributeCallbacks func(
	platformAttribCallback EGLAttribArrayCallback,
	surfaceAttribCallback EGLIntArrayCallback,
	contextAttribCallback EGLIntArrayCallback, userdata uintptr,
)

/**
* Set the swap interval for the current OpenGL context.
*
* Some systems allow specifying -1 for the interval, to enable adaptive
* vsync. Adaptive vsync works the same as vsync, but if you've already missed
* the vertical retrace for a given frame, it swaps buffers immediately, which
* might be less jarring for the user during occasional framerate drops. If an
* application requests adaptive vsync and the system does not support it,
* this function will fail and return false. In such a case, you should
* probably retry the call with 1 for the interval.
*
* Adaptive vsync is implemented for some glX drivers with
* GLX_EXT_swap_control_tear, and for some Windows drivers with
* WGL_EXT_swap_control_tear.
*
* Read more on the Khronos wiki:
* https://www.khronos.org/opengl/wiki/Swap_Interval#Adaptive_Vsync
*
* \param interval 0 for immediate updates, 1 for updates synchronized with
*                 the vertical retrace, -1 for adaptive vsync.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_GetSwapInterval
 */
//go:sdl3extern
var GL_SetSwapInterval func(interval int32) bool

/**
* Get the swap interval for the current OpenGL context.
*
* If the system can't determine the swap interval, or there isn't a valid
* current context, this function will set *interval to 0 as a safe default.
*
* \param interval output interval value. 0 if there is no vertical retrace
*                 synchronization, 1 if the buffer swap is synchronized with
*                 the vertical retrace, and -1 if late swaps happen
*                 immediately instead of waiting for the next retrace.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_SetSwapInterval
 */
//go:sdl3extern
var GL_GetSwapInterval func(interval *int32) bool

/**
* Update a window with OpenGL rendering.
*
* This is used with double-buffered OpenGL contexts, which are the default.
*
* On macOS, make sure you bind 0 to the draw framebuffer before swapping the
* window, otherwise nothing will happen. If you aren't using
* glBindFramebuffer(), this is the default and you won't have to do anything
* extra.
*
* \param window the window to change.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GL_SwapWindow func(window Window) bool

/**
* Delete an OpenGL context.
*
* \param context the OpenGL context to be deleted.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety This function should only be called on the main thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GL_CreateContext
 */
//go:sdl3extern
var GL_DestroyContext func(context GLContext) bool
