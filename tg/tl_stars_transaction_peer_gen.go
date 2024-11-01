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

// StarsTransactionPeerUnsupported represents TL type `starsTransactionPeerUnsupported#95f2bfe4`.
// Describes a Telegram Star¹ transaction that cannot be described using the current
// layer.
//
// Links:
//  1. https://core.telegram.org/api/stars
//
// See https://core.telegram.org/constructor/starsTransactionPeerUnsupported for reference.
type StarsTransactionPeerUnsupported struct {
}

// StarsTransactionPeerUnsupportedTypeID is TL type id of StarsTransactionPeerUnsupported.
const StarsTransactionPeerUnsupportedTypeID = 0x95f2bfe4

// construct implements constructor of StarsTransactionPeerClass.
func (s StarsTransactionPeerUnsupported) construct() StarsTransactionPeerClass { return &s }

// Ensuring interfaces in compile-time for StarsTransactionPeerUnsupported.
var (
	_ bin.Encoder     = &StarsTransactionPeerUnsupported{}
	_ bin.Decoder     = &StarsTransactionPeerUnsupported{}
	_ bin.BareEncoder = &StarsTransactionPeerUnsupported{}
	_ bin.BareDecoder = &StarsTransactionPeerUnsupported{}

	_ StarsTransactionPeerClass = &StarsTransactionPeerUnsupported{}
)

func (s *StarsTransactionPeerUnsupported) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarsTransactionPeerUnsupported) String() string {
	if s == nil {
		return "StarsTransactionPeerUnsupported(nil)"
	}
	type Alias StarsTransactionPeerUnsupported
	return fmt.Sprintf("StarsTransactionPeerUnsupported%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarsTransactionPeerUnsupported) TypeID() uint32 {
	return StarsTransactionPeerUnsupportedTypeID
}

// TypeName returns name of type in TL schema.
func (*StarsTransactionPeerUnsupported) TypeName() string {
	return "starsTransactionPeerUnsupported"
}

// TypeInfo returns info about TL type.
func (s *StarsTransactionPeerUnsupported) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starsTransactionPeerUnsupported",
		ID:   StarsTransactionPeerUnsupportedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarsTransactionPeerUnsupported) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerUnsupported#95f2bfe4 as nil")
	}
	b.PutID(StarsTransactionPeerUnsupportedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarsTransactionPeerUnsupported) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerUnsupported#95f2bfe4 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarsTransactionPeerUnsupported) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerUnsupported#95f2bfe4 to nil")
	}
	if err := b.ConsumeID(StarsTransactionPeerUnsupportedTypeID); err != nil {
		return fmt.Errorf("unable to decode starsTransactionPeerUnsupported#95f2bfe4: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarsTransactionPeerUnsupported) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerUnsupported#95f2bfe4 to nil")
	}
	return nil
}

// StarsTransactionPeerAppStore represents TL type `starsTransactionPeerAppStore#b457b375`.
// Describes a Telegram Star¹ transaction with the App Store, used when purchasing
// Telegram Stars through the App Store.
//
// Links:
//  1. https://core.telegram.org/api/stars
//
// See https://core.telegram.org/constructor/starsTransactionPeerAppStore for reference.
type StarsTransactionPeerAppStore struct {
}

// StarsTransactionPeerAppStoreTypeID is TL type id of StarsTransactionPeerAppStore.
const StarsTransactionPeerAppStoreTypeID = 0xb457b375

// construct implements constructor of StarsTransactionPeerClass.
func (s StarsTransactionPeerAppStore) construct() StarsTransactionPeerClass { return &s }

// Ensuring interfaces in compile-time for StarsTransactionPeerAppStore.
var (
	_ bin.Encoder     = &StarsTransactionPeerAppStore{}
	_ bin.Decoder     = &StarsTransactionPeerAppStore{}
	_ bin.BareEncoder = &StarsTransactionPeerAppStore{}
	_ bin.BareDecoder = &StarsTransactionPeerAppStore{}

	_ StarsTransactionPeerClass = &StarsTransactionPeerAppStore{}
)

