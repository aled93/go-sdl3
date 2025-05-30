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
 * # CategoryPixels
 *
 * SDL offers facilities for pixel management.
 *
 * Largely these facilities deal with pixel _format_: what does this set of
 * bits represent?
 *
 * If you mostly want to think of a pixel as some combination of red, green,
 * blue, and maybe alpha intensities, this is all pretty straightforward, and
 * in many cases, is enough information to build a perfectly fine game.
 *
 * However, the actual definition of a pixel is more complex than that:
 *
 * Pixels are a representation of a color in a particular color space.
 *
 * The first characteristic of a color space is the color type. SDL
 * understands two different color types, RGB and YCbCr, or in SDL also
 * referred to as YUV.
 *
 * RGB colors consist of red, green, and blue channels of color that are added
 * together to represent the colors we see on the screen.
 *
 * https://en.wikipedia.org/wiki/RGB_color_model
 *
 * YCbCr colors represent colors as a Y luma brightness component and red and
 * blue chroma color offsets. This color representation takes advantage of the
 * fact that the human eye is more sensitive to brightness than the color in
 * an image. The Cb and Cr components are often compressed and have lower
 * resolution than the luma component.
 *
 * https://en.wikipedia.org/wiki/YCbCr
 *
 * When the color information in YCbCr is compressed, the Y pixels are left at
 * full resolution and each Cr and Cb pixel represents an average of the color
 * information in a block of Y pixels. The chroma location determines where in
 * that block of pixels the color information is coming from.
 *
 * The color range defines how much of the pixel to use when converting a
 * pixel into a color on the display. When the full color range is used, the
 * entire numeric range of the pixel bits is significant. When narrow color
 * range is used, for historical reasons, the pixel uses only a portion of the
 * numeric range to represent colors.
 *
 * The color primaries and white point are a definition of the colors in the
 * color space relative to the standard XYZ color space.
 *
 * https://en.wikipedia.org/wiki/CIE_1931_color_space
 *
 * The transfer characteristic, or opto-electrical transfer function (OETF),
 * is the way a color is converted from mathematically linear space into a
 * non-linear output signals.
 *
 * https://en.wikipedia.org/wiki/Rec._709#Transfer_characteristics
 *
 * The matrix coefficients are used to convert between YCbCr and RGB colors.
 */
package sdl

import (
	"encoding/binary"
)

/**
* A fully opaque 8-bit alpha value.
*
* \since This macro is available since SDL 3.2.0.
*
* \sa SDL_ALPHA_TRANSPARENT
 */
const ALPHA_OPAQUE = 255

/**
* A fully opaque floating point alpha value.
*
* \since This macro is available since SDL 3.2.0.
*
* \sa ALPHA_TRANSPARENT_FLOAT
 */
const ALPHA_OPAQUE_FLOAT = 1.0

/**
* A fully transparent 8-bit alpha value.
*
* \since This macro is available since SDL 3.2.0.
*
* \sa SDL_ALPHA_OPAQUE
 */
const ALPHA_TRANSPARENT = 0

/**
* A fully transparent floating point alpha value.
*
* \since This macro is available since SDL 3.2.0.
*
* \sa SDL_ALPHA_OPAQUE_FLOAT
 */
const ALPHA_TRANSPARENT_FLOAT = 0.0

/**
* Pixel type.
*
* \since This enum is available since SDL 3.2.0.
 */
type PixelType int32

const (
	PixelTypeUnknown PixelType = iota
	PixelTypeIndex1
	PixelTypeIndex4
	PixelTypeIndex8
	PixelTypePacked8
	PixelTypePacked16
	PixelTypePacked32
	PixelTypeArrayu8
	PixelTypeArrayu16
	PixelTypeArrayu32
	PixelTypeArrayf16
	PixelTypeArrayf32
	/* appended at the end for compatibility with sdl2-compat:  */
	PixelTypeIndex2
)

/**
* Bitmap pixel order, high bit -> low bit.
*
* \since This enum is available since SDL 3.2.0.
 */
type BitmapOrder int32

const (
	BitmapOrderNone BitmapOrder = iota
	BitmapOrder4321
	BitmapOrder1234
)

/**
* Packed component order, high bit -> low bit.
*
* \since This enum is available since SDL 3.2.0.
 */
type PackedOrder int32

const (
	PackedOrderNone PackedOrder = iota
	PackedOrderXRGB
	PackedOrderRGBX
	PackedOrderARGB
	PackedOrderRGBA
	PackedOrderXBGR
	PackedOrderBGRX
	PackedOrderABGR
	PackedOrderBGRA
)

/**
* Array component order, low byte -> high byte.
*
* \since This enum is available since SDL 3.2.0.
 */
type ArrayOrder int32

const (
	ArrayOrderNone ArrayOrder = iota
	ArrayOrderRGB
	ArrayOrderRGBA
	ArrayOrderARGB
	ArrayOrderBGR
	ArrayOrderBGRA
	ArrayOrderABGR
)

/**
* Packed component layout.
*
* \since This enum is available since SDL 3.2.0.
 */
type PackedLayout int32

const (
	PackedLayoutNone PackedLayout = iota
	PackedLayout332
	PackedLayout4444
	PackedLayout1555
	PackedLayout5551
	PackedLayout565
	PackedLayout8888
	PackedLayout2101010
	PackedLayout1010102
)

/**
* A macro for defining custom FourCC pixel formats.
*
* For example, defining SDL_PIXELFORMAT_YV12 looks like this:
*
* ```c
* SDL_DEFINE_PIXELFOURCC('Y', 'V', '1', '2')
* ```
*
* \param A the first character of the FourCC code.
* \param B the second character of the FourCC code.
* \param C the third character of the FourCC code.
* \param D the fourth character of the FourCC code.
* \returns a format value in the style of SDL_PixelFormat.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func DEFINE_PIXELFOURCC(A, B, C, D byte) uint32 {
	return FOURCC(A, B, C, D)
}

/**
* A macro for defining custom non-FourCC pixel formats.
*
* For example, defining SDL_PIXELFORMAT_RGBA8888 looks like this:
*
* ```c
* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_RGBA, SDL_PACKEDLAYOUT_8888, 32, 4)
* ```
*
* \param type the type of the new format, probably a SDL_PixelType value.
* \param order the order of the new format, probably a SDL_BitmapOrder,
*              SDL_PackedOrder, or SDL_ArrayOrder value.
* \param layout the layout of the new format, probably an SDL_PackedLayout
*               value or zero.
* \param bits the number of bits per pixel of the new format.
* \param bytes the number of bytes per pixel of the new format.
* \returns a format value in the style of SDL_PixelFormat.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func DEFINE_PIXELFORMAT[O BitmapOrder | PackedOrder | ArrayOrder](typ PixelType,
	order O, layout PackedLayout, bits, bytes byte) PixelFormat {
	return PixelFormat((int32(1) << 28) | (int32(typ) << 24) | (int32(order) << 20) |
		(int32(layout) << 16) | (int32(bits) << 8) | (int32(bytes) << 0))
}

/**
* A macro to retrieve the flags of an SDL_PixelFormat.
*
* This macro is generally not needed directly by an app, which should use
* specific tests, like SDL_ISPIXELFORMAT_FOURCC, instead.
*
* \param format an SDL_PixelFormat to check.
* \returns the flags of `format`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func PIXELFLAG(format PixelFormat) int32 {
	return (int32(format) >> 28) & 0x0F
}

/**
* A macro to retrieve the type of an SDL_PixelFormat.
*
* This is usually a value from the SDL_PixelType enumeration.
*
* \param format an SDL_PixelFormat to check.
* \returns the type of `format`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func PIXELTYPE(format PixelFormat) PixelType {
	return PixelType((format >> 24) & 0x0F)
}

/**
* A macro to retrieve the order of an SDL_PixelFormat.
*
* This is usually a value from the SDL_BitmapOrder, SDL_PackedOrder, or
* SDL_ArrayOrder enumerations, depending on the format type.
*
* \param format an SDL_PixelFormat to check.
* \returns the order of `format`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func PIXELORDER(format PixelFormat) int32 {
	return (int32(format) >> 20) & 0x0F
}

/**
* A macro to retrieve the layout of an SDL_PixelFormat.
*
* This is usually a value from the SDL_PackedLayout enumeration, or zero if a
* layout doesn't make sense for the format type.
*
* \param format an SDL_PixelFormat to check.
* \returns the layout of `format`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func PIXELLAYOUT(format PixelFormat) PackedLayout {
	return PackedLayout((int32(format) >> 16) & 0x0F)
}

/**
* A macro to determine an SDL_PixelFormat's bits per pixel.
*
* Note that this macro double-evaluates its parameter, so do not use
* expressions with side-effects here.
*
* FourCC formats will report zero here, as it rarely makes sense to measure
* them per-pixel.
*
* \param format an SDL_PixelFormat to check.
* \returns the bits-per-pixel of `format`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
*
* \sa SDL_BYTESPERPIXEL
 */
