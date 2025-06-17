//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_GPUSupportsShaderFormats
func __SDL_GPUSupportsShaderFormats(int32, uintptr) int32

//go:wasmimport sdl3 SDL_GPUSupportsProperties
func __SDL_GPUSupportsProperties(int32) int32

//go:wasmimport sdl3 SDL_CreateGPUDevice
func __SDL_CreateGPUDevice(int32, int32, uintptr) uintptr

//go:wasmimport sdl3 SDL_CreateGPUDeviceWithProperties
func __SDL_CreateGPUDeviceWithProperties(int32) uintptr

//go:wasmimport sdl3 SDL_DestroyGPUDevice
func __SDL_DestroyGPUDevice(uintptr)

//go:wasmimport sdl3 SDL_GetNumGPUDrivers
func __SDL_GetNumGPUDrivers() int32

//go:wasmimport sdl3 SDL_GetGPUDriver
func __SDL_GetGPUDriver(int32) uintptr

//go:wasmimport sdl3 SDL_GetGPUDeviceDriver
func __SDL_GetGPUDeviceDriver(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetGPUShaderFormats
func __SDL_GetGPUShaderFormats(uintptr) int32

//go:wasmimport sdl3 SDL_GetGPUDeviceProperties
func __SDL_GetGPUDeviceProperties(uintptr) int32

//go:wasmimport sdl3 SDL_CreateGPUComputePipeline
func __SDL_CreateGPUComputePipeline(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_CreateGPUGraphicsPipeline
func __SDL_CreateGPUGraphicsPipeline(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_CreateGPUSampler
func __SDL_CreateGPUSampler(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_CreateGPUShader
func __SDL_CreateGPUShader(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_CreateGPUTexture
func __SDL_CreateGPUTexture(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_CreateGPUBuffer
func __SDL_CreateGPUBuffer(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_CreateGPUTransferBuffer
func __SDL_CreateGPUTransferBuffer(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_SetGPUBufferName
func __SDL_SetGPUBufferName(uintptr, uintptr, uintptr)

//go:wasmimport sdl3 SDL_SetGPUTextureName
func __SDL_SetGPUTextureName(uintptr, uintptr, uintptr)

//go:wasmimport sdl3 SDL_InsertGPUDebugLabel
func __SDL_InsertGPUDebugLabel(uintptr, uintptr)

//go:wasmimport sdl3 SDL_PushGPUDebugGroup
func __SDL_PushGPUDebugGroup(uintptr, uintptr)

//go:wasmimport sdl3 SDL_PopGPUDebugGroup
func __SDL_PopGPUDebugGroup(uintptr)

//go:wasmimport sdl3 SDL_ReleaseGPUTexture
func __SDL_ReleaseGPUTexture(uintptr, uintptr)

//go:wasmimport sdl3 SDL_ReleaseGPUSampler
func __SDL_ReleaseGPUSampler(uintptr, uintptr)

//go:wasmimport sdl3 SDL_ReleaseGPUBuffer
func __SDL_ReleaseGPUBuffer(uintptr, uintptr)

//go:wasmimport sdl3 SDL_ReleaseGPUTransferBuffer
func __SDL_ReleaseGPUTransferBuffer(uintptr, uintptr)

//go:wasmimport sdl3 SDL_ReleaseGPUComputePipeline
func __SDL_ReleaseGPUComputePipeline(uintptr, uintptr)

//go:wasmimport sdl3 SDL_ReleaseGPUShader
func __SDL_ReleaseGPUShader(uintptr, uintptr)

//go:wasmimport sdl3 SDL_ReleaseGPUGraphicsPipeline
func __SDL_ReleaseGPUGraphicsPipeline(uintptr, uintptr)

//go:wasmimport sdl3 SDL_AcquireGPUCommandBuffer
func __SDL_AcquireGPUCommandBuffer(uintptr) uintptr

//go:wasmimport sdl3 SDL_PushGPUVertexUniformData
func __SDL_PushGPUVertexUniformData(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_PushGPUFragmentUniformData
func __SDL_PushGPUFragmentUniformData(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_PushGPUComputeUniformData
func __SDL_PushGPUComputeUniformData(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_BeginGPURenderPass
func __SDL_BeginGPURenderPass(uintptr, uintptr, int32, uintptr) uintptr

//go:wasmimport sdl3 SDL_BindGPUGraphicsPipeline
func __SDL_BindGPUGraphicsPipeline(uintptr, uintptr)

//go:wasmimport sdl3 SDL_SetGPUViewport
func __SDL_SetGPUViewport(uintptr, uintptr)

//go:wasmimport sdl3 SDL_SetGPUScissor
func __SDL_SetGPUScissor(uintptr, uintptr)

//go:wasmimport sdl3 SDL_SetGPUBlendConstants
func __SDL_SetGPUBlendConstants(uintptr, uintptr)

//go:wasmimport sdl3 SDL_SetGPUStencilReference
func __SDL_SetGPUStencilReference(uintptr, int32)

//go:wasmimport sdl3 SDL_BindGPUVertexBuffers
func __SDL_BindGPUVertexBuffers(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_BindGPUIndexBuffer
func __SDL_BindGPUIndexBuffer(uintptr, uintptr, int32)

//go:wasmimport sdl3 SDL_BindGPUVertexSamplers
func __SDL_BindGPUVertexSamplers(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_BindGPUVertexStorageTextures
func __SDL_BindGPUVertexStorageTextures(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_BindGPUVertexStorageBuffers
func __SDL_BindGPUVertexStorageBuffers(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_BindGPUFragmentSamplers
func __SDL_BindGPUFragmentSamplers(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_BindGPUFragmentStorageTextures
func __SDL_BindGPUFragmentStorageTextures(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_BindGPUFragmentStorageBuffers
func __SDL_BindGPUFragmentStorageBuffers(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_DrawGPUIndexedPrimitives
func __SDL_DrawGPUIndexedPrimitives(uintptr, int32, int32, int32, int32, int32)

//go:wasmimport sdl3 SDL_DrawGPUPrimitives
func __SDL_DrawGPUPrimitives(uintptr, int32, int32, int32, int32)

//go:wasmimport sdl3 SDL_DrawGPUPrimitivesIndirect
func __SDL_DrawGPUPrimitivesIndirect(uintptr, uintptr, int32, int32)

//go:wasmimport sdl3 SDL_DrawGPUIndexedPrimitivesIndirect
func __SDL_DrawGPUIndexedPrimitivesIndirect(uintptr, uintptr, int32, int32)

//go:wasmimport sdl3 SDL_EndGPURenderPass
func __SDL_EndGPURenderPass(uintptr)

//go:wasmimport sdl3 SDL_BeginGPUComputePass
func __SDL_BeginGPUComputePass(uintptr, uintptr, int32, uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_BindGPUComputePipeline
func __SDL_BindGPUComputePipeline(uintptr, uintptr)

//go:wasmimport sdl3 SDL_BindGPUComputeSamplers
func __SDL_BindGPUComputeSamplers(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_BindGPUComputeStorageTextures
func __SDL_BindGPUComputeStorageTextures(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_BindGPUComputeStorageBuffers
func __SDL_BindGPUComputeStorageBuffers(uintptr, int32, uintptr, int32)

//go:wasmimport sdl3 SDL_DispatchGPUCompute
func __SDL_DispatchGPUCompute(uintptr, int32, int32, int32)

//go:wasmimport sdl3 SDL_DispatchGPUComputeIndirect
func __SDL_DispatchGPUComputeIndirect(uintptr, uintptr, int32)

//go:wasmimport sdl3 SDL_EndGPUComputePass
func __SDL_EndGPUComputePass(uintptr)

//go:wasmimport sdl3 SDL_MapGPUTransferBuffer
func __SDL_MapGPUTransferBuffer(uintptr, uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_UnmapGPUTransferBuffer
func __SDL_UnmapGPUTransferBuffer(uintptr, uintptr)

//go:wasmimport sdl3 SDL_BeginGPUCopyPass
func __SDL_BeginGPUCopyPass(uintptr) uintptr

//go:wasmimport sdl3 SDL_UploadToGPUTexture
func __SDL_UploadToGPUTexture(uintptr, uintptr, uintptr, int32)

//go:wasmimport sdl3 SDL_UploadToGPUBuffer
func __SDL_UploadToGPUBuffer(uintptr, uintptr, uintptr, int32)

//go:wasmimport sdl3 SDL_CopyGPUTextureToTexture
func __SDL_CopyGPUTextureToTexture(uintptr, uintptr, uintptr, int32, int32, int32, int32)

//go:wasmimport sdl3 SDL_CopyGPUBufferToBuffer
func __SDL_CopyGPUBufferToBuffer(uintptr, uintptr, uintptr, int32, int32)

//go:wasmimport sdl3 SDL_DownloadFromGPUTexture
func __SDL_DownloadFromGPUTexture(uintptr, uintptr, uintptr)

//go:wasmimport sdl3 SDL_DownloadFromGPUBuffer
func __SDL_DownloadFromGPUBuffer(uintptr, uintptr, uintptr)

//go:wasmimport sdl3 SDL_EndGPUCopyPass
func __SDL_EndGPUCopyPass(uintptr)

//go:wasmimport sdl3 SDL_GenerateMipmapsForGPUTexture
func __SDL_GenerateMipmapsForGPUTexture(uintptr, uintptr)

//go:wasmimport sdl3 SDL_BlitGPUTexture
func __SDL_BlitGPUTexture(uintptr, uintptr)

//go:wasmimport sdl3 SDL_WindowSupportsGPUSwapchainComposition
func __SDL_WindowSupportsGPUSwapchainComposition(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_WindowSupportsGPUPresentMode
func __SDL_WindowSupportsGPUPresentMode(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_ClaimWindowForGPUDevice
func __SDL_ClaimWindowForGPUDevice(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReleaseWindowFromGPUDevice
func __SDL_ReleaseWindowFromGPUDevice(uintptr, uintptr)

//go:wasmimport sdl3 SDL_SetGPUSwapchainParameters
func __SDL_SetGPUSwapchainParameters(uintptr, uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_SetGPUAllowedFramesInFlight
func __SDL_SetGPUAllowedFramesInFlight(uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetGPUSwapchainTextureFormat
func __SDL_GetGPUSwapchainTextureFormat(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_AcquireGPUSwapchainTexture
func __SDL_AcquireGPUSwapchainTexture(uintptr, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_WaitForGPUSwapchain
func __SDL_WaitForGPUSwapchain(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_WaitAndAcquireGPUSwapchainTexture
func __SDL_WaitAndAcquireGPUSwapchainTexture(uintptr, uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SubmitGPUCommandBuffer
func __SDL_SubmitGPUCommandBuffer(uintptr) int32

//go:wasmimport sdl3 SDL_SubmitGPUCommandBufferAndAcquireFence
func __SDL_SubmitGPUCommandBufferAndAcquireFence(uintptr) uintptr

//go:wasmimport sdl3 SDL_CancelGPUCommandBuffer
func __SDL_CancelGPUCommandBuffer(uintptr) int32

//go:wasmimport sdl3 SDL_WaitForGPUIdle
func __SDL_WaitForGPUIdle(uintptr) int32

//go:wasmimport sdl3 SDL_WaitForGPUFences
func __SDL_WaitForGPUFences(uintptr, int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_QueryGPUFence
func __SDL_QueryGPUFence(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReleaseGPUFence
func __SDL_ReleaseGPUFence(uintptr, uintptr)

//go:wasmimport sdl3 SDL_GPUTextureFormatTexelBlockSize
func __SDL_GPUTextureFormatTexelBlockSize(int32) int32

//go:wasmimport sdl3 SDL_GPUTextureSupportsFormat
func __SDL_GPUTextureSupportsFormat(uintptr, int32, int32, int32) int32

//go:wasmimport sdl3 SDL_GPUTextureSupportsSampleCount
func __SDL_GPUTextureSupportsSampleCount(uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_CalculateGPUTextureFormatSize
func __SDL_CalculateGPUTextureFormatSize(int32, int32, int32, int32) int32



func __gowrap__SDL_GPUSupportsShaderFormats(format_flags GPUShaderFormat, name string) (__result bool) {
	__result = (bool)(__SDL_GPUSupportsShaderFormats(*(*int32)(unsafe.Pointer(&format_flags)), ((*[2]uintptr)(unsafe.Pointer(&name)))[0]) != 0)
	runtime.KeepAlive(name)
	return
}

func __gowrap__SDL_GPUSupportsProperties(props PropertiesID) (__result bool) {
	__result = (bool)(__SDL_GPUSupportsProperties(*(*int32)(unsafe.Pointer(&props))) != 0)
	return
}

func __gowrap__SDL_CreateGPUDevice(format_flags GPUShaderFormat, debug_mode bool, name string) (__result GPUDevice) {
	__result = (GPUDevice)(unsafe.Pointer(__SDL_CreateGPUDevice(*(*int32)(unsafe.Pointer(&format_flags)), *(*int32)(unsafe.Pointer(&debug_mode)), ((*[2]uintptr)(unsafe.Pointer(&name)))[0])))
	runtime.KeepAlive(name)
	return
}

func __gowrap__SDL_CreateGPUDeviceWithProperties(props PropertiesID) (__result GPUDevice) {
	__result = (GPUDevice)(unsafe.Pointer(__SDL_CreateGPUDeviceWithProperties(*(*int32)(unsafe.Pointer(&props)))))
	return
}

func __gowrap__SDL_DestroyGPUDevice(device GPUDevice) {
	__SDL_DestroyGPUDevice(uintptr(unsafe.Pointer(device)))
}

func __gowrap__SDL_GetNumGPUDrivers() (__result int32) {
	__result = (int32)(__SDL_GetNumGPUDrivers())
	return
}

func __gowrap__SDL_GetGPUDriver(index int32) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGPUDriver(*(*int32)(unsafe.Pointer(&index)))).Collect()))
	return
}

func __gowrap__SDL_GetGPUDeviceDriver(device GPUDevice) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetGPUDeviceDriver(uintptr(unsafe.Pointer(device)))).Collect()))
	return
}

func __gowrap__SDL_GetGPUShaderFormats(device GPUDevice) (__result GPUShaderFormat) {
	__result = (GPUShaderFormat)(__SDL_GetGPUShaderFormats(uintptr(unsafe.Pointer(device))))
	return
}

func __gowrap__SDL_GetGPUDeviceProperties(device GPUDevice) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetGPUDeviceProperties(uintptr(unsafe.Pointer(device))))
	return
}

func __gowrap__SDL_CreateGPUComputePipeline(device GPUDevice, createinfo *GPUComputePipelineCreateInfo) (__result GPUComputePipeline) {
	__result = (GPUComputePipeline)(unsafe.Pointer(__SDL_CreateGPUComputePipeline(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(createinfo)))))
	return
}

func __gowrap__SDL_CreateGPUGraphicsPipeline(device GPUDevice, createinfo *GPUGraphicsPipelineCreateInfo) (__result GPUGraphicsPipeline) {
	__result = (GPUGraphicsPipeline)(unsafe.Pointer(__SDL_CreateGPUGraphicsPipeline(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(createinfo)))))
	return
}

func __gowrap__SDL_CreateGPUSampler(device GPUDevice, createinfo *GPUSamplerCreateInfo) (__result GPUSampler) {
	__result = (GPUSampler)(unsafe.Pointer(__SDL_CreateGPUSampler(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(createinfo)))))
	return
}

func __gowrap__SDL_CreateGPUShader(device GPUDevice, createinfo *GPUShaderCreateInfo) (__result GPUShader) {
	__result = (GPUShader)(unsafe.Pointer(__SDL_CreateGPUShader(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(createinfo)))))
	return
}

func __gowrap__SDL_CreateGPUTexture(device GPUDevice, createinfo *GPUTextureCreateInfo) (__result GPUTexture) {
	__result = (GPUTexture)(unsafe.Pointer(__SDL_CreateGPUTexture(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(createinfo)))))
	return
}

func __gowrap__SDL_CreateGPUBuffer(device GPUDevice, createinfo *GPUBufferCreateInfo) (__result GPUBuffer) {
	__result = (GPUBuffer)(unsafe.Pointer(__SDL_CreateGPUBuffer(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(createinfo)))))
	return
}

func __gowrap__SDL_CreateGPUTransferBuffer(device GPUDevice, createinfo *GPUTransferBufferCreateInfo) (__result GPUTransferBuffer) {
	__result = (GPUTransferBuffer)(unsafe.Pointer(__SDL_CreateGPUTransferBuffer(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(createinfo)))))
	return
}

func __gowrap__SDL_SetGPUBufferName(device GPUDevice, buffer GPUBuffer, text string) {
	__SDL_SetGPUBufferName(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(buffer)), ((*[2]uintptr)(unsafe.Pointer(&text)))[0])
	runtime.KeepAlive(text)
}

func __gowrap__SDL_SetGPUTextureName(device GPUDevice, texture GPUTexture, text string) {
	__SDL_SetGPUTextureName(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(texture)), ((*[2]uintptr)(unsafe.Pointer(&text)))[0])
	runtime.KeepAlive(text)
}

func __gowrap__SDL_InsertGPUDebugLabel(command_buffer GPUCommandBuffer, text string) {
	__SDL_InsertGPUDebugLabel(uintptr(unsafe.Pointer(command_buffer)), ((*[2]uintptr)(unsafe.Pointer(&text)))[0])
	runtime.KeepAlive(text)
}

func __gowrap__SDL_PushGPUDebugGroup(command_buffer GPUCommandBuffer, name string) {
	__SDL_PushGPUDebugGroup(uintptr(unsafe.Pointer(command_buffer)), ((*[2]uintptr)(unsafe.Pointer(&name)))[0])
	runtime.KeepAlive(name)
}

func __gowrap__SDL_PopGPUDebugGroup(command_buffer GPUCommandBuffer) {
	__SDL_PopGPUDebugGroup(uintptr(unsafe.Pointer(command_buffer)))
}

func __gowrap__SDL_ReleaseGPUTexture(device GPUDevice, texture GPUTexture) {
	__SDL_ReleaseGPUTexture(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(texture)))
}

func __gowrap__SDL_ReleaseGPUSampler(device GPUDevice, sampler GPUSampler) {
	__SDL_ReleaseGPUSampler(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(sampler)))
}

func __gowrap__SDL_ReleaseGPUBuffer(device GPUDevice, buffer GPUBuffer) {
	__SDL_ReleaseGPUBuffer(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(buffer)))
}

func __gowrap__SDL_ReleaseGPUTransferBuffer(device GPUDevice, transfer_buffer GPUTransferBuffer) {
	__SDL_ReleaseGPUTransferBuffer(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(transfer_buffer)))
}

func __gowrap__SDL_ReleaseGPUComputePipeline(device GPUDevice, compute_pipeline GPUComputePipeline) {
	__SDL_ReleaseGPUComputePipeline(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(compute_pipeline)))
}

func __gowrap__SDL_ReleaseGPUShader(device GPUDevice, shader GPUShader) {
	__SDL_ReleaseGPUShader(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(shader)))
}

func __gowrap__SDL_ReleaseGPUGraphicsPipeline(device GPUDevice, graphics_pipeline GPUGraphicsPipeline) {
	__SDL_ReleaseGPUGraphicsPipeline(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(graphics_pipeline)))
}

func __gowrap__SDL_AcquireGPUCommandBuffer(device GPUDevice) (__result GPUCommandBuffer) {
	__result = (GPUCommandBuffer)(unsafe.Pointer(__SDL_AcquireGPUCommandBuffer(uintptr(unsafe.Pointer(device)))))
	return
}

func __gowrap__SDL_PushGPUVertexUniformData(command_buffer GPUCommandBuffer, slot_index uint32, data uintptr, length uint32) {
	__SDL_PushGPUVertexUniformData(uintptr(unsafe.Pointer(command_buffer)), *(*int32)(unsafe.Pointer(&slot_index)), uintptr(unsafe.Pointer(data)), *(*int32)(unsafe.Pointer(&length)))
}

func __gowrap__SDL_PushGPUFragmentUniformData(command_buffer GPUCommandBuffer, slot_index uint32, data uintptr, length uint32) {
	__SDL_PushGPUFragmentUniformData(uintptr(unsafe.Pointer(command_buffer)), *(*int32)(unsafe.Pointer(&slot_index)), uintptr(unsafe.Pointer(data)), *(*int32)(unsafe.Pointer(&length)))
}

func __gowrap__SDL_PushGPUComputeUniformData(command_buffer GPUCommandBuffer, slot_index uint32, data uintptr, length uint32) {
	__SDL_PushGPUComputeUniformData(uintptr(unsafe.Pointer(command_buffer)), *(*int32)(unsafe.Pointer(&slot_index)), uintptr(unsafe.Pointer(data)), *(*int32)(unsafe.Pointer(&length)))
}

func __gowrap__SDL_BeginGPURenderPass(command_buffer GPUCommandBuffer, color_target_infos *GPUColorTargetInfo, num_color_targets uint32, depth_stencil_target_info *GPUDepthStencilTargetInfo) (__result GPURenderPass) {
	__result = (GPURenderPass)(unsafe.Pointer(__SDL_BeginGPURenderPass(uintptr(unsafe.Pointer(command_buffer)), uintptr(unsafe.Pointer(color_target_infos)), *(*int32)(unsafe.Pointer(&num_color_targets)), uintptr(unsafe.Pointer(depth_stencil_target_info)))))
	return
}

func __gowrap__SDL_BindGPUGraphicsPipeline(render_pass GPURenderPass, graphics_pipeline GPUGraphicsPipeline) {
	__SDL_BindGPUGraphicsPipeline(uintptr(unsafe.Pointer(render_pass)), uintptr(unsafe.Pointer(graphics_pipeline)))
}

func __gowrap__SDL_SetGPUViewport(render_pass GPURenderPass, viewport *GPUViewport) {
	__SDL_SetGPUViewport(uintptr(unsafe.Pointer(render_pass)), uintptr(unsafe.Pointer(viewport)))
}

func __gowrap__SDL_SetGPUScissor(render_pass GPURenderPass, scissor *Rect) {
	__SDL_SetGPUScissor(uintptr(unsafe.Pointer(render_pass)), uintptr(unsafe.Pointer(scissor)))
}

func __gowrap__SDL_SetGPUBlendConstants(render_pass GPURenderPass, blend_constants *FColor) {
	__SDL_SetGPUBlendConstants(uintptr(unsafe.Pointer(render_pass)), uintptr(unsafe.Pointer(blend_constants)))
}

func __gowrap__SDL_SetGPUStencilReference(render_pass GPURenderPass, reference uint8) {
	__SDL_SetGPUStencilReference(uintptr(unsafe.Pointer(render_pass)), *(*int32)(unsafe.Pointer(&reference)))
}

func __gowrap__SDL_BindGPUVertexBuffers(render_pass GPURenderPass, first_slot uint32, bindings *GPUBufferBinding, num_bindings uint32) {
	__SDL_BindGPUVertexBuffers(uintptr(unsafe.Pointer(render_pass)), *(*int32)(unsafe.Pointer(&first_slot)), uintptr(unsafe.Pointer(bindings)), *(*int32)(unsafe.Pointer(&num_bindings)))
}

func __gowrap__SDL_BindGPUIndexBuffer(render_pass GPURenderPass, binding *GPUBufferBinding, index_element_size GPUIndexElementSize) {
	__SDL_BindGPUIndexBuffer(uintptr(unsafe.Pointer(render_pass)), uintptr(unsafe.Pointer(binding)), *(*int32)(unsafe.Pointer(&index_element_size)))
}

func __gowrap__SDL_BindGPUVertexSamplers(render_pass GPURenderPass, first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32) {
	__SDL_BindGPUVertexSamplers(uintptr(unsafe.Pointer(render_pass)), *(*int32)(unsafe.Pointer(&first_slot)), uintptr(unsafe.Pointer(texture_sampler_bindings)), *(*int32)(unsafe.Pointer(&num_bindings)))
}

func __gowrap__SDL_BindGPUVertexStorageTextures(render_pass GPURenderPass, first_slot uint32, storage_textures []GPUTexture, num_bindings uint32) {
	__SDL_BindGPUVertexStorageTextures(uintptr(unsafe.Pointer(render_pass)), *(*int32)(unsafe.Pointer(&first_slot)), uintptr(unsafe.Pointer(&storage_textures[0])), *(*int32)(unsafe.Pointer(&num_bindings)))
	runtime.KeepAlive(storage_textures)
}

func __gowrap__SDL_BindGPUVertexStorageBuffers(render_pass GPURenderPass, first_slot uint32, storage_buffers []GPUBuffer, num_bindings uint32) {
	__SDL_BindGPUVertexStorageBuffers(uintptr(unsafe.Pointer(render_pass)), *(*int32)(unsafe.Pointer(&first_slot)), uintptr(unsafe.Pointer(&storage_buffers[0])), *(*int32)(unsafe.Pointer(&num_bindings)))
	runtime.KeepAlive(storage_buffers)
}

func __gowrap__SDL_BindGPUFragmentSamplers(render_pass GPURenderPass, first_slot uint32, texture_sampler_bindings []GPUTextureSamplerBinding, num_bindings uint32) {
	__SDL_BindGPUFragmentSamplers(uintptr(unsafe.Pointer(render_pass)), *(*int32)(unsafe.Pointer(&first_slot)), uintptr(unsafe.Pointer(&texture_sampler_bindings[0])), *(*int32)(unsafe.Pointer(&num_bindings)))
	runtime.KeepAlive(texture_sampler_bindings)
}

func __gowrap__SDL_BindGPUFragmentStorageTextures(render_pass GPURenderPass, first_slot uint32, storage_textures []GPUTexture, num_bindings uint32) {
	__SDL_BindGPUFragmentStorageTextures(uintptr(unsafe.Pointer(render_pass)), *(*int32)(unsafe.Pointer(&first_slot)), uintptr(unsafe.Pointer(&storage_textures[0])), *(*int32)(unsafe.Pointer(&num_bindings)))
	runtime.KeepAlive(storage_textures)
}

func __gowrap__SDL_BindGPUFragmentStorageBuffers(render_pass GPURenderPass, first_slot uint32, storage_buffers []GPUBuffer, num_bindings uint32) {
	__SDL_BindGPUFragmentStorageBuffers(uintptr(unsafe.Pointer(render_pass)), *(*int32)(unsafe.Pointer(&first_slot)), uintptr(unsafe.Pointer(&storage_buffers[0])), *(*int32)(unsafe.Pointer(&num_bindings)))
	runtime.KeepAlive(storage_buffers)
}

func __gowrap__SDL_DrawGPUIndexedPrimitives(render_pass GPURenderPass, num_indices uint32, num_instances uint32, first_index uint32, vertex_offset int32, first_instance uint32) {
	__SDL_DrawGPUIndexedPrimitives(uintptr(unsafe.Pointer(render_pass)), *(*int32)(unsafe.Pointer(&num_indices)), *(*int32)(unsafe.Pointer(&num_instances)), *(*int32)(unsafe.Pointer(&first_index)), *(*int32)(unsafe.Pointer(&vertex_offset)), *(*int32)(unsafe.Pointer(&first_instance)))
}

func __gowrap__SDL_DrawGPUPrimitives(render_pass GPURenderPass, num_vertices uint32, num_instances uint32, first_vertex uint32, first_instance uint32) {
	__SDL_DrawGPUPrimitives(uintptr(unsafe.Pointer(render_pass)), *(*int32)(unsafe.Pointer(&num_vertices)), *(*int32)(unsafe.Pointer(&num_instances)), *(*int32)(unsafe.Pointer(&first_vertex)), *(*int32)(unsafe.Pointer(&first_instance)))
}

func __gowrap__SDL_DrawGPUPrimitivesIndirect(render_pass GPURenderPass, buffer GPUBuffer, offset uint32, draw_count uint32) {
	__SDL_DrawGPUPrimitivesIndirect(uintptr(unsafe.Pointer(render_pass)), uintptr(unsafe.Pointer(buffer)), *(*int32)(unsafe.Pointer(&offset)), *(*int32)(unsafe.Pointer(&draw_count)))
}

func __gowrap__SDL_DrawGPUIndexedPrimitivesIndirect(render_pass GPURenderPass, buffer GPUBuffer, offset uint32, draw_count uint32) {
	__SDL_DrawGPUIndexedPrimitivesIndirect(uintptr(unsafe.Pointer(render_pass)), uintptr(unsafe.Pointer(buffer)), *(*int32)(unsafe.Pointer(&offset)), *(*int32)(unsafe.Pointer(&draw_count)))
}

func __gowrap__SDL_EndGPURenderPass(render_pass GPURenderPass) {
	__SDL_EndGPURenderPass(uintptr(unsafe.Pointer(render_pass)))
}

func __gowrap__SDL_BeginGPUComputePass(command_buffer GPUCommandBuffer, storage_texture_bindings []GPUStorageTextureReadWriteBinding, num_storage_texture_bindings uint32, storage_buffer_bindings []GPUStorageBufferReadWriteBinding, num_storage_buffer_bindings uint32) (__result GPUComputePass) {
	__result = (GPUComputePass)(unsafe.Pointer(__SDL_BeginGPUComputePass(uintptr(unsafe.Pointer(command_buffer)), uintptr(unsafe.Pointer(&storage_texture_bindings[0])), *(*int32)(unsafe.Pointer(&num_storage_texture_bindings)), uintptr(unsafe.Pointer(&storage_buffer_bindings[0])), *(*int32)(unsafe.Pointer(&num_storage_buffer_bindings)))))
	runtime.KeepAlive(storage_texture_bindings)
	runtime.KeepAlive(storage_buffer_bindings)
	return
}

func __gowrap__SDL_BindGPUComputePipeline(compute_pass GPUComputePass, compute_pipeline GPUComputePipeline) {
	__SDL_BindGPUComputePipeline(uintptr(unsafe.Pointer(compute_pass)), uintptr(unsafe.Pointer(compute_pipeline)))
}

func __gowrap__SDL_BindGPUComputeSamplers(compute_pass GPUComputePass, first_slot uint32, texture_sampler_bindings []GPUTextureSamplerBinding, num_bindings uint32) {
	__SDL_BindGPUComputeSamplers(uintptr(unsafe.Pointer(compute_pass)), *(*int32)(unsafe.Pointer(&first_slot)), uintptr(unsafe.Pointer(&texture_sampler_bindings[0])), *(*int32)(unsafe.Pointer(&num_bindings)))
	runtime.KeepAlive(texture_sampler_bindings)
}

func __gowrap__SDL_BindGPUComputeStorageTextures(compute_pass GPUComputePass, first_slot uint32, storage_textures []GPUTexture, num_bindings uint32) {
	__SDL_BindGPUComputeStorageTextures(uintptr(unsafe.Pointer(compute_pass)), *(*int32)(unsafe.Pointer(&first_slot)), uintptr(unsafe.Pointer(&storage_textures[0])), *(*int32)(unsafe.Pointer(&num_bindings)))
	runtime.KeepAlive(storage_textures)
}

func __gowrap__SDL_BindGPUComputeStorageBuffers(compute_pass GPUComputePass, first_slot uint32, storage_buffers []GPUBuffer, num_bindings uint32) {
	__SDL_BindGPUComputeStorageBuffers(uintptr(unsafe.Pointer(compute_pass)), *(*int32)(unsafe.Pointer(&first_slot)), uintptr(unsafe.Pointer(&storage_buffers[0])), *(*int32)(unsafe.Pointer(&num_bindings)))
	runtime.KeepAlive(storage_buffers)
}

func __gowrap__SDL_DispatchGPUCompute(compute_pass GPUComputePass, groupcount_x uint32, groupcount_y uint32, groupcount_z uint32) {
	__SDL_DispatchGPUCompute(uintptr(unsafe.Pointer(compute_pass)), *(*int32)(unsafe.Pointer(&groupcount_x)), *(*int32)(unsafe.Pointer(&groupcount_y)), *(*int32)(unsafe.Pointer(&groupcount_z)))
}

func __gowrap__SDL_DispatchGPUComputeIndirect(compute_pass GPUComputePass, buffer GPUBuffer, offset uint32) {
	__SDL_DispatchGPUComputeIndirect(uintptr(unsafe.Pointer(compute_pass)), uintptr(unsafe.Pointer(buffer)), *(*int32)(unsafe.Pointer(&offset)))
}

func __gowrap__SDL_EndGPUComputePass(compute_pass GPUComputePass) {
	__SDL_EndGPUComputePass(uintptr(unsafe.Pointer(compute_pass)))
}

func __gowrap__SDL_MapGPUTransferBuffer(device GPUDevice, transfer_buffer GPUTransferBuffer, cycle bool) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_MapGPUTransferBuffer(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(transfer_buffer)), *(*int32)(unsafe.Pointer(&cycle)))))
	return
}

func __gowrap__SDL_UnmapGPUTransferBuffer(device GPUDevice, transfer_buffer GPUTransferBuffer) {
	__SDL_UnmapGPUTransferBuffer(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(transfer_buffer)))
}

func __gowrap__SDL_BeginGPUCopyPass(command_buffer GPUCommandBuffer) (__result GPUCopyPass) {
	__result = (GPUCopyPass)(unsafe.Pointer(__SDL_BeginGPUCopyPass(uintptr(unsafe.Pointer(command_buffer)))))
	return
}

func __gowrap__SDL_UploadToGPUTexture(copy_pass GPUCopyPass, source *GPUTextureTransferInfo, destination *GPUTextureRegion, cycle bool) {
	__SDL_UploadToGPUTexture(uintptr(unsafe.Pointer(copy_pass)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)), *(*int32)(unsafe.Pointer(&cycle)))
}

func __gowrap__SDL_UploadToGPUBuffer(copy_pass GPUCopyPass, source *GPUTransferBufferLocation, destination *GPUBufferRegion, cycle bool) {
	__SDL_UploadToGPUBuffer(uintptr(unsafe.Pointer(copy_pass)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)), *(*int32)(unsafe.Pointer(&cycle)))
}

func __gowrap__SDL_CopyGPUTextureToTexture(copy_pass GPUCopyPass, source *GPUTextureLocation, destination *GPUTextureLocation, w uint32, h uint32, d uint32, cycle bool) {
	__SDL_CopyGPUTextureToTexture(uintptr(unsafe.Pointer(copy_pass)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)), *(*int32)(unsafe.Pointer(&w)), *(*int32)(unsafe.Pointer(&h)), *(*int32)(unsafe.Pointer(&d)), *(*int32)(unsafe.Pointer(&cycle)))
}

