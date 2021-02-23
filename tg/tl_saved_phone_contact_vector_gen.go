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

// SavedPhoneContactVector is a box for Vector<SavedContact>
type SavedPhoneContactVector struct {
	// Elements of Vector<SavedContact>
	Elems []SavedPhoneContact `schemaname:"Elems"`
}

// SavedPhoneContactVectorTypeID is TL type id of SavedPhoneContactVector.
const SavedPhoneContactVectorTypeID = bin.TypeVector

func (vec *SavedPhoneContactVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *SavedPhoneContactVector) String() string {
	if vec == nil {
		return "SavedPhoneContactVector(nil)"
	}
	type Alias SavedPhoneContactVector
	return fmt.Sprintf("SavedPhoneContactVector%+v", Alias(*vec))
}

// FillFrom fills SavedPhoneContactVector from given interface.
func (vec *SavedPhoneContactVector) FillFrom(from interface {
	GetElems() (value []SavedPhoneContact)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (vec *SavedPhoneContactVector) TypeID() uint32 {
	return SavedPhoneContactVectorTypeID
}

// SchemaName returns MTProto type name.
func (vec *SavedPhoneContactVector) SchemaName() string {
	return ""
}

// Encode implements bin.Encoder.
func (vec *SavedPhoneContactVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<SavedContact> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<SavedContact>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *SavedPhoneContactVector) GetElems() (value []SavedPhoneContact) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *SavedPhoneContactVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<SavedContact> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<SavedContact>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SavedPhoneContact
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<SavedContact>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for SavedPhoneContactVector.
var (
	_ bin.Encoder = &SavedPhoneContactVector{}
	_ bin.Decoder = &SavedPhoneContactVector{}
)
