package gst

/*
#include <stdlib.h>
#include <gst/gstbuffer.h>

void _golang_gst_set_dts( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_DTS(buffer) = value;
}

void _golang_gst_set_pts( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_PTS(buffer) = value;
}

void _golang_gst_set_duration( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_DURATION(buffer) = value;
}

void _golang_gst_set_offset( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_OFFSET(buffer) = value;
}

void _golang_gst_set_offset_end( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_OFFSET_END(buffer) = value;
}

*/
import "C"

import (
	"unsafe"
)

type GstBufferStruct C.GstBuffer

type Buffer struct {
	GstBuffer *GstBufferStruct
}

func NewBuffer() *Buffer {
	buffer := new(Buffer)
	buffer.GstBuffer = (*GstBufferStruct)(C.gst_buffer_new())
	return buffer
}

func NewBufferAllocate(size uint) *Buffer {
	buffer := new(Buffer)
	buffer.GstBuffer = (*GstBufferStruct)(C.gst_buffer_new_allocate(nil, C.gsize(size), nil))
	return buffer
}

func (b *Buffer) g() *C.GstBuffer {
	return (*C.GstBuffer)(b.GstBuffer)
}

func (b *Buffer) GetSize() uint {
	return (uint)(C.gst_buffer_get_size((*C.GstBuffer)(b.GstBuffer)))
}

func (b *Buffer) AppendMemory(memory *Memory) {
	C.gst_buffer_append_memory((*C.GstBuffer)(b.GstBuffer), (*C.GstMemory)(memory))
}

func (b *Buffer) MemSet(offset uint, val byte, size uint) int {
	return (int)(C.gst_buffer_memset((*C.GstBuffer)(b.GstBuffer), C.gsize(offset), C.guint8(val), C.gsize(size)))
}

func (b *Buffer) FillWithGoSlice(data []byte) int {
	dataLength := uint(len(data))
	return (int)(C.gst_buffer_fill((*C.GstBuffer)(b.GstBuffer), C.gsize(0), (C.gconstpointer)(C.CBytes(data)), C.gsize(dataLength)))
}

func (b *Buffer) Fill(offset uint, src unsafe.Pointer, size uint) int {
	return (int)(C.gst_buffer_fill((*C.GstBuffer)(b.GstBuffer), C.gsize(offset), C.gconstpointer(src), C.gsize(size)))
}

func (b *Buffer) SetDTS(value uint64) {
	C._golang_gst_set_dts((*C.GstBuffer)(b.GstBuffer), C.guint64(value))
}

func (b *Buffer) SetPTS(value uint64) {
	C._golang_gst_set_pts((*C.GstBuffer)(b.GstBuffer), C.guint64(value))
}

func (b *Buffer) SetDuration(value uint64) {
	C._golang_gst_set_duration((*C.GstBuffer)(b.GstBuffer), C.guint64(value))
}

func (b *Buffer) SetOffset(value uint64) {
	C._golang_gst_set_offset((*C.GstBuffer)(b.GstBuffer), C.guint64(value))
}

func (b *Buffer) SetOffsetEnd(value uint64) {
	C._golang_gst_set_offset_end((*C.GstBuffer)(b.GstBuffer), C.guint64(value))
}

func (b *Buffer) Unref() {
	C.gst_buffer_unref((*C.GstBuffer)(b.GstBuffer))
}

func (b *Buffer) Copy() *Buffer {
	buffer := new(Buffer)
	buffer.GstBuffer = (*GstBufferStruct)(C.gst_buffer_copy(b.g()))
	return buffer
}

type MapFlags int

// Type casting of the map flag types
const (
	MapRead     MapFlags = C.GST_MAP_READ      //  (1) – map for read access
	MapWrite    MapFlags = C.GST_MAP_WRITE     // (2) - map for write access
	MapFlagLast MapFlags = C.GST_MAP_FLAG_LAST // (65536) – first flag that can be used for custom purposes
)

// MapInfo is a go representation of a GstMapInfo.
type MapInfo struct {
	ptr *C.GstMapInfo
}

func wrapMapInfo(mapInfo *C.GstMapInfo) *MapInfo { return &MapInfo{ptr: mapInfo} }

func (m *MapInfo) Data() unsafe.Pointer {
	return unsafe.Pointer(m.ptr.data)
}

// Size returrns the size of this map.
func (m *MapInfo) Size() int64 {
	return int64(m.ptr.size)
}

func (m *MapInfo) Bytes() []byte {
	return C.GoBytes(m.Data(), (C.int)(m.Size()))
}

// func (m *MapInfo)Unmap() {
// 	C.gst_buffer_unmap()
// }

func (b *Buffer) BufferMap(flags MapFlags) *MapInfo {
	var mapinfo C.GstMapInfo

	C.gst_buffer_map(
		(*C.GstBuffer)(b.g()),
		(*C.GstMapInfo)(&mapinfo),
		C.GstMapFlags(flags),
	)

	return wrapMapInfo((*C.GstMapInfo)(&mapinfo))
}

func (b *Buffer) Bytes() []byte {
	mapinfo := b.BufferMap(MapRead)
	if mapinfo == nil {
		return nil
	}
	// defer b.Unmap()
	return mapinfo.Bytes()
}