func BITSPERPIXEL(format PixelFormat) int32 {
	if ISPIXELFORMAT_FOURCC(format) {
		return 0
	} else {
		return (int32(format) >> 8) & 0xFF
	}
}

/**
* A macro to determine an SDL_PixelFormat's bytes per pixel.
*
* Note that this macro double-evaluates its parameter, so do not use
* expressions with side-effects here.
*
* FourCC formats do their best here, but many of them don't have a meaningful
* measurement of bytes per pixel.
*
* \param format an SDL_PixelFormat to check.
* \returns the bytes-per-pixel of `format`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
*
* \sa SDL_BITSPERPIXEL
 */
func BYTESPERPIXEL(format PixelFormat) int32 {
	if ISPIXELFORMAT_FOURCC(format) {
		if (format == PixelFormatYUY2) ||
			(format == PixelFormatUYVY) ||
			(format == PixelFormatYVYU) ||
			(format == PixelFormatP010) {
			return 2
		} else {
			return 1
		}
	} else {
		return int32(format) & 0xFF
	}
}

/**
* A macro to determine if an SDL_PixelFormat is an indexed format.
*
* Note that this macro double-evaluates its parameter, so do not use
* expressions with side-effects here.
*
* \param format an SDL_PixelFormat to check.
* \returns true if the format is indexed, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISPIXELFORMAT_INDEXED(format PixelFormat) bool {
	return !ISPIXELFORMAT_FOURCC(format) &&
		((PIXELTYPE(format) == PixelTypeIndex1) ||
			(PIXELTYPE(format) == PixelTypeIndex2) ||
			(PIXELTYPE(format) == PixelTypeIndex4) ||
			(PIXELTYPE(format) == PixelTypeIndex8))
}

/**
* A macro to determine if an SDL_PixelFormat is a packed format.
*
* Note that this macro double-evaluates its parameter, so do not use
* expressions with side-effects here.
*
* \param format an SDL_PixelFormat to check.
* \returns true if the format is packed, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISPIXELFORMAT_PACKED(format PixelFormat) bool {
	return !ISPIXELFORMAT_FOURCC(format) &&
		((PIXELTYPE(format) == PixelTypePacked8) ||
			(PIXELTYPE(format) == PixelTypePacked16) ||
			(PIXELTYPE(format) == PixelTypePacked32))
}

/**
* A macro to determine if an SDL_PixelFormat is an array format.
*
* Note that this macro double-evaluates its parameter, so do not use
* expressions with side-effects here.
*
* \param format an SDL_PixelFormat to check.
* \returns true if the format is an array, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISPIXELFORMAT_ARRAY(format PixelFormat) bool {
	return !ISPIXELFORMAT_FOURCC(format) &&
		((PIXELTYPE(format) == PixelTypeArrayu8) ||
			(PIXELTYPE(format) == PixelTypeArrayu16) ||
			(PIXELTYPE(format) == PixelTypeArrayu32) ||
			(PIXELTYPE(format) == PixelTypeArrayf16) ||
			(PIXELTYPE(format) == PixelTypeArrayf32))
}

/**
* A macro to determine if an SDL_PixelFormat is a 10-bit format.
*
* Note that this macro double-evaluates its parameter, so do not use
* expressions with side-effects here.
*
* \param format an SDL_PixelFormat to check.
* \returns true if the format is 10-bit, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISPIXELFORMAT_10BIT(format PixelFormat) bool {
	return !ISPIXELFORMAT_FOURCC(format) &&
		((PIXELTYPE(format) == PixelTypePacked32) &&
			(PIXELLAYOUT(format) == PackedLayout2101010))
}

/**
* A macro to determine if an SDL_PixelFormat is a floating point format.
*
* Note that this macro double-evaluates its parameter, so do not use
* expressions with side-effects here.
*
* \param format an SDL_PixelFormat to check.
* \returns true if the format is 10-bit, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISPIXELFORMAT_FLOAT(format PixelFormat) bool {
	return !ISPIXELFORMAT_FOURCC(format) &&
		((PIXELTYPE(format) == PixelTypeArrayf16) ||
			(PIXELTYPE(format) == PixelTypeArrayf32))
}

/**
* A macro to determine if an SDL_PixelFormat has an alpha channel.
*
* Note that this macro double-evaluates its parameter, so do not use
* expressions with side-effects here.
*
* \param format an SDL_PixelFormat to check.
* \returns true if the format has alpha, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISPIXELFORMAT_ALPHA(format PixelFormat) bool {
	return (ISPIXELFORMAT_PACKED(format) &&
		((PackedOrder(PIXELORDER(format)) == PackedOrderARGB) ||
			(PackedOrder(PIXELORDER(format)) == PackedOrderRGBA) ||
			(PackedOrder(PIXELORDER(format)) == PackedOrderABGR) ||
			(PackedOrder(PIXELORDER(format)) == PackedOrderBGRA))) ||
		(ISPIXELFORMAT_ARRAY(format) &&
			((ArrayOrder(PIXELORDER(format)) == ArrayOrderARGB) ||
				(ArrayOrder(PIXELORDER(format)) == ArrayOrderRGBA) ||
				(ArrayOrder(PIXELORDER(format)) == ArrayOrderABGR) ||
				(ArrayOrder(PIXELORDER(format)) == ArrayOrderBGRA)))
}

/**
* A macro to determine if an SDL_PixelFormat is a "FourCC" format.
*
* This covers custom and other unusual formats.
*
* Note that this macro double-evaluates its parameter, so do not use
* expressions with side-effects here.
*
* \param format an SDL_PixelFormat to check.
* \returns true if the format has alpha, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISPIXELFORMAT_FOURCC(format PixelFormat) bool {
	/* The flag is set to 1 because 0x1? is not in the printable ASCII range */
	return format != 0 && (PIXELFLAG(format) != 1)
}

