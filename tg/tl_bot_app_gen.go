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

// BotAppNotModified represents TL type `botAppNotModified#5da674b7`.
//
// See https://core.telegram.org/constructor/botAppNotModified for reference.
type BotAppNotModified struct {
}

// BotAppNotModifiedTypeID is TL type id of BotAppNotModified.
const BotAppNotModifiedTypeID = 0x5da674b7

// construct implements constructor of BotAppClass.
func (b BotAppNotModified) construct() BotAppClass { return &b }

// Ensuring interfaces in compile-time for BotAppNotModified.
var (
	_ bin.Encoder     = &BotAppNotModified{}
	_ bin.Decoder     = &BotAppNotModified{}
	_ bin.BareEncoder = &BotAppNotModified{}
	_ bin.BareDecoder = &BotAppNotModified{}

	_ BotAppClass = &BotAppNotModified{}
)

func (b *BotAppNotModified) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotAppNotModified) String() string {
	if b == nil {
		return "BotAppNotModified(nil)"
	}
	type Alias BotAppNotModified
	return fmt.Sprintf("BotAppNotModified%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotAppNotModified) TypeID() uint32 {
	return BotAppNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*BotAppNotModified) TypeName() string {
	return "botAppNotModified"
}

// TypeInfo returns info about TL type.
func (b *BotAppNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botAppNotModified",
		ID:   BotAppNotModifiedTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotAppNotModified) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botAppNotModified#5da674b7 as nil")
	}
	buf.PutID(BotAppNotModifiedTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotAppNotModified) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botAppNotModified#5da674b7 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotAppNotModified) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botAppNotModified#5da674b7 to nil")
	}
	if err := buf.ConsumeID(BotAppNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode botAppNotModified#5da674b7: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotAppNotModified) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botAppNotModified#5da674b7 to nil")
	}
	return nil
}

// BotApp represents TL type `botApp#95fcd1d6`.
//
// See https://core.telegram.org/constructor/botApp for reference.
type BotApp struct {
	// Flags field of BotApp.
	Flags bin.Fields
	// ID field of BotApp.
	ID int64
	// AccessHash field of BotApp.
	AccessHash int64
	// ShortName field of BotApp.
	ShortName string
	// Title field of BotApp.
	Title string
	// Description field of BotApp.
	Description string
	// Photo field of BotApp.
	Photo PhotoClass
	// Document field of BotApp.
	//
	// Use SetDocument and GetDocument helpers.
	Document DocumentClass
	// Hash field of BotApp.
	Hash int64
}

// BotAppTypeID is TL type id of BotApp.
const BotAppTypeID = 0x95fcd1d6

// construct implements constructor of BotAppClass.
func (b BotApp) construct() BotAppClass { return &b }

// Ensuring interfaces in compile-time for BotApp.
var (
	_ bin.Encoder     = &BotApp{}
	_ bin.Decoder     = &BotApp{}
	_ bin.BareEncoder = &BotApp{}
	_ bin.BareDecoder = &BotApp{}

	_ BotAppClass = &BotApp{}
)

func (b *BotApp) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.ID == 0) {
		return false
	}
	if !(b.AccessHash == 0) {
		return false
	}
	if !(b.ShortName == "") {
		return false
	}
	if !(b.Title == "") {
		return false
	}
	if !(b.Description == "") {
		return false
	}
	if !(b.Photo == nil) {
		return false
	}
	if !(b.Document == nil) {
		return false
	}
	if !(b.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotApp) String() string {
	if b == nil {
		return "BotApp(nil)"
	}
	type Alias BotApp
	return fmt.Sprintf("BotApp%+v", Alias(*b))
}

// FillFrom fills BotApp from given interface.
func (b *BotApp) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetShortName() (value string)
	GetTitle() (value string)
	GetDescription() (value string)
	GetPhoto() (value PhotoClass)
	GetDocument() (value DocumentClass, ok bool)
	GetHash() (value int64)
}) {
	b.ID = from.GetID()
	b.AccessHash = from.GetAccessHash()
	b.ShortName = from.GetShortName()
	b.Title = from.GetTitle()
	b.Description = from.GetDescription()
	b.Photo = from.GetPhoto()
	if val, ok := from.GetDocument(); ok {
		b.Document = val
	}

	b.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotApp) TypeID() uint32 {
	return BotAppTypeID
}

// TypeName returns name of type in TL schema.
func (*BotApp) TypeName() string {
	return "botApp"
}

