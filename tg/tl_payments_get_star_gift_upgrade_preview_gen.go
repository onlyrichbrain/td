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

// PaymentsGetStarGiftUpgradePreviewRequest represents TL type `payments.getStarGiftUpgradePreview#9c9abcb1`.
//
// See https://core.telegram.org/method/payments.getStarGiftUpgradePreview for reference.
type PaymentsGetStarGiftUpgradePreviewRequest struct {
	// GiftID field of PaymentsGetStarGiftUpgradePreviewRequest.
	GiftID int64
}

// PaymentsGetStarGiftUpgradePreviewRequestTypeID is TL type id of PaymentsGetStarGiftUpgradePreviewRequest.
const PaymentsGetStarGiftUpgradePreviewRequestTypeID = 0x9c9abcb1

// Ensuring interfaces in compile-time for PaymentsGetStarGiftUpgradePreviewRequest.
var (
	_ bin.Encoder     = &PaymentsGetStarGiftUpgradePreviewRequest{}
	_ bin.Decoder     = &PaymentsGetStarGiftUpgradePreviewRequest{}
	_ bin.BareEncoder = &PaymentsGetStarGiftUpgradePreviewRequest{}
	_ bin.BareDecoder = &PaymentsGetStarGiftUpgradePreviewRequest{}
)

func (g *PaymentsGetStarGiftUpgradePreviewRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.GiftID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PaymentsGetStarGiftUpgradePreviewRequest) String() string {
	if g == nil {
		return "PaymentsGetStarGiftUpgradePreviewRequest(nil)"
	}
	type Alias PaymentsGetStarGiftUpgradePreviewRequest
	return fmt.Sprintf("PaymentsGetStarGiftUpgradePreviewRequest%+v", Alias(*g))
}

// FillFrom fills PaymentsGetStarGiftUpgradePreviewRequest from given interface.
func (g *PaymentsGetStarGiftUpgradePreviewRequest) FillFrom(from interface {
	GetGiftID() (value int64)
}) {
	g.GiftID = from.GetGiftID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetStarGiftUpgradePreviewRequest) TypeID() uint32 {
	return PaymentsGetStarGiftUpgradePreviewRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetStarGiftUpgradePreviewRequest) TypeName() string {
	return "payments.getStarGiftUpgradePreview"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetStarGiftUpgradePreviewRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getStarGiftUpgradePreview",
		ID:   PaymentsGetStarGiftUpgradePreviewRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GiftID",
			SchemaName: "gift_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PaymentsGetStarGiftUpgradePreviewRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getStarGiftUpgradePreview#9c9abcb1 as nil")
	}
	b.PutID(PaymentsGetStarGiftUpgradePreviewRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PaymentsGetStarGiftUpgradePreviewRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getStarGiftUpgradePreview#9c9abcb1 as nil")
	}
	b.PutLong(g.GiftID)
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetStarGiftUpgradePreviewRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getStarGiftUpgradePreview#9c9abcb1 to nil")
	}
	if err := b.ConsumeID(PaymentsGetStarGiftUpgradePreviewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getStarGiftUpgradePreview#9c9abcb1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PaymentsGetStarGiftUpgradePreviewRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getStarGiftUpgradePreview#9c9abcb1 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getStarGiftUpgradePreview#9c9abcb1: field gift_id: %w", err)
		}
		g.GiftID = value
	}
	return nil
}

// GetGiftID returns value of GiftID field.
func (g *PaymentsGetStarGiftUpgradePreviewRequest) GetGiftID() (value int64) {
	if g == nil {
		return
	}
	return g.GiftID
}

// PaymentsGetStarGiftUpgradePreview invokes method payments.getStarGiftUpgradePreview#9c9abcb1 returning error if any.
//
// See https://core.telegram.org/method/payments.getStarGiftUpgradePreview for reference.
func (c *Client) PaymentsGetStarGiftUpgradePreview(ctx context.Context, giftid int64) (*PaymentsStarGiftUpgradePreview, error) {
	var result PaymentsStarGiftUpgradePreview

	request := &PaymentsGetStarGiftUpgradePreviewRequest{
		GiftID: giftid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}