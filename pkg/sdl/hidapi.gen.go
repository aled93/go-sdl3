package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&HIDInit, lib, "SDL_hid_init")
        purego.RegisterLibFunc(&HIDExit, lib, "SDL_hid_exit")
        purego.RegisterLibFunc(&HIDDeviceChangeCount, lib, "SDL_hid_device_change_count")
        purego.RegisterLibFunc(&HIDEnumerate, lib, "SDL_hid_enumerate")
        purego.RegisterLibFunc(&HIDFreeEnumeration, lib, "SDL_hid_free_enumeration")
        purego.RegisterLibFunc(&HIDOpen, lib, "SDL_hid_open")
        purego.RegisterLibFunc(&HIDOpenPath, lib, "SDL_hid_open_path")
        purego.RegisterLibFunc(&HIDWrite, lib, "SDL_hid_write")
        purego.RegisterLibFunc(&HIDReadTimeout, lib, "SDL_hid_read_timeout")
        purego.RegisterLibFunc(&HIDRead, lib, "SDL_hid_read")
        purego.RegisterLibFunc(&HIDSetNonblocking, lib, "SDL_hid_set_nonblocking")
        purego.RegisterLibFunc(&HIDSendFeatureReport, lib, "SDL_hid_send_feature_report")
        purego.RegisterLibFunc(&HIDGetFeatureReport, lib, "SDL_hid_get_feature_report")
        purego.RegisterLibFunc(&HIDGetInputReport, lib, "SDL_hid_get_input_report")
        purego.RegisterLibFunc(&HIDclose, lib, "SDL_hid_close")
        purego.RegisterLibFunc(&HIDGetManufacturerString, lib, "SDL_hid_get_manufacturer_string")
        purego.RegisterLibFunc(&HIDGetProductString, lib, "SDL_hid_get_product_string")
        purego.RegisterLibFunc(&HIDGetSerialNumberString, lib, "SDL_hid_get_serial_number_string")
        purego.RegisterLibFunc(&HIDGetIndexedString, lib, "SDL_hid_get_indexed_string")
        purego.RegisterLibFunc(&HIDGetDeviceInfo, lib, "SDL_hid_get_device_info")
        purego.RegisterLibFunc(&HIDGetReportDescriptor, lib, "SDL_hid_get_report_descriptor")
        purego.RegisterLibFunc(&HIDBLESScan, lib, "SDL_hid_ble_scan")
        
    })
}
