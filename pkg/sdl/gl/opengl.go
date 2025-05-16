//go:build !ios

/* No OpenGL on iOS. */
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

/*
 * This is a simple file to encapsulate the OpenGL API headers.
 *
 * Define NO_SDL_GLEXT if you have your own version of glext.h and want
 * to disable the version included in SDL_opengl.h.
 */
package gl

/*
 * Mesa 3-D graphics library
 *
 * Copyright (C) 1999-2006  Brian Paul   All Rights Reserved.
 * Copyright (C) 2009  VMware, Inc.  All Rights Reserved.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included
 * in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
 * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
 * THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 */

/**********************************************************************
 * Begin system-specific stuff.
 */

const (
	VERSION_MAJOR = 1
	VERSION_MINOR = 3
)

/*
 * Datatypes
 */
type GLenum uint
type GLboolean uint8
type GLbitfield uint
type GLvoid byte
type GLbyte int8      /* 1-byte signed */
type GLshort int16    /* 2-byte signed */
type GLint int32      /* 4-byte signed */
type GLubyte uint8    /* 1-byte unsigned */
type GLushort uint16  /* 2-byte unsigned */
type GLuint uint32    /* 4-byte unsigned */
type GLsizei int32    /* 4-byte signed */
type GLfloat float32  /* single precision float */
type GLclampf float32 /* single precision float in [0,1] */
type GLdouble float64 /* double precision float */
type GLclampd float64 /* double precision float in [0,1] */

/*
 * Constants
 */

/* Boolean values */
const FALSE GLboolean = 0
const TRUE GLboolean = 1

/* Data types */
const BYTE = 0x1400
const UNSIGNED_BYTE = 0x1401
const SHORT = 0x1402
const UNSIGNED_SHORT = 0x1403
const INT = 0x1404
const UNSIGNED_INT = 0x1405
const FLOAT = 0x1406
const A2_BYTES = 0x1407
const A3_BYTES = 0x1408
const A4_BYTES = 0x1409
const DOUBLE = 0x140A

/* Primitives */
const POINTS = 0x0000
const LINES = 0x0001
const LINE_LOOP = 0x0002
const LINE_STRIP = 0x0003
const TRIANGLES = 0x0004
const TRIANGLE_STRIP = 0x0005
const TRIANGLE_FAN = 0x0006
const QUADS = 0x0007
const QUAD_STRIP = 0x0008
const POLYGON = 0x0009

/* Vertex Arrays */
const VERTEX_ARRAY = 0x8074
const NORMAL_ARRAY = 0x8075
const COLOR_ARRAY = 0x8076
const INDEX_ARRAY = 0x8077
const TEXTURE_COORD_ARRAY = 0x8078
const EDGE_FLAG_ARRAY = 0x8079
const VERTEX_ARRAY_SIZE = 0x807A
const VERTEX_ARRAY_TYPE = 0x807B
const VERTEX_ARRAY_STRIDE = 0x807C
const NORMAL_ARRAY_TYPE = 0x807E
const NORMAL_ARRAY_STRIDE = 0x807F
const COLOR_ARRAY_SIZE = 0x8081
const COLOR_ARRAY_TYPE = 0x8082
const COLOR_ARRAY_STRIDE = 0x8083
const INDEX_ARRAY_TYPE = 0x8085
const INDEX_ARRAY_STRIDE = 0x8086
const TEXTURE_COORD_ARRAY_SIZE = 0x8088
const TEXTURE_COORD_ARRAY_TYPE = 0x8089
const TEXTURE_COORD_ARRAY_STRIDE = 0x808A
const EDGE_FLAG_ARRAY_STRIDE = 0x808C
const VERTEX_ARRAY_POINTER = 0x808E
const NORMAL_ARRAY_POINTER = 0x808F
const COLOR_ARRAY_POINTER = 0x8090
const INDEX_ARRAY_POINTER = 0x8091
const TEXTURE_COORD_ARRAY_POINTER = 0x8092
const EDGE_FLAG_ARRAY_POINTER = 0x8093
const V2F = 0x2A20
const V3F = 0x2A21
const C4UB_V2F = 0x2A22
const C4UB_V3F = 0x2A23
const C3F_V3F = 0x2A24
const N3F_V3F = 0x2A25
const C4F_N3F_V3F = 0x2A26
const T2F_V3F = 0x2A27
const T4F_V4F = 0x2A28
const T2F_C4UB_V3F = 0x2A29
const T2F_C3F_V3F = 0x2A2A
const T2F_N3F_V3F = 0x2A2B
const T2F_C4F_N3F_V3F = 0x2A2C
const T4F_C4F_N3F_V4F = 0x2A2D

/* Matrix Mode */
const MATRIX_MODE = 0x0BA0
const MODELVIEW = 0x1700
const PROJECTION = 0x1701
const TEXTURE = 0x1702

/* Points */
const POINT_SMOOTH = 0x0B10
const POINT_SIZE = 0x0B11
const POINT_SIZE_GRANULARITY = 0x0B13
const POINT_SIZE_RANGE = 0x0B12

/* Lines */
const LINE_SMOOTH = 0x0B20
const LINE_STIPPLE = 0x0B24
const LINE_STIPPLE_PATTERN = 0x0B25
const LINE_STIPPLE_REPEAT = 0x0B26
const LINE_WIDTH = 0x0B21
const LINE_WIDTH_GRANULARITY = 0x0B23
const LINE_WIDTH_RANGE = 0x0B22

/* Polygons */
const POINT = 0x1B00
const LINE = 0x1B01
const FILL = 0x1B02
const CW = 0x0900
const CCW = 0x0901
const FRONT = 0x0404
const BACK = 0x0405
const POLYGON_MODE = 0x0B40
const POLYGON_SMOOTH = 0x0B41
const POLYGON_STIPPLE = 0x0B42
const EDGE_FLAG = 0x0B43
const CULL_FACE = 0x0B44
const CULL_FACE_MODE = 0x0B45
const FRONT_FACE = 0x0B46
const POLYGON_OFFSET_FACTOR = 0x8038
const POLYGON_OFFSET_UNITS = 0x2A00
const POLYGON_OFFSET_POINT = 0x2A01
const POLYGON_OFFSET_LINE = 0x2A02
const POLYGON_OFFSET_FILL = 0x8037

/* Display Lists */
const COMPILE = 0x1300
const COMPILE_AND_EXECUTE = 0x1301
const LIST_BASE = 0x0B32
const LIST_INDEX = 0x0B33
const LIST_MODE = 0x0B30

/* Depth buffer */
const NEVER = 0x0200
const LESS = 0x0201
const EQUAL = 0x0202
const LEQUAL = 0x0203
const GREATER = 0x0204
const NOTEQUAL = 0x0205
const GEQUAL = 0x0206
const ALWAYS = 0x0207
const DEPTH_TEST = 0x0B71
const DEPTH_BITS = 0x0D56
const DEPTH_CLEAR_VALUE = 0x0B73
const DEPTH_FUNC = 0x0B74
const DEPTH_RANGE = 0x0B70
const DEPTH_WRITEMASK = 0x0B72
const DEPTH_COMPONENT = 0x1902

