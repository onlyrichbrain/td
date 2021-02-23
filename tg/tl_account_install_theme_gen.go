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

// AccountInstallThemeRequest represents TL type `account.installTheme#7ae43737`.
// Install a theme
//
// See https://core.telegram.org/method/account.installTheme for reference.
type AccountInstallThemeRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether to install the dark version
	Dark bool `schemaname:"dark"`
	// Theme format, a string that identifies the theming engines supported by the client
	//
	// Use SetFormat and GetFormat helpers.
	Format string `schemaname:"format"`
	// Theme to install
	//
	// Use SetTheme and GetTheme helpers.
	Theme InputThemeClass `schemaname:"theme"`
}

// AccountInstallThemeRequestTypeID is TL type id of AccountInstallThemeRequest.
const AccountInstallThemeRequestTypeID = 0x7ae43737

func (i *AccountInstallThemeRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Dark == false) {
		return false
	}
	if !(i.Format == "") {
		return false
	}
	if !(i.Theme == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *AccountInstallThemeRequest) String() string {
	if i == nil {
		return "AccountInstallThemeRequest(nil)"
	}
	type Alias AccountInstallThemeRequest
	return fmt.Sprintf("AccountInstallThemeRequest%+v", Alias(*i))
}

// FillFrom fills AccountInstallThemeRequest from given interface.
func (i *AccountInstallThemeRequest) FillFrom(from interface {
	GetDark() (value bool)
	GetFormat() (value string, ok bool)
	GetTheme() (value InputThemeClass, ok bool)
}) {
	i.Dark = from.GetDark()
	if val, ok := from.GetFormat(); ok {
		i.Format = val
	}

	if val, ok := from.GetTheme(); ok {
		i.Theme = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *AccountInstallThemeRequest) TypeID() uint32 {
	return AccountInstallThemeRequestTypeID
}

// SchemaName returns MTProto type name.
func (i *AccountInstallThemeRequest) SchemaName() string {
	return "account.installTheme"
}

// Encode implements bin.Encoder.
func (i *AccountInstallThemeRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.installTheme#7ae43737 as nil")
	}
	b.PutID(AccountInstallThemeRequestTypeID)
	if !(i.Dark == false) {
		i.Flags.Set(0)
	}
	if !(i.Format == "") {
		i.Flags.Set(1)
	}
	if !(i.Theme == nil) {
		i.Flags.Set(1)
	}
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.installTheme#7ae43737: field flags: %w", err)
	}
	if i.Flags.Has(1) {
		b.PutString(i.Format)
	}
	if i.Flags.Has(1) {
		if i.Theme == nil {
			return fmt.Errorf("unable to encode account.installTheme#7ae43737: field theme is nil")
		}
		if err := i.Theme.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.installTheme#7ae43737: field theme: %w", err)
		}
	}
	return nil
}

// SetDark sets value of Dark conditional field.
func (i *AccountInstallThemeRequest) SetDark(value bool) {
	if value {
		i.Flags.Set(0)
		i.Dark = true
	} else {
		i.Flags.Unset(0)
		i.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (i *AccountInstallThemeRequest) GetDark() (value bool) {
	return i.Flags.Has(0)
}

// SetFormat sets value of Format conditional field.
func (i *AccountInstallThemeRequest) SetFormat(value string) {
	i.Flags.Set(1)
	i.Format = value
}

// GetFormat returns value of Format conditional field and
// boolean which is true if field was set.
func (i *AccountInstallThemeRequest) GetFormat() (value string, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Format, true
}

// SetTheme sets value of Theme conditional field.
func (i *AccountInstallThemeRequest) SetTheme(value InputThemeClass) {
	i.Flags.Set(1)
	i.Theme = value
}

// GetTheme returns value of Theme conditional field and
// boolean which is true if field was set.
func (i *AccountInstallThemeRequest) GetTheme() (value InputThemeClass, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Theme, true
}

// Decode implements bin.Decoder.
func (i *AccountInstallThemeRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.installTheme#7ae43737 to nil")
	}
	if err := b.ConsumeID(AccountInstallThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.installTheme#7ae43737: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.installTheme#7ae43737: field flags: %w", err)
		}
	}
	i.Dark = i.Flags.Has(0)
	if i.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.installTheme#7ae43737: field format: %w", err)
		}
		i.Format = value
	}
	if i.Flags.Has(1) {
		value, err := DecodeInputTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.installTheme#7ae43737: field theme: %w", err)
		}
		i.Theme = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountInstallThemeRequest.
var (
	_ bin.Encoder = &AccountInstallThemeRequest{}
	_ bin.Decoder = &AccountInstallThemeRequest{}
)

// AccountInstallTheme invokes method account.installTheme#7ae43737 returning error if any.
// Install a theme
//
// See https://core.telegram.org/method/account.installTheme for reference.
func (c *Client) AccountInstallTheme(ctx context.Context, request *AccountInstallThemeRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
