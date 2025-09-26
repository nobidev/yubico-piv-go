package ykpiv

/*
#cgo LDFLAGS: -lykpiv
#include "ykpiv/ykpiv.h"
*/
import "C"

type (
	State = *C.ykpiv_state
)

var _ error = ReturnCode(0)

type ReturnCode C.ykpiv_rc

func (rc ReturnCode) asC() C.ykpiv_rc {
	return C.ykpiv_rc(rc)
}

func (rc ReturnCode) isOK() bool {
	return rc.asC() == C.YKPIV_OK
}

func (rc ReturnCode) Error() string {
	var s *C.char = C.ykpiv_strerror(rc.asC())
	return C.GoString(s)
}

func cInit(state *State, verbose int) ReturnCode {
	var rc C.ykpiv_rc = C.ykpiv_init(state, C.int(verbose))
	return ReturnCode(rc)
}

func cDone(state State) ReturnCode {
	var rc C.ykpiv_rc = C.ykpiv_done(state)
	return ReturnCode(rc)
}

func cConnect(state State, reader string, encrypted bool) ReturnCode {
	var rc C.ykpiv_rc = C.ykpiv_connect_ex(state, C.CString(reader), C.bool(encrypted))
	return ReturnCode(rc)
}
