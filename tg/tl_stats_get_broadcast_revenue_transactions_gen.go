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

// StatsGetBroadcastRevenueTransactionsRequest represents TL type `stats.getBroadcastRevenueTransactions#70990b6d`.
// Fetch channel ad revenue transaction history »¹.
//
// Links:
//  1. https://core.telegram.org/api/revenue
//
// See https://core.telegram.org/method/stats.getBroadcastRevenueTransactions for reference.
type StatsGetBroadcastRevenueTransactionsRequest struct {
	// Peer field of StatsGetBroadcastRevenueTransactionsRequest.
	Peer InputPeerClass
	// Offset for pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Offset int
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// StatsGetBroadcastRevenueTransactionsRequestTypeID is TL type id of StatsGetBroadcastRevenueTransactionsRequest.
const StatsGetBroadcastRevenueTransactionsRequestTypeID = 0x70990b6d

// Ensuring interfaces in compile-time for StatsGetBroadcastRevenueTransactionsRequest.
var (
	_ bin.Encoder     = &StatsGetBroadcastRevenueTransactionsRequest{}
	_ bin.Decoder     = &StatsGetBroadcastRevenueTransactionsRequest{}
	_ bin.BareEncoder = &StatsGetBroadcastRevenueTransactionsRequest{}
	_ bin.BareDecoder = &StatsGetBroadcastRevenueTransactionsRequest{}
)

func (g *StatsGetBroadcastRevenueTransactionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StatsGetBroadcastRevenueTransactionsRequest) String() string {
	if g == nil {
		return "StatsGetBroadcastRevenueTransactionsRequest(nil)"
	}
	type Alias StatsGetBroadcastRevenueTransactionsRequest
	return fmt.Sprintf("StatsGetBroadcastRevenueTransactionsRequest%+v", Alias(*g))
}

// FillFrom fills StatsGetBroadcastRevenueTransactionsRequest from given interface.
func (g *StatsGetBroadcastRevenueTransactionsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetOffset() (value int)
	GetLimit() (value int)
}) {
	g.Peer = from.GetPeer()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGetBroadcastRevenueTransactionsRequest) TypeID() uint32 {
	return StatsGetBroadcastRevenueTransactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGetBroadcastRevenueTransactionsRequest) TypeName() string {
	return "stats.getBroadcastRevenueTransactions"
}

// TypeInfo returns info about TL type.
func (g *StatsGetBroadcastRevenueTransactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.getBroadcastRevenueTransactions",
		ID:   StatsGetBroadcastRevenueTransactionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *StatsGetBroadcastRevenueTransactionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getBroadcastRevenueTransactions#70990b6d as nil")
	}
	b.PutID(StatsGetBroadcastRevenueTransactionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StatsGetBroadcastRevenueTransactionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getBroadcastRevenueTransactions#70990b6d as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode stats.getBroadcastRevenueTransactions#70990b6d: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getBroadcastRevenueTransactions#70990b6d: field peer: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *StatsGetBroadcastRevenueTransactionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getBroadcastRevenueTransactions#70990b6d to nil")
	}
	if err := b.ConsumeID(StatsGetBroadcastRevenueTransactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.getBroadcastRevenueTransactions#70990b6d: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StatsGetBroadcastRevenueTransactionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getBroadcastRevenueTransactions#70990b6d to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getBroadcastRevenueTransactions#70990b6d: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getBroadcastRevenueTransactions#70990b6d: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getBroadcastRevenueTransactions#70990b6d: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *StatsGetBroadcastRevenueTransactionsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetOffset returns value of Offset field.
func (g *StatsGetBroadcastRevenueTransactionsRequest) GetOffset() (value int) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *StatsGetBroadcastRevenueTransactionsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// StatsGetBroadcastRevenueTransactions invokes method stats.getBroadcastRevenueTransactions#70990b6d returning error if any.
// Fetch channel ad revenue transaction history »¹.
//
// Links:
//  1. https://core.telegram.org/api/revenue
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//
// See https://core.telegram.org/method/stats.getBroadcastRevenueTransactions for reference.
func (c *Client) StatsGetBroadcastRevenueTransactions(ctx context.Context, request *StatsGetBroadcastRevenueTransactionsRequest) (*StatsBroadcastRevenueTransactions, error) {
	var result StatsBroadcastRevenueTransactions

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