func (s *StarsTransactionPeerAppStore) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarsTransactionPeerAppStore) String() string {
	if s == nil {
		return "StarsTransactionPeerAppStore(nil)"
	}
	type Alias StarsTransactionPeerAppStore
	return fmt.Sprintf("StarsTransactionPeerAppStore%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarsTransactionPeerAppStore) TypeID() uint32 {
	return StarsTransactionPeerAppStoreTypeID
}

// TypeName returns name of type in TL schema.
func (*StarsTransactionPeerAppStore) TypeName() string {
	return "starsTransactionPeerAppStore"
}

// TypeInfo returns info about TL type.
func (s *StarsTransactionPeerAppStore) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starsTransactionPeerAppStore",
		ID:   StarsTransactionPeerAppStoreTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarsTransactionPeerAppStore) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerAppStore#b457b375 as nil")
	}
	b.PutID(StarsTransactionPeerAppStoreTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarsTransactionPeerAppStore) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerAppStore#b457b375 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarsTransactionPeerAppStore) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerAppStore#b457b375 to nil")
	}
	if err := b.ConsumeID(StarsTransactionPeerAppStoreTypeID); err != nil {
		return fmt.Errorf("unable to decode starsTransactionPeerAppStore#b457b375: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarsTransactionPeerAppStore) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerAppStore#b457b375 to nil")
	}
	return nil
}

// StarsTransactionPeerPlayMarket represents TL type `starsTransactionPeerPlayMarket#7b560a0b`.
// Describes a Telegram Star¹ transaction with the Play Store, used when purchasing
// Telegram Stars through the Play Store.
//
// Links:
//  1. https://core.telegram.org/api/stars
//
// See https://core.telegram.org/constructor/starsTransactionPeerPlayMarket for reference.
type StarsTransactionPeerPlayMarket struct {
}

// StarsTransactionPeerPlayMarketTypeID is TL type id of StarsTransactionPeerPlayMarket.
const StarsTransactionPeerPlayMarketTypeID = 0x7b560a0b

// construct implements constructor of StarsTransactionPeerClass.
func (s StarsTransactionPeerPlayMarket) construct() StarsTransactionPeerClass { return &s }

// Ensuring interfaces in compile-time for StarsTransactionPeerPlayMarket.
var (
	_ bin.Encoder     = &StarsTransactionPeerPlayMarket{}
	_ bin.Decoder     = &StarsTransactionPeerPlayMarket{}
	_ bin.BareEncoder = &StarsTransactionPeerPlayMarket{}
	_ bin.BareDecoder = &StarsTransactionPeerPlayMarket{}

	_ StarsTransactionPeerClass = &StarsTransactionPeerPlayMarket{}
)

func (s *StarsTransactionPeerPlayMarket) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarsTransactionPeerPlayMarket) String() string {
	if s == nil {
		return "StarsTransactionPeerPlayMarket(nil)"
	}
	type Alias StarsTransactionPeerPlayMarket
	return fmt.Sprintf("StarsTransactionPeerPlayMarket%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarsTransactionPeerPlayMarket) TypeID() uint32 {
	return StarsTransactionPeerPlayMarketTypeID
}

// TypeName returns name of type in TL schema.
func (*StarsTransactionPeerPlayMarket) TypeName() string {
	return "starsTransactionPeerPlayMarket"
}

// TypeInfo returns info about TL type.
func (s *StarsTransactionPeerPlayMarket) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starsTransactionPeerPlayMarket",
		ID:   StarsTransactionPeerPlayMarketTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarsTransactionPeerPlayMarket) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerPlayMarket#7b560a0b as nil")
	}
	b.PutID(StarsTransactionPeerPlayMarketTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarsTransactionPeerPlayMarket) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerPlayMarket#7b560a0b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarsTransactionPeerPlayMarket) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerPlayMarket#7b560a0b to nil")
	}
	if err := b.ConsumeID(StarsTransactionPeerPlayMarketTypeID); err != nil {
		return fmt.Errorf("unable to decode starsTransactionPeerPlayMarket#7b560a0b: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarsTransactionPeerPlayMarket) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerPlayMarket#7b560a0b to nil")
	}
	return nil
}

