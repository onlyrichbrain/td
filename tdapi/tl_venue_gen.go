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

// Venue represents TL type `venue#3fcd1af9`.
type Venue struct {
	// Venue location; as defined by the sender
	Location Location
	// Venue name; as defined by the sender
	Title string
	// Venue address; as defined by the sender
	Address string
	// Provider of the venue database; as defined by the sender. Currently, only "foursquare"
	// and "gplaces" (Google Places) need to be supported
	Provider string
	// Identifier of the venue in the provider database; as defined by the sender
	ID string
	// Type of the venue in the provider database; as defined by the sender
	Type string
}

// VenueTypeID is TL type id of Venue.
const VenueTypeID = 0x3fcd1af9

// Ensuring interfaces in compile-time for Venue.
var (
	_ bin.Encoder     = &Venue{}
	_ bin.Decoder     = &Venue{}
	_ bin.BareEncoder = &Venue{}
	_ bin.BareDecoder = &Venue{}
)

func (v *Venue) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Location.Zero()) {
		return false
	}
	if !(v.Title == "") {
		return false
	}
	if !(v.Address == "") {
		return false
	}
	if !(v.Provider == "") {
		return false
	}
	if !(v.ID == "") {
		return false
	}
	if !(v.Type == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *Venue) String() string {
	if v == nil {
		return "Venue(nil)"
	}
	type Alias Venue
	return fmt.Sprintf("Venue%+v", Alias(*v))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Venue) TypeID() uint32 {
	return VenueTypeID
}

// TypeName returns name of type in TL schema.
func (*Venue) TypeName() string {
	return "venue"
}

// TypeInfo returns info about TL type.
func (v *Venue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "venue",
		ID:   VenueTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Address",
			SchemaName: "address",
		},
		{
			Name:       "Provider",
			SchemaName: "provider",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *Venue) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode venue#3fcd1af9 as nil")
	}
	b.PutID(VenueTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *Venue) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode venue#3fcd1af9 as nil")
	}
	if err := v.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode venue#3fcd1af9: field location: %w", err)
	}
	b.PutString(v.Title)
	b.PutString(v.Address)
	b.PutString(v.Provider)
	b.PutString(v.ID)
	b.PutString(v.Type)
	return nil
}

// Decode implements bin.Decoder.
func (v *Venue) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode venue#3fcd1af9 to nil")
	}
	if err := b.ConsumeID(VenueTypeID); err != nil {
		return fmt.Errorf("unable to decode venue#3fcd1af9: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *Venue) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode venue#3fcd1af9 to nil")
	}
	{
		if err := v.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode venue#3fcd1af9: field location: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode venue#3fcd1af9: field title: %w", err)
		}
		v.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode venue#3fcd1af9: field address: %w", err)
		}
		v.Address = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode venue#3fcd1af9: field provider: %w", err)
		}
		v.Provider = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode venue#3fcd1af9: field id: %w", err)
		}
		v.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode venue#3fcd1af9: field type: %w", err)
		}
		v.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (v *Venue) EncodeTDLibJSON(b tdjson.Encoder) error {
	if v == nil {
		return fmt.Errorf("can't encode venue#3fcd1af9 as nil")
	}
	b.ObjStart()
	b.PutID("venue")
	b.Comma()
	b.FieldStart("location")
	if err := v.Location.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode venue#3fcd1af9: field location: %w", err)
	}
	b.Comma()
	b.FieldStart("title")
	b.PutString(v.Title)
	b.Comma()
	b.FieldStart("address")
	b.PutString(v.Address)
	b.Comma()
	b.FieldStart("provider")
	b.PutString(v.Provider)
	b.Comma()
	b.FieldStart("id")
	b.PutString(v.ID)
	b.Comma()
	b.FieldStart("type")
	b.PutString(v.Type)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (v *Venue) DecodeTDLibJSON(b tdjson.Decoder) error {
	if v == nil {
		return fmt.Errorf("can't decode venue#3fcd1af9 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("venue"); err != nil {
				return fmt.Errorf("unable to decode venue#3fcd1af9: %w", err)
			}
		case "location":
			if err := v.Location.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode venue#3fcd1af9: field location: %w", err)
			}
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode venue#3fcd1af9: field title: %w", err)
			}
			v.Title = value
		case "address":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode venue#3fcd1af9: field address: %w", err)
			}
			v.Address = value
		case "provider":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode venue#3fcd1af9: field provider: %w", err)
			}
			v.Provider = value
		case "id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode venue#3fcd1af9: field id: %w", err)
			}
			v.ID = value
		case "type":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode venue#3fcd1af9: field type: %w", err)
			}
			v.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLocation returns value of Location field.
func (v *Venue) GetLocation() (value Location) {
	if v == nil {
		return
	}
	return v.Location
}

// GetTitle returns value of Title field.
func (v *Venue) GetTitle() (value string) {
	if v == nil {
		return
	}
	return v.Title
}

// GetAddress returns value of Address field.
func (v *Venue) GetAddress() (value string) {
	if v == nil {
		return
	}
	return v.Address
}

// GetProvider returns value of Provider field.
func (v *Venue) GetProvider() (value string) {
	if v == nil {
		return
	}
	return v.Provider
}

// GetID returns value of ID field.
func (v *Venue) GetID() (value string) {
	if v == nil {
		return
	}
	return v.ID
}

// GetType returns value of Type field.
func (v *Venue) GetType() (value string) {
	if v == nil {
		return
	}
	return v.Type
}
