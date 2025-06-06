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
 * # CategoryGamepad
 *
 * SDL provides a low-level joystick API, which just treats joysticks as an
 * arbitrary pile of buttons, axes, and hat switches. If you're planning to
 * write your own control configuration screen, this can give you a lot of
 * flexibility, but that's a lot of work, and most things that we consider
 * "joysticks" now are actually console-style gamepads. So SDL provides the
 * gamepad API on top of the lower-level joystick functionality.
 *
 * The difference between a joystick and a gamepad is that a gamepad tells you
 * _where_ a button or axis is on the device. You don't speak to gamepads in
 * terms of arbitrary numbers like "button 3" or "axis 2" but in standard
 * locations: the d-pad, the shoulder buttons, triggers, A/B/X/Y (or
 * X/O/Square/Triangle, if you will).
 *
 * One turns a joystick into a gamepad by providing a magic configuration
 * string, which tells SDL the details of a specific device: when you see this
 * specific hardware, if button 2 gets pressed, this is actually D-Pad Up,
 * etc.
 *
 * SDL has many popular controllers configured out of the box, and users can
 * add their own controller details through an environment variable if it's
 * otherwise unknown to SDL.
 *
 * In order to use these functions, SDL_Init() must have been called with the
 * SDL_INIT_GAMEPAD flag. This causes SDL to scan the system for gamepads, and
 * load appropriate drivers.
 *
 * If you would like to receive gamepad updates while the application is in
 * the background, you should set the following hint before calling
 * SDL_Init(): SDL_HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS
 *
 * Gamepads support various optional features such as rumble, color LEDs,
 * touchpad, gyro, etc. The support for these features varies depending on the
 * controller and OS support available. You can check for LED and rumble
 * capabilities at runtime by calling SDL_GetGamepadProperties() and checking
 * the various capability properties. You can check for touchpad by calling
 * SDL_GetNumGamepadTouchpads() and check for gyro and accelerometer by
 * calling SDL_GamepadHasSensor().
 *
 * By default SDL will try to use the most capable driver available, but you
 * can tune which OS drivers to use with the various joystick hints in
 * SDL_hints.h.
 *
 * Your application should always support gamepad hotplugging. On some
 * platforms like Xbox, Steam Deck, etc., this is a requirement for
 * certification. On other platforms, like macOS and Windows when using
 * Windows.Gaming.Input, controllers may not be available at startup and will
 * come in at some point after you've started processing events.
 */
package sdl

/**
 * The structure used to identify an SDL gamepad
 *
 * \since This struct is available since SDL 3.2.0.
 */
type Gamepad uintptr

/**
 * Standard gamepad types.
 *
 * This type does not necessarily map to first-party controllers from
 * Microsoft/Sony/Nintendo; in many cases, third-party controllers can report
 * as these, either because they were designed for a specific console, or they
 * simply most closely match that console's controllers (does it have A/B/X/Y
 * buttons or X/O/Square/Triangle? Does it have a touchpad? etc).
 */
type GamepadType int32

const (
	GamepadTypeUnknown GamepadType = iota
	GamepadTypeStandard
	GamepadTypeXbox360
	GamepadTypeXboxOne
	GamepadTypePS3
	GamepadTypePS4
	GamepadTypePS5
	GamepadTypeNintendoSwitchPro
	GamepadTypeNintendoSwitchJoyconLeft
	GamepadTypeNintendoSwitchJoyconRight
	GamepadTypeNintendoSwitchJoyconPair
	GamepadTypeGamecube
	GamepadTypeCount
)

/**
 * The list of buttons available on a gamepad
 *
 * For controllers that use a diamond pattern for the face buttons, the
 * south/east/west/north buttons below correspond to the locations in the
 * diamond pattern. For Xbox controllers, this would be A/B/X/Y, for Nintendo
 * Switch controllers, this would be B/A/Y/X, for GameCube controllers this
 * would be A/X/B/Y, for PlayStation controllers this would be
 * Cross/Circle/Square/Triangle.
 *
 * For controllers that don't use a diamond pattern for the face buttons, the
 * south/east/west/north buttons indicate the buttons labeled A, B, C, D, or
 * 1, 2, 3, 4, or for controllers that aren't labeled, they are the primary,
 * secondary, etc. buttons.
 *
 * The activate action is often the south button and the cancel action is
 * often the east button, but in some regions this is reversed, so your game
 * should allow remapping actions based on user preferences.
 *
 * You can query the labels for the face buttons using
 * SDL_GetGamepadButtonLabel()
 *
 * \since This enum is available since SDL 3.2.0.
 */
type GamepadButton int32

const (
	GamepadButtonInvalid GamepadButton = iota - 1
	GamepadButtonSouth                 /**< Bottom face button (e.g. Xbox A button) */
	GamepadButtonEast                  /**< Right face button (e.g. Xbox B button) */
	GamepadButtonWest                  /**< Left face button (e.g. Xbox X button) */
	GamepadButtonNorth                 /**< Top face button (e.g. Xbox Y button) */
	GamepadButtonBack
	GamepadButtonGuide
	GamepadButtonStart
	GamepadButtonLeftStick
	GamepadButtonRightStick
	GamepadButtonLeftShoulder
	GamepadButtonRightShoulder
	GamepadButtonDpadUp
	GamepadButtonDpadDown
	GamepadButtonDpadLeft
	GamepadButtonDpadRight
	GamepadButtonMisc1        /**< Additional button (e.g. Xbox Series X share button, PS5 microphone button, Nintendo Switch Pro capture button, Amazon Luna microphone button, Google Stadia capture button) */
	GamepadButtonRightPaddle1 /**< Upper or primary paddle, under your right hand (e.g. Xbox Elite paddle P1) */
	GamepadButtonLeftPaddle1  /**< Upper or primary paddle, under your left hand (e.g. Xbox Elite paddle P3) */
	GamepadButtonRightPaddle2 /**< Lower or secondary paddle, under your right hand (e.g. Xbox Elite paddle P2) */
	GamepadButtonLeftPaddle2  /**< Lower or secondary paddle, under your left hand (e.g. Xbox Elite paddle P4) */
	GamepadButtonTouchpad     /**< PS4/PS5 touchpad button */
	GamepadButtonMisc2        /**< Additional button */
	GamepadButtonMisc3        /**< Additional button */
	GamepadButtonMisc4        /**< Additional button */
	GamepadButtonMisc5        /**< Additional button */
	GamepadButtonMisc6        /**< Additional button */
	GamepadButtonCount
)