// StarsTransactionPeerPremiumBot represents TL type `starsTransactionPeerPremiumBot#250dbaf8`.
// Describes a Telegram Star¹ transaction made using @PremiumBot² (i.e. using the
// inputInvoiceStars³ flow described here »⁴).
//
// Links:
//  1. https://core.telegram.org/api/stars
//  2. https://t.me/premiumbot
//  3. https://core.telegram.org/constructor/inputInvoiceStars
//  4. https://core.telegram.org/api/stars#buying-or-gifting-stars
//
// See https://core.telegram.org/constructor/starsTransactionPeerPremiumBot for reference.
type StarsTransactionPeerPremiumBot struct {
}

// StarsTransactionPeerPremiumBotTypeID is TL type id of StarsTransactionPeerPremiumBot.
const StarsTransactionPeerPremiumBotTypeID = 0x250dbaf8

// construct implements constructor of StarsTransactionPeerClass.
func (s StarsTransactionPeerPremiumBot) construct() StarsTransactionPeerClass { return &s }

// Ensuring interfaces in compile-time for StarsTransactionPeerPremiumBot.
var (
	_ bin.Encoder     = &StarsTransactionPeerPremiumBot{}
	_ bin.Decoder     = &StarsTransactionPeerPremiumBot{}
	_ bin.BareEncoder = &StarsTransactionPeerPremiumBot{}
	_ bin.BareDecoder = &StarsTransactionPeerPremiumBot{}

	_ StarsTransactionPeerClass = &StarsTransactionPeerPremiumBot{}
)

func (s *StarsTransactionPeerPremiumBot) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarsTransactionPeerPremiumBot) String() string {
	if s == nil {
		return "StarsTransactionPeerPremiumBot(nil)"
	}
	type Alias StarsTransactionPeerPremiumBot
	return fmt.Sprintf("StarsTransactionPeerPremiumBot%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarsTransactionPeerPremiumBot) TypeID() uint32 {
	return StarsTransactionPeerPremiumBotTypeID
}

// TypeName returns name of type in TL schema.
func (*StarsTransactionPeerPremiumBot) TypeName() string {
	return "starsTransactionPeerPremiumBot"
}

// TypeInfo returns info about TL type.
func (s *StarsTransactionPeerPremiumBot) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starsTransactionPeerPremiumBot",
		ID:   StarsTransactionPeerPremiumBotTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarsTransactionPeerPremiumBot) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerPremiumBot#250dbaf8 as nil")
	}
	b.PutID(StarsTransactionPeerPremiumBotTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarsTransactionPeerPremiumBot) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerPremiumBot#250dbaf8 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarsTransactionPeerPremiumBot) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerPremiumBot#250dbaf8 to nil")
	}
	if err := b.ConsumeID(StarsTransactionPeerPremiumBotTypeID); err != nil {
		return fmt.Errorf("unable to decode starsTransactionPeerPremiumBot#250dbaf8: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarsTransactionPeerPremiumBot) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerPremiumBot#250dbaf8 to nil")
	}
	return nil
}

// StarsTransactionPeerFragment represents TL type `starsTransactionPeerFragment#e92fd902`.
// Describes a Telegram Star¹ transaction with Fragment², used when purchasing Telegram
// Stars through Fragment³.
//
// Links:
//  1. https://core.telegram.org/api/stars
//  2. https://fragment.com
//  3. https://fragment.com
//
// See https://core.telegram.org/constructor/starsTransactionPeerFragment for reference.
type StarsTransactionPeerFragment struct {
}

// StarsTransactionPeerFragmentTypeID is TL type id of StarsTransactionPeerFragment.
const StarsTransactionPeerFragmentTypeID = 0xe92fd902

// construct implements constructor of StarsTransactionPeerClass.
func (s StarsTransactionPeerFragment) construct() StarsTransactionPeerClass { return &s }

// Ensuring interfaces in compile-time for StarsTransactionPeerFragment.
var (
	_ bin.Encoder     = &StarsTransactionPeerFragment{}
	_ bin.Decoder     = &StarsTransactionPeerFragment{}
	_ bin.BareEncoder = &StarsTransactionPeerFragment{}
	_ bin.BareDecoder = &StarsTransactionPeerFragment{}

	_ StarsTransactionPeerClass = &StarsTransactionPeerFragment{}
)

