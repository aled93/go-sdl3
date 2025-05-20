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

/* WIKI CATEGORY: GPU */

/**
 * # CategoryGPU
 *
 * The GPU API offers a cross-platform way for apps to talk to modern graphics
 * hardware. It offers both 3D graphics and compute support, in the style of
 * Metal, Vulkan, and Direct3D 12.
 *
 * A basic workflow might be something like this:
 *
 * The app creates a GPU device with SDL_CreateGPUDevice(), and assigns it to
 * a window with SDL_ClaimWindowForGPUDevice()--although strictly speaking you
 * can render offscreen entirely, perhaps for image processing, and not use a
 * window at all.
 *
 * Next, the app prepares static data (things that are created once and used
 * over and over). For example:
 *
 * - Shaders (programs that run on the GPU): use SDL_CreateGPUShader().
 * - Vertex buffers (arrays of geometry data) and other rendering data: use
 *   SDL_CreateGPUBuffer() and SDL_UploadToGPUBuffer().
 * - Textures (images): use SDL_CreateGPUTexture() and
 *   SDL_UploadToGPUTexture().
 * - Samplers (how textures should be read from): use SDL_CreateGPUSampler().
 * - Render pipelines (precalculated rendering state): use
 *   SDL_CreateGPUGraphicsPipeline()
 *
 * To render, the app creates one or more command buffers, with
 * SDL_AcquireGPUCommandBuffer(). Command buffers collect rendering
 * instructions that will be submitted to the GPU in batch. Complex scenes can
 * use multiple command buffers, maybe configured across multiple threads in
 * parallel, as long as they are submitted in the correct order, but many apps
 * will just need one command buffer per frame.
 *
 * Rendering can happen to a texture (what other APIs call a "render target")
 * or it can happen to the swapchain texture (which is just a special texture
 * that represents a window's contents). The app can use
 * SDL_WaitAndAcquireGPUSwapchainTexture() to render to the window.
 *
 * Rendering actually happens in a Render Pass, which is encoded into a
 * command buffer. One can encode multiple render passes (or alternate between
 * render and compute passes) in a single command buffer, but many apps might
 * simply need a single render pass in a single command buffer. Render Passes
 * can render to up to four color textures and one depth texture
 * simultaneously. If the set of textures being rendered to needs to change,
 * the Render Pass must be ended and a new one must be begun.
 *
 * The app calls SDL_BeginGPURenderPass(). Then it sets states it needs for
 * each draw:
 *
 * - SDL_BindGPUGraphicsPipeline()
 * - SDL_SetGPUViewport()
 * - SDL_BindGPUVertexBuffers()
 * - SDL_BindGPUVertexSamplers()
 * - etc
 *
 * Then, make the actual draw commands with these states:
 *
 * - SDL_DrawGPUPrimitives()
 * - SDL_DrawGPUPrimitivesIndirect()
 * - SDL_DrawGPUIndexedPrimitivesIndirect()
 * - etc
 *
 * After all the drawing commands for a pass are complete, the app should call
 * SDL_EndGPURenderPass(). Once a render pass ends all render-related state is
 * reset.
 *
 * The app can begin new Render Passes and make new draws in the same command
 * buffer until the entire scene is rendered.
 *
 * Once all of the render commands for the scene are complete, the app calls
 * SDL_SubmitGPUCommandBuffer() to send it to the GPU for processing.
 *
 * If the app needs to read back data from texture or buffers, the API has an
 * efficient way of doing this, provided that the app is willing to tolerate
 * some latency. When the app uses SDL_DownloadFromGPUTexture() or
 * SDL_DownloadFromGPUBuffer(), submitting the command buffer with
 * SDL_SubmitGPUCommandBufferAndAcquireFence() will return a fence handle that
 * the app can poll or wait on in a thread. Once the fence indicates that the
 * command buffer is done processing, it is safe to read the downloaded data.
 * Make sure to call SDL_ReleaseGPUFence() when done with the fence.
 *
 * The API also has "compute" support. The app calls SDL_BeginGPUComputePass()
 * with compute-writeable textures and/or buffers, which can be written to in
 * a compute shader. Then it sets states it needs for the compute dispatches:
 *
 * - SDL_BindGPUComputePipeline()
 * - SDL_BindGPUComputeStorageBuffers()
 * - SDL_BindGPUComputeStorageTextures()
 *
 * Then, dispatch compute work:
 *
 * - SDL_DispatchGPUCompute()
 *
 * For advanced users, this opens up powerful GPU-driven workflows.
 *
 * Graphics and compute pipelines require the use of shaders, which as
 * mentioned above are small programs executed on the GPU. Each backend
 * (Vulkan, Metal, D3D12) requires a different shader format. When the app
 * creates the GPU device, the app lets the device know which shader formats
 * the app can provide. It will then select the appropriate backend depending
 * on the available shader formats and the backends available on the platform.
 * When creating shaders, the app must provide the correct shader format for
 * the selected backend. If you would like to learn more about why the API
 * works this way, there is a detailed
 * [blog post](https://moonside.games/posts/layers-all-the-way-down/)
 * explaining this situation.
 *
 * It is optimal for apps to pre-compile the shader formats they might use,
 * but for ease of use SDL provides a separate project,
 * [SDL_shadercross](https://github.com/libsdl-org/SDL_shadercross)
 * , for performing runtime shader cross-compilation. It also has a CLI
 * interface for offline precompilation as well.
 *
 * This is an extremely quick overview that leaves out several important
 * details. Already, though, one can see that GPU programming can be quite
 * complex! If you just need simple 2D graphics, the
 * [Render API](https://wiki.libsdl.org/SDL3/CategoryRender)
 * is much easier to use but still hardware-accelerated. That said, even for
 * 2D applications the performance benefits and expressiveness of the GPU API
 * are significant.
 *
 * The GPU API targets a feature set with a wide range of hardware support and
 * ease of portability. It is designed so that the app won't have to branch
 * itself by querying feature support. If you need cutting-edge features with
 * limited hardware support, this API is probably not for you.
 *
 * Examples demonstrating proper usage of this API can be found
 * [here](https://github.com/TheSpydog/SDL_gpu_examples)
 * .
 *
 * ## Performance considerations
 *
 * Here are some basic tips for maximizing your rendering performance.
 *
 * - Beginning a new render pass is relatively expensive. Use as few render
 *   passes as you can.
 * - Minimize the amount of state changes. For example, binding a pipeline is
 *   relatively cheap, but doing it hundreds of times when you don't need to
 *   will slow the performance significantly.
 * - Perform your data uploads as early as possible in the frame.
 * - Don't churn resources. Creating and releasing resources is expensive.
 *   It's better to create what you need up front and cache it.
 * - Don't use uniform buffers for large amounts of data (more than a matrix
 *   or so). Use a storage buffer instead.
 * - Use cycling correctly. There is a detailed explanation of cycling further
 *   below.
 * - Use culling techniques to minimize pixel writes. The less writing the GPU
 *   has to do the better. Culling can be a very advanced topic but even
 *   simple culling techniques can boost performance significantly.
 *
 * In general try to remember the golden rule of performance: doing things is
 * more expensive than not doing things. Don't Touch The Driver!
 *
 * ## FAQ
 *
 * **Question: When are you adding more advanced features, like ray tracing or
 * mesh shaders?**
 *
 * Answer: We don't have immediate plans to add more bleeding-edge features,
 * but we certainly might in the future, when these features prove worthwhile,
 * and reasonable to implement across several platforms and underlying APIs.
 * So while these things are not in the "never" category, they are definitely
 * not "near future" items either.
 *
 * **Question: Why is my shader not working?**
 *
 * Answer: A common oversight when using shaders is not properly laying out
 * the shader resources/registers correctly. The GPU API is very strict with
 * how it wants resources to be laid out and it's difficult for the API to
 * automatically validate shaders to see if they have a compatible layout. See
 * the documentation for SDL_CreateGPUShader() and
 * SDL_CreateGPUComputePipeline() for information on the expected layout.
 *
 * Another common issue is not setting the correct number of samplers,
 * textures, and buffers in SDL_GPUShaderCreateInfo. If possible use shader
 * reflection to extract the required information from the shader
 * automatically instead of manually filling in the struct's values.
 *
 * **Question: My application isn't performing very well. Is this the GPU
 * API's fault?**
 *
 * Answer: No. Long answer: The GPU API is a relatively thin layer over the
 * underlying graphics API. While it's possible that we have done something
 * inefficiently, it's very unlikely especially if you are relatively
 * inexperienced with GPU rendering. Please see the performance tips above and
 * make sure you are following them. Additionally, tools like
 * [RenderDoc](https://renderdoc.org/)
 * can be very helpful for diagnosing incorrect behavior and performance
 * issues.
 *
 * ## System Requirements
 *
 * ### Vulkan
 *
 * SDL driver name: "vulkan" (for use in SDL_CreateGPUDevice() and
 * SDL_PROP_GPU_DEVICE_CREATE_NAME_STRING)
 *
 * Supported on Windows, Linux, Nintendo Switch, and certain Android devices.
 * Requires Vulkan 1.0 with the following extensions and device features:
 *
 * - `VK_KHR_swapchain`
 * - `VK_KHR_maintenance1`
 * - `independentBlend`
 * - `imageCubeArray`
 * - `depthClamp`
 * - `shaderClipDistance`
 * - `drawIndirectFirstInstance`
 * - `sampleRateShading`
 *
 * ### D3D12
 *
 * SDL driver name: "direct3d12"
 *
 * Supported on Windows 10 or newer, Xbox One (GDK), and Xbox Series X|S
 * (GDK). Requires a GPU that supports DirectX 12 Feature Level 11_1.
 *
 * ### Metal
 *
 * SDL driver name: "metal"
 *
 * Supported on macOS 10.14+ and iOS/tvOS 13.0+. Hardware requirements vary by
 * operating system:
 *
 * - macOS requires an Apple Silicon or
 *   [Intel Mac2 family](https://developer.apple.com/documentation/metal/mtlfeatureset/mtlfeatureset_macos_gpufamily2_v1?language=objc)
 *   GPU
 * - iOS/tvOS requires an A9 GPU or newer
 * - iOS Simulator and tvOS Simulator are unsupported
 *
 * ## Coordinate System
 *
 * The GPU API uses a left-handed coordinate system, following the convention
 * of D3D12 and Metal. Specifically:
 *
 * - **Normalized Device Coordinates:** The lower-left corner has an x,y
 *   coordinate of `(-1.0, -1.0)`. The upper-right corner is `(1.0, 1.0)`. Z
 *   values range from `[0.0, 1.0]` where 0 is the near plane.
 * - **Viewport Coordinates:** The top-left corner has an x,y coordinate of
 *   `(0, 0)` and extends to the bottom-right corner at `(viewportWidth,
 *   viewportHeight)`. +Y is down.
 * - **Texture Coordinates:** The top-left corner has an x,y coordinate of
 *   `(0, 0)` and extends to the bottom-right corner at `(1.0, 1.0)`. +Y is
 *   down.
 *
 * If the backend driver differs from this convention (e.g. Vulkan, which has
 * an NDC that assumes +Y is down), SDL will automatically convert the
 * coordinate system behind the scenes, so you don't need to perform any
 * coordinate flipping logic in your shaders.
 *
 * ## Uniform Data
 *
 * Uniforms are for passing data to shaders. The uniform data will be constant
 * across all executions of the shader.
 *
 * There are 4 available uniform slots per shader stage (where the stages are
 * vertex, fragment, and compute). Uniform data pushed to a slot on a stage
 * keeps its value throughout the command buffer until you call the relevant
 * Push function on that slot again.
 *
 * For example, you could write your vertex shaders to read a camera matrix
 * from uniform binding slot 0, push the camera matrix at the start of the
 * command buffer, and that data will be used for every subsequent draw call.
 *
 * It is valid to push uniform data during a render or compute pass.
 *
 * Uniforms are best for pushing small amounts of data. If you are pushing
 * more than a matrix or two per call you should consider using a storage
 * buffer instead.
 *
 * ## A Note On Cycling
 *
 * When using a command buffer, operations do not occur immediately - they
 * occur some time after the command buffer is submitted.
 *
 * When a resource is used in a pending or active command buffer, it is
 * considered to be "bound". When a resource is no longer used in any pending
 * or active command buffers, it is considered to be "unbound".
 *
 * If data resources are bound, it is unspecified when that data will be
 * unbound unless you acquire a fence when submitting the command buffer and
 * wait on it. However, this doesn't mean you need to track resource usage
 * manually.
 *
 * All of the functions and structs that involve writing to a resource have a
 * "cycle" bool. SDL_GPUTransferBuffer, SDL_GPUBuffer, and SDL_GPUTexture all
 * effectively function as ring buffers on internal resources. When cycle is
 * true, if the resource is bound, the cycle rotates to the next unbound
 * internal resource, or if none are available, a new one is created. This
 * means you don't have to worry about complex state tracking and
 * synchronization as long as cycling is correctly employed.
 *
 * For example: you can call SDL_MapGPUTransferBuffer(), write texture data,
 * SDL_UnmapGPUTransferBuffer(), and then SDL_UploadToGPUTexture(). The next
 * time you write texture data to the transfer buffer, if you set the cycle
 * param to true, you don't have to worry about overwriting any data that is
 * not yet uploaded.
 *
 * Another example: If you are using a texture in a render pass every frame,
 * this can cause a data dependency between frames. If you set cycle to true
 * in the SDL_GPUColorTargetInfo struct, you can prevent this data dependency.
 *
 * Cycling will never undefine already bound data. When cycling, all data in
 * the resource is considered to be undefined for subsequent commands until
 * that data is written again. You must take care not to read undefined data.
 *
 * Note that when cycling a texture, the entire texture will be cycled, even
 * if only part of the texture is used in the call, so you must consider the
 * entire texture to contain undefined data after cycling.
 *
 * You must also take care not to overwrite a section of data that has been
 * referenced in a command without cycling first. It is OK to overwrite
 * unreferenced data in a bound resource without cycling, but overwriting a
 * section of data that has already been referenced will produce unexpected
 * results.
 *
 * ## Debugging
 *
 * At some point of your GPU journey, you will probably encounter issues that
 * are not traceable with regular debugger - for example, your code compiles
 * but you get an empty screen, or your shader fails in runtime.
 *
 * For debugging such cases, there are tools that allow visually inspecting
 * the whole GPU frame, every drawcall, every bound resource, memory buffers,
 * etc. They are the following, per platform:
 *
 * * For Windows/Linux, use
 *   [RenderDoc](https://renderdoc.org/)
 * * For MacOS (Metal), use Xcode built-in debugger (Open XCode, go to Debug >
 *   Debug Executable..., select your application, set "GPU Frame Capture" to
 *   "Metal" in scheme "Options" window, run your app, and click the small
 *   Metal icon on the bottom to capture a frame)
 *
 * Aside from that, you may want to enable additional debug layers to receive
 * more detailed error messages, based on your GPU backend:
 *
 * * For D3D12, the debug layer is an optional feature that can be installed
 *   via "Windows Settings -> System -> Optional features" and adding the
 *   "Graphics Tools" optional feature.
 * * For Vulkan, you will need to install Vulkan SDK on Windows, and on Linux,
 *   you usually have some sort of `vulkan-validation-layers` system package
 *   that should be installed.
 * * For Metal, it should be enough just to run the application from XCode to
 *   receive detailed errors or warnings in the output.
 *
 * Don't hesitate to use tools as RenderDoc when encountering runtime issues
 * or unexpected output on screen, quick GPU frame inspection can usually help
 * you fix the majority of such problems.
 */
package sdl

/* Type Declarations */

/**
 * An opaque handle representing the SDL_GPU context.
 *
 * \since This struct is available since SDL 3.2.0.
 */
type GPUDevice uintptr

/**
 * An opaque handle representing a buffer.
 *
 * Used for vertices, indices, indirect draw commands, and general compute
 * data.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUBuffer
 * \sa SDL_UploadToGPUBuffer
 * \sa SDL_DownloadFromGPUBuffer
 * \sa SDL_CopyGPUBufferToBuffer
 * \sa SDL_BindGPUVertexBuffers
 * \sa SDL_BindGPUIndexBuffer
 * \sa SDL_BindGPUVertexStorageBuffers
 * \sa SDL_BindGPUFragmentStorageBuffers
 * \sa SDL_DrawGPUPrimitivesIndirect
 * \sa SDL_DrawGPUIndexedPrimitivesIndirect
 * \sa SDL_BindGPUComputeStorageBuffers
 * \sa SDL_DispatchGPUComputeIndirect
 * \sa SDL_ReleaseGPUBuffer
 */
type GPUBuffer uintptr

/**
 * An opaque handle representing a transfer buffer.
 *
 * Used for transferring data to and from the device.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUTransferBuffer
 * \sa SDL_MapGPUTransferBuffer
 * \sa SDL_UnmapGPUTransferBuffer
 * \sa SDL_UploadToGPUBuffer
 * \sa SDL_UploadToGPUTexture
 * \sa SDL_DownloadFromGPUBuffer
 * \sa SDL_DownloadFromGPUTexture
 * \sa SDL_ReleaseGPUTransferBuffer
 */
type GPUTransferBuffer uintptr

/**
 * An opaque handle representing a texture.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUTexture
 * \sa SDL_UploadToGPUTexture
 * \sa SDL_DownloadFromGPUTexture
 * \sa SDL_CopyGPUTextureToTexture
 * \sa SDL_BindGPUVertexSamplers
 * \sa SDL_BindGPUVertexStorageTextures
 * \sa SDL_BindGPUFragmentSamplers
 * \sa SDL_BindGPUFragmentStorageTextures
 * \sa SDL_BindGPUComputeStorageTextures
 * \sa SDL_GenerateMipmapsForGPUTexture
 * \sa SDL_BlitGPUTexture
 * \sa SDL_ReleaseGPUTexture
 */
type GPUTexture uintptr

/**
 * An opaque handle representing a sampler.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUSampler
 * \sa SDL_BindGPUVertexSamplers
 * \sa SDL_BindGPUFragmentSamplers
 * \sa SDL_ReleaseGPUSampler
 */
type GPUSampler uintptr

/**
 * An opaque handle representing a compiled shader object.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 * \sa SDL_CreateGPUGraphicsPipeline
 * \sa SDL_ReleaseGPUShader
 */
type GPUShader uintptr

/**
 * An opaque handle representing a compute pipeline.
 *
 * Used during compute passes.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUComputePipeline
 * \sa SDL_BindGPUComputePipeline
 * \sa SDL_ReleaseGPUComputePipeline
 */
type GPUComputePipeline uintptr

/**
 * An opaque handle representing a graphics pipeline.
 *
 * Used during render passes.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 * \sa SDL_BindGPUGraphicsPipeline
 * \sa SDL_ReleaseGPUGraphicsPipeline
 */
type GPUGraphicsPipeline uintptr

