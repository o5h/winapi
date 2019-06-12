// +build windows

package user32

import (
	"unsafe"

	"github.com/o5h/winapi"
)

func MakeIntResource(id uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}

const (
	HWND_BOTTOM    = winapi.HWND(1)
	HWND_NOTOPMOST = -2
	HWND_TOP       = winapi.HWND(0)
	HWND_TOPMOST   = -1
)
const (
	BS_DEFPUSHBUTTON = 1
	COLOR_WINDOW     = 5
	COLOR_BTNFACE    = 15
	CW_USEDEFAULT    = 0x80000000 - 0x100000000 //// Default window position
)

func RegisterClassExW(wndclass *WNDCLASSEX) (winapi.ATOM, error) {
	r0, _, err := procRegisterClassExW.Call(uintptr(unsafe.Pointer(wndclass)))
	return winapi.ATOM(r0), err

}

//Window Styles
const (
	WS_OVERLAPPED   = 0
	WS_POPUP        = 0x80000000
	WS_CHILD        = 0x40000000
	WS_MINIMIZE     = 0x20000000
	WS_VISIBLE      = 0x10000000
	WS_DISABLED     = 0x08000000
	WS_CLIPSIBLINGS = 0x04000000
	WS_CLIPCHILDREN = 0x02000000
	WS_MAXIMIZE     = 0x01000000
	WS_CAPTION      = WS_BORDER | WS_DLGFRAME
	WS_BORDER       = 0x00800000
	WS_DLGFRAME     = 0x00400000
	WS_VSCROLL      = 0x00200000
	WS_HSCROLL      = 0x00100000
	WS_SYSMENU      = 0x00080000
	WS_THICKFRAME   = 0x00040000
	WS_GROUP        = 0x00020000
	WS_TABSTOP      = 0x00010000
	WS_MINIMIZEBOX  = 0x00020000
	WS_MAXIMIZEBOX  = 0x00010000
	WS_TILED        = WS_OVERLAPPED
	WS_ICONIC       = WS_MINIMIZE
	WS_SIZEBOX      = WS_THICKFRAME
	// Common Window Styles
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_TILEDWINDOW      = WS_OVERLAPPEDWINDOW
	WS_POPUPWINDOW      = WS_POPUP | WS_BORDER | WS_SYSMENU
	WS_CHILDWINDOW      = WS_CHILD
)

