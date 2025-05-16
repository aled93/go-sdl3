/*
  Simple DirectMedia Layer
  Copyright (C) 1997-2025 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/

/*
  Language bindings and additional adaptations by AleD, 2025
  This file contains modified code and/or translated interfaces.
  Original SDL3 source: https://github.com/libsdl-org/SDL
*/

/**
 * # CategoryRect
 *
 * Some helper functions for managing rectangles and 2D points, in both
 * integer and floating point versions.
 */
package sdl

import "math"

/**
* The structure that defines a point (using integers).
*
* \since This struct is available since SDL 3.2.0.
*
* \sa SDL_GetRectEnclosingPoints
* \sa SDL_PointInRect
 */
type Point struct {
	X, Y int
}

/**
* The structure that defines a point (using floating point values).
*
* \since This struct is available since SDL 3.2.0.
*
* \sa SDL_GetRectEnclosingPointsFloat
* \sa SDL_PointInRectFloat
 */
type FPoint struct {
	X, Y float32
}

/**
* A rectangle, with the origin at the upper left (using integers).
*
* \since This struct is available since SDL 3.2.0.
*
* \sa SDL_RectEmpty
* \sa SDL_RectsEqual
* \sa SDL_HasRectIntersection
* \sa SDL_GetRectIntersection
* \sa SDL_GetRectAndLineIntersection
* \sa SDL_GetRectUnion
* \sa SDL_GetRectEnclosingPoints
 */
type Rect struct {
	X, Y int
	W, H int
}

/**
* A rectangle, with the origin at the upper left (using floating point
* values).
*
* \since This struct is available since SDL 3.2.0.
*
* \sa SDL_RectEmptyFloat
* \sa SDL_RectsEqualFloat
* \sa SDL_RectsEqualEpsilon
* \sa SDL_HasRectIntersectionFloat
* \sa SDL_GetRectIntersectionFloat
* \sa SDL_GetRectAndLineIntersectionFloat
* \sa SDL_GetRectUnionFloat
* \sa SDL_GetRectEnclosingPointsFloat
* \sa SDL_PointInRectFloat
 */
type FRect struct {
	X float32
	Y float32
	W float32
	H float32
}

/**
* Convert an SDL_Rect to SDL_FRect
*
* \param rect a pointer to an SDL_Rect.
* \param frect a pointer filled in with the floating point representation of
*              `rect`.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
func RectToFRect(rect *Rect, frect *FRect) {
	frect.X = float32(rect.X)
	frect.Y = float32(rect.Y)
	frect.W = float32(rect.W)
	frect.H = float32(rect.H)
}

/**
* Determine whether a point resides inside a rectangle.
*
* A point is considered part of a rectangle if both `p` and `r` are not NULL,
* and `p`'s x and y coordinates are >= to the rectangle's top left corner,
* and < the rectangle's x+w and y+h. So a 1x1 rectangle considers point (0,0)
* as "inside" and (0,1) as not.
*
* Note that this is a forced-inline function in a header, and not a public
* API function available in the SDL library (which is to say, the code is
* embedded in the calling program and the linker and dynamic loader will not
* be able to find this function inside SDL itself).
*
* \param p the point to test.
* \param r the rectangle to test.
* \returns true if `p` is contained by `r`, false otherwise.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
func PointInRect(p *Point, r *Rect) bool {
	if p != nil && r != nil && (p.X >= r.X) && (p.X < (r.X + r.W)) && (p.Y >= r.Y) && (p.Y < (r.Y + r.H)) {
		return true
	} else {
		return false
	}
}

/**
* Determine whether a rectangle has no area.
*
* A rectangle is considered "empty" for this function if `r` is NULL, or if
* `r`'s width and/or height are <= 0.
*
* Note that this is a forced-inline function in a header, and not a public
* API function available in the SDL library (which is to say, the code is
* embedded in the calling program and the linker and dynamic loader will not
* be able to find this function inside SDL itself).
*
* \param r the rectangle to test.
* \returns true if the rectangle is "empty", false otherwise.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
func RectEmpty(r *Rect) bool {
	if (r == nil) || (r.W <= 0) || (r.H <= 0) {
		return true
	} else {
		return false
	}
}

/**
* Determine whether two rectangles are equal.
*
* Rectangles are considered equal if both are not NULL and each of their x,
* y, width and height match.
*
* Note that this is a forced-inline function in a header, and not a public
* API function available in the SDL library (which is to say, the code is
* embedded in the calling program and the linker and dynamic loader will not
* be able to find this function inside SDL itself).
*
* \param a the first rectangle to test.
* \param b the second rectangle to test.
* \returns true if the rectangles are equal, false otherwise.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
func RectsEqual(a *Rect, b *Rect) bool {
	if (a != nil) && (b != nil) && (a.X == b.X) && (a.Y == b.Y) && (a.W == b.W) && (a.H == b.H) {
		return true
	} else {
		return false
	}
}

/**
* Determine whether two rectangles intersect.
*
* If either pointer is NULL the function will return false.
*
* \param A an SDL_Rect structure representing the first rectangle.
* \param B an SDL_Rect structure representing the second rectangle.
* \returns true if there is an intersection, false otherwise.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetRectIntersection
 */
