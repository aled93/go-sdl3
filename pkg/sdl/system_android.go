package sdl

/*
* Platform specific functions for Android
 */

/**
* Get the Android Java Native Interface Environment of the current thread.
*
* This is the JNIEnv one needs to access the Java virtual machine from native
* code, and is needed for many Android APIs to be usable from C.
*
* The prototype of the function in SDL's code actually declare a void* return
* type, even if the implementation returns a pointer to a JNIEnv. The
* rationale being that the SDL headers can avoid including jni.h.
*
* \returns a pointer to Java native interface object (JNIEnv) to which the
*          current thread is attached, or NULL on failure; call
*          SDL_GetError() for more information.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetAndroidActivity
 */
//go:sdl3extern
var GetAndroidJNIEnv func() uintptr

/**
* Retrieve the Java instance of the Android activity class.
*
* The prototype of the function in SDL's code actually declares a void*
* return type, even if the implementation returns a jobject. The rationale
* being that the SDL headers can avoid including jni.h.
*
* The jobject returned by the function is a local reference and must be
* released by the caller. See the PushLocalFrame() and PopLocalFrame() or
* DeleteLocalRef() functions of the Java native interface:
*
* https://docs.oracle.com/javase/1.5.0/docs/guide/jni/spec/functions.html
*
* \returns the jobject representing the instance of the Activity class of the
*          Android application, or NULL on failure; call SDL_GetError() for
*          more information.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetAndroidJNIEnv
 */
//go:sdl3extern
var GetAndroidActivity func() uintptr

/**
* Query Android API level of the current device.
*
* - API level 35: Android 15 (VANILLA_ICE_CREAM)
* - API level 34: Android 14 (UPSIDE_DOWN_CAKE)
* - API level 33: Android 13 (TIRAMISU)
* - API level 32: Android 12L (S_V2)
* - API level 31: Android 12 (S)
* - API level 30: Android 11 (R)
* - API level 29: Android 10 (Q)
* - API level 28: Android 9 (P)
* - API level 27: Android 8.1 (O_MR1)
* - API level 26: Android 8.0 (O)
* - API level 25: Android 7.1 (N_MR1)
* - API level 24: Android 7.0 (N)
* - API level 23: Android 6.0 (M)
* - API level 22: Android 5.1 (LOLLIPOP_MR1)
* - API level 21: Android 5.0 (LOLLIPOP, L)
* - API level 20: Android 4.4W (KITKAT_WATCH)
* - API level 19: Android 4.4 (KITKAT)
* - API level 18: Android 4.3 (JELLY_BEAN_MR2)
* - API level 17: Android 4.2 (JELLY_BEAN_MR1)
* - API level 16: Android 4.1 (JELLY_BEAN)
* - API level 15: Android 4.0.3 (ICE_CREAM_SANDWICH_MR1)
* - API level 14: Android 4.0 (ICE_CREAM_SANDWICH)
* - API level 13: Android 3.2 (HONEYCOMB_MR2)
* - API level 12: Android 3.1 (HONEYCOMB_MR1)
* - API level 11: Android 3.0 (HONEYCOMB)
* - API level 10: Android 2.3.3 (GINGERBREAD_MR1)
*
* \returns the Android API level.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetAndroidSDKVersion func() int

/**
* Query if the application is running on a Chromebook.
*
* \returns true if this is a Chromebook, false otherwise.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var IsChromebook func() bool

/**
* Query if the application is running on a Samsung DeX docking station.
*
* \returns true if this is a DeX docking station, false otherwise.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var IsDeXMode func() bool

/**
* Trigger the Android system back button behavior.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SendAndroidBackButton func()

/**
* See the official Android developer guide for more information:
* http://developer.android.com/guide/topics/data/data-storage.html
*
* \since This macro is available since SDL 3.2.0.
 */
const ANDROID_EXTERNAL_STORAGE_READ = 0x01

/**
* See the official Android developer guide for more information:
* http://developer.android.com/guide/topics/data/data-storage.html
*
* \since This macro is available since SDL 3.2.0.
 */
const SDL_ANDROID_EXTERNAL_STORAGE_WRITE = 0x02

/**
* Get the path used for internal storage for this Android application.
*
* This path is unique to your application and cannot be written to by other
* applications.
*
* Your internal storage path is typically:
* `/data/data/your.app.package/files`.
*
* This is a C wrapper over `android.content.Context.getFilesDir()`:
*
* https://developer.android.com/reference/android/content/Context#getFilesDir()
*
* \returns the path used for internal storage or NULL on failure; call
*          SDL_GetError() for more information.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetAndroidExternalStoragePath
* \sa SDL_GetAndroidCachePath
 */
