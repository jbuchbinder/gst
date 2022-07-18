package gst

/*
#include <gst/gst.h>
#include <gio/gio.h>
*/
import "C"

type Socket C.GSocket

func (s *Socket) g() *C.GSocket {
	return (*C.GSocket)(s)
}
