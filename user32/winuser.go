// +build windows

package user32

import (
	"unsafe"

	"github.com/o5h/winapi"
)

var (
	IDI_APPLICATION = MakeIntResource(32512) //Default application icon.
	IDI_ERROR       = MakeIntResource(32513) //Hand-shaped icon.
	IDI_HAND        = MakeIntResource(32513) //Hand-shaped icon. Same as IDI_ERROR.
	IDI_QUESTION    = MakeIntResource(32514) //Question mark icon.
	IDI_WARNING     = MakeIntResource(32515) //Exclamation point icon.
	IDI_EXCLAMATION = MakeIntResource(32515) //Exclamation point icon. Same as IDI_WARNING.
	IDI_ASTERISK    = MakeIntResource(32516) //Asterisk icon. Same as IDI_INFORMATION.
	IDI_INFORMATION = MakeIntResource(32516) //Asterisk icon.
	IDI_WINLOGO     = MakeIntResource(32517) //Default application icon.Windows 2000:  Windows logo icon.
	IDI_SHIELD      = MakeIntResource(32518) //Security Shield icon.
)

func LoadIconW(instance winapi.HINSTANCE, iconName winapi.LPCWSTR) (winapi.HICON, error) {
	r0, _, lastErr := procLoadIconW.Call(uintptr(instance), uintptr(unsafe.Pointer(iconName)))
	return winapi.HICON(r0), lastErr
}

var (
	IDC_APPSTARTING = MakeIntResource(32650) //Standard arrow and small hourglass
	IDC_ARROW       = MakeIntResource(32512) //Standard arrow
	IDC_CROSS       = MakeIntResource(32515) //Crosshair
	IDC_HAND        = MakeIntResource(32649) //Hand
	IDC_HELP        = MakeIntResource(32651) //Arrow and question mark
	IDC_IBEAM       = MakeIntResource(32513) //I-beam
	IDC_ICON        = MakeIntResource(32641) //Obsolete for applications marked version 4.0 or later.
	IDC_NO          = MakeIntResource(32648) //Slashed circle
	IDC_SIZE        = MakeIntResource(32640) //Obsolete for applications marked version 4.0 or later. Use IDC_SIZEALL.
	IDC_SIZEALL     = MakeIntResource(32646) //Four-pointed arrow pointing north, south, east, and west
	IDC_SIZENESW    = MakeIntResource(32643) //Double-pointed arrow pointing northeast and southwest
	IDC_SIZENS      = MakeIntResource(32645) //Double-pointed arrow pointing north and south
	IDC_SIZENWSE    = MakeIntResource(32642) //Double-pointed arrow pointing northwest and southeast
	IDC_SIZEWE      = MakeIntResource(32644) //Double-pointed arrow pointing west and east
	IDC_UPARROW     = MakeIntResource(32516) //Vertical arrow
	IDC_WAIT        = MakeIntResource(32514) //Hourglass
)

func LoadCursorW(instance winapi.HINSTANCE, cursorName winapi.LPCWSTR) (winapi.HCURSOR, error) {
	r0, _, lastErr := procLoadCursorW.Call(uintptr(instance), uintptr(unsafe.Pointer(cursorName)))
	return winapi.HCURSOR(r0), lastErr
}

type WNDCLASSEX struct {
	Size       uint32
	Style      uint32
	WndProc    winapi.WNDPROC
	ClsExtra   int32
	WndExtra   int32
	Instance   winapi.HINSTANCE
	Icon       winapi.HICON
	Cursor     winapi.HCURSOR
	Background winapi.HBRUSH
	MenuName   winapi.LPCWSTR
	ClassName  winapi.LPCWSTR
	IconSm     winapi.HICON
}

type CREATESTRUCTW struct {
	CreateParams winapi.LPVOID
	Instance     winapi.HINSTANCE
	Menu         winapi.HMENU
	HwndParent   winapi.HWND
	Cy           int
	Cx           int
	Y            int
	X            int
	Style        winapi.LONG
	Name         winapi.LPCWSTR
	lass         winapi.LPCWSTR
	ExStyle      winapi.DWORD
}

const (
	SWP_ASYNCWINDOWPOS = 0x4000
	SWP_DEFERERASE     = 0x2000
	SWP_DRAWFRAME      = 0x0020
	SWP_FRAMECHANGED   = 0x0020
	SWP_HIDEWINDOW     = 0x0080
	SWP_NOACTIVATE     = 0x0010
	SWP_NOCOPYBITS     = 0x0100
	SWP_NOMOVE         = 0x0002
	SWP_NOOWNERZORDER  = 0x0200
	SWP_NOREDRAW       = 0x0008
	SWP_NOREPOSITION   = 0x0200
	SWP_NOSENDCHANGING = 0x0400
	SWP_NOSIZE         = 0x0001
	SWP_NOZORDER       = 0x0004
	SWP_SHOWWINDOW     = 0x0040
)

func DefWindowProcW(
	hwnd winapi.HWND,
	msg winapi.UINT,
	wparam winapi.WPARAM,
	lparam winapi.LPARAM) winapi.LRESULT {

	r0, _, _ := procDefWindowProcW.Call(
		uintptr(hwnd),
		uintptr(msg),
		uintptr(wparam),
		uintptr(lparam))
	return winapi.LRESULT(r0)
}

func DestroyWindow(hwnd winapi.HWND) (succeed bool, err error) {
	r0, _, err := procDestroyWindow.Call(uintptr(hwnd))
	return bool(r0 != 0), err
}

func PostQuitMessage(exitcode int32) {
	procPostQuitMessage.Call(uintptr(exitcode))
	return
}

type Point struct {
	X uintptr
	Y uintptr
}

type Rect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type Msg struct {
	Hwnd    winapi.HWND
	Message winapi.UINT
	Wparam  winapi.WPARAM
	Lparam  winapi.LPARAM
	Time    uint32
	Pt      Point
}

func GetMessageW(msg *Msg, hwnd winapi.HWND, MsgFilterMin uint32, MsgFilterMax uint32) (winapi.BOOL, error) {
	r0, _, err := procGetMessageW.Call(uintptr(unsafe.Pointer(msg)), uintptr(hwnd), uintptr(MsgFilterMin), uintptr(MsgFilterMax), 0, 0)
	return winapi.BOOL(r0), err
}

func TranslateMessage(msg *Msg) bool {
	r0, _, _ := procTranslateMessage.Call(uintptr(unsafe.Pointer(msg)))
	return r0 != 0
}

func DispatchMessageW(msg *Msg) winapi.LRESULT {
	r0, _, _ := procDispatchMessageW.Call(uintptr(unsafe.Pointer(msg)))
	return winapi.LRESULT(r0)
}
