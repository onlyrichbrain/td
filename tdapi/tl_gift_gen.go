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

// Gift represents TL type `gift#9dc9b8a1`.
type Gift struct {
	// Unique identifier of the gift
	ID int64
	// The sticker representing the gift
	Sticker Sticker
	// Number of Telegram Stars that must be paid for the gift
	StarCount int64
	// Number of Telegram Stars that can be claimed by the receiver instead of the gift by
	// default. If the gift was paid with just bought Telegram Stars, then full value can be
	// claimed
	DefaultSellStarCount int64
	// Number of remaining times the gift can be purchased by all users; 0 if not limited or
	// the gift was sold out
	RemainingCount int32
	// Number of total times the gift can be purchased by all users; 0 if not limited
	TotalCount int32
	// Point in time (Unix timestamp) when the gift was send for the first time; for sold out
	// gifts only
	FirstSendDate int32
	// Point in time (Unix timestamp) when the gift was send for the last time; for sold out
	// gifts only
	LastSendDate int32
}

// GiftTypeID is TL type id of Gift.
const GiftTypeID = 0x9dc9b8a1

// Ensuring interfaces in compile-time for Gift.
var (
	_ bin.Encoder     = &Gift{}
	_ bin.Decoder     = &Gift{}
	_ bin.BareEncoder = &Gift{}
	_ bin.BareDecoder = &Gift{}
)

func (g *Gift) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == 0) {
		return false
	}
	if !(g.Sticker.Zero()) {
		return false
	}
	if !(g.StarCount == 0) {
		return false
	}
	if !(g.DefaultSellStarCount == 0) {
		return false
	}
	if !(g.RemainingCount == 0) {
		return false
	}
	if !(g.TotalCount == 0) {
		return false
	}
	if !(g.FirstSendDate == 0) {
		return false
	}
	if !(g.LastSendDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *Gift) String() string {
	if g == nil {
		return "Gift(nil)"
	}
	type Alias Gift
	return fmt.Sprintf("Gift%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Gift) TypeID() uint32 {
	return GiftTypeID
}

// TypeName returns name of type in TL schema.
func (*Gift) TypeName() string {
	return "gift"
}

// TypeInfo returns info about TL type.
func (g *Gift) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "gift",
		ID:   GiftTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
		{
			Name:       "StarCount",
			SchemaName: "star_count",
		},
		{
			Name:       "DefaultSellStarCount",
			SchemaName: "default_sell_star_count",
		},
		{
			Name:       "RemainingCount",
			SchemaName: "remaining_count",
		},
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "FirstSendDate",
			SchemaName: "first_send_date",
		},
		{
			Name:       "LastSendDate",
			SchemaName: "last_send_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *Gift) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode gift#9dc9b8a1 as nil")
	}
	b.PutID(GiftTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *Gift) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode gift#9dc9b8a1 as nil")
	}
	b.PutLong(g.ID)
	if err := g.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode gift#9dc9b8a1: field sticker: %w", err)
	}
	b.PutInt53(g.StarCount)
	b.PutInt53(g.DefaultSellStarCount)
	b.PutInt32(g.RemainingCount)
	b.PutInt32(g.TotalCount)
	b.PutInt32(g.FirstSendDate)
	b.PutInt32(g.LastSendDate)
	return nil
}

