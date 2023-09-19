package main

import (
	"runtime"
	"unsafe"

	"github.com/o5h/winapi"
	"github.com/o5h/winapi/kernel32"
	"github.com/o5h/winapi/user32"
	"golang.org/x/sys/windows"
)

func main() {
	runtime.LockOSThread()
	hwnd := CreateWindow("Example 1", WndProc)
	user32.ShowWindow(hwnd, user32.SW_SHOW)
	user32.UpdateWindow(hwnd)
	user32.SetFocus(hwnd)
	user32.SetWindowPos(hwnd, user32.HWND_TOP, 0, 0, 1024, 768, user32.SWP_SHOWWINDOW)

	//main loop
	var m user32.Msg
	for {
		ok, _ := user32.GetMessageW(&m, 0, 0, 0)
		if ok == winapi.FALSE {
			break
		} else {
			user32.TranslateMessage(&m)
			user32.DispatchMessageW(&m)
		}
	}

}

func CreateWindow(title string, wndProc winapi.WindowProc) (hwnd winapi.HWND) {
	wndproc := winapi.WNDPROC(windows.NewCallback(wndProc))
	mh, _ := kernel32.GetModuleHandle(nil)
	myicon, _ := user32.LoadIconW(0, user32.IDI_APPLICATION)
	mycursor, _ := user32.LoadCursorW(0, user32.IDC_ARROW)

	var wc user32.WNDCLASSEX
	wc.Size = uint32(unsafe.Sizeof(wc))
	wc.WndProc = wndproc
	wc.Instance = winapi.HINSTANCE(mh)
	wc.Icon = myicon
	wc.Cursor = mycursor
	wc.Background = user32.COLOR_BTNFACE + 1
	wc.MenuName = nil
	wcname, _ := windows.UTF16PtrFromString("_WindowClass")
	wc.ClassName = wcname
	wc.IconSm = myicon
	user32.RegisterClassExW(&wc)

	windowTitle, _ := windows.UTF16PtrFromString(title)
	hwnd, _ = user32.CreateWindowExW(
		0,
		wcname,
		windowTitle,
		user32.WS_POPUP|user32.WS_CLIPSIBLINGS|user32.WS_CLIPCHILDREN|user32.WS_OVERLAPPEDWINDOW,
		user32.CW_USEDEFAULT,
		user32.CW_USEDEFAULT,
		user32.CW_USEDEFAULT,
		user32.CW_USEDEFAULT,
		winapi.HWND(0), winapi.HMENU(0), winapi.HINSTANCE(mh), winapi.LPVOID(0))
	return
}

func WndProc(hwnd winapi.HWND, msg winapi.UINT, wParam winapi.WPARAM, lparam winapi.LPARAM) (rc winapi.LRESULT) {
	switch msg {
	case user32.WM_CREATE:
		rc = user32.DefWindowProcW(hwnd, msg, wParam, lparam)
	case user32.WM_SIZE:
	case user32.WM_CLOSE:
		user32.DestroyWindow(hwnd)
	case user32.WM_DESTROY:
		user32.PostQuitMessage(0)
	case user32.WM_KEYDOWN:
	case user32.WM_KEYUP:
	default:
		rc = user32.DefWindowProcW(hwnd, msg, wParam, lparam)
	}
	return
}