/**
 * The set of gamepad button labels
 *
 * This isn't a complete set, just the face buttons to make it easy to show
 * button prompts.
 *
 * For a complete set, you should look at the button and gamepad type and have
 * a set of symbols that work well with your art style.
 *
 * \since This enum is available since SDL 3.2.0.
 */
type GamepadButtonLabel int32

const (
	GamepadButtonLabelUnknown GamepadButtonLabel = iota
	GamepadButtonLabelA
	GamepadButtonLabelB
	GamepadButtonLabelX
	GamepadButtonLabelY
	GamepadButtonLabelCross
	GamepadButtonLabelCircle
	GamepadButtonLabelSquare
	GamepadButtonLabelTriangle
)

/**
 * The list of axes available on a gamepad
 *
 * Thumbstick axis values range from SDL_JOYSTICK_AXIS_MIN to
 * SDL_JOYSTICK_AXIS_MAX, and are centered within ~8000 of zero, though
 * advanced UI will allow users to set or autodetect the dead zone, which
 * varies between gamepads.
 *
 * Trigger axis values range from 0 (released) to SDL_JOYSTICK_AXIS_MAX (fully
 * pressed) when reported by SDL_GetGamepadAxis(). Note that this is not the
 * same range that will be reported by the lower-level SDL_GetJoystickAxis().
 *
 * \since This enum is available since SDL 3.2.0.
 */
type GamepadAxis int32

const (
	GamepadAxisInvalid GamepadAxis = iota - 1
	GamepadAxisLeftx
	GamepadAxisLefty
	GamepadAxisRightX
	GamepadAxisRightY
	GamepadAxisLeftTrigger
	GamepadAxisRightTrigger
	GamepadAxisCount
)

/**
 * Types of gamepad control bindings.
 *
 * A gamepad is a collection of bindings that map arbitrary joystick buttons,
 * axes and hat switches to specific positions on a generic console-style
 * gamepad. This enum is used as part of SDL_GamepadBinding to specify those
 * mappings.
 *
 * \since This enum is available since SDL 3.2.0.
 */
type GamepadBindingType int32

const (
	GamepadBindTypeNone GamepadBindingType = iota
	GamepadBindTypeButton
	GamepadBindTypeAxis
	GamepadBindTypeHat
)

/**
 * A mapping between one joystick input to a gamepad control.
 *
 * A gamepad has a collection of several bindings, to say, for example, when
 * joystick button number 5 is pressed, that should be treated like the
 * gamepad's "start" button.
 *
 * SDL has these bindings built-in for many popular controllers, and can add
 * more with a simple text string. Those strings are parsed into a collection
 * of these structs to make it easier to operate on the data.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadBindings
 */
type GamepadBinding struct {
	InputType GamepadBindingType
	inputA    int32
	inputB    int32
	inputC    int32

	OutputType GamepadBindingType
	outputA    int32
	outputB    int32
	outputC    int32
}

func (b *GamepadBinding) InputButton() *int32 {
	return &b.inputA
}

func (b *GamepadBinding) InputAxis() *int32 {
	return &b.inputA
}

func (b *GamepadBinding) InputAxisMin() *int32 {
	return &b.inputB
}

func (b *GamepadBinding) InputAxisMax() *int32 {
	return &b.inputC
}

func (b *GamepadBinding) InputHat() *int32 {
	return &b.inputA
}

func (b *GamepadBinding) InputHatMask() *int32 {
	return &b.inputB
}

func (b *GamepadBinding) OutputButton() *GamepadButton {
	return (*GamepadButton)(&b.outputA)
}

func (b *GamepadBinding) OutputAxis() *GamepadAxis {
	return (*GamepadAxis)(&b.outputA)
}

func (b *GamepadBinding) OutputAxisMin() *int32 {
	return &b.outputB
}

func (b *GamepadBinding) OutputAxisMax() *int32 {
	return &b.outputC
}

/**
 * Add support for gamepads that SDL is unaware of or change the binding of an
 * existing gamepad.
 *
 * The mapping string has the format "GUID,name,mapping", where GUID is the
 * string value from SDL_GUIDToString(), name is the human readable string for
 * the device and mappings are gamepad mappings to joystick ones. Under
 * Windows there is a reserved GUID of "xinput" that covers all XInput
 * devices. The mapping format for joystick is:
 *
 * - `bX`: a joystick button, index X
 * - `hX.Y`: hat X with value Y
 * - `aX`: axis X of the joystick
 *
 * Buttons can be used as a gamepad axes and vice versa.
 *
 * If a device with this GUID is already plugged in, SDL will generate an
 * SDL_EVENT_GAMEPAD_ADDED event.
 *
 * This string shows an example of a valid mapping for a gamepad:
 *
 * ```c
 * "341a3608000000000000504944564944,Afterglow PS3 Controller,a:b1,b:b2,y:b3,x:b0,start:b9,guide:b12,back:b8,dpup:h0.1,dpleft:h0.8,dpdown:h0.4,dpright:h0.2,leftshoulder:b4,rightshoulder:b5,leftstick:b10,rightstick:b11,leftx:a0,lefty:a1,rightx:a2,righty:a3,lefttrigger:b6,righttrigger:b7"
 * ```
 *
 * \param mapping the mapping string.
 * \returns 1 if a new mapping is added, 0 if an existing mapping is updated,
 *          -1 on failure; call SDL_GetError() for more information.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddGamepadMappingsFromFile
 * \sa SDL_AddGamepadMappingsFromIO
 * \sa SDL_GetGamepadMapping
 * \sa SDL_GetGamepadMappingForGUID
 * \sa SDL_HINT_GAMECONTROLLERCONFIG
 * \sa SDL_HINT_GAMECONTROLLERCONFIG_FILE
 * \sa SDL_EVENT_GAMEPAD_ADDED
 */
