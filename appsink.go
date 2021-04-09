package gst

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-app-1.0
#cgo LDFLAGS: -lgstapp-1.0
#include <stdlib.h>
#include <string.h>
#include <gst/gst.h>
#include <gst/app/gstappsink.h>
*/
import "C"
import "unsafe"

type AppSink struct {
	*Element
}

func NewAppSink(name string) *AppSink {
	return &AppSink{ElementFactoryMake("appsink", name)}
}

func (a *AppSink) g() *C.GstAppSink {
	return (*C.GstAppSink)(a.GetPtr())
}

func (a *AppSink) SetCaps(caps *Caps) {
	p := unsafe.Pointer(caps) // HACK
	C.gst_app_sink_set_caps(a.g(), (*C.GstCaps)(p))
}

func (a *AppSink) PullSample() *Sample {
	return (*Sample)(C.gst_app_sink_pull_sample(a.g()))
}

func (a *AppSink) IsEOS() bool {
	return C.gst_app_sink_is_eos(a.g()) != 0
}
