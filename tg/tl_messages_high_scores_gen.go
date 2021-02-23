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

// MessagesHighScores represents TL type `messages.highScores#9a3bfd99`.
// Highscores in a game
//
// See https://core.telegram.org/constructor/messages.highScores for reference.
type MessagesHighScores struct {
	// Highscores
	Scores []HighScore `schemaname:"scores"`
	// Users, associated to the highscores
	Users []UserClass `schemaname:"users"`
}

// MessagesHighScoresTypeID is TL type id of MessagesHighScores.
const MessagesHighScoresTypeID = 0x9a3bfd99

func (h *MessagesHighScores) Zero() bool {
	if h == nil {
		return true
	}
	if !(h.Scores == nil) {
		return false
	}
	if !(h.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (h *MessagesHighScores) String() string {
	if h == nil {
		return "MessagesHighScores(nil)"
	}
	type Alias MessagesHighScores
	return fmt.Sprintf("MessagesHighScores%+v", Alias(*h))
}

// FillFrom fills MessagesHighScores from given interface.
func (h *MessagesHighScores) FillFrom(from interface {
	GetScores() (value []HighScore)
	GetUsers() (value []UserClass)
}) {
	h.Scores = from.GetScores()
	h.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (h *MessagesHighScores) TypeID() uint32 {
	return MessagesHighScoresTypeID
}

// SchemaName returns MTProto type name.
func (h *MessagesHighScores) SchemaName() string {
	return "messages.highScores"
}

// Encode implements bin.Encoder.
func (h *MessagesHighScores) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.highScores#9a3bfd99 as nil")
	}
	b.PutID(MessagesHighScoresTypeID)
	b.PutVectorHeader(len(h.Scores))
	for idx, v := range h.Scores {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.highScores#9a3bfd99: field scores element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(h.Users))
	for idx, v := range h.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.highScores#9a3bfd99: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.highScores#9a3bfd99: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetScores returns value of Scores field.
func (h *MessagesHighScores) GetScores() (value []HighScore) {
	return h.Scores
}

// GetUsers returns value of Users field.
func (h *MessagesHighScores) GetUsers() (value []UserClass) {
	return h.Users
}

// MapUsers returns field Users wrapped in UserClassSlice helper.
func (h *MessagesHighScores) MapUsers() (value UserClassSlice) {
	return UserClassSlice(h.Users)
}

// Decode implements bin.Decoder.
func (h *MessagesHighScores) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.highScores#9a3bfd99 to nil")
	}
	if err := b.ConsumeID(MessagesHighScoresTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field scores: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value HighScore
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field scores: %w", err)
			}
			h.Scores = append(h.Scores, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.highScores#9a3bfd99: field users: %w", err)
			}
			h.Users = append(h.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesHighScores.
var (
	_ bin.Encoder = &MessagesHighScores{}
	_ bin.Decoder = &MessagesHighScores{}
)
