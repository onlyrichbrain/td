// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// SetNetworkTypeRequest represents TL type `setNetworkType#d62de55e`.
type SetNetworkTypeRequest struct {
	// Type field of SetNetworkTypeRequest.
	Type NetworkTypeClass
}

// SetNetworkTypeRequestTypeID is TL type id of SetNetworkTypeRequest.
const SetNetworkTypeRequestTypeID = 0xd62de55e

// Ensuring interfaces in compile-time for SetNetworkTypeRequest.
var (
	_ bin.Encoder     = &SetNetworkTypeRequest{}
	_ bin.Decoder     = &SetNetworkTypeRequest{}
	_ bin.BareEncoder = &SetNetworkTypeRequest{}
	_ bin.BareDecoder = &SetNetworkTypeRequest{}
)

func (s *SetNetworkTypeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetNetworkTypeRequest) String() string {
	if s == nil {
		return "SetNetworkTypeRequest(nil)"
	}
	type Alias SetNetworkTypeRequest
	return fmt.Sprintf("SetNetworkTypeRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetNetworkTypeRequest) TypeID() uint32 {
	return SetNetworkTypeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetNetworkTypeRequest) TypeName() string {
	return "setNetworkType"
}

// TypeInfo returns info about TL type.
func (s *SetNetworkTypeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setNetworkType",
		ID:   SetNetworkTypeRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetNetworkTypeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setNetworkType#d62de55e as nil")
	}
	b.PutID(SetNetworkTypeRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetNetworkTypeRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setNetworkType#d62de55e as nil")
	}
	if s.Type == nil {
		return fmt.Errorf("unable to encode setNetworkType#d62de55e: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setNetworkType#d62de55e: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetNetworkTypeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setNetworkType#d62de55e to nil")
	}
	if err := b.ConsumeID(SetNetworkTypeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setNetworkType#d62de55e: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetNetworkTypeRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setNetworkType#d62de55e to nil")
	}
	{
		value, err := DecodeNetworkType(b)
		if err != nil {
			return fmt.Errorf("unable to decode setNetworkType#d62de55e: field type: %w", err)
		}
		s.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetNetworkTypeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setNetworkType#d62de55e as nil")
	}
	b.ObjStart()
	b.PutID("setNetworkType")
	b.Comma()
	b.FieldStart("type")
	if s.Type == nil {
		return fmt.Errorf("unable to encode setNetworkType#d62de55e: field type is nil")
	}
	if err := s.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setNetworkType#d62de55e: field type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetNetworkTypeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setNetworkType#d62de55e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setNetworkType"); err != nil {
				return fmt.Errorf("unable to decode setNetworkType#d62de55e: %w", err)
			}
		case "type":
			value, err := DecodeTDLibJSONNetworkType(b)
			if err != nil {
				return fmt.Errorf("unable to decode setNetworkType#d62de55e: field type: %w", err)
			}
			s.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetType returns value of Type field.
func (s *SetNetworkTypeRequest) GetType() (value NetworkTypeClass) {
	if s == nil {
		return
	}
	return s.Type
}

// SetNetworkType invokes method setNetworkType#d62de55e returning error if any.
func (c *Client) SetNetworkType(ctx context.Context, type_ NetworkTypeClass) error {
	var ok Ok

	request := &SetNetworkTypeRequest{
		Type: type_,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
