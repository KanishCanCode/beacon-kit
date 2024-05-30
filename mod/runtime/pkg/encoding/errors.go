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

package encoding

import "github.com/berachain/beacon-kit/mod/errors"

var (
	// ErrNilBeaconBlockInRequest is an error for when
	// the beacon block in an abci request is nil.
	ErrNilBeaconBlockInRequest = errors.New("nil beacon block in abci request")

	// ErrNoBeaconBlockInRequest is an error for when
	// there is no beacon block in an abci request.
	ErrNoBeaconBlockInRequest = errors.New("no beacon block in abci request")

	// ErrBzIndexOutOfBounds is an error for when the index
	// is out of bounds.
	ErrBzIndexOutOfBounds = errors.New("bzIndex out of bounds")

	// ErrNilABCIRequest is an error for when the abci request
	// is nil.
	ErrNilABCIRequest = errors.New("nil abci request")

	// ErrInvalidType is an error for when the type is invalid.
	ErrInvalidType = errors.New("invalid type")

	// ErrNilBlobSidecarsInRequest is an error for when
	// the blob sidecars in an abci request is nil.
	ErrNilBlobSidecarsInRequest = errors.New(
		"nil blob sidecars in abci request",
	)
)
