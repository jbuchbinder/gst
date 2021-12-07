package gst

/*
#include <stdlib.h>
#include <gst/gst.h>

extern GstPadProbeReturn goPadProbeFunc (GstPad * pad, GstPadProbeInfo * info, gpointer user_data);
extern void              goGDestroyNotifyFuncNoRun    (gpointer user_data);

GstPadProbeReturn
cgoPadProbeFunc (GstPad * pad, GstPadProbeInfo * info, gpointer user_data) {
	return goPadProbeFunc(pad, info, user_data);
}

void
cgoGDestroyNotifyFuncNoRun (gpointer user_data) {
	goGDestroyNotifyFuncNoRun(user_data);
}


#cgo pkg-config: gstreamer-1.0
#cgo LDFLAGS: -lm
#cgo CFLAGS: -Wno-deprecated-declarations
*/
import "C"

import (
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
	"github.com/ziutek/glib"
)

type PadProbeInfo struct {
	ptr *C.GstPadProbeInfo
}

func (p *PadProbeInfo) ID() uint32 { return uint32(p.ptr.id) }

func (p *PadProbeInfo) Type() PadProbeType { return PadProbeType(p.ptr._type) }

func (p *PadProbeInfo) Offset() uint64 { return uint64(p.ptr.offset) }

func (p *PadProbeInfo) Size() uint64 { return uint64(p.ptr.size) }

func (p *PadProbeInfo) GetBuffer() *Buffer {
	buf := C.gst_pad_probe_info_get_buffer(p.ptr)
	if buf == nil {
		return nil
	}
	return nil
}

func (p *PadProbeInfo) GetEvent() *Event {
	ev := C.gst_pad_probe_info_get_event(p.ptr)
	if ev == nil {
		return nil
	}

	return nil
	// return &Event(&ev)
}

func (p *PadProbeInfo) GetData() unsafe.Pointer {
	return unsafe.Pointer(p.ptr.data)
}

type PadLinkReturn C.GstPadLinkReturn

const (
	PAD_LINK_OK              = PadLinkReturn(C.GST_PAD_LINK_OK)
	PAD_LINK_WRONG_HIERARCHY = PadLinkReturn(C.GST_PAD_LINK_WRONG_HIERARCHY)
	PAD_LINK_WAS_LINKED      = PadLinkReturn(C.GST_PAD_LINK_WAS_LINKED)
	PAD_LINK_WRONG_DIRECTION = PadLinkReturn(C.GST_PAD_LINK_WRONG_DIRECTION)
	PAD_LINK_NOFORMAT        = PadLinkReturn(C.GST_PAD_LINK_NOFORMAT)
	PAD_LINK_NOSCHED         = PadLinkReturn(C.GST_PAD_LINK_NOSCHED)
	PAD_LINK_REFUSED         = PadLinkReturn(C.GST_PAD_LINK_REFUSED)
)

// PadProbeReturn casts GstPadProbeReturn
type PadProbeReturn int

// Type castings of ProbeReturns
const (
	PadProbeDrop      PadProbeReturn = C.GST_PAD_PROBE_DROP    // (0) – drop data in data probes. For push mode this means that the data item is not sent downstream. For pull mode, it means that the data item is not passed upstream. In both cases, no other probes are called for this item and GST_FLOW_OK or TRUE is returned to the caller.
	PadProbeOK        PadProbeReturn = C.GST_PAD_PROBE_OK      // (1) – normal probe return value. This leaves the probe in place, and defers decisions about dropping or passing data to other probes, if any. If there are no other probes, the default behaviour for the probe type applies ('block' for blocking probes, and 'pass' for non-blocking probes).
	PadProbeRemove    PadProbeReturn = C.GST_PAD_PROBE_REMOVE  // (2) – remove this probe.
	PadProbePass      PadProbeReturn = C.GST_PAD_PROBE_PASS    // (3) – pass the data item in the block probe and block on the next item.
	PadProbeUnhandled PadProbeReturn = C.GST_PAD_PROBE_HANDLED // (4) – Data has been handled in the probe and will not be forwarded further. For events and buffers this is the same behaviour as GST_PAD_PROBE_DROP (except that in this case you need to unref the buffer or event yourself). For queries it will also return TRUE to the caller. The probe can also modify the GstFlowReturn value by using the GST_PAD_PROBE_INFO_FLOW_RETURN() accessor. Note that the resulting query must contain valid entries. Since: 1.6
)

// PadProbeType casts GstPadProbeType
type PadProbeType int

