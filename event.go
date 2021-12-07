package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"
import "unsafe"

type GstEventStruct C.GstEvent

type Event struct {
	GstEvent *GstEventStruct
}

func Eos() *Event {
	event := new(Event)
	event.GstEvent = (*GstEventStruct)(C.gst_event_new_eos())
	return event
}

// EventType casts GstEventType
type EventType int

// Type castrings of EventTypes
const (
	EventUnknown                EventType = C.GST_EVENT_UNKNOWN           // (0) – unknown event
	EventFlushStart             EventType = C.GST_EVENT_FLUSH_START       // (2564)
	EventFlushStop              EventType = C.GST_EVENT_FLUSH_STOP        // (5127)
	EventStreamStart            EventType = C.GST_EVENT_STREAM_START      // (10254)
	EventCaps                   EventType = C.GST_EVENT_CAPS              // (12814)
	EventSegment                EventType = C.GST_EVENT_SEGMENT           // (17934)
	EventStreamCollection       EventType = C.GST_EVENT_STREAM_COLLECTION // (19230)
	EventTag                    EventType = C.GST_EVENT_TAG               // (20510)
	EventBuffersize             EventType = C.GST_EVENT_BUFFERSIZE        // (23054)
	EventSinkMessage            EventType = C.GST_EVENT_SINK_MESSAGE      // (25630)
	EventStreamGroupDone        EventType = C.GST_EVENT_STREAM_GROUP_DONE // (26894)
	EventEos                    EventType = C.GST_EVENT_EOS               // (28174)
	EventToc                    EventType = C.GST_EVENT_TOC
	EventProtection             EventType = C.GST_EVENT_PROTECTION
	EventSegmentDone            EventType = C.GST_EVENT_SEGMENT_DONE
	EventGap                    EventType = C.GST_EVENT_GAP
	EventQos                    EventType = C.GST_EVENT_QOS
	EventSeek                   EventType = C.GST_EVENT_SEEK
	EventNavigation             EventType = C.GST_EVENT_NAVIGATION
	EventLatency                EventType = C.GST_EVENT_LATENCY
	EventStep                   EventType = C.GST_EVENT_STEP
	EventReconfigure            EventType = C.GST_EVENT_RECONFIGURE
	EvnetTocSelect              EventType = C.GST_EVENT_TOC_SELECT
	EventSelectStreams          EventType = C.GST_EVENT_SELECT_STREAMS
	EventCustomUpstream         EventType = C.GST_EVENT_CUSTOM_UPSTREAM
	EventCustomDownstream       EventType = C.GST_EVENT_CUSTOM_DOWNSTREAM
	EventCustomDownstreamOob    EventType = C.GST_EVENT_CUSTOM_DOWNSTREAM_OOB
	EventCustomDownstreamSticky EventType = C.GST_EVENT_CUSTOM_DOWNSTREAM_STICKY
	EventCustomBoth             EventType = C.GST_EVENT_CUSTOM_BOTH
	EventCustomBothOob          EventType = C.GST_EVENT_CUSTOM_BOTH_OOB
)

func GetEventType(data unsafe.Pointer) EventType {
	return EventType((*C.GstEvent)(data)._type)
}
