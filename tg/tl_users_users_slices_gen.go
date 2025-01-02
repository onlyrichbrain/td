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

// UsersUsersClassArray is adapter for slice of UsersUsersClass.
type UsersUsersClassArray []UsersUsersClass

// Sort sorts slice of UsersUsersClass.
func (s UsersUsersClassArray) Sort(less func(a, b UsersUsersClass) bool) UsersUsersClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UsersUsersClass.
func (s UsersUsersClassArray) SortStable(less func(a, b UsersUsersClass) bool) UsersUsersClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UsersUsersClass.
func (s UsersUsersClassArray) Retain(keep func(x UsersUsersClass) bool) UsersUsersClassArray {
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
func (s UsersUsersClassArray) First() (v UsersUsersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UsersUsersClassArray) Last() (v UsersUsersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UsersUsersClassArray) PopFirst() (v UsersUsersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UsersUsersClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UsersUsersClassArray) Pop() (v UsersUsersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsUsersUsers returns copy with only UsersUsers constructors.
func (s UsersUsersClassArray) AsUsersUsers() (to UsersUsersArray) {
	for _, elem := range s {
		value, ok := elem.(*UsersUsers)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUsersUsersSlice returns copy with only UsersUsersSlice constructors.
func (s UsersUsersClassArray) AsUsersUsersSlice() (to UsersUsersSliceArray) {
	for _, elem := range s {
		value, ok := elem.(*UsersUsersSlice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// UsersUsersArray is adapter for slice of UsersUsers.
type UsersUsersArray []UsersUsers

// Sort sorts slice of UsersUsers.
func (s UsersUsersArray) Sort(less func(a, b UsersUsers) bool) UsersUsersArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UsersUsers.
func (s UsersUsersArray) SortStable(less func(a, b UsersUsers) bool) UsersUsersArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UsersUsers.
func (s UsersUsersArray) Retain(keep func(x UsersUsers) bool) UsersUsersArray {
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
func (s UsersUsersArray) First() (v UsersUsers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UsersUsersArray) Last() (v UsersUsers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UsersUsersArray) PopFirst() (v UsersUsers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UsersUsers
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UsersUsersArray) Pop() (v UsersUsers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UsersUsersSliceArray is adapter for slice of UsersUsersSlice.
type UsersUsersSliceArray []UsersUsersSlice

// Sort sorts slice of UsersUsersSlice.
func (s UsersUsersSliceArray) Sort(less func(a, b UsersUsersSlice) bool) UsersUsersSliceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UsersUsersSlice.
func (s UsersUsersSliceArray) SortStable(less func(a, b UsersUsersSlice) bool) UsersUsersSliceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UsersUsersSlice.
func (s UsersUsersSliceArray) Retain(keep func(x UsersUsersSlice) bool) UsersUsersSliceArray {
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
func (s UsersUsersSliceArray) First() (v UsersUsersSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UsersUsersSliceArray) Last() (v UsersUsersSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UsersUsersSliceArray) PopFirst() (v UsersUsersSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UsersUsersSlice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UsersUsersSliceArray) Pop() (v UsersUsersSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
