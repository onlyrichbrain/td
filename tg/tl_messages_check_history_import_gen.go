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

// MessagesCheckHistoryImportRequest represents TL type `messages.checkHistoryImport#43fe19f3`.
//
// See https://core.telegram.org/method/messages.checkHistoryImport for reference.
type MessagesCheckHistoryImportRequest struct {
	// ImportHead field of MessagesCheckHistoryImportRequest.
	ImportHead string `schemaname:"import_head"`
}

// MessagesCheckHistoryImportRequestTypeID is TL type id of MessagesCheckHistoryImportRequest.
const MessagesCheckHistoryImportRequestTypeID = 0x43fe19f3

func (c *MessagesCheckHistoryImportRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ImportHead == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesCheckHistoryImportRequest) String() string {
	if c == nil {
		return "MessagesCheckHistoryImportRequest(nil)"
	}
	type Alias MessagesCheckHistoryImportRequest
	return fmt.Sprintf("MessagesCheckHistoryImportRequest%+v", Alias(*c))
}

// FillFrom fills MessagesCheckHistoryImportRequest from given interface.
func (c *MessagesCheckHistoryImportRequest) FillFrom(from interface {
	GetImportHead() (value string)
}) {
	c.ImportHead = from.GetImportHead()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *MessagesCheckHistoryImportRequest) TypeID() uint32 {
	return MessagesCheckHistoryImportRequestTypeID
}

// SchemaName returns MTProto type name.
func (c *MessagesCheckHistoryImportRequest) SchemaName() string {
	return "messages.checkHistoryImport"
}

// Encode implements bin.Encoder.
func (c *MessagesCheckHistoryImportRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.checkHistoryImport#43fe19f3 as nil")
	}
	b.PutID(MessagesCheckHistoryImportRequestTypeID)
	b.PutString(c.ImportHead)
	return nil
}

// GetImportHead returns value of ImportHead field.
func (c *MessagesCheckHistoryImportRequest) GetImportHead() (value string) {
	return c.ImportHead
}

// Decode implements bin.Decoder.
func (c *MessagesCheckHistoryImportRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.checkHistoryImport#43fe19f3 to nil")
	}
	if err := b.ConsumeID(MessagesCheckHistoryImportRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.checkHistoryImport#43fe19f3: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.checkHistoryImport#43fe19f3: field import_head: %w", err)
		}
		c.ImportHead = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesCheckHistoryImportRequest.
var (
	_ bin.Encoder = &MessagesCheckHistoryImportRequest{}
	_ bin.Decoder = &MessagesCheckHistoryImportRequest{}
)

// MessagesCheckHistoryImport invokes method messages.checkHistoryImport#43fe19f3 returning error if any.
//
// See https://core.telegram.org/method/messages.checkHistoryImport for reference.
func (c *Client) MessagesCheckHistoryImport(ctx context.Context, importhead string) (*MessagesHistoryImportParsed, error) {
	var result MessagesHistoryImportParsed

	request := &MessagesCheckHistoryImportRequest{
		ImportHead: importhead,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
