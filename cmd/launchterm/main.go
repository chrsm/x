package main

import (
	"log"
	"os/exec"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"

	"github.com/AllenDang/w32"
)

// alternative to `launch-i3.vbs` that doesn't even launch i3 anymore :-)
func main() {
	var running bool
	log.Println("Starting terminal launch..")

	procs := make([]uint32, 1024)
	needed := uint32(0)
	if !w32.EnumProcesses(procs, 1024, &needed) {
		log.Fatal("failed to enumerate process list")
	}

	for i := uint32(0); i < needed/4; i++ {
		if procs[i] == 0 {
			continue
		}

		hnd, err := w32.OpenProcess(w32.PROCESS_QUERY_INFORMATION|w32.PROCESS_VM_READ, false, uintptr(procs[i]))
		if err != nil && err.Error() != "Access is denied." {
			log.Printf("failed to open proc(%d): %s", procs[i], err)
			continue
		}

		var (
			mod    w32.HMODULE
			needed uint32
		)
		if !enumProcessModules(hnd, &mod, &needed) {
			log.Printf("failed to enum proc(%d) modules", procs[i])
			continue
		}

		name := getModuleBaseName(hnd, mod)
		if name != "" {
			log.Printf("proc(%d) name: %s", procs[i], name)
		}

		if name == "vcxsrv.exe" {
			log.Printf("vcxsrv.exe is running: proc(%d)", procs[i])
			running = true
		}
	}

	if !running {
		// start vcxsrv
		cmd := exec.Command("vcxsrv.exe", ":0", "-multiwindow", "-nowgl")
		cmd.Start()
	}

	// now, launch a terminal in that session
	cmd := exec.Command("bash.exe", "-c", "-l", "DISPLAY=:0 alacritty")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("failed to run bash.exe: %s", err)
	}

	log.Println("peace out")
}

var (
	psapi = syscall.NewLazyDLL("psapi.dll")

	pEnumProcessModules = psapi.NewProc("EnumProcessModules")
	pGetModuleBaseName  = psapi.NewProc("GetModuleBaseNameW")

	szhmod = w32.HMODULE(0)
)

func enumProcessModules(hproc w32.HANDLE, mod *w32.HMODULE, needed *uint32) bool {
	ret, _, _ := pEnumProcessModules.Call(
		uintptr(hproc),
		uintptr(unsafe.Pointer(mod)),
		uintptr(unsafe.Sizeof(szhmod)),
		uintptr(unsafe.Pointer(needed)),
	)

	return ret != 0
}

func getModuleBaseName(hproc w32.HANDLE, mod w32.HMODULE) string {
	var name = make([]uint16, 32)

	ret, _, _ := pGetModuleBaseName.Call(
		uintptr(hproc),
		uintptr(mod),
		uintptr(unsafe.Pointer(&name[0])),
		uintptr(len(name)),
	)

	if ret != 0 {
		return windows.UTF16ToString(name[:ret])
	}

	return ""
}