const (
	WM_NULL              = winapi.UINT(0x0000)
	WM_CREATE            = winapi.UINT(0x0001)
	WM_DESTROY           = winapi.UINT(0x0002)
	WM_NCDESTROY         = winapi.UINT(0x0002)
	WM_MOVE              = winapi.UINT(0x0003)
	WM_SIZE              = winapi.UINT(0x0005)
	WM_ACTIVATE          = winapi.UINT(0x0006)
	WM_SETFOCUS          = winapi.UINT(0x0007)
	WM_KILLFOCUS         = winapi.UINT(0x0008)
	WM_ENABLE            = winapi.UINT(0x000A)
	WM_SETREDRAW         = winapi.UINT(0x000B)
	WM_SETTEXT           = winapi.UINT(0x000C)
	WM_GETTEXT           = winapi.UINT(0x000D)
	WM_GETTEXTLENGTH     = winapi.UINT(0x000E)
	WM_PAINT             = winapi.UINT(0x000F)
	WM_CLOSE             = winapi.UINT(0x0010)
	WM_QUIT              = winapi.UINT(0x0012)
	WM_ERASEBKGND        = winapi.UINT(0x0014)
	WM_SHOWWINDOW        = winapi.UINT(0x0018)
	WM_CTLCOLOR          = winapi.UINT(0x0019)
	WM_NEXTDLGCTL        = winapi.UINT(0x0028)
	WM_DRAWITEM          = winapi.UINT(0x002B)
	WM_MEASUREITEM       = winapi.UINT(0x002C)
	WM_DELETEITEM        = winapi.UINT(0x002D)
	WM_VKEYTOITEM        = winapi.UINT(0x002E)
	WM_CHARTOITEM        = winapi.UINT(0x002F)
	WM_SETFONT           = winapi.UINT(0x0030)
	WM_GETFONT           = winapi.UINT(0x0031)
	WM_COMPAREITEM       = winapi.UINT(0x0039)
	WM_WINDOWPOSCHANGED  = winapi.UINT(0x0047)
	WM_NOTIFY            = winapi.UINT(0x004E)
	WM_NCCALCSIZE        = winapi.UINT(0x0083)
	WM_NCHITTEST         = winapi.UINT(0x0084)
	WM_NCPAINT           = winapi.UINT(0x0085)
	WM_GETDLGCODE        = winapi.UINT(0x0087)
	WM_NCMOUSEMOVE       = winapi.UINT(0x00A0)
	WM_NCLBUTTONDOWN     = winapi.UINT(0x00A1)
	WM_NCLBUTTONUP       = winapi.UINT(0x00A2)
	WM_NCLBUTTONDBLCLK   = winapi.UINT(0x00A3)
	WM_NCRBUTTONDOWN     = winapi.UINT(0x00A4)
	WM_NCRBUTTONUP       = winapi.UINT(0x00A5)
	WM_NCRBUTTONDBLCLK   = winapi.UINT(0x00A6)
	WM_KEYFIRST          = winapi.UINT(0x0100)
	WM_KEYDOWN           = winapi.UINT(0x0100)
	WM_KEYUP             = winapi.UINT(0x0101)
	WM_CHAR              = winapi.UINT(0x0102)
	WM_DEADCHAR          = winapi.UINT(0x0103)
	WM_SYSKEYDOWN        = winapi.UINT(0x0104)
	WM_SYSKEYUP          = winapi.UINT(0x0105)
	WM_SYSCHAR           = winapi.UINT(0x0106)
	WM_SYSDEADCHAR       = winapi.UINT(0x0107)
	WM_KEYLAST           = winapi.UINT(0x0108)
	WM_INITDIALOG        = winapi.UINT(0x0110)
	WM_COMMAND           = winapi.UINT(0x0111)
	WM_SYSCOMMAND        = winapi.UINT(0x0112)
	WM_TIMER             = winapi.UINT(0x0113)
	WM_HSCROLL           = winapi.UINT(0x0114)
	WM_VSCROLL           = winapi.UINT(0x0115)
	WM_ENTERIDLE         = winapi.UINT(0x0121)
	WM_CTLCOLORMSGBOX    = winapi.UINT(0x0132)
	WM_CTLCOLOREDIT      = winapi.UINT(0x0133)
	WM_CTLCOLORLISTBOX   = winapi.UINT(0x0134)
	WM_CTLCOLORBTN       = winapi.UINT(0x0135)
	WM_CTLCOLORDLG       = winapi.UINT(0x0136)
	WM_CTLCOLORSCROLLBAR = winapi.UINT(0x0137)
	WM_CTLCOLORSTATIC    = winapi.UINT(0x0138)
	WM_MOUSEFIRST        = winapi.UINT(0x0200)
	WM_MOUSEMOVE         = winapi.UINT(0x0200)
	WM_LBUTTONDOWN       = winapi.UINT(0x0201)
	WM_LBUTTONUP         = winapi.UINT(0x0202)
	WM_LBUTTONDBLCLK     = winapi.UINT(0x0203)
	WM_RBUTTONDOWN       = winapi.UINT(0x0204)
	WM_RBUTTONUP         = winapi.UINT(0x0205)
	WM_RBUTTONDBLCLK     = winapi.UINT(0x0206)
	WM_MBUTTONDOWN       = winapi.UINT(0x0207)
	WM_MBUTTONUP         = winapi.UINT(0x0208)
	WM_MBUTTONDBLCLK     = winapi.UINT(0x0209)
	WM_MOUSEWHEEL        = winapi.UINT(0x020A)
	WM_MOUSELAST         = winapi.UINT(0x020A)
	WM_HOTKEY            = winapi.UINT(0x0312)
	WM_CARET_CREATE      = winapi.UINT(0x03E0)
	WM_CARET_DESTROY     = winapi.UINT(0x03E1)
	WM_CARET_BLINK       = winapi.UINT(0x03E2)
	WM_FDINPUT           = winapi.UINT(0x03F0)
	WM_FDOUTPUT          = winapi.UINT(0x03F1)
	WM_FDEXCEPT          = winapi.UINT(0x03F2)
	WM_USER              = winapi.UINT(0x0400)
)

const (
	WMSZ_BOTTOM      = 6 //	Bottom edge
	WMSZ_BOTTOMLEFT  = 7 //	Bottom-left corner
	WMSZ_BOTTOMRIGHT = 8 //	Bottom-right corner
	WMSZ_LEFT        = 1 //	Left edge
	WMSZ_RIGHT       = 2 //	Right edge
	WMSZ_TOP         = 3 //	Top edge
	WMSZ_TOPLEFT     = 4 //	Top-left corner
	WMSZ_TOPRIGHT    = 5 //	Top-right corner
)