//go:sdl3extern
var AddGamepadMapping func(mapping string) int32

/**
 * Load a set of gamepad mappings from an SDL_IOStream.
 *
 * You can call this function several times, if needed, to load different
 * database files.
 *
 * If a new mapping is loaded for an already known gamepad GUID, the later
 * version will overwrite the one currently loaded.
 *
 * Any new mappings for already plugged in controllers will generate
 * SDL_EVENT_GAMEPAD_ADDED events.
 *
 * Mappings not belonging to the current platform or with no platform field
 * specified will be ignored (i.e. mappings for Linux will be ignored in
 * Windows, etc).
 *
 * This function will load the text database entirely in memory before
 * processing it, so take this into consideration if you are in a memory
 * constrained environment.
 *
 * \param src the data stream for the mappings to be added.
 * \param closeio if true, calls SDL_CloseIO() on `src` before returning, even
 *                in the case of an error.
 * \returns the number of mappings added or -1 on failure; call SDL_GetError()
 *          for more information.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddGamepadMapping
 * \sa SDL_AddGamepadMappingsFromFile
 * \sa SDL_GetGamepadMapping
 * \sa SDL_GetGamepadMappingForGUID
 * \sa SDL_HINT_GAMECONTROLLERCONFIG
 * \sa SDL_HINT_GAMECONTROLLERCONFIG_FILE
 * \sa SDL_EVENT_GAMEPAD_ADDED
 */
//go:sdl3extern
var AddGamepadMappingsFromIO func(src IOStream, closeio bool) int32

/**
 * Load a set of gamepad mappings from a file.
 *
 * You can call this function several times, if needed, to load different
 * database files.
 *
 * If a new mapping is loaded for an already known gamepad GUID, the later
 * version will overwrite the one currently loaded.
 *
 * Any new mappings for already plugged in controllers will generate
 * SDL_EVENT_GAMEPAD_ADDED events.
 *
 * Mappings not belonging to the current platform or with no platform field
 * specified will be ignored (i.e. mappings for Linux will be ignored in
 * Windows, etc).
 *
 * \param file the mappings file to load.
 * \returns the number of mappings added or -1 on failure; call SDL_GetError()
 *          for more information.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddGamepadMapping
 * \sa SDL_AddGamepadMappingsFromIO
 * \sa SDL_GetGamepadMapping
 * \sa SDL_GetGamepadMappingForGUID
 * \sa SDL_HINT_GAMECONTROLLERCONFIG
 * \sa SDL_HINT_GAMECONTROLLERCONFIG_FILE
 * \sa SDL_EVENT_GAMEPAD_ADDED
 */
//go:sdl3extern
var AddGamepadMappingsFromFile func(file string) int32

/**
 * Reinitialize the SDL mapping database to its initial state.
 *
 * This will generate gamepad events as needed if device mappings change.
 *
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var ReloadGamepadMappings func() bool

/**
 * Get the current gamepad mappings.
 *
 * \param count a pointer filled in with the number of mappings returned, can
 *              be NULL.
 * \returns an array of the mapping strings, NULL-terminated, or NULL on
 *          failure; call SDL_GetError() for more information. This is a
 *          single allocation that should be freed with SDL_free() when it is
 *          no longer needed.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadMappings func(count *int32) []string

/**
 * Get the gamepad mapping string for a given GUID.
 *
 * \param guid a structure containing the GUID for which a mapping is desired.
 * \returns a mapping string or NULL on failure; call SDL_GetError() for more
 *          information. This should be freed with SDL_free() when it is no
 *          longer needed.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetJoystickGUIDForID
 * \sa SDL_GetJoystickGUID
 */
//go:sdl3extern
var GetGamepadMappingForGUID func(guid *GUID) string

/**
 * Get the current mapping of a gamepad.
 *
 * Details about mappings are discussed with SDL_AddGamepadMapping().
 *
 * \param gamepad the gamepad you want to get the current mapping for.
 * \returns a string that has the gamepad's mapping or NULL if no mapping is
 *          available; call SDL_GetError() for more information. This should
 *          be freed with SDL_free() when it is no longer needed.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddGamepadMapping
 * \sa SDL_GetGamepadMappingForID
 * \sa SDL_GetGamepadMappingForGUID
 * \sa SDL_SetGamepadMapping
 */
//go:sdl3extern
var GetGamepadMapping func(gamepad Gamepad) string

/**
 * Set the current mapping of a joystick or gamepad.
 *
 * Details about mappings are discussed with SDL_AddGamepadMapping().
 *
 * \param instance_id the joystick instance ID.
 * \param mapping the mapping to use for this device, or NULL to clear the
 *                mapping.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AddGamepadMapping
 * \sa SDL_GetGamepadMapping
 */
//go:sdl3extern
var SetGamepadMapping func(instance_id JoystickID, mapping string) bool

/**
 * Return whether a gamepad is currently connected.
 *
 * \returns true if a gamepad is connected, false otherwise.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepads
 */
//go:sdl3extern
var HasGamepad func() bool

/**
 * Get a list of currently connected gamepads.
 *
 * \param count a pointer filled in with the number of gamepads returned, may
 *              be NULL.
 * \returns a 0 terminated array of joystick instance IDs or NULL on failure;
 *          call SDL_GetError() for more information. This should be freed
 *          with SDL_free() when it is no longer needed.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_HasGamepad
 * \sa SDL_OpenGamepad
 */
