//go:build windows

package winapi

import (
	"golang.org/x/sys/windows"
)

type ATOM WORD
type BOOL int32
type DWORD uint32
type HANDLE windows.Handle
type HBRUSH HANDLE
type HCURSOR HICON
type HDC HANDLE
type HICON HANDLE
type HINSTANCE HANDLE
type HMENU HANDLE
type HMODULE HINSTANCE
type HWND HANDLE
type LARGE_INTEGER int64
type LONG_PTR uintptr
type LPARAM LONG_PTR
type LPCWSTR *uint16
type LRESULT LONG_PTR
type UINT uint32
type UINT_PTR uintptr
type WORD uint16
type WPARAM UINT_PTR
type LPVOID uintptr
type LONG int32
type HKL HANDLE

const (
	NULL  = 0
	TRUE  = BOOL(1)
	FALSE = BOOL(0)
)

type WindowProc func(HWND, UINT, WPARAM, LPARAM) LRESULT

type WNDPROC uintptr

func LOWORD(dword DWORD) WORD { return WORD(dword) }

func HIWORD(dword DWORD) WORD { return WORD(dword >> 16) }

func GET_X_LPARAM(lParam LPARAM) int32 { return int32(LOWORD(DWORD(lParam))) }

func GET_Y_LPARAM(lParam LPARAM) int32 { return int32(HIWORD(DWORD(lParam))) }

const WHEEL_DELTA = 120

func GET_WHEEL_DELTA_WPARAM(wParam WPARAM) int16 { return int16(HIWORD(DWORD(wParam))) }
