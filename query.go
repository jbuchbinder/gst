package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"
import "github.com/falinux/glib"

type SeekFlag C.GstSeekFlags

const (
	SEEK_FLAG_NONE                = SeekFlag(C.GST_SEEK_FLAG_NONE)
	SEEK_FLAG_FLUSH               = SeekFlag(C.GST_SEEK_FLAG_FLUSH)
	SEEK_FLAG_ACCURATE            = SeekFlag(C.GST_SEEK_FLAG_ACCURATE)
	SEEK_FLAG_KEY_UNIT            = SeekFlag(C.GST_SEEK_FLAG_KEY_UNIT)
	SEEK_FLAG_SEGMENT             = SeekFlag(C.GST_SEEK_FLAG_SEGMENT)
	SEEK_FLAG_TRICKMODE           = SeekFlag(C.GST_SEEK_FLAG_TRICKMODE)
	SEEK_FLAG_SKIP                = SeekFlag(C.GST_SEEK_FLAG_SKIP)
	SEEK_FLAG_SNAP_BEFORE         = SeekFlag(C.GST_SEEK_FLAG_SNAP_BEFORE)
	SEEK_FLAG_SNAP_AFTER          = SeekFlag(C.GST_SEEK_FLAG_SNAP_AFTER)
	SEEK_FLAG_SNAP_NEAREST        = SeekFlag(C.GST_SEEK_FLAG_SNAP_NEAREST)
	SEEK_FLAG_TRICKMODE_KEY_UNITS = SeekFlag(C.GST_SEEK_FLAG_TRICKMODE_KEY_UNITS)
	SEEK_FLAG_TRICKMODE_NO_AUDIO  = SeekFlag(C.GST_SEEK_FLAG_TRICKMODE_NO_AUDIO)
)

func (s *SeekFlag) g() *C.GstSeekFlags {
	return (*C.GstSeekFlags)(s)
}

type Query struct {
	GstObj
}

func (q *Query) g() *C.GstQuery {
	return (*C.GstQuery)(q.GetPtr())
}

func (q *Query) Unref() {
	C.gst_query_unref(q.g())
}

func (q *Query) ParseSeeking() (bool, int64, int64) {
	var s C.gboolean
	var start C.gint64
	var stop C.gint64
	C.gst_query_parse_seeking(q.g(), nil, &s, &start, &stop)

	var seekable bool
	if s == C.TRUE {
		seekable = true
	} else {
		seekable = false
	}
	return seekable, int64(start), int64(stop)
}

func QueryNewSeeking(f Format) *Query {
	ge := C.gst_query_new_seeking(C.GstFormat(f))
	if ge == nil {
		return nil
	}

	q := new(Query)
	q.SetPtr(glib.Pointer(ge))
	return q
}
