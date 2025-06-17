//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_IOFromFile
func __SDL_IOFromFile(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_IOFromMem
func __SDL_IOFromMem(uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_IOFromConstMem
func __SDL_IOFromConstMem(uintptr, int32) uintptr

//go:wasmimport sdl3 SDL_IOFromDynamicMem
func __SDL_IOFromDynamicMem() uintptr

//go:wasmimport sdl3 SDL_OpenIO
func __SDL_OpenIO(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_CloseIO
func __SDL_CloseIO(uintptr) int32

//go:wasmimport sdl3 SDL_GetIOProperties
func __SDL_GetIOProperties(uintptr) int32

//go:wasmimport sdl3 SDL_GetIOStatus
func __SDL_GetIOStatus(uintptr) int32

//go:wasmimport sdl3 SDL_GetIOSize
func __SDL_GetIOSize(uintptr) int64

//go:wasmimport sdl3 SDL_SeekIO
func __SDL_SeekIO(uintptr, int64, int32) int64

//go:wasmimport sdl3 SDL_TellIO
func __SDL_TellIO(uintptr) int64

//go:wasmimport sdl3 SDL_ReadIO
func __SDL_ReadIO(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_WriteIO
func __SDL_WriteIO(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_FlushIO
func __SDL_FlushIO(uintptr) int32

//go:wasmimport sdl3 SDL_LoadFile_IO
func __SDL_LoadFile_IO(uintptr, int32, int32) uintptr

//go:wasmimport sdl3 SDL_LoadFile
func __SDL_LoadFile(uintptr, uintptr) uintptr

//go:wasmimport sdl3 SDL_SaveFile_IO
func __SDL_SaveFile_IO(uintptr, uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_SaveFile
func __SDL_SaveFile(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_ReadU8
func __SDL_ReadU8(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadS8
func __SDL_ReadS8(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadU16LE
func __SDL_ReadU16LE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadS16LE
func __SDL_ReadS16LE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadU16BE
func __SDL_ReadU16BE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadS16BE
func __SDL_ReadS16BE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadU32LE
func __SDL_ReadU32LE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadS32LE
func __SDL_ReadS32LE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadU32BE
func __SDL_ReadU32BE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadS32BE
func __SDL_ReadS32BE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadU64LE
func __SDL_ReadU64LE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadS64LE
func __SDL_ReadS64LE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadU64BE
func __SDL_ReadU64BE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_ReadS64BE
func __SDL_ReadS64BE(uintptr, uintptr) int32

//go:wasmimport sdl3 SDL_WriteU8
func __SDL_WriteU8(uintptr, int32) int32

//go:wasmimport sdl3 SDL_WriteS8
func __SDL_WriteS8(uintptr, int32) int32

//go:wasmimport sdl3 SDL_WriteU16LE
func __SDL_WriteU16LE(uintptr, int32) int32

//go:wasmimport sdl3 SDL_WriteS16LE
func __SDL_WriteS16LE(uintptr, int32) int32

//go:wasmimport sdl3 SDL_WriteU16BE
func __SDL_WriteU16BE(uintptr, int32) int32

//go:wasmimport sdl3 SDL_WriteS16BE
func __SDL_WriteS16BE(uintptr, int32) int32

//go:wasmimport sdl3 SDL_WriteU32LE
func __SDL_WriteU32LE(uintptr, int32) int32

//go:wasmimport sdl3 SDL_WriteS32LE
func __SDL_WriteS32LE(uintptr, int32) int32

//go:wasmimport sdl3 SDL_WriteU32BE
func __SDL_WriteU32BE(uintptr, int32) int32

//go:wasmimport sdl3 SDL_WriteS32BE
func __SDL_WriteS32BE(uintptr, int32) int32

//go:wasmimport sdl3 SDL_WriteU64LE
func __SDL_WriteU64LE(uintptr, int64) int32

//go:wasmimport sdl3 SDL_WriteS64LE
func __SDL_WriteS64LE(uintptr, int64) int32

//go:wasmimport sdl3 SDL_WriteU64BE
func __SDL_WriteU64BE(uintptr, int64) int32

//go:wasmimport sdl3 SDL_WriteS64BE
func __SDL_WriteS64BE(uintptr, int64) int32



func __gowrap__SDL_IOFromFile(file string, mode string) (__result IOStream) {
	__result = (IOStream)(unsafe.Pointer(__SDL_IOFromFile(((*[2]uintptr)(unsafe.Pointer(&file)))[0], ((*[2]uintptr)(unsafe.Pointer(&mode)))[0])))
	runtime.KeepAlive(file)
	runtime.KeepAlive(mode)
	return
}

func __gowrap__SDL_IOFromMem(mem []byte, size size_t) (__result IOStream) {
	__result = (IOStream)(unsafe.Pointer(__SDL_IOFromMem(uintptr(unsafe.Pointer(&mem[0])), *(*int32)(unsafe.Pointer(&size)))))
	runtime.KeepAlive(mem)
	return
}

func __gowrap__SDL_IOFromConstMem(mem []byte, size size_t) (__result IOStream) {
	__result = (IOStream)(unsafe.Pointer(__SDL_IOFromConstMem(uintptr(unsafe.Pointer(&mem[0])), *(*int32)(unsafe.Pointer(&size)))))
	runtime.KeepAlive(mem)
	return
}

func __gowrap__SDL_IOFromDynamicMem() (__result IOStream) {
	__result = (IOStream)(unsafe.Pointer(__SDL_IOFromDynamicMem()))
	return
}

func __gowrap__SDL_OpenIO(iface *IOStreamInterface, userdata uintptr) (__result IOStream) {
	__result = (IOStream)(unsafe.Pointer(__SDL_OpenIO(uintptr(unsafe.Pointer(iface)), uintptr(unsafe.Pointer(userdata)))))
	return
}

func __gowrap__SDL_CloseIO(context IOStream) (__result bool) {
	__result = (bool)(__SDL_CloseIO(uintptr(unsafe.Pointer(context))) != 0)
	return
}

func __gowrap__SDL_GetIOProperties(context IOStream) (__result PropertiesID) {
	__result = (PropertiesID)(__SDL_GetIOProperties(uintptr(unsafe.Pointer(context))))
	return
}

func __gowrap__SDL_GetIOStatus(context IOStream) (__result IOStatus) {
	__result = (IOStatus)(__SDL_GetIOStatus(uintptr(unsafe.Pointer(context))))
	return
}

func __gowrap__SDL_GetIOSize(context IOStream) (__result int64) {
	__result = (int64)(__SDL_GetIOSize(uintptr(unsafe.Pointer(context))))
	return
}

func __gowrap__SDL_SeekIO(context IOStream, offset int64, whence IOWhence) (__result int64) {
	__result = (int64)(__SDL_SeekIO(uintptr(unsafe.Pointer(context)), *(*int64)(unsafe.Pointer(&offset)), *(*int32)(unsafe.Pointer(&whence))))
	return
}

func __gowrap__SDL_TellIO(context IOStream) (__result int64) {
	__result = (int64)(__SDL_TellIO(uintptr(unsafe.Pointer(context))))
	return
}

func __gowrap__SDL_ReadIO(context IOStream, ptr []byte, size size_t) (__result size_t) {
	__result = (size_t)(__SDL_ReadIO(uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(&ptr[0])), *(*int32)(unsafe.Pointer(&size))))
	runtime.KeepAlive(ptr)
	return
}

func __gowrap__SDL_WriteIO(context IOStream, ptr []byte, size size_t) (__result size_t) {
	__result = (size_t)(__SDL_WriteIO(uintptr(unsafe.Pointer(context)), uintptr(unsafe.Pointer(&ptr[0])), *(*int32)(unsafe.Pointer(&size))))
	runtime.KeepAlive(ptr)
	return
}

func __gowrap__SDL_FlushIO(context IOStream) (__result bool) {
	__result = (bool)(__SDL_FlushIO(uintptr(unsafe.Pointer(context))) != 0)
	return
}

func __gowrap__SDL_LoadFile_IO(src IOStream, datasize size_t, closeio bool) (__result uintptr) {
	__result = (uintptr)(unsafe.Pointer(__SDL_LoadFile_IO(uintptr(unsafe.Pointer(src)), *(*int32)(unsafe.Pointer(&datasize)), *(*int32)(unsafe.Pointer(&closeio)))))
	return
}

func __gowrap__SDL_LoadFile(file string, datasize *size_t) (__result []byte) {
	__result = ([]byte)(newCSliceRaw[byte](__SDL_LoadFile(((*[2]uintptr)(unsafe.Pointer(&file)))[0], uintptr(unsafe.Pointer(datasize)))).Collect())
	runtime.KeepAlive(file)
	return
}

func __gowrap__SDL_SaveFile_IO(src IOStream, data []byte, datasize size_t, closeio bool) (__result bool) {
	__result = (bool)(__SDL_SaveFile_IO(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(&data[0])), *(*int32)(unsafe.Pointer(&datasize)), *(*int32)(unsafe.Pointer(&closeio))) != 0)
	runtime.KeepAlive(data)
	return
}

func __gowrap__SDL_SaveFile(file string, data []byte, datasize size_t) (__result bool) {
	__result = (bool)(__SDL_SaveFile(((*[2]uintptr)(unsafe.Pointer(&file)))[0], uintptr(unsafe.Pointer(&data[0])), *(*int32)(unsafe.Pointer(&datasize))) != 0)
	runtime.KeepAlive(file)
	runtime.KeepAlive(data)
	return
}

func __gowrap__SDL_ReadU8(src IOStream, value *uint8) (__result bool) {
	__result = (bool)(__SDL_ReadU8(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadS8(src IOStream, value *int8) (__result bool) {
	__result = (bool)(__SDL_ReadS8(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadU16LE(src IOStream, value *uint16) (__result bool) {
	__result = (bool)(__SDL_ReadU16LE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadS16LE(src IOStream, value *int16) (__result bool) {
	__result = (bool)(__SDL_ReadS16LE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadU16BE(src IOStream, value *uint16) (__result bool) {
	__result = (bool)(__SDL_ReadU16BE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadS16BE(src IOStream, value *int16) (__result bool) {
	__result = (bool)(__SDL_ReadS16BE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadU32LE(src IOStream, value *uint32) (__result bool) {
	__result = (bool)(__SDL_ReadU32LE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadS32LE(src IOStream, value *int32) (__result bool) {
	__result = (bool)(__SDL_ReadS32LE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadU32BE(src IOStream, value *uint32) (__result bool) {
	__result = (bool)(__SDL_ReadU32BE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadS32BE(src IOStream, value *int32) (__result bool) {
	__result = (bool)(__SDL_ReadS32BE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadU64LE(src IOStream, value *uint64) (__result bool) {
	__result = (bool)(__SDL_ReadU64LE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadS64LE(src IOStream, value *int64) (__result bool) {
	__result = (bool)(__SDL_ReadS64LE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadU64BE(src IOStream, value *uint64) (__result bool) {
	__result = (bool)(__SDL_ReadU64BE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_ReadS64BE(src IOStream, value *int64) (__result bool) {
	__result = (bool)(__SDL_ReadS64BE(uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(value))) != 0)
	return
}

func __gowrap__SDL_WriteU8(dst IOStream, value uint8) (__result bool) {
	__result = (bool)(__SDL_WriteU8(uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteS8(dst IOStream, value int8) (__result bool) {
	__result = (bool)(__SDL_WriteS8(uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteU16LE(dst IOStream, value uint16) (__result bool) {
	__result = (bool)(__SDL_WriteU16LE(uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteS16LE(dst IOStream, value int16) (__result bool) {
	__result = (bool)(__SDL_WriteS16LE(uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteU16BE(dst IOStream, value uint16) (__result bool) {
	__result = (bool)(__SDL_WriteU16BE(uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteS16BE(dst IOStream, value int16) (__result bool) {
	__result = (bool)(__SDL_WriteS16BE(uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteU32LE(dst IOStream, value uint32) (__result bool) {
	__result = (bool)(__SDL_WriteU32LE(uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteS32LE(dst IOStream, value int32) (__result bool) {
	__result = (bool)(__SDL_WriteS32LE(uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteU32BE(dst IOStream, value uint32) (__result bool) {
	__result = (bool)(__SDL_WriteU32BE(uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteS32BE(dst IOStream, value int32) (__result bool) {
	__result = (bool)(__SDL_WriteS32BE(uintptr(unsafe.Pointer(dst)), *(*int32)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteU64LE(dst IOStream, value uint64) (__result bool) {
	__result = (bool)(__SDL_WriteU64LE(uintptr(unsafe.Pointer(dst)), *(*int64)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteS64LE(dst IOStream, value int64) (__result bool) {
	__result = (bool)(__SDL_WriteS64LE(uintptr(unsafe.Pointer(dst)), *(*int64)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteU64BE(dst IOStream, value uint64) (__result bool) {
	__result = (bool)(__SDL_WriteU64BE(uintptr(unsafe.Pointer(dst)), *(*int64)(unsafe.Pointer(&value))) != 0)
	return
}

func __gowrap__SDL_WriteS64BE(dst IOStream, value int64) (__result bool) {
	__result = (bool)(__SDL_WriteS64BE(uintptr(unsafe.Pointer(dst)), *(*int64)(unsafe.Pointer(&value))) != 0)
	return
}

 
func init() {
	IOFromFile = __gowrap__SDL_IOFromFile
	IOFromMem = __gowrap__SDL_IOFromMem
	IOFromConstMem = __gowrap__SDL_IOFromConstMem
	IOFromDynamicMem = __gowrap__SDL_IOFromDynamicMem
	OpenIO = __gowrap__SDL_OpenIO
	CloseIO = __gowrap__SDL_CloseIO
	GetIOProperties = __gowrap__SDL_GetIOProperties
	GetIOStatus = __gowrap__SDL_GetIOStatus
	GetIOSize = __gowrap__SDL_GetIOSize
	SeekIO = __gowrap__SDL_SeekIO
	TellIO = __gowrap__SDL_TellIO
	ReadIO = __gowrap__SDL_ReadIO
	WriteIO = __gowrap__SDL_WriteIO
	FlushIO = __gowrap__SDL_FlushIO
	LoadFile_IO = __gowrap__SDL_LoadFile_IO
	LoadFile = __gowrap__SDL_LoadFile
	SaveFile_IO = __gowrap__SDL_SaveFile_IO
	SaveFile = __gowrap__SDL_SaveFile
	ReadU8 = __gowrap__SDL_ReadU8
	ReadS8 = __gowrap__SDL_ReadS8
	ReadU16LE = __gowrap__SDL_ReadU16LE
	ReadS16LE = __gowrap__SDL_ReadS16LE
	ReadU16BE = __gowrap__SDL_ReadU16BE
	ReadS16BE = __gowrap__SDL_ReadS16BE
	ReadU32LE = __gowrap__SDL_ReadU32LE
	ReadS32LE = __gowrap__SDL_ReadS32LE
	ReadU32BE = __gowrap__SDL_ReadU32BE
	ReadS32BE = __gowrap__SDL_ReadS32BE
	ReadU64LE = __gowrap__SDL_ReadU64LE
	ReadS64LE = __gowrap__SDL_ReadS64LE
	ReadU64BE = __gowrap__SDL_ReadU64BE
	ReadS64BE = __gowrap__SDL_ReadS64BE
	WriteU8 = __gowrap__SDL_WriteU8
	WriteS8 = __gowrap__SDL_WriteS8
	WriteU16LE = __gowrap__SDL_WriteU16LE
	WriteS16LE = __gowrap__SDL_WriteS16LE
	WriteU16BE = __gowrap__SDL_WriteU16BE
	WriteS16BE = __gowrap__SDL_WriteS16BE
	WriteU32LE = __gowrap__SDL_WriteU32LE
	WriteS32LE = __gowrap__SDL_WriteS32LE
	WriteU32BE = __gowrap__SDL_WriteU32BE
	WriteS32BE = __gowrap__SDL_WriteS32BE
	WriteU64LE = __gowrap__SDL_WriteU64LE
	WriteS64LE = __gowrap__SDL_WriteS64LE
	WriteU64BE = __gowrap__SDL_WriteU64BE
	WriteS64BE = __gowrap__SDL_WriteS64BE
}
