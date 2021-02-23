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

// AccountSetPrivacyRequest represents TL type `account.setPrivacy#c9f81ce8`.
// Change privacy settings of current account
//
// See https://core.telegram.org/method/account.setPrivacy for reference.
type AccountSetPrivacyRequest struct {
	// Peers to which the privacy rules apply
	Key InputPrivacyKeyClass `schemaname:"key"`
	// New privacy rules
	Rules []InputPrivacyRuleClass `schemaname:"rules"`
}

// AccountSetPrivacyRequestTypeID is TL type id of AccountSetPrivacyRequest.
const AccountSetPrivacyRequestTypeID = 0xc9f81ce8

func (s *AccountSetPrivacyRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Key == nil) {
		return false
	}
	if !(s.Rules == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSetPrivacyRequest) String() string {
	if s == nil {
		return "AccountSetPrivacyRequest(nil)"
	}
	type Alias AccountSetPrivacyRequest
	return fmt.Sprintf("AccountSetPrivacyRequest%+v", Alias(*s))
}

// FillFrom fills AccountSetPrivacyRequest from given interface.
func (s *AccountSetPrivacyRequest) FillFrom(from interface {
	GetKey() (value InputPrivacyKeyClass)
	GetRules() (value []InputPrivacyRuleClass)
}) {
	s.Key = from.GetKey()
	s.Rules = from.GetRules()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AccountSetPrivacyRequest) TypeID() uint32 {
	return AccountSetPrivacyRequestTypeID
}

// SchemaName returns MTProto type name.
func (s *AccountSetPrivacyRequest) SchemaName() string {
	return "account.setPrivacy"
}

// Encode implements bin.Encoder.
func (s *AccountSetPrivacyRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setPrivacy#c9f81ce8 as nil")
	}
	b.PutID(AccountSetPrivacyRequestTypeID)
	if s.Key == nil {
		return fmt.Errorf("unable to encode account.setPrivacy#c9f81ce8: field key is nil")
	}
	if err := s.Key.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.setPrivacy#c9f81ce8: field key: %w", err)
	}
	b.PutVectorHeader(len(s.Rules))
	for idx, v := range s.Rules {
		if v == nil {
			return fmt.Errorf("unable to encode account.setPrivacy#c9f81ce8: field rules element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.setPrivacy#c9f81ce8: field rules element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetKey returns value of Key field.
func (s *AccountSetPrivacyRequest) GetKey() (value InputPrivacyKeyClass) {
	return s.Key
}

// GetRules returns value of Rules field.
func (s *AccountSetPrivacyRequest) GetRules() (value []InputPrivacyRuleClass) {
	return s.Rules
}

// MapRules returns field Rules wrapped in InputPrivacyRuleClassSlice helper.
func (s *AccountSetPrivacyRequest) MapRules() (value InputPrivacyRuleClassSlice) {
	return InputPrivacyRuleClassSlice(s.Rules)
}

// Decode implements bin.Decoder.
func (s *AccountSetPrivacyRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setPrivacy#c9f81ce8 to nil")
	}
	if err := b.ConsumeID(AccountSetPrivacyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.setPrivacy#c9f81ce8: %w", err)
	}
	{
		value, err := DecodeInputPrivacyKey(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.setPrivacy#c9f81ce8: field key: %w", err)
		}
		s.Key = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.setPrivacy#c9f81ce8: field rules: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPrivacyRule(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.setPrivacy#c9f81ce8: field rules: %w", err)
			}
			s.Rules = append(s.Rules, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSetPrivacyRequest.
var (
	_ bin.Encoder = &AccountSetPrivacyRequest{}
	_ bin.Decoder = &AccountSetPrivacyRequest{}
)

// AccountSetPrivacy invokes method account.setPrivacy#c9f81ce8 returning error if any.
// Change privacy settings of current account
//
// Possible errors:
//  400 PRIVACY_KEY_INVALID: The privacy key is invalid
//  400 PRIVACY_VALUE_INVALID: The specified privacy rule combination is invalid
//
// See https://core.telegram.org/method/account.setPrivacy for reference.
func (c *Client) AccountSetPrivacy(ctx context.Context, request *AccountSetPrivacyRequest) (*AccountPrivacyRules, error) {
	var result AccountPrivacyRules

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
