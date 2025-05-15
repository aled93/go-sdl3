package sdl

import "github.com/ebitengine/purego"

func init() {
    registerLoaderFunc(func(lib uintptr) {
        purego.RegisterLibFunc(&HasRectIntersection, lib, "SDL_HasRectIntersection")
        purego.RegisterLibFunc(&GetRectIntersection, lib, "SDL_GetRectIntersection")
        purego.RegisterLibFunc(&GetRectUnion, lib, "SDL_GetRectUnion")
        purego.RegisterLibFunc(&GetRectEnclosingPoints, lib, "SDL_GetRectEnclosingPoints")
        purego.RegisterLibFunc(&GetRectAndLineIntersection, lib, "SDL_GetRectAndLineIntersection")
        purego.RegisterLibFunc(&HasRectIntersectionFloat, lib, "SDL_HasRectIntersectionFloat")
        purego.RegisterLibFunc(&GetRectIntersectionFloat, lib, "SDL_GetRectIntersectionFloat")
        purego.RegisterLibFunc(&GetRectUnionFloat, lib, "SDL_GetRectUnionFloat")
        purego.RegisterLibFunc(&GetRectEnclosingPointsFloat, lib, "SDL_GetRectEnclosingPointsFloat")
        purego.RegisterLibFunc(&GetRectAndLineIntersectionFloat, lib, "SDL_GetRectAndLineIntersectionFloat")
        
    })
}
