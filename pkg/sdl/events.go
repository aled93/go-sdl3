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
 * # CategoryEvents
 *
 * Event queue management.
 *
 * It's extremely common--often required--that an app deal with SDL's event
 * queue. Almost all useful information about interactions with the real world
 * flow through here: the user interacting with the computer and app, hardware
 * coming and going, the system changing in some way, etc.
 *
 * An app generally takes a moment, perhaps at the start of a new frame, to
 * examine any events that have occured since the last time and process or
 * ignore them. This is generally done by calling SDL_PollEvent() in a loop
 * until it returns false (or, if using the main callbacks, events are
 * provided one at a time in calls to SDL_AppEvent() before the next call to
 * SDL_AppIterate(); in this scenario, the app does not call SDL_PollEvent()
 * at all).
 *
 * There is other forms of control, too: SDL_PeepEvents() has more
 * functionality at the cost of more complexity, and SDL_WaitEvent() can block
 * the process until something interesting happens, which might be beneficial
 * for certain types of programs on low-power hardware. One may also call
 * SDL_AddEventWatch() to set a callback when new events arrive.
 *
 * The app is free to generate their own events, too: SDL_PushEvent allows the
 * app to put events onto the queue for later retrieval; SDL_RegisterEvents
 * can guarantee that these events have a type that isn't in use by other
 * parts of the system.
 */
package sdl

import "unsafe"

/* General keyboard/mouse/pen state definitions */

/**
 * The types of events that can be delivered.
 *
 * \since This enum is available since SDL 3.2.0.
 */
type EventType int

const (
	ET_First EventType = 0 /**< Unused (do not remove) */

	/* Application events */
	ET_Quit EventType = 0x100 + iota /**< User-requested quit */

	/* These application events have special meaning on iOS and Android, see README-ios.md and README-android.md for details */
	ET_Terminating /**< The application is being terminated by the OS. This event must be handled in a callback set with SDL_AddEventWatch().
	  Called on iOS in applicationWillTerminate()
	  Called on Android in onDestroy()
	*/
	ET_LowMemory /**< The application is low on memory, free memory if possible. This event must be handled in a callback set with SDL_AddEventWatch().
	  Called on iOS in applicationDidReceiveMemoryWarning()
	  Called on Android in onTrimMemory()
	*/
	ET_WillEnterBackground /**< The application is about to enter the background. This event must be handled in a callback set with SDL_AddEventWatch().
	  Called on iOS in applicationWillResignActive()
	  Called on Android in onPause()
	*/
	ET_DidEnterBackground /**< The application did enter the background and may not get CPU for some time. This event must be handled in a callback set with SDL_AddEventWatch().
	  Called on iOS in applicationDidEnterBackground()
	  Called on Android in onPause()
	*/
	ET_WillEnterForeground /**< The application is about to enter the foreground. This event must be handled in a callback set with SDL_AddEventWatch().
	  Called on iOS in applicationWillEnterForeground()
	  Called on Android in onResume()
	*/
	ET_DidEnterForeground /**< The application is now interactive. This event must be handled in a callback set with SDL_AddEventWatch().
	  Called on iOS in applicationDidBecomeActive()
	  Called on Android in onResume()
	*/

	ET_LocaleChanged /**< The user's locale preferences have changed. */

	ET_SystemThemeChanged /**< The system theme changed */

	/* Display events */
	/* 0x150 was SDL_DISPLAYEVENT, reserve the number for sdl2-compat */
	ET_DisplayOrientation         EventType = 0x151 + iota /**< Display orientation has changed to data1 */
	ET_DisplayAdded                                        /**< Display has been added to the system */
	ET_DisplayRemoved                                      /**< Display has been removed from the system */
	ET_DisplayMoved                                        /**< Display has changed position */
	ET_DisplayDesktopModeChanged                           /**< Display has changed desktop mode */
	ET_DisplayCurrentModeChanged                           /**< Display has changed current mode */
	ET_DisplayContentScaleChanged                          /**< Display has changed content scale */
	ET_DisplayFirst                         = ET_DisplayOrientation
	ET_DisplayLast                          = ET_DisplayContentScaleChanged

	/* Window events */
	/* 0x200 was SDL_WINDOWEVENT, reserve the number for sdl2-compat */
	/* 0x201 was SDL_SYSWMEVENT, reserve the number for sdl2-compat */
	ET_WindowShown               EventType = 0x202 + iota /**< Window has been shown */
	ET_WindowHidden                                       /**< Window has been hidden */
	ET_WindowExposed                                      /**< Window has been exposed and should be redrawn, and can be redrawn directly from event watchers for this event */
	ET_WindowMoved                                        /**< Window has been moved to data1, data2 */
	ET_WindowResized                                      /**< Window has been resized to data1xdata2 */
	ET_WindowPixelSizeChanged                             /**< The pixel size of the window has changed to data1xdata2 */
	ET_WindowMetalViewResized                             /**< The pixel size of a Metal view associated with the window has changed */
	ET_WindowMinimized                                    /**< Window has been minimized */
	ET_WindowMaximized                                    /**< Window has been maximized */
	ET_WindowRestored                                     /**< Window has been restored to normal size and position */
	ET_WindowMouseEnter                                   /**< Window has gained mouse focus */
	ET_WindowMouseLeave                                   /**< Window has lost mouse focus */
	ET_WindowFocusGained                                  /**< Window has gained keyboard focus */
	ET_WindowFocusLost                                    /**< Window has lost keyboard focus */
	ET_WindowCloseRequested                               /**< The window manager requests that the window be closed */
	ET_WindowHitRest                                      /**< Window had a hit test that wasn't SDL_HITTEST_NORMAL */
	ET_WindowIccprofChanged                               /**< The ICC profile of the window's display has changed */
	ET_WindowDisplayChanged                               /**< Window has been moved to display data1 */
	ET_WindowDisplayScaleChanged                          /**< Window display scale has been changed */
	ET_WindowSafeAreaChanged                              /**< The window safe area has been changed */
	ET_WindowOccluded                                     /**< The window has been occluded */
	ET_WindowEnterFullscreen                              /**< The window has entered fullscreen mode */
	ET_WindowLeaveFullscreen                              /**< The window has left fullscreen mode */
	ET_WindowDestroyed                                    /**< The window with the associated ID is being or has been destroyed. If this message is being handled
	  in an event watcher, the window handle is still valid and can still be used to retrieve any properties
	  associated with the window. Otherwise, the handle has already been destroyed and all resources
	  associated with it are invalid */
	ET_WindowHdrStateChanged /**< Window HDR properties have changed */
	ET_WindowFirst           = ET_WindowShown
	ET_WindowLast            = ET_WindowHdrStateChanged

	/* Keyboard events */
	ET_KeyDown       = 0x300 + iota /**< Key pressed */
	ET_KeyUp                        /**< Key released */
	ET_TextEditing                  /**< Keyboard text editing (composition) */
	ET_TextInput                    /**< Keyboard text input */
	ET_KeymapChanged                /**< Keymap changed due to a system event such as an
	  input language or keyboard layout change. */
	ET_KeyboardAdded         /**< A new keyboard has been inserted into the system */
	ET_KeyboardRemoved       /**< A keyboard has been removed */
	ET_TextEditingCandidates /**< Keyboard text editing candidates */

	/* Mouse events */
	ET_MouseMotion     EventType = 0x400 + iota /**< Mouse moved */
	ET_MouseButtonDown                          /**< Mouse button pressed */
	ET_MouseButtonUp                            /**< Mouse button released */
	ET_MouseWheel                               /**< Mouse wheel motion */
	ET_MouseAdded                               /**< A new mouse has been inserted into the system */
	ET_MouseRemoved                             /**< A mouse has been removed */

	/* Joystick events */
	ET_JoystickAxisMotion     EventType = 0x600 + iota /**< Joystick axis motion */
	ET_JoystickBallMotion                              /**< Joystick trackball motion */
	ET_JoystickHatMotion                               /**< Joystick hat position change */
	ET_JoystickButtonDown                              /**< Joystick button pressed */
	ET_JoystickButtonUp                                /**< Joystick button released */
	ET_JoystickAdded                                   /**< A new joystick has been inserted into the system */
	ET_JoystickRemoved                                 /**< An opened joystick has been removed */
	ET_JoystickBatteryUpdated                          /**< Joystick battery level change */
	ET_JoystickUpdateComplete                          /**< Joystick update is complete */

	/* Gamepad events */
	ET_GamepadAxisMotion         EventType = 0x650 + iota /**< Gamepad axis motion */
	ET_GamepadButtonDown                                  /**< Gamepad button pressed */
	ET_GamepadButtonUp                                    /**< Gamepad button released */
	ET_GamepadAdded                                       /**< A new gamepad has been inserted into the system */
	ET_GamepadRemoved                                     /**< A gamepad has been removed */
	ET_GamepadRemapped                                    /**< The gamepad mapping was updated */
	ET_GamepadTouchpadDown                                /**< Gamepad touchpad was touched */
	ET_GamepadTouchpadMotion                              /**< Gamepad touchpad finger was moved */
	ET_GamepadTouchpadUp                                  /**< Gamepad touchpad finger was lifted */
	ET_GamepadSensorUpdate                                /**< Gamepad sensor was updated */
	ET_GamepadUpdateComplete                              /**< Gamepad update is complete */
	ET_GamepadSteamHandleUpdated                          /**< Gamepad Steam handle has changed */

	/* Touch events */
	ET_FingerDown EventType = 0x700
	ET_FingerUp
	ET_FingerMotion
	ET_FingerCanceled

	/* 0x800, 0x801, and 0x802 were the Gesture events from SDL2. Do not reuse these values! sdl2-compat needs them! */

	/* Clipboard events */
	ET_Clipboard_update EventType = 0x900 /**< The clipboard or primary selection changed */

	/* Drag and drop events */
	ET_DropFile     EventType = 0x1000 + iota /**< The system requests a file open */
	ET_DropText                               /**< text/plain drag-and-drop event */
	ET_DropBegin                              /**< A new set of drops is beginning (NULL filename) */
	ET_DropComplete                           /**< Current set of drops is now complete (NULL filename) */
	ET_DropPosition                           /**< Position while moving over the window */

	/* Audio hotplug events */
	ET_AudioDeviceAdded         EventType = 0x1100 + iota /**< A new audio device is available */
	ET_AudioDeviceRemoved                                 /**< An audio device has been removed. */
	ET_AudioDeviceFormatChanged                           /**< An audio device's format has been changed by the system. */

	/* Sensor events */
	ET_SensorUpdate EventType = 0x1200 /**< A sensor was updated */

	/* Pressure-sensitive pen events */
	ET_PenProximityIn  EventType = 0x1300 + iota /**< Pressure-sensitive pen has become available */
	ET_PenProximityOut                           /**< Pressure-sensitive pen has become unavailable */
	ET_PenDown                                   /**< Pressure-sensitive pen touched drawing surface */
	ET_PenUp                                     /**< Pressure-sensitive pen stopped touching drawing surface */
	ET_PenButtonDown                             /**< Pressure-sensitive pen button pressed */
	ET_PenButtonUp                               /**< Pressure-sensitive pen button released */
	ET_PenMotion                                 /**< Pressure-sensitive pen is moving on the tablet */
	ET_PenAxis                                   /**< Pressure-sensitive pen angle/pressure/etc changed */

	/* Camera hotplug events */
	ET_CameraDeviceAdded    = 0x1400 + iota /**< A new camera device is available */
	ET_CameraDeviceRemoved                  /**< A camera device has been removed. */
	ET_CameraDeviceApproved                 /**< A camera device has been approved for use by the user. */
	ET_CameraDeviceDenied                   /**< A camera device has been denied for use by the user. */

	/* Render events */
	ET_RenderTargetsReset = 0x2000 + iota /**< The render targets have been reset and their contents need to be updated */
	ET_RenderDeviceReset                  /**< The device has been reset and all textures need to be recreated */
	ET_RenderDeviceLost                   /**< The device has been lost and can't be recovered. */

	/* Reserved events for private platforms */
	ET_Private0 EventType = 0x4000 + iota
	ET_Private1
	ET_Private2
	ET_Private3

	/* Internal events */
	ET_PollSentinel EventType = 0x7F00 /**< Signals the end of an event poll cycle */

	/** Events ET_User through ET_Last are for your use,
	 *  and should be allocated with SDL_RegisterEvents()
	 */
	ET_User EventType = 0x8000

	/**
	 *  This last event is only for bounding internal arrays
	 */
	ET_Last EventType = 0xFFFF

	/* This just makes sure the enum is the size of Uint32 */
	ET_EnumPadding EventType = 0x7FFFFFFF
)

