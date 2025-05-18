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
	ScancodeUnknown Scancode = 0

	/**
	 *  \name Usage page 0x07
	 *
	 *  These values are from usage page 0x07 (USB keyboard page).
	 */
	/* @{ */

	ScancodeA Scancode = 4
	ScancodeB Scancode = 5
	ScancodeC Scancode = 6
	ScancodeD Scancode = 7
	ScancodeE Scancode = 8
	ScancodeF Scancode = 9
	ScancodeG Scancode = 10
	ScancodeH Scancode = 11
	ScancodeI Scancode = 12
	ScancodeJ Scancode = 13
	ScancodeK Scancode = 14
	ScancodeL Scancode = 15
	ScancodeM Scancode = 16
	ScancodeN Scancode = 17
	ScancodeO Scancode = 18
	ScancodeP Scancode = 19
	ScancodeQ Scancode = 20
	ScancodeR Scancode = 21
	ScancodeS Scancode = 22
	ScancodeT Scancode = 23
	ScancodeU Scancode = 24
	ScancodeV Scancode = 25
	ScancodeW Scancode = 26
	ScancodeX Scancode = 27
	ScancodeY Scancode = 28
	ScancodeZ Scancode = 29

	Scancode1 Scancode = 30
	Scancode2 Scancode = 31
	Scancode3 Scancode = 32
	Scancode4 Scancode = 33
	Scancode5 Scancode = 34
	Scancode6 Scancode = 35
	Scancode7 Scancode = 36
	Scancode8 Scancode = 37
	Scancode9 Scancode = 38
	Scancode0 Scancode = 39

	ScancodeReturn    Scancode = 40
	ScancodeEscape    Scancode = 41
	ScancodeBackspace Scancode = 42
	ScancodeTab       Scancode = 43
	ScancodeSpace     Scancode = 44

	ScancodeMinus        Scancode = 45
	ScancodeEquals       Scancode = 46
	ScancodeLeftbracket  Scancode = 47
	ScancodeRightbracket Scancode = 48
	ScancodeBackslash    Scancode = 49 /**< Located at the lower left of the return
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
	ScancodeNonUsHash Scancode = 50 /**< ISO USB keyboards actually use this code
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
	ScancodeSemicolon  Scancode = 51
	ScancodeApostrophe Scancode = 52
	ScancodeGrave      Scancode = 53 /**< Located in the top left corner (on both ANSI
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
	ScancodeComma  Scancode = 54
	ScancodePeriod Scancode = 55
	ScancodeSlash  Scancode = 56

	ScancodeCapsLock Scancode = 57

	ScancodeF1  Scancode = 58
	ScancodeF2  Scancode = 59
	ScancodeF3  Scancode = 60
	ScancodeF4  Scancode = 61
	ScancodeF5  Scancode = 62
	ScancodeF6  Scancode = 63
	ScancodeF7  Scancode = 64
	ScancodeF8  Scancode = 65
	ScancodeF9  Scancode = 66
	ScancodeF10 Scancode = 67
	ScancodeF11 Scancode = 68
	ScancodeF12 Scancode = 69

	ScancodePrintScreen Scancode = 70
	ScancodeScrollLock  Scancode = 71
	ScancodePause       Scancode = 72
	ScancodeInsert      Scancode = 73 /**< insert on PC, help on some Mac keyboards (but
	  does send code 73, not 117) */
	ScancodeHome     Scancode = 74
	ScancodePageUp   Scancode = 75
	ScancodeDelete   Scancode = 76
	ScancodeEnd      Scancode = 77
	ScancodePageDown Scancode = 78
	ScancodeRight    Scancode = 79
	ScancodeLeft     Scancode = 80
	ScancodeDown     Scancode = 81
	ScancodeUp       Scancode = 82

	ScancodeNumLockClear Scancode = 83 /**< num lock on PC, clear on Mac keyboards
	 */
	ScancodeKPDivide   Scancode = 84
	ScancodeKPMultiply Scancode = 85
	ScancodeKPMinus    Scancode = 86
	ScancodeKPPlus     Scancode = 87
	ScancodeKPEnter    Scancode = 88
	ScancodeKP1        Scancode = 89
	ScancodeKP2        Scancode = 90
	ScancodeKP3        Scancode = 91
	ScancodeKP4        Scancode = 92
	ScancodeKP5        Scancode = 93
	ScancodeKP6        Scancode = 94
	ScancodeKP7        Scancode = 95
	ScancodeKP8        Scancode = 96
	ScancodeKP9        Scancode = 97
	ScancodeKP0        Scancode = 98
	ScancodeKPPeriod   Scancode = 99

	ScancodeNonUsBackslash Scancode = 100 /**< This is the additional key that ISO
	 *   keyboards have over ANSI ones,
	 *   located between left shift and Y.
	 *   Produces GRAVE ACCENT and TILDE in a
	 *   US or UK Mac layout, REVERSE SOLIDUS
	 *   (backslash) and VERTICAL LINE in a
	 *   US or UK Windows layout, and
	 *   LESS-THAN SIGN and GREATER-THAN SIGN
	 *   in a Swiss German, German, or French
	 *   layout. */
	ScancodeApplication Scancode = 101 /**< windows contextual menu, compose */
	ScancodePower       Scancode = 102 /**< The USB document says this is a status flag,
	 *   not a physical key - but some Mac keyboards
	 *   do have a power key. */
	ScancodeKP_Equals  Scancode = 103
	ScancodeF13        Scancode = 104
	ScancodeF14        Scancode = 105
	ScancodeF15        Scancode = 106
	ScancodeF16        Scancode = 107
	ScancodeF17        Scancode = 108
	ScancodeF18        Scancode = 109
	ScancodeF19        Scancode = 110
	ScancodeF20        Scancode = 111
	ScancodeF21        Scancode = 112
	ScancodeF22        Scancode = 113
	ScancodeF23        Scancode = 114
	ScancodeF24        Scancode = 115
	ScancodeExecute    Scancode = 116
	ScancodeHelp       Scancode = 117 /**< AL Integrated Help Center */
	ScancodeMenu       Scancode = 118 /**< Menu (show menu) */
	ScancodeSelect     Scancode = 119
	ScancodeStop       Scancode = 120 /**< AC Stop */
	ScancodeAgain      Scancode = 121 /**< AC Redo/Repeat */
	ScancodeUndo       Scancode = 122 /**< AC Undo */
	ScancodeCut        Scancode = 123 /**< AC Cut */
	ScancodeCopy       Scancode = 124 /**< AC Copy */
	ScancodePaste      Scancode = 125 /**< AC Paste */
	ScancodeFind       Scancode = 126 /**< AC Find */
	ScancodeMute       Scancode = 127
	ScancodeVolumeUp   Scancode = 128
	ScancodeVolumeDown Scancode = 129
	/* not sure whether there's a reason to enable these */
	/* SCANCODE_LOCKINGCAPSLOCK Scancode = 130  */
	/* SCANCODE_LOCKINGNUMLOCK Scancode = 131 */
	/* SCANCODE_LOCKINGSCROLLLOCK Scancode = 132 */
	ScancodeKP_Comma       Scancode = 133
	ScancodeKP_EqualsAs400 Scancode = 134

	ScancodeInternational1 Scancode = 135 /**< used on Asian keyboards, see
	  footnotes in USB doc */
	ScancodeInternational2 Scancode = 136
	ScancodeInternational3 Scancode = 137 /**< Yen */
	ScancodeInternational4 Scancode = 138
	ScancodeInternational5 Scancode = 139
	ScancodeInternational6 Scancode = 140
	ScancodeInternational7 Scancode = 141
	ScancodeInternational8 Scancode = 142
	ScancodeInternational9 Scancode = 143
	ScancodeLang1          Scancode = 144 /**< Hangul/English toggle */
	ScancodeLang2          Scancode = 145 /**< Hanja conversion */
	ScancodeLang3          Scancode = 146 /**< Katakana */
	ScancodeLang4          Scancode = 147 /**< Hiragana */
	ScancodeLang5          Scancode = 148 /**< Zenkaku/Hankaku */
	ScancodeLang6          Scancode = 149 /**< reserved */
	ScancodeLang7          Scancode = 150 /**< reserved */
	ScancodeLang8          Scancode = 151 /**< reserved */
	ScancodeLang9          Scancode = 152 /**< reserved */

	ScancodeAltErase   Scancode = 153 /**< Erase-Eaze */
	ScancodeSysReq     Scancode = 154
	ScancodeCancel     Scancode = 155 /**< AC Cancel */
	ScancodeClear      Scancode = 156
	ScancodePrior      Scancode = 157
	ScancodeReturn2    Scancode = 158
	ScancodeSeparator  Scancode = 159
	ScancodeOut        Scancode = 160
	ScancodeOper       Scancode = 161
	ScancodeClearAgain Scancode = 162
	ScancodeCrsel      Scancode = 163
	ScancodeExsel      Scancode = 164

	ScancodeKP_00              Scancode = 176
	ScancodeKP_000             Scancode = 177
	ScancodeThousandsSeparator Scancode = 178
	ScancodeDecimalSeparator   Scancode = 179
	ScancodeCurrencyUnit       Scancode = 180
	ScancodeCurrencySubUnit    Scancode = 181
	ScancodeKPLeftParen        Scancode = 182
	ScancodeKPRightParen       Scancode = 183
	ScancodeKPLeftBrace        Scancode = 184
	ScancodeKPRightBrace       Scancode = 185
	ScancodeKPTab              Scancode = 186
	ScancodeKPBackspace        Scancode = 187
	ScancodeKPA                Scancode = 188
	ScancodeKPB                Scancode = 189
	ScancodeKPC                Scancode = 190
	ScancodeKPD                Scancode = 191
	ScancodeKPE                Scancode = 192
	ScancodeKPF                Scancode = 193
	ScancodeKPXor              Scancode = 194
	ScancodeKPPower            Scancode = 195
	ScancodeKPPercent          Scancode = 196
	ScancodeKPLess             Scancode = 197
	ScancodeKPGreater          Scancode = 198
	ScancodeKPAmpersand        Scancode = 199
	ScancodeKPDblAmpersand     Scancode = 200
	ScancodeKPVerticalBar      Scancode = 201
	ScancodeKPDblVerticalBar   Scancode = 202
	ScancodeKPColon            Scancode = 203
	ScancodeKPHash             Scancode = 204
	ScancodeKPSpace            Scancode = 205
	ScancodeKPAt               Scancode = 206
	ScancodeKPExclam           Scancode = 207
	ScancodeKPMemStore         Scancode = 208
	ScancodeKPMemRecall        Scancode = 209
	ScancodeKPMemClear         Scancode = 210
	ScancodeKPMemAdd           Scancode = 211
	ScancodeKPMemSubtract      Scancode = 212
	ScancodeKPMemMultiply      Scancode = 213
	ScancodeKPMemDivide        Scancode = 214
	ScancodeKPPlusMinus        Scancode = 215
	ScancodeKPClear            Scancode = 216
	ScancodeKPClearEntry       Scancode = 217
	ScancodeKPBinary           Scancode = 218
	ScancodeKPOctal            Scancode = 219
	ScancodeKPDecimal          Scancode = 220
	ScancodeKPHexadecimal      Scancode = 221

	ScancodeLCtrl  Scancode = 224
	ScancodeLShift Scancode = 225
	ScancodeLAlt   Scancode = 226 /**< alt, option */
	ScancodeLGui   Scancode = 227 /**< windows, command (apple), meta */
	ScancodeRCtrl  Scancode = 228
	ScancodeRShift Scancode = 229
	ScancodeRAlt   Scancode = 230 /**< alt gr, option */
	ScancodeRGui   Scancode = 231 /**< windows, command (apple), meta */

	ScancodeMode Scancode = 257 /**< I'm not sure if this is really not covered
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

	ScancodeSleep Scancode = 258 /**< Sleep */
	ScancodeWake  Scancode = 259 /**< Wake */

	ScancodeChannelIncrement Scancode = 260 /**< Channel Increment */
	ScancodeChannelDecrement Scancode = 261 /**< Channel Decrement */

	ScancodeMediaPlay          Scancode = 262 /**< Play */
	ScancodeMediaPause         Scancode = 263 /**< Pause */
	ScancodeMediaRecord        Scancode = 264 /**< Record */
	ScancodeMediaFastForward   Scancode = 265 /**< Fast Forward */
	ScancodeMediaRewind        Scancode = 266 /**< Rewind */
	ScancodeMediaNextTrack     Scancode = 267 /**< Next Track */
	ScancodeMediaPreviousTrack Scancode = 268 /**< Previous Track */
	ScancodeMediaStop          Scancode = 269 /**< Stop */
	ScancodeMediaEject         Scancode = 270 /**< Eject */
	ScancodeMediaPlayPause     Scancode = 271 /**< Play / Pause */
	ScancodeMediaSelect        Scancode = 272 /* Media Select */

	ScancodeACNew        Scancode = 273 /**< AC New */
	ScancodeACOpen       Scancode = 274 /**< AC Open */
	ScancodeACClose      Scancode = 275 /**< AC Close */
	ScancodeACExit       Scancode = 276 /**< AC Exit */
	ScancodeACSave       Scancode = 277 /**< AC Save */
	ScancodeACPrint      Scancode = 278 /**< AC Print */
	ScancodeACProperties Scancode = 279 /**< AC Properties */

	ScancodeACSearch    Scancode = 280 /**< AC Search */
	ScancodeACHome      Scancode = 281 /**< AC Home */
	ScancodeACBack      Scancode = 282 /**< AC Back */
	ScancodeACForward   Scancode = 283 /**< AC Forward */
	ScancodeACStop      Scancode = 284 /**< AC Stop */
	ScancodeACRefresh   Scancode = 285 /**< AC Refresh */
	ScancodeACBookmarks Scancode = 286 /**< AC Bookmarks */

	/* @} */ /* Usage page 0x0C */

	/**
	 *  \name Mobile keys
	 *
	 *  These are values that are often used on mobile phones.
	 */
	/* @{ */

	ScancodeSoftLeft Scancode = 287 /**< Usually situated below the display on phones and
	  used as a multi-function feature key for selecting
	  a software defined function shown on the bottom left
	  of the display. */
	ScancodeSoftRight Scancode = 288 /**< Usually situated below the display on phones and
	  used as a multi-function feature key for selecting
	  a software defined function shown on the bottom right
	  of the display. */
	ScancodeCall    Scancode = 289 /**< Used for accepting phone calls. */
	ScancodeEndCall Scancode = 290 /**< Used for rejecting phone calls. */

	/* @} */ /* Mobile keys */

	/* Add any other keys here. */

	ScancodeReserved Scancode = 400 /**< 400-500 reserved for dynamic keycodes */

	ScancodeCount Scancode = 512 /**< not a key, just marks the number of scancodes for array bounds */
)
