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
 * # CategoryBlendmode
 *
 * Blend modes decide how two colors will mix together. There are both
 * standard modes for basic needs and a means to create custom modes,
 * dictating what sort of math to do on what color components.
 */
package sdl

/**
* A set of blend modes used in drawing operations.
*
* These predefined blend modes are supported everywhere.
*
* Additional values may be obtained from SDL_ComposeCustomBlendMode.
*
* \since This datatype is available since SDL 3.2.0.
*
* \sa SDL_ComposeCustomBlendMode
 */
type BlendMode uint32

const (
	BM_None               BlendMode = 0x00000000 /**< no blending: dstRGBA = srcRGBA */
	BM_Blend              BlendMode = 0x00000001 /**< alpha blending: dstRGB = (srcRGB * srcA) + (dstRGB * (1-srcA)), dstA = srcA + (dstA * (1-srcA)) */
	BM_BlendPremultiplied BlendMode = 0x00000010 /**< pre-multiplied alpha blending: dstRGBA = srcRGBA + (dstRGBA * (1-srcA)) */
	BM_Add                BlendMode = 0x00000002 /**< additive blending: dstRGB = (srcRGB * srcA) + dstRGB, dstA = dstA */
	BM_AddPremultiplied   BlendMode = 0x00000020 /**< pre-multiplied additive blending: dstRGB = srcRGB + dstRGB, dstA = dstA */
	BM_Mod                BlendMode = 0x00000004 /**< color modulate: dstRGB = srcRGB * dstRGB, dstA = dstA */
	BM_Mul                BlendMode = 0x00000008 /**< color multiply: dstRGB = (srcRGB * dstRGB) + (dstRGB * (1-srcA)), dstA = dstA */
	BM_Invalid            BlendMode = 0x7FFFFFFF
)

/**
* The blend operation used when combining source and destination pixel
* components.
*
* \since This enum is available since SDL 3.2.0.
 */
type BlendOperation int

const (
	BO_Add         BlendOperation = 0x1 /**< dst + src: supported by all renderers */
	BO_Subtract    BlendOperation = 0x2 /**< src - dst : supported by D3D, OpenGL, OpenGLES, and Vulkan */
	BO_RevSubtract BlendOperation = 0x3 /**< dst - src : supported by D3D, OpenGL, OpenGLES, and Vulkan */
	BO_Minimum     BlendOperation = 0x4 /**< min(dst, src) : supported by D3D, OpenGL, OpenGLES, and Vulkan */
	BO_Maximum     BlendOperation = 0x5 /**< max(dst, src) : supported by D3D, OpenGL, OpenGLES, and Vulkan */
)

/**
* The normalized factor used to multiply pixel components.
*
* The blend factors are multiplied with the pixels from a drawing operation
* (src) and the pixels from the render target (dst) before the blend
* operation. The comma-separated factors listed above are always applied in
* the component order red, green, blue, and alpha.
*
* \since This enum is available since SDL 3.2.0.
 */
type BlendFactor int

const (
	BF_Zero             = 0x1 /**< 0, 0, 0, 0 */
	BF_One              = 0x2 /**< 1, 1, 1, 1 */
	BF_SrcColor         = 0x3 /**< srcR, srcG, srcB, srcA */
	BF_OneMinusSrcColor = 0x4 /**< 1-srcR, 1-srcG, 1-srcB, 1-srcA */
	BF_SrcAlpha         = 0x5 /**< srcA, srcA, srcA, srcA */
	BF_OneMinusSrcAlpha = 0x6 /**< 1-srcA, 1-srcA, 1-srcA, 1-srcA */
	BF_DstColor         = 0x7 /**< dstR, dstG, dstB, dstA */
	BF_OneMinusDstColor = 0x8 /**< 1-dstR, 1-dstG, 1-dstB, 1-dstA */
	BF_DstAlpha         = 0x9 /**< dstA, dstA, dstA, dstA */
	BF_OneMinusDstAlpha = 0xA /**< 1-dstA, 1-dstA, 1-dstA, 1-dstA */
)