//go:sdl3extern
var GetGamepads func(count *int32) *JoystickID

/**
 * Check if the given joystick is supported by the gamepad interface.
 *
 * \param instance_id the joystick instance ID.
 * \returns true if the given joystick is supported by the gamepad interface,
 *          false if it isn't or it's an invalid index.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetJoysticks
 * \sa SDL_OpenGamepad
 */
//go:sdl3extern
var IsGamepad func(instance_id JoystickID) bool

/**
 * Get the implementation dependent name of a gamepad.
 *
 * This can be called before any gamepads are opened.
 *
 * \param instance_id the joystick instance ID.
 * \returns the name of the selected gamepad. If no name can be found, this
 *          function returns NULL; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadName
 * \sa SDL_GetGamepads
 */
//go:sdl3extern
var GetGamepadNameForID func(instance_id JoystickID) string

/**
 * Get the implementation dependent path of a gamepad.
 *
 * This can be called before any gamepads are opened.
 *
 * \param instance_id the joystick instance ID.
 * \returns the path of the selected gamepad. If no path can be found, this
 *          function returns NULL; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadPath
 * \sa SDL_GetGamepads
 */
//go:sdl3extern
var GetGamepadPathForID func(instance_id JoystickID) string

/**
 * Get the player index of a gamepad.
 *
 * This can be called before any gamepads are opened.
 *
 * \param instance_id the joystick instance ID.
 * \returns the player index of a gamepad, or -1 if it's not available.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadPlayerIndex
 * \sa SDL_GetGamepads
 */
//go:sdl3extern
var GetGamepadPlayerIndexForID func(instance_id JoystickID) int32

/**
 * Get the implementation-dependent GUID of a gamepad.
 *
 * This can be called before any gamepads are opened.
 *
 * \param instance_id the joystick instance ID.
 * \returns the GUID of the selected gamepad. If called on an invalid index,
 *          this function returns a zero GUID.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GUIDToString
 * \sa SDL_GetGamepads
 */
//go:sdl3extern
var GetGamepadGUIDForID func(instance_id JoystickID) *GUID

/**
 * Get the USB vendor ID of a gamepad, if available.
 *
 * This can be called before any gamepads are opened. If the vendor ID isn't
 * available this function returns 0.
 *
 * \param instance_id the joystick instance ID.
 * \returns the USB vendor ID of the selected gamepad. If called on an invalid
 *          index, this function returns zero.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadVendor
 * \sa SDL_GetGamepads
 */
//go:sdl3extern
var GetGamepadVendorForID func(instance_id JoystickID) uint16

/**
 * Get the USB product ID of a gamepad, if available.
 *
 * This can be called before any gamepads are opened. If the product ID isn't
 * available this function returns 0.
 *
 * \param instance_id the joystick instance ID.
 * \returns the USB product ID of the selected gamepad. If called on an
 *          invalid index, this function returns zero.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadProduct
 * \sa SDL_GetGamepads
 */
//go:sdl3extern
var GetGamepadProductForID func(instance_id JoystickID) uint16

/**
 * Get the product version of a gamepad, if available.
 *
 * This can be called before any gamepads are opened. If the product version
 * isn't available this function returns 0.
 *
 * \param instance_id the joystick instance ID.
 * \returns the product version of the selected gamepad. If called on an
 *          invalid index, this function returns zero.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadProductVersion
 * \sa SDL_GetGamepads
 */
//go:sdl3extern
var GetGamepadProductVersionForID func(instance_id JoystickID) uint16

/**
 * Get the type of a gamepad.
 *
 * This can be called before any gamepads are opened.
 *
 * \param instance_id the joystick instance ID.
 * \returns the gamepad type.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadType
 * \sa SDL_GetGamepads
 * \sa SDL_GetRealGamepadTypeForID
 */
//go:sdl3extern
var GetGamepadTypeForID func(instance_id JoystickID) GamepadType

/**
 * Get the type of a gamepad, ignoring any mapping override.
 *
 * This can be called before any gamepads are opened.
 *
 * \param instance_id the joystick instance ID.
 * \returns the gamepad type.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadTypeForID
 * \sa SDL_GetGamepads
 * \sa SDL_GetRealGamepadType
 */
//go:sdl3extern
var GetRealGamepadTypeForID func(instance_id JoystickID) GamepadType

/**
 * Get the mapping of a gamepad.
 *
 * This can be called before any gamepads are opened.
 *
 * \param instance_id the joystick instance ID.
 * \returns the mapping string. Returns NULL if no mapping is available. This
 *          should be freed with SDL_free() when it is no longer needed.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepads
 * \sa SDL_GetGamepadMapping
 */
//go:sdl3extern
var GetGamepadMappingForID func(instance_id JoystickID) string

/**
 * Open a gamepad for use.
 *
 * \param instance_id the joystick instance ID.
 * \returns a gamepad identifier or NULL if an error occurred; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CloseGamepad
 * \sa SDL_IsGamepad
 */
//go:sdl3extern
var OpenGamepad func(instance_id JoystickID) Gamepad

/**
 * Get the SDL_Gamepad associated with a joystick instance ID, if it has been
 * opened.
 *
 * \param instance_id the joystick instance ID of the gamepad.
 * \returns an SDL_Gamepad on success or NULL on failure or if it hasn't been
 *          opened yet; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadFromID func(instance_id JoystickID) Gamepad

/**
 * Get the SDL_Gamepad associated with a player index.
 *
 * \param player_index the player index, which different from the instance ID.
 * \returns the SDL_Gamepad associated with a player index.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadPlayerIndex
 * \sa SDL_SetGamepadPlayerIndex
 */
//go:sdl3extern
var GetGamepadFromPlayerIndex func(player_index int32) Gamepad