/**
 * Fields shared by every event
 *
 * \since This struct is available since SDL 3.2.0.
 */
type CommonEvent struct {
	Type      uint32 /**< Event type, shared with all events, Uint32 to cover user events which are not in the SDL_EventType enumeration */
	Reserved  uint32
	Timestamp uint64 /**< In nanoseconds, populated using SDL_GetTicksNS() */
}

/**
 * Display state change event data (event.display.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type DisplayEvent struct {
	Type      EventType /**< SDL_DISPLAYEVENT_* */
	Reserved  uint32
	Timestamp uint64    /**< In nanoseconds, populated using SDL_GetTicksNS() */
	DisplayID DisplayID /**< The associated display */
	Data1     int32     /**< event dependent data */
	Data2     int32     /**< event dependent data */
}

/**
 * Window state change event data (event.window.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type WindowEvent struct {
	Type      EventType /**< SDL_EVENT_WINDOW_* */
	Reserved  uint32
	Timestamp uint64   /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID /**< The associated window */
	Data1     int32    /**< event dependent data */
	Data2     int32    /**< event dependent data */
}

/**
 * Keyboard device event structure (event.kdevice.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type KeyboardDeviceEvent struct {
	Type      EventType /**< SDL_EVENT_KEYBOARD_ADDED or SDL_EVENT_KEYBOARD_REMOVED */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     KeyboardID /**< The keyboard instance id */
}

/**
 * Keyboard button event structure (event.key.*)
 *
 * The `key` is the base SDL_Keycode generated by pressing the `scancode`
 * using the current keyboard layout, applying any options specified in
 * SDL_HINT_KEYCODE_OPTIONS. You can get the SDL_Keycode corresponding to the
 * event scancode and modifiers directly from the keyboard layout, bypassing
 * SDL_HINT_KEYCODE_OPTIONS, by calling SDL_GetKeyFromScancode().
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GetKeyFromScancode
 * \sa SDL_HINT_KEYCODE_OPTIONS
 */
