package gst

/*
#include <gst/gst.h>
#include <gio/gio.h>
*/
import "C"

type Socket struct {
	GstObj
}

func (s *Socket) g() *C.GSocket {
	return (*C.GSocket)(s.GetPtr())
}

func (s *Socket) AsSocket() *Socket {
	return s
}