/* Lighting */
const LIGHTING = 0x0B50
const LIGHT0 = 0x4000
const LIGHT1 = 0x4001
const LIGHT2 = 0x4002
const LIGHT3 = 0x4003
const LIGHT4 = 0x4004
const LIGHT5 = 0x4005
const LIGHT6 = 0x4006
const LIGHT7 = 0x4007
const SPOT_EXPONENT = 0x1205
const SPOT_CUTOFF = 0x1206
const CONSTANT_ATTENUATION = 0x1207
const LINEAR_ATTENUATION = 0x1208
const QUADRATIC_ATTENUATION = 0x1209
const AMBIENT = 0x1200
const DIFFUSE = 0x1201
const SPECULAR = 0x1202
const SHININESS = 0x1601
const EMISSION = 0x1600
const POSITION = 0x1203
const SPOT_DIRECTION = 0x1204
const AMBIENT_AND_DIFFUSE = 0x1602
const COLOR_INDEXES = 0x1603
const LIGHT_MODEL_TWO_SIDE = 0x0B52
const LIGHT_MODEL_LOCAL_VIEWER = 0x0B51
const LIGHT_MODEL_AMBIENT = 0x0B53
const FRONT_AND_BACK = 0x0408
const SHADE_MODEL = 0x0B54
const FLAT = 0x1D00
const SMOOTH = 0x1D01
const COLOR_MATERIAL = 0x0B57
const COLOR_MATERIAL_FACE = 0x0B55
const COLOR_MATERIAL_PARAMETER = 0x0B56
const NORMALIZE = 0x0BA1

/* User clipping planes */
const CLIP_PLANE0 = 0x3000
const CLIP_PLANE1 = 0x3001
const CLIP_PLANE2 = 0x3002
const CLIP_PLANE3 = 0x3003
const CLIP_PLANE4 = 0x3004
const CLIP_PLANE5 = 0x3005

/* Accumulation buffer */
const ACCUM_RED_BITS = 0x0D58
const ACCUM_GREEN_BITS = 0x0D59
const ACCUM_BLUE_BITS = 0x0D5A
const ACCUM_ALPHA_BITS = 0x0D5B
const ACCUM_CLEAR_VALUE = 0x0B80
const ACCUM = 0x0100
const ADD = 0x0104
const LOAD = 0x0101
const MULT = 0x0103
const RETURN = 0x0102

/* Alpha testing */
const ALPHA_TEST = 0x0BC0
const ALPHA_TEST_REF = 0x0BC2
const ALPHA_TEST_FUNC = 0x0BC1

/* Blending */
const BLEND = 0x0BE2
const BLEND_SRC = 0x0BE1
const BLEND_DST = 0x0BE0
const ZERO = 0
const ONE = 1
const SRC_COLOR = 0x0300
const ONE_MINUS_SRC_COLOR = 0x0301
const SRC_ALPHA = 0x0302
const ONE_MINUS_SRC_ALPHA = 0x0303
const DST_ALPHA = 0x0304
const ONE_MINUS_DST_ALPHA = 0x0305
const DST_COLOR = 0x0306
const ONE_MINUS_DST_COLOR = 0x0307
const SRC_ALPHA_SATURATE = 0x0308

/* Render Mode */
const FEEDBACK = 0x1C01
const RENDER = 0x1C00
const SELECT = 0x1C02

/* Feedback */
const A2D = 0x0600
const A3D = 0x0601
const A3D_COLOR = 0x0602
const A3D_COLOR_TEXTURE = 0x0603
const A4D_COLOR_TEXTURE = 0x0604
const POINT_TOKEN = 0x0701
const LINE_TOKEN = 0x0702
const LINE_RESET_TOKEN = 0x0707
const POLYGON_TOKEN = 0x0703
const BITMAP_TOKEN = 0x0704
const DRAW_PIXEL_TOKEN = 0x0705
const COPY_PIXEL_TOKEN = 0x0706
const PASS_THROUGH_TOKEN = 0x0700
const FEEDBACK_BUFFER_POINTER = 0x0DF0
const FEEDBACK_BUFFER_SIZE = 0x0DF1
const FEEDBACK_BUFFER_TYPE = 0x0DF2

/* Selection */
const SELECTION_BUFFER_POINTER = 0x0DF3
const SELECTION_BUFFER_SIZE = 0x0DF4

/* Fog */
const FOG = 0x0B60
const FOG_MODE = 0x0B65
const FOG_DENSITY = 0x0B62
const FOG_COLOR = 0x0B66
const FOG_INDEX = 0x0B61
const FOG_START = 0x0B63
const FOG_END = 0x0B64
const LINEAR = 0x2601
const EXP = 0x0800
const EXP2 = 0x0801

/* Logic Ops */
const LOGIC_OP = 0x0BF1
const INDEX_LOGIC_OP = 0x0BF1
const COLOR_LOGIC_OP = 0x0BF2
const LOGIC_OP_MODE = 0x0BF0
const CLEAR = 0x1500
const SET = 0x150F
const COPY = 0x1503
const COPY_INVERTED = 0x150C
const NOOP = 0x1505
const INVERT = 0x150A
const AND = 0x1501
const NAND = 0x150E
const OR = 0x1507
const NOR = 0x1508
const XOR = 0x1506
const EQUIV = 0x1509
const AND_REVERSE = 0x1502
const AND_INVERTED = 0x1504
const OR_REVERSE = 0x150B
const OR_INVERTED = 0x150D

/* Stencil */
const STENCIL_BITS = 0x0D57
const STENCIL_TEST = 0x0B90
const STENCIL_CLEAR_VALUE = 0x0B91
const STENCIL_FUNC = 0x0B92
const STENCIL_VALUE_MASK = 0x0B93
const STENCIL_FAIL = 0x0B94
const STENCIL_PASS_DEPTH_FAIL = 0x0B95
const STENCIL_PASS_DEPTH_PASS = 0x0B96
const STENCIL_REF = 0x0B97
const STENCIL_WRITEMASK = 0x0B98
const STENCIL_INDEX = 0x1901
const KEEP = 0x1E00
const REPLACE = 0x1E01
const INCR = 0x1E02
const DECR = 0x1E03

/* Buffers, Pixel Drawing/Reading */
const NONE = 0
const LEFT = 0x0406
const RIGHT = 0x0407