/**
 * Get the properties associated with an opened gamepad.
 *
 * These properties are shared with the underlying joystick object.
 *
 * The following read-only properties are provided by SDL:
 *
 * - `SDL_PROP_GAMEPAD_CAP_MONO_LED_BOOLEAN`: true if this gamepad has an LED
 *   that has adjustable brightness
 * - `SDL_PROP_GAMEPAD_CAP_RGB_LED_BOOLEAN`: true if this gamepad has an LED
 *   that has adjustable color
 * - `SDL_PROP_GAMEPAD_CAP_PLAYER_LED_BOOLEAN`: true if this gamepad has a
 *   player LED
 * - `SDL_PROP_GAMEPAD_CAP_RUMBLE_BOOLEAN`: true if this gamepad has
 *   left/right rumble
 * - `SDL_PROP_GAMEPAD_CAP_TRIGGER_RUMBLE_BOOLEAN`: true if this gamepad has
 *   simple trigger rumble
 *
 * \param gamepad a gamepad identifier previously returned by
 *                SDL_OpenGamepad().
 * \returns a valid property ID on success or 0 on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadProperties func(gamepad Gamepad) PropertiesID

const (
	PROP_GAMEPAD_CAP_MONO_LED_BOOLEAN       PropertyName = PROP_JOYSTICK_CAP_MONO_LED_BOOLEAN
	PROP_GAMEPAD_CAP_RGB_LED_BOOLEAN        PropertyName = PROP_JOYSTICK_CAP_RGB_LED_BOOLEAN
	PROP_GAMEPAD_CAP_PLAYER_LED_BOOLEAN     PropertyName = PROP_JOYSTICK_CAP_PLAYER_LED_BOOLEAN
	PROP_GAMEPAD_CAP_RUMBLE_BOOLEAN         PropertyName = PROP_JOYSTICK_CAP_RUMBLE_BOOLEAN
	PROP_GAMEPAD_CAP_TRIGGER_RUMBLE_BOOLEAN PropertyName = PROP_JOYSTICK_CAP_TRIGGER_RUMBLE_BOOLEAN
)

/**
 * Get the instance ID of an opened gamepad.
 *
 * \param gamepad a gamepad identifier previously returned by
 *                SDL_OpenGamepad().
 * \returns the instance ID of the specified gamepad on success or 0 on
 *          failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadID func(gamepad Gamepad) JoystickID

/**
 * Get the implementation-dependent name for an opened gamepad.
 *
 * \param gamepad a gamepad identifier previously returned by
 *                SDL_OpenGamepad().
 * \returns the implementation dependent name for the gamepad, or NULL if
 *          there is no name or the identifier passed is invalid.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadNameForID
 */
//go:sdl3extern
var GetGamepadName func(gamepad Gamepad) string

/**
 * Get the implementation-dependent path for an opened gamepad.
 *
 * \param gamepad a gamepad identifier previously returned by
 *                SDL_OpenGamepad().
 * \returns the implementation dependent path for the gamepad, or NULL if
 *          there is no path or the identifier passed is invalid.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadPathForID
 */
//go:sdl3extern
var GetGamepadPath func(gamepad Gamepad) string

/**
 * Get the type of an opened gamepad.
 *
 * \param gamepad the gamepad object to query.
 * \returns the gamepad type, or SDL_GAMEPAD_TYPE_UNKNOWN if it's not
 *          available.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadTypeForID
 */
//go:sdl3extern
var GetGamepadType func(gamepad Gamepad) GamepadType

/**
 * Get the type of an opened gamepad, ignoring any mapping override.
 *
 * \param gamepad the gamepad object to query.
 * \returns the gamepad type, or SDL_GAMEPAD_TYPE_UNKNOWN if it's not
 *          available.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetRealGamepadTypeForID
 */
//go:sdl3extern
var GetRealGamepadType func(gamepad Gamepad) GamepadType

/**
 * Get the player index of an opened gamepad.
 *
 * For XInput gamepads this returns the XInput user index.
 *
 * \param gamepad the gamepad object to query.
 * \returns the player index for gamepad, or -1 if it's not available.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SetGamepadPlayerIndex
 */
//go:sdl3extern
var GetGamepadPlayerIndex func(gamepad Gamepad) int32

/**
 * Set the player index of an opened gamepad.
 *
 * \param gamepad the gamepad object to adjust.
 * \param player_index player index to assign to this gamepad, or -1 to clear
 *                     the player index and turn off player LEDs.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadPlayerIndex
 */
//go:sdl3extern
var SetGamepadPlayerIndex func(gamepad Gamepad, player_index int32) bool

/**
 * Get the USB vendor ID of an opened gamepad, if available.
 *
 * If the vendor ID isn't available this function returns 0.
 *
 * \param gamepad the gamepad object to query.
 * \returns the USB vendor ID, or zero if unavailable.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadVendorForID
 */
//go:sdl3extern
var GetGamepadVendor func(gamepad Gamepad) uint16

/**
 * Get the USB product ID of an opened gamepad, if available.
 *
 * If the product ID isn't available this function returns 0.
 *
 * \param gamepad the gamepad object to query.
 * \returns the USB product ID, or zero if unavailable.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadProductForID
 */
//go:sdl3extern
var GetGamepadProduct func(gamepad Gamepad) uint16

/**
 * Get the product version of an opened gamepad, if available.
 *
 * If the product version isn't available this function returns 0.
 *
 * \param gamepad the gamepad object to query.
 * \returns the USB product version, or zero if unavailable.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadProductVersionForID
 */
//go:sdl3extern
var GetGamepadProductVersion func(gamepad Gamepad) uint16

