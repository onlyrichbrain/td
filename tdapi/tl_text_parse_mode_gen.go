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

// TextParseModeMarkdown represents TL type `textParseModeMarkdown#157648bf`.
type TextParseModeMarkdown struct {
	// Version of the parser: 0 or 1 - Telegram Bot API "Markdown" parse mode, 2 - Telegram
	// Bot API "MarkdownV2" parse mode
	Version int32
}

// TextParseModeMarkdownTypeID is TL type id of TextParseModeMarkdown.
const TextParseModeMarkdownTypeID = 0x157648bf

// construct implements constructor of TextParseModeClass.
func (t TextParseModeMarkdown) construct() TextParseModeClass { return &t }

// Ensuring interfaces in compile-time for TextParseModeMarkdown.
var (
	_ bin.Encoder     = &TextParseModeMarkdown{}
	_ bin.Decoder     = &TextParseModeMarkdown{}
	_ bin.BareEncoder = &TextParseModeMarkdown{}
	_ bin.BareDecoder = &TextParseModeMarkdown{}

	_ TextParseModeClass = &TextParseModeMarkdown{}
)

func (t *TextParseModeMarkdown) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Version == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TextParseModeMarkdown) String() string {
	if t == nil {
		return "TextParseModeMarkdown(nil)"
	}
	type Alias TextParseModeMarkdown
	return fmt.Sprintf("TextParseModeMarkdown%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TextParseModeMarkdown) TypeID() uint32 {
	return TextParseModeMarkdownTypeID
}

// TypeName returns name of type in TL schema.
func (*TextParseModeMarkdown) TypeName() string {
	return "textParseModeMarkdown"
}

// TypeInfo returns info about TL type.
func (t *TextParseModeMarkdown) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "textParseModeMarkdown",
		ID:   TextParseModeMarkdownTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Version",
			SchemaName: "version",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TextParseModeMarkdown) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textParseModeMarkdown#157648bf as nil")
	}
	b.PutID(TextParseModeMarkdownTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TextParseModeMarkdown) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textParseModeMarkdown#157648bf as nil")
	}
	b.PutInt32(t.Version)
	return nil
}

// Decode implements bin.Decoder.
func (t *TextParseModeMarkdown) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textParseModeMarkdown#157648bf to nil")
	}
	if err := b.ConsumeID(TextParseModeMarkdownTypeID); err != nil {
		return fmt.Errorf("unable to decode textParseModeMarkdown#157648bf: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TextParseModeMarkdown) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textParseModeMarkdown#157648bf to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode textParseModeMarkdown#157648bf: field version: %w", err)
		}
		t.Version = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TextParseModeMarkdown) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode textParseModeMarkdown#157648bf as nil")
	}
	b.ObjStart()
	b.PutID("textParseModeMarkdown")
	b.Comma()
	b.FieldStart("version")
	b.PutInt32(t.Version)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TextParseModeMarkdown) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode textParseModeMarkdown#157648bf to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("textParseModeMarkdown"); err != nil {
				return fmt.Errorf("unable to decode textParseModeMarkdown#157648bf: %w", err)
			}
		case "version":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode textParseModeMarkdown#157648bf: field version: %w", err)
			}
			t.Version = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetVersion returns value of Version field.
func (t *TextParseModeMarkdown) GetVersion() (value int32) {
	if t == nil {
		return
	}
	return t.Version
}

// TextParseModeHTML represents TL type `textParseModeHTML#62f4c5f3`.
type TextParseModeHTML struct {
}

// TextParseModeHTMLTypeID is TL type id of TextParseModeHTML.
const TextParseModeHTMLTypeID = 0x62f4c5f3

// construct implements constructor of TextParseModeClass.
func (t TextParseModeHTML) construct() TextParseModeClass { return &t }

// Ensuring interfaces in compile-time for TextParseModeHTML.
var (
	_ bin.Encoder     = &TextParseModeHTML{}
	_ bin.Decoder     = &TextParseModeHTML{}
	_ bin.BareEncoder = &TextParseModeHTML{}
	_ bin.BareDecoder = &TextParseModeHTML{}

	_ TextParseModeClass = &TextParseModeHTML{}
)

