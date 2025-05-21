package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &GetDateTimeLocalePreferences, "SDL_GetDateTimeLocalePreferences" },
            { &GetCurrentTime, "SDL_GetCurrentTime" },
            { &TimeToDateTime, "SDL_TimeToDateTime" },
            { &DateTimeToTime, "SDL_DateTimeToTime" },
            { &TimeToWindows, "SDL_TimeToWindows" },
            { &TimeFromWindows, "SDL_TimeFromWindows" },
            { &GetDaysInMonth, "SDL_GetDaysInMonth" },
            { &GetDayOfYear, "SDL_GetDayOfYear" },
            { &GetDayOfWeek, "SDL_GetDayOfWeek" },
        }
    })
}
