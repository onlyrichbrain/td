// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// ExportedMessageLink represents TL type `exportedMessageLink#5dab1af4`.
// Link to a message in a supergroup/channel
//
// See https://core.telegram.org/constructor/exportedMessageLink for reference.
type ExportedMessageLink struct {
	// URL
	Link string `schemaname:"link"`
	// Embed code
	HTML string `schemaname:"html"`
}

// ExportedMessageLinkTypeID is TL type id of ExportedMessageLink.
const ExportedMessageLinkTypeID = 0x5dab1af4

func (e *ExportedMessageLink) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Link == "") {
		return false
	}
	if !(e.HTML == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ExportedMessageLink) String() string {
	if e == nil {
		return "ExportedMessageLink(nil)"
	}
	type Alias ExportedMessageLink
	return fmt.Sprintf("ExportedMessageLink%+v", Alias(*e))
}

// FillFrom fills ExportedMessageLink from given interface.
func (e *ExportedMessageLink) FillFrom(from interface {
	GetLink() (value string)
	GetHTML() (value string)
}) {
	e.Link = from.GetLink()
	e.HTML = from.GetHTML()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *ExportedMessageLink) TypeID() uint32 {
	return ExportedMessageLinkTypeID
}

// SchemaName returns MTProto type name.
func (e *ExportedMessageLink) SchemaName() string {
	return "exportedMessageLink"
}

// Encode implements bin.Encoder.
func (e *ExportedMessageLink) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode exportedMessageLink#5dab1af4 as nil")
	}
	b.PutID(ExportedMessageLinkTypeID)
	b.PutString(e.Link)
	b.PutString(e.HTML)
	return nil
}

// GetLink returns value of Link field.
func (e *ExportedMessageLink) GetLink() (value string) {
	return e.Link
}

// GetHTML returns value of HTML field.
func (e *ExportedMessageLink) GetHTML() (value string) {
	return e.HTML
}

// Decode implements bin.Decoder.
func (e *ExportedMessageLink) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode exportedMessageLink#5dab1af4 to nil")
	}
	if err := b.ConsumeID(ExportedMessageLinkTypeID); err != nil {
		return fmt.Errorf("unable to decode exportedMessageLink#5dab1af4: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode exportedMessageLink#5dab1af4: field link: %w", err)
		}
		e.Link = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode exportedMessageLink#5dab1af4: field html: %w", err)
		}
		e.HTML = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ExportedMessageLink.
var (
	_ bin.Encoder = &ExportedMessageLink{}
	_ bin.Decoder = &ExportedMessageLink{}
)
