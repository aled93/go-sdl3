# Naming
- Remove all `SDL_` prefixes because in go all package usage already must be prefixed with package's name
  - Ensure that after removing prefixes first letter must be in upper case to be public.

# Functions

- Add line comment `//go:sdl3extern` to generate loading code for this function. If new name differs not only in case then provide original name in parenthesis without `SDL_` prefix:

        //go:sdl3extern(qsort_r)
        func QSortR func(...)

## Example
C:

    extern SDL_DECLSPEC SDL_DisplayMode ** SDLCALL SDL_GetFullscreenDisplayModes(SDL_DisplayID displayID, int *count);

Go:

    //go:sdl3extern
    var GetFullscreenDisplayModes func(displayID DisplayID, count *int) **DisplayMode

# Enum

- Default enum underlying type in C is int
- Rename enum variants from ALL_UPPERCASE to CamelCase
  - Prefix enum variants with enum name initials and underscore to prevent ident collision

## Example
C:

    typedef enum SDL_SystemTheme
    {
        SDL_SYSTEM_THEME_UNKNOWN,   /**< Unknown system theme */
        SDL_SYSTEM_THEME_LIGHT,     /**< Light colored system theme */
        SDL_SYSTEM_THEME_DARK       /**< Dark colored system theme */
    } SDL_SystemTheme;

Go:

    type SystemTheme int

    const (
        ST_Unknown SystemTheme = iota /**< Unknown system theme */
        ST_Light                      /**< Light colored system theme */
        ST_Dark                       /**< Dark colored system theme */
    )