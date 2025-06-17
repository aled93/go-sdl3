//go:build wasm

package sdl

import (
  "runtime"
  "unsafe"
)



//go:wasmimport sdl3 SDL_hid_init
func __SDL_hid_init() int32

//go:wasmimport sdl3 SDL_hid_exit
func __SDL_hid_exit() int32

//go:wasmimport sdl3 SDL_hid_device_change_count
func __SDL_hid_device_change_count() int32

//go:wasmimport sdl3 SDL_hid_enumerate
func __SDL_hid_enumerate(int32, int32) uintptr

//go:wasmimport sdl3 SDL_hid_free_enumeration
func __SDL_hid_free_enumeration(uintptr)

//go:wasmimport sdl3 SDL_hid_open
func __SDL_hid_open(int32, int32, uintptr) uintptr

//go:wasmimport sdl3 SDL_hid_open_path
func __SDL_hid_open_path(uintptr) uintptr

//go:wasmimport sdl3 SDL_hid_write
func __SDL_hid_write(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_hid_read_timeout
func __SDL_hid_read_timeout(uintptr, uintptr, int32, int32) int32

//go:wasmimport sdl3 SDL_hid_read
func __SDL_hid_read(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_hid_set_nonblocking
func __SDL_hid_set_nonblocking(uintptr, int32) int32

//go:wasmimport sdl3 SDL_hid_send_feature_report
func __SDL_hid_send_feature_report(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_hid_get_feature_report
func __SDL_hid_get_feature_report(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_hid_get_input_report
func __SDL_hid_get_input_report(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_hid_close
func __SDL_hid_close(uintptr) int32

//go:wasmimport sdl3 SDL_hid_get_manufacturer_string
func __SDL_hid_get_manufacturer_string(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_hid_get_product_string
func __SDL_hid_get_product_string(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_hid_get_serial_number_string
func __SDL_hid_get_serial_number_string(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_hid_get_indexed_string
func __SDL_hid_get_indexed_string(uintptr, int32, uintptr, int32) int32

//go:wasmimport sdl3 SDL_hid_get_device_info
func __SDL_hid_get_device_info(uintptr) uintptr

//go:wasmimport sdl3 SDL_hid_get_report_descriptor
func __SDL_hid_get_report_descriptor(uintptr, uintptr, int32) int32

//go:wasmimport sdl3 SDL_hid_ble_scan
func __SDL_hid_ble_scan(int32)



func __gowrap__SDL_hid_init() (__result int32) {
	__result = (int32)(__SDL_hid_init())
	return
}

func __gowrap__SDL_hid_exit() (__result int32) {
	__result = (int32)(__SDL_hid_exit())
	return
}

func __gowrap__SDL_hid_device_change_count() (__result uint32) {
	__result = (uint32)(__SDL_hid_device_change_count())
	return
}

func __gowrap__SDL_hid_enumerate(vendor_id uint16, product_id uint16) (__result *HIDDeviceInfo) {
	__result = (*HIDDeviceInfo)(unsafe.Pointer(__SDL_hid_enumerate(*(*int32)(unsafe.Pointer(&vendor_id)), *(*int32)(unsafe.Pointer(&product_id)))))
	return
}

func __gowrap__SDL_hid_free_enumeration(devs *HIDDeviceInfo) {
	__SDL_hid_free_enumeration(uintptr(unsafe.Pointer(devs)))
}

func __gowrap__SDL_hid_open(vendor_id uint16, product_id uint16, serial_number string) (__result HIDDevice) {
	__result = (HIDDevice)(unsafe.Pointer(__SDL_hid_open(*(*int32)(unsafe.Pointer(&vendor_id)), *(*int32)(unsafe.Pointer(&product_id)), ((*[2]uintptr)(unsafe.Pointer(&serial_number)))[0])))
	runtime.KeepAlive(serial_number)
	return
}

func __gowrap__SDL_hid_open_path(path string) (__result HIDDevice) {
	__result = (HIDDevice)(unsafe.Pointer(__SDL_hid_open_path(((*[2]uintptr)(unsafe.Pointer(&path)))[0])))
	runtime.KeepAlive(path)
	return
}

func __gowrap__SDL_hid_write(dev HIDDevice, data *byte, length size_t) (__result int32) {
	__result = (int32)(__SDL_hid_write(uintptr(unsafe.Pointer(dev)), uintptr(unsafe.Pointer(data)), *(*int32)(unsafe.Pointer(&length))))
	return
}

func __gowrap__SDL_hid_read_timeout(dev HIDDevice, data *byte, length size_t, milliseconds int32) (__result int32) {
	__result = (int32)(__SDL_hid_read_timeout(uintptr(unsafe.Pointer(dev)), uintptr(unsafe.Pointer(data)), *(*int32)(unsafe.Pointer(&length)), *(*int32)(unsafe.Pointer(&milliseconds))))
	return
}

func __gowrap__SDL_hid_read(dev HIDDevice, data *byte, length size_t) (__result int32) {
	__result = (int32)(__SDL_hid_read(uintptr(unsafe.Pointer(dev)), uintptr(unsafe.Pointer(data)), *(*int32)(unsafe.Pointer(&length))))
	return
}

func __gowrap__SDL_hid_set_nonblocking(dev HIDDevice, nonblock int32) (__result int32) {
	__result = (int32)(__SDL_hid_set_nonblocking(uintptr(unsafe.Pointer(dev)), *(*int32)(unsafe.Pointer(&nonblock))))
	return
}

func __gowrap__SDL_hid_send_feature_report(dev HIDDevice, data *byte, length size_t) (__result int32) {
	__result = (int32)(__SDL_hid_send_feature_report(uintptr(unsafe.Pointer(dev)), uintptr(unsafe.Pointer(data)), *(*int32)(unsafe.Pointer(&length))))
	return
}

func __gowrap__SDL_hid_get_feature_report(dev HIDDevice, data *byte, length size_t) (__result int32) {
	__result = (int32)(__SDL_hid_get_feature_report(uintptr(unsafe.Pointer(dev)), uintptr(unsafe.Pointer(data)), *(*int32)(unsafe.Pointer(&length))))
	return
}

func __gowrap__SDL_hid_get_input_report(dev HIDDevice, data *byte, length size_t) (__result int32) {
	__result = (int32)(__SDL_hid_get_input_report(uintptr(unsafe.Pointer(dev)), uintptr(unsafe.Pointer(data)), *(*int32)(unsafe.Pointer(&length))))
	return
}

func __gowrap__SDL_hid_close(dev HIDDevice) (__result int32) {
	__result = (int32)(__SDL_hid_close(uintptr(unsafe.Pointer(dev))))
	return
}

func __gowrap__SDL_hid_get_manufacturer_string(dev HIDDevice, string_ string, maxlen size_t) (__result int32) {
	__result = (int32)(__SDL_hid_get_manufacturer_string(uintptr(unsafe.Pointer(dev)), ((*[2]uintptr)(unsafe.Pointer(&string_)))[0], *(*int32)(unsafe.Pointer(&maxlen))))
	runtime.KeepAlive(string_)
	return
}

func __gowrap__SDL_hid_get_product_string(dev HIDDevice, string_ string, maxlen size_t) (__result int32) {
	__result = (int32)(__SDL_hid_get_product_string(uintptr(unsafe.Pointer(dev)), ((*[2]uintptr)(unsafe.Pointer(&string_)))[0], *(*int32)(unsafe.Pointer(&maxlen))))
	runtime.KeepAlive(string_)
	return
}

func __gowrap__SDL_hid_get_serial_number_string(dev HIDDevice, string_ string, maxlen size_t) (__result int32) {
	__result = (int32)(__SDL_hid_get_serial_number_string(uintptr(unsafe.Pointer(dev)), ((*[2]uintptr)(unsafe.Pointer(&string_)))[0], *(*int32)(unsafe.Pointer(&maxlen))))
	runtime.KeepAlive(string_)
	return
}

func __gowrap__SDL_hid_get_indexed_string(dev HIDDevice, string_index int32, string_ string, maxlen size_t) (__result int32) {
	__result = (int32)(__SDL_hid_get_indexed_string(uintptr(unsafe.Pointer(dev)), *(*int32)(unsafe.Pointer(&string_index)), ((*[2]uintptr)(unsafe.Pointer(&string_)))[0], *(*int32)(unsafe.Pointer(&maxlen))))
	runtime.KeepAlive(string_)
	return
}

func __gowrap__SDL_hid_get_device_info(dev HIDDevice) (__result *HIDDeviceInfo) {
	__result = (*HIDDeviceInfo)(unsafe.Pointer(__SDL_hid_get_device_info(uintptr(unsafe.Pointer(dev)))))
	return
}

func __gowrap__SDL_hid_get_report_descriptor(dev HIDDevice, buf *byte, buf_size size_t) (__result int32) {
	__result = (int32)(__SDL_hid_get_report_descriptor(uintptr(unsafe.Pointer(dev)), uintptr(unsafe.Pointer(buf)), *(*int32)(unsafe.Pointer(&buf_size))))
	return
}

func __gowrap__SDL_hid_ble_scan(active bool) {
	__SDL_hid_ble_scan(*(*int32)(unsafe.Pointer(&active)))
}

 
func init() {
	HIDInit = __gowrap__SDL_hid_init
	HIDExit = __gowrap__SDL_hid_exit
	HIDDeviceChangeCount = __gowrap__SDL_hid_device_change_count
	HIDEnumerate = __gowrap__SDL_hid_enumerate
	HIDFreeEnumeration = __gowrap__SDL_hid_free_enumeration
	HIDOpen = __gowrap__SDL_hid_open
	HIDOpenPath = __gowrap__SDL_hid_open_path
	HIDWrite = __gowrap__SDL_hid_write
	HIDReadTimeout = __gowrap__SDL_hid_read_timeout
	HIDRead = __gowrap__SDL_hid_read
	HIDSetNonblocking = __gowrap__SDL_hid_set_nonblocking
	HIDSendFeatureReport = __gowrap__SDL_hid_send_feature_report
	HIDGetFeatureReport = __gowrap__SDL_hid_get_feature_report
	HIDGetInputReport = __gowrap__SDL_hid_get_input_report
	HIDclose = __gowrap__SDL_hid_close
	HIDGetManufacturerString = __gowrap__SDL_hid_get_manufacturer_string
	HIDGetProductString = __gowrap__SDL_hid_get_product_string
	HIDGetSerialNumberString = __gowrap__SDL_hid_get_serial_number_string
	HIDGetIndexedString = __gowrap__SDL_hid_get_indexed_string
	HIDGetDeviceInfo = __gowrap__SDL_hid_get_device_info
	HIDGetReportDescriptor = __gowrap__SDL_hid_get_report_descriptor
	HIDBLESScan = __gowrap__SDL_hid_ble_scan
}
