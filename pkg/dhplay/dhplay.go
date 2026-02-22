package dhplay

/*
#cgo linux CFLAGS: -I../../vendor/linux/dhplay
#cgo linux LDFLAGS: -L../../vendor/linux/dhplay -ldhplay -Wl,--allow-shlib-undefined
#cgo windows CFLAGS: -I../../vendor/windows/dhplay
#cgo windows LDFLAGS: -L../../vendor/windows/dhplay -lplay

#include <stdbool.h>
#include "dhplay.h"
#include <stdlib.h>

// Forward declaration if not in header
extern BOOL CALLMETHOD PLAY_StartDataRecordEx(LONG nPort, char *sFileName, int idataType, void* fRecordDataCBFun, void* fListener, void* pUserData);
*/
import "C"
import "unsafe"

const (
	STREAME_REALTIME = 0
	STREAME_FILE     = 1

	DATA_RECORD_ORIGINAL          = 0
	DATA_RECORD_AVI               = 1
	DATA_RECORD_ASF               = 2
	DATA_RECORD_ORIGINAL_SEGMENT  = 3
	DATA_RECORD_RESIZE_AVI        = 4
	DATA_RECORD_MP4               = 5
	DATA_RECORD_RESIZE_MP4        = 6
	DATA_RECORD_MP4_NOSEEK        = 7
	DATA_RECORD_RESIZE_MP4_NOSEEK = 8
	DATA_RECORD_TS                = 9
	DATA_RECORD_PS                = 10
	DATA_RECORD_RESIZE_DAV        = 11
	DATA_RECORD_DAV               = 12
	DATA_RECORD_AAC               = 13
	DATA_RECORD_WAV               = 14
)

func SetStreamOpenMode(port int, mode int) bool {
	res := C.PLAY_SetStreamOpenMode(C.LONG(port), C.DWORD(mode))
	return res != 0
}

func OpenStream(port int, bufPoolSize int) bool {
	res := C.PLAY_OpenStream(C.LONG(port), nil, 0, C.DWORD(bufPoolSize))
	return res != 0
}

func Play(port int) bool {
	res := C.PLAY_Play(C.LONG(port), nil)
	return res != 0
}

func StartDataRecordEx(port int, fileName string, dataType int) bool {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	res := C.PLAY_StartDataRecordEx(C.LONG(port), cFileName, C.int(dataType), nil, nil, nil)
	return res != 0
}

func InputData(port int, data []byte) bool {
	if len(data) == 0 {
		return true
	}
	cData := (*C.BYTE)(unsafe.Pointer(&data[0]))
	res := C.PLAY_InputData(C.LONG(port), cData, C.DWORD(len(data)))
	return res != 0
}

func StopDataRecord(port int) bool {
	res := C.PLAY_StopDataRecord(C.LONG(port))
	return res != 0
}

func Stop(port int) bool {
	res := C.PLAY_Stop(C.LONG(port))
	return res != 0
}

func CloseStream(port int) bool {
	res := C.PLAY_CloseStream(C.LONG(port))
	return res != 0
}
