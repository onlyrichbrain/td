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

// GetChatScheduledMessagesRequest represents TL type `getChatScheduledMessages#df3d2ffb`.
type GetChatScheduledMessagesRequest struct {
	// Chat identifier
	ChatID int64
}

// GetChatScheduledMessagesRequestTypeID is TL type id of GetChatScheduledMessagesRequest.
const GetChatScheduledMessagesRequestTypeID = 0xdf3d2ffb

// Ensuring interfaces in compile-time for GetChatScheduledMessagesRequest.
var (
	_ bin.Encoder     = &GetChatScheduledMessagesRequest{}
	_ bin.Decoder     = &GetChatScheduledMessagesRequest{}
	_ bin.BareEncoder = &GetChatScheduledMessagesRequest{}
	_ bin.BareDecoder = &GetChatScheduledMessagesRequest{}
)

func (g *GetChatScheduledMessagesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatScheduledMessagesRequest) String() string {
	if g == nil {
		return "GetChatScheduledMessagesRequest(nil)"
	}
	type Alias GetChatScheduledMessagesRequest
	return fmt.Sprintf("GetChatScheduledMessagesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatScheduledMessagesRequest) TypeID() uint32 {
	return GetChatScheduledMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatScheduledMessagesRequest) TypeName() string {
	return "getChatScheduledMessages"
}

// TypeInfo returns info about TL type.
func (g *GetChatScheduledMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatScheduledMessages",
		ID:   GetChatScheduledMessagesRequestTypeID,
	}
	if g == nil {
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
func (g *GetChatScheduledMessagesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatScheduledMessages#df3d2ffb as nil")
	}
	b.PutID(GetChatScheduledMessagesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatScheduledMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatScheduledMessages#df3d2ffb as nil")
	}
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatScheduledMessagesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatScheduledMessages#df3d2ffb to nil")
	}
	if err := b.ConsumeID(GetChatScheduledMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatScheduledMessages#df3d2ffb: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatScheduledMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatScheduledMessages#df3d2ffb to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatScheduledMessages#df3d2ffb: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatScheduledMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatScheduledMessages#df3d2ffb as nil")
	}
	b.ObjStart()
	b.PutID("getChatScheduledMessages")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatScheduledMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatScheduledMessages#df3d2ffb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatScheduledMessages"); err != nil {
				return fmt.Errorf("unable to decode getChatScheduledMessages#df3d2ffb: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatScheduledMessages#df3d2ffb: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatScheduledMessagesRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetChatScheduledMessages invokes method getChatScheduledMessages#df3d2ffb returning error if any.
func (c *Client) GetChatScheduledMessages(ctx context.Context, chatid int64) (*Messages, error) {
	var result Messages

	request := &GetChatScheduledMessagesRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
