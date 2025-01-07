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

// AccountUpdateUsernameRequest represents TL type `account.updateUsername#3e0bdd7c`.
// Changes username for the current user.
//
// See https://core.telegram.org/method/account.updateUsername for reference.
type AccountUpdateUsernameRequest struct {
	// username or empty string if username is to be removedAccepted characters: a-z
	// (case-insensitive), 0-9 and underscores.Length: 5-32 characters.
	Username string
}

// AccountUpdateUsernameRequestTypeID is TL type id of AccountUpdateUsernameRequest.
const AccountUpdateUsernameRequestTypeID = 0x3e0bdd7c

// Ensuring interfaces in compile-time for AccountUpdateUsernameRequest.
var (
	_ bin.Encoder     = &AccountUpdateUsernameRequest{}
	_ bin.Decoder     = &AccountUpdateUsernameRequest{}
	_ bin.BareEncoder = &AccountUpdateUsernameRequest{}
	_ bin.BareDecoder = &AccountUpdateUsernameRequest{}
)

func (u *AccountUpdateUsernameRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateUsernameRequest) String() string {
	if u == nil {
		return "AccountUpdateUsernameRequest(nil)"
	}
	type Alias AccountUpdateUsernameRequest
	return fmt.Sprintf("AccountUpdateUsernameRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateUsernameRequest from given interface.
func (u *AccountUpdateUsernameRequest) FillFrom(from interface {
	GetUsername() (value string)
}) {
	u.Username = from.GetUsername()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateUsernameRequest) TypeID() uint32 {
	return AccountUpdateUsernameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateUsernameRequest) TypeName() string {
	return "account.updateUsername"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateUsernameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateUsername",
		ID:   AccountUpdateUsernameRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Username",
			SchemaName: "username",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUpdateUsernameRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateUsername#3e0bdd7c as nil")
	}
	b.PutID(AccountUpdateUsernameRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateUsernameRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateUsername#3e0bdd7c as nil")
	}
	b.PutString(u.Username)
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateUsernameRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateUsername#3e0bdd7c to nil")
	}
	if err := b.ConsumeID(AccountUpdateUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateUsername#3e0bdd7c: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateUsernameRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateUsername#3e0bdd7c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateUsername#3e0bdd7c: field username: %w", err)
		}
		u.Username = value
	}
	return nil
}

// GetUsername returns value of Username field.
func (u *AccountUpdateUsernameRequest) GetUsername() (value string) {
	if u == nil {
		return
	}
	return u.Username
}

// AccountUpdateUsername invokes method account.updateUsername#3e0bdd7c returning error if any.
// Changes username for the current user.
//
// Possible errors:
//
//	400 USERNAME_INVALID: The provided username is not valid.
//	400 USERNAME_NOT_MODIFIED: The username was not modified.
//	400 USERNAME_OCCUPIED: The provided username is already occupied.
//	400 USERNAME_PURCHASE_AVAILABLE: The specified username can be purchased on https://fragment.com.
//
// See https://core.telegram.org/method/account.updateUsername for reference.
func (c *Client) AccountUpdateUsername(ctx context.Context, username string) (UserClass, error) {
	var result UserBox

	request := &AccountUpdateUsernameRequest{
		Username: username,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.User, nil
}
