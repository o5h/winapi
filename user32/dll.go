//go:build windows

package user32

import "golang.org/x/sys/windows"

var (
	dll                   *windows.DLL
	procLoadCursorW       *windows.Proc
	procLoadIconW         *windows.Proc
	procRegisterClassExW  *windows.Proc
	procCreateWindowExW   *windows.Proc
	procShowWindow        *windows.Proc
	procUpdateWindow      *windows.Proc
	procSetFocus          *windows.Proc
	procDefWindowProcW    *windows.Proc
	procDestroyWindow     *windows.Proc
	procPostQuitMessage   *windows.Proc
	procGetMessageW       *windows.Proc
	procTranslateMessage  *windows.Proc
	procDispatchMessageW  *windows.Proc
	procSetWindowPos      *windows.Proc
	procGetWindowRect     *windows.Proc
	procGetDC             *windows.Proc
	procReleaseDC         *windows.Proc
	procShowCursor        *windows.Proc
	procPeekMessageW      *windows.Proc
	procSendMessageW      *windows.Proc
	procSetWindowLongPtrW *windows.Proc
	procGetWindowLongPtrW *windows.Proc
	procMapVirtualKeyExW  *windows.Proc
	procGetKeyboardLayout *windows.Proc
	procMapVirtualKeyW    *windows.Proc
)

func init() {
	dll = windows.MustLoadDLL("user32.dll")
	procLoadIconW = dll.MustFindProc("LoadIconW")
	procLoadCursorW = dll.MustFindProc("LoadCursorW")
	procRegisterClassExW = dll.MustFindProc("RegisterClassExW")
	procCreateWindowExW = dll.MustFindProc("CreateWindowExW")
	procShowWindow = dll.MustFindProc("ShowWindow")
	procUpdateWindow = dll.MustFindProc("UpdateWindow")
	procSetFocus = dll.MustFindProc("SetFocus")
	procDefWindowProcW = dll.MustFindProc("DefWindowProcW")
	procDestroyWindow = dll.MustFindProc("DestroyWindow")
	procPostQuitMessage = dll.MustFindProc("PostQuitMessage")
	procGetMessageW = dll.MustFindProc("GetMessageW")
	procTranslateMessage = dll.MustFindProc("TranslateMessage")
	procDispatchMessageW = dll.MustFindProc("DispatchMessageW")
	procSetWindowPos = dll.MustFindProc("SetWindowPos")
	procGetWindowRect = dll.MustFindProc("GetWindowRect")
	procGetDC = dll.MustFindProc("GetDC")
	procReleaseDC = dll.MustFindProc("ReleaseDC")
	procShowCursor = dll.MustFindProc("ShowCursor")
	procPeekMessageW = dll.MustFindProc("PeekMessageW")
	procSendMessageW = dll.MustFindProc("SendMessageW")
	procSetWindowLongPtrW = dll.MustFindProc("SetWindowLongPtrW")
	procGetWindowLongPtrW = dll.MustFindProc("GetWindowLongPtrW")
	procMapVirtualKeyExW = dll.MustFindProc("MapVirtualKeyExW")
	procGetKeyboardLayout = dll.MustFindProc("GetKeyboardLayout")
	procMapVirtualKeyW = dll.MustFindProc("MapVirtualKeyW")
}