/*GL_FRONT					0x0404 */
/*GL_BACK					0x0405 */
/*GL_FRONT_AND_BACK				0x0408 */
const FRONT_LEFT = 0x0400
const FRONT_RIGHT = 0x0401
const BACK_LEFT = 0x0402
const BACK_RIGHT = 0x0403
const AUX0 = 0x0409
const AUX1 = 0x040A
const AUX2 = 0x040B
const AUX3 = 0x040C
const COLOR_INDEX = 0x1900
const RED = 0x1903
const GREEN = 0x1904
const BLUE = 0x1905
const ALPHA = 0x1906
const LUMINANCE = 0x1909
const LUMINANCE_ALPHA = 0x190A
const ALPHA_BITS = 0x0D55
const RED_BITS = 0x0D52
const GREEN_BITS = 0x0D53
const BLUE_BITS = 0x0D54
const INDEX_BITS = 0x0D51
const SUBPIXEL_BITS = 0x0D50
const AUX_BUFFERS = 0x0C00
const READ_BUFFER = 0x0C02
const DRAW_BUFFER = 0x0C01
const DOUBLEBUFFER = 0x0C32
const STEREO = 0x0C33
const BITMAP = 0x1A00
const COLOR = 0x1800
const DEPTH = 0x1801
const STENCIL = 0x1802
const DITHER = 0x0BD0
const RGB = 0x1907
const RGBA = 0x1908

/* Implementation limits */
const MAX_LIST_NESTING = 0x0B31
const MAX_EVAL_ORDER = 0x0D30
const MAX_LIGHTS = 0x0D31
const MAX_CLIP_PLANES = 0x0D32
const MAX_TEXTURE_SIZE = 0x0D33
const MAX_PIXEL_MAP_TABLE = 0x0D34
const MAX_ATTRIB_STACK_DEPTH = 0x0D35
const MAX_MODELVIEW_STACK_DEPTH = 0x0D36
const MAX_NAME_STACK_DEPTH = 0x0D37
const MAX_PROJECTION_STACK_DEPTH = 0x0D38
const MAX_TEXTURE_STACK_DEPTH = 0x0D39
const MAX_VIEWPORT_DIMS = 0x0D3A
const MAX_CLIENT_ATTRIB_STACK_DEPTH = 0x0D3B

/* Gets */
const ATTRIB_STACK_DEPTH = 0x0BB0
const CLIENT_ATTRIB_STACK_DEPTH = 0x0BB1
const COLOR_CLEAR_VALUE = 0x0C22
const COLOR_WRITEMASK = 0x0C23
const CURRENT_INDEX = 0x0B01
const CURRENT_COLOR = 0x0B00
const CURRENT_NORMAL = 0x0B02
const CURRENT_RASTER_COLOR = 0x0B04
const CURRENT_RASTER_DISTANCE = 0x0B09
const CURRENT_RASTER_INDEX = 0x0B05
const CURRENT_RASTER_POSITION = 0x0B07
const CURRENT_RASTER_TEXTURE_COORDS = 0x0B06
const CURRENT_RASTER_POSITION_VALID = 0x0B08
const CURRENT_TEXTURE_COORDS = 0x0B03
const INDEX_CLEAR_VALUE = 0x0C20
const INDEX_MODE = 0x0C30
const INDEX_WRITEMASK = 0x0C21
const MODELVIEW_MATRIX = 0x0BA6
const MODELVIEW_STACK_DEPTH = 0x0BA3
const NAME_STACK_DEPTH = 0x0D70
const PROJECTION_MATRIX = 0x0BA7
const PROJECTION_STACK_DEPTH = 0x0BA4
const RENDER_MODE = 0x0C40
const RGBA_MODE = 0x0C31
const TEXTURE_MATRIX = 0x0BA8
const TEXTURE_STACK_DEPTH = 0x0BA5
const VIEWPORT = 0x0BA2

/* Evaluators */
const AUTO_NORMAL = 0x0D80
const MAP1_COLOR_4 = 0x0D90
const MAP1_INDEX = 0x0D91
const MAP1_NORMAL = 0x0D92
const MAP1_TEXTURE_COORD_1 = 0x0D93
const MAP1_TEXTURE_COORD_2 = 0x0D94
const MAP1_TEXTURE_COORD_3 = 0x0D95
const MAP1_TEXTURE_COORD_4 = 0x0D96
const MAP1_VERTEX_3 = 0x0D97
const MAP1_VERTEX_4 = 0x0D98
const MAP2_COLOR_4 = 0x0DB0
const MAP2_INDEX = 0x0DB1
const MAP2_NORMAL = 0x0DB2
const MAP2_TEXTURE_COORD_1 = 0x0DB3
const MAP2_TEXTURE_COORD_2 = 0x0DB4
const MAP2_TEXTURE_COORD_3 = 0x0DB5
const MAP2_TEXTURE_COORD_4 = 0x0DB6
const MAP2_VERTEX_3 = 0x0DB7
const MAP2_VERTEX_4 = 0x0DB8
const MAP1_GRID_DOMAIN = 0x0DD0
const MAP1_GRID_SEGMENTS = 0x0DD1
const MAP2_GRID_DOMAIN = 0x0DD2
const MAP2_GRID_SEGMENTS = 0x0DD3
const COEFF = 0x0A00
const ORDER = 0x0A01
const DOMAIN = 0x0A02

/* Hints */
const PERSPECTIVE_CORRECTION_HINT = 0x0C50
const POINT_SMOOTH_HINT = 0x0C51
const LINE_SMOOTH_HINT = 0x0C52
const POLYGON_SMOOTH_HINT = 0x0C53
const FOG_HINT = 0x0C54
const DONT_CARE = 0x1100
const FASTEST = 0x1101
const NICEST = 0x1102

/* Scissor box */
const SCISSOR_BOX = 0x0C10
const SCISSOR_TEST = 0x0C11

/* Pixel Mode / Transfer */
const MAP_COLOR = 0x0D10
const MAP_STENCIL = 0x0D11
const INDEX_SHIFT = 0x0D12
const INDEX_OFFSET = 0x0D13
const RED_SCALE = 0x0D14
const RED_BIAS = 0x0D15
const GREEN_SCALE = 0x0D18
const GREEN_BIAS = 0x0D19
const BLUE_SCALE = 0x0D1A
const BLUE_BIAS = 0x0D1B
const ALPHA_SCALE = 0x0D1C
const ALPHA_BIAS = 0x0D1D
const DEPTH_SCALE = 0x0D1E
const DEPTH_BIAS = 0x0D1F
const PIXEL_MAP_S_TO_S_SIZE = 0x0CB1
const PIXEL_MAP_I_TO_I_SIZE = 0x0CB0
const PIXEL_MAP_I_TO_R_SIZE = 0x0CB2
const PIXEL_MAP_I_TO_G_SIZE = 0x0CB3
const PIXEL_MAP_I_TO_B_SIZE = 0x0CB4
const PIXEL_MAP_I_TO_A_SIZE = 0x0CB5
const PIXEL_MAP_R_TO_R_SIZE = 0x0CB6
const PIXEL_MAP_G_TO_G_SIZE = 0x0CB7
const PIXEL_MAP_B_TO_B_SIZE = 0x0CB8
const PIXEL_MAP_A_TO_A_SIZE = 0x0CB9
const PIXEL_MAP_S_TO_S = 0x0C71
const PIXEL_MAP_I_TO_I = 0x0C70
const PIXEL_MAP_I_TO_R = 0x0C72
const PIXEL_MAP_I_TO_G = 0x0C73
const PIXEL_MAP_I_TO_B = 0x0C74
const PIXEL_MAP_I_TO_A = 0x0C75
const PIXEL_MAP_R_TO_R = 0x0C76
const PIXEL_MAP_G_TO_G = 0x0C77
const PIXEL_MAP_B_TO_B = 0x0C78
const PIXEL_MAP_A_TO_A = 0x0C79
const PACK_ALIGNMENT = 0x0D05
const PACK_LSB_FIRST = 0x0D01
const PACK_ROW_LENGTH = 0x0D02
const PACK_SKIP_PIXELS = 0x0D04
const PACK_SKIP_ROWS = 0x0D03
const PACK_SWAP_BYTES = 0x0D00
const UNPACK_ALIGNMENT = 0x0CF5
const UNPACK_LSB_FIRST = 0x0CF1
const UNPACK_ROW_LENGTH = 0x0CF2
const UNPACK_SKIP_PIXELS = 0x0CF4
const UNPACK_SKIP_ROWS = 0x0CF3
const UNPACK_SWAP_BYTES = 0x0CF0
const ZOOM_X = 0x0D16
const ZOOM_Y = 0x0D17