/**
* Compose a custom blend mode for renderers.
*
* The functions SDL_SetRenderDrawBlendMode and SDL_SetTextureBlendMode accept
* the SDL_BlendMode returned by this function if the renderer supports it.
*
* A blend mode controls how the pixels from a drawing operation (source) get
* combined with the pixels from the render target (destination). First, the
* components of the source and destination pixels get multiplied with their
* blend factors. Then, the blend operation takes the two products and
* calculates the result that will get stored in the render target.
*
* Expressed in pseudocode, it would look like this:
*
* ```c
* dstRGB = colorOperation(srcRGB * srcColorFactor, dstRGB * dstColorFactor);
* dstA = alphaOperation(srcA * srcAlphaFactor, dstA * dstAlphaFactor);
* ```
*
* Where the functions `colorOperation(src, dst)` and `alphaOperation(src,
* dst)` can return one of the following:
*
* - `src + dst`
* - `src - dst`
* - `dst - src`
* - `min(src, dst)`
* - `max(src, dst)`
*
* The red, green, and blue components are always multiplied with the first,
* second, and third components of the SDL_BlendFactor, respectively. The
* fourth component is not used.
*
* The alpha component is always multiplied with the fourth component of the
* SDL_BlendFactor. The other components are not used in the alpha
* calculation.
*
* Support for these blend modes varies for each renderer. To check if a
* specific SDL_BlendMode is supported, create a renderer and pass it to
* either SDL_SetRenderDrawBlendMode or SDL_SetTextureBlendMode. They will
* return with an error if the blend mode is not supported.
*
* This list describes the support of custom blend modes for each renderer.
* All renderers support the four blend modes listed in the SDL_BlendMode
* enumeration.
*
* - **direct3d**: Supports all operations with all factors. However, some
*   factors produce unexpected results with `SDL_BLENDOPERATION_MINIMUM` and
*   `SDL_BLENDOPERATION_MAXIMUM`.
* - **direct3d11**: Same as Direct3D 9.
* - **opengl**: Supports the `SDL_BLENDOPERATION_ADD` operation with all
*   factors. OpenGL versions 1.1, 1.2, and 1.3 do not work correctly here.
* - **opengles2**: Supports the `SDL_BLENDOPERATION_ADD`,
*   `SDL_BLENDOPERATION_SUBTRACT`, `SDL_BLENDOPERATION_REV_SUBTRACT`
*   operations with all factors.
* - **psp**: No custom blend mode support.
* - **software**: No custom blend mode support.
*
* Some renderers do not provide an alpha component for the default render
* target. The `SDL_BLENDFACTOR_DST_ALPHA` and
* `SDL_BLENDFACTOR_ONE_MINUS_DST_ALPHA` factors do not have an effect in this
* case.
*
* \param srcColorFactor the SDL_BlendFactor applied to the red, green, and
*                       blue components of the source pixels.
* \param dstColorFactor the SDL_BlendFactor applied to the red, green, and
*                       blue components of the destination pixels.
* \param colorOperation the SDL_BlendOperation used to combine the red,
*                       green, and blue components of the source and
*                       destination pixels.
* \param srcAlphaFactor the SDL_BlendFactor applied to the alpha component of
*                       the source pixels.
* \param dstAlphaFactor the SDL_BlendFactor applied to the alpha component of
*                       the destination pixels.
* \param alphaOperation the SDL_BlendOperation used to combine the alpha
*                       component of the source and destination pixels.
* \returns an SDL_BlendMode that represents the chosen factors and
*          operations.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_SetRenderDrawBlendMode
* \sa SDL_GetRenderDrawBlendMode
* \sa SDL_SetTextureBlendMode
* \sa SDL_GetTextureBlendMode
 */
//go:sdl3extern
var ComposeCustomBlendMode func(srcColorFactor BlendFactor,
	dstColorFactor BlendFactor,
	colorOperation BlendOperation,
	srcAlphaFactor BlendFactor,
	dstAlphaFactor BlendFactor,
	alphaOperation BlendOperation) BlendMode
