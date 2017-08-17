package windows

import (
  "syscall"
  "unsafe"
  "log"
)

type PsAPI struct{ *syscall.LazyDLL }
type procEnumProcesses struct{ *syscall.LazyProc }
type procEnumProcessModules struct{ *syscall.LazyProc }

func GetPsAPI() PsAPI {
  return PsAPI{syscall.NewLazyDLL("psapi.dll")}
}
func (psapi PsAPI) GetEnumProcesses() procEnumProcesses {
  return procEnumProcesses{psapi.NewProc("EnumProcesses")}
}
func (psapi PsAPI) GetEnumProcessModules() procEnumProcessModules {
  return procEnumProcessModules{psapi.NewProc("EnumProcessModules")}
}

func (proc procEnumProcesses) Exec(processIds *[]uint32) bool {
  /*
  pProcessIds [out]
    A pointer to an array that receives the list of process identifiers.
  cb [in]
    The size of the pProcessIds array, in bytes.
  pBytesReturned [out]
    The number of bytes returned in the pProcessIds array.
  */
  var bytesReturned uint32
  *processIds = (*processIds)[:cap(*processIds)]
	ret, _, _ := proc.Call(
		uintptr(unsafe.Pointer(&(*processIds)[0])),
		uintptr(4*cap(*processIds)),
		uintptr(unsafe.Pointer(&bytesReturned)),
  )
  *processIds = (*processIds)[:bytesReturned/4]
	return ret != 0
}

func (proc procEnumProcessModules) Exec(handle *uint32) bool {
    modules := make([]syscall.Handle, 2049)
    var needed uint32
    ret, _, _ := proc.Call(
        uintptr(unsafe.Pointer(handle)),
        uintptr(unsafe.Pointer(&modules)),
        uintptr(2048),
        uintptr(unsafe.Pointer(&needed)),
    )
    log.Println(needed)
    for i := uint32(0); i < needed; i++ {
        log.Println(modules[i])
    }
    return ret != 0
}
