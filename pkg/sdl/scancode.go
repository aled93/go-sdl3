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
 * # CategoryScancode
 *
 * Defines keyboard scancodes.
 *
 * Please refer to the Best Keyboard Practices document for details on what
 * this information means and how best to use it.
 *
 * https://wiki.libsdl.org/SDL3/BestKeyboardPractices
 */
package sdl

/**
 * The SDL keyboard scancode representation.
 *
 * An SDL scancode is the physical representation of a key on the keyboard,
 * independent of language and keyboard mapping.
 *
 * Values of this type are used to represent keyboard keys, among other places
 * in the `scancode` field of the SDL_KeyboardEvent structure.
 *
 * The values in this enumeration are based on the USB usage page standard:
 * https://usb.org/sites/default/files/hut1_5.pdf
 *
 * \since This enum is available since SDL 3.2.0.
 */
type Scancode int32

const (
	SC_Unknown Scancode = 0

	/**
	 *  \name Usage page 0x07
	 *
	 *  These values are from usage page 0x07 (USB keyboard page).
	 */
	/* @{ */

	SC_A Scancode = 4
	SC_B Scancode = 5
	SC_C Scancode = 6
	SC_D Scancode = 7
	SC_E Scancode = 8
	SC_F Scancode = 9
	SC_G Scancode = 10
	SC_H Scancode = 11
	SC_I Scancode = 12
	SC_J Scancode = 13
	SC_K Scancode = 14
	SC_L Scancode = 15
	SC_M Scancode = 16
	SC_N Scancode = 17
	SC_O Scancode = 18
	SC_P Scancode = 19
	SC_Q Scancode = 20
	SC_R Scancode = 21
	SC_S Scancode = 22
	SC_T Scancode = 23
	SC_U Scancode = 24
	SC_V Scancode = 25
	SC_W Scancode = 26
	SC_X Scancode = 27
	SC_Y Scancode = 28
	SC_Z Scancode = 29

	SC_1 Scancode = 30
	SC_2 Scancode = 31
	SC_3 Scancode = 32
	SC_4 Scancode = 33
	SC_5 Scancode = 34
	SC_6 Scancode = 35
	SC_7 Scancode = 36
	SC_8 Scancode = 37
	SC_9 Scancode = 38
	SC_0 Scancode = 39

	SC_Return    Scancode = 40
	SC_Escape    Scancode = 41
	SC_Backspace Scancode = 42
	SC_Tab       Scancode = 43
	SC_Space     Scancode = 44

	SC_Minus        Scancode = 45
	SC_Equals       Scancode = 46
	SC_Leftbracket  Scancode = 47
	SC_Rightbracket Scancode = 48
	SC_Backslash    Scancode = 49 /**< Located at the lower left of the return
	 *   key on ISO keyboards and at the right end
	 *   of the QWERTY row on ANSI keyboards.
	 *   Produces REVERSE SOLIDUS (backslash) and
	 *   VERTICAL LINE in a US layout, REVERSE
	 *   SOLIDUS and VERTICAL LINE in a UK Mac
	 *   layout, NUMBER SIGN and TILDE in a UK
	 *   Windows layout, DOLLAR SIGN and POUND SIGN
	 *   in a Swiss German layout, NUMBER SIGN and
	 *   APOSTROPHE in a German layout, GRAVE
	 *   ACCENT and POUND SIGN in a French Mac
	 *   layout, and ASTERISK and MICRO SIGN in a
	 *   French Windows layout.
	 */
	SC_NonUsHash Scancode = 50 /**< ISO USB keyboards actually use this code
	 *   instead of 49 for the same key, but all
	 *   OSes I've seen treat the two codes
	 *   identically. So, as an implementor, unless
	 *   your keyboard generates both of those
	 *   codes and your OS treats them differently,
	 *   you should generate SDL_SCANCODE_BACKSLASH
	 *   instead of this code. As a user, you
	 *   should not rely on this code because SDL
	 *   will never generate it with most (all?)
	 *   keyboards.
	 */
	SC_Semicolon  Scancode = 51
	SC_Apostrophe Scancode = 52
	SC_Grave      Scancode = 53 /**< Located in the top left corner (on both ANSI
	 *   and ISO keyboards). Produces GRAVE ACCENT and
	 *   TILDE in a US Windows layout and in US and UK
	 *   Mac layouts on ANSI keyboards, GRAVE ACCENT
	 *   and NOT SIGN in a UK Windows layout, SECTION
	 *   SIGN and PLUS-MINUS SIGN in US and UK Mac
	 *   layouts on ISO keyboards, SECTION SIGN and
	 *   DEGREE SIGN in a Swiss German layout (Mac:
	 *   only on ISO keyboards), CIRCUMFLEX ACCENT and
	 *   DEGREE SIGN in a German layout (Mac: only on
	 *   ISO keyboards), SUPERSCRIPT TWO and TILDE in a
	 *   French Windows layout, COMMERCIAL AT and
	 *   NUMBER SIGN in a French Mac layout on ISO
	 *   keyboards, and LESS-THAN SIGN and GREATER-THAN
	 *   SIGN in a Swiss German, German, or French Mac
	 *   layout on ANSI keyboards.
	 */
	SC_Comma  Scancode = 54
	SC_Period Scancode = 55
	SC_Slash  Scancode = 56

	SC_CapsLock Scancode = 57

	SC_F1  Scancode = 58
	SC_F2  Scancode = 59
	SC_F3  Scancode = 60
	SC_F4  Scancode = 61
	SC_F5  Scancode = 62
	SC_F6  Scancode = 63
	SC_F7  Scancode = 64
	SC_F8  Scancode = 65
	SC_F9  Scancode = 66
	SC_F10 Scancode = 67
	SC_F11 Scancode = 68
	SC_F12 Scancode = 69

	SC_PrintScreen Scancode = 70
	SC_ScrollLock  Scancode = 71
	SC_Pause       Scancode = 72
	SC_Insert      Scancode = 73 /**< insert on PC, help on some Mac keyboards (but
	  does send code 73, not 117) */
	SC_Home     Scancode = 74
	SC_PageUp   Scancode = 75
	SC_Delete   Scancode = 76
	SC_End      Scancode = 77
	SC_PageDown Scancode = 78
	SC_Right    Scancode = 79
	SC_Left     Scancode = 80
	SC_Down     Scancode = 81
	SC_Up       Scancode = 82

	SC_NumLockClear Scancode = 83 /**< num lock on PC, clear on Mac keyboards
	 */
	SC_KP_Divide   Scancode = 84
	SC_KP_Multiply Scancode = 85
	SC_KP_Minus    Scancode = 86
	SC_KP_Plus     Scancode = 87
	SC_KP_Enter    Scancode = 88
	SC_KP_1        Scancode = 89
	SC_KP_2        Scancode = 90
	SC_KP_3        Scancode = 91
	SC_KP_4        Scancode = 92
	SC_KP_5        Scancode = 93
	SC_KP_6        Scancode = 94
	SC_KP_7        Scancode = 95
	SC_KP_8        Scancode = 96
	SC_KP_9        Scancode = 97
	SC_KP_0        Scancode = 98
	SC_KP_Period   Scancode = 99

	SC_NonUsBackslash Scancode = 100 /**< This is the additional key that ISO
	 *   keyboards have over ANSI ones,
	 *   located between left shift and Y.
	 *   Produces GRAVE ACCENT and TILDE in a
	 *   US or UK Mac layout, REVERSE SOLIDUS
	 *   (backslash) and VERTICAL LINE in a
	 *   US or UK Windows layout, and
	 *   LESS-THAN SIGN and GREATER-THAN SIGN
	 *   in a Swiss German, German, or French
	 *   layout. */
	SC_Application Scancode = 101 /**< windows contextual menu, compose */
	SC_Power       Scancode = 102 /**< The USB document says this is a status flag,
	 *   not a physical key - but some Mac keyboards
	 *   do have a power key. */
	SC_KP_Equals  Scancode = 103
	SC_F13        Scancode = 104
	SC_F14        Scancode = 105
	SC_F15        Scancode = 106
	SC_F16        Scancode = 107
	SC_F17        Scancode = 108
	SC_F18        Scancode = 109
	SC_F19        Scancode = 110
	SC_F20        Scancode = 111
	SC_F21        Scancode = 112
	SC_F22        Scancode = 113
	SC_F23        Scancode = 114
	SC_F24        Scancode = 115
	SC_Execute    Scancode = 116
	SC_Help       Scancode = 117 /**< AL Integrated Help Center */
	SC_Menu       Scancode = 118 /**< Menu (show menu) */
	SC_Select     Scancode = 119
	SC_Stop       Scancode = 120 /**< AC Stop */
	SC_Again      Scancode = 121 /**< AC Redo/Repeat */
	SC_Undo       Scancode = 122 /**< AC Undo */
	SC_Cut        Scancode = 123 /**< AC Cut */
	SC_Copy       Scancode = 124 /**< AC Copy */
	SC_Paste      Scancode = 125 /**< AC Paste */
	SC_Find       Scancode = 126 /**< AC Find */
	SC_Mute       Scancode = 127
	SC_VolumeUp   Scancode = 128
	SC_VolumeDown Scancode = 129
	/* not sure whether there's a reason to enable these */
	/* SCANCODE_LOCKINGCAPSLOCK Scancode = 130  */
	/* SCANCODE_LOCKINGNUMLOCK Scancode = 131 */
	/* SCANCODE_LOCKINGSCROLLLOCK Scancode = 132 */
	SC_KP_Comma       Scancode = 133
	SC_KP_EqualsAs400 Scancode = 134

	SC_International1 Scancode = 135 /**< used on Asian keyboards, see
	  footnotes in USB doc */
	SC_International2 Scancode = 136
	SC_International3 Scancode = 137 /**< Yen */
	SC_International4 Scancode = 138
	SC_International5 Scancode = 139
	SC_International6 Scancode = 140
	SC_International7 Scancode = 141
	SC_International8 Scancode = 142
	SC_International9 Scancode = 143
	SC_Lang1          Scancode = 144 /**< Hangul/English toggle */
	SC_Lang2          Scancode = 145 /**< Hanja conversion */
	SC_Lang3          Scancode = 146 /**< Katakana */
	SC_Lang4          Scancode = 147 /**< Hiragana */
	SC_Lang5          Scancode = 148 /**< Zenkaku/Hankaku */
	SC_Lang6          Scancode = 149 /**< reserved */
	SC_Lang7          Scancode = 150 /**< reserved */
	SC_Lang8          Scancode = 151 /**< reserved */
	SC_Lang9          Scancode = 152 /**< reserved */

	SC_AltErase   Scancode = 153 /**< Erase-Eaze */
	SC_SysReq     Scancode = 154
	SC_Cancel     Scancode = 155 /**< AC Cancel */
	SC_Clear      Scancode = 156
	SC_Prior      Scancode = 157
	SC_Return2    Scancode = 158
	SC_Separator  Scancode = 159
	SC_Out        Scancode = 160
	SC_Oper       Scancode = 161
	SC_ClearAgain Scancode = 162
	SC_Crsel      Scancode = 163
	SC_Exsel      Scancode = 164

	SC_KP_00              Scancode = 176
	SC_KP_000             Scancode = 177
	SC_ThousandsSeparator Scancode = 178
	SC_DecimalSeparator   Scancode = 179
	SC_CurrencyUnit       Scancode = 180
	SC_CurrencySubUnit    Scancode = 181
	SC_KP_LeftParen       Scancode = 182
	SC_KP_RightParen      Scancode = 183
	SC_KP_LeftBrace       Scancode = 184
	SC_KP_RightBrace      Scancode = 185
	SC_KP_Tab             Scancode = 186
	SC_KP_Backspace       Scancode = 187
	SC_KP_A               Scancode = 188
	SC_KP_B               Scancode = 189
	SC_KP_C               Scancode = 190
	SC_KP_D               Scancode = 191
	SC_KP_E               Scancode = 192
	SC_KP_F               Scancode = 193
	SC_KP_Xor             Scancode = 194
	SC_KP_Power           Scancode = 195
	SC_KP_Percent         Scancode = 196
	SC_KP_Less            Scancode = 197
	SC_KP_Greater         Scancode = 198
	SC_KP_Ampersand       Scancode = 199
	SC_KP_DblAmpersand    Scancode = 200
	SC_KP_VerticalBar     Scancode = 201
	SC_KP_DblVerticalBar  Scancode = 202
	SC_KP_Colon           Scancode = 203
	SC_KP_Hash            Scancode = 204
	SC_KP_Space           Scancode = 205
	SC_KP_At              Scancode = 206
	SC_KP_Exclam          Scancode = 207
	SC_KP_MemStore        Scancode = 208
	SC_KP_MemRecall       Scancode = 209
	SC_KP_MemClear        Scancode = 210
	SC_KP_MemAdd          Scancode = 211
	SC_KP_MemSubtract     Scancode = 212
	SC_KP_MemMultiply     Scancode = 213
	SC_KP_MemDivide       Scancode = 214
	SC_KP_PlusMinus       Scancode = 215
	SC_KP_Clear           Scancode = 216
	SC_KP_ClearEntry      Scancode = 217
	SC_KP_Binary          Scancode = 218
	SC_KP_Octal           Scancode = 219
	SC_KP_Decimal         Scancode = 220
	SC_KP_Hexadecimal     Scancode = 221

	SC_LCTRL  Scancode = 224
	SC_LSHIFT Scancode = 225
	SC_LALT   Scancode = 226 /**< alt, option */
	SC_LGUI   Scancode = 227 /**< windows, command (apple), meta */
	SC_RCTRL  Scancode = 228
	SC_RSHIFT Scancode = 229
	SC_RALT   Scancode = 230 /**< alt gr, option */
	SC_RGUI   Scancode = 231 /**< windows, command (apple), meta */

	SC_MODE Scancode = 257 /**< I'm not sure if this is really not covered
	 *   by any of the above, but since there's a
	 *   special SDL_KMOD_MODE for it I'm adding it here
	 */

	/* @} */ /* Usage page 0x07 */

	/**
	 *  \name Usage page 0x0C
	 *
	 *  These values are mapped from usage page 0x0C (USB consumer page).
	 *
	 *  There are way more keys in the spec than we can represent in the
	 *  current scancode range, so pick the ones that commonly come up in
	 *  real world usage.
	 */
	/* @{ */

	SC_SLEEP Scancode = 258 /**< Sleep */
	SC_WAKE  Scancode = 259 /**< Wake */

	SC_CHANNEL_INCREMENT Scancode = 260 /**< Channel Increment */
	SC_CHANNEL_DECREMENT Scancode = 261 /**< Channel Decrement */

	SC_MEDIA_PLAY           Scancode = 262 /**< Play */
	SC_MEDIA_PAUSE          Scancode = 263 /**< Pause */
	SC_MEDIA_RECORD         Scancode = 264 /**< Record */
	SC_MEDIA_FAST_FORWARD   Scancode = 265 /**< Fast Forward */
	SC_MEDIA_REWIND         Scancode = 266 /**< Rewind */
	SC_MEDIA_NEXT_TRACK     Scancode = 267 /**< Next Track */
	SC_MEDIA_PREVIOUS_TRACK Scancode = 268 /**< Previous Track */
	SC_MEDIA_STOP           Scancode = 269 /**< Stop */
	SC_MEDIA_EJECT          Scancode = 270 /**< Eject */
	SC_MEDIA_PLAY_PAUSE     Scancode = 271 /**< Play / Pause */
	SC_MEDIA_SELECT         Scancode = 272 /* Media Select */

	SC_AC_NEW        Scancode = 273 /**< AC New */
	SC_AC_OPEN       Scancode = 274 /**< AC Open */
	SC_AC_CLOSE      Scancode = 275 /**< AC Close */
	SC_AC_EXIT       Scancode = 276 /**< AC Exit */
	SC_AC_SAVE       Scancode = 277 /**< AC Save */
	SC_AC_PRINT      Scancode = 278 /**< AC Print */
	SC_AC_PROPERTIES Scancode = 279 /**< AC Properties */

	SC_AC_SEARCH    Scancode = 280 /**< AC Search */
	SC_AC_HOME      Scancode = 281 /**< AC Home */
	SC_AC_BACK      Scancode = 282 /**< AC Back */
	SC_AC_FORWARD   Scancode = 283 /**< AC Forward */
	SC_AC_STOP      Scancode = 284 /**< AC Stop */
	SC_AC_REFRESH   Scancode = 285 /**< AC Refresh */
	SC_AC_BOOKMARKS Scancode = 286 /**< AC Bookmarks */

	/* @} */ /* Usage page 0x0C */

	/**
	 *  \name Mobile keys
	 *
	 *  These are values that are often used on mobile phones.
	 */
	/* @{ */

	SC_SOFTLEFT Scancode = 287 /**< Usually situated below the display on phones and
	  used as a multi-function feature key for selecting
	  a software defined function shown on the bottom left
	  of the display. */
	SC_SOFTRIGHT Scancode = 288 /**< Usually situated below the display on phones and
	  used as a multi-function feature key for selecting
	  a software defined function shown on the bottom right
	  of the display. */
	SC_CALL    Scancode = 289 /**< Used for accepting phone calls. */
	SC_ENDCALL Scancode = 290 /**< Used for rejecting phone calls. */

	/* @} */ /* Mobile keys */

	/* Add any other keys here. */

	SC_RESERVED Scancode = 400 /**< 400-500 reserved for dynamic keycodes */

	SC_COUNT Scancode = 512 /**< not a key, just marks the number of scancodes for array bounds */
)