func __gowrap__SDL_CopyGPUBufferToBuffer(copy_pass GPUCopyPass, source *GPUBufferLocation, destination *GPUBufferLocation, size uint32, cycle bool) {
	__SDL_CopyGPUBufferToBuffer(uintptr(unsafe.Pointer(copy_pass)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)), *(*int32)(unsafe.Pointer(&size)), *(*int32)(unsafe.Pointer(&cycle)))
}

func __gowrap__SDL_DownloadFromGPUTexture(copy_pass GPUCopyPass, source *GPUTextureRegion, destination *GPUTextureTransferInfo) {
	__SDL_DownloadFromGPUTexture(uintptr(unsafe.Pointer(copy_pass)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)))
}

func __gowrap__SDL_DownloadFromGPUBuffer(copy_pass GPUCopyPass, source *GPUBufferRegion, destination *GPUTransferBufferLocation) {
	__SDL_DownloadFromGPUBuffer(uintptr(unsafe.Pointer(copy_pass)), uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)))
}

func __gowrap__SDL_EndGPUCopyPass(copy_pass GPUCopyPass) {
	__SDL_EndGPUCopyPass(uintptr(unsafe.Pointer(copy_pass)))
}

func __gowrap__SDL_GenerateMipmapsForGPUTexture(command_buffer GPUCommandBuffer, texture GPUTexture) {
	__SDL_GenerateMipmapsForGPUTexture(uintptr(unsafe.Pointer(command_buffer)), uintptr(unsafe.Pointer(texture)))
}

