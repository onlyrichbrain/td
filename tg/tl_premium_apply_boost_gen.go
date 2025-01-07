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

// PremiumApplyBoostRequest represents TL type `premium.applyBoost#6b7da746`.
// Apply one or more boosts »¹ to a peer.
//
// Links:
//  1. https://core.telegram.org/api/boost
//
// See https://core.telegram.org/method/premium.applyBoost for reference.
type PremiumApplyBoostRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Which boost slots¹ to assign to this peer.
	//
	// Links:
	//  1) https://core.telegram.org/api/boost
	//
	// Use SetSlots and GetSlots helpers.
	Slots []int
	// The peer to boost.
	Peer InputPeerClass
}

// PremiumApplyBoostRequestTypeID is TL type id of PremiumApplyBoostRequest.
const PremiumApplyBoostRequestTypeID = 0x6b7da746

// Ensuring interfaces in compile-time for PremiumApplyBoostRequest.
var (
	_ bin.Encoder     = &PremiumApplyBoostRequest{}
	_ bin.Decoder     = &PremiumApplyBoostRequest{}
	_ bin.BareEncoder = &PremiumApplyBoostRequest{}
	_ bin.BareDecoder = &PremiumApplyBoostRequest{}
)

func (a *PremiumApplyBoostRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.Slots == nil) {
		return false
	}
	if !(a.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *PremiumApplyBoostRequest) String() string {
	if a == nil {
		return "PremiumApplyBoostRequest(nil)"
	}
	type Alias PremiumApplyBoostRequest
	return fmt.Sprintf("PremiumApplyBoostRequest%+v", Alias(*a))
}

// FillFrom fills PremiumApplyBoostRequest from given interface.
func (a *PremiumApplyBoostRequest) FillFrom(from interface {
	GetSlots() (value []int, ok bool)
	GetPeer() (value InputPeerClass)
}) {
	if val, ok := from.GetSlots(); ok {
		a.Slots = val
	}

	a.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumApplyBoostRequest) TypeID() uint32 {
	return PremiumApplyBoostRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumApplyBoostRequest) TypeName() string {
	return "premium.applyBoost"
}

// TypeInfo returns info about TL type.
func (a *PremiumApplyBoostRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premium.applyBoost",
		ID:   PremiumApplyBoostRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Slots",
			SchemaName: "slots",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *PremiumApplyBoostRequest) SetFlags() {
	if !(a.Slots == nil) {
		a.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (a *PremiumApplyBoostRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode premium.applyBoost#6b7da746 as nil")
	}
	b.PutID(PremiumApplyBoostRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *PremiumApplyBoostRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode premium.applyBoost#6b7da746 as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premium.applyBoost#6b7da746: field flags: %w", err)
	}
	if a.Flags.Has(0) {
		b.PutVectorHeader(len(a.Slots))
		for _, v := range a.Slots {
			b.PutInt(v)
		}
	}
	if a.Peer == nil {
		return fmt.Errorf("unable to encode premium.applyBoost#6b7da746: field peer is nil")
	}
	if err := a.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premium.applyBoost#6b7da746: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *PremiumApplyBoostRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode premium.applyBoost#6b7da746 to nil")
	}
	if err := b.ConsumeID(PremiumApplyBoostRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode premium.applyBoost#6b7da746: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *PremiumApplyBoostRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode premium.applyBoost#6b7da746 to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode premium.applyBoost#6b7da746: field flags: %w", err)
		}
	}
	if a.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode premium.applyBoost#6b7da746: field slots: %w", err)
		}

		if headerLen > 0 {
			a.Slots = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode premium.applyBoost#6b7da746: field slots: %w", err)
			}
			a.Slots = append(a.Slots, value)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode premium.applyBoost#6b7da746: field peer: %w", err)
		}
		a.Peer = value
	}
	return nil
}

// SetSlots sets value of Slots conditional field.
func (a *PremiumApplyBoostRequest) SetSlots(value []int) {
	a.Flags.Set(0)
	a.Slots = value
}

// GetSlots returns value of Slots conditional field and
// boolean which is true if field was set.
func (a *PremiumApplyBoostRequest) GetSlots() (value []int, ok bool) {
	if a == nil {
		return
	}
	if !a.Flags.Has(0) {
		return value, false
	}
	return a.Slots, true
}

// GetPeer returns value of Peer field.
func (a *PremiumApplyBoostRequest) GetPeer() (value InputPeerClass) {
	if a == nil {
		return
	}
	return a.Peer
}

// PremiumApplyBoost invokes method premium.applyBoost#6b7da746 returning error if any.
// Apply one or more boosts »¹ to a peer.
//
// Links:
//  1. https://core.telegram.org/api/boost
//
// Possible errors:
//
//	400 BOOSTS_EMPTY: No boost slots were specified.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 SLOTS_EMPTY: The specified slot list is empty.
//
// See https://core.telegram.org/method/premium.applyBoost for reference.
func (c *Client) PremiumApplyBoost(ctx context.Context, request *PremiumApplyBoostRequest) (*PremiumMyBoosts, error) {
	var result PremiumMyBoosts

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
