//go:build windows

package kernel32

import "golang.org/x/sys/windows"

var (
	dll                           *windows.DLL
	procGetModuleHandleW          *windows.Proc
	procGetLastError              *windows.Proc
	procQueryPerformanceFrequency *windows.Proc
	procQueryPerformanceCounter   *windows.Proc
	procSleep                     *windows.Proc
)

func init() {
	dll = windows.MustLoadDLL("kernel32.dll")
	procGetModuleHandleW = dll.MustFindProc("GetModuleHandleW")
	procGetLastError = dll.MustFindProc("GetLastError")
	procQueryPerformanceFrequency = dll.MustFindProc("QueryPerformanceFrequency")
	procQueryPerformanceCounter = dll.MustFindProc("QueryPerformanceCounter")
	procSleep = dll.MustFindProc("Sleep")
}
