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

// DeleteChatRequest represents TL type `deleteChat#f5cae05e`.
type DeleteChatRequest struct {
	// Chat identifier
	ChatID int64
}

// DeleteChatRequestTypeID is TL type id of DeleteChatRequest.
const DeleteChatRequestTypeID = 0xf5cae05e

// Ensuring interfaces in compile-time for DeleteChatRequest.
var (
	_ bin.Encoder     = &DeleteChatRequest{}
	_ bin.Decoder     = &DeleteChatRequest{}
	_ bin.BareEncoder = &DeleteChatRequest{}
	_ bin.BareDecoder = &DeleteChatRequest{}
)

func (d *DeleteChatRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteChatRequest) String() string {
	if d == nil {
		return "DeleteChatRequest(nil)"
	}
	type Alias DeleteChatRequest
	return fmt.Sprintf("DeleteChatRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteChatRequest) TypeID() uint32 {
	return DeleteChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteChatRequest) TypeName() string {
	return "deleteChat"
}

// TypeInfo returns info about TL type.
func (d *DeleteChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteChat",
		ID:   DeleteChatRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteChatRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChat#f5cae05e as nil")
	}
	b.PutID(DeleteChatRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteChatRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChat#f5cae05e as nil")
	}
	b.PutInt53(d.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteChatRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChat#f5cae05e to nil")
	}
	if err := b.ConsumeID(DeleteChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteChat#f5cae05e: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteChatRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChat#f5cae05e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteChat#f5cae05e: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChat#f5cae05e as nil")
	}
	b.ObjStart()
	b.PutID("deleteChat")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(d.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChat#f5cae05e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteChat"); err != nil {
				return fmt.Errorf("unable to decode deleteChat#f5cae05e: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteChat#f5cae05e: field chat_id: %w", err)
			}
			d.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (d *DeleteChatRequest) GetChatID() (value int64) {
	if d == nil {
		return
	}
	return d.ChatID
}

// DeleteChat invokes method deleteChat#f5cae05e returning error if any.
func (c *Client) DeleteChat(ctx context.Context, chatid int64) error {
	var ok Ok

	request := &DeleteChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