func __gowrap__SDL_BlitGPUTexture(command_buffer GPUCommandBuffer, info *GPUBlitInfo) {
	__SDL_BlitGPUTexture(uintptr(unsafe.Pointer(command_buffer)), uintptr(unsafe.Pointer(info)))
}

func __gowrap__SDL_WindowSupportsGPUSwapchainComposition(device GPUDevice, window Window, swapchain_composition GPUSwapchainComposition) (__result bool) {
	__result = (bool)(__SDL_WindowSupportsGPUSwapchainComposition(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&swapchain_composition))) != 0)
	return
}

func __gowrap__SDL_WindowSupportsGPUPresentMode(device GPUDevice, window Window, present_mode GPUPresentMode) (__result bool) {
	__result = (bool)(__SDL_WindowSupportsGPUPresentMode(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&present_mode))) != 0)
	return
}

func __gowrap__SDL_ClaimWindowForGPUDevice(device GPUDevice, window Window) (__result bool) {
	__result = (bool)(__SDL_ClaimWindowForGPUDevice(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_ReleaseWindowFromGPUDevice(device GPUDevice, window Window) {
	__SDL_ReleaseWindowFromGPUDevice(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(window)))
}

func __gowrap__SDL_SetGPUSwapchainParameters(device GPUDevice, window Window, swapchain_composition GPUSwapchainComposition, present_mode GPUPresentMode) (__result bool) {
	__result = (bool)(__SDL_SetGPUSwapchainParameters(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(window)), *(*int32)(unsafe.Pointer(&swapchain_composition)), *(*int32)(unsafe.Pointer(&present_mode))) != 0)
	return
}

func __gowrap__SDL_SetGPUAllowedFramesInFlight(device GPUDevice, allowed_frames_in_flight uint32) (__result bool) {
	__result = (bool)(__SDL_SetGPUAllowedFramesInFlight(uintptr(unsafe.Pointer(device)), *(*int32)(unsafe.Pointer(&allowed_frames_in_flight))) != 0)
	return
}

func __gowrap__SDL_GetGPUSwapchainTextureFormat(device GPUDevice, window Window) (__result GPUTextureFormat) {
	__result = (GPUTextureFormat)(__SDL_GetGPUSwapchainTextureFormat(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(window))))
	return
}