func (s *StarsTransactionPeerFragment) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarsTransactionPeerFragment) String() string {
	if s == nil {
		return "StarsTransactionPeerFragment(nil)"
	}
	type Alias StarsTransactionPeerFragment
	return fmt.Sprintf("StarsTransactionPeerFragment%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarsTransactionPeerFragment) TypeID() uint32 {
	return StarsTransactionPeerFragmentTypeID
}

// TypeName returns name of type in TL schema.
func (*StarsTransactionPeerFragment) TypeName() string {
	return "starsTransactionPeerFragment"
}

// TypeInfo returns info about TL type.
func (s *StarsTransactionPeerFragment) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starsTransactionPeerFragment",
		ID:   StarsTransactionPeerFragmentTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarsTransactionPeerFragment) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerFragment#e92fd902 as nil")
	}
	b.PutID(StarsTransactionPeerFragmentTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarsTransactionPeerFragment) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerFragment#e92fd902 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarsTransactionPeerFragment) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerFragment#e92fd902 to nil")
	}
	if err := b.ConsumeID(StarsTransactionPeerFragmentTypeID); err != nil {
		return fmt.Errorf("unable to decode starsTransactionPeerFragment#e92fd902: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarsTransactionPeerFragment) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerFragment#e92fd902 to nil")
	}
	return nil
}

// StarsTransactionPeer represents TL type `starsTransactionPeer#d80da15d`.
// Describes a Telegram Star¹ transaction with another peer.
//
// Links:
//  1. https://core.telegram.org/api/stars
//
// See https://core.telegram.org/constructor/starsTransactionPeer for reference.
type StarsTransactionPeer struct {
	// The peer.
	Peer PeerClass
}

// StarsTransactionPeerTypeID is TL type id of StarsTransactionPeer.
const StarsTransactionPeerTypeID = 0xd80da15d

// construct implements constructor of StarsTransactionPeerClass.
func (s StarsTransactionPeer) construct() StarsTransactionPeerClass { return &s }

// Ensuring interfaces in compile-time for StarsTransactionPeer.
var (
	_ bin.Encoder     = &StarsTransactionPeer{}
	_ bin.Decoder     = &StarsTransactionPeer{}
	_ bin.BareEncoder = &StarsTransactionPeer{}
	_ bin.BareDecoder = &StarsTransactionPeer{}

	_ StarsTransactionPeerClass = &StarsTransactionPeer{}
)

func (s *StarsTransactionPeer) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarsTransactionPeer) String() string {
	if s == nil {
		return "StarsTransactionPeer(nil)"
	}
	type Alias StarsTransactionPeer
	return fmt.Sprintf("StarsTransactionPeer%+v", Alias(*s))
}

// FillFrom fills StarsTransactionPeer from given interface.
func (s *StarsTransactionPeer) FillFrom(from interface {
	GetPeer() (value PeerClass)
}) {
	s.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarsTransactionPeer) TypeID() uint32 {
	return StarsTransactionPeerTypeID
}

// TypeName returns name of type in TL schema.
func (*StarsTransactionPeer) TypeName() string {
	return "starsTransactionPeer"
}

