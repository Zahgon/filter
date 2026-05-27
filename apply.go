// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package filter contains utility functions for filtering slices through the
// distributed application of a filter function.
//
// The package is an experiment to see how easy it is to write such things
// in Go. It is easy, but for loops are just as easy and more efficient.
//
// You should not use this package.
package filter // import "robpike.io/filter"

import (
	"reflect"
)

// Apply takes a slice of type []T and a function of type func(T) T. (If the
// input conditions are not satisfied, Apply panics.) It returns a newly
// allocated slice where each element is the result of calling the function on
// successive elements of the slice.
func Apply(slice, function interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// ApplyInPlace is like Apply, but overwrites the slice rather than returning a
// newly allocated slice.
func ApplyInPlace(slice, function interface{}) { _ = "STUB: not implemented"; return }

// Choose takes a slice of type []T and a function of type func(T) bool. (If
// the input conditions are not satisfied, Choose panics.) It returns a newly
// allocated slice containing only those elements of the input slice that
// satisfy the function.
func Choose(slice, function interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// Drop takes a slice of type []T and a function of type func(T) bool. (If the
// input conditions are not satisfied, Drop panics.) It returns a newly
// allocated slice containing only those elements of the input slice that do
// not satisfy the function, that is, it removes elements that satisfy the
// function.
func Drop(slice, function interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// ChooseInPlace is like Choose, but overwrites the slice rather than returning
// a newly allocated slice. Since ChooseInPlace must modify the header of the
// slice to set the new length, it takes as argument a pointer to a slice
// rather than a slice.
func ChooseInPlace(pointerToSlice, function interface{}) { _ = "STUB: not implemented"; return }

// DropInPlace is like Drop, but overwrites the slice rather than returning a
// newly allocated slice. Since DropInPlace must modify the header of the slice
// to set the new length, it takes as argument a pointer to a slice rather than
// a slice.
func DropInPlace(pointerToSlice, function interface{}) { _ = "STUB: not implemented"; return }

func apply(slice, function interface{}, inPlace bool) interface{} {
	_ = "STUB: not implemented"
	// Special case for strings, very common.
	return nil
}

// Outside the loop to avoid one allocation.

func chooseOrDropInPlace(slice, function interface{}, truth bool) {
	_ = "STUB: not implemented"
	return
}

var boolType = reflect.ValueOf(true).Type()

func chooseOrDrop(slice, function interface{}, inPlace, truth bool) (interface{}, int) {
	_ = "STUB: not implemented"
	// Special case for strings, very common.
	return nil, 0
}

// Outside the loop to avoid one allocation.

// goodFunc verifies that the function satisfies the signature, represented as a slice of types.
// The last type is the single result type; the others are the input types.
// A final type of nil means any result type is accepted.
func goodFunc(fn reflect.Value, types ...reflect.Type) bool {
	_ = "STUB: not implemented"
	return false
}

// Last type is return, the rest are ins.
