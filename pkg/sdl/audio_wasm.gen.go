//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_GetNumAudioDrivers
func __SDL_GetNumAudioDrivers() int32

//go:wasmimport sdl3 SDL_GetAudioDriver
func __SDL_GetAudioDriver(int32) uintptr

//go:wasmimport sdl3 SDL_GetCurrentAudioDriver
func __SDL_GetCurrentAudioDriver() uintptr

//go:wasmimport sdl3 SDL_GetAudioPlaybackDevices
func __SDL_GetAudioPlaybackDevices(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetAudioRecordingDevices
func __SDL_GetAudioRecordingDevices(uintptr) uintptr

//go:wasmimport sdl3 SDL_GetAudioDeviceName
func __SDL_GetAudioDeviceName(int32) uintptr

//go:wasmimport sdl3 SDL_GetAudioDeviceFormat
func __SDL_GetAudioDeviceFormat(int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetAudioDeviceChannelMap
func __SDL_GetAudioDeviceChannelMap(int32, int32) uintptr

//go:wasmimport sdl3 SDL_OpenAudioDevice
func __SDL_OpenAudioDevice(int32, uintptr) int32

//go:wasmimport sdl3 SDL_IsAudioDevicePhysical
func __SDL_IsAudioDevicePhysical(int32) int32

//go:wasmimport sdl3 SDL_IsAudioDevicePlayback
func __SDL_IsAudioDevicePlayback(int32) int32

//go:wasmimport sdl3 SDL_PauseAudioDevice
func __SDL_PauseAudioDevice(int32) int32

//go:wasmimport sdl3 SDL_ResumeAudioDevice
func __SDL_ResumeAudioDevice(int32) int32

//go:wasmimport sdl3 SDL_AudioDevicePaused
func __SDL_AudioDevicePaused(int32) int32

//go:wasmimport sdl3 SDL_GetAudioDeviceGain
func __SDL_GetAudioDeviceGain(int32) float32

//go:wasmimport sdl3 SDL_SetAudioDeviceGain
func __SDL_SetAudioDeviceGain(int32, float32) int32

//go:wasmimport sdl3 SDL_CloseAudioDevice
func __SDL_CloseAudioDevice(int32)

//go:wasmimport sdl3 SDL_BindAudioStreams
func __SDL_BindAudioStreams(int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_BindAudioStream
func __SDL_BindAudioStream(int32, uintptr) int32

//go:wasmimport sdl3 SDL_UnbindAudioStreams
func __SDL_UnbindAudioStreams(uintptr, int32)

//go:wasmimport sdl3 SDL_UnbindAudioStream
func __SDL_UnbindAudioStream(uintptr)

//go:wasmimport sdl3 SDL_GetAudioStreamDevice
func __SDL_GetAudioStreamDevice(uintptr) int32

//go:wasmimport sdl3 SDL_CreateAudioStream
func __SDL_CreateAudioStream(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_GetAudioStreamProperties
func __SDL_GetAudioStreamProperties(uintptr) int32

//go:wasmimport sdl3 SDL_GetAudioStreamFormat
func __SDL_GetAudioStreamFormat(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetAudioStreamFormat
func __SDL_SetAudioStreamFormat(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetAudioStreamFrequencyRatio
func __SDL_GetAudioStreamFrequencyRatio(uintptr) float32

//go:wasmimport sdl3 SDL_SetAudioStreamFrequencyRatio
func __SDL_SetAudioStreamFrequencyRatio(uintptr, float32) int32

//go:wasmimport sdl3 SDL_GetAudioStreamGain
func __SDL_GetAudioStreamGain(uintptr) float32

//go:wasmimport sdl3 SDL_SetAudioStreamGain
func __SDL_SetAudioStreamGain(uintptr, float32) int32

//go:wasmimport sdl3 SDL_GetAudioStreamInputChannelMap
func __SDL_GetAudioStreamInputChannelMap(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_GetAudioStreamOutputChannelMap
func __SDL_GetAudioStreamOutputChannelMap(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_SetAudioStreamInputChannelMap
func __SDL_SetAudioStreamInputChannelMap(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_SetAudioStreamOutputChannelMap
func __SDL_SetAudioStreamOutputChannelMap(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_PutAudioStreamData
func __SDL_PutAudioStreamData(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_PutAudioStreamPlanarData
func __SDL_PutAudioStreamPlanarData(uintptr, uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_GetAudioStreamData
func __SDL_GetAudioStreamData(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_GetAudioStreamAvailable
func __SDL_GetAudioStreamAvailable(uintptr) int32

//go:wasmimport sdl3 SDL_GetAudioStreamQueued
func __SDL_GetAudioStreamQueued(uintptr) int32

//go:wasmimport sdl3 SDL_FlushAudioStream
func __SDL_FlushAudioStream(uintptr) int32

//go:wasmimport sdl3 SDL_ClearAudioStream
func __SDL_ClearAudioStream(uintptr) int32

//go:wasmimport sdl3 SDL_PauseAudioStreamDevice
func __SDL_PauseAudioStreamDevice(uintptr) int32

//go:wasmimport sdl3 SDL_ResumeAudioStreamDevice
func __SDL_ResumeAudioStreamDevice(uintptr) int32

//go:wasmimport sdl3 SDL_AudioStreamDevicePaused
func __SDL_AudioStreamDevicePaused(uintptr) int32

//go:wasmimport sdl3 SDL_LockAudioStream
func __SDL_LockAudioStream(uintptr) int32

//go:wasmimport sdl3 SDL_UnlockAudioStream
func __SDL_UnlockAudioStream(uintptr) int32

//go:wasmimport sdl3 SDL_SetAudioStreamGetCallback
func __SDL_SetAudioStreamGetCallback(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetAudioStreamPutCallback
func __SDL_SetAudioStreamPutCallback(uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_DestroyAudioStream
func __SDL_DestroyAudioStream(uintptr)

//go:wasmimport sdl3 SDL_OpenAudioDeviceStream
func __SDL_OpenAudioDeviceStream(int32, uintptr, uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_SetAudioIterationCallbacks
func __SDL_SetAudioIterationCallbacks(int32, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_SetAudioPostmixCallback
func __SDL_SetAudioPostmixCallback(int32, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_LoadWAV_IO
func __SDL_LoadWAV_IO(uintptr, int32, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_LoadWAV
func __SDL_LoadWAV(uintptr, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_MixAudio
func __SDL_MixAudio(uintptr, uintptr, int32, int32, float32) int32

//go:wasmimport sdl3 SDL_ConvertAudioSamples
func __SDL_ConvertAudioSamples(uintptr, uintptr, int32, uintptr, uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_GetAudioFormatName
func __SDL_GetAudioFormatName(int32) uintptr

//go:wasmimport sdl3 SDL_GetSilenceValueForFormat
func __SDL_GetSilenceValueForFormat(int32) int32



func __gowrap__SDL_GetNumAudioDrivers() (__result int32) {
	__result = (int32)(__SDL_GetNumAudioDrivers())
	return
}

func __gowrap__SDL_GetAudioDriver(index int32) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetAudioDriver(*(*int32)(unsafe.Pointer(&index)))).Collect()))
	return
}

func __gowrap__SDL_GetCurrentAudioDriver() (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetCurrentAudioDriver()).Collect()))
	return
}

func __gowrap__SDL_GetAudioPlaybackDevices(count *int32) (__result *AudioDeviceID) {
	__result = (*AudioDeviceID)(unsafe.Pointer(__SDL_GetAudioPlaybackDevices(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetAudioRecordingDevices(count *int32) (__result *AudioDeviceID) {
	__result = (*AudioDeviceID)(unsafe.Pointer(__SDL_GetAudioRecordingDevices(uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetAudioDeviceName(devid AudioDeviceID) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetAudioDeviceName(*(*int32)(unsafe.Pointer(&devid)))).Collect()))
	return
}

func __gowrap__SDL_GetAudioDeviceFormat(devid AudioDeviceID, spec *AudioSpec, sample_frames *int32) (__result bool) {
	__result = (bool)(__SDL_GetAudioDeviceFormat(*(*int32)(unsafe.Pointer(&devid)), uintptr(unsafe.Pointer(spec)), uintptr(unsafe.Pointer(sample_frames))) != 0)
	return
}

func __gowrap__SDL_GetAudioDeviceChannelMap(devid AudioDeviceID, count int32) (__result *int32) {
	__result = (*int32)(unsafe.Pointer(__SDL_GetAudioDeviceChannelMap(*(*int32)(unsafe.Pointer(&devid)), *(*int32)(unsafe.Pointer(&count)))))
	return
}

func __gowrap__SDL_OpenAudioDevice(devid AudioDeviceID, spec *AudioSpec) (__result AudioDeviceID) {
	__result = (AudioDeviceID)(__SDL_OpenAudioDevice(*(*int32)(unsafe.Pointer(&devid)), uintptr(unsafe.Pointer(spec))))
	return
}

func __gowrap__SDL_IsAudioDevicePhysical(devid AudioDeviceID) (__result bool) {
	__result = (bool)(__SDL_IsAudioDevicePhysical(*(*int32)(unsafe.Pointer(&devid))) != 0)
	return
}

func __gowrap__SDL_IsAudioDevicePlayback(devid AudioDeviceID) (__result bool) {
	__result = (bool)(__SDL_IsAudioDevicePlayback(*(*int32)(unsafe.Pointer(&devid))) != 0)
	return
}

func __gowrap__SDL_PauseAudioDevice(devid AudioDeviceID) (__result bool) {
	__result = (bool)(__SDL_PauseAudioDevice(*(*int32)(unsafe.Pointer(&devid))) != 0)
	return
}

func __gowrap__SDL_ResumeAudioDevice(devid AudioDeviceID) (__result bool) {
	__result = (bool)(__SDL_ResumeAudioDevice(*(*int32)(unsafe.Pointer(&devid))) != 0)
	return
}

func __gowrap__SDL_AudioDevicePaused(devid AudioDeviceID) (__result bool) {
	__result = (bool)(__SDL_AudioDevicePaused(*(*int32)(unsafe.Pointer(&devid))) != 0)
	return
}

func __gowrap__SDL_GetAudioDeviceGain(devid AudioDeviceID) (__result float32) {
	__result = (float32)(__SDL_GetAudioDeviceGain(*(*int32)(unsafe.Pointer(&devid))))
	return
}

func __gowrap__SDL_SetAudioDeviceGain(devid AudioDeviceID, gain float32) (__result bool) {
	__result = (bool)(__SDL_SetAudioDeviceGain(*(*int32)(unsafe.Pointer(&devid)), *(*float32)(unsafe.Pointer(&gain))) != 0)
	return
}

func __gowrap__SDL_CloseAudioDevice(devid AudioDeviceID) {
	__SDL_CloseAudioDevice(*(*int32)(unsafe.Pointer(&devid)))
}

func __gowrap__SDL_BindAudioStreams(devid AudioDeviceID, streams []AudioStream, num_streams int32) (__result bool) {
	__result = (bool)(__SDL_BindAudioStreams(*(*int32)(unsafe.Pointer(&devid)), uintptr(unsafe.Pointer(&streams[0])), *(*int32)(unsafe.Pointer(&num_streams))) != 0)
	runtime.KeepAlive(streams)
	return
}

func __gowrap__SDL_BindAudioStream(devid AudioDeviceID, stream AudioStream) (__result bool) {
	__result = (bool)(__SDL_BindAudioStream(*(*int32)(unsafe.Pointer(&devid)), uintptr(unsafe.Pointer(stream))) != 0)
	return
}

func __gowrap__SDL_UnbindAudioStreams(streams []AudioStream, num_streams int32) {
	__SDL_UnbindAudioStreams(uintptr(unsafe.Pointer(&streams[0])), *(*int32)(unsafe.Pointer(&num_streams)))
	runtime.KeepAlive(streams)
}

func __gowrap__SDL_UnbindAudioStream(stream AudioStream) {
	__SDL_UnbindAudioStream(uintptr(unsafe.Pointer(stream)))
}

func __gowrap__SDL_GetAudioStreamDevice(stream AudioStream) (__result AudioDeviceID) {
	__result = (AudioDeviceID)(__SDL_GetAudioStreamDevice(uintptr(unsafe.Pointer(stream))))
	return
}

func __gowrap__SDL_CreateAudioStream(src_spec *AudioSpec, dst_spec *AudioSpec) (__result AudioStream) {
	__result = (AudioStream)(unsafe.Pointer(__SDL_CreateAudioStream(uintptr(unsafe.Pointer(src_spec)), uintptr(unsafe.Pointer(dst_spec)))))
	return
}

func __gowrap__SDL_GetAudioStreamProperties(stream AudioStream) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetAudioStreamProperties(uintptr(unsafe.Pointer(stream))))
	return
}

func __gowrap__SDL_GetAudioStreamFormat(stream AudioStream, src_spec *AudioSpec, dst_spec *AudioSpec) (__result bool) {
	__result = (bool)(__SDL_GetAudioStreamFormat(uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(src_spec)), uintptr(unsafe.Pointer(dst_spec))) != 0)
	return
}

func __gowrap__SDL_SetAudioStreamFormat(stream AudioStream, src_spec *AudioSpec, dst_spec *AudioSpec) (__result bool) {
	__result = (bool)(__SDL_SetAudioStreamFormat(uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(src_spec)), uintptr(unsafe.Pointer(dst_spec))) != 0)
	return
}

func __gowrap__SDL_GetAudioStreamFrequencyRatio(stream AudioStream) (__result float32) {
	__result = (float32)(__SDL_GetAudioStreamFrequencyRatio(uintptr(unsafe.Pointer(stream))))
	return
}

func __gowrap__SDL_SetAudioStreamFrequencyRatio(stream AudioStream, ratio float32) (__result bool) {
	__result = (bool)(__SDL_SetAudioStreamFrequencyRatio(uintptr(unsafe.Pointer(stream)), *(*float32)(unsafe.Pointer(&ratio))) != 0)
	return
}

func __gowrap__SDL_GetAudioStreamGain(stream AudioStream) (__result float32) {
	__result = (float32)(__SDL_GetAudioStreamGain(uintptr(unsafe.Pointer(stream))))
	return
}

func __gowrap__SDL_SetAudioStreamGain(stream AudioStream, gain float32) (__result bool) {
	__result = (bool)(__SDL_SetAudioStreamGain(uintptr(unsafe.Pointer(stream)), *(*float32)(unsafe.Pointer(&gain))) != 0)
	return
}

func __gowrap__SDL_GetAudioStreamInputChannelMap(stream AudioStream, count *int32) (__result *int32) {
	__result = (*int32)(unsafe.Pointer(__SDL_GetAudioStreamInputChannelMap(uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_GetAudioStreamOutputChannelMap(stream AudioStream, count *int32) (__result *int32) {
	__result = (*int32)(unsafe.Pointer(__SDL_GetAudioStreamOutputChannelMap(uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(count)))))
	return
}

func __gowrap__SDL_SetAudioStreamInputChannelMap(stream AudioStream, chmap *int32, count int32) (__result bool) {
	__result = (bool)(__SDL_SetAudioStreamInputChannelMap(uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(chmap)), *(*int32)(unsafe.Pointer(&count))) != 0)
	return
}

func __gowrap__SDL_SetAudioStreamOutputChannelMap(stream AudioStream, chmap *int32, count int32) (__result bool) {
	__result = (bool)(__SDL_SetAudioStreamOutputChannelMap(uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(chmap)), *(*int32)(unsafe.Pointer(&count))) != 0)
	return
}

func __gowrap__SDL_PutAudioStreamData(stream AudioStream, buf []byte, len int32) (__result bool) {
	__result = (bool)(__SDL_PutAudioStreamData(uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&buf[0])), *(*int32)(unsafe.Pointer(&len))) != 0)
	runtime.KeepAlive(buf)
	return
}

func __gowrap__SDL_PutAudioStreamPlanarData(stream AudioStream, channel_buffers []*byte, num_channels int32, num_samples int32) (__result bool) {
	__result = (bool)(__SDL_PutAudioStreamPlanarData(uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&channel_buffers[0])), *(*int32)(unsafe.Pointer(&num_channels)), *(*int32)(unsafe.Pointer(&num_samples))) != 0)
	runtime.KeepAlive(channel_buffers)
	return
}

func __gowrap__SDL_GetAudioStreamData(stream AudioStream, buf []byte, len int32) (__result int32) {
	__result = (int32)(__SDL_GetAudioStreamData(uintptr(unsafe.Pointer(stream)), uintptr(unsafe.Pointer(&buf[0])), *(*int32)(unsafe.Pointer(&len))))
	runtime.KeepAlive(buf)
	return
}

func __gowrap__SDL_GetAudioStreamAvailable(stream AudioStream) (__result int32) {
	__result = (int32)(__SDL_GetAudioStreamAvailable(uintptr(unsafe.Pointer(stream))))
	return
}

func __gowrap__SDL_GetAudioStreamQueued(stream AudioStream) (__result int32) {
	__result = (int32)(__SDL_GetAudioStreamQueued(uintptr(unsafe.Pointer(stream))))
	return
}

func __gowrap__SDL_FlushAudioStream(stream AudioStream) (__result bool) {
	__result = (bool)(__SDL_FlushAudioStream(uintptr(unsafe.Pointer(stream))) != 0)
	return
}

func __gowrap__SDL_ClearAudioStream(stream AudioStream) (__result bool) {
	__result = (bool)(__SDL_ClearAudioStream(uintptr(unsafe.Pointer(stream))) != 0)
	return
}

func __gowrap__SDL_PauseAudioStreamDevice(stream AudioStream) (__result bool) {
	__result = (bool)(__SDL_PauseAudioStreamDevice(uintptr(unsafe.Pointer(stream))) != 0)
	return
}

func __gowrap__SDL_ResumeAudioStreamDevice(stream AudioStream) (__result bool) {
	__result = (bool)(__SDL_ResumeAudioStreamDevice(uintptr(unsafe.Pointer(stream))) != 0)
	return
}

func __gowrap__SDL_AudioStreamDevicePaused(stream AudioStream) (__result bool) {
	__result = (bool)(__SDL_AudioStreamDevicePaused(uintptr(unsafe.Pointer(stream))) != 0)
	return
}

func __gowrap__SDL_LockAudioStream(stream AudioStream) (__result bool) {
	__result = (bool)(__SDL_LockAudioStream(uintptr(unsafe.Pointer(stream))) != 0)
	return
}

func __gowrap__SDL_UnlockAudioStream(stream AudioStream) (__result bool) {
	__result = (bool)(__SDL_UnlockAudioStream(uintptr(unsafe.Pointer(stream))) != 0)
	return
}

func __gowrap__SDL_SetAudioStreamGetCallback(stream AudioStream, callback AudioStreamCallback, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_SetAudioStreamGetCallback(uintptr(unsafe.Pointer(stream)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	return
}

func __gowrap__SDL_SetAudioStreamPutCallback(stream AudioStream, callback AudioStreamCallback, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_SetAudioStreamPutCallback(uintptr(unsafe.Pointer(stream)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	return
}

func __gowrap__SDL_DestroyAudioStream(stream AudioStream) {
	__SDL_DestroyAudioStream(uintptr(unsafe.Pointer(stream)))
}

func __gowrap__SDL_OpenAudioDeviceStream(devid AudioDeviceID, spec *AudioSpec, callback AudioStreamCallback, userdata uintptr) (__result AudioStream) {
	__result = (AudioStream)(unsafe.Pointer(__SDL_OpenAudioDeviceStream(*(*int32)(unsafe.Pointer(&devid)), uintptr(unsafe.Pointer(spec)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata)))))
	return
}

func __gowrap__SDL_SetAudioIterationCallbacks(devid AudioDeviceID, start AudioStreamCallback, end AudioStreamCallback, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_SetAudioIterationCallbacks(*(*int32)(unsafe.Pointer(&devid)), 0 /* TODO: callbacks */, 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	return
}

func __gowrap__SDL_SetAudioPostmixCallback(devid AudioDeviceID, callback AudioPostmixCallback, userdata uintptr) (__result bool) {
	__result = (bool)(__SDL_SetAudioPostmixCallback(*(*int32)(unsafe.Pointer(&devid)), 0 /* TODO: callbacks */, uintptr(unsafe.Pointer(userdata))) != 0)
	return
}

func __gowrap__SDL_LoadWAV_IO(src IOStream, closeio bool, spec *AudioSpec, audio_buf [][]uint8, audio_len []uint32) (__result bool) {
	__result = (bool)(__SDL_LoadWAV_IO(uintptr(unsafe.Pointer(src)), *(*int32)(unsafe.Pointer(&closeio)), uintptr(unsafe.Pointer(spec)), uintptr(unsafe.Pointer(&audio_buf[0])), uintptr(unsafe.Pointer(&audio_len[0]))) != 0)
	runtime.KeepAlive(audio_buf)
	runtime.KeepAlive(audio_len)
	return
}

func __gowrap__SDL_LoadWAV(path string, spec *AudioSpec, audio_buf [][]uint8, audio_len []uint32) (__result bool) {
	__result = (bool)(__SDL_LoadWAV(((*[2]uintptr)(unsafe.Pointer(&path)))[0], uintptr(unsafe.Pointer(spec)), uintptr(unsafe.Pointer(&audio_buf[0])), uintptr(unsafe.Pointer(&audio_len[0]))) != 0)
	runtime.KeepAlive(path)
	runtime.KeepAlive(audio_buf)
	runtime.KeepAlive(audio_len)
	return
}

func __gowrap__SDL_MixAudio(dst []uint8, src []uint8, format AudioFormat, len uint32, volume float32) (__result bool) {
	__result = (bool)(__SDL_MixAudio(uintptr(unsafe.Pointer(&dst[0])), uintptr(unsafe.Pointer(&src[0])), *(*int32)(unsafe.Pointer(&format)), *(*int32)(unsafe.Pointer(&len)), *(*float32)(unsafe.Pointer(&volume))) != 0)
	runtime.KeepAlive(dst)
	runtime.KeepAlive(src)
	return
}

func __gowrap__SDL_ConvertAudioSamples(src_spec *AudioSpec, src_data []uint8, src_len int32, dst_spec *AudioSpec, dst_data [][]uint8, dst_len []int32) (__result bool) {
	__result = (bool)(__SDL_ConvertAudioSamples(uintptr(unsafe.Pointer(src_spec)), uintptr(unsafe.Pointer(&src_data[0])), *(*int32)(unsafe.Pointer(&src_len)), uintptr(unsafe.Pointer(dst_spec)), uintptr(unsafe.Pointer(&dst_data[0])), uintptr(unsafe.Pointer(&dst_len[0]))) != 0)
	runtime.KeepAlive(src_data)
	runtime.KeepAlive(dst_data)
	runtime.KeepAlive(dst_len)
	return
}

func __gowrap__SDL_GetAudioFormatName(format AudioFormat) (__result string) {
	__result = (string)(string(newCSliceRaw[byte](__SDL_GetAudioFormatName(*(*int32)(unsafe.Pointer(&format)))).Collect()))
	return
}

func __gowrap__SDL_GetSilenceValueForFormat(format AudioFormat) (__result int32) {
	__result = (int32)(__SDL_GetSilenceValueForFormat(*(*int32)(unsafe.Pointer(&format))))
	return
}

 
func init() {
	GetNumAudioDrivers = __gowrap__SDL_GetNumAudioDrivers
	GetAudioDriver = __gowrap__SDL_GetAudioDriver
	GetCurrentAudioDriver = __gowrap__SDL_GetCurrentAudioDriver
	GetAudioPlaybackDevices = __gowrap__SDL_GetAudioPlaybackDevices
	GetAudioRecordingDevices = __gowrap__SDL_GetAudioRecordingDevices
	GetAudioDeviceName = __gowrap__SDL_GetAudioDeviceName
	GetAudioDeviceFormat = __gowrap__SDL_GetAudioDeviceFormat
	GetAudioDeviceChannelMap = __gowrap__SDL_GetAudioDeviceChannelMap
	OpenAudioDevice = __gowrap__SDL_OpenAudioDevice
	IsAudioDevicePhysical = __gowrap__SDL_IsAudioDevicePhysical
	IsAudioDevicePlayback = __gowrap__SDL_IsAudioDevicePlayback
	PauseAudioDevice = __gowrap__SDL_PauseAudioDevice
	ResumeAudioDevice = __gowrap__SDL_ResumeAudioDevice
	AudioDevicePaused = __gowrap__SDL_AudioDevicePaused
	GetAudioDeviceGain = __gowrap__SDL_GetAudioDeviceGain
	SetAudioDeviceGain = __gowrap__SDL_SetAudioDeviceGain
	CloseAudioDevice = __gowrap__SDL_CloseAudioDevice
	BindAudioStreams = __gowrap__SDL_BindAudioStreams
	BindAudioStream = __gowrap__SDL_BindAudioStream
	UnbindAudioStreams = __gowrap__SDL_UnbindAudioStreams
	UnbindAudioStream = __gowrap__SDL_UnbindAudioStream
	GetAudioStreamDevice = __gowrap__SDL_GetAudioStreamDevice
	CreateAudioStream = __gowrap__SDL_CreateAudioStream
	GetAudioStreamProperties = __gowrap__SDL_GetAudioStreamProperties
	GetAudioStreamFormat = __gowrap__SDL_GetAudioStreamFormat
	SetAudioStreamFormat = __gowrap__SDL_SetAudioStreamFormat
	GetAudioStreamFrequencyRatio = __gowrap__SDL_GetAudioStreamFrequencyRatio
	SetAudioStreamFrequencyRatio = __gowrap__SDL_SetAudioStreamFrequencyRatio
	GetAudioStreamGain = __gowrap__SDL_GetAudioStreamGain
	SetAudioStreamGain = __gowrap__SDL_SetAudioStreamGain
	GetAudioStreamInputChannelMap = __gowrap__SDL_GetAudioStreamInputChannelMap
	GetAudioStreamOutputChannelMap = __gowrap__SDL_GetAudioStreamOutputChannelMap
	SetAudioStreamInputChannelMap = __gowrap__SDL_SetAudioStreamInputChannelMap
	SetAudioStreamOutputChannelMap = __gowrap__SDL_SetAudioStreamOutputChannelMap
	PutAudioStreamData = __gowrap__SDL_PutAudioStreamData
	PutAudioStreamPlanarData = __gowrap__SDL_PutAudioStreamPlanarData
	GetAudioStreamData = __gowrap__SDL_GetAudioStreamData
	GetAudioStreamAvailable = __gowrap__SDL_GetAudioStreamAvailable
	GetAudioStreamQueued = __gowrap__SDL_GetAudioStreamQueued
	FlushAudioStream = __gowrap__SDL_FlushAudioStream
	ClearAudioStream = __gowrap__SDL_ClearAudioStream
	PauseAudioStreamDevice = __gowrap__SDL_PauseAudioStreamDevice
	ResumeAudioStreamDevice = __gowrap__SDL_ResumeAudioStreamDevice
	AudioStreamDevicePaused = __gowrap__SDL_AudioStreamDevicePaused
	LockAudioStream = __gowrap__SDL_LockAudioStream
	UnlockAudioStream = __gowrap__SDL_UnlockAudioStream
	SetAudioStreamGetCallback = __gowrap__SDL_SetAudioStreamGetCallback
	SetAudioStreamPutCallback = __gowrap__SDL_SetAudioStreamPutCallback
	DestroyAudioStream = __gowrap__SDL_DestroyAudioStream
	OpenAudioDeviceStream = __gowrap__SDL_OpenAudioDeviceStream
	SetAudioIterationCallbacks = __gowrap__SDL_SetAudioIterationCallbacks
	SetAudioPostmixCallback = __gowrap__SDL_SetAudioPostmixCallback
	LoadWAV_IO = __gowrap__SDL_LoadWAV_IO
	LoadWAV = __gowrap__SDL_LoadWAV
	MixAudio = __gowrap__SDL_MixAudio
	ConvertAudioSamples = __gowrap__SDL_ConvertAudioSamples
	GetAudioFormatName = __gowrap__SDL_GetAudioFormatName
	GetSilenceValueForFormat = __gowrap__SDL_GetSilenceValueForFormat
}
