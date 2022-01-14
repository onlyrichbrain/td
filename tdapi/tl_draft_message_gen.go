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

// DraftMessage represents TL type `draftMessage#51d71500`.
type DraftMessage struct {
	// Identifier of the message to reply to; 0 if none
	ReplyToMessageID int64
	// Point in time (Unix timestamp) when the draft was created
	Date int32
	// Content of the message draft; must be of the type inputMessageText
	InputMessageText InputMessageContentClass
}

// DraftMessageTypeID is TL type id of DraftMessage.
const DraftMessageTypeID = 0x51d71500

// Ensuring interfaces in compile-time for DraftMessage.
var (
	_ bin.Encoder     = &DraftMessage{}
	_ bin.Decoder     = &DraftMessage{}
	_ bin.BareEncoder = &DraftMessage{}
	_ bin.BareDecoder = &DraftMessage{}
)

func (d *DraftMessage) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ReplyToMessageID == 0) {
		return false
	}
	if !(d.Date == 0) {
		return false
	}
	if !(d.InputMessageText == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DraftMessage) String() string {
	if d == nil {
		return "DraftMessage(nil)"
	}
	type Alias DraftMessage
	return fmt.Sprintf("DraftMessage%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DraftMessage) TypeID() uint32 {
	return DraftMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*DraftMessage) TypeName() string {
	return "draftMessage"
}

// TypeInfo returns info about TL type.
func (d *DraftMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "draftMessage",
		ID:   DraftMessageTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ReplyToMessageID",
			SchemaName: "reply_to_message_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "InputMessageText",
			SchemaName: "input_message_text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DraftMessage) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode draftMessage#51d71500 as nil")
	}
	b.PutID(DraftMessageTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DraftMessage) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode draftMessage#51d71500 as nil")
	}
	b.PutInt53(d.ReplyToMessageID)
	b.PutInt32(d.Date)
	if d.InputMessageText == nil {
		return fmt.Errorf("unable to encode draftMessage#51d71500: field input_message_text is nil")
	}
	if err := d.InputMessageText.Encode(b); err != nil {
		return fmt.Errorf("unable to encode draftMessage#51d71500: field input_message_text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DraftMessage) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode draftMessage#51d71500 to nil")
	}
	if err := b.ConsumeID(DraftMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode draftMessage#51d71500: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DraftMessage) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode draftMessage#51d71500 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode draftMessage#51d71500: field reply_to_message_id: %w", err)
		}
		d.ReplyToMessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode draftMessage#51d71500: field date: %w", err)
		}
		d.Date = value
	}
	{
		value, err := DecodeInputMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode draftMessage#51d71500: field input_message_text: %w", err)
		}
		d.InputMessageText = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DraftMessage) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode draftMessage#51d71500 as nil")
	}
	b.ObjStart()
	b.PutID("draftMessage")
	b.Comma()
	b.FieldStart("reply_to_message_id")
	b.PutInt53(d.ReplyToMessageID)
	b.Comma()
	b.FieldStart("date")
	b.PutInt32(d.Date)
	b.Comma()
	b.FieldStart("input_message_text")
	if d.InputMessageText == nil {
		return fmt.Errorf("unable to encode draftMessage#51d71500: field input_message_text is nil")
	}
	if err := d.InputMessageText.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode draftMessage#51d71500: field input_message_text: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DraftMessage) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode draftMessage#51d71500 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("draftMessage"); err != nil {
				return fmt.Errorf("unable to decode draftMessage#51d71500: %w", err)
			}
		case "reply_to_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode draftMessage#51d71500: field reply_to_message_id: %w", err)
			}
			d.ReplyToMessageID = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode draftMessage#51d71500: field date: %w", err)
			}
			d.Date = value
		case "input_message_text":
			value, err := DecodeTDLibJSONInputMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode draftMessage#51d71500: field input_message_text: %w", err)
			}
			d.InputMessageText = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetReplyToMessageID returns value of ReplyToMessageID field.
func (d *DraftMessage) GetReplyToMessageID() (value int64) {
	if d == nil {
		return
	}
	return d.ReplyToMessageID
}

// GetDate returns value of Date field.
func (d *DraftMessage) GetDate() (value int32) {
	if d == nil {
		return
	}
	return d.Date
}

// GetInputMessageText returns value of InputMessageText field.
func (d *DraftMessage) GetInputMessageText() (value InputMessageContentClass) {
	if d == nil {
		return
	}
	return d.InputMessageText
}
