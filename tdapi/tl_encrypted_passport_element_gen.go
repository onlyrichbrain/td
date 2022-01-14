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

// EncryptedPassportElement represents TL type `encryptedPassportElement#262d248`.
type EncryptedPassportElement struct {
	// Type of Telegram Passport element
	Type PassportElementTypeClass
	// Encrypted JSON-encoded data about the user
	Data []byte
	// The front side of an identity document
	FrontSide DatedFile
	// The reverse side of an identity document; may be null
	ReverseSide DatedFile
	// Selfie with the document; may be null
	Selfie DatedFile
	// List of files containing a certified English translation of the document
	Translation []DatedFile
	// List of attached files
	Files []DatedFile
	// Unencrypted data, phone number or email address
	Value string
	// Hash of the entire element
	Hash string
}

// EncryptedPassportElementTypeID is TL type id of EncryptedPassportElement.
const EncryptedPassportElementTypeID = 0x262d248

// Ensuring interfaces in compile-time for EncryptedPassportElement.
var (
	_ bin.Encoder     = &EncryptedPassportElement{}
	_ bin.Decoder     = &EncryptedPassportElement{}
	_ bin.BareEncoder = &EncryptedPassportElement{}
	_ bin.BareDecoder = &EncryptedPassportElement{}
)

func (e *EncryptedPassportElement) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Type == nil) {
		return false
	}
	if !(e.Data == nil) {
		return false
	}
	if !(e.FrontSide.Zero()) {
		return false
	}
	if !(e.ReverseSide.Zero()) {
		return false
	}
	if !(e.Selfie.Zero()) {
		return false
	}
	if !(e.Translation == nil) {
		return false
	}
	if !(e.Files == nil) {
		return false
	}
	if !(e.Value == "") {
		return false
	}
	if !(e.Hash == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedPassportElement) String() string {
	if e == nil {
		return "EncryptedPassportElement(nil)"
	}
	type Alias EncryptedPassportElement
	return fmt.Sprintf("EncryptedPassportElement%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedPassportElement) TypeID() uint32 {
	return EncryptedPassportElementTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedPassportElement) TypeName() string {
	return "encryptedPassportElement"
}

// TypeInfo returns info about TL type.
func (e *EncryptedPassportElement) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedPassportElement",
		ID:   EncryptedPassportElementTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Data",
			SchemaName: "data",
		},
		{
			Name:       "FrontSide",
			SchemaName: "front_side",
		},
		{
			Name:       "ReverseSide",
			SchemaName: "reverse_side",
		},
		{
			Name:       "Selfie",
			SchemaName: "selfie",
		},
		{
			Name:       "Translation",
			SchemaName: "translation",
		},
		{
			Name:       "Files",
			SchemaName: "files",
		},
		{
			Name:       "Value",
			SchemaName: "value",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EncryptedPassportElement) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedPassportElement#262d248 as nil")
	}
	b.PutID(EncryptedPassportElementTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EncryptedPassportElement) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedPassportElement#262d248 as nil")
	}
	if e.Type == nil {
		return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field type is nil")
	}
	if err := e.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field type: %w", err)
	}
	b.PutBytes(e.Data)
	if err := e.FrontSide.Encode(b); err != nil {
		return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field front_side: %w", err)
	}
	if err := e.ReverseSide.Encode(b); err != nil {
		return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field reverse_side: %w", err)
	}
	if err := e.Selfie.Encode(b); err != nil {
		return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field selfie: %w", err)
	}
	b.PutInt(len(e.Translation))
	for idx, v := range e.Translation {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare encryptedPassportElement#262d248: field translation element with index %d: %w", idx, err)
		}
	}
	b.PutInt(len(e.Files))
	for idx, v := range e.Files {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare encryptedPassportElement#262d248: field files element with index %d: %w", idx, err)
		}
	}
	b.PutString(e.Value)
	b.PutString(e.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedPassportElement) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedPassportElement#262d248 to nil")
	}
	if err := b.ConsumeID(EncryptedPassportElementTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedPassportElement#262d248: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EncryptedPassportElement) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedPassportElement#262d248 to nil")
	}
	{
		value, err := DecodePassportElementType(b)
		if err != nil {
			return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field type: %w", err)
		}
		e.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field data: %w", err)
		}
		e.Data = value
	}
	{
		if err := e.FrontSide.Decode(b); err != nil {
			return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field front_side: %w", err)
		}
	}
	{
		if err := e.ReverseSide.Decode(b); err != nil {
			return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field reverse_side: %w", err)
		}
	}
	{
		if err := e.Selfie.Decode(b); err != nil {
			return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field selfie: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field translation: %w", err)
		}

		if headerLen > 0 {
			e.Translation = make([]DatedFile, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value DatedFile
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare encryptedPassportElement#262d248: field translation: %w", err)
			}
			e.Translation = append(e.Translation, value)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field files: %w", err)
		}

		if headerLen > 0 {
			e.Files = make([]DatedFile, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value DatedFile
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare encryptedPassportElement#262d248: field files: %w", err)
			}
			e.Files = append(e.Files, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field value: %w", err)
		}
		e.Value = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field hash: %w", err)
		}
		e.Hash = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EncryptedPassportElement) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedPassportElement#262d248 as nil")
	}
	b.ObjStart()
	b.PutID("encryptedPassportElement")
	b.Comma()
	b.FieldStart("type")
	if e.Type == nil {
		return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field type is nil")
	}
	if err := e.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("data")
	b.PutBytes(e.Data)
	b.Comma()
	b.FieldStart("front_side")
	if err := e.FrontSide.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field front_side: %w", err)
	}
	b.Comma()
	b.FieldStart("reverse_side")
	if err := e.ReverseSide.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field reverse_side: %w", err)
	}
	b.Comma()
	b.FieldStart("selfie")
	if err := e.Selfie.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field selfie: %w", err)
	}
	b.Comma()
	b.FieldStart("translation")
	b.ArrStart()
	for idx, v := range e.Translation {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field translation element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("files")
	b.ArrStart()
	for idx, v := range e.Files {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode encryptedPassportElement#262d248: field files element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("value")
	b.PutString(e.Value)
	b.Comma()
	b.FieldStart("hash")
	b.PutString(e.Hash)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EncryptedPassportElement) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedPassportElement#262d248 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("encryptedPassportElement"); err != nil {
				return fmt.Errorf("unable to decode encryptedPassportElement#262d248: %w", err)
			}
		case "type":
			value, err := DecodeTDLibJSONPassportElementType(b)
			if err != nil {
				return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field type: %w", err)
			}
			e.Type = value
		case "data":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field data: %w", err)
			}
			e.Data = value
		case "front_side":
			if err := e.FrontSide.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field front_side: %w", err)
			}
		case "reverse_side":
			if err := e.ReverseSide.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field reverse_side: %w", err)
			}
		case "selfie":
			if err := e.Selfie.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field selfie: %w", err)
			}
		case "translation":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value DatedFile
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field translation: %w", err)
				}
				e.Translation = append(e.Translation, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field translation: %w", err)
			}
		case "files":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value DatedFile
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field files: %w", err)
				}
				e.Files = append(e.Files, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field files: %w", err)
			}
		case "value":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field value: %w", err)
			}
			e.Value = value
		case "hash":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode encryptedPassportElement#262d248: field hash: %w", err)
			}
			e.Hash = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetType returns value of Type field.