/**
 * Get the firmware version of an opened gamepad, if available.
 *
 * If the firmware version isn't available this function returns 0.
 *
 * \param gamepad the gamepad object to query.
 * \returns the gamepad firmware version, or zero if unavailable.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadFirmwareVersion func(gamepad Gamepad) uint16

/**
 * Get the serial number of an opened gamepad, if available.
 *
 * Returns the serial number of the gamepad, or NULL if it is not available.
 *
 * \param gamepad the gamepad object to query.
 * \returns the serial number, or NULL if unavailable.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadSerial func(gamepad Gamepad) string

/**
 * Get the Steam Input handle of an opened gamepad, if available.
 *
 * Returns an InputHandle_t for the gamepad that can be used with Steam Input
 * API: https://partner.steamgames.com/doc/api/ISteamInput
 *
 * \param gamepad the gamepad object to query.
 * \returns the gamepad handle, or 0 if unavailable.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadSteamHandle func(gamepad Gamepad) uint64

/**
 * Get the connection state of a gamepad.
 *
 * \param gamepad the gamepad object to query.
 * \returns the connection state on success or
 *          `SDL_JOYSTICK_CONNECTION_INVALID` on failure; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadConnectionState func(gamepad Gamepad) JoystickConnectionState

/**
 * Get the battery state of a gamepad.
 *
 * You should never take a battery status as absolute truth. Batteries
 * (especially failing batteries) are delicate hardware, and the values
 * reported here are best estimates based on what that hardware reports. It's
 * not uncommon for older batteries to lose stored power much faster than it
 * reports, or completely drain when reporting it has 20 percent left, etc.
 *
 * \param gamepad the gamepad object to query.
 * \param percent a pointer filled in with the percentage of battery life
 *                left, between 0 and 100, or NULL to ignore. This will be
 *                filled in with -1 we can't determine a value or there is no
 *                battery.
 * \returns the current battery state.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadPowerInfo func(gamepad Gamepad, percent *int32) PowerState

/**
 * Check if a gamepad has been opened and is currently connected.
 *
 * \param gamepad a gamepad identifier previously returned by
 *                SDL_OpenGamepad().
 * \returns true if the gamepad has been opened and is currently connected, or
 *          false if not.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GamepadConnected func(gamepad Gamepad) bool

/**
 * Get the underlying joystick from a gamepad.
 *
 * This function will give you a SDL_Joystick object, which allows you to use
 * the SDL_Joystick functions with a SDL_Gamepad object. This would be useful
 * for getting a joystick's position at any given time, even if it hasn't
 * moved (moving it would produce an event, which would have the axis' value).
 *
 * The pointer returned is owned by the SDL_Gamepad. You should not call
 * SDL_CloseJoystick() on it, for example, since doing so will likely cause
 * SDL to crash.
 *
 * \param gamepad the gamepad object that you want to get a joystick from.
 * \returns an SDL_Joystick object, or NULL on failure; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadJoystick func(gamepad Gamepad) Joystick

/**
 * Set the state of gamepad event processing.
 *
 * If gamepad events are disabled, you must call SDL_UpdateGamepads() yourself
 * and check the state of the gamepad when you want gamepad information.
 *
 * \param enabled whether to process gamepad events or not.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GamepadEventsEnabled
 * \sa SDL_UpdateGamepads
 */
//go:sdl3extern
var SetGamepadEventsEnabled func(enabled bool)

/**
 * Query the state of gamepad event processing.
 *
 * If gamepad events are disabled, you must call SDL_UpdateGamepads() yourself
 * and check the state of the gamepad when you want gamepad information.
 *
 * \returns true if gamepad events are being processed, false otherwise.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SetGamepadEventsEnabled
 */
//go:sdl3extern
var GamepadEventsEnabled func() bool

/**
 * Get the SDL joystick layer bindings for a gamepad.
 *
 * \param gamepad a gamepad.
 * \param count a pointer filled in with the number of bindings returned.
 * \returns a NULL terminated array of pointers to bindings or NULL on
 *          failure; call SDL_GetError() for more information. This is a
 *          single allocation that should be freed with SDL_free() when it is
 *          no longer needed.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadBindings func(gamepad Gamepad, count *int32) []*GamepadBinding

/**
 * Manually pump gamepad updates if not using the loop.
 *
 * This function is called automatically by the event loop if events are
 * enabled. Under such circumstances, it will not be necessary to call this
 * function.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var UpdateGamepads func()

/**
 * Convert a string into SDL_GamepadType enum.
 *
 * This function is called internally to translate SDL_Gamepad mapping strings
 * for the underlying joystick device into the consistent SDL_Gamepad mapping.
 * You do not normally need to call this function unless you are parsing
 * SDL_Gamepad mappings in your own code.
 *
 * \param str string representing a type_ GamepadType.
 * \returns the SDL_GamepadType enum corresponding to the input string, or
 *          `SDL_GAMEPAD_TYPE_UNKNOWN` if no match was found.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadStringForType
 */
//go:sdl3extern
var GetGamepadTypeFromString func(str string) GamepadType

/**
 * Convert from an SDL_GamepadType enum to a string.
 *
 * \param type an enum value for a given SDL_GamepadType.
 * \returns a string for the given type, or NULL if an invalid type is
 *          specified. The string returned is of the format used by
 *          SDL_Gamepad mapping strings.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadTypeFromString
 */
//go:sdl3extern
var GetGamepadStringForType func(type_ GamepadType) string

/**
 * Convert a string into SDL_GamepadAxis enum.
 *
 * This function is called internally to translate SDL_Gamepad mapping strings
 * for the underlying joystick device into the consistent SDL_Gamepad mapping.
 * You do not normally need to call this function unless you are parsing
 * SDL_Gamepad mappings in your own code.
 *
 * Note specially that "righttrigger" and "lefttrigger" map to
 * `SDL_GAMEPAD_AXIS_RIGHT_TRIGGER` and `SDL_GAMEPAD_AXIS_LEFT_TRIGGER`,
 * respectively.
 *
 * \param str string representing a SDL_Gamepad axis.
 * \returns the SDL_GamepadAxis enum corresponding to the input string, or
 *          `SDL_GAMEPAD_AXIS_INVALID` if no match was found.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadStringForAxis
 */
//go:sdl3extern
var GetGamepadAxisFromString func(str string) GamepadAxis

