//go:build windows

package user32

import "golang.org/x/sys/windows"

var (
	dll                   *windows.DLL
	procCreateWindowExW   *windows.Proc
	procDefWindowProcW    *windows.Proc
	procDestroyWindow     *windows.Proc
	procDispatchMessageW  *windows.Proc
	procGetDC             *windows.Proc
	procGetKeyboardLayout *windows.Proc
	procGetMessageW       *windows.Proc
	procGetSystemMetrics  *windows.Proc
	procGetWindowLongPtrW *windows.Proc
	procGetWindowRect     *windows.Proc
	procLoadCursorW       *windows.Proc
	procLoadIconW         *windows.Proc
	procMapVirtualKeyExW  *windows.Proc
	procMapVirtualKeyW    *windows.Proc
	procPeekMessageW      *windows.Proc
	procPostQuitMessage   *windows.Proc
	procRegisterClassExW  *windows.Proc
	procReleaseDC         *windows.Proc
	procSendMessageW      *windows.Proc
	procSetFocus          *windows.Proc
	procSetWindowLongPtrW *windows.Proc
	procSetWindowPos      *windows.Proc
	procShowCursor        *windows.Proc
	procShowWindow        *windows.Proc
	procTranslateMessage  *windows.Proc
	procUpdateWindow      *windows.Proc
)

func init() {
	dll = windows.MustLoadDLL("user32.dll")
	procCreateWindowExW = dll.MustFindProc("CreateWindowExW")
	procDefWindowProcW = dll.MustFindProc("DefWindowProcW")
	procDestroyWindow = dll.MustFindProc("DestroyWindow")
	procDispatchMessageW = dll.MustFindProc("DispatchMessageW")
	procGetDC = dll.MustFindProc("GetDC")
	procGetKeyboardLayout = dll.MustFindProc("GetKeyboardLayout")
	procGetMessageW = dll.MustFindProc("GetMessageW")
	procGetSystemMetrics = dll.MustFindProc("GetSystemMetrics")
	procGetWindowLongPtrW = dll.MustFindProc("GetWindowLongPtrW")
	procGetWindowRect = dll.MustFindProc("GetWindowRect")
	procLoadCursorW = dll.MustFindProc("LoadCursorW")
	procLoadIconW = dll.MustFindProc("LoadIconW")
	procMapVirtualKeyExW = dll.MustFindProc("MapVirtualKeyExW")
	procMapVirtualKeyW = dll.MustFindProc("MapVirtualKeyW")
	procPeekMessageW = dll.MustFindProc("PeekMessageW")
	procPostQuitMessage = dll.MustFindProc("PostQuitMessage")
	procRegisterClassExW = dll.MustFindProc("RegisterClassExW")
	procReleaseDC = dll.MustFindProc("ReleaseDC")
	procSendMessageW = dll.MustFindProc("SendMessageW")
	procSetFocus = dll.MustFindProc("SetFocus")
	procSetWindowLongPtrW = dll.MustFindProc("SetWindowLongPtrW")
	procSetWindowPos = dll.MustFindProc("SetWindowPos")
	procShowCursor = dll.MustFindProc("ShowCursor")
	procShowWindow = dll.MustFindProc("ShowWindow")
	procTranslateMessage = dll.MustFindProc("TranslateMessage")
	procUpdateWindow = dll.MustFindProc("UpdateWindow")
}
