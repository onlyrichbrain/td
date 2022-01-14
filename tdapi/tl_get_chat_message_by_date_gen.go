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

// GetChatMessageByDateRequest represents TL type `getChatMessageByDate#3f557136`.
type GetChatMessageByDateRequest struct {
	// Chat identifier
	ChatID int64
	// Point in time (Unix timestamp) relative to which to search for messages
	Date int32
}

// GetChatMessageByDateRequestTypeID is TL type id of GetChatMessageByDateRequest.
const GetChatMessageByDateRequestTypeID = 0x3f557136

// Ensuring interfaces in compile-time for GetChatMessageByDateRequest.
var (
	_ bin.Encoder     = &GetChatMessageByDateRequest{}
	_ bin.Decoder     = &GetChatMessageByDateRequest{}
	_ bin.BareEncoder = &GetChatMessageByDateRequest{}
	_ bin.BareDecoder = &GetChatMessageByDateRequest{}
)

func (g *GetChatMessageByDateRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatMessageByDateRequest) String() string {
	if g == nil {
		return "GetChatMessageByDateRequest(nil)"
	}
	type Alias GetChatMessageByDateRequest
	return fmt.Sprintf("GetChatMessageByDateRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatMessageByDateRequest) TypeID() uint32 {
	return GetChatMessageByDateRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatMessageByDateRequest) TypeName() string {
	return "getChatMessageByDate"
}

// TypeInfo returns info about TL type.
func (g *GetChatMessageByDateRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatMessageByDate",
		ID:   GetChatMessageByDateRequestTypeID,
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
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatMessageByDateRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessageByDate#3f557136 as nil")
	}
	b.PutID(GetChatMessageByDateRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatMessageByDateRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessageByDate#3f557136 as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt32(g.Date)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatMessageByDateRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessageByDate#3f557136 to nil")
	}
	if err := b.ConsumeID(GetChatMessageByDateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatMessageByDate#3f557136: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatMessageByDateRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessageByDate#3f557136 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessageByDate#3f557136: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessageByDate#3f557136: field date: %w", err)
		}
		g.Date = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatMessageByDateRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessageByDate#3f557136 as nil")
	}
	b.ObjStart()
	b.PutID("getChatMessageByDate")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("date")
	b.PutInt32(g.Date)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatMessageByDateRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessageByDate#3f557136 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatMessageByDate"); err != nil {
				return fmt.Errorf("unable to decode getChatMessageByDate#3f557136: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessageByDate#3f557136: field chat_id: %w", err)
			}
			g.ChatID = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessageByDate#3f557136: field date: %w", err)
			}
			g.Date = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatMessageByDateRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetDate returns value of Date field.
func (g *GetChatMessageByDateRequest) GetDate() (value int32) {
	if g == nil {
		return
	}
	return g.Date
}

// GetChatMessageByDate invokes method getChatMessageByDate#3f557136 returning error if any.
func (c *Client) GetChatMessageByDate(ctx context.Context, request *GetChatMessageByDateRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
