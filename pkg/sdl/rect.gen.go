package sdl

func init() {
    registerLoaderFunc(func() []externFunc {
        return []externFunc {
            { &HasRectIntersection, "SDL_HasRectIntersection" },
            { &GetRectIntersection, "SDL_GetRectIntersection" },
            { &GetRectUnion, "SDL_GetRectUnion" },
            { &GetRectEnclosingPoints, "SDL_GetRectEnclosingPoints" },
            { &GetRectAndLineIntersection, "SDL_GetRectAndLineIntersection" },
            { &HasRectIntersectionFloat, "SDL_HasRectIntersectionFloat" },
            { &GetRectIntersectionFloat, "SDL_GetRectIntersectionFloat" },
            { &GetRectUnionFloat, "SDL_GetRectUnionFloat" },
            { &GetRectEnclosingPointsFloat, "SDL_GetRectEnclosingPointsFloat" },
            { &GetRectAndLineIntersectionFloat, "SDL_GetRectAndLineIntersectionFloat" },
        }
    })
}
