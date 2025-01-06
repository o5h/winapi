//go:build windows

package kernel32

import "golang.org/x/sys/windows"

var (
	dll                           *windows.DLL
	procGetLastError              *windows.Proc
	procGetModuleHandleW          *windows.Proc
	procQueryPerformanceCounter   *windows.Proc
	procQueryPerformanceFrequency *windows.Proc
	procSleep                     *windows.Proc
)

func init() {
	dll = windows.MustLoadDLL("kernel32.dll")
	procGetLastError = dll.MustFindProc("GetLastError")
	procGetModuleHandleW = dll.MustFindProc("GetModuleHandleW")
	procQueryPerformanceCounter = dll.MustFindProc("QueryPerformanceCounter")
	procQueryPerformanceFrequency = dll.MustFindProc("QueryPerformanceFrequency")
	procSleep = dll.MustFindProc("Sleep")
}
