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

// AccountSetReactionsNotifySettingsRequest represents TL type `account.setReactionsNotifySettings#316ce548`.
// Change the reaction notification settings »¹.
//
// Links:
//  1. https://core.telegram.org/api/reactions#notifications-about-reactions
//
// See https://core.telegram.org/method/account.setReactionsNotifySettings for reference.
type AccountSetReactionsNotifySettingsRequest struct {
	// New reaction notification settings.
	Settings ReactionsNotifySettings
}

// AccountSetReactionsNotifySettingsRequestTypeID is TL type id of AccountSetReactionsNotifySettingsRequest.
const AccountSetReactionsNotifySettingsRequestTypeID = 0x316ce548

// Ensuring interfaces in compile-time for AccountSetReactionsNotifySettingsRequest.
var (
	_ bin.Encoder     = &AccountSetReactionsNotifySettingsRequest{}
	_ bin.Decoder     = &AccountSetReactionsNotifySettingsRequest{}
	_ bin.BareEncoder = &AccountSetReactionsNotifySettingsRequest{}
	_ bin.BareDecoder = &AccountSetReactionsNotifySettingsRequest{}
)

func (s *AccountSetReactionsNotifySettingsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSetReactionsNotifySettingsRequest) String() string {
	if s == nil {
		return "AccountSetReactionsNotifySettingsRequest(nil)"
	}
	type Alias AccountSetReactionsNotifySettingsRequest
	return fmt.Sprintf("AccountSetReactionsNotifySettingsRequest%+v", Alias(*s))
}

// FillFrom fills AccountSetReactionsNotifySettingsRequest from given interface.
func (s *AccountSetReactionsNotifySettingsRequest) FillFrom(from interface {
	GetSettings() (value ReactionsNotifySettings)
}) {
	s.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSetReactionsNotifySettingsRequest) TypeID() uint32 {
	return AccountSetReactionsNotifySettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSetReactionsNotifySettingsRequest) TypeName() string {
	return "account.setReactionsNotifySettings"
}

// TypeInfo returns info about TL type.
func (s *AccountSetReactionsNotifySettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.setReactionsNotifySettings",
		ID:   AccountSetReactionsNotifySettingsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSetReactionsNotifySettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setReactionsNotifySettings#316ce548 as nil")
	}
	b.PutID(AccountSetReactionsNotifySettingsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSetReactionsNotifySettingsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setReactionsNotifySettings#316ce548 as nil")
	}
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.setReactionsNotifySettings#316ce548: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSetReactionsNotifySettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setReactionsNotifySettings#316ce548 to nil")
	}
	if err := b.ConsumeID(AccountSetReactionsNotifySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.setReactionsNotifySettings#316ce548: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSetReactionsNotifySettingsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setReactionsNotifySettings#316ce548 to nil")
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.setReactionsNotifySettings#316ce548: field settings: %w", err)
		}
	}
	return nil
}

// GetSettings returns value of Settings field.
func (s *AccountSetReactionsNotifySettingsRequest) GetSettings() (value ReactionsNotifySettings) {
	if s == nil {
		return
	}
	return s.Settings
}

// AccountSetReactionsNotifySettings invokes method account.setReactionsNotifySettings#316ce548 returning error if any.
// Change the reaction notification settings »¹.
//
// Links:
//  1. https://core.telegram.org/api/reactions#notifications-about-reactions
//
// See https://core.telegram.org/method/account.setReactionsNotifySettings for reference.
func (c *Client) AccountSetReactionsNotifySettings(ctx context.Context, settings ReactionsNotifySettings) (*ReactionsNotifySettings, error) {
	var result ReactionsNotifySettings

	request := &AccountSetReactionsNotifySettingsRequest{
		Settings: settings,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
