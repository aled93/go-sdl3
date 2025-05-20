# SDL3 binding for GoLang

Goal of this project is to provide complete SDL3 API usable with CGo or PureGo

# Status

Full API translated (except ones that have similar functionality in go standard packages, see [STATUS.MD](/STATUS.MD)).

Need test that everything (un-)marshaling correctly

# Installation

1. Build SDL3 as shared library, see how to do it [in SDL3 repo](https://github.com/libsdl-org/SDL/blob/main/INSTALL.md#to-build-and-use-sdl)
2. `go get https://github.com/aled93/go-sdl3`