/* Note: If you modify this enum, update SDL_GetPixelFormatName() */

/**
* Pixel format.
*
* SDL's pixel formats have the following naming convention:
*
* - Names with a list of components and a single bit count, such as RGB24 and
*   ABGR32, define a platform-independent encoding into bytes in the order
*   specified. For example, in RGB24 data, each pixel is encoded in 3 bytes
*   (red, green, blue) in that order, and in ABGR32 data, each pixel is
*   encoded in 4 bytes alpha, blue, green, red) in that order. Use these
*   names if the property of a format that is important to you is the order
*   of the bytes in memory or on disk.
* - Names with a bit count per component, such as ARGB8888 and XRGB1555, are
*   "packed" into an appropriately-sized integer in the platform's native
*   endianness. For example, ARGB8888 is a sequence of 32-bit integers; in
*   each integer, the most significant bits are alpha, and the least
*   significant bits are blue. On a little-endian CPU such as x86, the least
*   significant bits of each integer are arranged first in memory, but on a
*   big-endian CPU such as s390x, the most significant bits are arranged
*   first. Use these names if the property of a format that is important to
*   you is the meaning of each bit position within a native-endianness
*   integer.
* - In indexed formats such as INDEX4LSB, each pixel is represented by
*   encoding an index into the palette into the indicated number of bits,
*   with multiple pixels packed into each byte if appropriate. In LSB
*   formats, the first (leftmost) pixel is stored in the least-significant
*   bits of the byte; in MSB formats, it's stored in the most-significant
*   bits. INDEX8 does not need LSB/MSB variants, because each pixel exactly
*   fills one byte.
*
* The 32-bit byte-array encodings such as RGBA32 are aliases for the
* appropriate 8888 encoding for the current platform. For example, RGBA32 is
* an alias for ABGR8888 on little-endian CPUs like x86, or an alias for
* RGBA8888 on big-endian CPUs.
*
* \since This enum is available since SDL 3.2.0.
 */
type PixelFormat int32

const (
	PixelFormatUnknown   PixelFormat = 0
	PixelFormatIndex1LSB PixelFormat = 0x11100100
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX1, SDL_BITMAPORDER_4321, 0, 1, 0), */
	PixelFormatIndex1MSB PixelFormat = 0x11200100
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX1, SDL_BITMAPORDER_1234, 0, 1, 0), */
	PixelFormatIndex2LSB PixelFormat = 0x1c100200
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX2, SDL_BITMAPORDER_4321, 0, 2, 0), */
	PixelFormatIndex2MSB PixelFormat = 0x1c200200
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX2, SDL_BITMAPORDER_1234, 0, 2, 0), */
	PixelFormatIndex4LSB PixelFormat = 0x12100400
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX4, SDL_BITMAPORDER_4321, 0, 4, 0), */
	PixelFormatIndex4MSB PixelFormat = 0x12200400
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX4, SDL_BITMAPORDER_1234, 0, 4, 0), */
	PixelFormatIndex8 PixelFormat = 0x13000801
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX8, 0, 0, 8, 1), */
	PixelFormatRGB332 PixelFormat = 0x14110801
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED8, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_332, 8, 1), */
	PixelFormatXRGB4444 PixelFormat = 0x15120c02
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_4444, 12, 2), */
	PixelFormatXBGR4444 PixelFormat = 0x15520c02
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_4444, 12, 2), */
	PixelFormatXRGB1555 PixelFormat = 0x15130f02
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_1555, 15, 2), */
	PixelFormatXBGR1555 PixelFormat = 0x15530f02
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_1555, 15, 2), */
	PixelFormatARGB4444 PixelFormat = 0x15321002
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_4444, 16, 2), */
	PixelFormatRGBA4444 PixelFormat = 0x15421002
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_RGBA, SDL_PACKEDLAYOUT_4444, 16, 2), */
	PixelFormatABGR4444 PixelFormat = 0x15721002
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ABGR, SDL_PACKEDLAYOUT_4444, 16, 2), */
	PixelFormatBGRA4444 PixelFormat = 0x15821002
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_BGRA, SDL_PACKEDLAYOUT_4444, 16, 2), */
	PixelFormatARGB1555 PixelFormat = 0x15331002
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_1555, 16, 2), */
	PixelFormatRGBA5551 PixelFormat = 0x15441002
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_RGBA, SDL_PACKEDLAYOUT_5551, 16, 2), */
	PixelFormatABGR1555 PixelFormat = 0x15731002
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ABGR, SDL_PACKEDLAYOUT_1555, 16, 2), */
	PixelFormatBGRA5551 PixelFormat = 0x15841002
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_BGRA, SDL_PACKEDLAYOUT_5551, 16, 2), */
	PixelFormatRGB565 PixelFormat = 0x15151002
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_565, 16, 2), */
	PixelFormatBGR565 PixelFormat = 0x15551002
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_565, 16, 2), */
	PixelFormatRGB24 PixelFormat = 0x17101803
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU8, SDL_ARRAYORDER_RGB, 0, 24, 3), */
	PixelFormatBGR24 PixelFormat = 0x17401803
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU8, SDL_ARRAYORDER_BGR, 0, 24, 3), */
	PixelFormatXRGB8888 PixelFormat = 0x16161804
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_8888, 24, 4), */
	PixelFormatRGBX8888 PixelFormat = 0x16261804
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_RGBX, SDL_PACKEDLAYOUT_8888, 24, 4), */
	PixelFormatXBGR8888 PixelFormat = 0x16561804
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_8888, 24, 4), */
	PixelFormatBGRX8888 PixelFormat = 0x16661804
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_BGRX, SDL_PACKEDLAYOUT_8888, 24, 4), */
	PixelFormatARGB8888 PixelFormat = 0x16362004
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_8888, 32, 4), */
	PixelFormatRGBA8888 PixelFormat = 0x16462004
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_RGBA, SDL_PACKEDLAYOUT_8888, 32, 4), */
	PixelFormatABGR8888 PixelFormat = 0x16762004
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ABGR, SDL_PACKEDLAYOUT_8888, 32, 4), */
	PixelFormatBGRA8888 PixelFormat = 0x16862004
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_BGRA, SDL_PACKEDLAYOUT_8888, 32, 4), */
	PixelFormatXRGB2101010 PixelFormat = 0x16172004
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_2101010, 32, 4), */
	PixelFormatXBGR2101010 PixelFormat = 0x16572004
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_2101010, 32, 4), */
	PixelFormatARGB2101010 PixelFormat = 0x16372004
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_2101010, 32, 4), */
	PixelFormatABGR2101010 PixelFormat = 0x16772004
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ABGR, SDL_PACKEDLAYOUT_2101010, 32, 4), */
	PixelFormatRGB48 PixelFormat = 0x18103006
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU16, SDL_ARRAYORDER_RGB, 0, 48, 6), */
	PixelFormatBGR48 PixelFormat = 0x18403006
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU16, SDL_ARRAYORDER_BGR, 0, 48, 6), */
	PixelFormatRGBA64 PixelFormat = 0x18204008
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU16, SDL_ARRAYORDER_RGBA, 0, 64, 8), */
	PixelFormatARGB64 PixelFormat = 0x18304008
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU16, SDL_ARRAYORDER_ARGB, 0, 64, 8), */
	PixelFormatBGRA64 PixelFormat = 0x18504008
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU16, SDL_ARRAYORDER_BGRA, 0, 64, 8), */
	PixelFormatABGR64 PixelFormat = 0x18604008
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU16, SDL_ARRAYORDER_ABGR, 0, 64, 8), */
	PixelFormatRGB48Float PixelFormat = 0x1a103006
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF16, SDL_ARRAYORDER_RGB, 0, 48, 6), */
	PixelFormatBGR48Float PixelFormat = 0x1a403006
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF16, SDL_ARRAYORDER_BGR, 0, 48, 6), */
	PixelFormatRGBA64Float PixelFormat = 0x1a204008
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF16, SDL_ARRAYORDER_RGBA, 0, 64, 8), */
	PixelFormatARGB64Float PixelFormat = 0x1a304008
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF16, SDL_ARRAYORDER_ARGB, 0, 64, 8), */
	PixelFormatBGRA64Float PixelFormat = 0x1a504008
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF16, SDL_ARRAYORDER_BGRA, 0, 64, 8), */
	PixelFormatABGR64Float PixelFormat = 0x1a604008
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF16, SDL_ARRAYORDER_ABGR, 0, 64, 8), */
	PixelFormatRGB96Float PixelFormat = 0x1b10600c
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF32, SDL_ARRAYORDER_RGB, 0, 96, 12), */
	PixelFormatBGR96Float PixelFormat = 0x1b40600c
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF32, SDL_ARRAYORDER_BGR, 0, 96, 12), */
	PixelFormatRGBA128Float PixelFormat = 0x1b208010
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF32, SDL_ARRAYORDER_RGBA, 0, 128, 16), */
	PixelFormatARGB128Float PixelFormat = 0x1b308010
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF32, SDL_ARRAYORDER_ARGB, 0, 128, 16), */
	PixelFormatBGRA128Float PixelFormat = 0x1b508010
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF32, SDL_ARRAYORDER_BGRA, 0, 128, 16), */
	PixelFormatABGR128Float PixelFormat = 0x1b608010
	/* SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYF32, SDL_ARRAYORDER_ABGR, 0, 128, 16), */

	PixelFormatYV12 PixelFormat = 0x32315659 /**< Planar mode: Y + V + U  (3 planes) */
	/* SDL_DEFINE_PIXELFOURCC('Y', 'V', '1', '2'), */
	PixelFormatIYUV PixelFormat = 0x56555949 /**< Planar mode: Y + U + V  (3 planes) */
	/* SDL_DEFINE_PIXELFOURCC('I', 'Y', 'U', 'V'), */
	PixelFormatYUY2 PixelFormat = 0x32595559 /**< Packed mode: Y0+U0+Y1+V0 (1 plane) */
	/* SDL_DEFINE_PIXELFOURCC('Y', 'U', 'Y', '2'), */
	PixelFormatUYVY PixelFormat = 0x59565955 /**< Packed mode: U0+Y0+V0+Y1 (1 plane) */
	/* SDL_DEFINE_PIXELFOURCC('U', 'Y', 'V', 'Y'), */
	PixelFormatYVYU PixelFormat = 0x55595659 /**< Packed mode: Y0+V0+Y1+U0 (1 plane) */
	/* SDL_DEFINE_PIXELFOURCC('Y', 'V', 'Y', 'U'), */
	PixelFormatNV12 PixelFormat = 0x3231564e /**< Planar mode: Y + U/V interleaved  (2 planes) */
	/* SDL_DEFINE_PIXELFOURCC('N', 'V', '1', '2'), */
	PixelFormatNV21 PixelFormat = 0x3132564e /**< Planar mode: Y + V/U interleaved  (2 planes) */
	/* SDL_DEFINE_PIXELFOURCC('N', 'V', '2', '1'), */
	PixelFormatP010 PixelFormat = 0x30313050 /**< Planar mode: Y + U/V interleaved  (2 planes) */
	/* SDL_DEFINE_PIXELFOURCC('P', '0', '1', '0'), */
	PixelFormatExternalOES PixelFormat = 0x2053454f /**< Android video texture format */
	/* SDL_DEFINE_PIXELFOURCC('O', 'E', 'S', ' ') */

	PixelFormatMJPG PixelFormat = 0x47504a4d /**< Motion JPEG */
	/* SDL_DEFINE_PIXELFOURCC('M', 'J', 'P', 'G') */
)