// TypeInfo returns info about TL type.
func (s *StarsTransactionPeer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starsTransactionPeer",
		ID:   StarsTransactionPeerTypeID,
	}
	if s == nil {
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
func (s *StarsTransactionPeer) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeer#d80da15d as nil")
	}
	b.PutID(StarsTransactionPeerTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarsTransactionPeer) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeer#d80da15d as nil")
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode starsTransactionPeer#d80da15d: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starsTransactionPeer#d80da15d: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarsTransactionPeer) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeer#d80da15d to nil")
	}
	if err := b.ConsumeID(StarsTransactionPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode starsTransactionPeer#d80da15d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarsTransactionPeer) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeer#d80da15d to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode starsTransactionPeer#d80da15d: field peer: %w", err)
		}
		s.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *StarsTransactionPeer) GetPeer() (value PeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// StarsTransactionPeerAds represents TL type `starsTransactionPeerAds#60682812`.
// Describes a Telegram Star¹ transaction used to pay for Telegram ads as specified here
// »².
//
// Links:
//  1. https://core.telegram.org/api/stars
//  2. https://core.telegram.org/api/stars#paying-for-ads
//
// See https://core.telegram.org/constructor/starsTransactionPeerAds for reference.
type StarsTransactionPeerAds struct {
}

// StarsTransactionPeerAdsTypeID is TL type id of StarsTransactionPeerAds.
const StarsTransactionPeerAdsTypeID = 0x60682812

// construct implements constructor of StarsTransactionPeerClass.
func (s StarsTransactionPeerAds) construct() StarsTransactionPeerClass { return &s }

// Ensuring interfaces in compile-time for StarsTransactionPeerAds.
var (
	_ bin.Encoder     = &StarsTransactionPeerAds{}
	_ bin.Decoder     = &StarsTransactionPeerAds{}
	_ bin.BareEncoder = &StarsTransactionPeerAds{}
	_ bin.BareDecoder = &StarsTransactionPeerAds{}

	_ StarsTransactionPeerClass = &StarsTransactionPeerAds{}
)

func (s *StarsTransactionPeerAds) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarsTransactionPeerAds) String() string {
	if s == nil {
		return "StarsTransactionPeerAds(nil)"
	}
	type Alias StarsTransactionPeerAds
	return fmt.Sprintf("StarsTransactionPeerAds%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarsTransactionPeerAds) TypeID() uint32 {
	return StarsTransactionPeerAdsTypeID
}

// TypeName returns name of type in TL schema.
func (*StarsTransactionPeerAds) TypeName() string {
	return "starsTransactionPeerAds"
}

// TypeInfo returns info about TL type.
func (s *StarsTransactionPeerAds) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starsTransactionPeerAds",
		ID:   StarsTransactionPeerAdsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarsTransactionPeerAds) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerAds#60682812 as nil")
	}
	b.PutID(StarsTransactionPeerAdsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarsTransactionPeerAds) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerAds#60682812 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarsTransactionPeerAds) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerAds#60682812 to nil")
	}
	if err := b.ConsumeID(StarsTransactionPeerAdsTypeID); err != nil {
		return fmt.Errorf("unable to decode starsTransactionPeerAds#60682812: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarsTransactionPeerAds) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerAds#60682812 to nil")
	}
	return nil
}

// StarsTransactionPeerAPI represents TL type `starsTransactionPeerAPI#f9677aad`.
//
// See https://core.telegram.org/constructor/starsTransactionPeerAPI for reference.
type StarsTransactionPeerAPI struct {
}

// StarsTransactionPeerAPITypeID is TL type id of StarsTransactionPeerAPI.
const StarsTransactionPeerAPITypeID = 0xf9677aad

// construct implements constructor of StarsTransactionPeerClass.
func (s StarsTransactionPeerAPI) construct() StarsTransactionPeerClass { return &s }

// Ensuring interfaces in compile-time for StarsTransactionPeerAPI.
var (
	_ bin.Encoder     = &StarsTransactionPeerAPI{}
	_ bin.Decoder     = &StarsTransactionPeerAPI{}
	_ bin.BareEncoder = &StarsTransactionPeerAPI{}
	_ bin.BareDecoder = &StarsTransactionPeerAPI{}

	_ StarsTransactionPeerClass = &StarsTransactionPeerAPI{}
)

func (s *StarsTransactionPeerAPI) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarsTransactionPeerAPI) String() string {
	if s == nil {
		return "StarsTransactionPeerAPI(nil)"
	}
	type Alias StarsTransactionPeerAPI
	return fmt.Sprintf("StarsTransactionPeerAPI%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarsTransactionPeerAPI) TypeID() uint32 {
	return StarsTransactionPeerAPITypeID
}

// TypeName returns name of type in TL schema.
func (*StarsTransactionPeerAPI) TypeName() string {
	return "starsTransactionPeerAPI"
}

// TypeInfo returns info about TL type.
func (s *StarsTransactionPeerAPI) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starsTransactionPeerAPI",
		ID:   StarsTransactionPeerAPITypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarsTransactionPeerAPI) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerAPI#f9677aad as nil")
	}
	b.PutID(StarsTransactionPeerAPITypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarsTransactionPeerAPI) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransactionPeerAPI#f9677aad as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarsTransactionPeerAPI) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerAPI#f9677aad to nil")
	}
	if err := b.ConsumeID(StarsTransactionPeerAPITypeID); err != nil {
		return fmt.Errorf("unable to decode starsTransactionPeerAPI#f9677aad: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarsTransactionPeerAPI) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransactionPeerAPI#f9677aad to nil")
	}
	return nil
}

