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

// ToggleGroupCallMuteNewParticipantsRequest represents TL type `toggleGroupCallMuteNewParticipants#10eec1c2`.
type ToggleGroupCallMuteNewParticipantsRequest struct {
	// Group call identifier
	GroupCallID int32
	// New value of the mute_new_participants setting
	MuteNewParticipants bool
}

// ToggleGroupCallMuteNewParticipantsRequestTypeID is TL type id of ToggleGroupCallMuteNewParticipantsRequest.
const ToggleGroupCallMuteNewParticipantsRequestTypeID = 0x10eec1c2

// Ensuring interfaces in compile-time for ToggleGroupCallMuteNewParticipantsRequest.
var (
	_ bin.Encoder     = &ToggleGroupCallMuteNewParticipantsRequest{}
	_ bin.Decoder     = &ToggleGroupCallMuteNewParticipantsRequest{}
	_ bin.BareEncoder = &ToggleGroupCallMuteNewParticipantsRequest{}
	_ bin.BareDecoder = &ToggleGroupCallMuteNewParticipantsRequest{}
)

func (t *ToggleGroupCallMuteNewParticipantsRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.GroupCallID == 0) {
		return false
	}
	if !(t.MuteNewParticipants == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleGroupCallMuteNewParticipantsRequest) String() string {
	if t == nil {
		return "ToggleGroupCallMuteNewParticipantsRequest(nil)"
	}
	type Alias ToggleGroupCallMuteNewParticipantsRequest
	return fmt.Sprintf("ToggleGroupCallMuteNewParticipantsRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleGroupCallMuteNewParticipantsRequest) TypeID() uint32 {
	return ToggleGroupCallMuteNewParticipantsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleGroupCallMuteNewParticipantsRequest) TypeName() string {
	return "toggleGroupCallMuteNewParticipants"
}

// TypeInfo returns info about TL type.
func (t *ToggleGroupCallMuteNewParticipantsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleGroupCallMuteNewParticipants",
		ID:   ToggleGroupCallMuteNewParticipantsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
		{
			Name:       "MuteNewParticipants",
			SchemaName: "mute_new_participants",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleGroupCallMuteNewParticipantsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleGroupCallMuteNewParticipants#10eec1c2 as nil")
	}
	b.PutID(ToggleGroupCallMuteNewParticipantsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleGroupCallMuteNewParticipantsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleGroupCallMuteNewParticipants#10eec1c2 as nil")
	}
	b.PutInt32(t.GroupCallID)
	b.PutBool(t.MuteNewParticipants)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleGroupCallMuteNewParticipantsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleGroupCallMuteNewParticipants#10eec1c2 to nil")
	}
	if err := b.ConsumeID(ToggleGroupCallMuteNewParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleGroupCallMuteNewParticipants#10eec1c2: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleGroupCallMuteNewParticipantsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleGroupCallMuteNewParticipants#10eec1c2 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode toggleGroupCallMuteNewParticipants#10eec1c2: field group_call_id: %w", err)
		}
		t.GroupCallID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleGroupCallMuteNewParticipants#10eec1c2: field mute_new_participants: %w", err)
		}
		t.MuteNewParticipants = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleGroupCallMuteNewParticipantsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleGroupCallMuteNewParticipants#10eec1c2 as nil")
	}
	b.ObjStart()
	b.PutID("toggleGroupCallMuteNewParticipants")
	b.Comma()
	b.FieldStart("group_call_id")
	b.PutInt32(t.GroupCallID)
	b.Comma()
	b.FieldStart("mute_new_participants")
	b.PutBool(t.MuteNewParticipants)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleGroupCallMuteNewParticipantsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleGroupCallMuteNewParticipants#10eec1c2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleGroupCallMuteNewParticipants"); err != nil {
				return fmt.Errorf("unable to decode toggleGroupCallMuteNewParticipants#10eec1c2: %w", err)
			}
		case "group_call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode toggleGroupCallMuteNewParticipants#10eec1c2: field group_call_id: %w", err)
			}
			t.GroupCallID = value
		case "mute_new_participants":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleGroupCallMuteNewParticipants#10eec1c2: field mute_new_participants: %w", err)
			}
			t.MuteNewParticipants = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGroupCallID returns value of GroupCallID field.
func (t *ToggleGroupCallMuteNewParticipantsRequest) GetGroupCallID() (value int32) {
	if t == nil {
		return
	}
	return t.GroupCallID
}

// GetMuteNewParticipants returns value of MuteNewParticipants field.
func (t *ToggleGroupCallMuteNewParticipantsRequest) GetMuteNewParticipants() (value bool) {
	if t == nil {
		return
	}
	return t.MuteNewParticipants
}

// ToggleGroupCallMuteNewParticipants invokes method toggleGroupCallMuteNewParticipants#10eec1c2 returning error if any.
func (c *Client) ToggleGroupCallMuteNewParticipants(ctx context.Context, request *ToggleGroupCallMuteNewParticipantsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