const (
	SC_CLOSE        = 0xF060     //	Closes the window.
	SC_CONTEXTHELP  = 0xF180     //	Changes the cursor to a question mark with a pointer. If the user then clicks a control in the dialog box, the control receives a WM_HELP message.
	SC_DEFAULT      = 0xF160     //	Selects the default item; the user double-clicked the window menu.
	SC_HOTKEY       = 0xF150     //	Activates the window associated with the application-specified hot key. The lParam parameter identifies the window to activate.
	SC_HSCROLL      = 0xF080     //	Scrolls horizontally.
	SCF_ISSECURE    = 0x00000001 //	Indicates whether the screen saver is secure.
	SC_KEYMENU      = 0xF100     //	Retrieves the window menu as a result of a keystroke. For more information, see the Remarks section.
	SC_MAXIMIZE     = 0xF030     //	Maximizes the window.
	SC_MINIMIZE     = 0xF020     //	Minimizes the window.
	SC_MONITORPOWER = 0xF170
	SC_MOUSEMENU    = 0xF090 //	Retrieves the window menu as a result of a mouse click.
	SC_MOVE         = 0xF010 //	Moves the window.
	SC_NEXTWINDOW   = 0xF040 //	Moves to the next window.
	SC_PREVWINDOW   = 0xF050 //	Moves to the previous window.
	SC_RESTORE      = 0xF120 //	Restores the window to its normal position and size.
	SC_SCREENSAVE   = 0xF140 //	Executes the screen saver application specified in the [boot] section of the System.ini file.
	SC_SIZE         = 0xF000 //	Sizes the window.
	SC_TASKLIST     = 0xF130 //	Activates the Start menu.
	SC_VSCROLL      = 0xF070 //	Scrolls vertically.
)

func CreateWindowExW(
	exstyle uint32,
	className winapi.LPCWSTR,
	windowName winapi.LPCWSTR,
	style uint32,
	x int32,
	y int32,
	width int32,
	height int32,
	wndparent winapi.HWND,
	menu winapi.HMENU,
	instance winapi.HINSTANCE,
	param winapi.LPVOID) (winapi.HWND, error) {

	r0, _, err :=
		procCreateWindowExW.Call(
			uintptr(exstyle),
			uintptr(unsafe.Pointer(className)),
			uintptr(unsafe.Pointer(windowName)),
			uintptr(style),
			uintptr(x),
			uintptr(y),
			uintptr(width),
			uintptr(height),
			uintptr(wndparent),
			uintptr(menu),
			uintptr(instance),
			uintptr(param))
	return winapi.HWND(r0), err

}

const (
	SW_FORCEMINIMIZE   = 11 //Minimizes a window, even if the thread that owns the window is not responding. This flag should only be used when minimizing windows from a different thread.
	SW_HIDE            = 0  //Hides the window and activates another window.
	SW_MAXIMIZE        = 3  //Maximizes the specified window.
	SW_MINIMIZE        = 6  //Minimizes the specified window and activates the next top-level window in the Z order.
	SW_RESTORE         = 9  //Activates and displays the window. If the window is minimized or maximized, the system restores it to its original size and position. An application should specify this flag when restoring a minimized window.
	SW_SHOW            = 5  //Activates the window and displays it in its current size and position.
	SW_SHOWDEFAULT     = 10 //Sets the show state based on the SW_ value specified in the STARTUPINFO structure passed to the CreateProcess function by the program that started the application.
	SW_SHOWMAXIMIZED   = 3  //Activates the window and displays it as a maximized window.
	SW_SHOWMINIMIZED   = 2  //Activates the window and displays it as a minimized window.
	SW_SHOWMINNOACTIVE = 7  //Displays the window as a minimized window. This value is similar to SW_SHOWMINIMIZED, except the window is not activated.
	SW_SHOWNA          = 8  //Displays the window in its current size and position. This value is similar to SW_SHOW, except that the window is not activated.
	SW_SHOWNOACTIVATE  = 4  //Displays a window in its most recent size and position. This value is similar to SW_SHOWNORMAL, except that the window is not activated.
	SW_SHOWNORMAL      = 1  //Activates and displays a window. If the window is minimized or maximized, the system restores it to its original size and position. An application should specify this flag when displaying the window for the first time.
)

func ShowWindow(hwnd winapi.HWND, cmdshow int32) bool {
	r0, _, _ := procShowWindow.Call(uintptr(hwnd), uintptr(cmdshow))
	return bool(r0 != 0)
}

func UpdateWindow(hwnd winapi.HWND) bool {
	r0, _, _ := procUpdateWindow.Call(uintptr(hwnd))
	return bool(r0 != 0)
}

func SetFocus(hwnd winapi.HWND) winapi.HWND {
	r0, _, _ := procSetFocus.Call(uintptr(hwnd))
	return winapi.HWND(r0)
}