// StarsTransactionPeerClassName is schema name of StarsTransactionPeerClass.
const StarsTransactionPeerClassName = "StarsTransactionPeer"

// StarsTransactionPeerClass represents StarsTransactionPeer generic type.
//
// See https://core.telegram.org/type/StarsTransactionPeer for reference.
//
// Example:
//
//	g, err := tg.DecodeStarsTransactionPeer(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.StarsTransactionPeerUnsupported: // starsTransactionPeerUnsupported#95f2bfe4
//	case *tg.StarsTransactionPeerAppStore: // starsTransactionPeerAppStore#b457b375
//	case *tg.StarsTransactionPeerPlayMarket: // starsTransactionPeerPlayMarket#7b560a0b
//	case *tg.StarsTransactionPeerPremiumBot: // starsTransactionPeerPremiumBot#250dbaf8
//	case *tg.StarsTransactionPeerFragment: // starsTransactionPeerFragment#e92fd902
//	case *tg.StarsTransactionPeer: // starsTransactionPeer#d80da15d
//	case *tg.StarsTransactionPeerAds: // starsTransactionPeerAds#60682812
//	case *tg.StarsTransactionPeerAPI: // starsTransactionPeerAPI#f9677aad
//	default: panic(v)
//	}
type StarsTransactionPeerClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StarsTransactionPeerClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeStarsTransactionPeer implements binary de-serialization for StarsTransactionPeerClass.
func DecodeStarsTransactionPeer(buf *bin.Buffer) (StarsTransactionPeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StarsTransactionPeerUnsupportedTypeID:
		// Decoding starsTransactionPeerUnsupported#95f2bfe4.
		v := StarsTransactionPeerUnsupported{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarsTransactionPeerClass: %w", err)
		}
		return &v, nil
	case StarsTransactionPeerAppStoreTypeID:
		// Decoding starsTransactionPeerAppStore#b457b375.
		v := StarsTransactionPeerAppStore{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarsTransactionPeerClass: %w", err)
		}
		return &v, nil
	case StarsTransactionPeerPlayMarketTypeID:
		// Decoding starsTransactionPeerPlayMarket#7b560a0b.
		v := StarsTransactionPeerPlayMarket{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarsTransactionPeerClass: %w", err)
		}
		return &v, nil
	case StarsTransactionPeerPremiumBotTypeID:
		// Decoding starsTransactionPeerPremiumBot#250dbaf8.
		v := StarsTransactionPeerPremiumBot{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarsTransactionPeerClass: %w", err)
		}
		return &v, nil
	case StarsTransactionPeerFragmentTypeID:
		// Decoding starsTransactionPeerFragment#e92fd902.
		v := StarsTransactionPeerFragment{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarsTransactionPeerClass: %w", err)
		}
		return &v, nil
	case StarsTransactionPeerTypeID:
		// Decoding starsTransactionPeer#d80da15d.
		v := StarsTransactionPeer{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarsTransactionPeerClass: %w", err)
		}
		return &v, nil
	case StarsTransactionPeerAdsTypeID:
		// Decoding starsTransactionPeerAds#60682812.
		v := StarsTransactionPeerAds{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarsTransactionPeerClass: %w", err)
		}
		return &v, nil
	case StarsTransactionPeerAPITypeID:
		// Decoding starsTransactionPeerAPI#f9677aad.
		v := StarsTransactionPeerAPI{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarsTransactionPeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StarsTransactionPeerClass: %w", bin.NewUnexpectedID(id))
	}
}

// StarsTransactionPeer boxes the StarsTransactionPeerClass providing a helper.
type StarsTransactionPeerBox struct {
	StarsTransactionPeer StarsTransactionPeerClass
}

// Decode implements bin.Decoder for StarsTransactionPeerBox.
func (b *StarsTransactionPeerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StarsTransactionPeerBox to nil")
	}
	v, err := DecodeStarsTransactionPeer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StarsTransactionPeer = v
	return nil
}

// Encode implements bin.Encode for StarsTransactionPeerBox.
func (b *StarsTransactionPeerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StarsTransactionPeer == nil {
		return fmt.Errorf("unable to encode StarsTransactionPeerClass as nil")
	}
	return b.StarsTransactionPeer.Encode(buf)
}