/* Texture mapping */
const TEXTURE_ENV = 0x2300
const TEXTURE_ENV_MODE = 0x2200
const TEXTURE_1D = 0x0DE0
const TEXTURE_2D = 0x0DE1
const TEXTURE_WRAP_S = 0x2802
const TEXTURE_WRAP_T = 0x2803
const TEXTURE_MAG_FILTER = 0x2800
const TEXTURE_MIN_FILTER = 0x2801
const TEXTURE_ENV_COLOR = 0x2201
const TEXTURE_GEN_S = 0x0C60
const TEXTURE_GEN_T = 0x0C61
const TEXTURE_GEN_R = 0x0C62
const TEXTURE_GEN_Q = 0x0C63
const TEXTURE_GEN_MODE = 0x2500
const TEXTURE_BORDER_COLOR = 0x1004
const TEXTURE_WIDTH = 0x1000
const TEXTURE_HEIGHT = 0x1001
const TEXTURE_BORDER = 0x1005
const TEXTURE_COMPONENTS = 0x1003
const TEXTURE_RED_SIZE = 0x805C
const TEXTURE_GREEN_SIZE = 0x805D
const TEXTURE_BLUE_SIZE = 0x805E
const TEXTURE_ALPHA_SIZE = 0x805F
const TEXTURE_LUMINANCE_SIZE = 0x8060
const TEXTURE_INTENSITY_SIZE = 0x8061
const NEAREST_MIPMAP_NEAREST = 0x2700
const NEAREST_MIPMAP_LINEAR = 0x2702
const LINEAR_MIPMAP_NEAREST = 0x2701
const LINEAR_MIPMAP_LINEAR = 0x2703
const OBJECT_LINEAR = 0x2401
const OBJECT_PLANE = 0x2501
const EYE_LINEAR = 0x2400
const EYE_PLANE = 0x2502
const SPHERE_MAP = 0x2402
const DECAL = 0x2101
const MODULATE = 0x2100
const NEAREST = 0x2600
const REPEAT = 0x2901
const CLAMP = 0x2900
const S = 0x2000
const T = 0x2001
const R = 0x2002
const Q = 0x2003

/* Utility */
const VENDOR = 0x1F00
const RENDERER = 0x1F01
const VERSION = 0x1F02
const EXTENSIONS = 0x1F03

/* Errors */
const NO_ERROR = 0
const INVALID_ENUM = 0x0500
const INVALID_VALUE = 0x0501
const INVALID_OPERATION = 0x0502
const STACK_OVERFLOW = 0x0503
const STACK_UNDERFLOW = 0x0504
const OUT_OF_MEMORY = 0x0505

/* glPush/PopAttrib bits */
const CURRENT_BIT = 0x00000001
const POINT_BIT = 0x00000002
const LINE_BIT = 0x00000004
const POLYGON_BIT = 0x00000008
const POLYGON_STIPPLE_BIT = 0x00000010
const PIXEL_MODE_BIT = 0x00000020
const LIGHTING_BIT = 0x00000040
const FOG_BIT = 0x00000080
const DEPTH_BUFFER_BIT = 0x00000100
const ACCUM_BUFFER_BIT = 0x00000200
const STENCIL_BUFFER_BIT = 0x00000400
const VIEWPORT_BIT = 0x00000800
const TRANSFORM_BIT = 0x00001000
const ENABLE_BIT = 0x00002000
const COLOR_BUFFER_BIT = 0x00004000
const HINT_BIT = 0x00008000
const EVAL_BIT = 0x00010000
const LIST_BIT = 0x00020000
const TEXTURE_BIT = 0x00040000
const SCISSOR_BIT = 0x00080000
const ALL_ATTRIB_BITS = 0x000FFFFF

/* OpenGL 1.1 */
const PROXY_TEXTURE_1D = 0x8063
const PROXY_TEXTURE_2D = 0x8064
const TEXTURE_PRIORITY = 0x8066
const TEXTURE_RESIDENT = 0x8067
const TEXTURE_BINDING_1D = 0x8068
const TEXTURE_BINDING_2D = 0x8069
const TEXTURE_INTERNAL_FORMAT = 0x1003
const ALPHA4 = 0x803B
const ALPHA8 = 0x803C
const ALPHA12 = 0x803D
const ALPHA16 = 0x803E
const LUMINANCE4 = 0x803F
const LUMINANCE8 = 0x8040
const LUMINANCE12 = 0x8041
const LUMINANCE16 = 0x8042
const LUMINANCE4_ALPHA4 = 0x8043
const LUMINANCE6_ALPHA2 = 0x8044
const LUMINANCE8_ALPHA8 = 0x8045
const LUMINANCE12_ALPHA4 = 0x8046
const LUMINANCE12_ALPHA12 = 0x8047
const LUMINANCE16_ALPHA16 = 0x8048
const INTENSITY = 0x8049
const INTENSITY4 = 0x804A
const INTENSITY8 = 0x804B
const INTENSITY12 = 0x804C
const INTENSITY16 = 0x804D
const R3_G3_B2 = 0x2A10
const RGB4 = 0x804F
const RGB5 = 0x8050
const RGB8 = 0x8051
const RGB10 = 0x8052
const RGB12 = 0x8053
const RGB16 = 0x8054
const RGBA2 = 0x8055
const RGBA4 = 0x8056
const RGB5_A1 = 0x8057
const RGBA8 = 0x8058
const RGB10_A2 = 0x8059
const RGBA12 = 0x805A
const RGBA16 = 0x805B
const CLIENT_PIXEL_STORE_BIT = 0x00000001
const CLIENT_VERTEX_ARRAY_BIT = 0x00000002
const ALL_CLIENT_ATTRIB_BITS = 0xFFFFFFFF
const CLIENT_ALL_ATTRIB_BITS = 0xFFFFFFFF

/*
 * OpenGL 1.2
 */