// Decode implements bin.Decoder.
func (g *Gift) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode gift#9dc9b8a1 to nil")
	}
	if err := b.ConsumeID(GiftTypeID); err != nil {
		return fmt.Errorf("unable to decode gift#9dc9b8a1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *Gift) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode gift#9dc9b8a1 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode gift#9dc9b8a1: field id: %w", err)
		}
		g.ID = value
	}
	{
		if err := g.Sticker.Decode(b); err != nil {
			return fmt.Errorf("unable to decode gift#9dc9b8a1: field sticker: %w", err)
		}
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode gift#9dc9b8a1: field star_count: %w", err)
		}
		g.StarCount = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode gift#9dc9b8a1: field default_sell_star_count: %w", err)
		}
		g.DefaultSellStarCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode gift#9dc9b8a1: field remaining_count: %w", err)
		}
		g.RemainingCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode gift#9dc9b8a1: field total_count: %w", err)
		}
		g.TotalCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode gift#9dc9b8a1: field first_send_date: %w", err)
		}
		g.FirstSendDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode gift#9dc9b8a1: field last_send_date: %w", err)
		}
		g.LastSendDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *Gift) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode gift#9dc9b8a1 as nil")
	}
	b.ObjStart()
	b.PutID("gift")
	b.Comma()
	b.FieldStart("id")
	b.PutLong(g.ID)
	b.Comma()
	b.FieldStart("sticker")
	if err := g.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode gift#9dc9b8a1: field sticker: %w", err)
	}
	b.Comma()
	b.FieldStart("star_count")
	b.PutInt53(g.StarCount)
	b.Comma()
	b.FieldStart("default_sell_star_count")
	b.PutInt53(g.DefaultSellStarCount)
	b.Comma()
	b.FieldStart("remaining_count")
	b.PutInt32(g.RemainingCount)
	b.Comma()
	b.FieldStart("total_count")
	b.PutInt32(g.TotalCount)
	b.Comma()
	b.FieldStart("first_send_date")
	b.PutInt32(g.FirstSendDate)
	b.Comma()
	b.FieldStart("last_send_date")
	b.PutInt32(g.LastSendDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *Gift) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode gift#9dc9b8a1 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("gift"); err != nil {
				return fmt.Errorf("unable to decode gift#9dc9b8a1: %w", err)
			}
		case "id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode gift#9dc9b8a1: field id: %w", err)
			}
			g.ID = value
		case "sticker":
			if err := g.Sticker.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode gift#9dc9b8a1: field sticker: %w", err)
			}
		case "star_count":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode gift#9dc9b8a1: field star_count: %w", err)
			}
			g.StarCount = value
		case "default_sell_star_count":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode gift#9dc9b8a1: field default_sell_star_count: %w", err)
			}
			g.DefaultSellStarCount = value
		case "remaining_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode gift#9dc9b8a1: field remaining_count: %w", err)
			}
			g.RemainingCount = value
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode gift#9dc9b8a1: field total_count: %w", err)
			}
			g.TotalCount = value
		case "first_send_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode gift#9dc9b8a1: field first_send_date: %w", err)
			}
			g.FirstSendDate = value
		case "last_send_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode gift#9dc9b8a1: field last_send_date: %w", err)
			}
			g.LastSendDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (g *Gift) GetID() (value int64) {
	if g == nil {
		return
	}
	return g.ID
}

// GetSticker returns value of Sticker field.
func (g *Gift) GetSticker() (value Sticker) {
	if g == nil {
		return
	}
	return g.Sticker
}

// GetStarCount returns value of StarCount field.
func (g *Gift) GetStarCount() (value int64) {
	if g == nil {
		return
	}
	return g.StarCount
}

// GetDefaultSellStarCount returns value of DefaultSellStarCount field.
func (g *Gift) GetDefaultSellStarCount() (value int64) {
	if g == nil {
		return
	}
	return g.DefaultSellStarCount
}

// GetRemainingCount returns value of RemainingCount field.
func (g *Gift) GetRemainingCount() (value int32) {
	if g == nil {
		return
	}
	return g.RemainingCount
}

// GetTotalCount returns value of TotalCount field.
func (g *Gift) GetTotalCount() (value int32) {
	if g == nil {
		return
	}
	return g.TotalCount
}

// GetFirstSendDate returns value of FirstSendDate field.
func (g *Gift) GetFirstSendDate() (value int32) {
	if g == nil {
		return
	}
	return g.FirstSendDate
}

// GetLastSendDate returns value of LastSendDate field.
func (g *Gift) GetLastSendDate() (value int32) {
	if g == nil {
		return
	}
	return g.LastSendDate
}