func __gowrap__SDL_AcquireGPUSwapchainTexture(command_buffer GPUCommandBuffer, window Window, swapchain_texture GPUTexture, swapchain_texture_width *uint32, swapchain_texture_height *uint32) (__result bool) {
	__result = (bool)(__SDL_AcquireGPUSwapchainTexture(uintptr(unsafe.Pointer(command_buffer)), uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(swapchain_texture)), uintptr(unsafe.Pointer(swapchain_texture_width)), uintptr(unsafe.Pointer(swapchain_texture_height))) != 0)
	return
}

func __gowrap__SDL_WaitForGPUSwapchain(device GPUDevice, window Window) (__result bool) {
	__result = (bool)(__SDL_WaitForGPUSwapchain(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(window))) != 0)
	return
}

func __gowrap__SDL_WaitAndAcquireGPUSwapchainTexture(command_buffer GPUCommandBuffer, window Window, swapchain_texture GPUTexture, swapchain_texture_width *uint32, swapchain_texture_height *uint32) (__result bool) {
	__result = (bool)(__SDL_WaitAndAcquireGPUSwapchainTexture(uintptr(unsafe.Pointer(command_buffer)), uintptr(unsafe.Pointer(window)), uintptr(unsafe.Pointer(swapchain_texture)), uintptr(unsafe.Pointer(swapchain_texture_width)), uintptr(unsafe.Pointer(swapchain_texture_height))) != 0)
	return
}