// Type castings of PadProbeTypes
const (
	PadProbeTypeInvalid         PadProbeType = C.GST_PAD_PROBE_TYPE_INVALID          // (0) – invalid probe type
	PadProbeTypeIdle            PadProbeType = C.GST_PAD_PROBE_TYPE_IDLE             // (1) – probe idle pads and block while the callback is called
	PadProbeTypeBlock           PadProbeType = C.GST_PAD_PROBE_TYPE_BLOCK            // (2) – probe and block pads
	PadProbeTypeBuffer          PadProbeType = C.GST_PAD_PROBE_TYPE_BUFFER           // (16) – probe buffers
	PadProbeTypeBufferList      PadProbeType = C.GST_PAD_PROBE_TYPE_BUFFER_LIST      // (32) – probe buffer lists
	PadProbeTypeEventDownstream PadProbeType = C.GST_PAD_PROBE_TYPE_EVENT_DOWNSTREAM // (64) – probe downstream events
	PadProbeTypeEventUpstream   PadProbeType = C.GST_PAD_PROBE_TYPE_EVENT_UPSTREAM   // (128) – probe upstream events
	PadProbeTypeEventFlush      PadProbeType = C.GST_PAD_PROBE_TYPE_EVENT_FLUSH      // (256) – probe flush events. This probe has to be explicitly enabled and is not included in the @GST_PAD_PROBE_TYPE_EVENT_DOWNSTREAM or @GST_PAD_PROBE_TYPE_EVENT_UPSTREAM probe types.
	PadProbeTypeQueryDownstream PadProbeType = C.GST_PAD_PROBE_TYPE_QUERY_DOWNSTREAM // (512) – probe downstream queries
	PadProbeTypeQueryUpstream   PadProbeType = C.GST_PAD_PROBE_TYPE_QUERY_UPSTREAM   // (1024) – probe upstream queries
	PadProbeTypePush            PadProbeType = C.GST_PAD_PROBE_TYPE_PUSH             // (4096) – probe push
	PadProbeTypePull            PadProbeType = C.GST_PAD_PROBE_TYPE_PULL             // (8192) – probe pull
	PadProbeTypeBlocking        PadProbeType = C.GST_PAD_PROBE_TYPE_BLOCKING         // (3) – probe and block at the next opportunity, at data flow or when idle
	PadProbeTypeDataDownstream  PadProbeType = C.GST_PAD_PROBE_TYPE_DATA_DOWNSTREAM  // (112) – probe downstream data (buffers, buffer lists, and events)
	PadProbeTypeDataUpstream    PadProbeType = C.GST_PAD_PROBE_TYPE_DATA_UPSTREAM    // (128) – probe upstream data (events)
	PadProbeTypeDataBoth        PadProbeType = C.GST_PAD_PROBE_TYPE_DATA_BOTH        // (240) – probe upstream and downstream data (buffers, buffer lists, and events)
	PadProbeTypeBlockDownstream PadProbeType = C.GST_PAD_PROBE_TYPE_BLOCK_DOWNSTREAM // (114) – probe and block downstream data (buffers, buffer lists, and events)
	PadProbeTypeBlockUpstream   PadProbeType = C.GST_PAD_PROBE_TYPE_BLOCK_UPSTREAM   // (130) – probe and block upstream data (events)
	PadProbeTypeEventBoth       PadProbeType = C.GST_PAD_PROBE_TYPE_EVENT_BOTH       // (192) – probe upstream and downstream events
	PadProbeTypeQueryBoth       PadProbeType = C.GST_PAD_PROBE_TYPE_QUERY_BOTH       // (1536) – probe upstream and downstream queries
	PadProbeTypeAllBoth         PadProbeType = C.GST_PAD_PROBE_TYPE_ALL_BOTH         // (1776) – probe upstream events and queries and downstream buffers, buffer lists, events and queries
	PadProbeTypeScheduling      PadProbeType = C.GST_PAD_PROBE_TYPE_SCHEDULING       // (12288) – probe push and pull
)

func (p PadLinkReturn) String() string {
	switch p {
	case PAD_LINK_OK:
		return "PAD_LINK_OK"
	case PAD_LINK_WRONG_HIERARCHY:
		return "PAD_LINK_WRONG_HIERARCHY"
	case PAD_LINK_WAS_LINKED:
		return "PAD_LINK_WAS_LINKED"
	case PAD_LINK_WRONG_DIRECTION:
		return "PAD_LINK_WRONG_DIRECTION"
	case PAD_LINK_NOFORMAT:
		return "PAD_LINK_NOFORMAT"
	case PAD_LINK_NOSCHED:
		return "PAD_LINK_NOSCHED"
	case PAD_LINK_REFUSED:
		return "PAD_LINK_REFUSED"
	}
	panic("Wrong value of PadLinkReturn variable")
}