type KeyboardEvent struct {
	Type      EventType /**< SDL_EVENT_KEY_DOWN or SDL_EVENT_KEY_UP */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID   /**< The window with keyboard focus, if any */
	Which     KeyboardID /**< The keyboard instance id, or 0 if unknown or virtual */
	Scancode  Scancode   /**< SDL physical key code */
	Key       Keycode    /**< SDL virtual key code */
	Mod       Keymod     /**< current key modifiers */
	Raw       uint16     /**< The platform dependent scancode for this event */
	Down      bool       /**< true if the key is pressed */
	Repeat    bool       /**< true if this is a key repeat */
}

/**
 * Keyboard text editing event structure (event.edit.*)
 *
 * The start cursor is the position, in UTF-8 characters, where new typing
 * will be inserted into the editing text. The length is the number of UTF-8
 * characters that will be replaced by new typing.
 *
 * \since This struct is available since SDL 3.2.0.
 */
type TextEditingEvent struct {
	Type      EventType /**< SDL_EVENT_TEXT_EDITING */
	Reserved  uint32
	Timestamp uint64   /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID /**< The window with keyboard focus, if any */
	Text      string   /**< The editing text */
	Start     int32    /**< The start cursor of selected editing text, or -1 if not set */
	Length    int32    /**< The length of selected editing text, or -1 if not set */
}

/**
 * Keyboard IME candidates event structure (event.edit_candidates.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type TextEditingCandidatesEvent struct {
	Type              EventType /**< SDL_EVENT_TEXT_EDITING_CANDIDATES */
	Reserved          uint32
	Timestamp         uint64   /**< In nanoseconds, populated using SDL_GetTicksNS() */
	windowID          WindowID /**< The window with keyboard focus, if any */
	Candidates        []string /**< The list of candidates, or NULL if there are no candidates available */
	NumCandidates     int32    /**< The number of strings in `candidates` */
	SelectedCandidate int32    /**< The index of the selected candidate, or -1 if no candidate is selected */
	Horizontal        bool     /**< true if the list is horizontal, false if it's vertical */
	Padding1          uint8
	Padding2          uint8
	Padding3          uint8
}

/**
 * Keyboard text input event structure (event.text.*)
 *
 * This event will never be delivered unless text input is enabled by calling
 * SDL_StartTextInput(). Text input is disabled by default!
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_StartTextInput
 * \sa SDL_StopTextInput
 */
type TextInputEvent struct {
	Type      EventType /**< SDL_EVENT_TEXT_INPUT */
	Reserved  uint32
	Timestamp uint64   /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID /**< The window with keyboard focus, if any */
	Text      string   /**< The input text, UTF-8 encoded */
}

/**
 * Mouse device event structure (event.mdevice.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type MouseDeviceEvent struct {
	Type      EventType /**< SDL_EVENT_MOUSE_ADDED or SDL_EVENT_MOUSE_REMOVED */
	Reserved  uint32
	Timestamp uint64  /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     MouseID /**< The mouse instance id */
}

/**
 * Mouse motion event structure (event.motion.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type MouseMotionEvent struct {
	Type      EventType /**< SDL_EVENT_MOUSE_MOTION */
	Reserved  uint32
	Timestamp uint64           /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID         /**< The window with mouse focus, if any */
	Which     MouseID          /**< The mouse instance id in relative mode, SDL_TOUCH_MOUSEID for touch events, or 0 */
	State     MouseButtonFlags /**< The current button state */
	X         float32          /**< X coordinate, relative to window */
	Y         float32          /**< Y coordinate, relative to window */
	XRel      float32          /**< The relative motion in the X direction */
	YRel      float32          /**< The relative motion in the Y direction */
}

/**
 * Mouse button event structure (event.button.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type MouseButtonEvent struct {
	Type      EventType /**< SDL_EVENT_MOUSE_BUTTON_DOWN or SDL_EVENT_MOUSE_BUTTON_UP */
	Reserved  uint32
	Timestamp uint64   /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID /**< The window with mouse focus, if any */
	Which     MouseID  /**< The mouse instance id in relative mode, SDL_TOUCH_MOUSEID for touch events, or 0 */
	Button    uint8    /**< The mouse button index */
	Down      bool     /**< true if the button is pressed */
	Clicks    uint8    /**< 1 for single-click, 2 for double-click, etc. */
	Padding   uint8
	X         float32 /**< X coordinate, relative to window */
	Y         float32 /**< Y coordinate, relative to window */
}

/**
 * Mouse wheel event structure (event.wheel.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type MouseWheelEvent struct {
	Type      EventType /**< SDL_EVENT_MOUSE_WHEEL */
	Reserved  uint32
	Timestamp uint64              /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID            /**< The window with mouse focus, if any */
	Which     MouseID             /**< The mouse instance id in relative mode or 0 */
	X         float32             /**< The amount scrolled horizontally, positive to the right and negative to the left */
	Y         float32             /**< The amount scrolled vertically, positive away from the user and negative toward the user */
	Direction MouseWheelDirection /**< Set to one of the SDL_MOUSEWHEEL_* defines. When FLIPPED the values in X and Y will be opposite. Multiply by -1 to change them back */
	MouseX    float32             /**< X coordinate, relative to window */
	MouseY    float32             /**< Y coordinate, relative to window */
	IntegerX  int32               /**< The amount scrolled horizontally, accumulated to whole scroll "ticks" (added in 3.2.12) */
	IntegerY  int32               /**< The amount scrolled vertically, accumulated to whole scroll "ticks" (added in 3.2.12) */
}

/**
 * Joystick axis motion event structure (event.jaxis.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type JoyAxisEvent struct {
	Type      EventType /**< SDL_EVENT_JOYSTICK_AXIS_MOTION */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     JoystickID /**< The joystick instance id */
	Axis      uint8      /**< The joystick axis index */
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Value     int16 /**< The axis value (range: -32768 to 32767) */
	Padding4  uint16
}

/**
 * Joystick trackball motion event structure (event.jball.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type JoyBallEvent struct {
	Type      EventType /**< SDL_EVENT_JOYSTICK_BALL_MOTION */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     JoystickID /**< The joystick instance id */
	Ball      uint8      /**< The joystick trackball index */
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	XRel      int16 /**< The relative motion in the X direction */
	YRel      int16 /**< The relative motion in the Y direction */
}

/**
 * Joystick hat position change event structure (event.jhat.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type JoyHatEvent struct {
	Type      EventType /**< SDL_EVENT_JOYSTICK_HAT_MOTION */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     JoystickID /**< The joystick instance id */
	Hat       uint8      /**< The joystick hat index */
	Value     uint8      /**< The hat position value.
	 *   \sa SDL_HAT_LEFTUP SDL_HAT_UP SDL_HAT_RIGHTUP
	 *   \sa SDL_HAT_LEFT SDL_HAT_CENTERED SDL_HAT_RIGHT
	 *   \sa SDL_HAT_LEFTDOWN SDL_HAT_DOWN SDL_HAT_RIGHTDOWN
	 *
	 *   Note that zero means the POV is centered.
	 */
	Padding1 uint8
	Padding2 uint8
}

/**
 * Joystick button event structure (event.jbutton.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type JoyButtonEvent struct {
	Type      EventType /**< SDL_EVENT_JOYSTICK_BUTTON_DOWN or SDL_EVENT_JOYSTICK_BUTTON_UP */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     JoystickID /**< The joystick instance id */
	Button    uint8      /**< The joystick button index */
	Down      bool       /**< true if the button is pressed */
	Padding1  uint8
	Padding2  uint8
}

