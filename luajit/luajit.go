package luajit

/*
#cgo LDFLAGS: -lluajit-5.1 -DLUAJIT_ENABLE_LUA52COMPAT

#include <stdlib.h>
#include <stdio.h>
#include <luajit-2.0/lua.h>
#include <luajit-2.0/lualib.h>
#include <luajit-2.0/lauxlib.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

type LuaState *C.struct_lua_State

type LState struct {
	vm LuaState
}

func NewState() *LState {
	vm := C.luaL_newstate()
	C.luaL_openlibs(vm)

	ls := &LState{vm}
	// ls.expose(get_number)

	return ls
}

// DoString executes the string `src` as lua code.
func (l *LState) DoString(src string) error {
	code := C.CString(src)
	defer C.free(unsafe.Pointer(code))

	if C.luaL_loadstring(l.vm, code) != 0 {
		return l.error()
	}

	return l.Run()
}

func (l *LState) DoFile(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	if C.luaL_loadfile(l.vm, cpath) != 0 {
		return l.error()
	}

	return l.Run()
}

func (l *LState) Run() error {
	if C.lua_pcall(l.vm, 0, 0, 0) != 0 {
		return l.error()
	}

	return nil
}

func (l *LState) Close() {
	C.lua_close(l.vm)
}

func (l *LState) error() error {
	return errors.New(C.GoString(C.lua_tolstring(l.vm, -1, nil)))
}

const exposedFuncs = "gluajit.exposed"

func (l *LState) expose(fn interface{}) error {
	str := C.CString(exposedFuncs)

	ptr := C.lua_newuserdata(l.vm, 4)
	_ = ptr

	// C.luaL_register(l.vm, str, ptr)
	C.lua_getfield(l.vm, -1001, str)
	C.lua_setmetatable(l.vm, -2)

	return nil
}
