package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
	"github.com/ziutek/glib"
)

//export goPadProbeFunc
func goPadProbeFunc(gstPad *C.GstPad, info *C.GstPadProbeInfo, userData C.gpointer) C.GstPadProbeReturn {
	ud := (*UserData)(unsafe.Pointer(userData))
	cbIface := gopointer.Restore(ud.cb)
	cbFunc := cbIface.(PadProbeCallback)
	pad := new(Pad)
	pad.SetPtr(glib.Pointer(gstPad))
	return C.GstPadProbeReturn(cbFunc(pad, &PadProbeInfo{info}, ud.ptr))
}

//export goGDestroyNotifyFuncNoRun
func goGDestroyNotifyFuncNoRun(ptr C.gpointer) {
	gopointer.Unref(unsafe.Pointer(ptr))
}