/**
 * Joystick device event structure (event.jdevice.*)
 *
 * SDL will send JOYSTICK_ADDED events for devices that are already plugged in
 * during SDL_Init.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GamepadDeviceEvent
 */
type JoyDeviceEvent struct {
	Type      EventType /**< SDL_EVENT_JOYSTICK_ADDED or SDL_EVENT_JOYSTICK_REMOVED or SDL_EVENT_JOYSTICK_UPDATE_COMPLETE */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     JoystickID /**< The joystick instance id */
}

/**
 * Joystick battery level change event structure (event.jbattery.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type JoyBatteryEvent struct {
	Type      EventType /**< SDL_EVENT_JOYSTICK_BATTERY_UPDATED */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     JoystickID /**< The joystick instance id */
	State     PowerState /**< The joystick battery state */
	Percent   int        /**< The joystick battery percent charge remaining */
}

/**
 * Gamepad axis motion event structure (event.gaxis.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type GamepadAxisEvent struct {
	Type      EventType /**< SDL_EVENT_GAMEPAD_AXIS_MOTION */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     JoystickID /**< The joystick instance id */
	Axis      uint8      /**< The gamepad axis (SDL_GamepadAxis) */
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Value     int16 /**< The axis value (range: -32768 to 32767) */
	Padding4  uint16
}

/**
 * Gamepad button event structure (event.gbutton.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type GamepadButtonEvent struct {
	Type      EventType /**< SDL_EVENT_GAMEPAD_BUTTON_DOWN or SDL_EVENT_GAMEPAD_BUTTON_UP */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     JoystickID /**< The joystick instance id */
	Button    uint8      /**< The gamepad button (SDL_GamepadButton) */
	Down      bool       /**< true if the button is pressed */
	Padding1  uint8
	Padding2  uint8
}

/**
 * Gamepad device event structure (event.gdevice.*)
 *
 * Joysticks that are supported gamepads receive both an SDL_JoyDeviceEvent
 * and an SDL_GamepadDeviceEvent.
 *
 * SDL will send GAMEPAD_ADDED events for joysticks that are already plugged
 * in during SDL_Init() and are recognized as gamepads. It will also send
 * events for joysticks that get gamepad mappings at runtime.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_JoyDeviceEvent
 */
type GamepadDeviceEvent struct {
	Type      EventType /**< SDL_EVENT_GAMEPAD_ADDED, SDL_EVENT_GAMEPAD_REMOVED, or SDL_EVENT_GAMEPAD_REMAPPED, SDL_EVENT_GAMEPAD_UPDATE_COMPLETE or SDL_EVENT_GAMEPAD_STEAM_HANDLE_UPDATED */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     JoystickID /**< The joystick instance id */
}

/**
 * Gamepad touchpad event structure (event.gtouchpad.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type GamepadTouchpadEvent struct {
	Type      EventType /**< SDL_EVENT_GAMEPAD_TOUCHPAD_DOWN or SDL_EVENT_GAMEPAD_TOUCHPAD_MOTION or SDL_EVENT_GAMEPAD_TOUCHPAD_UP */
	Reserved  uint32
	Timestamp uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     JoystickID /**< The joystick instance id */
	Touchpad  int32      /**< The index of the touchpad */
	Finger    int32      /**< The index of the finger on the touchpad */
	X         float32    /**< Normalized in the range 0...1 with 0 being on the left */
	Y         float32    /**< Normalized in the range 0...1 with 0 being at the top */
	Pressure  float32    /**< Normalized in the range 0...1 */
}

/**
 * Gamepad sensor event structure (event.gsensor.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type GamepadSensorEvent struct {
	Type            EventType /**< SDL_EVENT_GAMEPAD_SENSOR_UPDATE */
	Reserved        uint32
	Timestamp       uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which           JoystickID /**< The joystick instance id */
	Sensor          int32      /**< The type of the sensor, one of the values of SDL_SensorType */
	Data            [3]float32 /**< Up to 3 values from the sensor, as defined in SDL_sensor.h */
	SensorTimestamp uint64     /**< The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock */
}

/**
 * Audio device event structure (event.adevice.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type AudioDeviceEvent struct {
	Type      EventType /**< SDL_EVENT_AUDIO_DEVICE_ADDED, or SDL_EVENT_AUDIO_DEVICE_REMOVED, or SDL_EVENT_AUDIO_DEVICE_FORMAT_CHANGED */
	Reserved  uint32
	Timestamp uint64        /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     AudioDeviceID /**< SDL_AudioDeviceID for the device being added or removed or changing */
	Recording bool          /**< false if a playback device, true if a recording device. */
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
}

/**
 * Camera device event structure (event.cdevice.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type CameraDeviceEvent struct {
	Type      EventType /**< SDL_EVENT_CAMERA_DEVICE_ADDED, SDL_EVENT_CAMERA_DEVICE_REMOVED, SDL_EVENT_CAMERA_DEVICE_APPROVED, SDL_EVENT_CAMERA_DEVICE_DENIED */
	Reserved  uint32
	Timestamp uint64   /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     CameraID /**< SDL_CameraID for the device being added or removed or changing */
}

/**
 * Renderer event structure (event.render.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type RenderEvent struct {
	Type      EventType /**< SDL_EVENT_RENDER_TARGETS_RESET, SDL_EVENT_RENDER_DEVICE_RESET, SDL_EVENT_RENDER_DEVICE_LOST */
	Reserved  uint32
	Timestamp uint64   /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID /**< The window containing the renderer in question. */
}

/**
 * Touch finger event structure (event.tfinger.*)
 *
 * Coordinates in this event are normalized. `x` and `y` are normalized to a
 * range between 0.0f and 1.0f, relative to the window, so (0,0) is the top
 * left and (1,1) is the bottom right. Delta coordinates `dx` and `dy` are
 * normalized in the ranges of -1.0f (traversed all the way from the bottom or
 * right to all the way up or left) to 1.0f (traversed all the way from the
 * top or left to all the way down or right).
 *
 * Note that while the coordinates are _normalized_, they are not _clamped_,
 * which means in some circumstances you can get a value outside of this
 * range. For example, a renderer using logical presentation might give a
 * negative value when the touch is in the letterboxing. Some platforms might
 * report a touch outside of the window, which will also be outside of the
 * range.
 *
 * \since This struct is available since SDL 3.2.0.
 */
type TouchFingerEvent struct {
	Type      EventType /**< SDL_EVENT_FINGER_DOWN, SDL_EVENT_FINGER_UP, SDL_EVENT_FINGER_MOTION, or SDL_EVENT_FINGER_CANCELED */
	Reserved  uint32
	Timestamp uint64  /**< In nanoseconds, populated using SDL_GetTicksNS() */
	TouchID   TouchID /**< The touch device id */
	FingerID  FingerID
	X         float32  /**< Normalized in the range 0...1 */
	Y         float32  /**< Normalized in the range 0...1 */
	Dx        float32  /**< Normalized in the range -1...1 */
	Dy        float32  /**< Normalized in the range -1...1 */
	Pressure  float32  /**< Normalized in the range 0...1 */
	WindowID  WindowID /**< The window underneath the finger, if any */
}