//go:sdl3extern
var GetAndroidInternalStoragePath func() *byte

/**
* Get the current state of external storage for this Android application.
*
* The current state of external storage, a bitmask of these values:
* `SDL_ANDROID_EXTERNAL_STORAGE_READ`, `SDL_ANDROID_EXTERNAL_STORAGE_WRITE`.
*
* If external storage is currently unavailable, this will return 0.
*
* \returns the current state of external storage, or 0 if external storage is
*          currently unavailable.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetAndroidExternalStoragePath
 */
//go:sdl3extern
var GetAndroidExternalStorageState func() uint32

/**
* Get the path used for external storage for this Android application.
*
* This path is unique to your application, but is public and can be written
* to by other applications.
*
* Your external storage path is typically:
* `/storage/sdcard0/Android/data/your.app.package/files`.
*
* This is a C wrapper over `android.content.Context.getExternalFilesDir()`:
*
* https://developer.android.com/reference/android/content/Context#getExternalFilesDir()
*
* \returns the path used for external storage for this application on success
*          or NULL on failure; call SDL_GetError() for more information.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetAndroidExternalStorageState
* \sa SDL_GetAndroidInternalStoragePath
* \sa SDL_GetAndroidCachePath
 */
//go:sdl3extern
var GetAndroidExternalStoragePath func() *byte

/**
* Get the path used for caching data for this Android application.
*
* This path is unique to your application, but is public and can be written
* to by other applications.
*
* Your cache path is typically: `/data/data/your.app.package/cache/`.
*
* This is a C wrapper over `android.content.Context.getCacheDir()`:
*
* https://developer.android.com/reference/android/content/Context#getCacheDir()
*
* \returns the path used for caches for this application on success or NULL
*          on failure; call SDL_GetError() for more information.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetAndroidInternalStoragePath
* \sa SDL_GetAndroidExternalStoragePath
 */
//go:sdl3extern
var GetAndroidCachePath func() *byte

/**
* Callback that presents a response from a SDL_RequestAndroidPermission call.
*
* \param userdata an app-controlled pointer that is passed to the callback.
* \param permission the Android-specific permission name that was requested.
* \param granted true if permission is granted, false if denied.
*
* \since This datatype is available since SDL 3.2.0.
*
* \sa SDL_RequestAndroidPermission
 */
type RequestAndroidPermissionCallback func(userdata uintptr, permission *byte, granted bool)

/**
* Request permissions at runtime, asynchronously.
*
* You do not need to call this for built-in functionality of SDL; recording
* from a microphone or reading images from a camera, using standard SDL APIs,
* will manage permission requests for you.
*
* This function never blocks. Instead, the app-supplied callback will be
* called when a decision has been made. This callback may happen on a
* different thread, and possibly much later, as it might wait on a user to
* respond to a system dialog. If permission has already been granted for a
* specific entitlement, the callback will still fire, probably on the current
* thread and before this function returns.
*
* If the request submission fails, this function returns -1 and the callback
* will NOT be called, but this should only happen in catastrophic conditions,
* like memory running out. Normally there will be a yes or no to the request
* through the callback.
*
* For the `permission` parameter, choose a value from here:
*
* https://developer.android.com/reference/android/Manifest.permission
*
* \param permission the permission to request.
* \param cb the callback to trigger when the request has a response.
* \param userdata an app-controlled pointer that is passed to the callback.
* \returns true if the request was submitted, false if there was an error
*          submitting. The result of the request is only ever reported
*          through the callback, not this return value.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var RequestAndroidPermission func(permission *byte, cb RequestAndroidPermissionCallback, userdata uintptr) bool

/**
* Shows an Android toast notification.
*
* Toasts are a sort of lightweight notification that are unique to Android.
*
* https://developer.android.com/guide/topics/ui/notifiers/toasts
*
* Shows toast in UI thread.
*
* For the `gravity` parameter, choose a value from here, or -1 if you don't
* have a preference:
*
* https://developer.android.com/reference/android/view/Gravity
*
* \param message text message to be shown.
* \param duration 0=short, 1=long.
* \param gravity where the notification should appear on the screen.
* \param xoffset set this parameter only when gravity >=0.
* \param yoffset set this parameter only when gravity >=0.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var ShowAndroidToast func(message *byte, duration, gravity, xoffset, yoffset int) bool

/**
* Send a user command to SDLActivity.
*
* Override "boolean onUnhandledMessage(Message msg)" to handle the message.
*
* \param command user command that must be greater or equal to 0x8000.
* \param param user parameter.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var SendAndroidMessage func(command uint32, param int) bool