func __gowrap__SDL_SubmitGPUCommandBuffer(command_buffer GPUCommandBuffer) (__result bool) {
	__result = (bool)(__SDL_SubmitGPUCommandBuffer(uintptr(unsafe.Pointer(command_buffer))) != 0)
	return
}

func __gowrap__SDL_SubmitGPUCommandBufferAndAcquireFence(command_buffer GPUCommandBuffer) (__result GPUFence) {
	__result = (GPUFence)(unsafe.Pointer(__SDL_SubmitGPUCommandBufferAndAcquireFence(uintptr(unsafe.Pointer(command_buffer)))))
	return
}

func __gowrap__SDL_CancelGPUCommandBuffer(command_buffer GPUCommandBuffer) (__result bool) {
	__result = (bool)(__SDL_CancelGPUCommandBuffer(uintptr(unsafe.Pointer(command_buffer))) != 0)
	return
}

func __gowrap__SDL_WaitForGPUIdle(device GPUDevice) (__result bool) {
	__result = (bool)(__SDL_WaitForGPUIdle(uintptr(unsafe.Pointer(device))) != 0)
	return
}

func __gowrap__SDL_WaitForGPUFences(device GPUDevice, wait_all bool, fences []GPUFence, num_fences uint32) (__result bool) {
	__result = (bool)(__SDL_WaitForGPUFences(uintptr(unsafe.Pointer(device)), *(*int32)(unsafe.Pointer(&wait_all)), uintptr(unsafe.Pointer(&fences[0])), *(*int32)(unsafe.Pointer(&num_fences))) != 0)
	runtime.KeepAlive(fences)
	return
}