/**
 * Pressure-sensitive pen proximity event structure (event.pproximity.*)
 *
 * When a pen becomes visible to the system (it is close enough to a tablet,
 * etc), SDL will send an SDL_EVENT_PEN_PROXIMITY_IN event with the new pen's
 * ID. This ID is valid until the pen leaves proximity again (has been removed
 * from the tablet's area, the tablet has been unplugged, etc). If the same
 * pen reenters proximity again, it will be given a new ID.
 *
 * Note that "proximity" means "close enough for the tablet to know the tool
 * is there." The pen touching and lifting off from the tablet while not
 * leaving the area are handled by SDL_EVENT_PEN_DOWN and SDL_EVENT_PEN_UP.
 *
 * \since This struct is available since SDL 3.2.0.
 */
type PenProximityEvent struct {
	Type      EventType /**< SDL_EVENT_PEN_PROXIMITY_IN or SDL_EVENT_PEN_PROXIMITY_OUT */
	Reserved  uint32
	Timestamp uint64   /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID /**< The window with pen focus, if any */
	Which     PenID    /**< The pen instance id */
}

/**
 * Pressure-sensitive pen motion event structure (event.pmotion.*)
 *
 * Depending on the hardware, you may get motion events when the pen is not
 * touching a tablet, for tracking a pen even when it isn't drawing. You
 * should listen for SDL_EVENT_PEN_DOWN and SDL_EVENT_PEN_UP events, or check
 * `pen_state & SDL_PEN_INPUT_DOWN` to decide if a pen is "drawing" when
 * dealing with pen motion.
 *
 * \since This struct is available since SDL 3.2.0.
 */
type PenMotionEvent struct {
	Type      EventType /**< SDL_EVENT_PEN_MOTION */
	Reserved  uint32
	Timestamp uint64        /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID      /**< The window with pen focus, if any */
	Which     PenID         /**< The pen instance id */
	PenState  PenInputFlags /**< Complete pen input state at time of event */
	X         float32       /**< X coordinate, relative to window */
	Y         float32       /**< Y coordinate, relative to window */
}

/**
 * Pressure-sensitive pen touched event structure (event.ptouch.*)
 *
 * These events come when a pen touches a surface (a tablet, etc), or lifts
 * off from one.
 *
 * \since This struct is available since SDL 3.2.0.
 */
type PenTouchEvent struct {
	Type      EventType /**< SDL_EVENT_PEN_DOWN or SDL_EVENT_PEN_UP */
	Reserved  uint32
	Timestamp uint64        /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID      /**< The window with pen focus, if any */
	Which     PenID         /**< The pen instance id */
	PenState  PenInputFlags /**< Complete pen input state at time of event */
	X         float32       /**< X coordinate, relative to window */
	Y         float32       /**< Y coordinate, relative to window */
	Eraser    bool          /**< true if eraser end is used (not all pens support this). */
	Down      bool          /**< true if the pen is touching or false if the pen is lifted off */
}

/**
 * Pressure-sensitive pen button event structure (event.pbutton.*)
 *
 * This is for buttons on the pen itself that the user might click. The pen
 * itself pressing down to draw triggers a SDL_EVENT_PEN_DOWN event instead.
 *
 * \since This struct is available since SDL 3.2.0.
 */
type PenButtonEvent struct {
	Type      EventType /**< SDL_EVENT_PEN_BUTTON_DOWN or SDL_EVENT_PEN_BUTTON_UP */
	Reserved  uint32
	Timestamp uint64        /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID      /**< The window with mouse focus, if any */
	Which     PenID         /**< The pen instance id */
	PenState  PenInputFlags /**< Complete pen input state at time of event */
	X         float32       /**< X coordinate, relative to window */
	Y         float32       /**< Y coordinate, relative to window */
	Button    uint8         /**< The pen button index (first button is 1). */
	Down      bool          /**< true if the button is pressed */
}

/**
 * Pressure-sensitive pen pressure / angle event structure (event.paxis.*)
 *
 * You might get some of these events even if the pen isn't touching the
 * tablet.
 *
 * \since This struct is available since SDL 3.2.0.
 */
type PenAxisEvent struct {
	Type      EventType /**< SDL_EVENT_PEN_AXIS */
	Reserved  uint32
	Timestamp uint64        /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID      /**< The window with pen focus, if any */
	Which     PenID         /**< The pen instance id */
	PenState  PenInputFlags /**< Complete pen input state at time of event */
	X         float32       /**< X coordinate, relative to window */
	Y         float32       /**< Y coordinate, relative to window */
	Axis      PenAxis       /**< Axis that has changed */
	Value     float32       /**< New value of axis */
}

/**
 * An event used to drop text or request a file open by the system
 * (event.drop.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type DropEvent struct {
	Type      EventType /**< SDL_EVENT_DROP_BEGIN or SDL_EVENT_DROP_FILE or SDL_EVENT_DROP_TEXT or SDL_EVENT_DROP_COMPLETE or SDL_EVENT_DROP_POSITION */
	Reserved  uint32
	Timestamp uint64   /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID /**< The window that was dropped on, if any */
	X         float32  /**< X coordinate, relative to window (not on begin) */
	Y         float32  /**< Y coordinate, relative to window (not on begin) */
	Source    string   /**< The source app that sent this drop event, or NULL if that isn't available */
	Data      string   /**< The text for SDL_EVENT_DROP_TEXT and the file name for SDL_EVENT_DROP_FILE, NULL for other events */
}

/**
 * An event triggered when the clipboard contents have changed
 * (event.clipboard.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type ClipboardEvent struct {
	Type         EventType /**< SDL_EVENT_CLIPBOARD_UPDATE */
	Reserved     uint32
	Timestamp    uint64   /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Owner        bool     /**< are we owning the clipboard (internal update) */
	NumMimeTypes int32    /**< number of mime types */
	MimeTypes    []string /**< current mime types */
}

/**
 * Sensor event structure (event.sensor.*)
 *
 * \since This struct is available since SDL 3.2.0.
 */
type SensorEvent struct {
	Type            EventType /**< SDL_EVENT_SENSOR_UPDATE */
	Reserved        uint32
	Timestamp       uint64     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which           SensorID   /**< The instance ID of the sensor */
	Data            [6]float32 /**< Up to 6 values from the sensor - additional values can be queried using SDL_GetSensorData() */
	SensorTimestamp uint64     /**< The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock */
}

/**
 * The "quit requested" event
 *
 * \since This struct is available since SDL 3.2.0.
 */
type QuitEvent struct {
	Type      EventType /**< SDL_EVENT_QUIT */
	Reserved  uint32
	Timestamp uint64 /**< In nanoseconds, populated using SDL_GetTicksNS() */
}

/**
 * A user-defined event type (event.user.*)
 *
 * This event is unique; it is never created by SDL, but only by the
 * application. The event can be pushed onto the event queue using
 * SDL_PushEvent(). The contents of the structure members are completely up to
 * the programmer; the only requirement is that '''type''' is a value obtained
 * from SDL_RegisterEvents().
 *
 * \since This struct is available since SDL 3.2.0.
 */
type UserEvent struct {
	Type      uint32 /**< SDL_EVENT_USER through SDL_EVENT_LAST-1, Uint32 because these are not in the SDL_EventType enumeration */
	Reserved  uint32
	Timestamp uint64   /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  WindowID /**< The associated window if any */
	Code      int32    /**< User defined event code */
	Data1     uintptr  /**< User defined data pointer */
	Data2     uintptr  /**< User defined data pointer */
}

