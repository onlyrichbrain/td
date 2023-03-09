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

// GetDefaultChatPhotoCustomEmojiStickersRequest represents TL type `getDefaultChatPhotoCustomEmojiStickers#e9917765`.
type GetDefaultChatPhotoCustomEmojiStickersRequest struct {
}

// GetDefaultChatPhotoCustomEmojiStickersRequestTypeID is TL type id of GetDefaultChatPhotoCustomEmojiStickersRequest.
const GetDefaultChatPhotoCustomEmojiStickersRequestTypeID = 0xe9917765

// Ensuring interfaces in compile-time for GetDefaultChatPhotoCustomEmojiStickersRequest.
var (
	_ bin.Encoder     = &GetDefaultChatPhotoCustomEmojiStickersRequest{}
	_ bin.Decoder     = &GetDefaultChatPhotoCustomEmojiStickersRequest{}
	_ bin.BareEncoder = &GetDefaultChatPhotoCustomEmojiStickersRequest{}
	_ bin.BareDecoder = &GetDefaultChatPhotoCustomEmojiStickersRequest{}
)

func (g *GetDefaultChatPhotoCustomEmojiStickersRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetDefaultChatPhotoCustomEmojiStickersRequest) String() string {
	if g == nil {
		return "GetDefaultChatPhotoCustomEmojiStickersRequest(nil)"
	}
	type Alias GetDefaultChatPhotoCustomEmojiStickersRequest
	return fmt.Sprintf("GetDefaultChatPhotoCustomEmojiStickersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetDefaultChatPhotoCustomEmojiStickersRequest) TypeID() uint32 {
	return GetDefaultChatPhotoCustomEmojiStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetDefaultChatPhotoCustomEmojiStickersRequest) TypeName() string {
	return "getDefaultChatPhotoCustomEmojiStickers"
}

// TypeInfo returns info about TL type.
func (g *GetDefaultChatPhotoCustomEmojiStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getDefaultChatPhotoCustomEmojiStickers",
		ID:   GetDefaultChatPhotoCustomEmojiStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetDefaultChatPhotoCustomEmojiStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDefaultChatPhotoCustomEmojiStickers#e9917765 as nil")
	}
	b.PutID(GetDefaultChatPhotoCustomEmojiStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetDefaultChatPhotoCustomEmojiStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDefaultChatPhotoCustomEmojiStickers#e9917765 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetDefaultChatPhotoCustomEmojiStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDefaultChatPhotoCustomEmojiStickers#e9917765 to nil")
	}
	if err := b.ConsumeID(GetDefaultChatPhotoCustomEmojiStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getDefaultChatPhotoCustomEmojiStickers#e9917765: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetDefaultChatPhotoCustomEmojiStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDefaultChatPhotoCustomEmojiStickers#e9917765 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetDefaultChatPhotoCustomEmojiStickersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getDefaultChatPhotoCustomEmojiStickers#e9917765 as nil")
	}
	b.ObjStart()
	b.PutID("getDefaultChatPhotoCustomEmojiStickers")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetDefaultChatPhotoCustomEmojiStickersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getDefaultChatPhotoCustomEmojiStickers#e9917765 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getDefaultChatPhotoCustomEmojiStickers"); err != nil {
				return fmt.Errorf("unable to decode getDefaultChatPhotoCustomEmojiStickers#e9917765: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDefaultChatPhotoCustomEmojiStickers invokes method getDefaultChatPhotoCustomEmojiStickers#e9917765 returning error if any.
func (c *Client) GetDefaultChatPhotoCustomEmojiStickers(ctx context.Context) (*Stickers, error) {
	var result Stickers

	request := &GetDefaultChatPhotoCustomEmojiStickersRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}