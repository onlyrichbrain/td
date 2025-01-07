// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// PaymentsGetStarGiftsRequest represents TL type `payments.getStarGifts#c4563590`.
// Get a list of available gifts, see here »¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/gifts
//
// See https://core.telegram.org/method/payments.getStarGifts for reference.
type PaymentsGetStarGiftsRequest struct {
	// Hash used for caching, for more info click here¹.The hash may be generated locally by
	// using the ids of the returned or stored sticker starGift²s.
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	//  2) https://core.telegram.org/constructor/starGift
	Hash int
}

// PaymentsGetStarGiftsRequestTypeID is TL type id of PaymentsGetStarGiftsRequest.
const PaymentsGetStarGiftsRequestTypeID = 0xc4563590

// Ensuring interfaces in compile-time for PaymentsGetStarGiftsRequest.
var (
	_ bin.Encoder     = &PaymentsGetStarGiftsRequest{}
	_ bin.Decoder     = &PaymentsGetStarGiftsRequest{}
	_ bin.BareEncoder = &PaymentsGetStarGiftsRequest{}
	_ bin.BareDecoder = &PaymentsGetStarGiftsRequest{}
)

func (g *PaymentsGetStarGiftsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PaymentsGetStarGiftsRequest) String() string {
	if g == nil {
		return "PaymentsGetStarGiftsRequest(nil)"
	}
	type Alias PaymentsGetStarGiftsRequest
	return fmt.Sprintf("PaymentsGetStarGiftsRequest%+v", Alias(*g))
}

// FillFrom fills PaymentsGetStarGiftsRequest from given interface.
func (g *PaymentsGetStarGiftsRequest) FillFrom(from interface {
	GetHash() (value int)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetStarGiftsRequest) TypeID() uint32 {
	return PaymentsGetStarGiftsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetStarGiftsRequest) TypeName() string {
	return "payments.getStarGifts"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetStarGiftsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getStarGifts",
		ID:   PaymentsGetStarGiftsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PaymentsGetStarGiftsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getStarGifts#c4563590 as nil")
	}
	b.PutID(PaymentsGetStarGiftsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PaymentsGetStarGiftsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getStarGifts#c4563590 as nil")
	}
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetStarGiftsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getStarGifts#c4563590 to nil")
	}
	if err := b.ConsumeID(PaymentsGetStarGiftsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getStarGifts#c4563590: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PaymentsGetStarGiftsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getStarGifts#c4563590 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getStarGifts#c4563590: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *PaymentsGetStarGiftsRequest) GetHash() (value int) {
	if g == nil {
		return
	}
	return g.Hash
}

// PaymentsGetStarGifts invokes method payments.getStarGifts#c4563590 returning error if any.
// Get a list of available gifts, see here »¹ for more info.
//
// Links:
//  1. https://core.telegram.org/api/gifts
//
// See https://core.telegram.org/method/payments.getStarGifts for reference.
func (c *Client) PaymentsGetStarGifts(ctx context.Context, hash int) (PaymentsStarGiftsClass, error) {
	var result PaymentsStarGiftsBox

	request := &PaymentsGetStarGiftsRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.StarGifts, nil
}