var (
	/* Aliases for RGBA byte arrays of color data, for the current platform */
	PixelFormatRGBA32 PixelFormat
	PixelFormatARGB32 PixelFormat
	PixelFormatBGRA32 PixelFormat
	PixelFormatABGR32 PixelFormat
	PixelFormatRGBX32 PixelFormat
	PixelFormatXRGB32 PixelFormat
	PixelFormatBGRX32 PixelFormat
	PixelFormatXBGR32 PixelFormat
)

func init() {
	if binary.NativeEndian.Uint16([]byte{0x12, 0x34}) == uint16(0x3412) {
		// little endian
		PixelFormatRGBA32 = PixelFormatABGR8888
		PixelFormatARGB32 = PixelFormatBGRA8888
		PixelFormatBGRA32 = PixelFormatARGB8888
		PixelFormatABGR32 = PixelFormatRGBA8888
		PixelFormatRGBX32 = PixelFormatXBGR8888
		PixelFormatXRGB32 = PixelFormatBGRX8888
		PixelFormatBGRX32 = PixelFormatXRGB8888
		PixelFormatXBGR32 = PixelFormatRGBX8888
	} else {
		// big endian
		PixelFormatRGBA32 = PixelFormatRGBA8888
		PixelFormatARGB32 = PixelFormatARGB8888
		PixelFormatBGRA32 = PixelFormatBGRA8888
		PixelFormatABGR32 = PixelFormatABGR8888
		PixelFormatRGBX32 = PixelFormatRGBX8888
		PixelFormatXRGB32 = PixelFormatXRGB8888
		PixelFormatBGRX32 = PixelFormatBGRX8888
		PixelFormatXBGR32 = PixelFormatXBGR8888
	}
}

/**
* Colorspace color type.
*
* \since This enum is available since SDL 3.2.0.
 */
type ColorType int32

const (
	ColorTypeUnknown ColorType = 0
	ColorTypeRGB     ColorType = 1
	ColorTypeYCBCR   ColorType = 2
)

/**
* Colorspace color range, as described by
* https://www.itu.int/rec/R-REC-BT.2100-2-201807-I/en
*
* \since This enum is available since SDL 3.2.0.
 */
type ColorRange int32

const (
	ColorRangeUnknown ColorRange = 0
	ColorRangeLimited ColorRange = 1 /**< Narrow range, e.g. 16-235 for 8-bit RGB and luma, and 16-240 for 8-bit chroma */
	ColorRangeFull    ColorRange = 2 /**< Full range, e.g. 0-255 for 8-bit RGB and luma, and 1-255 for 8-bit chroma */
)