const RESCALE_NORMAL = 0x803A
const CLAMP_TO_EDGE = 0x812F
const MAX_ELEMENTS_VERTICES = 0x80E8
const MAX_ELEMENTS_INDICES = 0x80E9
const BGR = 0x80E0
const BGRA = 0x80E1
const UNSIGNED_BYTE_3_3_2 = 0x8032
const UNSIGNED_BYTE_2_3_3_REV = 0x8362
const UNSIGNED_SHORT_5_6_5 = 0x8363
const UNSIGNED_SHORT_5_6_5_REV = 0x8364
const UNSIGNED_SHORT_4_4_4_4 = 0x8033
const UNSIGNED_SHORT_4_4_4_4_REV = 0x8365
const UNSIGNED_SHORT_5_5_5_1 = 0x8034
const UNSIGNED_SHORT_1_5_5_5_REV = 0x8366
const UNSIGNED_INT_8_8_8_8 = 0x8035
const UNSIGNED_INT_8_8_8_8_REV = 0x8367
const UNSIGNED_INT_10_10_10_2 = 0x8036
const UNSIGNED_INT_2_10_10_10_REV = 0x8368
const LIGHT_MODEL_COLOR_CONTROL = 0x81F8
const SINGLE_COLOR = 0x81F9
const SEPARATE_SPECULAR_COLOR = 0x81FA
const TEXTURE_MIN_LOD = 0x813A
const TEXTURE_MAX_LOD = 0x813B
const TEXTURE_BASE_LEVEL = 0x813C
const TEXTURE_MAX_LEVEL = 0x813D
const SMOOTH_POINT_SIZE_RANGE = 0x0B12
const SMOOTH_POINT_SIZE_GRANULARITY = 0x0B13
const SMOOTH_LINE_WIDTH_RANGE = 0x0B22
const SMOOTH_LINE_WIDTH_GRANULARITY = 0x0B23
const ALIASED_POINT_SIZE_RANGE = 0x846D
const ALIASED_LINE_WIDTH_RANGE = 0x846E
const PACK_SKIP_IMAGES = 0x806B
const PACK_IMAGE_HEIGHT = 0x806C
const UNPACK_SKIP_IMAGES = 0x806D
const UNPACK_IMAGE_HEIGHT = 0x806E
const TEXTURE_3D = 0x806F
const PROXY_TEXTURE_3D = 0x8070
const TEXTURE_DEPTH = 0x8071
const TEXTURE_WRAP_R = 0x8072
const MAX_3D_TEXTURE_SIZE = 0x8073
const TEXTURE_BINDING_3D = 0x806A

/*
 * GL_ARB_imaging
 */

const CONSTANT_COLOR = 0x8001
const ONE_MINUS_CONSTANT_COLOR = 0x8002
const CONSTANT_ALPHA = 0x8003
const ONE_MINUS_CONSTANT_ALPHA = 0x8004
const COLOR_TABLE = 0x80D0
const POST_CONVOLUTION_COLOR_TABLE = 0x80D1
const POST_COLOR_MATRIX_COLOR_TABLE = 0x80D2
const PROXY_COLOR_TABLE = 0x80D3
const PROXY_POST_CONVOLUTION_COLOR_TABLE = 0x80D4
const PROXY_POST_COLOR_MATRIX_COLOR_TABLE = 0x80D5
const COLOR_TABLE_SCALE = 0x80D6
const COLOR_TABLE_BIAS = 0x80D7
const COLOR_TABLE_FORMAT = 0x80D8
const COLOR_TABLE_WIDTH = 0x80D9
const COLOR_TABLE_RED_SIZE = 0x80DA
const COLOR_TABLE_GREEN_SIZE = 0x80DB
const COLOR_TABLE_BLUE_SIZE = 0x80DC
const COLOR_TABLE_ALPHA_SIZE = 0x80DD
const COLOR_TABLE_LUMINANCE_SIZE = 0x80DE
const COLOR_TABLE_INTENSITY_SIZE = 0x80DF
const CONVOLUTION_1D = 0x8010
const CONVOLUTION_2D = 0x8011
const SEPARABLE_2D = 0x8012
const CONVOLUTION_BORDER_MODE = 0x8013
const CONVOLUTION_FILTER_SCALE = 0x8014
const CONVOLUTION_FILTER_BIAS = 0x8015
const REDUCE = 0x8016
const CONVOLUTION_FORMAT = 0x8017
const CONVOLUTION_WIDTH = 0x8018
const CONVOLUTION_HEIGHT = 0x8019
const MAX_CONVOLUTION_WIDTH = 0x801A
const MAX_CONVOLUTION_HEIGHT = 0x801B
const POST_CONVOLUTION_RED_SCALE = 0x801C
const POST_CONVOLUTION_GREEN_SCALE = 0x801D
const POST_CONVOLUTION_BLUE_SCALE = 0x801E
const POST_CONVOLUTION_ALPHA_SCALE = 0x801F
const POST_CONVOLUTION_RED_BIAS = 0x8020
const POST_CONVOLUTION_GREEN_BIAS = 0x8021
const POST_CONVOLUTION_BLUE_BIAS = 0x8022
const POST_CONVOLUTION_ALPHA_BIAS = 0x8023
const CONSTANT_BORDER = 0x8151
const REPLICATE_BORDER = 0x8153
const CONVOLUTION_BORDER_COLOR = 0x8154
const COLOR_MATRIX = 0x80B1
const COLOR_MATRIX_STACK_DEPTH = 0x80B2
const MAX_COLOR_MATRIX_STACK_DEPTH = 0x80B3
const POST_COLOR_MATRIX_RED_SCALE = 0x80B4
const POST_COLOR_MATRIX_GREEN_SCALE = 0x80B5
const POST_COLOR_MATRIX_BLUE_SCALE = 0x80B6
const POST_COLOR_MATRIX_ALPHA_SCALE = 0x80B7
const POST_COLOR_MATRIX_RED_BIAS = 0x80B8
const POST_COLOR_MATRIX_GREEN_BIAS = 0x80B9
const POST_COLOR_MATRIX_BLUE_BIAS = 0x80BA
const POST_COLOR_MATRIX_ALPHA_BIAS = 0x80BB
const HISTOGRAM = 0x8024
const PROXY_HISTOGRAM = 0x8025
const HISTOGRAM_WIDTH = 0x8026
const HISTOGRAM_FORMAT = 0x8027
const HISTOGRAM_RED_SIZE = 0x8028
const HISTOGRAM_GREEN_SIZE = 0x8029
const HISTOGRAM_BLUE_SIZE = 0x802A
const HISTOGRAM_ALPHA_SIZE = 0x802B
const HISTOGRAM_LUMINANCE_SIZE = 0x802C
const HISTOGRAM_SINK = 0x802D
const MINMAX = 0x802E
const MINMAX_FORMAT = 0x802F
const MINMAX_SINK = 0x8030
const TABLE_TOO_LARGE = 0x8031
const BLEND_EQUATION = 0x8009
const MIN = 0x8007
const MAX = 0x8008
const FUNC_ADD = 0x8006
const FUNC_SUBTRACT = 0x800A
const FUNC_REVERSE_SUBTRACT = 0x800B
const BLEND_COLOR = 0x8005

