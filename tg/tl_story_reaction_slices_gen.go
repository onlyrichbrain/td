//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// StoryReactionClassArray is adapter for slice of StoryReactionClass.
type StoryReactionClassArray []StoryReactionClass

// Sort sorts slice of StoryReactionClass.
func (s StoryReactionClassArray) Sort(less func(a, b StoryReactionClass) bool) StoryReactionClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StoryReactionClass.
func (s StoryReactionClassArray) SortStable(less func(a, b StoryReactionClass) bool) StoryReactionClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StoryReactionClass.
func (s StoryReactionClassArray) Retain(keep func(x StoryReactionClass) bool) StoryReactionClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s StoryReactionClassArray) First() (v StoryReactionClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StoryReactionClassArray) Last() (v StoryReactionClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StoryReactionClassArray) PopFirst() (v StoryReactionClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StoryReactionClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StoryReactionClassArray) Pop() (v StoryReactionClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsStoryReaction returns copy with only StoryReaction constructors.
func (s StoryReactionClassArray) AsStoryReaction() (to StoryReactionArray) {
	for _, elem := range s {
		value, ok := elem.(*StoryReaction)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsStoryReactionPublicForward returns copy with only StoryReactionPublicForward constructors.
func (s StoryReactionClassArray) AsStoryReactionPublicForward() (to StoryReactionPublicForwardArray) {
	for _, elem := range s {
		value, ok := elem.(*StoryReactionPublicForward)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsStoryReactionPublicRepost returns copy with only StoryReactionPublicRepost constructors.
func (s StoryReactionClassArray) AsStoryReactionPublicRepost() (to StoryReactionPublicRepostArray) {
	for _, elem := range s {
		value, ok := elem.(*StoryReactionPublicRepost)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// StoryReactionArray is adapter for slice of StoryReaction.
type StoryReactionArray []StoryReaction

// Sort sorts slice of StoryReaction.
func (s StoryReactionArray) Sort(less func(a, b StoryReaction) bool) StoryReactionArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StoryReaction.
func (s StoryReactionArray) SortStable(less func(a, b StoryReaction) bool) StoryReactionArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StoryReaction.
func (s StoryReactionArray) Retain(keep func(x StoryReaction) bool) StoryReactionArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s StoryReactionArray) First() (v StoryReaction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StoryReactionArray) Last() (v StoryReaction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StoryReactionArray) PopFirst() (v StoryReaction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StoryReaction
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StoryReactionArray) Pop() (v StoryReaction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of StoryReaction by Date.
func (s StoryReactionArray) SortByDate() StoryReactionArray {
	return s.Sort(func(a, b StoryReaction) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of StoryReaction by Date.
func (s StoryReactionArray) SortStableByDate() StoryReactionArray {
	return s.SortStable(func(a, b StoryReaction) bool {
		return a.GetDate() < b.GetDate()
	})
}

// StoryReactionPublicForwardArray is adapter for slice of StoryReactionPublicForward.
type StoryReactionPublicForwardArray []StoryReactionPublicForward

// Sort sorts slice of StoryReactionPublicForward.
func (s StoryReactionPublicForwardArray) Sort(less func(a, b StoryReactionPublicForward) bool) StoryReactionPublicForwardArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StoryReactionPublicForward.
func (s StoryReactionPublicForwardArray) SortStable(less func(a, b StoryReactionPublicForward) bool) StoryReactionPublicForwardArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StoryReactionPublicForward.
func (s StoryReactionPublicForwardArray) Retain(keep func(x StoryReactionPublicForward) bool) StoryReactionPublicForwardArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s StoryReactionPublicForwardArray) First() (v StoryReactionPublicForward, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StoryReactionPublicForwardArray) Last() (v StoryReactionPublicForward, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StoryReactionPublicForwardArray) PopFirst() (v StoryReactionPublicForward, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StoryReactionPublicForward
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StoryReactionPublicForwardArray) Pop() (v StoryReactionPublicForward, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// StoryReactionPublicRepostArray is adapter for slice of StoryReactionPublicRepost.
type StoryReactionPublicRepostArray []StoryReactionPublicRepost

// Sort sorts slice of StoryReactionPublicRepost.
func (s StoryReactionPublicRepostArray) Sort(less func(a, b StoryReactionPublicRepost) bool) StoryReactionPublicRepostArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StoryReactionPublicRepost.
func (s StoryReactionPublicRepostArray) SortStable(less func(a, b StoryReactionPublicRepost) bool) StoryReactionPublicRepostArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StoryReactionPublicRepost.
func (s StoryReactionPublicRepostArray) Retain(keep func(x StoryReactionPublicRepost) bool) StoryReactionPublicRepostArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s StoryReactionPublicRepostArray) First() (v StoryReactionPublicRepost, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StoryReactionPublicRepostArray) Last() (v StoryReactionPublicRepost, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StoryReactionPublicRepostArray) PopFirst() (v StoryReactionPublicRepost, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StoryReactionPublicRepost
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StoryReactionPublicRepostArray) Pop() (v StoryReactionPublicRepost, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
