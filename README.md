#cgo-exmaple

There are five examples of using cgo.

## builtin

C code is built in Go program.

usage: `cd builtin && go build && ./builtin`

## c-lib

Using C static library in Go program.

The function `add` was written in C and built in static library: `add.a`.

usage: `cd c-lib && make && ./c-lib`

## c-so

Using C shared library in Go program.

The function `add` was written in C and built in shared library: `add.so`.

usage: `cd c-so && make && ./c-so`

## go-lib

Calling Go function in C program.

The function `add` was written in Go and built in static library.

Then called by `main.c`.

usage: `cd go-lib && make && ./go-lib`

## go-so

Calling go function in c program.

The function `add` was written in Go and built in shared library.

Then called by `main.c`.

usage: `cd go-so && make && ./go-so`

