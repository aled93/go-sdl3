# Generate loader
This generator generates loader code for all functions marked with `//go:sdl3extern` using purego

# Usage
Run this program in root of repository when any function marked with `//go:sdl3extern` added, removed or changed:

    go run ./cmd/genloader