//go:sdl3staticextern
var glColorTable func(target GLenum, internalformat GLenum,
	width GLsizei, format GLenum,
	type_ GLenum, table *GLvoid)

//go:sdl3staticextern
var glColorSubTable func(target GLenum,
	start GLsizei, count GLsizei,
	format GLenum, type_ GLenum,
	data *GLvoid)

//go:sdl3staticextern
var glColorTableParameteriv func(target GLenum, pname GLenum,
	params *GLint)

//go:sdl3staticextern
var glColorTableParameterfv func(target GLenum, pname GLenum,
	params *GLfloat)

//go:sdl3staticextern
var glCopyColorSubTable func(target GLenum, start GLsizei,
	x GLint, y GLint, width GLsizei)

//go:sdl3staticextern
var glCopyColorTable func(target GLenum, internalformat GLenum,
	x GLint, y GLint, width GLsizei)

//go:sdl3staticextern
var glGetColorTable func(target GLenum, format GLenum,
	type_ GLenum, table *GLvoid)

//go:sdl3staticextern
var glGetColorTableParameterfv func(target GLenum, pname GLenum,
	params *GLfloat)

//go:sdl3staticextern
var glGetColorTableParameteriv func(target GLenum, pname GLenum,
	params *GLint)

//go:sdl3staticextern
var glBlendEquation func(mode GLenum)

//go:sdl3staticextern
var glBlendColor func(red GLclampf, green GLclampf,
	blue GLclampf, alpha GLclampf)

//go:sdl3staticextern
var glHistogram func(target GLenum, width GLsizei,
	internalformat GLenum, sink GLboolean)

//go:sdl3staticextern
var glResetHistogram func(target GLenum)

//go:sdl3staticextern
var glGetHistogram func(target GLenum, reset GLboolean,
	format GLenum, type_ GLenum,
	values *GLvoid)

//go:sdl3staticextern
var glGetHistogramParameterfv func(target GLenum, pname GLenum,
	params *GLfloat)

//go:sdl3staticextern
var glGetHistogramParameteriv func(target GLenum, pname GLenum,
	params *GLint)

//go:sdl3staticextern
var glMinmax func(target GLenum, internalformat GLenum,
	sink GLboolean)

//go:sdl3staticextern
var glResetMinmax func(target GLenum)

//go:sdl3staticextern
var glGetMinmax func(target GLenum, reset GLboolean,
	format GLenum, types GLenum,
	values *GLvoid)

//go:sdl3staticextern
var glGetMinmaxParameterfv func(target GLenum, pname GLenum,
	params *GLfloat)

//go:sdl3staticextern
var glGetMinmaxParameteriv func(target GLenum, pname GLenum,
	params *GLint)

//go:sdl3staticextern
var glConvolutionFilter1D func(target GLenum,
	internalformat GLenum, width GLsizei, format GLenum, type_ GLenum,
	image *GLvoid)

//go:sdl3staticextern
var glConvolutionFilter2D func(target GLenum,
	internalformat GLenum, width GLsizei, height GLsizei, format GLenum,
	type_ GLenum, image *GLvoid)

//go:sdl3staticextern
var glConvolutionParameterf func(target GLenum, pname GLenum,
	params GLfloat)

//go:sdl3staticextern
var glConvolutionParameterfv func(target GLenum, pname GLenum,
	params *GLfloat)

//go:sdl3staticextern
var glConvolutionParameteri func(target GLenum, pname GLenum,
	params GLint)

//go:sdl3staticextern
var glConvolutionParameteriv func(target GLenum, pname GLenum,
	params *GLint)

//go:sdl3staticextern
var glCopyConvolutionFilter1D func(target GLenum,
	internalformat GLenum, x GLint, y GLint, width GLsizei)

//go:sdl3staticextern
var glCopyConvolutionFilter2D func(target GLenum,
	internalformat GLenum, x GLint, y GLint, width GLsizei,
	height GLsizei)

//go:sdl3staticextern
var glGetConvolutionFilter func(target GLenum, format GLenum,
	type_ GLenum, image *GLvoid)

//go:sdl3staticextern
var glGetConvolutionParameterfv func(target GLenum, pname GLenum,
	params *GLfloat)

//go:sdl3staticextern
var glGetConvolutionParameteriv func(target GLenum, pname GLenum,
	params *GLint)

//go:sdl3staticextern
var glSeparableFilter2D func(target GLenum,
	internalformat GLenum, width GLsizei, height GLsizei, format GLenum,
	type_ GLenum, row *GLvoid, column *GLvoid)

//go:sdl3staticextern
var glGetSeparableFilter func(target GLenum, format GLenum,
	type_ GLenum, row *GLvoid, column *GLvoid, span *GLvoid)

/*
 * OpenGL 1.3
 */

/* multitexture */
const TEXTURE0 = 0x84C0
const TEXTURE1 = 0x84C1
const TEXTURE2 = 0x84C2
const TEXTURE3 = 0x84C3
const TEXTURE4 = 0x84C4
const TEXTURE5 = 0x84C5
const TEXTURE6 = 0x84C6
const TEXTURE7 = 0x84C7
const TEXTURE8 = 0x84C8
const TEXTURE9 = 0x84C9
const TEXTURE10 = 0x84CA
const TEXTURE11 = 0x84CB
const TEXTURE12 = 0x84CC
const TEXTURE13 = 0x84CD
const TEXTURE14 = 0x84CE
const TEXTURE15 = 0x84CF
const TEXTURE16 = 0x84D0
const TEXTURE17 = 0x84D1
const TEXTURE18 = 0x84D2
const TEXTURE19 = 0x84D3
const TEXTURE20 = 0x84D4
const TEXTURE21 = 0x84D5
const TEXTURE22 = 0x84D6
const TEXTURE23 = 0x84D7
const TEXTURE24 = 0x84D8
const TEXTURE25 = 0x84D9
const TEXTURE26 = 0x84DA
const TEXTURE27 = 0x84DB
const TEXTURE28 = 0x84DC
const TEXTURE29 = 0x84DD
const TEXTURE30 = 0x84DE
const TEXTURE31 = 0x84DF
const ACTIVE_TEXTURE = 0x84E0
const CLIENT_ACTIVE_TEXTURE = 0x84E1
const MAX_TEXTURE_UNITS = 0x84E2

/* texture_cube_map */
const NORMAL_MAP = 0x8511
const REFLECTION_MAP = 0x8512
const TEXTURE_CUBE_MAP = 0x8513
const TEXTURE_BINDING_CUBE_MAP = 0x8514
const TEXTURE_CUBE_MAP_POSITIVE_X = 0x8515
const TEXTURE_CUBE_MAP_NEGATIVE_X = 0x8516
const TEXTURE_CUBE_MAP_POSITIVE_Y = 0x8517
const TEXTURE_CUBE_MAP_NEGATIVE_Y = 0x8518
const TEXTURE_CUBE_MAP_POSITIVE_Z = 0x8519
const TEXTURE_CUBE_MAP_NEGATIVE_Z = 0x851A
const PROXY_TEXTURE_CUBE_MAP = 0x851B
const MAX_CUBE_MAP_TEXTURE_SIZE = 0x851C