/**
* Colorspace color primaries, as described by
* https://www.itu.int/rec/T-REC-H.273-201612-S/en
*
* \since This enum is available since SDL 3.2.0.
 */
type ColorPrimaries int32

const (
	ColorPrimariesUnknown     ColorPrimaries = 0
	ColorPrimariesBT709       ColorPrimaries = 1 /**< ITU-R BT.709-6 */
	ColorPrimariesUnspecified ColorPrimaries = 2
	ColorPrimariesBT470M      ColorPrimaries = 4  /**< ITU-R BT.470-6 System M */
	ColorPrimariesBT470BG     ColorPrimaries = 5  /**< ITU-R BT.470-6 System B, G / ITU-R BT.601-7 625 */
	ColorPrimariesBT601       ColorPrimaries = 6  /**< ITU-R BT.601-7 525, SMPTE 170M */
	ColorPrimariesSMPTE240    ColorPrimaries = 7  /**< SMPTE 240M, functionally the same as SDL_COLOR_PRIMARIES_BT601 */
	ColorPrimariesGenericFilm ColorPrimaries = 8  /**< Generic film (color filters using Illuminant C) */
	ColorPrimariesBT2020      ColorPrimaries = 9  /**< ITU-R BT.2020-2 / ITU-R BT.2100-0 */
	ColorPrimariesXYZ         ColorPrimaries = 10 /**< SMPTE ST 428-1 */
	ColorPrimariesSMPTE431    ColorPrimaries = 11 /**< SMPTE RP 431-2 */
	ColorPrimariesSMPTE432    ColorPrimaries = 12 /**< SMPTE EG 432-1 / DCI P3 */
	ColorPrimariesEBU3213     ColorPrimaries = 22 /**< EBU Tech. 3213-E */
	ColorPrimariesCustom      ColorPrimaries = 31
)

/**
* Colorspace transfer characteristics.
*
* These are as described by https://www.itu.int/rec/T-REC-H.273-201612-S/en
*
* \since This enum is available since SDL 3.2.0.
 */
type TransferCharacteristics int32

const (
	TransferCharacteristicsUnknown       TransferCharacteristics = 0
	TransferCharacteristicsBT709         TransferCharacteristics = 1 /**< Rec. ITU-R BT.709-6 / ITU-R BT1361 */
	TransferCharacteristicsUnspecified   TransferCharacteristics = 2
	TransferCharacteristicsGamma22       TransferCharacteristics = 4 /**< ITU-R BT.470-6 System M / ITU-R BT1700 625 PAL & SECAM */
	TransferCharacteristicsGamma28       TransferCharacteristics = 5 /**< ITU-R BT.470-6 System B, G */
	TransferCharacteristicsBT601         TransferCharacteristics = 6 /**< SMPTE ST 170M / ITU-R BT.601-7 525 or 625 */
	TransferCharacteristicsSMPTE240      TransferCharacteristics = 7 /**< SMPTE ST 240M */
	TransferCharacteristicsLinear        TransferCharacteristics = 8
	TransferCharacteristicsLog100        TransferCharacteristics = 9
	TransferCharacteristicsLog100_Sqrt10 TransferCharacteristics = 10
	TransferCharacteristicsIEC61966      TransferCharacteristics = 11 /**< IEC 61966-2-4 */
	TransferCharacteristicsBT1361        TransferCharacteristics = 12 /**< ITU-R BT1361 Extended Colour Gamut */
	TransferCharacteristicsSRGB          TransferCharacteristics = 13 /**< IEC 61966-2-1 (sRGB or sYCC) */
	TransferCharacteristicsBT2020_10Bit  TransferCharacteristics = 14 /**< ITU-R BT2020 for 10-bit system */
	TransferCharacteristicsBT2020_12Bit  TransferCharacteristics = 15 /**< ITU-R BT2020 for 12-bit system */
	TransferCharacteristicsPQ            TransferCharacteristics = 16 /**< SMPTE ST 2084 for 10-, 12-, 14- and 16-bit systems */
	TransferCharacteristicsSMPTE428      TransferCharacteristics = 17 /**< SMPTE ST 428-1 */
	TransferCharacteristicsHLG           TransferCharacteristics = 18 /**< ARIB STD-B67, known as "hybrid log-gamma" (HLG) */
	TransferCharacteristicsCustom        TransferCharacteristics = 31
)

/**
* Colorspace matrix coefficients.
*
* These are as described by https://www.itu.int/rec/T-REC-H.273-201612-S/en
*
* \since This enum is available since SDL 3.2.0.
 */
type MatrixCoefficients int32

const (
	MatrixCoefficientsIdentity         MatrixCoefficients = 0
	MatrixCoefficientsBT709            MatrixCoefficients = 1 /**< ITU-R BT.709-6 */
	MatrixCoefficientsUnspecified      MatrixCoefficients = 2
	MatrixCoefficientsFCC              MatrixCoefficients = 4 /**< US FCC Title 47 */
	MatrixCoefficientsBT470BG          MatrixCoefficients = 5 /**< ITU-R BT.470-6 System B, G / ITU-R BT.601-7 625, functionally the same as SDL_MATRIX_COEFFICIENTS_BT601 */
	MatrixCoefficientsBT601            MatrixCoefficients = 6 /**< ITU-R BT.601-7 525 */
	MatrixCoefficientsSMPTE240         MatrixCoefficients = 7 /**< SMPTE 240M */
	MatrixCoefficientsYCGCO            MatrixCoefficients = 8
	MatrixCoefficientsBT2020NCL        MatrixCoefficients = 9  /**< ITU-R BT.2020-2 non-constant luminance */
	MatrixCoefficientsBT2020CL         MatrixCoefficients = 10 /**< ITU-R BT.2020-2 constant luminance */
	MatrixCoefficientsSMPTE2085        MatrixCoefficients = 11 /**< SMPTE ST 2085 */
	MatrixCoefficientsChromaDerivedNCL MatrixCoefficients = 12
	MatrixCoefficientsChromaDerivedCL  MatrixCoefficients = 13
	MatrixCoefficientsICTCP            MatrixCoefficients = 14 /**< ITU-R BT.2100-0 ICTCP */
	MatrixCoefficientsCustom           MatrixCoefficients = 31
)

/**
* Colorspace chroma sample location.
*
* \since This enum is available since SDL 3.2.0.
 */
type ChromaLocation int32

const (
	ChromaLocationNone    ChromaLocation = 0 /**< RGB, no chroma sampling */
	ChromaLocationLeft    ChromaLocation = 1 /**< In MPEG-2, MPEG-4, and AVC, Cb and Cr are taken on midpoint of the left-edge of the 2x2 square. In other words, they have the same horizontal location as the top-left pixel, but is shifted one-half pixel down vertically. */
	ChromaLocationCenter  ChromaLocation = 2 /**< In JPEG/JFIF, H.261, and MPEG-1, Cb and Cr are taken at the center of the 2x2 square. In other words, they are offset one-half pixel to the right and one-half pixel down compared to the top-left pixel. */
	ChromaLocationTopLeft ChromaLocation = 3 /**< In HEVC for BT.2020 and BT.2100 content (in particular on Blu-rays), Cb and Cr are sampled at the same location as the group's top-left Y pixel ("co-sited", "co-located"). */
)