/**
 * The structure for all events in SDL.
 *
 * The SDL_Event structure is the core of all event handling in SDL. SDL_Event
 * is a union of all event structures used in SDL.
 *
 * \since This struct is available since SDL 3.2.0.
 */
type Event struct {
	Type EventType
	data [128]byte // must be >= largest event struct
}

func (e *Event) AsCommonEvent() CommonEvent {
	return *(*CommonEvent)(unsafe.Pointer(e))
}

func (e *Event) AsDisplayEvent() DisplayEvent {
	return *(*DisplayEvent)(unsafe.Pointer(e))
}

func (e *Event) AsWindowEvent() WindowEvent {
	return *(*WindowEvent)(unsafe.Pointer(e))
}

func (e *Event) AsKeyboardDeviceEvent() KeyboardDeviceEvent {
	return *(*KeyboardDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) AsKeyboardEvent() KeyboardEvent {
	return *(*KeyboardEvent)(unsafe.Pointer(e))
}

func (e *Event) AsTextEditingEvent() TextEditingEvent {
	return *(*TextEditingEvent)(unsafe.Pointer(e))
}

func (e *Event) AsTextEditingCandidatesEvent() TextEditingCandidatesEvent {
	return *(*TextEditingCandidatesEvent)(unsafe.Pointer(e))
}

func (e *Event) AsTextInputEvent() TextInputEvent {
	return *(*TextInputEvent)(unsafe.Pointer(e))
}

func (e *Event) AsMouseDeviceEvent() MouseDeviceEvent {
	return *(*MouseDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) AsMouseMotionEvent() MouseMotionEvent {
	return *(*MouseMotionEvent)(unsafe.Pointer(e))
}

func (e *Event) AsMouseButtonEvent() MouseButtonEvent {
	return *(*MouseButtonEvent)(unsafe.Pointer(e))
}

func (e *Event) AsMouseWheelEvent() MouseWheelEvent {
	return *(*MouseWheelEvent)(unsafe.Pointer(e))
}

func (e *Event) AsJoyDeviceEvent() JoyDeviceEvent {
	return *(*JoyDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) AsJoyAxisEvent() JoyAxisEvent {
	return *(*JoyAxisEvent)(unsafe.Pointer(e))
}

func (e *Event) AsJoyBallEvent() JoyBallEvent {
	return *(*JoyBallEvent)(unsafe.Pointer(e))
}

func (e *Event) AsJoyHatEvent() JoyHatEvent {
	return *(*JoyHatEvent)(unsafe.Pointer(e))
}

func (e *Event) AsJoyButtonEvent() JoyButtonEvent {
	return *(*JoyButtonEvent)(unsafe.Pointer(e))
}

func (e *Event) AsJoyBatteryEvent() JoyBatteryEvent {
	return *(*JoyBatteryEvent)(unsafe.Pointer(e))
}

func (e *Event) AsGamepadDeviceEvent() GamepadDeviceEvent {
	return *(*GamepadDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) AsGamepadAxisEvent() GamepadAxisEvent {
	return *(*GamepadAxisEvent)(unsafe.Pointer(e))
}

func (e *Event) AsGamepadButtonEvent() GamepadButtonEvent {
	return *(*GamepadButtonEvent)(unsafe.Pointer(e))
}

func (e *Event) AsGamepadTouchpadEvent() GamepadTouchpadEvent {
	return *(*GamepadTouchpadEvent)(unsafe.Pointer(e))
}

func (e *Event) AsGamepadSensorEvent() GamepadSensorEvent {
	return *(*GamepadSensorEvent)(unsafe.Pointer(e))
}

func (e *Event) AsAudioDeviceEvent() AudioDeviceEvent {
	return *(*AudioDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) AsCameraDeviceEvent() CameraDeviceEvent {
	return *(*CameraDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) AsSensorEvent() SensorEvent {
	return *(*SensorEvent)(unsafe.Pointer(e))
}

func (e *Event) AsQuitEvent() QuitEvent {
	return *(*QuitEvent)(unsafe.Pointer(e))
}

func (e *Event) AsUserEvent() UserEvent {
	return *(*UserEvent)(unsafe.Pointer(e))
}

func (e *Event) AsTouchFingerEvent() TouchFingerEvent {
	return *(*TouchFingerEvent)(unsafe.Pointer(e))
}

func (e *Event) AsPenProximityEvent() PenProximityEvent {
	return *(*PenProximityEvent)(unsafe.Pointer(e))
}

func (e *Event) AsPenTouchEvent() PenTouchEvent {
	return *(*PenTouchEvent)(unsafe.Pointer(e))
}

func (e *Event) AsPenMotionEvent() PenMotionEvent {
	return *(*PenMotionEvent)(unsafe.Pointer(e))
}

func (e *Event) AsPenButtonEvent() PenButtonEvent {
	return *(*PenButtonEvent)(unsafe.Pointer(e))
}

func (e *Event) AsPenAxisEvent() PenAxisEvent {
	return *(*PenAxisEvent)(unsafe.Pointer(e))
}

func (e *Event) AsRenderEvent() RenderEvent {
	return *(*RenderEvent)(unsafe.Pointer(e))
}

func (e *Event) AsDropEvent() DropEvent {
	return *(*DropEvent)(unsafe.Pointer(e))
}

func (e *Event) AsClipboardEvent() ClipboardEvent {
	return *(*ClipboardEvent)(unsafe.Pointer(e))
}

/* Function prototypes */

/**
 * Pump the event loop, gathering events from the input devices.
 *
 * This function updates the event queue and internal input device state.
 *
 * SDL_PumpEvents() gathers all the pending input information from devices and
 * places it in the event queue. Without calls to SDL_PumpEvents() no events
 * would ever be placed on the queue. Often the need for calls to
 * SDL_PumpEvents() is hidden from the user since SDL_PollEvent() and
 * SDL_WaitEvent() implicitly call SDL_PumpEvents(). However, if you are not
 * polling or waiting for events (e.g. you are filtering them), then you must
 * call SDL_PumpEvents() to force an event queue update.
 *
 * \threadsafety This function should only be called on the main thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_PollEvent
 * \sa SDL_WaitEvent
 */
//go:sdl3extern
var PumpEvents func()

/* @{ */

/**
 * The type of action to request from SDL_PeepEvents().
 *
 * \since This enum is available since SDL 3.2.0.
 */
type EventAction int

const (
	EA_AddEvent  EventAction = iota /**< Add events to the back of the queue. */
	EA_PeekEvent                    /**< Check but don't remove events from the queue front. */
	EA_GetEvent                     /**< Retrieve/remove events from the front of the queue. */
)

/**
 * Check the event queue for messages and optionally return them.
 *
 * `action` may be any of the following:
 *
 * - `SDL_ADDEVENT`: up to `numevents` events will be added to the back of the
 *   event queue.
 * - `SDL_PEEKEVENT`: `numevents` events at the front of the event queue,
 *   within the specified minimum and maximum type, will be returned to the
 *   caller and will _not_ be removed from the queue. If you pass NULL for
 *   `events`, then `numevents` is ignored and the total number of matching
 *   events will be returned.
 * - `SDL_GETEVENT`: up to `numevents` events at the front of the event queue,
 *   within the specified minimum and maximum type, will be returned to the
 *   caller and will be removed from the queue.
 *
 * You may have to call SDL_PumpEvents() before calling this function.
 * Otherwise, the events may not be ready to be filtered when you call
 * SDL_PeepEvents().
 *
 * \param events destination buffer for the retrieved events, may be NULL to
 *               leave the events in the queue and return the number of events
 *               that would have been stored.
 * \param numevents if action is SDL_ADDEVENT, the number of events to add
 *                  back to the event queue; if action is SDL_PEEKEVENT or
 *                  SDL_GETEVENT, the maximum number of events to retrieve.
 * \param action action to take; see [Remarks](#remarks) for details.
 * \param minType minimum value of the event type to be considered;
 *                SDL_EVENT_FIRST is a safe choice.
 * \param maxType maximum value of the event type to be considered;
 *                SDL_EVENT_LAST is a safe choice.
 * \returns the number of events actually stored or -1 on failure; call
 *          SDL_GetError() for more information.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_PollEvent
 * \sa SDL_PumpEvents
 * \sa SDL_PushEvent
 */
//go:sdl3extern
var PeepEvents func(events *Event, numevents int, action EventAction, minType, maxType uint32) int

/* @} */

/**
 * Check for the existence of a certain event type in the event queue.
 *
 * If you need to check for a range of event types, use SDL_HasEvents()
 * instead.
 *
 * \param type the type of event to be queried; see SDL_EventType for details.
 * \returns true if events matching `type` are present, or false if events
 *          matching `type` are not present.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_HasEvents
 */
//go:sdl3extern
var HasEvent func(eventType EventType) bool

/**
 * Check for the existence of certain event types in the event queue.
 *
 * If you need to check for a single event type, use SDL_HasEvent() instead.
 *
 * \param minType the low end of event type to be queried, inclusive; see
 *                SDL_EventType for details.
 * \param maxType the high end of event type to be queried, inclusive; see
 *                SDL_EventType for details.
 * \returns true if events with type >= `minType` and <= `maxType` are
 *          present, or false if not.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_HasEvents
 */
//go:sdl3extern
var HasEvents func(minType, maxType EventType) bool

/**
 * Clear events of a specific type from the event queue.
 *
 * This will unconditionally remove any events from the queue that match
 * `type`. If you need to remove a range of event types, use SDL_FlushEvents()
 * instead.
 *
 * It's also normal to just ignore events you don't care about in your event
 * loop without calling this function.
 *
 * This function only affects currently queued events. If you want to make
 * sure that all pending OS events are flushed, you can call SDL_PumpEvents()
 * on the main thread immediately before the flush call.
 *
 * If you have user events with custom data that needs to be freed, you should
 * use SDL_PeepEvents() to remove and clean up those events before calling
 * this function.
 *
 * \param type the type of event to be cleared; see SDL_EventType for details.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_FlushEvents
 */
//go:sdl3extern
var FlushEvent func(eventType EventType)

/**
 * Clear events of a range of types from the event queue.
 *
 * This will unconditionally remove any events from the queue that are in the
 * range of `minType` to `maxType`, inclusive. If you need to remove a single
 * event type, use SDL_FlushEvent() instead.
 *
 * It's also normal to just ignore events you don't care about in your event
 * loop without calling this function.
 *
 * This function only affects currently queued events. If you want to make
 * sure that all pending OS events are flushed, you can call SDL_PumpEvents()
 * on the main thread immediately before the flush call.
 *
 * \param minType the low end of event type to be cleared, inclusive; see
 *                SDL_EventType for details.
 * \param maxType the high end of event type to be cleared, inclusive; see
 *                SDL_EventType for details.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_FlushEvent
 */
//go:sdl3extern
var FlushEvents func(minType, maxType EventType)

/**
 * Poll for currently pending events.
 *
 * If `event` is not NULL, the next event is removed from the queue and stored
 * in the SDL_Event structure pointed to by `event`. The 1 returned refers to
 * this event, immediately stored in the SDL Event structure -- not an event
 * to follow.
 *
 * If `event` is NULL, it simply returns 1 if there is an event in the queue,
 * but will not remove it from the queue.
 *
 * As this function may implicitly call SDL_PumpEvents(), you can only call
 * this function in the thread that set the video mode.
 *
 * SDL_PollEvent() is the favored way of receiving system events since it can
 * be done from the main loop and does not suspend the main loop while waiting
 * on an event to be posted.
 *
 * The common practice is to fully process the event queue once every frame,
 * usually as a first step before updating the game's state:
 *
 * ```c
 * while (game_is_still_running) {
 *     SDL_Event event;
 *     while (SDL_PollEvent(&event)) {  // poll until all events are handled!
 *         // decide what to do with this event.
 *     }
 *
 *     // update game state, draw the current frame
 * }
 * ```
 *
 * \param event the SDL_Event structure to be filled with the next event from
 *              the queue, or NULL.
 * \returns true if this got an event or false if there are none available.
 *
 * \threadsafety This function should only be called on the main thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_PushEvent
 * \sa SDL_WaitEvent
 * \sa SDL_WaitEventTimeout
 */
//go:sdl3extern
var PollEvent func(event *Event) bool

/**
 * Wait indefinitely for the next available event.
 *
 * If `event` is not NULL, the next event is removed from the queue and stored
 * in the SDL_Event structure pointed to by `event`.
 *
 * As this function may implicitly call SDL_PumpEvents(), you can only call
 * this function in the thread that initialized the video subsystem.
 *
 * \param event the SDL_Event structure to be filled in with the next event
 *              from the queue, or NULL.
 * \returns true on success or false if there was an error while waiting for
 *          events; call SDL_GetError() for more information.
 *
 * \threadsafety This function should only be called on the main thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_PollEvent
 * \sa SDL_PushEvent
 * \sa SDL_WaitEventTimeout
 */
//go:sdl3extern
var WaitEvent func(event *Event) bool

/**
 * Wait until the specified timeout (in milliseconds) for the next available
 * event.
 *
 * If `event` is not NULL, the next event is removed from the queue and stored
 * in the SDL_Event structure pointed to by `event`.
 *
 * As this function may implicitly call SDL_PumpEvents(), you can only call
 * this function in the thread that initialized the video subsystem.
 *
 * The timeout is not guaranteed, the actual wait time could be longer due to
 * system scheduling.
 *
 * \param event the SDL_Event structure to be filled in with the next event
 *              from the queue, or NULL.
 * \param timeoutMS the maximum number of milliseconds to wait for the next
 *                  available event.
 * \returns true if this got an event or false if the timeout elapsed without
 *          any events available.
 *
 * \threadsafety This function should only be called on the main thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_PollEvent
 * \sa SDL_PushEvent
 * \sa SDL_WaitEvent
 */
//go:sdl3extern
var WaitEventTimeout func(event *Event, timeoutMS int32) bool

/**
 * Add an event to the event queue.
 *
 * The event queue can actually be used as a two way communication channel.
 * Not only can events be read from the queue, but the user can also push
 * their own events onto it. `event` is a pointer to the event structure you
 * wish to push onto the queue. The event is copied into the queue, and the
 * caller may dispose of the memory pointed to after SDL_PushEvent() returns.
 *
 * Note: Pushing device input events onto the queue doesn't modify the state
 * of the device within SDL.
 *
 * Note: Events pushed onto the queue with SDL_PushEvent() get passed through
 * the event filter but events added with SDL_PeepEvents() do not.
 *
 * For pushing application-specific events, please use SDL_RegisterEvents() to
 * get an event type that does not conflict with other code that also wants
 * its own custom event types.
 *
 * \param event the SDL_Event to be added to the queue.
 * \returns true on success, false if the event was filtered or on failure;
 *          call SDL_GetError() for more information. A common reason for
 *          error is the event queue being full.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_PeepEvents
 * \sa SDL_PollEvent
 * \sa SDL_RegisterEvents
 */
//go:sdl3extern
var PushEvent func(event *Event) bool

/**
 * A function pointer used for callbacks that watch the event queue.
 *
 * \param userdata what was passed as `userdata` to SDL_SetEventFilter() or
 *                 SDL_AddEventWatch, etc.
 * \param event the event that triggered the callback.
 * \returns true to permit event to be added to the queue, and false to
 *          disallow it. When used with SDL_AddEventWatch, the return value is
 *          ignored.
 *
 * \threadsafety SDL may call this callback at any time from any thread; the
 *               application is responsible for locking resources the callback
 *               touches that need to be protected.
 *
 * \since This datatype is available since SDL 3.2.0.
 *
 * \sa SDL_SetEventFilter
 * \sa SDL_AddEventWatch
 */
type EventFilter func(userdata uintptr, event *Event) bool

/**
 * Set up a filter to process all events before they are added to the internal
 * event queue.
 *
 * If you just want to see events without modifying them or preventing them
 * from being queued, you should use SDL_AddEventWatch() instead.
 *
 * If the filter function returns true when called, then the event will be
 * added to the internal queue. If it returns false, then the event will be
 * dropped from the queue, but the internal state will still be updated. This
 * allows selective filtering of dynamically arriving events.
 *
 * **WARNING**: Be very careful of what you do in the event filter function,
 * as it may run in a different thread!
 *
 * On platforms that support it, if the quit event is generated by an
 * interrupt signal (e.g. pressing Ctrl-C), it will be delivered to the
 * application at the next event poll.
 *
 * Note: Disabled events never make it to the event filter function; see
 * SDL_SetEventEnabled().
 *
 * Note: Events pushed onto the queue with SDL_PushEvent() get passed through
 * the event filter, but events pushed onto the queue with SDL_PeepEvents() do
 * not.
 *
 * \param filter an SDL_EventFilter function to call when an event happens.
 * \param userdata a pointer that is passed to `filter`.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddEventWatch
 * \sa SDL_SetEventEnabled
 * \sa SDL_GetEventFilter
 * \sa SDL_PeepEvents
 * \sa SDL_PushEvent
 */
//go:sdl3extern
var SetEventFilter func(filter EventFilter, userdata uintptr)

/**
 * Query the current event filter.
 *
 * This function can be used to "chain" filters, by saving the existing filter
 * before replacing it with a function that will call that saved filter.
 *
 * \param filter the current callback function will be stored here.
 * \param userdata the pointer that is passed to the current event filter will
 *                 be stored here.
 * \returns true on success or false if there is no event filter set.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SetEventFilter
 */
//go:sdl3extern
var GetEventFilter func(filter EventFilter, userdata *uintptr) bool

/**
 * Add a callback to be triggered when an event is added to the event queue.
 *
 * `filter` will be called when an event happens, and its return value is
 * ignored.
 *
 * **WARNING**: Be very careful of what you do in the event filter function,
 * as it may run in a different thread!
 *
 * If the quit event is generated by a signal (e.g. SIGINT), it will bypass
 * the internal queue and be delivered to the watch callback immediately, and
 * arrive at the next event poll.
 *
 * Note: the callback is called for events posted by the user through
 * SDL_PushEvent(), but not for disabled events, nor for events by a filter
 * callback set with SDL_SetEventFilter(), nor for events posted by the user
 * through SDL_PeepEvents().
 *
 * \param filter an SDL_EventFilter function to call when an event happens.
 * \param userdata a pointer that is passed to `filter`.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_RemoveEventWatch
 * \sa SDL_SetEventFilter
 */
//go:sdl3extern
var AddEventWatch func(filter EventFilter, userdata uintptr) bool

/**
 * Remove an event watch callback added with SDL_AddEventWatch().
 *
 * This function takes the same input as SDL_AddEventWatch() to identify and
 * delete the corresponding callback.
 *
 * \param filter the function originally passed to SDL_AddEventWatch().
 * \param userdata the pointer originally passed to SDL_AddEventWatch().
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddEventWatch
 */
//go:sdl3extern
var RemoveEventWatch func(filter EventFilter, userdata uintptr)

/**
 * Run a specific filter function on the current event queue, removing any
 * events for which the filter returns false.
 *
 * See SDL_SetEventFilter() for more information. Unlike SDL_SetEventFilter(),
 * this function does not change the filter permanently, it only uses the
 * supplied filter until this function returns.
 *
 * \param filter the SDL_EventFilter function to call when an event happens.
 * \param userdata a pointer that is passed to `filter`.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetEventFilter
 * \sa SDL_SetEventFilter
 */
//go:sdl3extern
var FilterEvents func(filter EventFilter, userdata uintptr)

/**
 * Set the state of processing events by type.
 *
 * \param type the type of event; see SDL_EventType for details.
 * \param enabled whether to process the event or not.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_EventEnabled
 */
//go:sdl3extern
var SetEventEnabled func(eventType EventType, enabled bool)

/**
 * Query the state of processing events by type.
 *
 * \param type the type of event; see SDL_EventType for details.
 * \returns true if the event is being processed, false otherwise.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SetEventEnabled
 */
//go:sdl3extern
var EventEnabled func(eventType EventType) bool

/**
 * Allocate a set of user-defined events, and return the beginning event
 * number for that set of events.
 *
 * \param numevents the number of events to be allocated.
 * \returns the beginning event number, or 0 if numevents is invalid or if
 *          there are not enough user-defined events left.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_PushEvent
 */
//go:sdl3extern
var RegisterEvents func(numevents int) uint32

/**
 * Get window associated with an event.
 *
 * \param event an event containing a `windowID`.
 * \returns the associated window on success or NULL if there is none.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_PollEvent
 * \sa SDL_WaitEvent
 * \sa SDL_WaitEventTimeout
 */
//go:sdl3extern
var GetWindowFromEvent func(event *Event) *Window

/**
 * Generate a human-readable description of an event.
 *
 * This will fill `buf` with a null-terminated string that might look
 * something like this:
 *
 * ```
 * SDL_EVENT_MOUSE_MOTION (timestamp=1140256324 windowid=2 which=0 state=0 x=492.99 y=139.09 xrel=52 yrel=6)
 * ```
 *
 * The exact format of the string is not guaranteed; it is intended for
 * logging purposes, to be read by a human, and not parsed by a computer.
 *
 * The returned value follows the same rules as SDL_snprintf(): `buf` will
 * always be NULL-terminated (unless `buflen` is zero), and will be truncated
 * if `buflen` is too small. The return code is the number of bytes needed for
 * the complete string, not counting the NULL-terminator, whether the string
 * was truncated or not. Unlike SDL_snprintf(), though, this function never
 * returns -1.
 *
 * \param event an event to describe. May be NULL.
 * \param buf the buffer to fill with the description string. May be NULL.
 * \param buflen the maximum bytes that can be written to `buf`.
 * \returns number of bytes needed for the full string, not counting the
 *          null-terminator byte.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.4.0.
 */
//go:sdl3extern
var GetEventDescription func(event *Event, buf []byte, buflen int) int