/* texture_compression */
const COMPRESSED_ALPHA = 0x84E9
const COMPRESSED_LUMINANCE = 0x84EA
const COMPRESSED_LUMINANCE_ALPHA = 0x84EB
const COMPRESSED_INTENSITY = 0x84EC
const COMPRESSED_RGB = 0x84ED
const COMPRESSED_RGBA = 0x84EE
const TEXTURE_COMPRESSION_HINT = 0x84EF
const TEXTURE_COMPRESSED_IMAGE_SIZE = 0x86A0
const TEXTURE_COMPRESSED = 0x86A1
const NUM_COMPRESSED_TEXTURE_FORMATS = 0x86A2
const COMPRESSED_TEXTURE_FORMATS = 0x86A3

/* multisample */
const MULTISAMPLE = 0x809D
const SAMPLE_ALPHA_TO_COVERAGE = 0x809E
const SAMPLE_ALPHA_TO_ONE = 0x809F
const SAMPLE_COVERAGE = 0x80A0
const SAMPLE_BUFFERS = 0x80A8
const SAMPLES = 0x80A9
const SAMPLE_COVERAGE_VALUE = 0x80AA
const SAMPLE_COVERAGE_INVERT = 0x80AB
const MULTISAMPLE_BIT = 0x20000000

/* transpose_matrix */
const TRANSPOSE_MODELVIEW_MATRIX = 0x84E3
const TRANSPOSE_PROJECTION_MATRIX = 0x84E4
const TRANSPOSE_TEXTURE_MATRIX = 0x84E5
const TRANSPOSE_COLOR_MATRIX = 0x84E6

/* texture_env_combine */
const COMBINE = 0x8570
const COMBINE_RGB = 0x8571
const COMBINE_ALPHA = 0x8572
const SOURCE0_RGB = 0x8580
const SOURCE1_RGB = 0x8581
const SOURCE2_RGB = 0x8582
const SOURCE0_ALPHA = 0x8588
const SOURCE1_ALPHA = 0x8589
const SOURCE2_ALPHA = 0x858A
const OPERAND0_RGB = 0x8590
const OPERAND1_RGB = 0x8591
const OPERAND2_RGB = 0x8592
const OPERAND0_ALPHA = 0x8598
const OPERAND1_ALPHA = 0x8599
const OPERAND2_ALPHA = 0x859A
const RGB_SCALE = 0x8573
const ADD_SIGNED = 0x8574
const INTERPOLATE = 0x8575
const SUBTRACT = 0x84E7
const CONSTANT = 0x8576
const PRIMARY_COLOR = 0x8577
const PREVIOUS = 0x8578

/* texture_env_dot3 */
const DOT3_RGB = 0x86AE
const DOT3_RGBA = 0x86AF

/* texture_border_clamp */
const CLAMP_TO_BORDER = 0x812D

//go:sdl3staticextern
var glActiveTexture func(texture GLenum)

//go:sdl3staticextern
var glClientActiveTexture func(texture GLenum)

//go:sdl3staticextern
var glCompressedTexImage1D func(target GLenum, level GLint, internalformat GLenum, width GLsizei, border GLint, imageSize GLsizei, data *GLvoid)

//go:sdl3staticextern
var glCompressedTexImage2D func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, imageSize GLsizei, data *GLvoid)

//go:sdl3staticextern
var glCompressedTexImage3D func(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, depth GLsizei, border GLint, imageSize GLsizei, data *GLvoid)

//go:sdl3staticextern
var glCompressedTexSubImage1D func(target GLenum, level GLint, xoffset GLint, width GLsizei, format GLenum, imageSize GLsizei, data *GLvoid)

//go:sdl3staticextern
var glCompressedTexSubImage2D func(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, imageSize GLsizei, data *GLvoid)

//go:sdl3staticextern
var glCompressedTexSubImage3D func(target GLenum, level GLint, xoffset GLint, yoffset GLint, zoffset GLint, width GLsizei, height GLsizei, depth GLsizei, format GLenum, imageSize GLsizei, data *GLvoid)

//go:sdl3staticextern
var glGetCompressedTexImage func(target GLenum, lod GLint, img *GLvoid)

//go:sdl3staticextern
var glMultiTexCoord1d func(target GLenum, s GLdouble)

//go:sdl3staticextern
var glMultiTexCoord1dv func(target GLenum, v *GLdouble)

//go:sdl3staticextern
var glMultiTexCoord1f func(target GLenum, s GLfloat)

//go:sdl3staticextern
var glMultiTexCoord1fv func(target GLenum, v *GLfloat)

//go:sdl3staticextern
var glMultiTexCoord1i func(target GLenum, s GLint)

//go:sdl3staticextern
var glMultiTexCoord1iv func(target GLenum, v *GLint)

//go:sdl3staticextern
var glMultiTexCoord1s func(target GLenum, s GLshort)

//go:sdl3staticextern
var glMultiTexCoord1sv func(target GLenum, v *GLshort)

//go:sdl3staticextern
var glMultiTexCoord2d func(target GLenum, s GLdouble, t GLdouble)

//go:sdl3staticextern
var glMultiTexCoord2dv func(target GLenum, v *GLdouble)

//go:sdl3staticextern
var glMultiTexCoord2f func(target GLenum, s GLfloat, t GLfloat)

//go:sdl3staticextern
var glMultiTexCoord2fv func(target GLenum, v *GLfloat)

//go:sdl3staticextern
var glMultiTexCoord2i func(target GLenum, s GLint, t GLint)

//go:sdl3staticextern
var glMultiTexCoord2iv func(target GLenum, v *GLint)

//go:sdl3staticextern
var glMultiTexCoord2s func(target GLenum, s GLshort, t GLshort)

//go:sdl3staticextern
var glMultiTexCoord2sv func(target GLenum, v *GLshort)

//go:sdl3staticextern
var glMultiTexCoord3d func(target GLenum, s GLdouble, t GLdouble, r GLdouble)

//go:sdl3staticextern
var glMultiTexCoord3dv func(target GLenum, v *GLdouble)

//go:sdl3staticextern
var glMultiTexCoord3f func(target GLenum, s GLfloat, t GLfloat, r GLfloat)

//go:sdl3staticextern
var glMultiTexCoord3fv func(target GLenum, v *GLfloat)

//go:sdl3staticextern
var glMultiTexCoord3i func(target GLenum, s GLint, t GLint, r GLint)

//go:sdl3staticextern
var glMultiTexCoord3iv func(target GLenum, v *GLint)

//go:sdl3staticextern
var glMultiTexCoord3s func(target GLenum, s GLshort, t GLshort, r GLshort)

//go:sdl3staticextern
var glMultiTexCoord3sv func(target GLenum, v *GLshort)

//go:sdl3staticextern
var glMultiTexCoord4d func(target GLenum, s GLdouble, t GLdouble, r GLdouble, q GLdouble)

//go:sdl3staticextern
var glMultiTexCoord4dv func(target GLenum, v *GLdouble)

