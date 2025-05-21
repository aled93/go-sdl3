package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &HIDInit, "SDL_hid_init" },
            { &HIDExit, "SDL_hid_exit" },
            { &HIDDeviceChangeCount, "SDL_hid_device_change_count" },
            { &HIDEnumerate, "SDL_hid_enumerate" },
            { &HIDFreeEnumeration, "SDL_hid_free_enumeration" },
            { &HIDOpen, "SDL_hid_open" },
            { &HIDOpenPath, "SDL_hid_open_path" },
            { &HIDWrite, "SDL_hid_write" },
            { &HIDReadTimeout, "SDL_hid_read_timeout" },
            { &HIDRead, "SDL_hid_read" },
            { &HIDSetNonblocking, "SDL_hid_set_nonblocking" },
            { &HIDSendFeatureReport, "SDL_hid_send_feature_report" },
            { &HIDGetFeatureReport, "SDL_hid_get_feature_report" },
            { &HIDGetInputReport, "SDL_hid_get_input_report" },
            { &HIDclose, "SDL_hid_close" },
            { &HIDGetManufacturerString, "SDL_hid_get_manufacturer_string" },
            { &HIDGetProductString, "SDL_hid_get_product_string" },
            { &HIDGetSerialNumberString, "SDL_hid_get_serial_number_string" },
            { &HIDGetIndexedString, "SDL_hid_get_indexed_string" },
            { &HIDGetDeviceInfo, "SDL_hid_get_device_info" },
            { &HIDGetReportDescriptor, "SDL_hid_get_report_descriptor" },
            { &HIDBLESScan, "SDL_hid_ble_scan" },
        }
    })
}