/**
 * Convert from an SDL_GamepadAxis enum to a string.
 *
 * \param axis an enum value for a given SDL_GamepadAxis.
 * \returns a string for the given axis, or NULL if an invalid axis is
 *          specified. The string returned is of the format used by
 *          SDL_Gamepad mapping strings.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadAxisFromString
 */
//go:sdl3extern
var GetGamepadStringForAxis func(axis GamepadAxis) string

/**
 * Query whether a gamepad has a given axis.
 *
 * This merely reports whether the gamepad's mapping defined this axis, as
 * that is all the information SDL has about the physical device.
 *
 * \param gamepad a gamepad.
 * \param axis an axis enum value (an SDL_GamepadAxis value).
 * \returns true if the gamepad has this axis, false otherwise.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GamepadHasButton
 * \sa SDL_GetGamepadAxis
 */
//go:sdl3extern
var GamepadHasAxis func(gamepad Gamepad, axis GamepadAxis) bool

/**
 * Get the current state of an axis control on a gamepad.
 *
 * The axis indices start at index 0.
 *
 * For thumbsticks, the state is a value ranging from -32768 (up/left) to
 * 32767 (down/right).
 *
 * Triggers range from 0 when released to 32767 when fully pressed, and never
 * return a negative value. Note that this differs from the value reported by
 * the lower-level SDL_GetJoystickAxis(), which normally uses the full range.
 *
 * Note that for invalid gamepads or axes, this will return 0. Zero is also a
 * valid value in normal operation; usually it means a centered axis.
 *
 * \param gamepad a gamepad.
 * \param axis an axis index (one of the SDL_GamepadAxis values).
 * \returns axis state.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GamepadHasAxis
 * \sa SDL_GetGamepadButton
 */
//go:sdl3extern
var GetGamepadAxis func(gamepad Gamepad, axis GamepadAxis) int16

/**
 * Convert a string into an SDL_GamepadButton enum.
 *
 * This function is called internally to translate SDL_Gamepad mapping strings
 * for the underlying joystick device into the consistent SDL_Gamepad mapping.
 * You do not normally need to call this function unless you are parsing
 * SDL_Gamepad mappings in your own code.
 *
 * \param str string representing a SDL_Gamepad axis.
 * \returns the SDL_GamepadButton enum corresponding to the input string, or
 *          `SDL_GAMEPAD_BUTTON_INVALID` if no match was found.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadStringForButton
 */
//go:sdl3extern
var GetGamepadButtonFromString func(str string) GamepadButton

/**
 * Convert from an SDL_GamepadButton enum to a string.
 *
 * \param button an enum value for a given SDL_GamepadButton.
 * \returns a string for the given button, or NULL if an invalid button is
 *          specified. The string returned is of the format used by
 *          SDL_Gamepad mapping strings.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadButtonFromString
 */
//go:sdl3extern
var GetGamepadStringForButton func(button GamepadButton) string

/**
 * Query whether a gamepad has a given button.
 *
 * This merely reports whether the gamepad's mapping defined this button, as
 * that is all the information SDL has about the physical device.
 *
 * \param gamepad a gamepad.
 * \param button a button enum value (an SDL_GamepadButton value).
 * \returns true if the gamepad has this button, false otherwise.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GamepadHasAxis
 */
//go:sdl3extern
var GamepadHasButton func(gamepad Gamepad, button GamepadButton) bool

/**
 * Get the current state of a button on a gamepad.
 *
 * \param gamepad a gamepad.
 * \param button a button index (one of the SDL_GamepadButton values).
 * \returns true if the button is pressed, false otherwise.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GamepadHasButton
 * \sa SDL_GetGamepadAxis
 */
//go:sdl3extern
var GetGamepadButton func(gamepad Gamepad, button GamepadButton) bool

/**
 * Get the label of a button on a gamepad.
 *
 * \param type the type of gamepad to check.
 * \param button a button index (one of the SDL_GamepadButton values).
 * \returns the SDL_GamepadButtonLabel enum corresponding to the button label.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadButtonLabel
 */
//go:sdl3extern
var GetGamepadButtonLabelForType func(type_ GamepadType, button GamepadButton) GamepadButtonLabel

/**
 * Get the label of a button on a gamepad.
 *
 * \param gamepad a gamepad.
 * \param button a button index (one of the SDL_GamepadButton values).
 * \returns the SDL_GamepadButtonLabel enum corresponding to the button label.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadButtonLabelForType
 */
//go:sdl3extern
var GetGamepadButtonLabel func(gamepad Gamepad, button GamepadButton) GamepadButtonLabel

/**
 * Get the number of touchpads on a gamepad.
 *
 * \param gamepad a gamepad.
 * \returns number of touchpads.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetNumGamepadTouchpadFingers
 */
//go:sdl3extern
var GetNumGamepadTouchpads func(gamepad Gamepad) int32

/**
 * Get the number of supported simultaneous fingers on a touchpad on a game
 * gamepad.
 *
 * \param gamepad a gamepad.
 * \param touchpad a touchpad.
 * \returns number of supported simultaneous fingers.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadTouchpadFinger
 * \sa SDL_GetNumGamepadTouchpads
 */
//go:sdl3extern
var GetNumGamepadTouchpadFingers func(gamepad Gamepad, touchpad int32) int32

/**
 * Get the current state of a finger on a touchpad on a gamepad.
 *
 * \param gamepad a gamepad.
 * \param touchpad a touchpad.
 * \param finger a finger.
 * \param down a pointer filled with true if the finger is down, false
 *             otherwise, may be NULL.
 * \param x a pointer filled with the x position, normalized 0 to 1, with the
 *          origin in the upper left, may be NULL.
 * \param y a pointer filled with the y position, normalized 0 to 1, with the
 *          origin in the upper left, may be NULL.
 * \param pressure a pointer filled with pressure value, may be NULL.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetNumGamepadTouchpadFingers
 */
//go:sdl3extern
var GetGamepadTouchpadFinger func(gamepad Gamepad, touchpad, finger int32, down *bool, x, y, pressure *float32) bool

