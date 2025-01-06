//go:build windows

package kernel32

import (
	"unsafe"

	"github.com/o5h/winapi"
)

func GetLastError() winapi.DWORD {
	r0, _, _ := procGetLastError.Call()
	return winapi.DWORD(r0)
}

func GetModuleHandle(moduleName winapi.LPCWSTR) (winapi.HMODULE, error) {
	r0, _, lastErr := procGetModuleHandleW.Call(uintptr(unsafe.Pointer(moduleName)))
	return winapi.HMODULE(r0), lastErr
}

func QueryPerformanceFrequency(lpFrequency *winapi.LARGE_INTEGER) winapi.BOOL {
	r0, _, _ := procQueryPerformanceFrequency.Call(uintptr(unsafe.Pointer(lpFrequency)))
	return winapi.BOOL(r0)
}

func QueryPerformanceCounter(lpPerformanceCount *winapi.LARGE_INTEGER) winapi.BOOL {
	r0, _, _ := procQueryPerformanceCounter.Call(uintptr(unsafe.Pointer(lpPerformanceCount)))
	return winapi.BOOL(r0)
}

func Sleep(milliseconds winapi.DWORD) {
	procSleep.Call(uintptr(milliseconds))
}