/**
 * An opaque handle representing a command buffer.
 *
 * Most state is managed via command buffers. When setting state using a
 * command buffer, that state is local to the command buffer.
 *
 * Commands only begin execution on the GPU once SDL_SubmitGPUCommandBuffer is
 * called. Once the command buffer is submitted, it is no longer valid to use
 * it.
 *
 * Command buffers are executed in submission order. If you submit command
 * buffer A and then command buffer B all commands in A will begin executing
 * before any command in B begins executing.
 *
 * In multi-threading scenarios, you should only access a command buffer on
 * the thread you acquired it from.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_AcquireGPUCommandBuffer
 * \sa SDL_SubmitGPUCommandBuffer
 * \sa SDL_SubmitGPUCommandBufferAndAcquireFence
 */
type GPUCommandBuffer uintptr

/**
 * An opaque handle representing a render pass.
 *
 * This handle is transient and should not be held or referenced after
 * SDL_EndGPURenderPass is called.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_BeginGPURenderPass
 * \sa SDL_EndGPURenderPass
 */
type GPURenderPass uintptr

/**
 * An opaque handle representing a compute pass.
 *
 * This handle is transient and should not be held or referenced after
 * SDL_EndGPUComputePass is called.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_BeginGPUComputePass
 * \sa SDL_EndGPUComputePass
 */
type GPUComputePass uintptr

/**
 * An opaque handle representing a copy pass.
 *
 * This handle is transient and should not be held or referenced after
 * SDL_EndGPUCopyPass is called.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_BeginGPUCopyPass
 * \sa SDL_EndGPUCopyPass
 */
type GPUCopyPass uintptr

/**
 * An opaque handle representing a fence.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_SubmitGPUCommandBufferAndAcquireFence
 * \sa SDL_QueryGPUFence
 * \sa SDL_WaitForGPUFences
 * \sa SDL_ReleaseGPUFence
 */
type GPUFence uintptr

/**
 * Specifies the primitive topology of a graphics pipeline.
 *
 * If you are using POINTLIST you must include a point size output in the
 * vertex shader.
 *
 * - For HLSL compiling to SPIRV you must decorate a float32 output with
 *   [[vk::builtin("PointSize")]].
 * - For GLSL you must set the gl_PointSize builtin.
 * - For MSL you must include a float32 output with the [[point_size]]
 *   decorator.
 *
 * Note that sized point topology is totally unsupported on D3D12. Any size
 * other than 1 will be ignored. In general, you should avoid using point
 * topology for both compatibility and performance reasons. You WILL regret
 * using it.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUPrimitiveType int32

const (
	GPUPrimitiveTypeTriangleList  GPUPrimitiveType = iota /**< A series of separate triangles. */
	GPUPrimitiveTypeTriangleStrip                         /**< A series of connected triangles. */
	GPUPrimitiveTypeLineList                              /**< A series of separate lines. */
	GPUPrimitiveTypeLineStrip                             /**< A series of connected lines. */
	GPUPrimitiveTypePointList                             /**< A series of separate points. */
)

/**
 * Specifies how the contents of a texture attached to a render pass are
 * treated at the beginning of the render pass.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_BeginGPURenderPass
 */
type GPULoadOp int32

const (
	GPULoadOpLoad     GPULoadOp = iota /**< The previous contents of the texture will be preserved. */
	GPULoadOpClear                     /**< The contents of the texture will be cleared to a color. */
	GPULoadOpDontCare                  /**< The previous contents of the texture need not be preserved. The contents will be undefined. */
)

/**
 * Specifies how the contents of a texture attached to a render pass are
 * treated at the end of the render pass.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_BeginGPURenderPass
 */
type GPUStoreOp int32

const (
	GPUStoreOpStore           GPUStoreOp = iota /**< The contents generated during the render pass will be written to memory. */
	GPUStoreOpDontCare                          /**< The contents generated during the render pass are not needed and may be discarded. The contents will be undefined. */
	GPUStoreOpResolve                           /**< The multisample contents generated during the render pass will be resolved to a non-multisample texture. The contents in the multisample texture may then be discarded and will be undefined. */
	GPUStoreOpResolveAndStore                   /**< The multisample contents generated during the render pass will be resolved to a non-multisample texture. The contents in the multisample texture will be written to memory. */
)

/**
 * Specifies the size of elements in an index buffer.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUIndexElementSize int32

const (
	GPUIndexElementSize16Bit GPUIndexElementSize = iota /**< The index elements are 16-bit. */
	GPUIndexElementSize32Bit                            /**< The index elements are 32-bit. */
)

/**
 * Specifies the pixel format of a texture.
 *
 * Texture format support varies depending on driver, hardware, and usage
 * flags. In general, you should use SDL_GPUTextureSupportsFormat to query if
 * a format is supported before using it. However, there are a few guaranteed
 * formats.
 *
 * FIXME: Check universal support for 32-bit component formats FIXME: Check
 * universal support for SIMULTANEOUS_READ_WRITE
 *
 * For SAMPLER usage, the following formats are universally supported:
 *
 * - R8G8B8A8_UNorm
 * - B8G8R8A8_UNorm
 * - R8_UNorm
 * - R8_SNORM
 * - R8G8_UNorm
 * - R8G8_SNORM
 * - R8G8B8A8_SNORM
 * - R16_Float
 * - R16G16_Float
 * - R16G16B16A16_Float
 * - R32_Float
 * - R32G32_Float
 * - R32G32B32A32_Float
 * - R11G11B10_UFLOAT
 * - R8G8B8A8_UNorm_SRGB
 * - B8G8R8A8_UNorm_SRGB
 * - D16_UNorm
 *
 * For COLOR_TARGET usage, the following formats are universally supported:
 *
 * - R8G8B8A8_UNorm
 * - B8G8R8A8_UNorm
 * - R8_UNorm
 * - R16_Float
 * - R16G16_Float
 * - R16G16B16A16_Float
 * - R32_Float
 * - R32G32_Float
 * - R32G32B32A32_Float
 * - R8_UInt
 * - R8G8_UInt
 * - R8G8B8A8_UInt
 * - R16_UInt
 * - R16G16_UInt
 * - R16G16B16A16_UInt
 * - R8_INT
 * - R8G8_INT
 * - R8G8B8A8_INT
 * - R16_INT
 * - R16G16_INT
 * - R16G16B16A16_INT
 * - R8G8B8A8_UNorm_SRGB
 * - B8G8R8A8_UNorm_SRGB
 *
 * For STORAGE usages, the following formats are universally supported:
 *
 * - R8G8B8A8_UNorm
 * - R8G8B8A8_SNORM
 * - R16G16B16A16_Float
 * - R32_Float
 * - R32G32_Float
 * - R32G32B32A32_Float
 * - R8G8B8A8_UInt
 * - R16G16B16A16_UInt
 * - R8G8B8A8_INT
 * - R16G16B16A16_INT
 *
 * For DEPTH_STENCIL_TARGET usage, the following formats are universally
 * supported:
 *
 * - D16_UNorm
 * - Either (but not necessarily both!) D24_UNorm or D32_Float
 * - Either (but not necessarily both!) D24_UNorm_S8_UInt or D32_Float_S8_UInt
 *
 * Unless D16_UNorm is sufficient for your purposes, always check which of
 * D24/D32 is supported before creating a depth-stencil texture!
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUTexture
 * \sa SDL_GPUTextureSupportsFormat
 */
type GPUTextureFormat int32

const (
	GPUTextureFormatInvalid GPUTextureFormat = iota

	/* Unsigned Normalized Float Color Formats */
	GPUTextureFormatA8_UNorm
	GPUTextureFormatR8_UNorm
	GPUTextureFormatR8G8_UNorm
	GPUTextureFormatR8G8B8A8_UNorm
	GPUTextureFormatR16_UNorm
	GPUTextureFormatR16G16_UNorm
	GPUTextureFormatR16G16B16A16_UNorm
	GPUTextureFormatR10G10B10A2_UNorm
	GPUTextureFormatB5G6R5_UNorm
	GPUTextureFormatB5G5R5A1_UNorm
	GPUTextureFormatB4G4R4A4_UNorm
	GPUTextureFormatB8G8R8A8_UNorm
	/* Compressed Unsigned Normalized Float Color Formats */
	GPUTextureFormatBC1_RGBA_UNorm
	GPUTextureFormatBC2_RGBA_UNorm
	GPUTextureFormatBC3_RGBA_UNorm
	GPUTextureFormatBC4_R_UNorm
	GPUTextureFormatBC5_RG_UNorm
	GPUTextureFormatBC7_RGBA_UNorm
	/* Compressed Signed Float Color Formats */
	GPUTextureFormatBC6H_RGB_Float
	/* Compressed Unsigned Float Color Formats */
	GPUTextureFormatBC6H_RGB_UFloat
	/* Signed Normalized Float Color Formats  */
	GPUTextureFormatR8_SNorm
	GPUTextureFormatR8G8_SNorm
	GPUTextureFormatR8G8B8A8_SNorm
	GPUTextureFormatR16_SNorm
	GPUTextureFormatR16G16_SNorm
	GPUTextureFormatR16G16B16A16_SNorm
	/* Signed Float Color Formats */
	GPUTextureFormatR16_Float
	GPUTextureFormatR16G16_Float
	GPUTextureFormatR16G16B16A16_Float
	GPUTextureFormatR32_Float
	GPUTextureFormatR32G32_Float
	GPUTextureFormatR32G32B32A32_Float
	/* Unsigned Float Color Formats */
	GPUTextureFormatR11G11B10_UFloat
	/* Unsigned Integer Color Formats */
	GPUTextureFormatR8_UInt
	GPUTextureFormatR8G8_UInt
	GPUTextureFormatR8G8B8A8_UInt
	GPUTextureFormatR16_UInt
	GPUTextureFormatR16G16_UInt
	GPUTextureFormatR16G16B16A16_UInt
	GPUTextureFormatR32_UInt
	GPUTextureFormatR32G32_UInt
	GPUTextureFormatR32G32B32A32_UInt
	/* Signed Integer Color Formats */
	GPUTextureFormatR8_Int
	GPUTextureFormatR8G8_Int
	GPUTextureFormatR8G8B8A8_Int
	GPUTextureFormatR16_Int
	GPUTextureFormatR16G16_Int
	GPUTextureFormatR16G16B16A16_Int
	GPUTextureFormatR32_Int
	GPUTextureFormatR32G32_Int
	GPUTextureFormatR32G32B32A32_Int
	/* SRGB Unsigned Normalized Color Formats */
	GPUTextureFormatR8G8B8A8_UNorm_SRGB
	GPUTextureFormatB8G8R8A8_UNorm_SRGB
	/* Compressed SRGB Unsigned Normalized Color Formats */
	GPUTextureFormatBC1_RGBA_UNorm_SRGB
	GPUTextureFormatBC2_RGBA_UNorm_SRGB
	GPUTextureFormatBC3_RGBA_UNorm_SRGB
	GPUTextureFormatBC7_RGBA_UNorm_SRGB
	/* Depth Formats */
	GPUTextureFormatD16_UNorm
	GPUTextureFormatD24_UNorm
	GPUTextureFormatD32_Float
	GPUTextureFormatD24_UNorm_S8_UInt
	GPUTextureFormatD32_Float_S8_UInt
	/* Compressed ASTC Normalized Float Color Formats*/
	GPUTextureFormatASTC_4x4_UNorm
	GPUTextureFormatASTC_5x4_UNorm
	GPUTextureFormatASTC_5x5_UNorm
	GPUTextureFormatASTC_6x5_UNorm
	GPUTextureFormatASTC_6x6_UNorm
	GPUTextureFormatASTC_8x5_UNorm
	GPUTextureFormatASTC_8x6_UNorm
	GPUTextureFormatASTC_8x8_UNorm
	GPUTextureFormatASTC_10x5_UNorm
	GPUTextureFormatASTC_10x6_UNorm
	GPUTextureFormatASTC_10x8_UNorm
	GPUTextureFormatASTC_10x10_UNorm
	GPUTextureFormatASTC_12x10_UNorm
	GPUTextureFormatASTC_12x12_UNorm
	/* Compressed SRGB ASTC Normalized Float Color Formats*/
	GPUTextureFormatASTC_4x4_UNorm_SRGB
	GPUTextureFormatASTC_5x4_UNorm_SRGB
	GPUTextureFormatASTC_5x5_UNorm_SRGB
	GPUTextureFormatASTC_6x5_UNorm_SRGB
	GPUTextureFormatASTC_6x6_UNorm_SRGB
	GPUTextureFormatASTC_8x5_UNorm_SRGB
	GPUTextureFormatASTC_8x6_UNorm_SRGB
	GPUTextureFormatASTC_8x8_UNorm_SRGB
	GPUTextureFormatASTC_10x5_UNorm_SRGB
	GPUTextureFormatASTC_10x6_UNorm_SRGB
	GPUTextureFormatASTC_10x8_UNorm_SRGB
	GPUTextureFormatASTC_10x10_UNorm_SRGB
	GPUTextureFormatASTC_12x10_UNorm_SRGB
	GPUTextureFormatASTC_12x12_UNorm_SRGB
	/* Compressed ASTC Signed Float Color Formats*/
	GPUTextureFormatASTC_4x4_Float
	GPUTextureFormatASTC_5x4_Float
	GPUTextureFormatASTC_5x5_Float
	GPUTextureFormatASTC_6x5_Float
	GPUTextureFormatASTC_6x6_Float
	GPUTextureFormatASTC_8x5_Float
	GPUTextureFormatASTC_8x6_Float
	GPUTextureFormatASTC_8x8_Float
	GPUTextureFormatASTC_10x5_Float
	GPUTextureFormatASTC_10x6_Float
	GPUTextureFormatASTC_10x8_Float
	GPUTextureFormatASTC_10x10_Float
	GPUTextureFormatASTC_12x10_Float
	GPUTextureFormatASTC_12x12_Float
)

/**
 * Specifies how a texture is intended to be used by the client.
 *
 * A texture must have at least one usage flag. Note that some usage flag
 * combinations are invalid.
 *
 * With regards to compute storage usage, READ | WRITE means that you can have
 * shader A that only writes into the texture and shader B that only reads
 * from the texture and bind the same texture to either shader respectively.
 * SIMULTANEOUS means that you can do reads and writes within the same shader
 * or compute pass. It also implies that atomic ops can be used, since those
 * are read-modify-write operations. If you use SIMULTANEOUS, you are
 * responsible for avoiding data races, as there is no data synchronization
 * within a compute pass. Note that SIMULTANEOUS usage is only supported by a
 * limited number of texture formats.
 *
 * \since This datatype is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUTexture
 */
type GPUTextureUsageFlags uint32

const (
	GPUTextureUsageSampler                             GPUTextureUsageFlags = 1 << 0 /**< Texture supports sampling. */
	GPUTextureUsageColorTarget                         GPUTextureUsageFlags = 1 << 1 /**< Texture is a color render target. */
	GPUTextureUsageDepthStencilTarget                  GPUTextureUsageFlags = 1 << 2 /**< Texture is a depth stencil target. */
	GPUTextureUsageGraphicsStorageRead                 GPUTextureUsageFlags = 1 << 3 /**< Texture supports storage reads in graphics stages. */
	GPUTextureUsageComputeStorageRead                  GPUTextureUsageFlags = 1 << 4 /**< Texture supports storage reads in the compute stage. */
	GPUTextureUsageComputeStorageWrite                 GPUTextureUsageFlags = 1 << 5 /**< Texture supports storage writes in the compute stage. */
	GPUTextureUsageComputeStorageSimultaneousReadWrite GPUTextureUsageFlags = 1 << 6 /**< Texture supports reads and writes in the same compute shader. This is NOT equivalent to READ | WRITE. */
)

/**
 * Specifies the type of a texture.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUTexture
 */
type GPUTextureType int32

const (
	GPUTextureType2D        GPUTextureType = iota /**< The texture is a 2-dimensional image. */
	GPUTextureType2DArray                         /**< The texture is a 2-dimensional array image. */
	GPUTextureType3D                              /**< The texture is a 3-dimensional image. */
	GPUTextureTypeCube                            /**< The texture is a cube image. */
	GPUTextureTypeCubeArray                       /**< The texture is a cube array image. */
)

/**
 * Specifies the sample count of a texture.
 *
 * Used in multisampling. Note that this value only applies when the texture
 * is used as a render target.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUTexture
 * \sa SDL_GPUTextureSupportsSampleCount
 */
type GPUSampleCount int32

const (
	GPUSampleCount1 GPUSampleCount = iota /**< No multisampling. */
	GPUSampleCount2                       /**< MSAA 2x */
	GPUSampleCount4                       /**< MSAA 4x */
	GPUSampleCount8                       /**< MSAA 8x */
)

/**
 * Specifies the face of a cube map.
 *
 * Can be passed in as the layer field in texture-related structs.
 *
 * \since This enum is available since SDL 3.2.0.
 */
type GPUCubeMapFace int32

const (
	GPUCubemapFacePositiveX GPUCubeMapFace = iota
	GPUCubemapFaceNegativeX
	GPUCubemapFacePositiveY
	GPUCubemapFaceNegativeY
	GPUCubemapFacePositiveZ
	GPUCubemapFaceNegativeZ
)

/**
 * Specifies how a buffer is intended to be used by the client.
 *
 * A buffer must have at least one usage flag. Note that some usage flag
 * combinations are invalid.
 *
 * Unlike textures, READ | WRITE can be used for simultaneous read-write
 * usage. The same data synchronization concerns as textures apply.
 *
 * If you use a STORAGE flag, the data in the buffer must respect std140
 * layout conventions. In practical terms this means you must ensure that vec3
 * and vec4 fields are 16-byte aligned.
 *
 * \since This datatype is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUBuffer
 */
type GPUBufferUsageFlags uint32

const (
	GPUBufferUsageVertex              GPUBufferUsageFlags = 1 << 0 /**< Buffer is a vertex buffer. */
	GPUBufferUsageIndex               GPUBufferUsageFlags = 1 << 1 /**< Buffer is an index buffer. */
	GPUBufferUsageIndirect            GPUBufferUsageFlags = 1 << 2 /**< Buffer is an indirect buffer. */
	GPUBufferUsageGraphicsStorageRead GPUBufferUsageFlags = 1 << 3 /**< Buffer supports storage reads in graphics stages. */
	GPUBufferUsageComputeStorageRead  GPUBufferUsageFlags = 1 << 4 /**< Buffer supports storage reads in the compute stage. */
	GPUBufferUsageComputeStorageWrite GPUBufferUsageFlags = 1 << 5 /**< Buffer supports storage writes in the compute stage. */
)

/**
 * Specifies how a transfer buffer is intended to be used by the client.
 *
 * Note that mapping and copying FROM an upload transfer buffer or TO a
 * download transfer buffer is undefined behavior.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUTransferBuffer
 */
type GPUTransferBufferUsage int32

const (
	GPUTransferBufferUsageUpload GPUTransferBufferUsage = iota
	GPUTransferBufferUsageDownload
)