func (t *TextParseModeHTML) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TextParseModeHTML) String() string {
	if t == nil {
		return "TextParseModeHTML(nil)"
	}
	type Alias TextParseModeHTML
	return fmt.Sprintf("TextParseModeHTML%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TextParseModeHTML) TypeID() uint32 {
	return TextParseModeHTMLTypeID
}

// TypeName returns name of type in TL schema.
func (*TextParseModeHTML) TypeName() string {
	return "textParseModeHTML"
}

// TypeInfo returns info about TL type.
func (t *TextParseModeHTML) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "textParseModeHTML",
		ID:   TextParseModeHTMLTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TextParseModeHTML) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textParseModeHTML#62f4c5f3 as nil")
	}
	b.PutID(TextParseModeHTMLTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TextParseModeHTML) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textParseModeHTML#62f4c5f3 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextParseModeHTML) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textParseModeHTML#62f4c5f3 to nil")
	}
	if err := b.ConsumeID(TextParseModeHTMLTypeID); err != nil {
		return fmt.Errorf("unable to decode textParseModeHTML#62f4c5f3: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TextParseModeHTML) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textParseModeHTML#62f4c5f3 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TextParseModeHTML) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode textParseModeHTML#62f4c5f3 as nil")
	}
	b.ObjStart()
	b.PutID("textParseModeHTML")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TextParseModeHTML) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode textParseModeHTML#62f4c5f3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("textParseModeHTML"); err != nil {
				return fmt.Errorf("unable to decode textParseModeHTML#62f4c5f3: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// TextParseModeClassName is schema name of TextParseModeClass.
const TextParseModeClassName = "TextParseMode"

// TextParseModeClass represents TextParseMode generic type.
//
// Example:
//  g, err := tdapi.DecodeTextParseMode(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.TextParseModeMarkdown: // textParseModeMarkdown#157648bf
//  case *tdapi.TextParseModeHTML: // textParseModeHTML#62f4c5f3
//  default: panic(v)
//  }
type TextParseModeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() TextParseModeClass

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

// DecodeTextParseMode implements binary de-serialization for TextParseModeClass.
func DecodeTextParseMode(buf *bin.Buffer) (TextParseModeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case TextParseModeMarkdownTypeID:
		// Decoding textParseModeMarkdown#157648bf.
		v := TextParseModeMarkdown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TextParseModeClass: %w", err)
		}
		return &v, nil
	case TextParseModeHTMLTypeID:
		// Decoding textParseModeHTML#62f4c5f3.
		v := TextParseModeHTML{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TextParseModeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode TextParseModeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONTextParseMode implements binary de-serialization for TextParseModeClass.
func DecodeTDLibJSONTextParseMode(buf tdjson.Decoder) (TextParseModeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "textParseModeMarkdown":
		// Decoding textParseModeMarkdown#157648bf.
		v := TextParseModeMarkdown{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TextParseModeClass: %w", err)
		}
		return &v, nil
	case "textParseModeHTML":
		// Decoding textParseModeHTML#62f4c5f3.
		v := TextParseModeHTML{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TextParseModeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode TextParseModeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// TextParseMode boxes the TextParseModeClass providing a helper.
type TextParseModeBox struct {
	TextParseMode TextParseModeClass
}

// Decode implements bin.Decoder for TextParseModeBox.
func (b *TextParseModeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode TextParseModeBox to nil")
	}
	v, err := DecodeTextParseMode(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TextParseMode = v
	return nil
}

// Encode implements bin.Encode for TextParseModeBox.
func (b *TextParseModeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.TextParseMode == nil {
		return fmt.Errorf("unable to encode TextParseModeClass as nil")
	}
	return b.TextParseMode.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for TextParseModeBox.
func (b *TextParseModeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode TextParseModeBox to nil")
	}
	v, err := DecodeTDLibJSONTextParseMode(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TextParseMode = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for TextParseModeBox.
func (b *TextParseModeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.TextParseMode == nil {
		return fmt.Errorf("unable to encode TextParseModeClass as nil")
	}
	return b.TextParseMode.EncodeTDLibJSON(buf)
}