//go:sdl3extern
var HasRectIntersection func(A *Rect, B *Rect) bool

/**
* Calculate the intersection of two rectangles.
*
* If `result` is NULL then this function will return false.
*
* \param A an SDL_Rect structure representing the first rectangle.
* \param B an SDL_Rect structure representing the second rectangle.
* \param result an SDL_Rect structure filled in with the intersection of
*               rectangles `A` and `B`.
* \returns true if there is an intersection, false otherwise.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_HasRectIntersection
 */
//go:sdl3extern
var GetRectIntersection func(A *Rect, B *Rect, result *Rect) bool

/**
* Calculate the union of two rectangles.
*
* \param A an SDL_Rect structure representing the first rectangle.
* \param B an SDL_Rect structure representing the second rectangle.
* \param result an SDL_Rect structure filled in with the union of rectangles
*               `A` and `B`.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetRectUnion func(A *Rect, B *Rect, result *Rect) bool

/**
* Calculate a minimal rectangle enclosing a set of points.
*
* If `clip` is not NULL then only points inside of the clipping rectangle are
* considered.
*
* \param points an array of SDL_Point structures representing points to be
*               enclosed.
* \param count the number of structures in the `points` array.
* \param clip an SDL_Rect used for clipping or NULL to enclose all points.
* \param result an SDL_Rect structure filled in with the minimal enclosing
*               rectangle.
* \returns true if any points were enclosed or false if all the points were
*          outside of the clipping rectangle.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetRectEnclosingPoints func(points *Point, count int, clip *Rect, result *Rect) bool

/**
* Calculate the intersection of a rectangle and line segment.
*
* This function is used to clip a line segment to a rectangle. A line segment
* contained entirely within the rectangle or that does not intersect will
* remain unchanged. A line segment that crosses the rectangle at either or
* both ends will be clipped to the boundary of the rectangle and the new
* coordinates saved in `X1`, `Y1`, `X2`, and/or `Y2` as necessary.
*
* \param rect an SDL_Rect structure representing the rectangle to intersect.
* \param X1 a pointer to the starting X-coordinate of the line.
* \param Y1 a pointer to the starting Y-coordinate of the line.
* \param X2 a pointer to the ending X-coordinate of the line.
* \param Y2 a pointer to the ending Y-coordinate of the line.
* \returns true if there is an intersection, false otherwise.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetRectAndLineIntersection func(rect *Rect, X1 *int, Y1 *int, X2 *int, Y2 *int) bool

/* SDL_FRect versions... */

/**
* Determine whether a point resides inside a floating point rectangle.
*
* A point is considered part of a rectangle if both `p` and `r` are not NULL,
* and `p`'s x and y coordinates are >= to the rectangle's top left corner,
* and <= the rectangle's x+w and y+h. So a 1x1 rectangle considers point
* (0,0) and (0,1) as "inside" and (0,2) as not.
*
* Note that this is a forced-inline function in a header, and not a public
* API function available in the SDL library (which is to say, the code is
* embedded in the calling program and the linker and dynamic loader will not
* be able to find this function inside SDL itself).
*
* \param p the point to test.
* \param r the rectangle to test.
* \returns true if `p` is contained by `r`, false otherwise.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
func PointInRectFloat(p *FPoint, r *FRect) bool {
	if (p != nil) && (r != nil) && (p.X >= r.X) && (p.X <= (r.X + r.W)) &&
		(p.Y >= r.Y) && (p.Y <= (r.Y + r.H)) {
		return true
	} else {
		return false
	}
}

/**
* Determine whether a floating point rectangle takes no space.
*
* A rectangle is considered "empty" for this function if `r` is NULL, or if
* `r`'s width and/or height are < 0.0f.
*
* Note that this is a forced-inline function in a header, and not a public
* API function available in the SDL library (which is to say, the code is
* embedded in the calling program and the linker and dynamic loader will not
* be able to find this function inside SDL itself).
*
* \param r the rectangle to test.
* \returns true if the rectangle is "empty", false otherwise.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
 */
func RectEmptyFloat(r *FRect) bool {
	if (r == nil) || (r.W < 0.0) || (r.H < 0.0) {
		return true
	} else {
		return false
	}
}

/**
* Determine whether two floating point rectangles are equal, within some
* given epsilon.
*
* Rectangles are considered equal if both are not NULL and each of their x,
* y, width and height are within `epsilon` of each other. If you don't know
* what value to use for `epsilon`, you should call the SDL_RectsEqualFloat
* function instead.
*
* Note that this is a forced-inline function in a header, and not a public
* API function available in the SDL library (which is to say, the code is
* embedded in the calling program and the linker and dynamic loader will not
* be able to find this function inside SDL itself).
*
* \param a the first rectangle to test.
* \param b the second rectangle to test.
* \param epsilon the epsilon value for comparison.
* \returns true if the rectangles are equal, false otherwise.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_RectsEqualFloat
 */
