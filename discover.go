package gst

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-pbutils-1.0
#include <string.h>
#include <stdlib.h>
#include <gst/gst.h>
#include <gst/pbutils/pbutils.h>
*/
import "C"

import (
	"time"
	"unsafe"

	"github.com/falinux/glib"
)

type Discoverer struct {
	GstObj
}

func (d *Discoverer) g() *C.GstDiscoverer {
	return (*C.GstDiscoverer)(d.GetPtr())
}

func NewDiscoverer(timeout int64) (*Discoverer, error) {
	var Cerr *C.GError
	discover := C.gst_discoverer_new(C.GstClockTime(timeout*C.GST_SECOND), &Cerr)
	if Cerr != nil {
		err := *(*glib.Error)(unsafe.Pointer(Cerr))
		C.g_error_free(Cerr)
		return nil, &err
	}
	d := new(Discoverer)
	d.SetPtr(glib.Pointer(discover))

	return d, nil
}

type DiscovererInfo C.GstDiscovererInfo

func (i *DiscovererInfo) g() *C.GstDiscovererInfo {
	return (*C.GstDiscovererInfo)(i)
}

func (d *Discoverer) DiscoverUri(uri string) (*DiscovererInfo, error) {
	u := (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(u))

	var Cerr *C.GError
	info := (*DiscovererInfo)(C.gst_discoverer_discover_uri(d.g(), u, &Cerr))
	if Cerr != nil {
		err := *(*glib.Error)(unsafe.Pointer(Cerr))
		C.g_error_free(Cerr)
		return nil, &err
	}

	return info, nil
}

type DiscovererResult C.GstDiscovererResult

const (
	DISCOVERER_OK              = DiscovererResult(C.GST_DISCOVERER_OK)
	DISCOVERER_URI_INVALID     = DiscovererResult(C.GST_DISCOVERER_URI_INVALID)
	DISCOVERER_ERROR           = DiscovererResult(C.GST_DISCOVERER_ERROR)
	DISCOVERER_TIMEOUT         = DiscovererResult(C.GST_DISCOVERER_TIMEOUT)
	DISCOVERER_BUSY            = DiscovererResult(C.GST_DISCOVERER_BUSY)
	DISCOVERER_MISSING_PLUGINS = DiscovererResult(C.GST_DISCOVERER_MISSING_PLUGINS)
)

func (s DiscovererResult) String() string {
	switch s {
	case DISCOVERER_OK:
		return "DISCOVERER_OK"
	case DISCOVERER_URI_INVALID:
		return "DISCOVERER_URI_INVALID"
	case DISCOVERER_ERROR:
		return "DISCOVERER_ERROR"
	case DISCOVERER_TIMEOUT:
		return "DISCOVERER_TIMEOUT"
	case DISCOVERER_BUSY:
		return "DISCOVERER_BUSY"
	case DISCOVERER_MISSING_PLUGINS:
		return "DISCOVERER_MISSING_PLUGINS"
	}
	panic("Unknown state")
}

func (i *DiscovererInfo) GetResult() DiscovererResult {
	return DiscovererResult(C.gst_discoverer_info_get_result(i.g()))
}

func (i *DiscovererInfo) GetDuration() time.Duration {
	return time.Duration(C.gst_discoverer_info_get_duration(i.g()))
}

func GetFilenameToUri(file string) (string, error) {
	s := (*C.gchar)(C.CString(file))
	defer C.free(unsafe.Pointer(s))

	var Cerr *C.GError
	str := C.gst_filename_to_uri(s, &Cerr)
	if Cerr != nil {
		err := *(*glib.Error)(unsafe.Pointer(Cerr))
		C.g_error_free(Cerr)
		return "", &err
	}

	uri := C.GoString((*C.char)(str))
	defer C.g_free(C.gpointer(str))
	return uri, nil
}