func SetWindowPos(hwnd winapi.HWND, hWndInsertAfter winapi.HWND, x int32, y int32, cx int32, cy int32, uFlags winapi.UINT) bool {
	r0, _, _ := procSetWindowPos.Call(
		uintptr(hwnd),
		uintptr(hWndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(uFlags))
	return r0 != 0
}

func GetDC(hwnd winapi.HWND) (winapi.HDC, error) {
	r0, _, err := procGetDC.Call(uintptr(hwnd))
	return winapi.HDC(r0), err
}

func ReleaseDC(hwnd winapi.HWND, dc winapi.HDC) (int, error) {
	r0, _, err := procReleaseDC.Call(uintptr(hwnd), uintptr(dc))
	return int(r0), err
}

func GetWindowRect(hwnd winapi.HWND, rect *Rect) (bool, error) {
	r0, _, err := procGetWindowRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(rect)))
	return r0 != 0, err
}

func ShowCursor(show winapi.BOOL) (int, error) {
	r0, _, err := procShowCursor.Call(uintptr(show))
	return int(r0), err
}

const (
	PM_NOREMOVE = 0x0000
	PM_REMOVE   = 0x0001
	PM_NOYIELD  = 0x0002
)

func PeekMessageW(msg *Msg,
	hWnd winapi.HWND,
	wMsgFilterMin winapi.UINT,
	wMsgFilterMax winapi.UINT,
	wRemoveMsg winapi.UINT) (winapi.BOOL, error) {

	r0, _, err := procPeekMessageW.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg))

	return winapi.BOOL(r0), err
}

func SendMessageW(hWnd winapi.HWND, msg winapi.UINT, wparam uintptr, lparam uintptr) (winapi.LRESULT, error) {
	r0, _, err := procSendMessageW.Call(
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wparam),
		uintptr(lparam))
	return winapi.LRESULT(r0), err
}

const (
	GWL_EXSTYLE    = -20 //Sets a new extended window style.
	GWLP_HINSTANCE = -6  //Sets a new application instance handle.
	GWLP_ID        = -12 //Sets a new identifier of the child window. The window cannot be a top-level window.
	GWL_STYLE      = -16 //Sets a new window style.
	GWLP_USERDATA  = -21 //Sets the user data associated with the window. This data is intended for use by the application that created the window. Its value is initially zero.
	GWLP_WNDPROC   = -4  //Sets a new address for the window procedure.
)

func SetWindowLongPtrW(hWnd winapi.HWND, nIndex int, dwNewLong winapi.LONG_PTR) (winapi.LONG_PTR, error) {
	r0, _, err := procSetWindowLongPtrW.Call(
		uintptr(hWnd),
		uintptr(nIndex),
		uintptr(dwNewLong))
	return winapi.LONG_PTR(r0), err
}

func GetWindowLongPtrW(hWnd winapi.HWND, nIndex int) (winapi.LONG_PTR, error) {
	r0, _, err := procGetWindowLongPtrW.Call(
		uintptr(hWnd),
		uintptr(nIndex))
	return winapi.LONG_PTR(r0), err
}

const (
	MAPVK_VK_TO_CHAR   = 2
	MAPVK_VK_TO_VSC    = 0
	MAPVK_VK_TO_VSC_EX = 4
	MAPVK_VSC_TO_VK    = 1
	MAPVK_VSC_TO_VK_EX = 3
)

func MapVirtualKeyExW(uCode winapi.UINT, uMapType winapi.UINT, dwhkl winapi.HKL) winapi.UINT {
	r0, _, _ := procMapVirtualKeyExW.Call(
		uintptr(uCode),
		uintptr(uMapType),
		uintptr(dwhkl))
	return winapi.UINT(r0)
}

func MapVirtualKeyW(uCode winapi.UINT, uMapType winapi.UINT) winapi.UINT {
	r0, _, _ := procMapVirtualKeyW.Call(
		uintptr(uCode),
		uintptr(uMapType))
	return winapi.UINT(r0)
}

const (
	KLF_ACTIVATE      = 0x00000001
	KLF_NOTELLSHELL   = 0x00000080
	KLF_REORDER       = 0x00000008
	KLF_REPLACELANG   = 0x00000010
	KLF_SUBSTITUTE_OK = 0x00000002
	KLF_SETFORPROCESS = 0x00000100
)

func GetKeyboardLayout(idThread winapi.DWORD) winapi.HKL {
	r0, _, _ := procGetKeyboardLayout.Call(uintptr(idThread))
	return winapi.HKL(r0)
}
