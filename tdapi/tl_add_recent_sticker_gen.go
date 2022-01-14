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

// AddRecentStickerRequest represents TL type `addRecentSticker#a7e5d89e`.
type AddRecentStickerRequest struct {
	// Pass true to add the sticker to the list of stickers recently attached to photo or
	// video files; pass false to add the sticker to the list of recently sent stickers
	IsAttached bool
	// Sticker file to add
	Sticker InputFileClass
}

// AddRecentStickerRequestTypeID is TL type id of AddRecentStickerRequest.
const AddRecentStickerRequestTypeID = 0xa7e5d89e

// Ensuring interfaces in compile-time for AddRecentStickerRequest.
var (
	_ bin.Encoder     = &AddRecentStickerRequest{}
	_ bin.Decoder     = &AddRecentStickerRequest{}
	_ bin.BareEncoder = &AddRecentStickerRequest{}
	_ bin.BareDecoder = &AddRecentStickerRequest{}
)

func (a *AddRecentStickerRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.IsAttached == false) {
		return false
	}
	if !(a.Sticker == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddRecentStickerRequest) String() string {
	if a == nil {
		return "AddRecentStickerRequest(nil)"
	}
	type Alias AddRecentStickerRequest
	return fmt.Sprintf("AddRecentStickerRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddRecentStickerRequest) TypeID() uint32 {
	return AddRecentStickerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddRecentStickerRequest) TypeName() string {
	return "addRecentSticker"
}

// TypeInfo returns info about TL type.
func (a *AddRecentStickerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addRecentSticker",
		ID:   AddRecentStickerRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsAttached",
			SchemaName: "is_attached",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddRecentStickerRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addRecentSticker#a7e5d89e as nil")
	}
	b.PutID(AddRecentStickerRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddRecentStickerRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addRecentSticker#a7e5d89e as nil")
	}
	b.PutBool(a.IsAttached)
	if a.Sticker == nil {
		return fmt.Errorf("unable to encode addRecentSticker#a7e5d89e: field sticker is nil")
	}
	if err := a.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode addRecentSticker#a7e5d89e: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AddRecentStickerRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addRecentSticker#a7e5d89e to nil")
	}
	if err := b.ConsumeID(AddRecentStickerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addRecentSticker#a7e5d89e: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddRecentStickerRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addRecentSticker#a7e5d89e to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode addRecentSticker#a7e5d89e: field is_attached: %w", err)
		}
		a.IsAttached = value
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode addRecentSticker#a7e5d89e: field sticker: %w", err)
		}
		a.Sticker = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AddRecentStickerRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addRecentSticker#a7e5d89e as nil")
	}
	b.ObjStart()
	b.PutID("addRecentSticker")
	b.Comma()
	b.FieldStart("is_attached")
	b.PutBool(a.IsAttached)
	b.Comma()
	b.FieldStart("sticker")
	if a.Sticker == nil {
		return fmt.Errorf("unable to encode addRecentSticker#a7e5d89e: field sticker is nil")
	}
	if err := a.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode addRecentSticker#a7e5d89e: field sticker: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AddRecentStickerRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addRecentSticker#a7e5d89e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("addRecentSticker"); err != nil {
				return fmt.Errorf("unable to decode addRecentSticker#a7e5d89e: %w", err)
			}
		case "is_attached":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode addRecentSticker#a7e5d89e: field is_attached: %w", err)
			}
			a.IsAttached = value
		case "sticker":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode addRecentSticker#a7e5d89e: field sticker: %w", err)
			}
			a.Sticker = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIsAttached returns value of IsAttached field.
func (a *AddRecentStickerRequest) GetIsAttached() (value bool) {
	if a == nil {
		return
	}
	return a.IsAttached
}

// GetSticker returns value of Sticker field.
func (a *AddRecentStickerRequest) GetSticker() (value InputFileClass) {
	if a == nil {
		return
	}
	return a.Sticker
}

// AddRecentSticker invokes method addRecentSticker#a7e5d89e returning error if any.
func (c *Client) AddRecentSticker(ctx context.Context, request *AddRecentStickerRequest) (*Stickers, error) {
	var result Stickers

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
