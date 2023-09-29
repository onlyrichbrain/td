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

// StoriesGetBoostsStatusRequest represents TL type `stories.getBoostsStatus#4c449472`.
//
// See https://core.telegram.org/method/stories.getBoostsStatus for reference.
type StoriesGetBoostsStatusRequest struct {
	// Peer field of StoriesGetBoostsStatusRequest.
	Peer InputPeerClass
}

// StoriesGetBoostsStatusRequestTypeID is TL type id of StoriesGetBoostsStatusRequest.
const StoriesGetBoostsStatusRequestTypeID = 0x4c449472

// Ensuring interfaces in compile-time for StoriesGetBoostsStatusRequest.
var (
	_ bin.Encoder     = &StoriesGetBoostsStatusRequest{}
	_ bin.Decoder     = &StoriesGetBoostsStatusRequest{}
	_ bin.BareEncoder = &StoriesGetBoostsStatusRequest{}
	_ bin.BareDecoder = &StoriesGetBoostsStatusRequest{}
)

func (g *StoriesGetBoostsStatusRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StoriesGetBoostsStatusRequest) String() string {
	if g == nil {
		return "StoriesGetBoostsStatusRequest(nil)"
	}
	type Alias StoriesGetBoostsStatusRequest
	return fmt.Sprintf("StoriesGetBoostsStatusRequest%+v", Alias(*g))
}

// FillFrom fills StoriesGetBoostsStatusRequest from given interface.
func (g *StoriesGetBoostsStatusRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	g.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesGetBoostsStatusRequest) TypeID() uint32 {
	return StoriesGetBoostsStatusRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesGetBoostsStatusRequest) TypeName() string {
	return "stories.getBoostsStatus"
}

// TypeInfo returns info about TL type.
func (g *StoriesGetBoostsStatusRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.getBoostsStatus",
		ID:   StoriesGetBoostsStatusRequestTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *StoriesGetBoostsStatusRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getBoostsStatus#4c449472 as nil")
	}
	b.PutID(StoriesGetBoostsStatusRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StoriesGetBoostsStatusRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stories.getBoostsStatus#4c449472 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode stories.getBoostsStatus#4c449472: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.getBoostsStatus#4c449472: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *StoriesGetBoostsStatusRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getBoostsStatus#4c449472 to nil")
	}
	if err := b.ConsumeID(StoriesGetBoostsStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.getBoostsStatus#4c449472: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StoriesGetBoostsStatusRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stories.getBoostsStatus#4c449472 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.getBoostsStatus#4c449472: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *StoriesGetBoostsStatusRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// StoriesGetBoostsStatus invokes method stories.getBoostsStatus#4c449472 returning error if any.
//
// See https://core.telegram.org/method/stories.getBoostsStatus for reference.
func (c *Client) StoriesGetBoostsStatus(ctx context.Context, peer InputPeerClass) (*StoriesBoostsStatus, error) {
	var result StoriesBoostsStatus

	request := &StoriesGetBoostsStatusRequest{
		Peer: peer,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}