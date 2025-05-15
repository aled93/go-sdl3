package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&GetDateTimeLocalePreferences, lib, "SDL_GetDateTimeLocalePreferences")
        purego.RegisterLibFunc(&GetCurrentTime, lib, "SDL_GetCurrentTime")
        purego.RegisterLibFunc(&TimeToDateTime, lib, "SDL_TimeToDateTime")
        purego.RegisterLibFunc(&DateTimeToTime, lib, "SDL_DateTimeToTime")
        purego.RegisterLibFunc(&TimeToWindows, lib, "SDL_TimeToWindows")
        purego.RegisterLibFunc(&TimeFromWindows, lib, "SDL_TimeFromWindows")
        purego.RegisterLibFunc(&GetDaysInMonth, lib, "SDL_GetDaysInMonth")
        purego.RegisterLibFunc(&GetDayOfYear, lib, "SDL_GetDayOfYear")
        purego.RegisterLibFunc(&GetDayOfWeek, lib, "SDL_GetDayOfWeek")
        
    })
}