/**
 * Specifies which stage a shader program corresponds to.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
type GPUShaderStage int32

const (
	GPUShaderStageVertex GPUShaderStage = iota
	GPUShaderStageFragment
)

/**
 * Specifies the format of shader code.
 *
 * Each format corresponds to a specific backend that accepts it.
 *
 * \since This datatype is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
type GPUShaderFormat uint32

const (
	GPUShaderFormatInvalid  GPUShaderFormat = 0
	GPUShaderFormatPrivate  GPUShaderFormat = 1 << 0 /**< Shaders for NDA'd platforms. */
	GPUShaderFormatSpirV    GPUShaderFormat = 1 << 1 /**< SPIR-V shaders for Vulkan. */
	GPUShaderFormatDXBC     GPUShaderFormat = 1 << 2 /**< DXBC SM5_1 shaders for D3D12. */
	GPUShaderFormatDXIL     GPUShaderFormat = 1 << 3 /**< DXIL SM6_0 shaders for D3D12. */
	GPUShaderFormatMSL      GPUShaderFormat = 1 << 4 /**< MSL shaders for Metal. */
	GPUShaderFormatMetalLib GPUShaderFormat = 1 << 5 /**< Precompiled metallib shaders for Metal. */
)

/**
 * Specifies the format of a vertex attribute.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUVertexElementFormat int32

const (
	GPUVertexElementFormatIinvalid GPUVertexElementFormat = iota

	/* 32-bit Signed Integers */
	GPUVertexElementFormatInt
	GPUVertexElementFormatInt2
	GPUVertexElementFormatInt3
	GPUVertexElementFormatInt4

	/* 32-bit Unsigned Integers */
	GPUVertexElementFormatUInt
	GPUVertexElementFormatUInt2
	GPUVertexElementFormatUInt3
	GPUVertexElementFormatUInt4

	/* 32-bit Floats */
	GPUVertexElementFormatFloat
	GPUVertexElementFormatFloat2
	GPUVertexElementFormatFloat3
	GPUVertexElementFormatFloat4

	/* 8-bit Signed Integers */
	GPUVertexElementFormatByte2
	GPUVertexElementFormatByte4

	/* 8-bit Unsigned Integers */
	GPUVertexElementFormatUByte2
	GPUVertexElementFormatUByte4

	/* 8-bit Signed Normalized */
	GPUVertexElementFormatByte2Norm
	GPUVertexElementFormatByte4Norm

	/* 8-bit Unsigned Normalized */
	GPUVertexElementFormatUByte2Norm
	GPUVertexElementFormatUByte4Norm

	/* 16-bit Signed Integers */
	GPUVertexElementFormatShort2
	GPUVertexElementFormatShort4

	/* 16-bit Unsigned Integers */
	GPUVertexElementFormatUShort2
	GPUVertexElementFormatUShort4

	/* 16-bit Signed Normalized */
	GPUVertexElementFormatShort2Norm
	GPUVertexElementFormatShort4Norm

	/* 16-bit Unsigned Normalized */
	GPUVertexElementFormatUShort2Norm
	GPUVertexElementFormatUShort4Norm

	/* 16-bit Floats */
	GPUVertexElementFormatHalf2
	GPUVertexElementFormatHalf4
)

/**
 * Specifies the rate at which vertex attributes are pulled from buffers.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUVertexInputRate int32

const (
	GPUVertexInputRateVertex   GPUVertexInputRate = iota /**< Attribute addressing is a function of the vertex index. */
	GPUVertexInputRateInstance                           /**< Attribute addressing is a function of the instance index. */
)

/**
 * Specifies the fill mode of the graphics pipeline.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUFillMode int32

const (
	GPUFillModeFill GPUFillMode = iota /**< Polygons will be rendered via rasterization. */
	GPUFillModeLine                    /**< Polygon edges will be drawn as line segments. */
)

/**
 * Specifies the facing direction in which triangle faces will be culled.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUCullMode int32

const (
	GPUCullModeNone  GPUCullMode = iota /**< No triangles are culled. */
	GPUCullModeFront                    /**< Front-facing triangles are culled. */
	GPUCullModeBack                     /**< Back-facing triangles are culled. */
)

/**
 * Specifies the vertex winding that will cause a triangle to be determined to
 * be front-facing.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUFrontFace int32

const (
	GPUFrontFaceCounterClockwise GPUFrontFace = iota /**< A triangle with counter-clockwise vertex winding will be considered front-facing. */
	GPUFrontFaceClockwise                            /**< A triangle with clockwise vertex winding will be considered front-facing. */
)

/**
 * Specifies a comparison operator for depth, stencil and sampler operations.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUCompareOp int32

const (
	GPUCompareOpI              GPUCompareOp = iota
	GPUCompareOpNever                       /**< The comparison always evaluates false. */
	GPUCompareOpLess                        /**< The comparison evaluates reference < test. */
	GPUCompareOpEqual                       /**< The comparison evaluates reference == test. */
	GPUCompareOpLessOrEqual                 /**< The comparison evaluates reference <= test. */
	GPUCompareOpGreater                     /**< The comparison evaluates reference > test. */
	GPUCompareOpNotEqual                    /**< The comparison evaluates reference != test. */
	GPUCompareOpGreaterOrEqual              /**< The comparison evalutes reference >= test. */
	GPUCompareOpAlways                      /**< The comparison always evaluates true. */
)

/**
 * Specifies what happens to a stored stencil value if stencil tests fail or
 * pass.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUStencilOp int32

const (
	GPUStencilOpInvalid           GPUStencilOp = iota
	GPUStencilOpKeep                           /**< Keeps the current value. */
	GPUStencilOpZero                           /**< Sets the value to 0. */
	GPUStencilOpReplace                        /**< Sets the value to reference. */
	GPUStencilOpIncrementAndClamp              /**< Increments the current value and clamps to the maximum value. */
	GPUStencilOpDecrementAndClamp              /**< Decrements the current value and clamps to 0. */
	GPUStencilOpInvert                         /**< Bitwise-inverts the current value. */
	GPUStencilOpIncrementAndWrap               /**< Increments the current value and wraps back to 0. */
	GPUStencilOpDecrementAndWrap               /**< Decrements the current value and wraps to the maximum value. */
)

/**
 * Specifies the operator to be used when pixels in a render target are
 * blended with existing pixels in the texture.
 *
 * The source color is the value written by the fragment shader. The
 * destination color is the value currently existing in the texture.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUBlendOp int32

const (
	GPUBlendOpInvalid         GPUBlendOp = iota
	GPUBlendOpAdd                        /**< (source * source_factor) + (destination * destination_factor) */
	GPUBlendOpSubtract                   /**< (source * source_factor) - (destination * destination_factor) */
	GPUBlendOpReverseSubtract            /**< (destination * destination_factor) - (source * source_factor) */
	GPUBlendOpMin                        /**< min(source, destination) */
	GPUBlendOpMax                        /**< max(source, destination) */
)

/**
 * Specifies a blending factor to be used when pixels in a render target are
 * blended with existing pixels in the texture.
 *
 * The source color is the value written by the fragment shader. The
 * destination color is the value currently existing in the texture.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUBlendFactor int32

const (
	GPUBlendFactorInvalid               GPUBlendFactor = iota
	GPUBlendFactorZero                                 /**< 0 */
	GPUBlendFactorOne                                  /**< 1 */
	GPUBlendFactorSrcColor                             /**< source color */
	GPUBlendFactorOneMinusSrcColor                     /**< 1 - source color */
	GPUBlendFactorDstColor                             /**< destination color */
	GPUBlendFactorOneMinusDstColor                     /**< 1 - destination color */
	GPUBlendFactorSrcAlpha                             /**< source alpha */
	GPUBlendFactorOneMinusSrcAlpha                     /**< 1 - source alpha */
	GPUBlendFactorDstAlpha                             /**< destination alpha */
	GPUBlendFactorOneMinusDstAlpha                     /**< 1 - destination alpha */
	GPUBlendFactorConstantColor                        /**< blend constant */
	GPUBlendFactorOneMinusConstantColor                /**< 1 - blend constant */
	GPUBlendFactorSrcAlphaSaturate                     /**< min(source alpha, 1 - destination alpha) */
)

/**
 * Specifies which color components are written in a graphics pipeline.
 *
 * \since This datatype is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 */
type GPUColorComponentFlags uint8

const (
	GPUColorComponentR GPUColorComponentFlags = 1 << 0 /**< the red component */
	GPUColorComponentG GPUColorComponentFlags = 1 << 1 /**< the green component */
	GPUColorComponentB GPUColorComponentFlags = 1 << 2 /**< the blue component */
	GPUColorComponentA GPUColorComponentFlags = 1 << 3 /**< the alpha component */
)

/**
 * Specifies a filter operation used by a sampler.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUSampler
 */
type GPUFilter int32

const (
	GPUFilterNearest GPUFilter = iota /**< Point filtering. */
	GPUFilterLinear                   /**< Linear filtering. */
)

/**
 * Specifies a mipmap mode used by a sampler.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUSampler
 */
type GPUSamplerMipmapMode int32

const (
	GPUSamplerMipmapModeNearest GPUSamplerMipmapMode = iota /**< Point filtering. */
	GPUSamplerMipmapModeLinear                              /**< Linear filtering. */
)

/**
 * Specifies behavior of texture sampling when the coordinates exceed the 0-1
 * range.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUSampler
 */
type GPUSamplerAddressMode int32

const (
	GPUSamplerAddressModeRepeat         GPUSamplerAddressMode = iota /**< Specifies that the coordinates will wrap around. */
	GPUSamplerAddressModeMirroredRepeat                              /**< Specifies that the coordinates will wrap around mirrored. */
	GPUSamplerAddressModeClampToEdge                                 /**< Specifies that the coordinates will clamp to the 0-1 range. */
)

/**
 * Specifies the timing that will be used to present swapchain textures to the
 * OS.
 *
 * VSYNC mode will always be supported. IMMEDIATE and MAILBOX modes may not be
 * supported on certain systems.
 *
 * It is recommended to query SDL_WindowSupportsGPUPresentMode after claiming
 * the window if you wish to change the present mode to IMMEDIATE or MAILBOX.
 *
 * - VSYNC: Waits for vblank before presenting. No tearing is possible. If
 *   there is a pending image to present, the new image is enqueued for
 *   presentation. Disallows tearing at the cost of visual latency.
 * - IMMEDIATE: Immediately presents. Lowest latency option, but tearing may
 *   occur.
 * - MAILBOX: Waits for vblank before presenting. No tearing is possible. If
 *   there is a pending image to present, the pending image is replaced by the
 *   new image. Similar to VSYNC, but with reduced visual latency.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_SetGPUSwapchainParameters
 * \sa SDL_WindowSupportsGPUPresentMode
 * \sa SDL_WaitAndAcquireGPUSwapchainTexture
 */
type GPUPresentMode int32

const (
	GPUPresentModeVVync GPUPresentMode = iota
	GPUPresentModeImmediate
	GPUPresentModeMailbox
)

/**
 * Specifies the texture format and colorspace of the swapchain textures.
 *
 * SDR will always be supported. Other compositions may not be supported on
 * certain systems.
 *
 * It is recommended to query SDL_WindowSupportsGPUSwapchainComposition after
 * claiming the window if you wish to change the swapchain composition from
 * SDR.
 *
 * - SDR: B8G8R8A8 or R8G8B8A8 swapchain. Pixel values are in sRGB encoding.
 * - SDR_LINEAR: B8G8R8A8_SRGB or R8G8B8A8_SRGB swapchain. Pixel values are
 *   stored in memory in sRGB encoding but accessed in shaders in "linear
 *   sRGB" encoding which is sRGB but with a linear transfer function.
 * - HDR_EXTENDED_LINEAR: R16G16B16A16_Float swapchain. Pixel values are in
 *   extended linear sRGB encoding and permits values outside of the [0, 1]
 *   range.
 * - HDR10_ST2084: A2R10G10B10 or A2B10G10R10 swapchain. Pixel values are in
 *   BT.2020 ST2084 (PQ) encoding.
 *
 * \since This enum is available since SDL 3.2.0.
 *
 * \sa SDL_SetGPUSwapchainParameters
 * \sa SDL_WindowSupportsGPUSwapchainComposition
 * \sa SDL_WaitAndAcquireGPUSwapchainTexture
 */
type GPUSwapchainComposition int32

const (
	GPUSwapchainCompositionSDR GPUSwapchainComposition = iota
	GPUSwapchainCompositionSDRLinear
	GPUSwapchainCompositionHDRExtendedLinear
	GPUSwapchainCompositionHDR10_ST2084
)

/* Structures */

/**
 * A structure specifying a viewport.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_SetGPUViewport
 */
type GPUViewport struct {
	X        float32 /**< The left offset of the viewport. */
	Y        float32 /**< The top offset of the viewport. */
	W        float32 /**< The width of the viewport. */
	H        float32 /**< The height of the viewport. */
	MinDepth float32 /**< The minimum depth of the viewport. */
	MaxDepth float32 /**< The maximum depth of the viewport. */
}

/**
 * A structure specifying parameters related to transferring data to or from a
 * texture.
 *
 * If either of `pixels_per_row` or `rows_per_layer` is zero, then width and
 * height of passed SDL_GPUTextureRegion to SDL_UploadToGPUTexture or
 * SDL_DownloadFromGPUTexture are used as default values respectively and data
 * is considered to be tightly packed.
 *
 * **WARNING**: Direct3D 12 requires texture data row pitch to be 256 byte
 * aligned, and offsets to be aligned to 512 bytes. If they are not, SDL will
 * make a temporary copy of the data that is properly aligned, but this adds
 * overhead to the transfer process. Apps can avoid this by aligning their
 * data appropriately, or using a different GPU backend than Direct3D 12.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_UploadToGPUTexture
 * \sa SDL_DownloadFromGPUTexture
 */
type GPUTextureTransferInfo struct {
	TransferBuffer GPUTransferBuffer /**< The transfer buffer used in the transfer operation. */
	Offset         uint32            /**< The starting byte of the image data in the transfer buffer. */
	PixelsPerRow   uint32            /**< The number of pixels from one row to the next. */
	RowsPerLayer   uint32            /**< The number of rows from one layer/depth-slice to the next. */
}

/**
 * A structure specifying a location in a transfer buffer.
 *
 * Used when transferring buffer data to or from a transfer buffer.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_UploadToGPUBuffer
 * \sa SDL_DownloadFromGPUBuffer
 */
type GPUTransferBufferLocation struct {
	TransferBuffer GPUTransferBuffer /**< The transfer buffer used in the transfer operation. */
	Offset         uint32            /**< The starting byte of the buffer data in the transfer buffer. */
}

/**
 * A structure specifying a location in a texture.
 *
 * Used when copying data from one texture to another.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CopyGPUTextureToTexture
 */
type GPUTextureLocation struct {
	Texture  GPUTexture /**< The texture used in the copy operation. */
	MipLevel uint32     /**< The mip level index of the location. */
	Layer    uint32     /**< The layer index of the location. */
	X        uint32     /**< The left offset of the location. */
	Y        uint32     /**< The top offset of the location. */
	Z        uint32     /**< The front offset of the location. */
}

/**
 * A structure specifying a region of a texture.
 *
 * Used when transferring data to or from a texture.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_UploadToGPUTexture
 * \sa SDL_DownloadFromGPUTexture
 * \sa SDL_CreateGPUTexture
 */
type GPUTextureRegion struct {
	Texture  GPUTexture /**< The texture used in the copy operation. */
	MipLevel uint32     /**< The mip level index to transfer. */
	Layer    uint32     /**< The layer index to transfer. */
	X        uint32     /**< The left offset of the region. */
	Y        uint32     /**< The top offset of the region. */
	Z        uint32     /**< The front offset of the region. */
	W        uint32     /**< The width of the region. */
	H        uint32     /**< The height of the region. */
	D        uint32     /**< The depth of the region. */
}

/**
 * A structure specifying a region of a texture used in the blit operation.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_BlitGPUTexture
 */
type GPUBlitRegion struct {
	Texture           GPUTexture /**< The texture. */
	MipLevel          uint32     /**< The mip level index of the region. */
	LayerOrDepthPlane uint32     /**< The layer index or depth plane of the region. This value is treated as a layer index on 2D array and cube textures, and as a depth plane on 3D textures. */
	X                 uint32     /**< The left offset of the region. */
	Y                 uint32     /**< The top offset of the region.  */
	W                 uint32     /**< The width of the region. */
	H                 uint32     /**< The height of the region. */
}

/**
 * A structure specifying a location in a buffer.
 *
 * Used when copying data between buffers.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CopyGPUBufferToBuffer
 */
type GPUBufferLocation struct {
	Buffer GPUBuffer /**< The buffer. */
	Offset uint32    /**< The starting byte within the buffer. */
}

/**
 * A structure specifying a region of a buffer.
 *
 * Used when transferring data to or from buffers.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_UploadToGPUBuffer
 * \sa SDL_DownloadFromGPUBuffer
 */
type GPUBufferRegion struct {
	Buffer GPUBuffer /**< The buffer. */
	Offset uint32    /**< The starting byte within the buffer. */
	Size   uint32    /**< The size in bytes of the region. */
}

/**
 * A structure specifying the parameters of an indirect draw command.
 *
 * Note that the `first_vertex` and `first_instance` parameters are NOT
 * compatible with built-in vertex/instance ID variables in shaders (for
 * example, SV_VertexID); GPU APIs and shader languages do not define these
 * built-in variables consistently, so if your shader depends on them, the
 * only way to keep behavior consistent and portable is to always pass 0 for
 * the correlating parameter in the draw calls.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_DrawGPUPrimitivesIndirect
 */
type GPUIndirectDrawCommand struct {
	NumVertices   uint32 /**< The number of vertices to draw. */
	NumInstances  uint32 /**< The number of instances to draw. */
	FirstVertex   uint32 /**< The index of the first vertex to draw. */
	FirstInstance uint32 /**< The ID of the first instance to draw. */
}

/**
 * A structure specifying the parameters of an indexed indirect draw command.
 *
 * Note that the `first_vertex` and `first_instance` parameters are NOT
 * compatible with built-in vertex/instance ID variables in shaders (for
 * example, SV_VertexID); GPU APIs and shader languages do not define these
 * built-in variables consistently, so if your shader depends on them, the
 * only way to keep behavior consistent and portable is to always pass 0 for
 * the correlating parameter in the draw calls.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_DrawGPUIndexedPrimitivesIndirect
 */
type GPUIndexedIndirectDrawCommand struct {
	NumIndices    uint32 /**< The number of indices to draw per instance. */
	NumInstances  uint32 /**< The number of instances to draw. */
	FirstIndex    uint32 /**< The base index within the index buffer. */
	VertexOffset  int32  /**< The value added to the vertex index before indexing into the vertex buffer. */
	FirstInstance uint32 /**< The ID of the first instance to draw. */
}

/**
 * A structure specifying the parameters of an indexed dispatch command.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_DispatchGPUComputeIndirect
 */
