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

// ConnectionStateWaitingForNetwork represents TL type `connectionStateWaitingForNetwork#650dd758`.
type ConnectionStateWaitingForNetwork struct {
}

// ConnectionStateWaitingForNetworkTypeID is TL type id of ConnectionStateWaitingForNetwork.
const ConnectionStateWaitingForNetworkTypeID = 0x650dd758

// construct implements constructor of ConnectionStateClass.
func (c ConnectionStateWaitingForNetwork) construct() ConnectionStateClass { return &c }

// Ensuring interfaces in compile-time for ConnectionStateWaitingForNetwork.
var (
	_ bin.Encoder     = &ConnectionStateWaitingForNetwork{}
	_ bin.Decoder     = &ConnectionStateWaitingForNetwork{}
	_ bin.BareEncoder = &ConnectionStateWaitingForNetwork{}
	_ bin.BareDecoder = &ConnectionStateWaitingForNetwork{}

	_ ConnectionStateClass = &ConnectionStateWaitingForNetwork{}
)

func (c *ConnectionStateWaitingForNetwork) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ConnectionStateWaitingForNetwork) String() string {
	if c == nil {
		return "ConnectionStateWaitingForNetwork(nil)"
	}
	type Alias ConnectionStateWaitingForNetwork
	return fmt.Sprintf("ConnectionStateWaitingForNetwork%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ConnectionStateWaitingForNetwork) TypeID() uint32 {
	return ConnectionStateWaitingForNetworkTypeID
}

// TypeName returns name of type in TL schema.
func (*ConnectionStateWaitingForNetwork) TypeName() string {
	return "connectionStateWaitingForNetwork"
}

// TypeInfo returns info about TL type.
func (c *ConnectionStateWaitingForNetwork) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "connectionStateWaitingForNetwork",
		ID:   ConnectionStateWaitingForNetworkTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ConnectionStateWaitingForNetwork) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateWaitingForNetwork#650dd758 as nil")
	}
	b.PutID(ConnectionStateWaitingForNetworkTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ConnectionStateWaitingForNetwork) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateWaitingForNetwork#650dd758 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ConnectionStateWaitingForNetwork) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateWaitingForNetwork#650dd758 to nil")
	}
	if err := b.ConsumeID(ConnectionStateWaitingForNetworkTypeID); err != nil {
		return fmt.Errorf("unable to decode connectionStateWaitingForNetwork#650dd758: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ConnectionStateWaitingForNetwork) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateWaitingForNetwork#650dd758 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ConnectionStateWaitingForNetwork) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateWaitingForNetwork#650dd758 as nil")
	}
	b.ObjStart()
	b.PutID("connectionStateWaitingForNetwork")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ConnectionStateWaitingForNetwork) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateWaitingForNetwork#650dd758 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("connectionStateWaitingForNetwork"); err != nil {
				return fmt.Errorf("unable to decode connectionStateWaitingForNetwork#650dd758: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ConnectionStateConnectingToProxy represents TL type `connectionStateConnectingToProxy#fa721359`.
type ConnectionStateConnectingToProxy struct {
}

// ConnectionStateConnectingToProxyTypeID is TL type id of ConnectionStateConnectingToProxy.
const ConnectionStateConnectingToProxyTypeID = 0xfa721359

// construct implements constructor of ConnectionStateClass.
func (c ConnectionStateConnectingToProxy) construct() ConnectionStateClass { return &c }

// Ensuring interfaces in compile-time for ConnectionStateConnectingToProxy.
var (
	_ bin.Encoder     = &ConnectionStateConnectingToProxy{}
	_ bin.Decoder     = &ConnectionStateConnectingToProxy{}
	_ bin.BareEncoder = &ConnectionStateConnectingToProxy{}
	_ bin.BareDecoder = &ConnectionStateConnectingToProxy{}

	_ ConnectionStateClass = &ConnectionStateConnectingToProxy{}
)

func (c *ConnectionStateConnectingToProxy) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ConnectionStateConnectingToProxy) String() string {
	if c == nil {
		return "ConnectionStateConnectingToProxy(nil)"
	}
	type Alias ConnectionStateConnectingToProxy
	return fmt.Sprintf("ConnectionStateConnectingToProxy%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ConnectionStateConnectingToProxy) TypeID() uint32 {
	return ConnectionStateConnectingToProxyTypeID
}

// TypeName returns name of type in TL schema.
func (*ConnectionStateConnectingToProxy) TypeName() string {
	return "connectionStateConnectingToProxy"
}

// TypeInfo returns info about TL type.
func (c *ConnectionStateConnectingToProxy) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "connectionStateConnectingToProxy",
		ID:   ConnectionStateConnectingToProxyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ConnectionStateConnectingToProxy) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateConnectingToProxy#fa721359 as nil")
	}
	b.PutID(ConnectionStateConnectingToProxyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ConnectionStateConnectingToProxy) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateConnectingToProxy#fa721359 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ConnectionStateConnectingToProxy) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateConnectingToProxy#fa721359 to nil")
	}
	if err := b.ConsumeID(ConnectionStateConnectingToProxyTypeID); err != nil {
		return fmt.Errorf("unable to decode connectionStateConnectingToProxy#fa721359: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ConnectionStateConnectingToProxy) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateConnectingToProxy#fa721359 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ConnectionStateConnectingToProxy) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateConnectingToProxy#fa721359 as nil")
	}
	b.ObjStart()
	b.PutID("connectionStateConnectingToProxy")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ConnectionStateConnectingToProxy) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateConnectingToProxy#fa721359 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("connectionStateConnectingToProxy"); err != nil {
				return fmt.Errorf("unable to decode connectionStateConnectingToProxy#fa721359: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ConnectionStateConnecting represents TL type `connectionStateConnecting#b29bfa62`.
type ConnectionStateConnecting struct {
}

// ConnectionStateConnectingTypeID is TL type id of ConnectionStateConnecting.
const ConnectionStateConnectingTypeID = 0xb29bfa62

// construct implements constructor of ConnectionStateClass.
func (c ConnectionStateConnecting) construct() ConnectionStateClass { return &c }

// Ensuring interfaces in compile-time for ConnectionStateConnecting.
var (
	_ bin.Encoder     = &ConnectionStateConnecting{}
	_ bin.Decoder     = &ConnectionStateConnecting{}
	_ bin.BareEncoder = &ConnectionStateConnecting{}
	_ bin.BareDecoder = &ConnectionStateConnecting{}

	_ ConnectionStateClass = &ConnectionStateConnecting{}
)

func (c *ConnectionStateConnecting) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ConnectionStateConnecting) String() string {
	if c == nil {
		return "ConnectionStateConnecting(nil)"
	}
	type Alias ConnectionStateConnecting
	return fmt.Sprintf("ConnectionStateConnecting%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ConnectionStateConnecting) TypeID() uint32 {
	return ConnectionStateConnectingTypeID
}

// TypeName returns name of type in TL schema.
func (*ConnectionStateConnecting) TypeName() string {
	return "connectionStateConnecting"
}

// TypeInfo returns info about TL type.
func (c *ConnectionStateConnecting) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "connectionStateConnecting",
		ID:   ConnectionStateConnectingTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ConnectionStateConnecting) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateConnecting#b29bfa62 as nil")
	}
	b.PutID(ConnectionStateConnectingTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ConnectionStateConnecting) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateConnecting#b29bfa62 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ConnectionStateConnecting) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateConnecting#b29bfa62 to nil")
	}
	if err := b.ConsumeID(ConnectionStateConnectingTypeID); err != nil {
		return fmt.Errorf("unable to decode connectionStateConnecting#b29bfa62: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ConnectionStateConnecting) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateConnecting#b29bfa62 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ConnectionStateConnecting) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateConnecting#b29bfa62 as nil")
	}
	b.ObjStart()
	b.PutID("connectionStateConnecting")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ConnectionStateConnecting) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateConnecting#b29bfa62 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("connectionStateConnecting"); err != nil {
				return fmt.Errorf("unable to decode connectionStateConnecting#b29bfa62: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ConnectionStateUpdating represents TL type `connectionStateUpdating#f4c9c2b7`.
type ConnectionStateUpdating struct {
}

// ConnectionStateUpdatingTypeID is TL type id of ConnectionStateUpdating.
const ConnectionStateUpdatingTypeID = 0xf4c9c2b7

// construct implements constructor of ConnectionStateClass.
func (c ConnectionStateUpdating) construct() ConnectionStateClass { return &c }

// Ensuring interfaces in compile-time for ConnectionStateUpdating.
var (
	_ bin.Encoder     = &ConnectionStateUpdating{}
	_ bin.Decoder     = &ConnectionStateUpdating{}
	_ bin.BareEncoder = &ConnectionStateUpdating{}
	_ bin.BareDecoder = &ConnectionStateUpdating{}

	_ ConnectionStateClass = &ConnectionStateUpdating{}
)

func (c *ConnectionStateUpdating) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ConnectionStateUpdating) String() string {
	if c == nil {
		return "ConnectionStateUpdating(nil)"
	}
	type Alias ConnectionStateUpdating
	return fmt.Sprintf("ConnectionStateUpdating%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ConnectionStateUpdating) TypeID() uint32 {
	return ConnectionStateUpdatingTypeID
}

// TypeName returns name of type in TL schema.
func (*ConnectionStateUpdating) TypeName() string {
	return "connectionStateUpdating"
}

// TypeInfo returns info about TL type.
func (c *ConnectionStateUpdating) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "connectionStateUpdating",
		ID:   ConnectionStateUpdatingTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ConnectionStateUpdating) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateUpdating#f4c9c2b7 as nil")
	}
	b.PutID(ConnectionStateUpdatingTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ConnectionStateUpdating) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateUpdating#f4c9c2b7 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ConnectionStateUpdating) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateUpdating#f4c9c2b7 to nil")
	}
	if err := b.ConsumeID(ConnectionStateUpdatingTypeID); err != nil {
		return fmt.Errorf("unable to decode connectionStateUpdating#f4c9c2b7: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ConnectionStateUpdating) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateUpdating#f4c9c2b7 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ConnectionStateUpdating) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateUpdating#f4c9c2b7 as nil")
	}
	b.ObjStart()
	b.PutID("connectionStateUpdating")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ConnectionStateUpdating) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateUpdating#f4c9c2b7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("connectionStateUpdating"); err != nil {
				return fmt.Errorf("unable to decode connectionStateUpdating#f4c9c2b7: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ConnectionStateReady represents TL type `connectionStateReady#2e5b4ec`.
type ConnectionStateReady struct {
}

// ConnectionStateReadyTypeID is TL type id of ConnectionStateReady.
const ConnectionStateReadyTypeID = 0x2e5b4ec

// construct implements constructor of ConnectionStateClass.
func (c ConnectionStateReady) construct() ConnectionStateClass { return &c }

// Ensuring interfaces in compile-time for ConnectionStateReady.
var (
	_ bin.Encoder     = &ConnectionStateReady{}
	_ bin.Decoder     = &ConnectionStateReady{}
	_ bin.BareEncoder = &ConnectionStateReady{}
	_ bin.BareDecoder = &ConnectionStateReady{}

	_ ConnectionStateClass = &ConnectionStateReady{}
)

func (c *ConnectionStateReady) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ConnectionStateReady) String() string {
	if c == nil {
		return "ConnectionStateReady(nil)"
	}
	type Alias ConnectionStateReady
	return fmt.Sprintf("ConnectionStateReady%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ConnectionStateReady) TypeID() uint32 {
	return ConnectionStateReadyTypeID
}

// TypeName returns name of type in TL schema.
func (*ConnectionStateReady) TypeName() string {
	return "connectionStateReady"
}

// TypeInfo returns info about TL type.
func (c *ConnectionStateReady) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "connectionStateReady",
		ID:   ConnectionStateReadyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ConnectionStateReady) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateReady#2e5b4ec as nil")
	}
	b.PutID(ConnectionStateReadyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ConnectionStateReady) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateReady#2e5b4ec as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ConnectionStateReady) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateReady#2e5b4ec to nil")
	}
	if err := b.ConsumeID(ConnectionStateReadyTypeID); err != nil {
		return fmt.Errorf("unable to decode connectionStateReady#2e5b4ec: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ConnectionStateReady) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateReady#2e5b4ec to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ConnectionStateReady) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode connectionStateReady#2e5b4ec as nil")
	}
	b.ObjStart()
	b.PutID("connectionStateReady")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ConnectionStateReady) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode connectionStateReady#2e5b4ec to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("connectionStateReady"); err != nil {
				return fmt.Errorf("unable to decode connectionStateReady#2e5b4ec: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ConnectionStateClassName is schema name of ConnectionStateClass.
const ConnectionStateClassName = "ConnectionState"

// ConnectionStateClass represents ConnectionState generic type.
//
// Example:
//  g, err := tdapi.DecodeConnectionState(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.ConnectionStateWaitingForNetwork: // connectionStateWaitingForNetwork#650dd758
//  case *tdapi.ConnectionStateConnectingToProxy: // connectionStateConnectingToProxy#fa721359
//  case *tdapi.ConnectionStateConnecting: // connectionStateConnecting#b29bfa62
//  case *tdapi.ConnectionStateUpdating: // connectionStateUpdating#f4c9c2b7
//  case *tdapi.ConnectionStateReady: // connectionStateReady#2e5b4ec
//  default: panic(v)
//  }
type ConnectionStateClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ConnectionStateClass

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

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeConnectionState implements binary de-serialization for ConnectionStateClass.
func DecodeConnectionState(buf *bin.Buffer) (ConnectionStateClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ConnectionStateWaitingForNetworkTypeID:
		// Decoding connectionStateWaitingForNetwork#650dd758.
		v := ConnectionStateWaitingForNetwork{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", err)
		}
		return &v, nil
	case ConnectionStateConnectingToProxyTypeID:
		// Decoding connectionStateConnectingToProxy#fa721359.
		v := ConnectionStateConnectingToProxy{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", err)
		}
		return &v, nil
	case ConnectionStateConnectingTypeID:
		// Decoding connectionStateConnecting#b29bfa62.
		v := ConnectionStateConnecting{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", err)
		}
		return &v, nil
	case ConnectionStateUpdatingTypeID:
		// Decoding connectionStateUpdating#f4c9c2b7.
		v := ConnectionStateUpdating{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", err)
		}
		return &v, nil
	case ConnectionStateReadyTypeID:
		// Decoding connectionStateReady#2e5b4ec.
		v := ConnectionStateReady{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONConnectionState implements binary de-serialization for ConnectionStateClass.
func DecodeTDLibJSONConnectionState(buf tdjson.Decoder) (ConnectionStateClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "connectionStateWaitingForNetwork":
		// Decoding connectionStateWaitingForNetwork#650dd758.
		v := ConnectionStateWaitingForNetwork{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", err)
		}
		return &v, nil
	case "connectionStateConnectingToProxy":
		// Decoding connectionStateConnectingToProxy#fa721359.
		v := ConnectionStateConnectingToProxy{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", err)
		}
		return &v, nil
	case "connectionStateConnecting":
		// Decoding connectionStateConnecting#b29bfa62.
		v := ConnectionStateConnecting{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", err)
		}
		return &v, nil
	case "connectionStateUpdating":
		// Decoding connectionStateUpdating#f4c9c2b7.
		v := ConnectionStateUpdating{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", err)
		}
		return &v, nil
	case "connectionStateReady":
		// Decoding connectionStateReady#2e5b4ec.
		v := ConnectionStateReady{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ConnectionStateClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// ConnectionState boxes the ConnectionStateClass providing a helper.
type ConnectionStateBox struct {
	ConnectionState ConnectionStateClass
}

// Decode implements bin.Decoder for ConnectionStateBox.
func (b *ConnectionStateBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ConnectionStateBox to nil")
	}
	v, err := DecodeConnectionState(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ConnectionState = v
	return nil
}

// Encode implements bin.Encode for ConnectionStateBox.
func (b *ConnectionStateBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ConnectionState == nil {
		return fmt.Errorf("unable to encode ConnectionStateClass as nil")
	}
	return b.ConnectionState.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for ConnectionStateBox.
func (b *ConnectionStateBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode ConnectionStateBox to nil")
	}
	v, err := DecodeTDLibJSONConnectionState(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ConnectionState = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for ConnectionStateBox.
func (b *ConnectionStateBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.ConnectionState == nil {
		return fmt.Errorf("unable to encode ConnectionStateClass as nil")
	}
	return b.ConnectionState.EncodeTDLibJSON(buf)
}