/* Colorspace definition */

/**
* A macro for defining custom SDL_Colorspace formats.
*
* For example, defining SDL_COLORSPACE_SRGB looks like this:
*
* ```c
* SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_RGB,
*                       SDL_COLOR_RANGE_FULL,
*                       SDL_COLOR_PRIMARIES_BT709,
*                       SDL_TRANSFER_CHARACTERISTICS_SRGB,
*                       SDL_MATRIX_COEFFICIENTS_IDENTITY,
*                       SDL_CHROMA_LOCATION_NONE)
* ```
*
* \param type the type of the new format, probably an SDL_ColorType value.
* \param range the range of the new format, probably a SDL_ColorRange value.
* \param primaries the primaries of the new format, probably an
*                  SDL_ColorPrimaries value.
* \param transfer the transfer characteristics of the new format, probably an
*                 SDL_TransferCharacteristics value.
* \param matrix the matrix coefficients of the new format, probably an
*               SDL_MatrixCoefficients value.
* \param chroma the chroma sample location of the new format, probably an
*               SDL_ChromaLocation value.
* \returns a format value in the style of SDL_Colorspace.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func DEFINE_COLORSPACE(type_ ColorType, range_ ColorRange,
	primaries ColorPrimaries, transfer TransferCharacteristics,
	matrix MatrixCoefficients, chroma ChromaLocation) Colorspace {
	return Colorspace((uint32(type_) << 28) | (uint32(range_) << 24) | (uint32(chroma) << 20) |
		(uint32(primaries) << 10) | (uint32(transfer) << 5) | (uint32(matrix) << 0))
}

/**
* A macro to retrieve the type of an SDL_Colorspace.
*
* \param cspace an SDL_Colorspace to check.
* \returns the SDL_ColorType for `cspace`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func COLORSPACETYPE(cspace Colorspace) ColorType {
	return ColorType((cspace >> 28) & 0x0F)
}

/**
* A macro to retrieve the range of an SDL_Colorspace.
*
* \param cspace an SDL_Colorspace to check.
* \returns the SDL_ColorRange of `cspace`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func COLORSPACERANGE(cspace Colorspace) ColorRange {
	return ColorRange((cspace >> 24) & 0x0F)
}

/**
* A macro to retrieve the chroma sample location of an SDL_Colorspace.
*
* \param cspace an SDL_Colorspace to check.
* \returns the SDL_ChromaLocation of `cspace`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func COLORSPACECHROMA(cspace Colorspace) ChromaLocation {
	return ChromaLocation((cspace >> 20) & 0x0F)
}

/**
* A macro to retrieve the primaries of an SDL_Colorspace.
*
* \param cspace an SDL_Colorspace to check.
* \returns the SDL_ColorPrimaries of `cspace`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func COLORSPACEPRIMARIES(cspace Colorspace) ColorPrimaries {
	return ColorPrimaries((cspace >> 10) & 0x1F)
}

/**
* A macro to retrieve the transfer characteristics of an SDL_Colorspace.
*
* \param cspace an SDL_Colorspace to check.
* \returns the SDL_TransferCharacteristics of `cspace`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func COLORSPACETRANSFER(cspace Colorspace) TransferCharacteristics {
	return TransferCharacteristics((cspace >> 5) & 0x1F)
}

/**
* A macro to retrieve the matrix coefficients of an SDL_Colorspace.
*
* \param cspace an SDL_Colorspace to check.
* \returns the SDL_MatrixCoefficients of `cspace`.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func COLORSPACEMATRIX(cspace Colorspace) MatrixCoefficients {
	return MatrixCoefficients(cspace & 0x1F)
}

/**
* A macro to determine if an SDL_Colorspace uses BT601 (or BT470BG) matrix
* coefficients.
*
* Note that this macro double-evaluates its parameter, so do not use
* expressions with side-effects here.
*
* \param cspace an SDL_Colorspace to check.
* \returns true if BT601 or BT470BG, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISCOLORSPACE_MATRIX_BT601(cspace Colorspace) bool {
	return COLORSPACEMATRIX(cspace) == MatrixCoefficientsBT601 || COLORSPACEMATRIX(cspace) == MatrixCoefficientsBT470BG
}

/**
* A macro to determine if an SDL_Colorspace uses BT709 matrix coefficients.
*
* \param cspace an SDL_Colorspace to check.
* \returns true if BT709, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISCOLORSPACE_MATRIX_BT709(cspace Colorspace) bool {
	return COLORSPACEMATRIX(cspace) == MatrixCoefficientsBT709
}

/**
* A macro to determine if an SDL_Colorspace uses BT2020_NCL matrix
* coefficients.
*
* \param cspace an SDL_Colorspace to check.
* \returns true if BT2020_NCL, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISCOLORSPACE_MATRIX_BT2020_NCL(cspace Colorspace) bool {
	return COLORSPACEMATRIX(cspace) == MatrixCoefficientsBT2020NCL
}

/**
* A macro to determine if an SDL_Colorspace has a limited range.
*
* \param cspace an SDL_Colorspace to check.
* \returns true if limited range, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISCOLORSPACE_LIMITED_RANGE(cspace Colorspace) bool {
	return COLORSPACERANGE(cspace) != ColorRangeFull
}

/**
* A macro to determine if an SDL_Colorspace has a full range.
*
* \param cspace an SDL_Colorspace to check.
* \returns true if full range, false otherwise.
*
* \threadsafety It is safe to call this macro from any thread.
*
* \since This macro is available since SDL 3.2.0.
 */
func ISCOLORSPACE_FULL_RANGE(cspace Colorspace) bool {
	return COLORSPACERANGE(cspace) == ColorRangeFull
}

/**
* Colorspace definitions.
*
* Since similar colorspaces may vary in their details (matrix, transfer
* function, etc.), this is not an exhaustive list, but rather a
* representative sample of the kinds of colorspaces supported in SDL.
*
* \since This enum is available since SDL 3.2.0.
*
* \sa SDL_ColorPrimaries
* \sa SDL_ColorRange
* \sa SDL_ColorType
* \sa SDL_MatrixCoefficients
* \sa SDL_TransferCharacteristics
 */
type Colorspace int32