func RectsEqualEpsilon(a *FRect, b *FRect, epsilon float32) bool {
	if (a != nil) && (b != nil) && ((a == b) ||
		((float32(math.Abs(float64(a.X-b.X))) <= epsilon) &&
			(float32(math.Abs(float64(a.Y-b.Y))) <= epsilon) &&
			(float32(math.Abs(float64(a.W-b.W))) <= epsilon) &&
			(float32(math.Abs(float64(a.H-b.H))) <= epsilon))) {
		return true
	} else {
		return false
	}
}

/**
* Determine whether two floating point rectangles are equal, within a default
* epsilon.
*
* Rectangles are considered equal if both are not NULL and each of their x,
* y, width and height are within SDL_FLT_EPSILON of each other. This is often
* a reasonable way to compare two floating point rectangles and deal with the
* slight precision variations in floating point calculations that tend to pop
* up.
*
* Note that this is a forced-inline function in a header, and not a public
* API function available in the SDL library (which is to say, the code is
* embedded in the calling program and the linker and dynamic loader will not
* be able to find this function inside SDL itself).
*
* \param a the first rectangle to test.
* \param b the second rectangle to test.
* \returns true if the rectangles are equal, false otherwise.
*
* \threadsafety It is safe to call this function from any thread.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_RectsEqualEpsilon
 */
func RectsEqualFloat(a *FRect, b *FRect) bool {
	return RectsEqualEpsilon(a, b, FLT_EPSILON)
}

/**
* Determine whether two rectangles intersect with float precision.
*
* If either pointer is NULL the function will return false.
*
* \param A an SDL_FRect structure representing the first rectangle.
* \param B an SDL_FRect structure representing the second rectangle.
* \returns true if there is an intersection, false otherwise.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_GetRectIntersection
 */
//go:sdl3extern
var HasRectIntersectionFloat func(A *FRect, B *FRect) bool

/**
* Calculate the intersection of two rectangles with float precision.
*
* If `result` is NULL then this function will return false.
*
* \param A an SDL_FRect structure representing the first rectangle.
* \param B an SDL_FRect structure representing the second rectangle.
* \param result an SDL_FRect structure filled in with the intersection of
*               rectangles `A` and `B`.
* \returns true if there is an intersection, false otherwise.
*
* \since This function is available since SDL 3.2.0.
*
* \sa SDL_HasRectIntersectionFloat
 */
//go:sdl3extern
var GetRectIntersectionFloat func(A *FRect, B *FRect, result *FRect) bool

/**
* Calculate the union of two rectangles with float precision.
*
* \param A an SDL_FRect structure representing the first rectangle.
* \param B an SDL_FRect structure representing the second rectangle.
* \param result an SDL_FRect structure filled in with the union of rectangles
*               `A` and `B`.
* \returns true on success or false on failure; call SDL_GetError() for more
*          information.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetRectUnionFloat func(A *FRect, B *FRect, result *FRect) bool

/**
* Calculate a minimal rectangle enclosing a set of points with float
* precision.
*
* If `clip` is not NULL then only points inside of the clipping rectangle are
* considered.
*
* \param points an array of SDL_FPoint structures representing points to be
*               enclosed.
* \param count the number of structures in the `points` array.
* \param clip an SDL_FRect used for clipping or NULL to enclose all points.
* \param result an SDL_FRect structure filled in with the minimal enclosing
*               rectangle.
* \returns true if any points were enclosed or false if all the points were
*          outside of the clipping rectangle.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetRectEnclosingPointsFloat func(points *FPoint, count int, clip *FRect, result *FRect) bool

/**
* Calculate the intersection of a rectangle and line segment with float
* precision.
*
* This function is used to clip a line segment to a rectangle. A line segment
* contained entirely within the rectangle or that does not intersect will
* remain unchanged. A line segment that crosses the rectangle at either or
* both ends will be clipped to the boundary of the rectangle and the new
* coordinates saved in `X1`, `Y1`, `X2`, and/or `Y2` as necessary.
*
* \param rect an SDL_FRect structure representing the rectangle to intersect.
* \param X1 a pointer to the starting X-coordinate of the line.
* \param Y1 a pointer to the starting Y-coordinate of the line.
* \param X2 a pointer to the ending X-coordinate of the line.
* \param Y2 a pointer to the ending Y-coordinate of the line.
* \returns true if there is an intersection, false otherwise.
*
* \since This function is available since SDL 3.2.0.
 */
//go:sdl3extern
var GetRectAndLineIntersectionFloat func(rect *FRect, X1 *float32, Y1 *float32, X2 *float32, Y2 *float32) bool
