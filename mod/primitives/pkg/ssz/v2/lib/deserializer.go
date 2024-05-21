// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package ssz

import (
	"errors"
	"reflect"

	ssz "github.com/berachain/beacon-kit/mod/primitives/pkg/ssz"
)

type Deserializer struct{}

func NewDeserializer() Deserializer {
	return Deserializer{}
}

// UnmarshalSSZ is the top-level function to unmarshal an SSZ encoded buffer
// into the provided interface.
func (d *Deserializer) UnmarshalSSZ(
	val interface{},
	data []byte,
) (interface{}, error) {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	decodedValue, err := d.Unmarshal(v.Interface(), data)
	if err != nil {
		return nil, err
	}
	return decodedValue.Interface(), nil
}

// Unmarshal is a recursive function that determines the type of the value and
// unmarshals the data accordingly.
func (d *Deserializer) Unmarshal(
	c interface{},
	data []byte,
) (reflect.Value, error) {
	typ := reflect.TypeOf(c)
	val := reflect.ValueOf(c)
	k := typ.Kind()

	size := DetermineSize(val)
	buf := data[:size]

	if IsUintLike(k) {
		return RouteUintUnmarshal(k, buf), nil
	}
	switch k {
	case reflect.Bool:
		return reflect.ValueOf(ssz.UnmarshalBool[bool](buf)), nil
	case reflect.Ptr:
		elem, err := d.Unmarshal(typ.Elem(), data)
		if err != nil {
			return reflect.Value{}, err
		}
		ptr := reflect.New(typ.Elem())
		ptr.Elem().Set(elem)
		return ptr, nil
	case reflect.Slice, reflect.Array:
		return d.unmarshalArrayOrSlice(typ, val, data)
	// case reflect.Struct:
	// 	return d.unmarshalStruct(typ, data)
	default:
		return reflect.Value{}, errors.New("unsupported type")
	}
}

// unmarshalArrayOrSlice unmarshals an array or slice type.
func (d *Deserializer) unmarshalArrayOrSlice(
	typ reflect.Type,
	val reflect.Value,
	data []byte,
) (reflect.Value, error) {
	if typ.Elem().Kind() == reflect.Uint8 {
		return d.UnmarshalByteArray(typ, data), nil
	}

	// TODO: Code below will not function as expected due to unknown size

	//#nosec:G701 // if lenData is greater than int limit,
	// we cant call makeslice, which will panic with -ve nums.
	lenData := len(data) / int(DetermineSize(val.Elem()))
	slice := reflect.MakeSlice(typ, lenData, lenData)

	if IsNDimensionalArrayLike(typ) {
		prevDataIndex := uint64(0)
		for i := range val.Len() {
			// This doesnt work as we are not passing the struct field extras
			// like ssz-size
			size := DetermineSize(val.Index(i))
			cur := data[prevDataIndex:size]
			prevDataIndex += size

			elem, err := d.Unmarshal(
				typ.Elem(),
				cur)
			if err != nil {
				return reflect.Value{}, err
			}

			slice.Index(i).Set(elem)
		}
		return slice, nil
	}
	// This also wont work for the same reason as the above
	// for i := range lenData {
	// 	item := slice.Index(i)
	// 	size := DetermineSize(item)
	// 	elem, err := d.Unmarshal(
	// 		typ.Elem(),
	// 		data[i*int(size):(i+1)*int(size)])

	// 	if err != nil {
	// 		return reflect.Value{}, err
	// 	}
	// 	slice.Index(i).Set(elem)
	// }
	return slice, nil
}

func (d *Deserializer) UnmarshalByteArray(
	typ reflect.Type,
	data []byte,
) reflect.Value {
	return reflect.ValueOf(data).Convert(typ)
}

// todo
// // unmarshalStruct unmarshals a struct type.
// func (d *Deserializer) unmarshalStruct(typ reflect.Type, data []byte)
// (reflect.Value, error) {
// 	v := reflect.New(typ).Elem()
// 	offset := uint64(0)

// 	fixedParts := make(map[int][2]int) // map of [start, end] of fixed sizes
// 	variableParts := make(map[int]int) // map of [size] of variable sizes

// 	// Analyze and collect fixed and variable fields.
// 	for i := 0; i < v.NumField(); i++ {
// 		field := typ.Field(i)

// 		if hasUndefinedSizeTag(field) && isVariableSizeType(field.Type) {
// 			variableParts[i] = 0
// 		} else {
// 			size := determineFixedSize(v.Field(i), field.Type)
// 			if size == 0 {
// 				return reflect.Value{}, errors.New("unexpected 0 size")
// 			}
// 			fixedParts[i] = [2]int{int(offset), int(offset + size)}
// 			offset += size
// 		}
// 	}

// 	// Calculate sizes for variable parts from the fixed positions.
// 	for idx := range variableParts {
// 		if offset >= uint64(len(data)) {
// 			break
// 		}

// 		readOffset := fastssz.readOffset(data, offset)
// 		actualSize := determineSize(data[offset : offset+BytesPerLengthOffset])
// 		if (actualSize * BytesPerLengthOffset) != readOffset {
// 			return reflect.Value{}, errors.New("invalid size read from offset")
// 		}
// 		variableParts[idx] = int(actualSize)
// 		offset += BytesPerLengthOffset
// 	}

// 	// Unmarshal fixed parts
// 	for idx, span := range fixedParts {
// 		fieldData := data[span[0]:span[1]]
// 		fieldVal, err := d.Unmarshal(typ.Field(idx).Type, fieldData)
// 		if err != nil {
// 			return reflect.Value{}, err
// 		}
// 		v.Field(idx).Set(fieldVal)
// 	}

// 	// Unmarshal variable parts
// 	for idx, size := range variableParts {
// 		fieldData := data[offset : offset+uint64(size)]
// 		fieldVal, err := d.Unmarshal(typ.Field(idx).Type, fieldData)
// 		if err != nil {
// 			return reflect.Value{}, err
// 		}
// 		v.Field(idx).Set(fieldVal)
// 		offset += uint64(size)
// 	}

// 	return v, nil
// }