const (
	ColorspaceUnknown Colorspace = 0

	/* sRGB is a gamma corrected colorspace, and the default colorspace for SDL rendering and 8-bit RGB surfaces */
	ColorspaceSRGB Colorspace = 0x120005a0 /**< Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G22_NONE_P709 */
	/* SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_RGB,
	SDL_COLOR_RANGE_FULL,
	SDL_COLOR_PRIMARIES_BT709,
	SDL_TRANSFER_CHARACTERISTICS_SRGB,
	SDL_MATRIX_COEFFICIENTS_IDENTITY,
	SDL_CHROMA_LOCATION_NONE), */

	/* This is a linear colorspace and the default colorspace for floating point surfaces. On Windows this is the scRGB colorspace, and on Apple platforms this is kCGColorSpaceExtendedLinearSRGB for EDR content */
	ColorspaceSRGBLinear Colorspace = 0x12000500 /**< Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G10_NONE_P709  */
	/* SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_RGB,
	SDL_COLOR_RANGE_FULL,
	SDL_COLOR_PRIMARIES_BT709,
	SDL_TRANSFER_CHARACTERISTICS_LINEAR,
	SDL_MATRIX_COEFFICIENTS_IDENTITY,
	SDL_CHROMA_LOCATION_NONE), */

	/* HDR10 is a non-linear HDR colorspace and the default colorspace for 10-bit surfaces */
	ColorspaceHDR10 Colorspace = 0x12002600 /**< Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G2084_NONE_P2020  */
	/* SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_RGB,
	SDL_COLOR_RANGE_FULL,
	SDL_COLOR_PRIMARIES_BT2020,
	SDL_TRANSFER_CHARACTERISTICS_PQ,
	SDL_MATRIX_COEFFICIENTS_IDENTITY,
	SDL_CHROMA_LOCATION_NONE), */

	ColorspaceJPEG Colorspace = 0x220004c6 /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_FULL_G22_NONE_P709_X601 */
	/* SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	SDL_COLOR_RANGE_FULL,
	SDL_COLOR_PRIMARIES_BT709,
	SDL_TRANSFER_CHARACTERISTICS_BT601,
	SDL_MATRIX_COEFFICIENTS_BT601,
	SDL_CHROMA_LOCATION_NONE), */

	ColorspaceBT601Limited Colorspace = 0x211018c6 /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P601 */
	/* SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	SDL_COLOR_RANGE_LIMITED,
	SDL_COLOR_PRIMARIES_BT601,
	SDL_TRANSFER_CHARACTERISTICS_BT601,
	SDL_MATRIX_COEFFICIENTS_BT601,
	SDL_CHROMA_LOCATION_LEFT), */

	ColorspaceBT601Full Colorspace = 0x221018c6 /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P601 */
	/* SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	SDL_COLOR_RANGE_FULL,
	SDL_COLOR_PRIMARIES_BT601,
	SDL_TRANSFER_CHARACTERISTICS_BT601,
	SDL_MATRIX_COEFFICIENTS_BT601,
	SDL_CHROMA_LOCATION_LEFT), */

	ColorspaceBT709Limited Colorspace = 0x21100421 /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P709 */
	/* SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	SDL_COLOR_RANGE_LIMITED,
	SDL_COLOR_PRIMARIES_BT709,
	SDL_TRANSFER_CHARACTERISTICS_BT709,
	SDL_MATRIX_COEFFICIENTS_BT709,
	SDL_CHROMA_LOCATION_LEFT), */

	ColorspaceBT709Full Colorspace = 0x22100421 /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P709 */
	/* SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	SDL_COLOR_RANGE_FULL,
	SDL_COLOR_PRIMARIES_BT709,
	SDL_TRANSFER_CHARACTERISTICS_BT709,
	SDL_MATRIX_COEFFICIENTS_BT709,
	SDL_CHROMA_LOCATION_LEFT), */

	ColorspaceBT2020Limited Colorspace = 0x21102609 /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P2020 */
	/* SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	SDL_COLOR_RANGE_LIMITED,
	SDL_COLOR_PRIMARIES_BT2020,
	SDL_TRANSFER_CHARACTERISTICS_PQ,
	SDL_MATRIX_COEFFICIENTS_BT2020_NCL,
	SDL_CHROMA_LOCATION_LEFT), */

	ColorspaceBT2020Full Colorspace = 0x22102609 /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_FULL_G22_LEFT_P2020 */
	/* SDL_DEFINE_COLORSPACE(SDL_COLOR_TYPE_YCBCR,
	SDL_COLOR_RANGE_FULL,
	SDL_COLOR_PRIMARIES_BT2020,
	SDL_TRANSFER_CHARACTERISTICS_PQ,
	SDL_MATRIX_COEFFICIENTS_BT2020_NCL,
	SDL_CHROMA_LOCATION_LEFT), */

	ColorspaceRGBDefault Colorspace = ColorspaceSRGB /**< The default colorspace for RGB surfaces if no colorspace is specified */
	ColorspaceYUVDefault Colorspace = ColorspaceJPEG /**< The default colorspace for YUV surfaces if no colorspace is specified */
)

/**
* A structure that represents a color as RGBA components.
*
* The bits of this structure can be directly reinterpreted as an
* integer-packed color which uses the SDL_PIXELFORMAT_RGBA32 format
* (SDL_PIXELFORMAT_ABGR8888 on little-endian systems and
* SDL_PIXELFORMAT_RGBA8888 on big-endian systems).
*
* \since This struct is available since SDL 3.2.0.
 */
type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

/**
* The bits of this structure can be directly reinterpreted as a float-packed
* color which uses the SDL_PIXELFORMAT_RGBA128_FLOAT format
*
* \since This struct is available since SDL 3.2.0.
 */
type FColor struct {
	R float32
	G float32
	B float32
	A float32
}

/**
* A set of indexed colors representing a palette.
*
* \since This struct is available since SDL 3.2.0.
*
* \sa SDL_SetPaletteColors
 */
type Palette struct {
	NColors  int32  /**< number of elements in `colors`. */
	Colors   *Color /**< an array of colors, `ncolors` long. */
	Version  uint32 /**< internal use only, do not touch. */
	Refcount int32  /**< internal use only, do not touch. */
}

/**
* Details about the format of a pixel.
*
* \since This struct is available since SDL 3.2.0.
 */
type PixelFormatDetails struct {
	Format        PixelFormat
	BitsPerPixel  uint8
	BytesPerPixel uint8
	Padding       [2]uint8
	RMask         uint32
	GMask         uint32
	BMask         uint32
	AMask         uint32
	RBits         uint8
	GBits         uint8
	BBits         uint8
	ABits         uint8
	RShift        uint8
	GShift        uint8
	BShift        uint8
	AShift        uint8
}

/**
* Get the human readable name of a pixel format.
*
* \param format the pixel format to query.
* \returns the human readable name of the specified pixel format or
*          "SDL_PIXELFORMAT_UNKNOWN" if the format isn't recognized.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetPixelFormatName func(format PixelFormat) *byte

/**
* Convert one of the enumerated pixel formats to a bpp value and RGBA masks.
*
* \param format one of the SDL_PixelFormat values.
* \param bpp a bits per pixel value; usually 15, 16, or 32.
* \param Rmask a pointer filled in with the red mask for the format.
* \param Gmask a pointer filled in with the green mask for the format.
* \param Bmask a pointer filled in with the blue mask for the format.
* \param Amask a pointer filled in with the alpha mask for the format.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetPixelFormatForMasks
 */
//go:sdl3extern
var GetMasksForPixelFormat func(format PixelFormat, bpp *int32, Rmask, Gmask, Bmask, Amask *uint32) bool

/**
* Convert a bpp value and RGBA masks to an enumerated pixel format.
*
* This will return `SDL_PIXELFORMAT_UNKNOWN` if the conversion wasn't
* possible.
*
* \param bpp a bits per pixel value; usually 15, 16, or 32.
* \param Rmask the red mask for the format.
* \param Gmask the green mask for the format.
* \param Bmask the blue mask for the format.
* \param Amask the alpha mask for the format.
* \returns the SDL_PixelFormat value corresponding to the format masks, or
*          SDL_PIXELFORMAT_UNKNOWN if there isn't a match.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetMasksForPixelFormat
 */
