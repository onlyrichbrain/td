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

// ToggleSupergroupSignMessagesRequest represents TL type `toggleSupergroupSignMessages#44efd524`.
type ToggleSupergroupSignMessagesRequest struct {
	// Identifier of the channel
	SupergroupID int64
	// New value of sign_messages
	SignMessages bool
}

// ToggleSupergroupSignMessagesRequestTypeID is TL type id of ToggleSupergroupSignMessagesRequest.
const ToggleSupergroupSignMessagesRequestTypeID = 0x44efd524

// Ensuring interfaces in compile-time for ToggleSupergroupSignMessagesRequest.
var (
	_ bin.Encoder     = &ToggleSupergroupSignMessagesRequest{}
	_ bin.Decoder     = &ToggleSupergroupSignMessagesRequest{}
	_ bin.BareEncoder = &ToggleSupergroupSignMessagesRequest{}
	_ bin.BareDecoder = &ToggleSupergroupSignMessagesRequest{}
)

func (t *ToggleSupergroupSignMessagesRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.SupergroupID == 0) {
		return false
	}
	if !(t.SignMessages == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleSupergroupSignMessagesRequest) String() string {
	if t == nil {
		return "ToggleSupergroupSignMessagesRequest(nil)"
	}
	type Alias ToggleSupergroupSignMessagesRequest
	return fmt.Sprintf("ToggleSupergroupSignMessagesRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleSupergroupSignMessagesRequest) TypeID() uint32 {
	return ToggleSupergroupSignMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleSupergroupSignMessagesRequest) TypeName() string {
	return "toggleSupergroupSignMessages"
}

// TypeInfo returns info about TL type.
func (t *ToggleSupergroupSignMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleSupergroupSignMessages",
		ID:   ToggleSupergroupSignMessagesRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
		{
			Name:       "SignMessages",
			SchemaName: "sign_messages",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleSupergroupSignMessagesRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupSignMessages#44efd524 as nil")
	}
	b.PutID(ToggleSupergroupSignMessagesRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleSupergroupSignMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupSignMessages#44efd524 as nil")
	}
	b.PutInt53(t.SupergroupID)
	b.PutBool(t.SignMessages)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleSupergroupSignMessagesRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupSignMessages#44efd524 to nil")
	}
	if err := b.ConsumeID(ToggleSupergroupSignMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleSupergroupSignMessages#44efd524: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleSupergroupSignMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupSignMessages#44efd524 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSupergroupSignMessages#44efd524: field supergroup_id: %w", err)
		}
		t.SupergroupID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSupergroupSignMessages#44efd524: field sign_messages: %w", err)
		}
		t.SignMessages = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleSupergroupSignMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupSignMessages#44efd524 as nil")
	}
	b.ObjStart()
	b.PutID("toggleSupergroupSignMessages")
	b.Comma()
	b.FieldStart("supergroup_id")
	b.PutInt53(t.SupergroupID)
	b.Comma()
	b.FieldStart("sign_messages")
	b.PutBool(t.SignMessages)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleSupergroupSignMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupSignMessages#44efd524 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleSupergroupSignMessages"); err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupSignMessages#44efd524: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupSignMessages#44efd524: field supergroup_id: %w", err)
			}
			t.SupergroupID = value
		case "sign_messages":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupSignMessages#44efd524: field sign_messages: %w", err)
			}
			t.SignMessages = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (t *ToggleSupergroupSignMessagesRequest) GetSupergroupID() (value int64) {
	if t == nil {
		return
	}
	return t.SupergroupID
}

// GetSignMessages returns value of SignMessages field.
func (t *ToggleSupergroupSignMessagesRequest) GetSignMessages() (value bool) {
	if t == nil {
		return
	}
	return t.SignMessages
}

// ToggleSupergroupSignMessages invokes method toggleSupergroupSignMessages#44efd524 returning error if any.
func (c *Client) ToggleSupergroupSignMessages(ctx context.Context, request *ToggleSupergroupSignMessagesRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
