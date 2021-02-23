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

// AccountGetMultiWallPapersRequest represents TL type `account.getMultiWallPapers#65ad71dc`.
// Get info about multiple wallpapers
//
// See https://core.telegram.org/method/account.getMultiWallPapers for reference.
type AccountGetMultiWallPapersRequest struct {
	// Wallpapers to fetch info about
	Wallpapers []InputWallPaperClass `schemaname:"wallpapers"`
}

// AccountGetMultiWallPapersRequestTypeID is TL type id of AccountGetMultiWallPapersRequest.
const AccountGetMultiWallPapersRequestTypeID = 0x65ad71dc

func (g *AccountGetMultiWallPapersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Wallpapers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetMultiWallPapersRequest) String() string {
	if g == nil {
		return "AccountGetMultiWallPapersRequest(nil)"
	}
	type Alias AccountGetMultiWallPapersRequest
	return fmt.Sprintf("AccountGetMultiWallPapersRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetMultiWallPapersRequest from given interface.
func (g *AccountGetMultiWallPapersRequest) FillFrom(from interface {
	GetWallpapers() (value []InputWallPaperClass)
}) {
	g.Wallpapers = from.GetWallpapers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *AccountGetMultiWallPapersRequest) TypeID() uint32 {
	return AccountGetMultiWallPapersRequestTypeID
}

// SchemaName returns MTProto type name.
func (g *AccountGetMultiWallPapersRequest) SchemaName() string {
	return "account.getMultiWallPapers"
}

// Encode implements bin.Encoder.
func (g *AccountGetMultiWallPapersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getMultiWallPapers#65ad71dc as nil")
	}
	b.PutID(AccountGetMultiWallPapersRequestTypeID)
	b.PutVectorHeader(len(g.Wallpapers))
	for idx, v := range g.Wallpapers {
		if v == nil {
			return fmt.Errorf("unable to encode account.getMultiWallPapers#65ad71dc: field wallpapers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.getMultiWallPapers#65ad71dc: field wallpapers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetWallpapers returns value of Wallpapers field.
func (g *AccountGetMultiWallPapersRequest) GetWallpapers() (value []InputWallPaperClass) {
	return g.Wallpapers
}

// MapWallpapers returns field Wallpapers wrapped in InputWallPaperClassSlice helper.
func (g *AccountGetMultiWallPapersRequest) MapWallpapers() (value InputWallPaperClassSlice) {
	return InputWallPaperClassSlice(g.Wallpapers)
}

// Decode implements bin.Decoder.
func (g *AccountGetMultiWallPapersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getMultiWallPapers#65ad71dc to nil")
	}
	if err := b.ConsumeID(AccountGetMultiWallPapersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getMultiWallPapers#65ad71dc: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.getMultiWallPapers#65ad71dc: field wallpapers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputWallPaper(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.getMultiWallPapers#65ad71dc: field wallpapers: %w", err)
			}
			g.Wallpapers = append(g.Wallpapers, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetMultiWallPapersRequest.
var (
	_ bin.Encoder = &AccountGetMultiWallPapersRequest{}
	_ bin.Decoder = &AccountGetMultiWallPapersRequest{}
)

// AccountGetMultiWallPapers invokes method account.getMultiWallPapers#65ad71dc returning error if any.
// Get info about multiple wallpapers
//
// See https://core.telegram.org/method/account.getMultiWallPapers for reference.
func (c *Client) AccountGetMultiWallPapers(ctx context.Context, wallpapers []InputWallPaperClass) ([]WallPaperClass, error) {
	var result WallPaperClassVector

	request := &AccountGetMultiWallPapersRequest{
		Wallpapers: wallpapers,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}
