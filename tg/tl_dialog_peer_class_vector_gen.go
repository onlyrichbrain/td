// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// DialogPeerClassVector is a box for Vector<DialogPeer>
type DialogPeerClassVector struct {
	// Elements of Vector<DialogPeer>
	Elems []DialogPeerClass `schemaname:"Elems"`
}

// DialogPeerClassVectorTypeID is TL type id of DialogPeerClassVector.
const DialogPeerClassVectorTypeID = bin.TypeVector

func (vec *DialogPeerClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *DialogPeerClassVector) String() string {
	if vec == nil {
		return "DialogPeerClassVector(nil)"
	}
	type Alias DialogPeerClassVector
	return fmt.Sprintf("DialogPeerClassVector%+v", Alias(*vec))
}

// FillFrom fills DialogPeerClassVector from given interface.
func (vec *DialogPeerClassVector) FillFrom(from interface {
	GetElems() (value []DialogPeerClass)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (vec *DialogPeerClassVector) TypeID() uint32 {
	return DialogPeerClassVectorTypeID
}

// SchemaName returns MTProto type name.
func (vec *DialogPeerClassVector) SchemaName() string {
	return ""
}

// Encode implements bin.Encoder.
func (vec *DialogPeerClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<DialogPeer> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<DialogPeer>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<DialogPeer>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *DialogPeerClassVector) GetElems() (value []DialogPeerClass) {
	return vec.Elems
}

// MapElems returns field Elems wrapped in DialogPeerClassSlice helper.
func (vec *DialogPeerClassVector) MapElems() (value DialogPeerClassSlice) {
	return DialogPeerClassSlice(vec.Elems)
}

// Decode implements bin.Decoder.
func (vec *DialogPeerClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<DialogPeer> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<DialogPeer>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDialogPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<DialogPeer>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for DialogPeerClassVector.
var (
	_ bin.Encoder = &DialogPeerClassVector{}
	_ bin.Decoder = &DialogPeerClassVector{}
)