/**
 * Return whether a gamepad has a particular sensor.
 *
 * \param gamepad the gamepad to query.
 * \param type the type of sensor to query.
 * \returns true if the sensor exists, false otherwise.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadSensorData
 * \sa SDL_GetGamepadSensorDataRate
 * \sa SDL_SetGamepadSensorEnabled
 */
//go:sdl3extern
var GamepadHasSensor func(gamepad Gamepad, type_ SensorType) bool

/**
 * Set whether data reporting for a gamepad sensor is enabled.
 *
 * \param gamepad the gamepad to update.
 * \param type the type of sensor to enable/disable.
 * \param enabled whether data reporting should be enabled.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GamepadHasSensor
 * \sa SDL_GamepadSensorEnabled
 */
//go:sdl3extern
var SetGamepadSensorEnabled func(gamepad Gamepad, type_ SensorType, enabled bool) bool

/**
 * Query whether sensor data reporting is enabled for a gamepad.
 *
 * \param gamepad the gamepad to query.
 * \param type the type of sensor to query.
 * \returns true if the sensor is enabled, false otherwise.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SetGamepadSensorEnabled
 */
//go:sdl3extern
var GamepadSensorEnabled func(gamepad Gamepad, type_ SensorType) bool

/**
 * Get the data rate (number of events per second) of a gamepad sensor.
 *
 * \param gamepad the gamepad to query.
 * \param type the type of sensor to query.
 * \returns the data rate, or 0.0f if the data rate is not available.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadSensorDataRate func(gamepad Gamepad, type_ SensorType) float32

/**
 * Get the current state of a gamepad sensor.
 *
 * The number of values and interpretation of the data is sensor dependent.
 * See SDL_sensor.h for the details for each type of sensor.
 *
 * \param gamepad the gamepad to query.
 * \param type the type of sensor to query.
 * \param data a pointer filled with the current sensor state.
 * \param num_values the number of values to write to data.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGamepadSensorData func(gamepad Gamepad, type_ SensorType, data *float32, num_values int32) bool

/**
 * Start a rumble effect on a gamepad.
 *
 * Each call to this function cancels any previous rumble effect, and calling
 * it with 0 intensity stops any rumbling.
 *
 * This function requires you to process SDL events or call
 * SDL_UpdateJoysticks() to update rumble state.
 *
 * \param gamepad the gamepad to vibrate.
 * \param low_frequency_rumble the intensity of the low frequency (left)
 *                             rumble motor, from 0 to 0xFFFF.
 * \param high_frequency_rumble the intensity of the high frequency (right)
 *                              rumble motor, from 0 to 0xFFFF.
 * \param duration_ms the duration of the rumble effect, in milliseconds.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var RumbleGamepad func(gamepad Gamepad, low_frequency_rumble, high_frequency_rumble uint16, duration_ms uint32) bool

/**
 * Start a rumble effect in the gamepad's triggers.
 *
 * Each call to this function cancels any previous trigger rumble effect, and
 * calling it with 0 intensity stops any rumbling.
 *
 * Note that this is rumbling of the _triggers_ and not the gamepad as a
 * whole. This is currently only supported on Xbox One gamepads. If you want
 * the (more common) whole-gamepad rumble, use SDL_RumbleGamepad() instead.
 *
 * This function requires you to process SDL events or call
 * SDL_UpdateJoysticks() to update rumble state.
 *
 * \param gamepad the gamepad to vibrate.
 * \param left_rumble the intensity of the left trigger rumble motor, from 0
 *                    to 0xFFFF.
 * \param right_rumble the intensity of the right trigger rumble motor, from 0
 *                     to 0xFFFF.
 * \param duration_ms the duration of the rumble effect, in milliseconds.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_RumbleGamepad
 */
//go:sdl3extern
var RumbleGamepadTriggers func(gamepad Gamepad, left_rumble, right_rumble uint16, duration_ms uint32) bool

/**
 * Update a gamepad's LED color.
 *
 * An example of a joystick LED is the light on the back of a PlayStation 4's
 * DualShock 4 controller.
 *
 * For gamepads with a single color LED, the maximum of the RGB values will be
 * used as the LED brightness.
 *
 * \param gamepad the gamepad to update.
 * \param red the intensity of the red LED.
 * \param green the intensity of the green LED.
 * \param blue the intensity of the blue LED.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetGamepadLED func(gamepad Gamepad, red, green, blue uint8) bool

/**
 * Send a gamepad specific effect packet.
 *
 * \param gamepad the gamepad to affect.
 * \param data the data to send to the gamepad.
 * \param size the size of the data to send to the gamepad.
 * \returns true on success or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SendGamepadEffect func(gamepad Gamepad, data []byte, size int32) bool

/**
 * Close a gamepad previously opened with SDL_OpenGamepad().
 *
 * \param gamepad a gamepad identifier previously returned by
 *                SDL_OpenGamepad().
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_OpenGamepad
 */
//go:sdl3extern
var CloseGamepad func(gamepad Gamepad)

/**
 * Return the sfSymbolsName for a given button on a gamepad on Apple
 * platforms.
 *
 * \param gamepad the gamepad to query.
 * \param button a button on the gamepad.
 * \returns the sfSymbolsName or NULL if the name can't be found.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadAppleSFSymbolsNameForAxis
 */
//go:sdl3extern
var GetGamepadAppleSFSymbolsNameForButton func(gamepad Gamepad, button GamepadButton) string

/**
 * Return the sfSymbolsName for a given axis on a gamepad on Apple platforms.
 *
 * \param gamepad the gamepad to query.
 * \param axis an axis on the gamepad.
 * \returns the sfSymbolsName or NULL if the name can't be found.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGamepadAppleSFSymbolsNameForButton
 */
//go:sdl3extern
var GetGamepadAppleSFSymbolsNameForAxis func(gamepad Gamepad, axis GamepadAxis) string