type PadDirection C.GstPadDirection

const (
	PAD_UNKNOWN = PadDirection(C.GST_PAD_UNKNOWN)
	PAD_SRC     = PadDirection(C.GST_PAD_SRC)
	PAD_SINK    = PadDirection(C.GST_PAD_SINK)
)

func (p PadDirection) g() C.GstPadDirection {
	return C.GstPadDirection(p)
}

func (p PadDirection) String() string {
	switch p {
	case PAD_UNKNOWN:
		return "PAD_UNKNOWN"
	case PAD_SRC:
		return "PAD_SRC"
	case PAD_SINK:
		return "PAD_SINK"
	}
	panic("Wrong value of PadDirection variable")
}

type Pad struct {
	GstObj
}

func (p *Pad) g() *C.GstPad {
	return (*C.GstPad)(p.GetPtr())
}

func (p *Pad) AsPad() *Pad {
	return p
}

func (p *Pad) CanLink(sink_pad *Pad) bool {
	return C.gst_pad_can_link(p.g(), sink_pad.g()) != 0
}

func (p *Pad) Link(sink_pad *Pad) PadLinkReturn {
	return PadLinkReturn(C.gst_pad_link(p.g(), sink_pad.g()))
}

// Unlink() is a wrapper around gst_pad_unlink().
func (p *Pad) Unlink(sink_pad *Pad) bool {
	c := C.gst_pad_unlink(p.g(), sink_pad.g())
	return c != 0
}

func (p *Pad) SetActive(active bool) bool {
	var a C.gboolean
	if active {
		a = C.gboolean(1)
	}
	c := C.gst_pad_set_active(p.g(), a)
	return c != 0
}

func (p *Pad) GetCurrentCaps() *Caps {
	// unref the caps when you're done
	return (*Caps)(C.gst_pad_get_current_caps((*C.GstPad)(p.GetPtr())))
}

type PadProbeCallback func(*Pad, *PadProbeInfo, unsafe.Pointer) PadProbeReturn

type UserData struct {
	cb  unsafe.Pointer // callback
	ptr unsafe.Pointer
}

func (p *Pad) AddProbe(mask PadProbeType, f PadProbeCallback, ptr unsafe.Pointer) uint64 {
	user := new(UserData)
	user.cb = gopointer.Save(f)
	user.ptr = ptr
	ret := C.gst_pad_add_probe(
		p.g(),
		C.GstPadProbeType(mask),
		C.GstPadProbeCallback(C.cgoPadProbeFunc),
		(C.gpointer)(unsafe.Pointer(user)),
		// C.GDestroyNotify(C.cgoGDestroyNotifyFuncNoRun),
		nil,
	)

	return uint64(ret)
}

func (p *Pad) RemoveProbe(id uint64) {
	C.gst_pad_remove_probe(p.g(), C.gulong(id))
}

func (p *Pad) SendEvent(event *Event) bool {
	return C.gst_pad_send_event(p.g(), (*C.GstEvent)(event.GstEvent)) != 0
}

type GhostPad struct {
	Pad
}

func (p *GhostPad) g() *C.GstGhostPad {
	return (*C.GstGhostPad)(p.GetPtr())
}

func (p *GhostPad) AsGhostPad() *GhostPad {
	return p
}

func (p *GhostPad) SetTarget(new_target *Pad) bool {
	return C.gst_ghost_pad_set_target(p.g(), new_target.g()) != 0
}

func (p *GhostPad) GetTarget() *Pad {
	r := new(Pad)
	r.SetPtr(glib.Pointer(C.gst_ghost_pad_get_target(p.g())))
	return r
}

func (p *GhostPad) Construct() bool {
	return C.gst_ghost_pad_construct(p.g()) != 0
}

func NewGhostPad(name string, target *Pad) *GhostPad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(GhostPad)
	p.SetPtr(glib.Pointer(C.gst_ghost_pad_new(s, target.g())))
	return p
}

func NewGhostPadNoTarget(name string, dir PadDirection) *GhostPad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(GhostPad)
	p.SetPtr(glib.Pointer(C.gst_ghost_pad_new_no_target(s, dir.g())))
	return p
}