//go:sdl3extern
var GetPixelFormatForMasks func(bpp int32, Rmask, Gmask, Bmask, Amask uint32) PixelFormat

/**
* Create an SDL_PixelFormatDetails structure corresponding to a pixel format.
*
* Returned structure may come from a shared global cache (i.e. not newly
* allocated), and hence should not be modified, especially the palette. Weird
* errors such as `Blit combination not supported` may occur.
*
* \param format one of the SDL_PixelFormat values.
* \returns a pointer to a SDL_PixelFormatDetails structure or NULL on
*          failure; call SDL_GetError() for more information.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetPixelFormatDetails func(format PixelFormat) *PixelFormatDetails

/**
* Create a palette structure with the specified number of color entries.
*
* The palette entries are initialized to white.
*
* \param ncolors represents the number of color entries in the color palette.
* \returns a new SDL_Palette structure on success or NULL on failure (e.g. if
*          there wasn't enough memory); call SDL_GetError() for more
*          information.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_DestroyPalette
* \sa SDL_SetPaletteColors
* \sa SDL_SetSurfacePalette
 */
//go:sdl3extern
var CreatePalette func(ncolors int32) *Palette

/**
* Set a range of colors in a palette.
*
* \param palette the SDL_Palette structure to modify.
* \param colors an array of SDL_Color structures to copy into the palette.
* \param firstcolor the index of the first palette entry to modify.
* \param ncolors the number of entries to modify.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety It is safe to call this function from any thread, as long as
*               the palette is not modified or destroyed in another thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetPaletteColors func(palette *Palette, colors *Color, firstcolor int32, ncolors int32) bool

/**
* Free a palette created with SDL_CreatePalette().
*
* \param palette the SDL_Palette structure to be freed.
*
* \threadsafety It is safe to call this function from any thread, as long as
*               the palette is not modified or destroyed in another thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_CreatePalette
 */
//go:sdl3extern
var DestroyPalette func(palette *Palette)

/**
* Map an RGB triple to an opaque pixel value for a given pixel format.
*
* This function maps the RGB color value to the specified pixel format and
* returns the pixel value best approximating the given RGB color value for
* the given pixel format.
*
* If the format has a palette (8-bit) the index of the closest matching color
* in the palette will be returned.
*
* If the specified pixel format has an alpha component it will be returned as
* all 1 bits (fully opaque).
*
* If the pixel format bpp (color depth) is less than 32-bpp then the unused
* upper bits of the return value can safely be ignored (e.g., with a 16-bpp
* format the return value can be assigned to a Uint16, and similarly a Uint8
* for an 8-bpp format).
*
* \param format a pointer to SDL_PixelFormatDetails describing the pixel
*               format.
* \param palette an optional palette for indexed formats, may be NULL.
* \param r the red component of the pixel in the range 0-255.
* \param g the green component of the pixel in the range 0-255.
* \param b the blue component of the pixel in the range 0-255.
* \returns a pixel value.
*
* \threadsafety It is safe to call this function from any thread, as long as
*               the palette is not modified.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetPixelFormatDetails
* \sa SDL_GetRGB
* \sa SDL_MapRGBA
* \sa SDL_MapSurfaceRGB
 */
//go:sdl3extern
var MapRGB func(format *PixelFormatDetails, palette *Palette, r, g, b uint8) uint32

/**
* Map an RGBA quadruple to a pixel value for a given pixel format.
*
* This function maps the RGBA color value to the specified pixel format and
* returns the pixel value best approximating the given RGBA color value for
* the given pixel format.
*
* If the specified pixel format has no alpha component the alpha value will
* be ignored (as it will be in formats with a palette).
*
* If the format has a palette (8-bit) the index of the closest matching color
* in the palette will be returned.
*
* If the pixel format bpp (color depth) is less than 32-bpp then the unused
* upper bits of the return value can safely be ignored (e.g., with a 16-bpp
* format the return value can be assigned to a Uint16, and similarly a Uint8
* for an 8-bpp format).
*
* \param format a pointer to SDL_PixelFormatDetails describing the pixel
*               format.
* \param palette an optional palette for indexed formats, may be NULL.
* \param r the red component of the pixel in the range 0-255.
* \param g the green component of the pixel in the range 0-255.
* \param b the blue component of the pixel in the range 0-255.
* \param a the alpha component of the pixel in the range 0-255.
* \returns a pixel value.
*
* \threadsafety It is safe to call this function from any thread, as long as
*               the palette is not modified.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetPixelFormatDetails
* \sa SDL_GetRGBA
* \sa SDL_MapRGB
* \sa SDL_MapSurfaceRGBA
 */
//go:sdl3extern
var MapRGBA func(format *PixelFormatDetails, palette *Palette, r, g, b, a uint8) uint32

/**
* Get RGB values from a pixel in the specified format.
*
* This function uses the entire 8-bit [0..255] range when converting color
* components from pixel formats with less than 8-bits per RGB component
* (e.g., a completely white pixel in 16-bit RGB565 format would return [0xff,
* 0xff, 0xff] not [0xf8, 0xfc, 0xf8]).
*
* \param pixelvalue a pixel value.
* \param format a pointer to SDL_PixelFormatDetails describing the pixel
*               format.
* \param palette an optional palette for indexed formats, may be NULL.
* \param r a pointer filled in with the red component, may be NULL.
* \param g a pointer filled in with the green component, may be NULL.
* \param b a pointer filled in with the blue component, may be NULL.
*
* \threadsafety It is safe to call this function from any thread, as long as
*               the palette is not modified.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetPixelFormatDetails
* \sa SDL_GetRGBA
* \sa SDL_MapRGB
* \sa SDL_MapRGBA
 */
//go:sdl3extern
var GetRGB func(pixelvalue uint32, format *PixelFormatDetails, palette *Palette, r, g, b *uint8)

/**
* Get RGBA values from a pixel in the specified format.
*
* This function uses the entire 8-bit [0..255] range when converting color
* components from pixel formats with less than 8-bits per RGB component
* (e.g., a completely white pixel in 16-bit RGB565 format would return [0xff,
* 0xff, 0xff] not [0xf8, 0xfc, 0xf8]).
*
* If the surface has no alpha component, the alpha will be returned as 0xff
* (100% opaque).
*
* \param pixelvalue a pixel value.
* \param format a pointer to SDL_PixelFormatDetails describing the pixel
*               format.
* \param palette an optional palette for indexed formats, may be NULL.
* \param r a pointer filled in with the red component, may be NULL.
* \param g a pointer filled in with the green component, may be NULL.
* \param b a pointer filled in with the blue component, may be NULL.
* \param a a pointer filled in with the alpha component, may be NULL.
*
* \threadsafety It is safe to call this function from any thread, as long as
*               the palette is not modified.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetPixelFormatDetails
* \sa SDL_GetRGB
* \sa SDL_MapRGB
* \sa SDL_MapRGBA
 */
//go:sdl3extern
var GetRGBA func(pixelvalue uint32, format *PixelFormatDetails, palette *Palette, r, g, b, a uint8)