type GPUIndirectDispatchCommand struct {
	GroupCountX uint32 /**< The number of local workgroups to dispatch in the X dimension. */
	GroupCountY uint32 /**< The number of local workgroups to dispatch in the Y dimension. */
	GroupCountZ uint32 /**< The number of local workgroups to dispatch in the Z dimension. */
}

/* State structures */

/**
 * A structure specifying the parameters of a sampler.
 *
 * Note that mip_lod_bias is a no-op for the Metal driver. For Metal, LOD bias
 * must be applied via shader instead.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUSampler
 * \sa SDL_GPUFilter
 * \sa SDL_GPUSamplerMipmapMode
 * \sa SDL_GPUSamplerAddressMode
 * \sa SDL_GPUCompareOp
 */
type GPUSamplerCreateInfo struct {
	MinFilter        GPUFilter             /**< The minification filter to apply to lookups. */
	MagFilter        GPUFilter             /**< The magnification filter to apply to lookups. */
	MipmapMode       GPUSamplerMipmapMode  /**< The mipmap filter to apply to lookups. */
	AddressModeU     GPUSamplerAddressMode /**< The addressing mode for U coordinates outside [0, 1). */
	AddressModeV     GPUSamplerAddressMode /**< The addressing mode for V coordinates outside [0, 1). */
	AddressModeW     GPUSamplerAddressMode /**< The addressing mode for W coordinates outside [0, 1). */
	MipLodBias       float32               /**< The bias to be added to mipmap LOD calculation. */
	MaxAnisotropy    float32               /**< The anisotropy value clamp used by the sampler. If enable_anisotropy is false, this is ignored. */
	CompareOp        GPUCompareOp          /**< The comparison operator to apply to fetched data before filtering. */
	MinLod           float32               /**< Clamps the minimum of the computed LOD value. */
	MaxLod           float32               /**< Clamps the maximum of the computed LOD value. */
	EnableAnisotropy bool                  /**< true to enable anisotropic filtering. */
	EnableCompare    bool                  /**< true to enable comparison against a reference value during lookups. */
	Padding1         uint8
	Padding2         uint8

	Props PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

/**
 * A structure specifying the parameters of vertex buffers used in a graphics
 * pipeline.
 *
 * When you call SDL_BindGPUVertexBuffers, you specify the binding slots of
 * the vertex buffers. For example if you called SDL_BindGPUVertexBuffers with
 * a first_slot of 2 and num_bindings of 3, the binding slots 2, 3, 4 would be
 * used by the vertex buffers you pass in.
 *
 * Vertex attributes are linked to buffers via the buffer_slot field of
 * SDL_GPUVertexAttribute. For example, if an attribute has a buffer_slot of
 * 0, then that attribute belongs to the vertex buffer bound at slot 0.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GPUVertexAttribute
 * \sa SDL_GPUVertexInputRate
 */
type GPUVertexBufferDescription struct {
	Slot             uint32             /**< The binding slot of the vertex buffer. */
	Pitch            uint32             /**< The byte pitch between consecutive elements of the vertex buffer. */
	InputRate        GPUVertexInputRate /**< Whether attribute addressing is a function of the vertex index or instance index. */
	InstanceStepRate uint32             /**< Reserved for future use. Must be set to 0. */
}

/**
 * A structure specifying a vertex attribute.
 *
 * All vertex attribute locations provided to an SDL_GPUVertexInputState must
 * be unique.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GPUVertexBufferDescription
 * \sa SDL_GPUVertexInputState
 * \sa SDL_GPUVertexElementFormat
 */
type GPUVertexAttribute struct {
	Location   uint32                 /**< The shader input location index. */
	BufferSlot uint32                 /**< The binding slot of the associated vertex buffer. */
	Format     GPUVertexElementFormat /**< The size and type of the attribute data. */
	Offset     uint32                 /**< The byte offset of this attribute relative to the start of the vertex element. */
}

/**
 * A structure specifying the parameters of a graphics pipeline vertex input
 * state.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GPUGraphicsPipelineCreateInfo
 * \sa SDL_GPUVertexBufferDescription
 * \sa SDL_GPUVertexAttribute
 */
type GPUVertexInputState struct {
	VertexBufferDescriptions *GPUVertexBufferDescription /**< A pointer to an array of vertex buffer descriptions. */
	NumVertexBuffers         uint32                      /**< The number of vertex buffer descriptions in the above array. */
	VertexAttributes         *GPUVertexAttribute         /**< A pointer to an array of vertex attribute descriptions. */
	NumVertexAttributes      uint32                      /**< The number of vertex attribute descriptions in the above array. */
}

/**
 * A structure specifying the stencil operation state of a graphics pipeline.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GPUDepthStencilState
 */
type GPUStencilOpState struct {
	FailOp      GPUStencilOp /**< The action performed on samples that fail the stencil test. */
	PassOp      GPUStencilOp /**< The action performed on samples that pass the depth and stencil tests. */
	DepthFailOp GPUStencilOp /**< The action performed on samples that pass the stencil test and fail the depth test. */
	CompareOp   GPUCompareOp /**< The comparison operator used in the stencil test. */
}

/**
 * A structure specifying the blend state of a color target.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GPUColorTargetDescription
 */
type GPUColorTargetBlendState struct {
	SrcColorBlendFactor  GPUBlendFactor         /**< The value to be multiplied by the source RGB value. */
	DstColorBlendFactor  GPUBlendFactor         /**< The value to be multiplied by the destination RGB value. */
	ColorBlendOp         GPUBlendOp             /**< The blend operation for the RGB components. */
	SrcAlphaBlendFactor  GPUBlendFactor         /**< The value to be multiplied by the source alpha. */
	DstAlphaBlendFactor  GPUBlendFactor         /**< The value to be multiplied by the destination alpha. */
	AlphaBlendOp         GPUBlendOp             /**< The blend operation for the alpha component. */
	ColorWriteMask       GPUColorComponentFlags /**< A bitmask specifying which of the RGBA components are enabled for writing. Writes to all channels if enable_color_write_mask is false. */
	EnableBlend          bool                   /**< Whether blending is enabled for the color target. */
	EnableColorWriteMask bool                   /**< Whether the color write mask is enabled. */
	Padding1             uint8
	Padding2             uint8
}

/**
 * A structure specifying code and metadata for creating a shader object.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
type GPUShaderCreateInfo struct {
	CodeSize           size_t          /**< The size in bytes of the code pointed to. */
	Code               *byte           /**< A pointer to shader code. */
	EntryPoint         string          /**< A pointer to a null-terminated UTF-8 string specifying the entry point function name for the shader. */
	Format             GPUShaderFormat /**< The format of the shader code. */
	Stage              GPUShaderStage  /**< The stage the shader program corresponds to. */
	NumSamplers        uint32          /**< The number of samplers defined in the shader. */
	NumStorageTextures uint32          /**< The number of storage textures defined in the shader. */
	NumStorageBuffers  uint32          /**< The number of storage buffers defined in the shader. */
	NumUniformBuffers  uint32          /**< The number of uniform buffers defined in the shader. */

	Props PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

/**
 * A structure specifying the parameters of a texture.
 *
 * Usage flags can be bitwise OR'd together for combinations of usages. Note
 * that certain usage combinations are invalid, for example SAMPLER and
 * GRAPHICS_STORAGE.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUTexture
 * \sa SDL_GPUTextureType
 * \sa SDL_GPUTextureFormat
 * \sa SDL_GPUTextureUsageFlags
 * \sa SDL_GPUSampleCount
 */
type GPUTextureCreateInfo struct {
	Type              GPUTextureType       /**< The base dimensionality of the texture. */
	Format            GPUTextureFormat     /**< The pixel format of the texture. */
	Usage             GPUTextureUsageFlags /**< How the texture is intended to be used by the client. */
	Width             uint32               /**< The width of the texture. */
	Height            uint32               /**< The height of the texture. */
	LayerCountOrDepth uint32               /**< The layer count or depth of the texture. This value is treated as a layer count on 2D array textures, and as a depth value on 3D textures. */
	NumLevels         uint32               /**< The number of mip levels in the texture. */
	SampleCount       GPUSampleCount       /**< The number of samples per texel. Only applies if the texture is used as a render target. */

	Props PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

/**
 * A structure specifying the parameters of a buffer.
 *
 * Usage flags can be bitwise OR'd together for combinations of usages. Note
 * that certain combinations are invalid, for example VERTEX and INDEX.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUBuffer
 * \sa SDL_GPUBufferUsageFlags
 */
type GPUBufferCreateInfo struct {
	Usage GPUBufferUsageFlags /**< How the buffer is intended to be used by the client. */
	Size  uint32              /**< The size in bytes of the buffer. */

	Props PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

/**
 * A structure specifying the parameters of a transfer buffer.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUTransferBuffer
 */
type GPUTransferBufferCreateInfo struct {
	Usage GPUTransferBufferUsage /**< How the transfer buffer is intended to be used by the client. */
	Size  uint32                 /**< The size in bytes of the transfer buffer. */

	Props PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

/* Pipeline state structures */

/**
 * A structure specifying the parameters of the graphics pipeline rasterizer
 * state.
 *
 * Note that SDL_GPU_FILLMODE_LINE is not supported on many Android devices.
 * For those devices, the fill mode will automatically fall back to FILL.
 *
 * Also note that the D3D12 driver will enable depth clamping even if
 * enable_depth_clip is true. If you need this clamp+clip behavior, consider
 * enabling depth clip and then manually clamping depth in your fragment
 * shaders on Metal and Vulkan.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GPUGraphicsPipelineCreateInfo
 */
type GPURasterizerState struct {
	FillMode                GPUFillMode  /**< Whether polygons will be filled in or drawn as lines. */
	CullMode                GPUCullMode  /**< The facing direction in which triangles will be culled. */
	FrontFace               GPUFrontFace /**< The vertex winding that will cause a triangle to be determined as front-facing. */
	DepthBiasConstantFactor float32      /**< A scalar factor controlling the depth value added to each fragment. */
	DepthBiasClamp          float32      /**< The maximum depth bias of a fragment. */
	DepthBiasSlopeFactor    float32      /**< A scalar factor applied to a fragment's slope in depth calculations. */
	EnableDepthBias         bool         /**< true to bias fragment depth values. */
	EnableDepthClip         bool         /**< true to enable depth clip, false to enable depth clamp. */
	Padding1                uint8
	Padding2                uint8
}

/**
 * A structure specifying the parameters of the graphics pipeline multisample
 * state.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GPUGraphicsPipelineCreateInfo
 */
type GPUMultisampleState struct {
	SampleCount           GPUSampleCount /**< The number of samples to be used in rasterization. */
	SampleMask            uint32         /**< Reserved for future use. Must be set to 0. */
	EnableMask            bool           /**< Reserved for future use. Must be set to false. */
	EnableAlphaToCoverage bool           /**< true enables the alpha-to-coverage feature. */
	Padding2              uint8
	Padding3              uint8
}

/**
 * A structure specifying the parameters of the graphics pipeline depth
 * stencil state.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GPUGraphicsPipelineCreateInfo
 */
type GPUDepthStencilState struct {
	CompareOp         GPUCompareOp      /**< The comparison operator used for depth testing. */
	BackStencilState  GPUStencilOpState /**< The stencil op state for back-facing triangles. */
	FrontStencilState GPUStencilOpState /**< The stencil op state for front-facing triangles. */
	CompareMask       uint8             /**< Selects the bits of the stencil values participating in the stencil test. */
	WriteMask         uint8             /**< Selects the bits of the stencil values updated by the stencil test. */
	EnableDepthTest   bool              /**< true enables the depth test. */
	EnableDepthWrite  bool              /**< true enables depth writes. Depth writes are always disabled when enable_depth_test is false. */
	EnableStencilTest bool              /**< true enables the stencil test. */
	Padding1          uint8
	Padding2          uint8
	Padding3          uint8
}

/**
 * A structure specifying the parameters of color targets used in a graphics
 * pipeline.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GPUGraphicsPipelineTargetInfo
 */
type GPUColorTargetDescription struct {
	Format     GPUTextureFormat         /**< The pixel format of the texture to be used as a color target. */
	BlendState GPUColorTargetBlendState /**< The blend state to be used for the color target. */
}

/**
 * A structure specifying the descriptions of render targets used in a
 * graphics pipeline.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_GPUGraphicsPipelineCreateInfo
 * \sa SDL_GPUColorTargetDescription
 * \sa SDL_GPUTextureFormat
 */
type GPUGraphicsPipelineTargetInfo struct {
	ColorTargetDescriptions *GPUColorTargetDescription /**< A pointer to an array of color target descriptions. */
	NumColorTargets         uint32                     /**< The number of color target descriptions in the above array. */
	DepthStencilFormat      GPUTextureFormat           /**< The pixel format of the depth-stencil target. Ignored if has_depth_stencil_target is false. */
	HasDepthStencilTarget   bool                       /**< true specifies that the pipeline uses a depth-stencil target. */
	Padding1                uint8
	Padding2                uint8
	Padding3                uint8
}

/**
 * A structure specifying the parameters of a graphics pipeline state.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 * \sa SDL_GPUShader
 * \sa SDL_GPUVertexInputState
 * \sa SDL_GPUPrimitiveType
 * \sa SDL_GPURasterizerState
 * \sa SDL_GPUMultisampleState
 * \sa SDL_GPUDepthStencilState
 * \sa SDL_GPUGraphicsPipelineTargetInfo
 */
type GPUGraphicsPipelineCreateInfo struct {
	VertexShader      GPUShader                     /**< The vertex shader used by the graphics pipeline. */
	FragmentShader    GPUShader                     /**< The fragment shader used by the graphics pipeline. */
	VertexInputState  GPUVertexInputState           /**< The vertex layout of the graphics pipeline. */
	PrimitiveType     GPUPrimitiveType              /**< The primitive topology of the graphics pipeline. */
	RasterizerState   GPURasterizerState            /**< The rasterizer state of the graphics pipeline. */
	MultisampleState  GPUMultisampleState           /**< The multisample state of the graphics pipeline. */
	DepthStencilState GPUDepthStencilState          /**< The depth-stencil state of the graphics pipeline. */
	TargetInfo        GPUGraphicsPipelineTargetInfo /**< Formats and blend modes for the render targets of the graphics pipeline. */

	Props PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

/**
 * A structure specifying the parameters of a compute pipeline state.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUComputePipeline
 * \sa SDL_GPUShaderFormat
 */
type GPUComputePipelineCreateInfo struct {
	CodeSize                    size_t          /**< The size in bytes of the compute shader code pointed to. */
	Code                        *byte           /**< A pointer to compute shader code. */
	EntryPoint                  string          /**< A pointer to a null-terminated UTF-8 string specifying the entry point function name for the shader. */
	Format                      GPUShaderFormat /**< The format of the compute shader code. */
	NumSamplers                 uint32          /**< The number of samplers defined in the shader. */
	NumReadonlyStorageTextures  uint32          /**< The number of readonly storage textures defined in the shader. */
	NumReadonlyStorageBuffers   uint32          /**< The number of readonly storage buffers defined in the shader. */
	NumReadwriteStorageTextures uint32          /**< The number of read-write storage textures defined in the shader. */
	NumReadwriteStorageBuffers  uint32          /**< The number of read-write storage buffers defined in the shader. */
	NumUniformBuffers           uint32          /**< The number of uniform buffers defined in the shader. */
	ThreadCountX                uint32          /**< The number of threads in the X dimension. This should match the value in the shader. */
	ThreadCountY                uint32          /**< The number of threads in the Y dimension. This should match the value in the shader. */
	ThreadCountZ                uint32          /**< The number of threads in the Z dimension. This should match the value in the shader. */

	Props PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

/**
 * A structure specifying the parameters of a color target used by a render
 * pass.
 *
 * The load_op field determines what is done with the texture at the beginning
 * of the render pass.
 *
 * - LOAD: Loads the data currently in the texture. Not recommended for
 *   multisample textures as it requires significant memory bandwidth.
 * - CLEAR: Clears the texture to a single color.
 * - DONT_CARE: The driver will do whatever it wants with the texture memory.
 *   This is a good option if you know that every single pixel will be touched
 *   in the render pass.
 *
 * The store_op field determines what is done with the color results of the
 * render pass.
 *
 * - STORE: Stores the results of the render pass in the texture. Not
 *   recommended for multisample textures as it requires significant memory
 *   bandwidth.
 * - DONT_CARE: The driver will do whatever it wants with the texture memory.
 *   This is often a good option for depth/stencil textures.
 * - RESOLVE: Resolves a multisample texture into resolve_texture, which must
 *   have a sample count of 1. Then the driver may discard the multisample
 *   texture memory. This is the most performant method of resolving a
 *   multisample target.
 * - RESOLVE_AND_STORE: Resolves a multisample texture into the
 *   resolve_texture, which must have a sample count of 1. Then the driver
 *   stores the multisample texture's contents. Not recommended as it requires
 *   significant memory bandwidth.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_BeginGPURenderPass
 */
type GPUColorTargetInfo struct {
	Texture             GPUTexture /**< The texture that will be used as a color target by a render pass. */
	MipLevel            uint32     /**< The mip level to use as a color target. */
	LayerOrDepthPlane   uint32     /**< The layer index or depth plane to use as a color target. This value is treated as a layer index on 2D array and cube textures, and as a depth plane on 3D textures. */
	ClearColor          FColor     /**< The color to clear the color target to at the start of the render pass. Ignored if SDL_GPU_LOADOP_CLEAR is not used. */
	LoadOp              GPULoadOp  /**< What is done with the contents of the color target at the beginning of the render pass. */
	StoreOp             GPUStoreOp /**< What is done with the results of the render pass. */
	ResolveTexture      GPUTexture /**< The texture that will receive the results of a multisample resolve operation. Ignored if a RESOLVE* store_op is not used. */
	ResolveMipLevel     uint32     /**< The mip level of the resolve texture to use for the resolve operation. Ignored if a RESOLVE* store_op is not used. */
	ResolveLayer        uint32     /**< The layer index of the resolve texture to use for the resolve operation. Ignored if a RESOLVE* store_op is not used. */
	Cycle               bool       /**< true cycles the texture if the texture is bound and load_op is not LOAD */
	CycleResolveTexture bool       /**< true cycles the resolve texture if the resolve texture is bound. Ignored if a RESOLVE* store_op is not used. */
	Padding1            uint8
	Padding2            uint8
}

/**
 * A structure specifying the parameters of a depth-stencil target used by a
 * render pass.
 *
 * The load_op field determines what is done with the depth contents of the
 * texture at the beginning of the render pass.
 *
 * - LOAD: Loads the depth values currently in the texture.
 * - CLEAR: Clears the texture to a single depth.
 * - DONT_CARE: The driver will do whatever it wants with the memory. This is
 *   a good option if you know that every single pixel will be touched in the
 *   render pass.
 *
 * The store_op field determines what is done with the depth results of the
 * render pass.
 *
 * - STORE: Stores the depth results in the texture.
 * - DONT_CARE: The driver will do whatever it wants with the depth results.
 *   This is often a good option for depth/stencil textures that don't need to
 *   be reused again.
 *
 * The stencil_load_op field determines what is done with the stencil contents
 * of the texture at the beginning of the render pass.
 *
 * - LOAD: Loads the stencil values currently in the texture.
 * - CLEAR: Clears the stencil values to a single value.
 * - DONT_CARE: The driver will do whatever it wants with the memory. This is
 *   a good option if you know that every single pixel will be touched in the
 *   render pass.
 *
 * The stencil_store_op field determines what is done with the stencil results
 * of the render pass.
 *
 * - STORE: Stores the stencil results in the texture.
 * - DONT_CARE: The driver will do whatever it wants with the stencil results.
 *   This is often a good option for depth/stencil textures that don't need to
 *   be reused again.
 *
 * Note that depth/stencil targets do not support multisample resolves.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_BeginGPURenderPass
 */
type GPUDepthStencilTargetInfo struct {
	Texture        GPUTexture /**< The texture that will be used as the depth stencil target by the render pass. */
	ClearDepth     float32    /**< The value to clear the depth component to at the beginning of the render pass. Ignored if SDL_GPU_LOADOP_CLEAR is not used. */
	LoadOp         GPULoadOp  /**< What is done with the depth contents at the beginning of the render pass. */
	StoreOp        GPUStoreOp /**< What is done with the depth results of the render pass. */
	StencilLoadOp  GPULoadOp  /**< What is done with the stencil contents at the beginning of the render pass. */
	StencilStoreOp GPUStoreOp /**< What is done with the stencil results of the render pass. */
	Cycle          bool       /**< true cycles the texture if the texture is bound and any load ops are not LOAD */
	ClearStencil   uint8      /**< The value to clear the stencil component to at the beginning of the render pass. Ignored if SDL_GPU_LOADOP_CLEAR is not used. */
	Padding1       uint8
	Padding2       uint8
}

/**
 * A structure containing parameters for a blit command.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_BlitGPUTexture
 */
type GPUBlitInfo struct {
	Source      GPUBlitRegion /**< The source region for the blit. */
	Destination GPUBlitRegion /**< The destination region for the blit. */
	LoadOp      GPULoadOp     /**< What is done with the contents of the destination before the blit. */
	ClearColor  FColor        /**< The color to clear the destination region to before the blit. Ignored if load_op is not SDL_GPU_LOADOP_CLEAR. */
	FlipMode    FlipMode      /**< The flip mode for the source region. */
	Filter      GPUFilter     /**< The filter mode used when blitting. */
	Cycle       bool          /**< true cycles the destination texture if it is already bound. */
	Padding1    uint8
	Padding2    uint8
	Padding3    uint8
}

/* Binding structs */

/**
 * A structure specifying parameters in a buffer binding call.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_BindGPUVertexBuffers
 * \sa SDL_BindGPUIndexBuffer
 */
type GPUBufferBinding struct {
	Buffer GPUBuffer /**< The buffer to bind. Must have been created with SDL_GPU_BUFFERUSAGE_VERTEX for SDL_BindGPUVertexBuffers, or SDL_GPU_BUFFERUSAGE_INDEX for SDL_BindGPUIndexBuffer. */
	Offset uint32    /**< The starting byte of the data to bind in the buffer. */
}

/**
 * A structure specifying parameters in a sampler binding call.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_BindGPUVertexSamplers
 * \sa SDL_BindGPUFragmentSamplers
 */
type GPUTextureSamplerBinding struct {
	Texture GPUTexture /**< The texture to bind. Must have been created with SDL_GPU_TEXTUREUSAGE_SAMPLER. */
	Sampler GPUSampler /**< The sampler to bind. */
}

/**
 * A structure specifying parameters related to binding buffers in a compute
 * pass.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_BeginGPUComputePass
 */
type GPUStorageBufferReadWriteBinding struct {
	Buffer   GPUBuffer /**< The buffer to bind. Must have been created with SDL_GPU_BUFFERUSAGE_COMPUTE_STORAGE_WRITE. */
	Cycle    bool      /**< true cycles the buffer if it is already bound. */
	Padding1 uint8
	Padding2 uint8
	Padding3 uint8
}

/**
 * A structure specifying parameters related to binding textures in a compute
 * pass.
 *
 * \since This struct is available since SDL 3.2.0.
 *
 * \sa SDL_BeginGPUComputePass
 */
type GPUStorageTextureReadWriteBinding struct {
	Texture  GPUTexture /**< The texture to bind. Must have been created with SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE or SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_SIMULTANEOUS_READ_WRITE. */
	MipLevel uint32     /**< The mip level index to bind. */
	Layer    uint32     /**< The layer index to bind. */
	Cycle    bool       /**< true cycles the texture if it is already bound. */
	Padding1 uint8
	Padding2 uint8
	Padding3 uint8
}

/* Functions */

/* Device */

/**
 * Checks for GPU runtime support.
 *
 * \param format_flags a bitflag indicating which shader formats the app is
 *                     able to provide.
 * \param name the preferred GPU driver, or NULL to let SDL pick the optimal
 *             driver.
 * \returns true if supported, false otherwise.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUDevice
 */
//go:sdl3extern
var GPUSupportsShaderFormats func(format_flags GPUShaderFormat, name string) bool

/**
 * Checks for GPU runtime support.
 *
 * \param props the properties to use.
 * \returns true if supported, false otherwise.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUDeviceWithProperties
 */
//go:sdl3extern
var GPUSupportsProperties func(props PropertiesID) bool

/**
 * Creates a GPU context.
 *
 * The GPU driver name can be one of the following:
 *
 * - "vulkan": [Vulkan](CategoryGPU#vulkan)
 * - "direct3d12": [D3D12](CategoryGPU#d3d12)
 * - "metal": [Metal](CategoryGPU#metal)
 * - NULL: let SDL pick the optimal driver
 *
 * \param format_flags a bitflag indicating which shader formats the app is
 *                     able to provide.
 * \param debug_mode enable debug mode properties and validations.
 * \param name the preferred GPU driver, or NULL to let SDL pick the optimal
 *             driver.
 * \returns a GPU context on success or NULL on failure; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUDeviceWithProperties
 * \sa SDL_GetGPUShaderFormats
 * \sa SDL_GetGPUDeviceDriver
 * \sa SDL_DestroyGPUDevice
 * \sa SDL_GPUSupportsShaderFormats
 */
//go:sdl3extern
var CreateGPUDevice func(
	format_flags GPUShaderFormat,
	debug_mode bool,
	name string) GPUDevice

/**
 * Creates a GPU context.
 *
 * These are the supported properties:
 *
 * - `SDL_PROP_GPU_DEVICE_CREATE_DEBUGMODE_BOOLEAN`: enable debug mode
 *   properties and validations, defaults to true.
 * - `SDL_PROP_GPU_DEVICE_CREATE_PREFERLOWPOWER_BOOLEAN`: enable to prefer
 *   energy efficiency over maximum GPU performance, defaults to false.
 * - `SDL_PROP_GPU_DEVICE_CREATE_VERBOSE_BOOLEAN`: enable to automatically log
 *   useful debug information on device creation, defaults to true.
 * - `SDL_PROP_GPU_DEVICE_CREATE_NAME_STRING`: the name of the GPU driver to
 *   use, if a specific one is desired.
 *
 * These are the current shader format properties:
 *
 * - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_PRIVATE_BOOLEAN`: The app is able to
 *   provide shaders for an NDA platform.
 * - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_SPIRV_BOOLEAN`: The app is able to
 *   provide SPIR-V shaders if applicable.
 * - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_DXBC_BOOLEAN`: The app is able to
 *   provide DXBC shaders if applicable
 * - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_DXIL_BOOLEAN`: The app is able to
 *   provide DXIL shaders if applicable.
 * - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_MSL_BOOLEAN`: The app is able to
 *   provide MSL shaders if applicable.
 * - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_METALLIB_BOOLEAN`: The app is able to
 *   provide Metal shader libraries if applicable.
 *
 * With the D3D12 renderer:
 *
 * - `SDL_PROP_GPU_DEVICE_CREATE_D3D12_SEMANTIC_NAME_STRING`: the prefix to
 *   use for all vertex semantics, default is "TEXCOORD".
 *
 * \param props the properties to use.
 * \returns a GPU context on success or NULL on failure; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGPUShaderFormats
 * \sa SDL_GetGPUDeviceDriver
 * \sa SDL_DestroyGPUDevice
 * \sa SDL_GPUSupportsProperties
 */
//go:sdl3extern
var CreateGPUDeviceWithProperties func(
	props PropertiesID) GPUDevice

const (
	PROP_GPU_DEVICE_CREATE_DEBUGMODE_BOOLEAN          PropertyName = "SDL.gpu.device.create.debugmode"
	PROP_GPU_DEVICE_CREATE_PREFERLOWPOWER_BOOLEAN     PropertyName = "SDL.gpu.device.create.preferlowpower"
	PROP_GPU_DEVICE_CREATE_VERBOSE_BOOLEAN            PropertyName = "SDL.gpu.device.create.verbose"
	PROP_GPU_DEVICE_CREATE_NAME_STRING                PropertyName = "SDL.gpu.device.create.name"
	PROP_GPU_DEVICE_CREATE_SHADERS_PRIVATE_BOOLEAN    PropertyName = "SDL.gpu.device.create.shaders.private"
	PROP_GPU_DEVICE_CREATE_SHADERS_SPIRV_BOOLEAN      PropertyName = "SDL.gpu.device.create.shaders.spirv"
	PROP_GPU_DEVICE_CREATE_SHADERS_DXBC_BOOLEAN       PropertyName = "SDL.gpu.device.create.shaders.dxbc"
	PROP_GPU_DEVICE_CREATE_SHADERS_DXIL_BOOLEAN       PropertyName = "SDL.gpu.device.create.shaders.dxil"
	PROP_GPU_DEVICE_CREATE_SHADERS_MSL_BOOLEAN        PropertyName = "SDL.gpu.device.create.shaders.msl"
	PROP_GPU_DEVICE_CREATE_SHADERS_METALLIB_BOOLEAN   PropertyName = "SDL.gpu.device.create.shaders.metallib"
	PROP_GPU_DEVICE_CREATE_D3D12_SEMANTIC_NAME_STRING PropertyName = "SDL.gpu.device.create.d3d12.semantic"
)

/**
 * Destroys a GPU context previously returned by SDL_CreateGPUDevice.
 *
 * \param device a GPU Context to destroy.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUDevice
 */
//go:sdl3extern
var DestroyGPUDevice func(device GPUDevice)

/**
 * Get the number of GPU drivers compiled into SDL.
 *
 * \returns the number of built in GPU drivers.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetGPUDriver
 */
//go:sdl3extern
var GetNumGPUDrivers func() int32

/**
 * Get the name of a built in GPU driver.
 *
 * The GPU drivers are presented in the order in which they are normally
 * checked during initialization.
 *
 * The names of drivers are all simple, low-ASCII identifiers, like "vulkan",
 * "metal" or "direct3d12". These never have Unicode characters, and are not
 * meant to be proper names.
 *
 * \param index the index of a GPU driver.
 * \returns the name of the GPU driver with the given **index**.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GetNumGPUDrivers
 */
//go:sdl3extern
var GetGPUDriver func(index int32) string

/**
 * Returns the name of the backend used to create this GPU context.
 *
 * \param device a GPU context to query.
 * \returns the name of the device's driver, or NULL on error.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGPUDeviceDriver func(device GPUDevice) string

/**
 * Returns the supported shader formats for this GPU context.
 *
 * \param device a GPU context to query.
 * \returns a bitflag indicating which shader formats the driver is able to
 *          consume.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGPUShaderFormats func(device GPUDevice) GPUShaderFormat

/**
 * Get the properties associated with a GPU device.
 *
 * All properties are optional and may differ between GPU backends and SDL
 * versions.
 *
 * The following properties are provided by SDL:
 *
 * `SDL_PROP_GPU_DEVICE_NAME_STRING`: Contains the name of the underlying
 * device as reported by the system driver. This string has no standardized
 * format, is highly inconsistent between hardware devices and drivers, and is
 * able to change at any time. Do not attempt to parse this string as it is
 * bound to fail at some point in the future when system drivers are updated,
 * new hardware devices are introduced, or when SDL adds new GPU backends or
 * modifies existing ones.
 *
 * Strings that have been found in the wild include:
 *
 * - GTX 970
 * - GeForce GTX 970
 * - NVIDIA GeForce GTX 970
 * - Microsoft Direct3D12 (NVIDIA GeForce GTX 970)
 * - NVIDIA Graphics Device
 * - GeForce GPU
 * - P106-100
 * - AMD 15D8:C9
 * - AMD Custom GPU 0405
 * - AMD Radeon (TM) Graphics
 * - ASUS Radeon RX 470 Series
 * - Intel(R) Arc(tm) A380 Graphics (DG2)
 * - Virtio-GPU Venus (NVIDIA TITAN V)
 * - SwiftShader Device (LLVM 16.0.0)
 * - llvmpipe (LLVM 15.0.4, 256 bits)
 * - Microsoft Basic Render Driver
 * - unknown device
 *
 * The above list shows that the same device can have different formats, the
 * vendor name may or may not appear in the string, the included vendor name
 * may not be the vendor of the chipset on the device, some manufacturers
 * include pseudo-legal marks while others don't, some devices may not use a
 * marketing name in the string, the device string may be wrapped by the name
 * of a translation interface, the device may be emulated in software, or the
 * string may contain generic text that does not identify the device at all.
 *
 * `SDL_PROP_GPU_DEVICE_DRIVER_NAME_STRING`: Contains the self-reported name
 * of the underlying system driver.
 *
 * Strings that have been found in the wild include:
 *
 * - Intel Corporation
 * - Intel open-source Mesa driver
 * - Qualcomm Technologies Inc. Adreno Vulkan Driver
 * - MoltenVK
 * - Mali-G715
 * - venus
 *
 * `SDL_PROP_GPU_DEVICE_DRIVER_VERSION_STRING`: Contains the self-reported
 * version of the underlying system driver. This is a relatively short version
 * string in an unspecified format. If SDL_PROP_GPU_DEVICE_DRIVER_INFO_STRING
 * is available then that property should be preferred over this one as it may
 * contain additional information that is useful for identifying the exact
 * driver version used.
 *
 * Strings that have been found in the wild include:
 *
 * - 53.0.0
 * - 0.405.2463
 * - 32.0.15.6614
 *
 * `SDL_PROP_GPU_DEVICE_DRIVER_INFO_STRING`: Contains the detailed version
 * information of the underlying system driver as reported by the driver. This
 * is an arbitrary string with no standardized format and it may contain
 * newlines. This property should be preferred over
 * SDL_PROP_GPU_DEVICE_DRIVER_VERSION_STRING if it is available as it usually
 * contains the same information but in a format that is easier to read.
 *
 * Strings that have been found in the wild include:
 *
 * - 101.6559
 * - 1.2.11
 * - Mesa 21.2.2 (LLVM 12.0.1)
 * - Mesa 22.2.0-devel (git-f226222 2022-04-14 impish-oibaf-ppa)
 * - v1.r53p0-00eac0.824c4f31403fb1fbf8ee1042422c2129
 *
 * This string has also been observed to be a multiline string (which has a
 * trailing newline):
 *
 * ```
 * Driver Build: 85da404, I46ff5fc46f, 1606794520
 * Date: 11/30/20
 * Compiler Version: EV031.31.04.01
 * Driver Branch: promo490_3_Google
 * ```
 *
 * \param device a GPU context to query.
 * \returns a valid property ID on success or 0 on failure; call
 *          SDL_GetError() for more information.
 *
 * \threadsafety It is safe to call this function from any thread.
 *
 * \since This function is available since SDL 3.4.0.
 */
//go:sdl3extern
var GetGPUDeviceProperties func(device GPUDevice) PropertiesID

const (
	PROP_GPU_DEVICE_NAME_STRING           PropertyName = "SDL.gpu.device.name"
	PROP_GPU_DEVICE_DRIVER_NAME_STRING    PropertyName = "SDL.gpu.device.driver_name"
	PROP_GPU_DEVICE_DRIVER_VERSION_STRING PropertyName = "SDL.gpu.device.driver_version"
	PROP_GPU_DEVICE_DRIVER_INFO_STRING    PropertyName = "SDL.gpu.device.driver_info"
)

/* State Creation */

/**
 * Creates a pipeline object to be used in a compute workflow.
 *
 * Shader resource bindings must be authored to follow a particular order
 * depending on the shader format.
 *
 * For SPIR-V shaders, use the following resource sets:
 *
 * - 0: Sampled textures, followed by read-only storage textures, followed by
 *   read-only storage buffers
 * - 1: Read-write storage textures, followed by read-write storage buffers
 * - 2: Uniform buffers
 *
 * For DXBC and DXIL shaders, use the following register order:
 *
 * - (t[n], space0): Sampled textures, followed by read-only storage textures,
 *   followed by read-only storage buffers
 * - (u[n], space1): Read-write storage textures, followed by read-write
 *   storage buffers
 * - (b[n], space2): Uniform buffers
 *
 * For MSL/metallib, use the following order:
 *
 * - [[buffer]]: Uniform buffers, followed by read-only storage buffers,
 *   followed by read-write storage buffers
 * - [[texture]]: Sampled textures, followed by read-only storage textures,
 *   followed by read-write storage textures
 *
 * There are optional properties that can be provided through `props`. These
 * are the supported properties:
 *
 * - `SDL_PROP_GPU_COMPUTEPIPELINE_CREATE_NAME_STRING`: a name that can be
 *   displayed in debugging tools.
 *
 * \param device a GPU Context.
 * \param createinfo a struct describing the state of the compute pipeline to
 *                   create.
 * \returns a compute pipeline object on success, or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_BindGPUComputePipeline
 * \sa SDL_ReleaseGPUComputePipeline
 */
//go:sdl3extern
var CreateGPUComputePipeline func(
	device GPUDevice,
	createinfo *GPUComputePipelineCreateInfo) GPUComputePipeline

const PROP_GPU_COMPUTEPIPELINE_CREATE_NAME_STRING PropertyName = "SDL.gpu.computepipeline.create.name"

/**
 * Creates a pipeline object to be used in a graphics workflow.
 *
 * There are optional properties that can be provided through `props`. These
 * are the supported properties:
 *
 * - `SDL_PROP_GPU_GRAPHICSPIPELINE_CREATE_NAME_STRING`: a name that can be
 *   displayed in debugging tools.
 *
 * \param device a GPU Context.
 * \param createinfo a struct describing the state of the graphics pipeline to
 *                   create.
 * \returns a graphics pipeline object on success, or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 * \sa SDL_BindGPUGraphicsPipeline
 * \sa SDL_ReleaseGPUGraphicsPipeline
 */
//go:sdl3extern
var CreateGPUGraphicsPipeline func(
	device GPUDevice,
	createinfo *GPUGraphicsPipelineCreateInfo) GPUGraphicsPipeline

const PROP_GPU_GRAPHICSPIPELINE_CREATE_NAME_STRING PropertyName = "SDL.gpu.graphicspipeline.create.name"

/**
 * Creates a sampler object to be used when binding textures in a graphics
 * workflow.
 *
 * There are optional properties that can be provided through `props`. These
 * are the supported properties:
 *
 * - `SDL_PROP_GPU_SAMPLER_CREATE_NAME_STRING`: a name that can be displayed
 *   in debugging tools.
 *
 * \param device a GPU Context.
 * \param createinfo a struct describing the state of the sampler to create.
 * \returns a sampler object on success, or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_BindGPUVertexSamplers
 * \sa SDL_BindGPUFragmentSamplers
 * \sa SDL_ReleaseGPUSampler
 */
//go:sdl3extern
var CreateGPUSampler func(
	device GPUDevice,
	createinfo *GPUSamplerCreateInfo) GPUSampler

const PROP_GPU_SAMPLER_CREATE_NAME_STRING PropertyName = "SDL.gpu.sampler.create.name"

/**
 * Creates a shader to be used when creating a graphics pipeline.
 *
 * Shader resource bindings must be authored to follow a particular order
 * depending on the shader format.
 *
 * For SPIR-V shaders, use the following resource sets:
 *
 * For vertex shaders:
 *
 * - 0: Sampled textures, followed by storage textures, followed by storage
 *   buffers
 * - 1: Uniform buffers
 *
 * For fragment shaders:
 *
 * - 2: Sampled textures, followed by storage textures, followed by storage
 *   buffers
 * - 3: Uniform buffers
 *
 * For DXBC and DXIL shaders, use the following register order:
 *
 * For vertex shaders:
 *
 * - (t[n], space0): Sampled textures, followed by storage textures, followed
 *   by storage buffers
 * - (s[n], space0): Samplers with indices corresponding to the sampled
 *   textures
 * - (b[n], space1): Uniform buffers
 *
 * For pixel shaders:
 *
 * - (t[n], space2): Sampled textures, followed by storage textures, followed
 *   by storage buffers
 * - (s[n], space2): Samplers with indices corresponding to the sampled
 *   textures
 * - (b[n], space3): Uniform buffers
 *
 * For MSL/metallib, use the following order:
 *
 * - [[texture]]: Sampled textures, followed by storage textures
 * - [[sampler]]: Samplers with indices corresponding to the sampled textures
 * - [[buffer]]: Uniform buffers, followed by storage buffers. Vertex buffer 0
 *   is bound at [[buffer(14)]], vertex buffer 1 at [[buffer(15)]], and so on.
 *   Rather than manually authoring vertex buffer indices, use the
 *   [[stage_in]] attribute which will automatically use the vertex input
 *   information from the SDL_GPUGraphicsPipeline.
 *
 * Shader semantics other than system-value semantics do not matter in D3D12
 * and for ease of use the SDL implementation assumes that non system-value
 * semantics will all be TEXCOORD. If you are using HLSL as the shader source
 * language, your vertex semantics should start at TEXCOORD0 and increment
 * like so: TEXCOORD1, TEXCOORD2, etc. If you wish to change the semantic
 * prefix to something other than TEXCOORD you can use
 * SDL_PROP_GPU_DEVICE_CREATE_D3D12_SEMANTIC_NAME_STRING with
 * SDL_CreateGPUDeviceWithProperties().
 *
 * There are optional properties that can be provided through `props`. These
 * are the supported properties:
 *
 * - `SDL_PROP_GPU_SHADER_CREATE_NAME_STRING`: a name that can be displayed in
 *   debugging tools.
 *
 * \param device a GPU Context.
 * \param createinfo a struct describing the state of the shader to create.
 * \returns a shader object on success, or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUGraphicsPipeline
 * \sa SDL_ReleaseGPUShader
 */
//go:sdl3extern
var CreateGPUShader func(
	device GPUDevice,
	createinfo *GPUShaderCreateInfo) GPUShader

const PROP_GPU_SHADER_CREATE_NAME_STRING PropertyName = "SDL.gpu.shader.create.name"

/**
 * Creates a texture object to be used in graphics or compute workflows.
 *
 * The contents of this texture are undefined until data is written to the
 * texture.
 *
 * Note that certain combinations of usage flags are invalid. For example, a
 * texture cannot have both the SAMPLER and GRAPHICS_STORAGE_READ flags.
 *
 * If you request a sample count higher than the hardware supports, the
 * implementation will automatically fall back to the highest available sample
 * count.
 *
 * There are optional properties that can be provided through
 * SDL_GPUTextureCreateInfo's `props`. These are the supported properties:
 *
 * - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_R_Float`: (Direct3D 12 only) if
 *   the texture usage is SDL_GPU_TEXTUREUSAGE_COLOR_TARGET, clear the texture
 *   to a color with this red intensity. Defaults to zero.
 * - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_G_Float`: (Direct3D 12 only) if
 *   the texture usage is SDL_GPU_TEXTUREUSAGE_COLOR_TARGET, clear the texture
 *   to a color with this green intensity. Defaults to zero.
 * - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_B_Float`: (Direct3D 12 only) if
 *   the texture usage is SDL_GPU_TEXTUREUSAGE_COLOR_TARGET, clear the texture
 *   to a color with this blue intensity. Defaults to zero.
 * - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_A_Float`: (Direct3D 12 only) if
 *   the texture usage is SDL_GPU_TEXTUREUSAGE_COLOR_TARGET, clear the texture
 *   to a color with this alpha intensity. Defaults to zero.
 * - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_DEPTH_Float`: (Direct3D 12 only)
 *   if the texture usage is SDL_GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET, clear
 *   the texture to a depth of this value. Defaults to zero.
 * - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_STENCIL_NUMBER`: (Direct3D 12
 *   only) if the texture usage is SDL_GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET,
 *   clear the texture to a stencil of this uint8 value. Defaults to zero.
 * - `SDL_PROP_GPU_TEXTURE_CREATE_NAME_STRING`: a name that can be displayed
 *   in debugging tools.
 *
 * \param device a GPU Context.
 * \param createinfo a struct describing the state of the texture to create.
 * \returns a texture object on success, or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_UploadToGPUTexture
 * \sa SDL_DownloadFromGPUTexture
 * \sa SDL_BindGPUVertexSamplers
 * \sa SDL_BindGPUVertexStorageTextures
 * \sa SDL_BindGPUFragmentSamplers
 * \sa SDL_BindGPUFragmentStorageTextures
 * \sa SDL_BindGPUComputeStorageTextures
 * \sa SDL_BlitGPUTexture
 * \sa SDL_ReleaseGPUTexture
 * \sa SDL_GPUTextureSupportsFormat
 */
//go:sdl3extern
var CreateGPUTexture func(
	device GPUDevice,
	createinfo *GPUTextureCreateInfo) GPUTexture

const PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_R_Float PropertyName = "SDL.gpu.texture.create.d3d12.clear.r"
const PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_G_Float PropertyName = "SDL.gpu.texture.create.d3d12.clear.g"
const PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_B_Float PropertyName = "SDL.gpu.texture.create.d3d12.clear.b"
const PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_A_Float PropertyName = "SDL.gpu.texture.create.d3d12.clear.a"
const PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_DEPTH_Float PropertyName = "SDL.gpu.texture.create.d3d12.clear.depth"
const PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_STENCIL_NUMBER PropertyName = "SDL.gpu.texture.create.d3d12.clear.stencil"
const PROP_GPU_TEXTURE_CREATE_NAME_STRING PropertyName = "SDL.gpu.texture.create.name"

/**
 * Creates a buffer object to be used in graphics or compute workflows.
 *
 * The contents of this buffer are undefined until data is written to the
 * buffer.
 *
 * Note that certain combinations of usage flags are invalid. For example, a
 * buffer cannot have both the VERTEX and INDEX flags.
 *
 * If you use a STORAGE flag, the data in the buffer must respect std140
 * layout conventions. In practical terms this means you must ensure that vec3
 * and vec4 fields are 16-byte aligned.
 *
 * For better understanding of underlying concepts and memory management with
 * SDL GPU API, you may refer
 * [this blog post](https://moonside.games/posts/sdl-gpu-concepts-cycling/)
 * .
 *
 * There are optional properties that can be provided through `props`. These
 * are the supported properties:
 *
 * - `SDL_PROP_GPU_BUFFER_CREATE_NAME_STRING`: a name that can be displayed in
 *   debugging tools.
 *
 * \param device a GPU Context.
 * \param createinfo a struct describing the state of the buffer to create.
 * \returns a buffer object on success, or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_UploadToGPUBuffer
 * \sa SDL_DownloadFromGPUBuffer
 * \sa SDL_CopyGPUBufferToBuffer
 * \sa SDL_BindGPUVertexBuffers
 * \sa SDL_BindGPUIndexBuffer
 * \sa SDL_BindGPUVertexStorageBuffers
 * \sa SDL_BindGPUFragmentStorageBuffers
 * \sa SDL_DrawGPUPrimitivesIndirect
 * \sa SDL_DrawGPUIndexedPrimitivesIndirect
 * \sa SDL_BindGPUComputeStorageBuffers
 * \sa SDL_DispatchGPUComputeIndirect
 * \sa SDL_ReleaseGPUBuffer
 */
//go:sdl3extern
var CreateGPUBuffer func(
	device GPUDevice,
	createinfo *GPUBufferCreateInfo) GPUBuffer

const PROP_GPU_BUFFER_CREATE_NAME_STRING PropertyName = "SDL.gpu.buffer.create.name"

/**
 * Creates a transfer buffer to be used when uploading to or downloading from
 * graphics resources.
 *
 * Download buffers can be particularly expensive to create, so it is good
 * practice to reuse them if data will be downloaded regularly.
 *
 * There are optional properties that can be provided through `props`. These
 * are the supported properties:
 *
 * - `SDL_PROP_GPU_TRANSFERBUFFER_CREATE_NAME_STRING`: a name that can be
 *   displayed in debugging tools.
 *
 * \param device a GPU Context.
 * \param createinfo a struct describing the state of the transfer buffer to
 *                   create.
 * \returns a transfer buffer on success, or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_UploadToGPUBuffer
 * \sa SDL_DownloadFromGPUBuffer
 * \sa SDL_UploadToGPUTexture
 * \sa SDL_DownloadFromGPUTexture
 * \sa SDL_ReleaseGPUTransferBuffer
 */
//go:sdl3extern
var CreateGPUTransferBuffer func(
	device GPUDevice,
	createinfo *GPUTransferBufferCreateInfo) GPUTransferBuffer

const PROP_GPU_TRANSFERBUFFER_CREATE_NAME_STRING PropertyName = "SDL.gpu.transferbuffer.create.name"

/* Debug Naming */

/**
 * Sets an arbitrary string constant to label a buffer.
 *
 * You should use SDL_PROP_GPU_BUFFER_CREATE_NAME_STRING with
 * SDL_CreateGPUBuffer instead of this function to avoid thread safety issues.
 *
 * \param device a GPU Context.
 * \param buffer a buffer to attach the name to.
 * \param text a UTF-8 string constant to mark as the name of the buffer.
 *
 * \threadsafety This function is not thread safe, you must make sure the
 *               buffer is not simultaneously used by any other thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUBuffer
 */
//go:sdl3extern
var SetGPUBufferName func(
	device GPUDevice,
	buffer GPUBuffer,
	text string)

/**
 * Sets an arbitrary string constant to label a texture.
 *
 * You should use SDL_PROP_GPU_TEXTURE_CREATE_NAME_STRING with
 * SDL_CreateGPUTexture instead of this function to avoid thread safety
 * issues.
 *
 * \param device a GPU Context.
 * \param texture a texture to attach the name to.
 * \param text a UTF-8 string constant to mark as the name of the texture.
 *
 * \threadsafety This function is not thread safe, you must make sure the
 *               texture is not simultaneously used by any other thread.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUTexture
 */
//go:sdl3extern
var SetGPUTextureName func(
	device GPUDevice,
	texture GPUTexture,
	text string)

/**
 * Inserts an arbitrary string label into the command buffer callstream.
 *
 * Useful for debugging.
 *
 * \param command_buffer a command buffer.
 * \param text a UTF-8 string constant to insert as the label.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var InsertGPUDebugLabel func(
	command_buffer GPUCommandBuffer,
	text string)

/**
 * Begins a debug group with an arbitary name.
 *
 * Used for denoting groups of calls when viewing the command buffer
 * callstream in a graphics debugging tool.
 *
 * Each call to SDL_PushGPUDebugGroup must have a corresponding call to
 * SDL_PopGPUDebugGroup.
 *
 * On some backends (e.g. Metal), pushing a debug group during a
 * render/blit/compute pass will create a group that is scoped to the native
 * pass rather than the command buffer. For best results, if you push a debug
 * group during a pass, always pop it in the same pass.
 *
 * \param command_buffer a command buffer.
 * \param name a UTF-8 string constant that names the group.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_PopGPUDebugGroup
 */
//go:sdl3extern
var PushGPUDebugGroup func(
	command_buffer GPUCommandBuffer,
	name string)

/**
 * Ends the most-recently pushed debug group.
 *
 * \param command_buffer a command buffer.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_PushGPUDebugGroup
 */
//go:sdl3extern
var PopGPUDebugGroup func(
	command_buffer GPUCommandBuffer)

/* Disposal */

/**
 * Frees the given texture as soon as it is safe to do so.
 *
 * You must not reference the texture after calling this function.
 *
 * \param device a GPU context.
 * \param texture a texture to be destroyed.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var ReleaseGPUTexture func(
	device GPUDevice,
	texture GPUTexture)

/**
 * Frees the given sampler as soon as it is safe to do so.
 *
 * You must not reference the sampler after calling this function.
 *
 * \param device a GPU context.
 * \param sampler a sampler to be destroyed.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var ReleaseGPUSampler func(
	device GPUDevice,
	sampler GPUSampler)

/**
 * Frees the given buffer as soon as it is safe to do so.
 *
 * You must not reference the buffer after calling this function.
 *
 * \param device a GPU context.
 * \param buffer a buffer to be destroyed.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var ReleaseGPUBuffer func(
	device GPUDevice,
	buffer GPUBuffer)

/**
 * Frees the given transfer buffer as soon as it is safe to do so.
 *
 * You must not reference the transfer buffer after calling this function.
 *
 * \param device a GPU context.
 * \param transfer_buffer a transfer buffer to be destroyed.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var ReleaseGPUTransferBuffer func(
	device GPUDevice,
	transfer_buffer GPUTransferBuffer)

/**
 * Frees the given compute pipeline as soon as it is safe to do so.
 *
 * You must not reference the compute pipeline after calling this function.
 *
 * \param device a GPU context.
 * \param compute_pipeline a compute pipeline to be destroyed.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var ReleaseGPUComputePipeline func(
	device GPUDevice,
	compute_pipeline GPUComputePipeline)

/**
 * Frees the given shader as soon as it is safe to do so.
 *
 * You must not reference the shader after calling this function.
 *
 * \param device a GPU context.
 * \param shader a shader to be destroyed.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var ReleaseGPUShader func(
	device GPUDevice,
	shader GPUShader)

/**
 * Frees the given graphics pipeline as soon as it is safe to do so.
 *
 * You must not reference the graphics pipeline after calling this function.
 *
 * \param device a GPU context.
 * \param graphics_pipeline a graphics pipeline to be destroyed.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var ReleaseGPUGraphicsPipeline func(
	device GPUDevice,
	graphics_pipeline GPUGraphicsPipeline)

/**
 * Acquire a command buffer.
 *
 * This command buffer is managed by the implementation and should not be
 * freed by the user. The command buffer may only be used on the thread it was
 * acquired on. The command buffer should be submitted on the thread it was
 * acquired on.
 *
 * It is valid to acquire multiple command buffers on the same thread at once.
 * In fact a common design pattern is to acquire two command buffers per frame
 * where one is dedicated to render and compute passes and the other is
 * dedicated to copy passes and other preparatory work such as generating
 * mipmaps. Interleaving commands between the two command buffers reduces the
 * total amount of passes overall which improves rendering performance.
 *
 * \param device a GPU context.
 * \returns a command buffer, or NULL on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SubmitGPUCommandBuffer
 * \sa SDL_SubmitGPUCommandBufferAndAcquireFence
 */
//go:sdl3extern
var AcquireGPUCommandBuffer func(
	device GPUDevice) GPUCommandBuffer

/* Uniform Data */

/**
 * Pushes data to a vertex uniform slot on the command buffer.
 *
 * Subsequent draw calls will use this uniform data.
 *
 * The data being pushed must respect std140 layout conventions. In practical
 * terms this means you must ensure that vec3 and vec4 fields are 16-byte
 * aligned.
 *
 * For detailed information about accessing uniform data from a shader, please
 * refer to SDL_CreateGPUShader.
 *
 * \param command_buffer a command buffer.
 * \param slot_index the vertex uniform slot to push data to.
 * \param data client data to write.
 * \param length the length of the data to write.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var PushGPUVertexUniformData func(
	command_buffer GPUCommandBuffer,
	slot_index uint32,
	data uintptr,
	length uint32)

/**
 * Pushes data to a fragment uniform slot on the command buffer.
 *
 * Subsequent draw calls will use this uniform data.
 *
 * The data being pushed must respect std140 layout conventions. In practical
 * terms this means you must ensure that vec3 and vec4 fields are 16-byte
 * aligned.
 *
 * \param command_buffer a command buffer.
 * \param slot_index the fragment uniform slot to push data to.
 * \param data client data to write.
 * \param length the length of the data to write.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var PushGPUFragmentUniformData func(
	command_buffer GPUCommandBuffer,
	slot_index uint32,
	data uintptr,
	length uint32)

/**
 * Pushes data to a uniform slot on the command buffer.
 *
 * Subsequent draw calls will use this uniform data.
 *
 * The data being pushed must respect std140 layout conventions. In practical
 * terms this means you must ensure that vec3 and vec4 fields are 16-byte
 * aligned.
 *
 * \param command_buffer a command buffer.
 * \param slot_index the uniform slot to push data to.
 * \param data client data to write.
 * \param length the length of the data to write.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var PushGPUComputeUniformData func(
	command_buffer GPUCommandBuffer,
	slot_index uint32,
	data uintptr,
	length uint32)

/* Graphics State */

/**
 * Begins a render pass on a command buffer.
 *
 * A render pass consists of a set of texture subresources (or depth slices in
 * the 3D texture case) which will be rendered to during the render pass,
 * along with corresponding clear values and load/store operations. All
 * operations related to graphics pipelines must take place inside of a render
 * pass. A default viewport and scissor state are automatically set when this
 * is called. You cannot begin another render pass, or begin a compute pass or
 * copy pass until you have ended the render pass.
 *
 * \param command_buffer a command buffer.
 * \param color_target_infos an array of texture subresources with
 *                           corresponding clear values and load/store ops.
 * \param num_color_targets the number of color targets in the
 *                          color_target_infos array.
 * \param depth_stencil_target_info a texture subresource with corresponding
 *                                  clear value and load/store ops, may be
 *                                  NULL.
 * \returns a render pass handle.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_EndGPURenderPass
 */
//go:sdl3extern
var BeginGPURenderPass func(
	command_buffer GPUCommandBuffer,
	color_target_infos *GPUColorTargetInfo,
	num_color_targets uint32,
	depth_stencil_target_info *GPUDepthStencilTargetInfo) GPURenderPass

/**
 * Binds a graphics pipeline on a render pass to be used in rendering.
 *
 * A graphics pipeline must be bound before making any draw calls.
 *
 * \param render_pass a render pass handle.
 * \param graphics_pipeline the graphics pipeline to bind.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var BindGPUGraphicsPipeline func(
	render_pass GPURenderPass,
	graphics_pipeline GPUGraphicsPipeline)

/**
 * Sets the current viewport state on a command buffer.
 *
 * \param render_pass a render pass handle.
 * \param viewport the viewport to set.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetGPUViewport func(
	render_pass GPURenderPass,
	viewport *GPUViewport)

/**
 * Sets the current scissor state on a command buffer.
 *
 * \param render_pass a render pass handle.
 * \param scissor the scissor area to set.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetGPUScissor func(
	render_pass GPURenderPass,
	scissor *Rect)

/**
 * Sets the current blend constants on a command buffer.
 *
 * \param render_pass a render pass handle.
 * \param blend_constants the blend constant color.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_GPU_BLENDFACTOR_CONSTANT_COLOR
 * \sa SDL_GPU_BLENDFACTOR_ONE_MINUS_CONSTANT_COLOR
 */
//go:sdl3extern
var SetGPUBlendConstants func(
	render_pass GPURenderPass,
	blend_constants *FColor)

/**
 * Sets the current stencil reference value on a command buffer.
 *
 * \param render_pass a render pass handle.
 * \param reference the stencil reference value to set.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetGPUStencilReference func(
	render_pass GPURenderPass,
	reference uint8)

/**
 * Binds vertex buffers on a command buffer for use with subsequent draw
 * calls.
 *
 * \param render_pass a render pass handle.
 * \param first_slot the vertex buffer slot to begin binding from.
 * \param bindings an array of SDL_GPUBufferBinding structs containing vertex
 *                 buffers and offset values.
 * \param num_bindings the number of bindings in the bindings array.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var BindGPUVertexBuffers func(
	render_pass GPURenderPass,
	first_slot uint32,
	bindings *GPUBufferBinding,
	num_bindings uint32)

/**
 * Binds an index buffer on a command buffer for use with subsequent draw
 * calls.
 *
 * \param render_pass a render pass handle.
 * \param binding a pointer to a struct containing an index buffer and offset.
 * \param index_element_size whether the index values in the buffer are 16- or
 *                           32-bit.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var BindGPUIndexBuffer func(
	render_pass GPURenderPass,
	binding *GPUBufferBinding,
	index_element_size GPUIndexElementSize)

/**
 * Binds texture-sampler pairs for use on the vertex shader.
 *
 * The textures must have been created with SDL_GPU_TEXTUREUSAGE_SAMPLER.
 *
 * Be sure your shader is set up according to the requirements documented in
 * SDL_CreateGPUShader().
 *
 * \param render_pass a render pass handle.
 * \param first_slot the vertex sampler slot to begin binding from.
 * \param texture_sampler_bindings an array of texture-sampler binding
 *                                 structs.
 * \param num_bindings the number of texture-sampler pairs to bind from the
 *                     array.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
//go:sdl3extern
var BindGPUVertexSamplers func(
	render_pass GPURenderPass,
	first_slot uint32,
	texture_sampler_bindings *GPUTextureSamplerBinding,
	num_bindings uint32)

/**
 * Binds storage textures for use on the vertex shader.
 *
 * These textures must have been created with
 * SDL_GPU_TEXTUREUSAGE_GRAPHICS_STORAGE_READ.
 *
 * Be sure your shader is set up according to the requirements documented in
 * SDL_CreateGPUShader().
 *
 * \param render_pass a render pass handle.
 * \param first_slot the vertex storage texture slot to begin binding from.
 * \param storage_textures an array of storage textures.
 * \param num_bindings the number of storage texture to bind from the array.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
//go:sdl3extern
var BindGPUVertexStorageTextures func(
	render_pass GPURenderPass,
	first_slot uint32,
	storage_textures []GPUTexture,
	num_bindings uint32)

/**
 * Binds storage buffers for use on the vertex shader.
 *
 * These buffers must have been created with
 * SDL_GPU_BUFFERUSAGE_GRAPHICS_STORAGE_READ.
 *
 * Be sure your shader is set up according to the requirements documented in
 * SDL_CreateGPUShader().
 *
 * \param render_pass a render pass handle.
 * \param first_slot the vertex storage buffer slot to begin binding from.
 * \param storage_buffers an array of buffers.
 * \param num_bindings the number of buffers to bind from the array.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
//go:sdl3extern
var BindGPUVertexStorageBuffers func(
	render_pass GPURenderPass,
	first_slot uint32,
	storage_buffers []GPUBuffer,
	num_bindings uint32)

/**
 * Binds texture-sampler pairs for use on the fragment shader.
 *
 * The textures must have been created with SDL_GPU_TEXTUREUSAGE_SAMPLER.
 *
 * Be sure your shader is set up according to the requirements documented in
 * SDL_CreateGPUShader().
 *
 * \param render_pass a render pass handle.
 * \param first_slot the fragment sampler slot to begin binding from.
 * \param texture_sampler_bindings an array of texture-sampler binding
 *                                 structs.
 * \param num_bindings the number of texture-sampler pairs to bind from the
 *                     array.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
//go:sdl3extern
var BindGPUFragmentSamplers func(
	render_pass GPURenderPass,
	first_slot uint32,
	texture_sampler_bindings []GPUTextureSamplerBinding,
	num_bindings uint32)

/**
 * Binds storage textures for use on the fragment shader.
 *
 * These textures must have been created with
 * SDL_GPU_TEXTUREUSAGE_GRAPHICS_STORAGE_READ.
 *
 * Be sure your shader is set up according to the requirements documented in
 * SDL_CreateGPUShader().
 *
 * \param render_pass a render pass handle.
 * \param first_slot the fragment storage texture slot to begin binding from.
 * \param storage_textures an array of storage textures.
 * \param num_bindings the number of storage textures to bind from the array.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
//go:sdl3extern
var BindGPUFragmentStorageTextures func(
	render_pass GPURenderPass,
	first_slot uint32,
	storage_textures []GPUTexture,
	num_bindings uint32)

/**
 * Binds storage buffers for use on the fragment shader.
 *
 * These buffers must have been created with
 * SDL_GPU_BUFFERUSAGE_GRAPHICS_STORAGE_READ.
 *
 * Be sure your shader is set up according to the requirements documented in
 * SDL_CreateGPUShader().
 *
 * \param render_pass a render pass handle.
 * \param first_slot the fragment storage buffer slot to begin binding from.
 * \param storage_buffers an array of storage buffers.
 * \param num_bindings the number of storage buffers to bind from the array.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
//go:sdl3extern
var BindGPUFragmentStorageBuffers func(
	render_pass GPURenderPass,
	first_slot uint32,
	storage_buffers []GPUBuffer,
	num_bindings uint32)

/* Drawing */

/**
 * Draws data using bound graphics state with an index buffer and instancing
 * enabled.
 *
 * You must not call this function before binding a graphics pipeline.
 *
 * Note that the `first_vertex` and `first_instance` parameters are NOT
 * compatible with built-in vertex/instance ID variables in shaders (for
 * example, SV_VertexID); GPU APIs and shader languages do not define these
 * built-in variables consistently, so if your shader depends on them, the
 * only way to keep behavior consistent and portable is to always pass 0 for
 * the correlating parameter in the draw calls.
 *
 * \param render_pass a render pass handle.
 * \param num_indices the number of indices to draw per instance.
 * \param num_instances the number of instances to draw.
 * \param first_index the starting index within the index buffer.
 * \param vertex_offset value added to vertex index before indexing into the
 *                      vertex buffer.
 * \param first_instance the ID of the first instance to draw.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var DrawGPUIndexedPrimitives func(
	render_pass GPURenderPass,
	num_indices uint32,
	num_instances uint32,
	first_index uint32,
	vertex_offset int32,
	first_instance uint32)

/**
 * Draws data using bound graphics state.
 *
 * You must not call this function before binding a graphics pipeline.
 *
 * Note that the `first_vertex` and `first_instance` parameters are NOT
 * compatible with built-in vertex/instance ID variables in shaders (for
 * example, SV_VertexID); GPU APIs and shader languages do not define these
 * built-in variables consistently, so if your shader depends on them, the
 * only way to keep behavior consistent and portable is to always pass 0 for
 * the correlating parameter in the draw calls.
 *
 * \param render_pass a render pass handle.
 * \param num_vertices the number of vertices to draw.
 * \param num_instances the number of instances that will be drawn.
 * \param first_vertex the index of the first vertex to draw.
 * \param first_instance the ID of the first instance to draw.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var DrawGPUPrimitives func(
	render_pass GPURenderPass,
	num_vertices uint32,
	num_instances uint32,
	first_vertex uint32,
	first_instance uint32)

/**
 * Draws data using bound graphics state and with draw parameters set from a
 * buffer.
 *
 * The buffer must consist of tightly-packed draw parameter sets that each
 * match the layout of SDL_GPUIndirectDrawCommand. You must not call this
 * function before binding a graphics pipeline.
 *
 * \param render_pass a render pass handle.
 * \param buffer a buffer containing draw parameters.
 * \param offset the offset to start reading from the draw buffer.
 * \param draw_count the number of draw parameter sets that should be read
 *                   from the draw buffer.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var DrawGPUPrimitivesIndirect func(
	render_pass GPURenderPass,
	buffer GPUBuffer,
	offset uint32,
	draw_count uint32)

/**
 * Draws data using bound graphics state with an index buffer enabled and with
 * draw parameters set from a buffer.
 *
 * The buffer must consist of tightly-packed draw parameter sets that each
 * match the layout of SDL_GPUIndexedIndirectDrawCommand. You must not call
 * this function before binding a graphics pipeline.
 *
 * \param render_pass a render pass handle.
 * \param buffer a buffer containing draw parameters.
 * \param offset the offset to start reading from the draw buffer.
 * \param draw_count the number of draw parameter sets that should be read
 *                   from the draw buffer.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var DrawGPUIndexedPrimitivesIndirect func(
	render_pass GPURenderPass,
	buffer GPUBuffer,
	offset uint32,
	draw_count uint32)

/**
 * Ends the given render pass.
 *
 * All bound graphics state on the render pass command buffer is unset. The
 * render pass handle is now invalid.
 *
 * \param render_pass a render pass handle.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var EndGPURenderPass func(
	render_pass GPURenderPass)

/* Compute Pass */

/**
 * Begins a compute pass on a command buffer.
 *
 * A compute pass is defined by a set of texture subresources and buffers that
 * may be written to by compute pipelines. These textures and buffers must
 * have been created with the COMPUTE_STORAGE_WRITE bit or the
 * COMPUTE_STORAGE_SIMULTANEOUS_READ_WRITE bit. If you do not create a texture
 * with COMPUTE_STORAGE_SIMULTANEOUS_READ_WRITE, you must not read from the
 * texture in the compute pass. All operations related to compute pipelines
 * must take place inside of a compute pass. You must not begin another
 * compute pass, or a render pass or copy pass before ending the compute pass.
 *
 * A VERY IMPORTANT NOTE - Reads and writes in compute passes are NOT
 * implicitly synchronized. This means you may cause data races by both
 * reading and writing a resource region in a compute pass, or by writing
 * multiple times to a resource region. If your compute work depends on
 * reading the completed output from a previous dispatch, you MUST end the
 * current compute pass and begin a new one before you can safely access the
 * data. Otherwise you will receive unexpected results. Reading and writing a
 * texture in the same compute pass is only supported by specific texture
 * formats. Make sure you check the format support!
 *
 * \param command_buffer a command buffer.
 * \param storage_texture_bindings an array of writeable storage texture
 *                                 binding structs.
 * \param num_storage_texture_bindings the number of storage textures to bind
 *                                     from the array.
 * \param storage_buffer_bindings an array of writeable storage buffer binding
 *                                structs.
 * \param num_storage_buffer_bindings the number of storage buffers to bind
 *                                    from the array.
 * \returns a compute pass handle.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_EndGPUComputePass
 */
//go:sdl3extern
var BeginGPUComputePass func(
	command_buffer GPUCommandBuffer,
	storage_texture_bindings []GPUStorageTextureReadWriteBinding,
	num_storage_texture_bindings uint32,
	storage_buffer_bindings []GPUStorageBufferReadWriteBinding,
	num_storage_buffer_bindings uint32) GPUComputePass

/**
 * Binds a compute pipeline on a command buffer for use in compute dispatch.
 *
 * \param compute_pass a compute pass handle.
 * \param compute_pipeline a compute pipeline to bind.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var BindGPUComputePipeline func(
	compute_pass GPUComputePass,
	compute_pipeline GPUComputePipeline)

/**
 * Binds texture-sampler pairs for use on the compute shader.
 *
 * The textures must have been created with SDL_GPU_TEXTUREUSAGE_SAMPLER.
 *
 * Be sure your shader is set up according to the requirements documented in
 * SDL_CreateGPUShader().
 *
 * \param compute_pass a compute pass handle.
 * \param first_slot the compute sampler slot to begin binding from.
 * \param texture_sampler_bindings an array of texture-sampler binding
 *                                 structs.
 * \param num_bindings the number of texture-sampler bindings to bind from the
 *                     array.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
//go:sdl3extern
var BindGPUComputeSamplers func(
	compute_pass GPUComputePass,
	first_slot uint32,
	texture_sampler_bindings []GPUTextureSamplerBinding,
	num_bindings uint32)

/**
 * Binds storage textures as readonly for use on the compute pipeline.
 *
 * These textures must have been created with
 * SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_READ.
 *
 * Be sure your shader is set up according to the requirements documented in
 * SDL_CreateGPUShader().
 *
 * \param compute_pass a compute pass handle.
 * \param first_slot the compute storage texture slot to begin binding from.
 * \param storage_textures an array of storage textures.
 * \param num_bindings the number of storage textures to bind from the array.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
//go:sdl3extern
var BindGPUComputeStorageTextures func(
	compute_pass GPUComputePass,
	first_slot uint32,
	storage_textures []GPUTexture,
	num_bindings uint32)

/**
 * Binds storage buffers as readonly for use on the compute pipeline.
 *
 * These buffers must have been created with
 * SDL_GPU_BUFFERUSAGE_COMPUTE_STORAGE_READ.
 *
 * Be sure your shader is set up according to the requirements documented in
 * SDL_CreateGPUShader().
 *
 * \param compute_pass a compute pass handle.
 * \param first_slot the compute storage buffer slot to begin binding from.
 * \param storage_buffers an array of storage buffer binding structs.
 * \param num_bindings the number of storage buffers to bind from the array.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_CreateGPUShader
 */
//go:sdl3extern
var BindGPUComputeStorageBuffers func(
	compute_pass GPUComputePass,
	first_slot uint32,
	storage_buffers []GPUBuffer,
	num_bindings uint32)

/**
 * Dispatches compute work.
 *
 * You must not call this function before binding a compute pipeline.
 *
 * A VERY IMPORTANT NOTE If you dispatch multiple times in a compute pass, and
 * the dispatches write to the same resource region as each other, there is no
 * guarantee of which order the writes will occur. If the write order matters,
 * you MUST end the compute pass and begin another one.
 *
 * \param compute_pass a compute pass handle.
 * \param groupcount_x number of local workgroups to dispatch in the X
 *                     dimension.
 * \param groupcount_y number of local workgroups to dispatch in the Y
 *                     dimension.
 * \param groupcount_z number of local workgroups to dispatch in the Z
 *                     dimension.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var DispatchGPUCompute func(
	compute_pass GPUComputePass,
	groupcount_x uint32,
	groupcount_y uint32,
	groupcount_z uint32)

/**
 * Dispatches compute work with parameters set from a buffer.
 *
 * The buffer layout should match the layout of
 * SDL_GPUIndirectDispatchCommand. You must not call this function before
 * binding a compute pipeline.
 *
 * A VERY IMPORTANT NOTE If you dispatch multiple times in a compute pass, and
 * the dispatches write to the same resource region as each other, there is no
 * guarantee of which order the writes will occur. If the write order matters,
 * you MUST end the compute pass and begin another one.
 *
 * \param compute_pass a compute pass handle.
 * \param buffer a buffer containing dispatch parameters.
 * \param offset the offset to start reading from the dispatch buffer.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var DispatchGPUComputeIndirect func(
	compute_pass GPUComputePass,
	buffer GPUBuffer,
	offset uint32)

/**
 * Ends the current compute pass.
 *
 * All bound compute state on the command buffer is unset. The compute pass
 * handle is now invalid.
 *
 * \param compute_pass a compute pass handle.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var EndGPUComputePass func(
	compute_pass GPUComputePass)

/* TransferBuffer Data */

/**
 * Maps a transfer buffer into application address space.
 *
 * You must unmap the transfer buffer before encoding upload commands. The
 * memory is owned by the graphics driver - do NOT call SDL_free() on the
 * returned pointer.
 *
 * \param device a GPU context.
 * \param transfer_buffer a transfer buffer.
 * \param cycle if true, cycles the transfer buffer if it is already bound.
 * \returns the address of the mapped transfer buffer memory, or NULL on
 *          failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var MapGPUTransferBuffer func(
	device GPUDevice,
	transfer_buffer GPUTransferBuffer,
	cycle bool) uintptr

/**
 * Unmaps a previously mapped transfer buffer.
 *
 * \param device a GPU context.
 * \param transfer_buffer a previously mapped transfer buffer.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var UnmapGPUTransferBuffer func(
	device GPUDevice,
	transfer_buffer GPUTransferBuffer)

/* Copy Pass */

/**
 * Begins a copy pass on a command buffer.
 *
 * All operations related to copying to or from buffers or textures take place
 * inside a copy pass. You must not begin another copy pass, or a render pass
 * or compute pass before ending the copy pass.
 *
 * \param command_buffer a command buffer.
 * \returns a copy pass handle.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var BeginGPUCopyPass func(
	command_buffer GPUCommandBuffer) GPUCopyPass

/**
 * Uploads data from a transfer buffer to a texture.
 *
 * The upload occurs on the GPU timeline. You may assume that the upload has
 * finished in subsequent commands.
 *
 * You must align the data in the transfer buffer to a multiple of the texel
 * size of the texture format.
 *
 * \param copy_pass a copy pass handle.
 * \param source the source transfer buffer with image layout information.
 * \param destination the destination texture region.
 * \param cycle if true, cycles the texture if the texture is bound, otherwise
 *              overwrites the data.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var UploadToGPUTexture func(
	copy_pass GPUCopyPass,
	source *GPUTextureTransferInfo,
	destination *GPUTextureRegion,
	cycle bool)

/**
 * Uploads data from a transfer buffer to a buffer.
 *
 * The upload occurs on the GPU timeline. You may assume that the upload has
 * finished in subsequent commands.
 *
 * \param copy_pass a copy pass handle.
 * \param source the source transfer buffer with offset.
 * \param destination the destination buffer with offset and size.
 * \param cycle if true, cycles the buffer if it is already bound, otherwise
 *              overwrites the data.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var UploadToGPUBuffer func(
	copy_pass GPUCopyPass,
	source *GPUTransferBufferLocation,
	destination *GPUBufferRegion,
	cycle bool)

/**
 * Performs a texture-to-texture copy.
 *
 * This copy occurs on the GPU timeline. You may assume the copy has finished
 * in subsequent commands.
 *
 * \param copy_pass a copy pass handle.
 * \param source a source texture region.
 * \param destination a destination texture region.
 * \param w the width of the region to copy.
 * \param h the height of the region to copy.
 * \param d the depth of the region to copy.
 * \param cycle if true, cycles the destination texture if the destination
 *              texture is bound, otherwise overwrites the data.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var CopyGPUTextureToTexture func(
	copy_pass GPUCopyPass,
	source *GPUTextureLocation,
	destination *GPUTextureLocation,
	w uint32,
	h uint32,
	d uint32,
	cycle bool)

/**
 * Performs a buffer-to-buffer copy.
 *
 * This copy occurs on the GPU timeline. You may assume the copy has finished
 * in subsequent commands.
 *
 * \param copy_pass a copy pass handle.
 * \param source the buffer and offset to copy from.
 * \param destination the buffer and offset to copy to.
 * \param size the length of the buffer to copy.
 * \param cycle if true, cycles the destination buffer if it is already bound,
 *              otherwise overwrites the data.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var CopyGPUBufferToBuffer func(
	copy_pass GPUCopyPass,
	source *GPUBufferLocation,
	destination *GPUBufferLocation,
	size uint32,
	cycle bool)

/**
 * Copies data from a texture to a transfer buffer on the GPU timeline.
 *
 * This data is not guaranteed to be copied until the command buffer fence is
 * signaled.
 *
 * \param copy_pass a copy pass handle.
 * \param source the source texture region.
 * \param destination the destination transfer buffer with image layout
 *                    information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var DownloadFromGPUTexture func(
	copy_pass GPUCopyPass,
	source *GPUTextureRegion,
	destination *GPUTextureTransferInfo)

/**
 * Copies data from a buffer to a transfer buffer on the GPU timeline.
 *
 * This data is not guaranteed to be copied until the command buffer fence is
 * signaled.
 *
 * \param copy_pass a copy pass handle.
 * \param source the source buffer with offset and size.
 * \param destination the destination transfer buffer with offset.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var DownloadFromGPUBuffer func(
	copy_pass GPUCopyPass,
	source *GPUBufferRegion,
	destination *GPUTransferBufferLocation)

/**
 * Ends the current copy pass.
 *
 * \param copy_pass a copy pass handle.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var EndGPUCopyPass func(
	copy_pass GPUCopyPass)

/**
 * Generates mipmaps for the given texture.
 *
 * This function must not be called inside of any pass.
 *
 * \param command_buffer a command_buffer.
 * \param texture a texture with more than 1 mip level.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GenerateMipmapsForGPUTexture func(
	command_buffer GPUCommandBuffer,
	texture GPUTexture)

/**
 * Blits from a source texture region to a destination texture region.
 *
 * This function must not be called inside of any pass.
 *
 * \param command_buffer a command buffer.
 * \param info the blit info struct containing the blit parameters.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var BlitGPUTexture func(
	command_buffer GPUCommandBuffer,
	info *GPUBlitInfo)

/* Submission/Presentation */

/**
 * Determines whether a swapchain composition is supported by the window.
 *
 * The window must be claimed before calling this function.
 *
 * \param device a GPU context.
 * \param window an SDL_Window.
 * \param swapchain_composition the swapchain composition to check.
 * \returns true if supported, false if unsupported.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_ClaimWindowForGPUDevice
 */
//go:sdl3extern
var WindowSupportsGPUSwapchainComposition func(
	device GPUDevice,
	window Window,
	swapchain_composition GPUSwapchainComposition) bool

/**
 * Determines whether a presentation mode is supported by the window.
 *
 * The window must be claimed before calling this function.
 *
 * \param device a GPU context.
 * \param window an SDL_Window.
 * \param present_mode the presentation mode to check.
 * \returns true if supported, false if unsupported.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_ClaimWindowForGPUDevice
 */
//go:sdl3extern
var WindowSupportsGPUPresentMode func(
	device GPUDevice,
	window Window,
	present_mode GPUPresentMode) bool

/**
 * Claims a window, creating a swapchain structure for it.
 *
 * This must be called before SDL_AcquireGPUSwapchainTexture is called using
 * the window. You should only call this function from the thread that created
 * the window.
 *
 * The swapchain will be created with SDL_GPU_SWAPCHAINCOMPOSITION_SDR and
 * SDL_GPU_PRESENTMODE_VSYNC. If you want to have different swapchain
 * parameters, you must call SDL_SetGPUSwapchainParameters after claiming the
 * window.
 *
 * \param device a GPU context.
 * \param window an SDL_Window.
 * \returns true on success, or false on failure; call SDL_GetError() for more
 *          information.
 *
 * \threadsafety This function should only be called from the thread that
 *               created the window.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_WaitAndAcquireGPUSwapchainTexture
 * \sa SDL_ReleaseWindowFromGPUDevice
 * \sa SDL_WindowSupportsGPUPresentMode
 * \sa SDL_WindowSupportsGPUSwapchainComposition
 */
//go:sdl3extern
var ClaimWindowForGPUDevice func(
	device GPUDevice,
	window Window) bool

/**
 * Unclaims a window, destroying its swapchain structure.
 *
 * \param device a GPU context.
 * \param window an SDL_Window that has been claimed.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_ClaimWindowForGPUDevice
 */
//go:sdl3extern
var ReleaseWindowFromGPUDevice func(
	device GPUDevice,
	window Window)

/**
 * Changes the swapchain parameters for the given claimed window.
 *
 * This function will fail if the requested present mode or swapchain
 * composition are unsupported by the device. Check if the parameters are
 * supported via SDL_WindowSupportsGPUPresentMode /
 * SDL_WindowSupportsGPUSwapchainComposition prior to calling this function.
 *
 * SDL_GPU_PRESENTMODE_VSYNC with SDL_GPU_SWAPCHAINCOMPOSITION_SDR is always
 * supported.
 *
 * \param device a GPU context.
 * \param window an SDL_Window that has been claimed.
 * \param swapchain_composition the desired composition of the swapchain.
 * \param present_mode the desired present mode for the swapchain.
 * \returns true if successful, false on error; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_WindowSupportsGPUPresentMode
 * \sa SDL_WindowSupportsGPUSwapchainComposition
 */
//go:sdl3extern
var SetGPUSwapchainParameters func(
	device GPUDevice,
	window Window,
	swapchain_composition GPUSwapchainComposition,
	present_mode GPUPresentMode) bool

/**
 * Configures the maximum allowed number of frames in flight.
 *
 * The default value when the device is created is 2. This means that after
 * you have submitted 2 frames for presentation, if the GPU has not finished
 * working on the first frame, SDL_AcquireGPUSwapchainTexture() will fill the
 * swapchain texture pointer with NULL, and
 * SDL_WaitAndAcquireGPUSwapchainTexture() will block.
 *
 * Higher values increase throughput at the expense of visual latency. Lower
 * values decrease visual latency at the expense of throughput.
 *
 * Note that calling this function will stall and flush the command queue to
 * prevent synchronization issues.
 *
 * The minimum value of allowed frames in flight is 1, and the maximum is 3.
 *
 * \param device a GPU context.
 * \param allowed_frames_in_flight the maximum number of frames that can be
 *                                 pending on the GPU.
 * \returns true if successful, false on error; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SetGPUAllowedFramesInFlight func(
	device GPUDevice,
	allowed_frames_in_flight uint32) bool

/**
 * Obtains the texture format of the swapchain for the given window.
 *
 * Note that this format can change if the swapchain parameters change.
 *
 * \param device a GPU context.
 * \param window an SDL_Window that has been claimed.
 * \returns the texture format of the swapchain.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetGPUSwapchainTextureFormat func(
	device GPUDevice,
	window Window) GPUTextureFormat

/**
 * Acquire a texture to use in presentation.
 *
 * When a swapchain texture is acquired on a command buffer, it will
 * automatically be submitted for presentation when the command buffer is
 * submitted. The swapchain texture should only be referenced by the command
 * buffer used to acquire it.
 *
 * This function will fill the swapchain texture handle with NULL if too many
 * frames are in flight. This is not an error. This NULL pointer should not be
 * passed back into SDL. Instead, it should be considered as an indication to
 * wait until the swapchain is available.
 *
 * If you use this function, it is possible to create a situation where many
 * command buffers are allocated while the rendering context waits for the GPU
 * to catch up, which will cause memory usage to grow. You should use
 * SDL_WaitAndAcquireGPUSwapchainTexture() unless you know what you are doing
 * with timing.
 *
 * The swapchain texture is managed by the implementation and must not be
 * freed by the user. You MUST NOT call this function from any thread other
 * than the one that created the window.
 *
 * \param command_buffer a command buffer.
 * \param window a window that has been claimed.
 * \param swapchain_texture a pointer filled in with a swapchain texture
 *                          handle.
 * \param swapchain_texture_width a pointer filled in with the swapchain
 *                                texture width, may be NULL.
 * \param swapchain_texture_height a pointer filled in with the swapchain
 *                                 texture height, may be NULL.
 * \returns true on success, false on error; call SDL_GetError() for more
 *          information.
 *
 * \threadsafety This function should only be called from the thread that
 *               created the window.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_ClaimWindowForGPUDevice
 * \sa SDL_SubmitGPUCommandBuffer
 * \sa SDL_SubmitGPUCommandBufferAndAcquireFence
 * \sa SDL_CancelGPUCommandBuffer
 * \sa SDL_GetWindowSizeInPixels
 * \sa SDL_WaitForGPUSwapchain
 * \sa SDL_WaitAndAcquireGPUSwapchainTexture
 * \sa SDL_SetGPUAllowedFramesInFlight
 */
//go:sdl3extern
var AcquireGPUSwapchainTexture func(
	command_buffer GPUCommandBuffer,
	window Window,
	swapchain_texture GPUTexture,
	swapchain_texture_width *uint32,
	swapchain_texture_height *uint32) bool

/**
 * Blocks the thread until a swapchain texture is available to be acquired.
 *
 * \param device a GPU context.
 * \param window a window that has been claimed.
 * \returns true on success, false on failure; call SDL_GetError() for more
 *          information.
 *
 * \threadsafety This function should only be called from the thread that
 *               created the window.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AcquireGPUSwapchainTexture
 * \sa SDL_WaitAndAcquireGPUSwapchainTexture
 * \sa SDL_SetGPUAllowedFramesInFlight
 */
//go:sdl3extern
var WaitForGPUSwapchain func(
	device GPUDevice,
	window Window) bool

/**
 * Blocks the thread until a swapchain texture is available to be acquired,
 * and then acquires it.
 *
 * When a swapchain texture is acquired on a command buffer, it will
 * automatically be submitted for presentation when the command buffer is
 * submitted. The swapchain texture should only be referenced by the command
 * buffer used to acquire it. It is an error to call
 * SDL_CancelGPUCommandBuffer() after a swapchain texture is acquired.
 *
 * This function can fill the swapchain texture handle with NULL in certain
 * cases, for example if the window is minimized. This is not an error. You
 * should always make sure to check whether the pointer is NULL before
 * actually using it.
 *
 * The swapchain texture is managed by the implementation and must not be
 * freed by the user. You MUST NOT call this function from any thread other
 * than the one that created the window.
 *
 * The swapchain texture is write-only and cannot be used as a sampler or for
 * another reading operation.
 *
 * \param command_buffer a command buffer.
 * \param window a window that has been claimed.
 * \param swapchain_texture a pointer filled in with a swapchain texture
 *                          handle.
 * \param swapchain_texture_width a pointer filled in with the swapchain
 *                                texture width, may be NULL.
 * \param swapchain_texture_height a pointer filled in with the swapchain
 *                                 texture height, may be NULL.
 * \returns true on success, false on error; call SDL_GetError() for more
 *          information.
 *
 * \threadsafety This function should only be called from the thread that
 *               created the window.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SubmitGPUCommandBuffer
 * \sa SDL_SubmitGPUCommandBufferAndAcquireFence
 * \sa SDL_AcquireGPUSwapchainTexture
 */
//go:sdl3extern
var WaitAndAcquireGPUSwapchainTexture func(
	command_buffer GPUCommandBuffer,
	window Window,
	swapchain_texture GPUTexture,
	swapchain_texture_width *uint32,
	swapchain_texture_height *uint32) bool

/**
 * Submits a command buffer so its commands can be processed on the GPU.
 *
 * It is invalid to use the command buffer after this is called.
 *
 * This must be called from the thread the command buffer was acquired on.
 *
 * All commands in the submission are guaranteed to begin executing before any
 * command in a subsequent submission begins executing.
 *
 * \param command_buffer a command buffer.
 * \returns true on success, false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AcquireGPUCommandBuffer
 * \sa SDL_WaitAndAcquireGPUSwapchainTexture
 * \sa SDL_AcquireGPUSwapchainTexture
 * \sa SDL_SubmitGPUCommandBufferAndAcquireFence
 */
//go:sdl3extern
var SubmitGPUCommandBuffer func(
	command_buffer GPUCommandBuffer) bool

/**
 * Submits a command buffer so its commands can be processed on the GPU, and
 * acquires a fence associated with the command buffer.
 *
 * You must release this fence when it is no longer needed or it will cause a
 * leak. It is invalid to use the command buffer after this is called.
 *
 * This must be called from the thread the command buffer was acquired on.
 *
 * All commands in the submission are guaranteed to begin executing before any
 * command in a subsequent submission begins executing.
 *
 * \param command_buffer a command buffer.
 * \returns a fence associated with the command buffer, or NULL on failure;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_AcquireGPUCommandBuffer
 * \sa SDL_WaitAndAcquireGPUSwapchainTexture
 * \sa SDL_AcquireGPUSwapchainTexture
 * \sa SDL_SubmitGPUCommandBuffer
 * \sa SDL_ReleaseGPUFence
 */
//go:sdl3extern
var SubmitGPUCommandBufferAndAcquireFence func(
	command_buffer GPUCommandBuffer) GPUFence

/**
 * Cancels a command buffer.
 *
 * None of the enqueued commands are executed.
 *
 * It is an error to call this function after a swapchain texture has been
 * acquired.
 *
 * This must be called from the thread the command buffer was acquired on.
 *
 * You must not reference the command buffer after calling this function.
 *
 * \param command_buffer a command buffer.
 * \returns true on success, false on error; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_WaitAndAcquireGPUSwapchainTexture
 * \sa SDL_AcquireGPUCommandBuffer
 * \sa SDL_AcquireGPUSwapchainTexture
 */
//go:sdl3extern
var CancelGPUCommandBuffer func(
	command_buffer GPUCommandBuffer) bool

/**
 * Blocks the thread until the GPU is completely idle.
 *
 * \param device a GPU context.
 * \returns true on success, false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_WaitForGPUFences
 */
//go:sdl3extern
var WaitForGPUIdle func(
	device GPUDevice) bool

/**
 * Blocks the thread until the given fences are signaled.
 *
 * \param device a GPU context.
 * \param wait_all if 0, wait for any fence to be signaled, if 1, wait for all
 *                 fences to be signaled.
 * \param fences an array of fences to wait on.
 * \param num_fences the number of fences in the fences array.
 * \returns true on success, false on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SubmitGPUCommandBufferAndAcquireFence
 * \sa SDL_WaitForGPUIdle
 */
//go:sdl3extern
var WaitForGPUFences func(
	device GPUDevice,
	wait_all bool,
	fences []GPUFence,
	num_fences uint32) bool

/**
 * Checks the status of a fence.
 *
 * \param device a GPU context.
 * \param fence a fence.
 * \returns true if the fence is signaled, false if it is not.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SubmitGPUCommandBufferAndAcquireFence
 */
//go:sdl3extern
var QueryGPUFence func(
	device GPUDevice,
	fence GPUFence) bool

/**
 * Releases a fence obtained from SDL_SubmitGPUCommandBufferAndAcquireFence.
 *
 * You must not reference the fence after calling this function.
 *
 * \param device a GPU context.
 * \param fence a fence.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_SubmitGPUCommandBufferAndAcquireFence
 */
//go:sdl3extern
var ReleaseGPUFence func(
	device GPUDevice,
	fence GPUFence)

/* Format Info */

/**
 * Obtains the texel block size for a texture format.
 *
 * \param format the texture format you want to know the texel size of.
 * \returns the texel block size of the texture format.
 *
 * \since This function is available since SDL 3.2.0.
 *
 * \sa SDL_UploadToGPUTexture
 */
//go:sdl3extern
var GPUTextureFormatTexelBlockSize func(
	format GPUTextureFormat) uint32

/**
 * Determines whether a texture format is supported for a given type and
 * usage.
 *
 * \param device a GPU context.
 * \param format the texture format to check.
 * \param type the type of texture (2D, 3D, Cube).
 * \param usage a bitmask of all usage scenarios to check.
 * \returns whether the texture format is supported for this type and usage.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GPUTextureSupportsFormat func(
	device GPUDevice,
	format GPUTextureFormat,
	type_ GPUTextureType,
	usage GPUTextureUsageFlags) bool

/**
 * Determines if a sample count for a texture format is supported.
 *
 * \param device a GPU context.
 * \param format the texture format to check.
 * \param sample_count the sample count to check.
 * \returns whether the sample count is supported for this texture format.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GPUTextureSupportsSampleCount func(
	device GPUDevice,
	format GPUTextureFormat,
	sample_count GPUSampleCount) bool

/**
 * Calculate the size in bytes of a texture format with dimensions.
 *
 * \param format a texture format.
 * \param width width in pixels.
 * \param height height in pixels.
 * \param depth_or_layer_count depth for 3D textures or layer count otherwise.
 * \returns the size of a texture with this format and dimensions.
 *
 * \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var CalculateGPUTextureFormatSize func(
	format GPUTextureFormat,
	width uint32,
	height uint32,
	depth_or_layer_count uint32) uint32