//go:sdl3staticextern
var glMultiTexCoord4f func(target GLenum, s GLfloat, t GLfloat, r GLfloat, q GLfloat)

//go:sdl3staticextern
var glMultiTexCoord4fv func(target GLenum, v *GLfloat)

//go:sdl3staticextern
var glMultiTexCoord4i func(target GLenum, s GLint, t GLint, r GLint, q GLint)

//go:sdl3staticextern
var glMultiTexCoord4iv func(target GLenum, v *GLint)

//go:sdl3staticextern
var glMultiTexCoord4s func(target GLenum, s GLshort, t GLshort, r GLshort, q GLshort)

//go:sdl3staticextern
var glMultiTexCoord4sv func(target GLenum, v *GLshort)

//go:sdl3staticextern
var glLoadTransposeMatrixd func(m [16]GLdouble)

//go:sdl3staticextern
var glLoadTransposeMatrixf func(m [16]GLfloat)

//go:sdl3staticextern
var glMultTransposeMatrixd func(m [16]GLdouble)

//go:sdl3staticextern
var glMultTransposeMatrixf func(m [16]GLfloat)

//go:sdl3staticextern
var glSampleCoverage func(value GLclampf, invert GLboolean)

/*
 * GL_ARB_multitexture (ARB extension 1 and OpenGL 1.2.1)
 */
const ARB_multitexture = 1

const TEXTURE0_ARB = 0x84C0
const TEXTURE1_ARB = 0x84C1
const TEXTURE2_ARB = 0x84C2
const TEXTURE3_ARB = 0x84C3
const TEXTURE4_ARB = 0x84C4
const TEXTURE5_ARB = 0x84C5
const TEXTURE6_ARB = 0x84C6
const TEXTURE7_ARB = 0x84C7
const TEXTURE8_ARB = 0x84C8
const TEXTURE9_ARB = 0x84C9
const TEXTURE10_ARB = 0x84CA
const TEXTURE11_ARB = 0x84CB
const TEXTURE12_ARB = 0x84CC
const TEXTURE13_ARB = 0x84CD
const TEXTURE14_ARB = 0x84CE
const TEXTURE15_ARB = 0x84CF
const TEXTURE16_ARB = 0x84D0
const TEXTURE17_ARB = 0x84D1
const TEXTURE18_ARB = 0x84D2
const TEXTURE19_ARB = 0x84D3
const TEXTURE20_ARB = 0x84D4
const TEXTURE21_ARB = 0x84D5
const TEXTURE22_ARB = 0x84D6
const TEXTURE23_ARB = 0x84D7
const TEXTURE24_ARB = 0x84D8
const TEXTURE25_ARB = 0x84D9
const TEXTURE26_ARB = 0x84DA
const TEXTURE27_ARB = 0x84DB
const TEXTURE28_ARB = 0x84DC
const TEXTURE29_ARB = 0x84DD
const TEXTURE30_ARB = 0x84DE
const TEXTURE31_ARB = 0x84DF
const ACTIVE_TEXTURE_ARB = 0x84E0
const CLIENT_ACTIVE_TEXTURE_ARB = 0x84E1
const MAX_TEXTURE_UNITS_ARB = 0x84E2

//go:sdl3staticextern
var glActiveTextureARB func(texture GLenum)

//go:sdl3staticextern
var glClientActiveTextureARB func(texture GLenum)

//go:sdl3staticextern
var glMultiTexCoord1dARB func(target GLenum, s GLdouble)

//go:sdl3staticextern
var glMultiTexCoord1dvARB func(target GLenum, v *GLdouble)

//go:sdl3staticextern
var glMultiTexCoord1fARB func(target GLenum, s GLfloat)

//go:sdl3staticextern
var glMultiTexCoord1fvARB func(target GLenum, v *GLfloat)

//go:sdl3staticextern
var glMultiTexCoord1iARB func(target GLenum, s GLint)

//go:sdl3staticextern
var glMultiTexCoord1ivARB func(target GLenum, v *GLint)

//go:sdl3staticextern
var glMultiTexCoord1sARB func(target GLenum, s GLshort)

//go:sdl3staticextern
var glMultiTexCoord1svARB func(target GLenum, v *GLshort)

//go:sdl3staticextern
var glMultiTexCoord2dARB func(target GLenum, s GLdouble, t GLdouble)

//go:sdl3staticextern
var glMultiTexCoord2dvARB func(target GLenum, v *GLdouble)

//go:sdl3staticextern
var glMultiTexCoord2fARB func(target GLenum, s GLfloat, t GLfloat)

//go:sdl3staticextern
var glMultiTexCoord2fvARB func(target GLenum, v *GLfloat)

//go:sdl3staticextern
var glMultiTexCoord2iARB func(target GLenum, s GLint, t GLint)

//go:sdl3staticextern
var glMultiTexCoord2ivARB func(target GLenum, v *GLint)

//go:sdl3staticextern
var glMultiTexCoord2sARB func(target GLenum, s GLshort, t GLshort)

//go:sdl3staticextern
var glMultiTexCoord2svARB func(target GLenum, v *GLshort)

//go:sdl3staticextern
var glMultiTexCoord3dARB func(target GLenum, s GLdouble, t GLdouble, r GLdouble)

//go:sdl3staticextern
var glMultiTexCoord3dvARB func(target GLenum, v *GLdouble)

//go:sdl3staticextern
var glMultiTexCoord3fARB func(target GLenum, s GLfloat, t GLfloat, r GLfloat)

//go:sdl3staticextern
var glMultiTexCoord3fvARB func(target GLenum, v *GLfloat)

//go:sdl3staticextern
var glMultiTexCoord3iARB func(target GLenum, s GLint, t GLint, r GLint)

//go:sdl3staticextern
var glMultiTexCoord3ivARB func(target GLenum, v *GLint)

//go:sdl3staticextern
var glMultiTexCoord3sARB func(target GLenum, s GLshort, t GLshort, r GLshort)

//go:sdl3staticextern
var glMultiTexCoord3svARB func(target GLenum, v *GLshort)

//go:sdl3staticextern
var glMultiTexCoord4dARB func(target GLenum, s GLdouble, t GLdouble, r GLdouble, q GLdouble)

//go:sdl3staticextern
var glMultiTexCoord4dvARB func(target GLenum, v *GLdouble)

//go:sdl3staticextern
var glMultiTexCoord4fARB func(target GLenum, s GLfloat, t GLfloat, r GLfloat, q GLfloat)

//go:sdl3staticextern
var glMultiTexCoord4fvARB func(target GLenum, v *GLfloat)

//go:sdl3staticextern
var glMultiTexCoord4iARB func(target GLenum, s GLint, t GLint, r GLint, q GLint)

//go:sdl3staticextern
var glMultiTexCoord4ivARB func(target GLenum, v *GLint)

//go:sdl3staticextern
var glMultiTexCoord4sARB func(target GLenum, s GLshort, t GLshort, r GLshort, q GLshort)

//go:sdl3staticextern
var glMultiTexCoord4svARB func(target GLenum, v *GLshort)

/*
 * End system-specific stuff
 **********************************************************************/