func __gowrap__SDL_QueryGPUFence(device GPUDevice, fence GPUFence) (__result bool) {
	__result = (bool)(__SDL_QueryGPUFence(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(fence))) != 0)
	return
}

func __gowrap__SDL_ReleaseGPUFence(device GPUDevice, fence GPUFence) {
	__SDL_ReleaseGPUFence(uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(fence)))
}

func __gowrap__SDL_GPUTextureFormatTexelBlockSize(format GPUTextureFormat) (__result uint32) {
	__result = (uint32)(__SDL_GPUTextureFormatTexelBlockSize(*(*int32)(unsafe.Pointer(&format))))
	return
}

func __gowrap__SDL_GPUTextureSupportsFormat(device GPUDevice, format GPUTextureFormat, type_ GPUTextureType, usage GPUTextureUsageFlags) (__result bool) {
	__result = (bool)(__SDL_GPUTextureSupportsFormat(uintptr(unsafe.Pointer(device)), *(*int32)(unsafe.Pointer(&format)), *(*int32)(unsafe.Pointer(&type_)), *(*int32)(unsafe.Pointer(&usage))) != 0)
	return
}

func __gowrap__SDL_GPUTextureSupportsSampleCount(device GPUDevice, format GPUTextureFormat, sample_count GPUSampleCount) (__result bool) {
	__result = (bool)(__SDL_GPUTextureSupportsSampleCount(uintptr(unsafe.Pointer(device)), *(*int32)(unsafe.Pointer(&format)), *(*int32)(unsafe.Pointer(&sample_count))) != 0)
	return
}

