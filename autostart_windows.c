#define UNICODE
#define _UNICODE
#include <windows.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include <objbase.h>
#include <shlobj.h>

uint64_t CreateShortcut(char *shortcutA, char *path, char *args) {
	IShellLink*   pISL;
	IPersistFile* pIPF;
	HRESULT       hr;

	CoInitializeEx(NULL, COINIT_MULTITHREADED);

	// Shortcut filename: convert ANSI to unicode
	WORD shortcutW[MAX_PATH];
	int nChar = MultiByteToWideChar(CP_UTF8, 0, shortcutA, -1, shortcutW, MAX_PATH);

	hr = CoCreateInstance(&CLSID_ShellLink, NULL, CLSCTX_INPROC_SERVER, &IID_IShellLink, (LPVOID*)&pISL);
	if (!SUCCEEDED(hr)) {
		return hr+0x01000000;
	}

	// Convert path to unicode
	WCHAR pathW[MAX_PATH];
	int pathLen = MultiByteToWideChar(CP_UTF8, 0, path, -1, pathW, MAX_PATH);

	// See https://msdn.microsoft.com/en-us/library/windows/desktop/bb774950(v=vs.85).aspx
	hr = pISL->lpVtbl->SetPath(pISL, pathW);
	if (!SUCCEEDED(hr)) {
		return hr+0x02000000;
	}

	// Convert args to unicode
	WCHAR argsW[MAX_PATH];
	int argsLen = MultiByteToWideChar(CP_UTF8, 0, args, -1, argsW, MAX_PATH);

	hr = pISL->lpVtbl->SetArguments(pISL, argsW);
	if (!SUCCEEDED(hr)) {
		return hr+0x03000000;
	}

	// Save the shortcut
	hr = pISL->lpVtbl->QueryInterface(pISL, &IID_IPersistFile, (void**)&pIPF);
	if (!SUCCEEDED(hr)) {
		return hr+0x04000000;
	}

	hr = pIPF->lpVtbl->Save(pIPF, shortcutW, FALSE);
	if (!SUCCEEDED(hr)) {
		return hr+0x05000000;
	}

	pIPF->lpVtbl->Release(pIPF);
	pISL->lpVtbl->Release(pISL);

	return 0x0;
}
