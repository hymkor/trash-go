package trash

import (
	"unsafe"

	"golang.org/x/sys/windows"

	"github.com/hymkor/trash-go/internal/constant"
)

var shell32Dll = windows.NewLazySystemDLL("Shell32.dll")

var shFileOperationWProc = shell32Dll.NewProc("SHFileOperationW")

type _ShFileOpStruct struct {
	hwnd                  uintptr
	wFunc                 uintptr
	pFrom                 uintptr
	pTo                   uintptr
	fileOpFlags           uintptr
	fAnyOperationsAborted uintptr
	hNameMappings         uintptr
	lpszProgressTitle     uintptr
}

func throw(filenames ...string) error {
	pFromData := make([]uint16, 0, 256)
	for _, fn := range filenames {
		u, err := windows.UTF16FromString(fn)
		if err != nil {
			return err
		}
		pFromData = append(pFromData, u...)
	}
	pFromData = append(pFromData, 0)

	title := []uint16{0, 0}

	param := &_ShFileOpStruct{
		wFunc:             constant.FO_DELETE,
		pFrom:             uintptr(unsafe.Pointer(&pFromData[0])),
		fileOpFlags:       (constant.FOF_ALLOWUNDO | constant.FOF_NOCONFIRMATION),
		lpszProgressTitle: uintptr(unsafe.Pointer(&title[0])),
	}

	shFileOperationWProc.Call(uintptr(unsafe.Pointer(param)))
	return nil
}