func __gowrap__SDL_CalculateGPUTextureFormatSize(format GPUTextureFormat, width uint32, height uint32, depth_or_layer_count uint32) (__result uint32) {
	__result = (uint32)(__SDL_CalculateGPUTextureFormatSize(*(*int32)(unsafe.Pointer(&format)), *(*int32)(unsafe.Pointer(&width)), *(*int32)(unsafe.Pointer(&height)), *(*int32)(unsafe.Pointer(&depth_or_layer_count))))
	return
}

 
func init() {
	GPUSupportsShaderFormats = __gowrap__SDL_GPUSupportsShaderFormats
	GPUSupportsProperties = __gowrap__SDL_GPUSupportsProperties
	CreateGPUDevice = __gowrap__SDL_CreateGPUDevice
	CreateGPUDeviceWithProperties = __gowrap__SDL_CreateGPUDeviceWithProperties
	DestroyGPUDevice = __gowrap__SDL_DestroyGPUDevice
	GetNumGPUDrivers = __gowrap__SDL_GetNumGPUDrivers
	GetGPUDriver = __gowrap__SDL_GetGPUDriver
	GetGPUDeviceDriver = __gowrap__SDL_GetGPUDeviceDriver
	GetGPUShaderFormats = __gowrap__SDL_GetGPUShaderFormats
	GetGPUDeviceProperties = __gowrap__SDL_GetGPUDeviceProperties
	CreateGPUComputePipeline = __gowrap__SDL_CreateGPUComputePipeline
	CreateGPUGraphicsPipeline = __gowrap__SDL_CreateGPUGraphicsPipeline
	CreateGPUSampler = __gowrap__SDL_CreateGPUSampler
	CreateGPUShader = __gowrap__SDL_CreateGPUShader
	CreateGPUTexture = __gowrap__SDL_CreateGPUTexture
	CreateGPUBuffer = __gowrap__SDL_CreateGPUBuffer
	CreateGPUTransferBuffer = __gowrap__SDL_CreateGPUTransferBuffer
	SetGPUBufferName = __gowrap__SDL_SetGPUBufferName
	SetGPUTextureName = __gowrap__SDL_SetGPUTextureName
	InsertGPUDebugLabel = __gowrap__SDL_InsertGPUDebugLabel
	PushGPUDebugGroup = __gowrap__SDL_PushGPUDebugGroup
	PopGPUDebugGroup = __gowrap__SDL_PopGPUDebugGroup
	ReleaseGPUTexture = __gowrap__SDL_ReleaseGPUTexture
	ReleaseGPUSampler = __gowrap__SDL_ReleaseGPUSampler
	ReleaseGPUBuffer = __gowrap__SDL_ReleaseGPUBuffer
	ReleaseGPUTransferBuffer = __gowrap__SDL_ReleaseGPUTransferBuffer
	ReleaseGPUComputePipeline = __gowrap__SDL_ReleaseGPUComputePipeline
	ReleaseGPUShader = __gowrap__SDL_ReleaseGPUShader
	ReleaseGPUGraphicsPipeline = __gowrap__SDL_ReleaseGPUGraphicsPipeline
	AcquireGPUCommandBuffer = __gowrap__SDL_AcquireGPUCommandBuffer
	PushGPUVertexUniformData = __gowrap__SDL_PushGPUVertexUniformData
	PushGPUFragmentUniformData = __gowrap__SDL_PushGPUFragmentUniformData
	PushGPUComputeUniformData = __gowrap__SDL_PushGPUComputeUniformData
	BeginGPURenderPass = __gowrap__SDL_BeginGPURenderPass
	BindGPUGraphicsPipeline = __gowrap__SDL_BindGPUGraphicsPipeline
	SetGPUViewport = __gowrap__SDL_SetGPUViewport
	SetGPUScissor = __gowrap__SDL_SetGPUScissor
	SetGPUBlendConstants = __gowrap__SDL_SetGPUBlendConstants
	SetGPUStencilReference = __gowrap__SDL_SetGPUStencilReference
	BindGPUVertexBuffers = __gowrap__SDL_BindGPUVertexBuffers
	BindGPUIndexBuffer = __gowrap__SDL_BindGPUIndexBuffer
	BindGPUVertexSamplers = __gowrap__SDL_BindGPUVertexSamplers
	BindGPUVertexStorageTextures = __gowrap__SDL_BindGPUVertexStorageTextures
	BindGPUVertexStorageBuffers = __gowrap__SDL_BindGPUVertexStorageBuffers
	BindGPUFragmentSamplers = __gowrap__SDL_BindGPUFragmentSamplers
	BindGPUFragmentStorageTextures = __gowrap__SDL_BindGPUFragmentStorageTextures
	BindGPUFragmentStorageBuffers = __gowrap__SDL_BindGPUFragmentStorageBuffers
	DrawGPUIndexedPrimitives = __gowrap__SDL_DrawGPUIndexedPrimitives
	DrawGPUPrimitives = __gowrap__SDL_DrawGPUPrimitives
	DrawGPUPrimitivesIndirect = __gowrap__SDL_DrawGPUPrimitivesIndirect
	DrawGPUIndexedPrimitivesIndirect = __gowrap__SDL_DrawGPUIndexedPrimitivesIndirect
	EndGPURenderPass = __gowrap__SDL_EndGPURenderPass
	BeginGPUComputePass = __gowrap__SDL_BeginGPUComputePass
	BindGPUComputePipeline = __gowrap__SDL_BindGPUComputePipeline
	BindGPUComputeSamplers = __gowrap__SDL_BindGPUComputeSamplers
	BindGPUComputeStorageTextures = __gowrap__SDL_BindGPUComputeStorageTextures
	BindGPUComputeStorageBuffers = __gowrap__SDL_BindGPUComputeStorageBuffers
	DispatchGPUCompute = __gowrap__SDL_DispatchGPUCompute
	DispatchGPUComputeIndirect = __gowrap__SDL_DispatchGPUComputeIndirect
	EndGPUComputePass = __gowrap__SDL_EndGPUComputePass
	MapGPUTransferBuffer = __gowrap__SDL_MapGPUTransferBuffer
	UnmapGPUTransferBuffer = __gowrap__SDL_UnmapGPUTransferBuffer
	BeginGPUCopyPass = __gowrap__SDL_BeginGPUCopyPass
	UploadToGPUTexture = __gowrap__SDL_UploadToGPUTexture
	UploadToGPUBuffer = __gowrap__SDL_UploadToGPUBuffer
	CopyGPUTextureToTexture = __gowrap__SDL_CopyGPUTextureToTexture
	CopyGPUBufferToBuffer = __gowrap__SDL_CopyGPUBufferToBuffer
	DownloadFromGPUTexture = __gowrap__SDL_DownloadFromGPUTexture
	DownloadFromGPUBuffer = __gowrap__SDL_DownloadFromGPUBuffer
	EndGPUCopyPass = __gowrap__SDL_EndGPUCopyPass
	GenerateMipmapsForGPUTexture = __gowrap__SDL_GenerateMipmapsForGPUTexture
	BlitGPUTexture = __gowrap__SDL_BlitGPUTexture
	WindowSupportsGPUSwapchainComposition = __gowrap__SDL_WindowSupportsGPUSwapchainComposition
	WindowSupportsGPUPresentMode = __gowrap__SDL_WindowSupportsGPUPresentMode
	ClaimWindowForGPUDevice = __gowrap__SDL_ClaimWindowForGPUDevice
	ReleaseWindowFromGPUDevice = __gowrap__SDL_ReleaseWindowFromGPUDevice
	SetGPUSwapchainParameters = __gowrap__SDL_SetGPUSwapchainParameters
	SetGPUAllowedFramesInFlight = __gowrap__SDL_SetGPUAllowedFramesInFlight
	GetGPUSwapchainTextureFormat = __gowrap__SDL_GetGPUSwapchainTextureFormat
	AcquireGPUSwapchainTexture = __gowrap__SDL_AcquireGPUSwapchainTexture
	WaitForGPUSwapchain = __gowrap__SDL_WaitForGPUSwapchain
	WaitAndAcquireGPUSwapchainTexture = __gowrap__SDL_WaitAndAcquireGPUSwapchainTexture
	SubmitGPUCommandBuffer = __gowrap__SDL_SubmitGPUCommandBuffer
	SubmitGPUCommandBufferAndAcquireFence = __gowrap__SDL_SubmitGPUCommandBufferAndAcquireFence
	CancelGPUCommandBuffer = __gowrap__SDL_CancelGPUCommandBuffer
	WaitForGPUIdle = __gowrap__SDL_WaitForGPUIdle
	WaitForGPUFences = __gowrap__SDL_WaitForGPUFences
	QueryGPUFence = __gowrap__SDL_QueryGPUFence
	ReleaseGPUFence = __gowrap__SDL_ReleaseGPUFence
	GPUTextureFormatTexelBlockSize = __gowrap__SDL_GPUTextureFormatTexelBlockSize
	GPUTextureSupportsFormat = __gowrap__SDL_GPUTextureSupportsFormat
	GPUTextureSupportsSampleCount = __gowrap__SDL_GPUTextureSupportsSampleCount
	CalculateGPUTextureFormatSize = __gowrap__SDL_CalculateGPUTextureFormatSize
}