func (e *EncryptedPassportElement) GetType() (value PassportElementTypeClass) {
	if e == nil {
		return
	}
	return e.Type
}

// GetData returns value of Data field.
func (e *EncryptedPassportElement) GetData() (value []byte) {
	if e == nil {
		return
	}
	return e.Data
}

// GetFrontSide returns value of FrontSide field.
func (e *EncryptedPassportElement) GetFrontSide() (value DatedFile) {
	if e == nil {
		return
	}
	return e.FrontSide
}

// GetReverseSide returns value of ReverseSide field.
func (e *EncryptedPassportElement) GetReverseSide() (value DatedFile) {
	if e == nil {
		return
	}
	return e.ReverseSide
}

// GetSelfie returns value of Selfie field.
func (e *EncryptedPassportElement) GetSelfie() (value DatedFile) {
	if e == nil {
		return
	}
	return e.Selfie
}

// GetTranslation returns value of Translation field.
func (e *EncryptedPassportElement) GetTranslation() (value []DatedFile) {
	if e == nil {
		return
	}
	return e.Translation
}

// GetFiles returns value of Files field.
func (e *EncryptedPassportElement) GetFiles() (value []DatedFile) {
	if e == nil {
		return
	}
	return e.Files
}

// GetValue returns value of Value field.
func (e *EncryptedPassportElement) GetValue() (value string) {
	if e == nil {
		return
	}
	return e.Value
}

// GetHash returns value of Hash field.
func (e *EncryptedPassportElement) GetHash() (value string) {
	if e == nil {
		return
	}
	return e.Hash
}