// TypeInfo returns info about TL type.
func (b *BotApp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botApp",
		ID:   BotAppTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "Document",
			SchemaName: "document",
			Null:       !b.Flags.Has(0),
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (b *BotApp) SetFlags() {
	if !(b.Document == nil) {
		b.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (b *BotApp) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botApp#95fcd1d6 as nil")
	}
	buf.PutID(BotAppTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotApp) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botApp#95fcd1d6 as nil")
	}
	b.SetFlags()
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botApp#95fcd1d6: field flags: %w", err)
	}
	buf.PutLong(b.ID)
	buf.PutLong(b.AccessHash)
	buf.PutString(b.ShortName)
	buf.PutString(b.Title)
	buf.PutString(b.Description)
	if b.Photo == nil {
		return fmt.Errorf("unable to encode botApp#95fcd1d6: field photo is nil")
	}
	if err := b.Photo.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botApp#95fcd1d6: field photo: %w", err)
	}
	if b.Flags.Has(0) {
		if b.Document == nil {
			return fmt.Errorf("unable to encode botApp#95fcd1d6: field document is nil")
		}
		if err := b.Document.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botApp#95fcd1d6: field document: %w", err)
		}
	}
	buf.PutLong(b.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (b *BotApp) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botApp#95fcd1d6 to nil")
	}
	if err := buf.ConsumeID(BotAppTypeID); err != nil {
		return fmt.Errorf("unable to decode botApp#95fcd1d6: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotApp) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botApp#95fcd1d6 to nil")
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botApp#95fcd1d6: field flags: %w", err)
		}
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode botApp#95fcd1d6: field id: %w", err)
		}
		b.ID = value
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode botApp#95fcd1d6: field access_hash: %w", err)
		}
		b.AccessHash = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botApp#95fcd1d6: field short_name: %w", err)
		}
		b.ShortName = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botApp#95fcd1d6: field title: %w", err)
		}
		b.Title = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botApp#95fcd1d6: field description: %w", err)
		}
		b.Description = value
	}
	{
		value, err := DecodePhoto(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botApp#95fcd1d6: field photo: %w", err)
		}
		b.Photo = value
	}
	if b.Flags.Has(0) {
		value, err := DecodeDocument(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botApp#95fcd1d6: field document: %w", err)
		}
		b.Document = value
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode botApp#95fcd1d6: field hash: %w", err)
		}
		b.Hash = value
	}
	return nil
}

// GetID returns value of ID field.
func (b *BotApp) GetID() (value int64) {
	if b == nil {
		return
	}
	return b.ID
}

// GetAccessHash returns value of AccessHash field.
func (b *BotApp) GetAccessHash() (value int64) {
	if b == nil {
		return
	}
	return b.AccessHash
}

// GetShortName returns value of ShortName field.
func (b *BotApp) GetShortName() (value string) {
	if b == nil {
		return
	}
	return b.ShortName
}

// GetTitle returns value of Title field.
func (b *BotApp) GetTitle() (value string) {
	if b == nil {
		return
	}
	return b.Title
}

// GetDescription returns value of Description field.
func (b *BotApp) GetDescription() (value string) {
	if b == nil {
		return
	}
	return b.Description
}

// GetPhoto returns value of Photo field.
func (b *BotApp) GetPhoto() (value PhotoClass) {
	if b == nil {
		return
	}
	return b.Photo
}

// SetDocument sets value of Document conditional field.
func (b *BotApp) SetDocument(value DocumentClass) {
	b.Flags.Set(0)
	b.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (b *BotApp) GetDocument() (value DocumentClass, ok bool) {
	if b == nil {
		return
	}
	if !b.Flags.Has(0) {
		return value, false
	}
	return b.Document, true
}

// GetHash returns value of Hash field.
func (b *BotApp) GetHash() (value int64) {
	if b == nil {
		return
	}
	return b.Hash
}

// BotAppClassName is schema name of BotAppClass.
const BotAppClassName = "BotApp"

// BotAppClass represents BotApp generic type.
//
// See https://core.telegram.org/type/BotApp for reference.
//
// Example:
//
//	g, err := tg.DecodeBotApp(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.BotAppNotModified: // botAppNotModified#5da674b7
//	case *tg.BotApp: // botApp#95fcd1d6
//	default: panic(v)
//	}
type BotAppClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() BotAppClass

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

	// AsModified tries to map BotAppClass to BotApp.
	AsModified() (*BotApp, bool)
}

// AsInputBotAppID tries to map BotApp to InputBotAppID.
func (b *BotApp) AsInputBotAppID() *InputBotAppID {
	value := new(InputBotAppID)
	value.ID = b.GetID()
	value.AccessHash = b.GetAccessHash()
	value.AccessHash = b.GetHash()

	return value
}

// AsModified tries to map BotAppNotModified to BotApp.
func (b *BotAppNotModified) AsModified() (*BotApp, bool) {
	return nil, false
}

// AsModified tries to map BotApp to BotApp.
func (b *BotApp) AsModified() (*BotApp, bool) {
	return b, true
}

// DecodeBotApp implements binary de-serialization for BotAppClass.
func DecodeBotApp(buf *bin.Buffer) (BotAppClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BotAppNotModifiedTypeID:
		// Decoding botAppNotModified#5da674b7.
		v := BotAppNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotAppClass: %w", err)
		}
		return &v, nil
	case BotAppTypeID:
		// Decoding botApp#95fcd1d6.
		v := BotApp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotAppClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BotAppClass: %w", bin.NewUnexpectedID(id))
	}
}

// BotApp boxes the BotAppClass providing a helper.
type BotAppBox struct {
	BotApp BotAppClass
}

// Decode implements bin.Decoder for BotAppBox.
func (b *BotAppBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BotAppBox to nil")
	}
	v, err := DecodeBotApp(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BotApp = v
	return nil
}

// Encode implements bin.Encode for BotAppBox.
func (b *BotAppBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BotApp == nil {
		return fmt.Errorf("unable to encode BotAppClass as nil")
	}
	return b.BotApp.Encode(buf)
}