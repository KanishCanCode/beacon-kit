// Code generated by fastssz. DO NOT EDIT.
// Hash: 51760ae3ada8d13d59ffa4b38767685ad249bb416c6a9890f69e77b8a9128146
// Version: 0.1.3
package types

import (
	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the ExecutableDataDeneb object
func (e *ExecutableDataDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ExecutableDataDeneb object to a target array
func (e *ExecutableDataDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(528)

	// Field (0) 'ParentHash'
	dst = append(dst, e.ParentHash[:]...)

	// Field (1) 'FeeRecipient'
	dst = append(dst, e.FeeRecipient[:]...)

	// Field (2) 'StateRoot'
	dst = append(dst, e.StateRoot[:]...)

	// Field (3) 'ReceiptsRoot'
	dst = append(dst, e.ReceiptsRoot[:]...)

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("ExecutableDataDeneb.LogsBloom", size, 256)
		return
	}
	dst = append(dst, e.LogsBloom...)

	// Field (5) 'Random'
	dst = append(dst, e.Random[:]...)

	// Field (6) 'Number'
	dst = ssz.MarshalUint64(dst, uint64(e.Number))

	// Field (7) 'GasLimit'
	dst = ssz.MarshalUint64(dst, uint64(e.GasLimit))

	// Field (8) 'GasUsed'
	dst = ssz.MarshalUint64(dst, uint64(e.GasUsed))

	// Field (9) 'Timestamp'
	dst = ssz.MarshalUint64(dst, uint64(e.Timestamp))

	// Offset (10) 'ExtraData'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.ExtraData)

	// Field (11) 'BaseFeePerGas'
	dst = append(dst, e.BaseFeePerGas[:]...)

	// Field (12) 'BlockHash'
	dst = append(dst, e.BlockHash[:]...)

	// Offset (13) 'Transactions'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(e.Transactions); ii++ {
		offset += 4
		offset += len(e.Transactions[ii])
	}

	// Offset (14) 'Withdrawals'
	dst = ssz.WriteOffset(dst, offset)

	// Field (15) 'BlobGasUsed'
	dst = ssz.MarshalUint64(dst, uint64(e.BlobGasUsed))

	// Field (16) 'ExcessBlobGas'
	dst = ssz.MarshalUint64(dst, uint64(e.ExcessBlobGas))

	// Field (10) 'ExtraData'
	if size := len(e.ExtraData); size > 32 {
		err = ssz.ErrBytesLengthFn("ExecutableDataDeneb.ExtraData", size, 32)
		return
	}
	dst = append(dst, e.ExtraData...)

	// Field (13) 'Transactions'
	if size := len(e.Transactions); size > 1048576 {
		err = ssz.ErrListTooBigFn("ExecutableDataDeneb.Transactions", size, 1048576)
		return
	}
	{
		offset = 4 * len(e.Transactions)
		for ii := 0; ii < len(e.Transactions); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(e.Transactions[ii])
		}
	}
	for ii := 0; ii < len(e.Transactions); ii++ {
		if size := len(e.Transactions[ii]); size > 1073741824 {
			err = ssz.ErrBytesLengthFn("ExecutableDataDeneb.Transactions[ii]", size, 1073741824)
			return
		}
		dst = append(dst, e.Transactions[ii]...)
	}

	// Field (14) 'Withdrawals'
	if size := len(e.Withdrawals); size > 16 {
		err = ssz.ErrListTooBigFn("ExecutableDataDeneb.Withdrawals", size, 16)
		return
	}
	for ii := 0; ii < len(e.Withdrawals); ii++ {
		if dst, err = e.Withdrawals[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the ExecutableDataDeneb object
func (e *ExecutableDataDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 528 {
		return ssz.ErrSize
	}

	tail := buf
	var o10, o13, o14 uint64

	// Field (0) 'ParentHash'
	copy(e.ParentHash[:], buf[0:32])

	// Field (1) 'FeeRecipient'
	copy(e.FeeRecipient[:], buf[32:52])

	// Field (2) 'StateRoot'
	copy(e.StateRoot[:], buf[52:84])

	// Field (3) 'ReceiptsRoot'
	copy(e.ReceiptsRoot[:], buf[84:116])

	// Field (4) 'LogsBloom'
	if cap(e.LogsBloom) == 0 {
		e.LogsBloom = make([]byte, 0, len(buf[116:372]))
	}
	e.LogsBloom = append(e.LogsBloom, buf[116:372]...)

	// Field (5) 'Random'
	copy(e.Random[:], buf[372:404])

	// Field (6) 'Number'
	e.Number = math.U64(ssz.UnmarshallUint64(buf[404:412]))

	// Field (7) 'GasLimit'
	e.GasLimit = math.U64(ssz.UnmarshallUint64(buf[412:420]))

	// Field (8) 'GasUsed'
	e.GasUsed = math.U64(ssz.UnmarshallUint64(buf[420:428]))

	// Field (9) 'Timestamp'
	e.Timestamp = math.U64(ssz.UnmarshallUint64(buf[428:436]))

	// Offset (10) 'ExtraData'
	if o10 = ssz.ReadOffset(buf[436:440]); o10 > size {
		return ssz.ErrOffset
	}

	if o10 < 528 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (11) 'BaseFeePerGas'
	copy(e.BaseFeePerGas[:], buf[440:472])

	// Field (12) 'BlockHash'
	copy(e.BlockHash[:], buf[472:504])

	// Offset (13) 'Transactions'
	if o13 = ssz.ReadOffset(buf[504:508]); o13 > size || o10 > o13 {
		return ssz.ErrOffset
	}

	// Offset (14) 'Withdrawals'
	if o14 = ssz.ReadOffset(buf[508:512]); o14 > size || o13 > o14 {
		return ssz.ErrOffset
	}

	// Field (15) 'BlobGasUsed'
	e.BlobGasUsed = math.U64(ssz.UnmarshallUint64(buf[512:520]))

	// Field (16) 'ExcessBlobGas'
	e.ExcessBlobGas = math.U64(ssz.UnmarshallUint64(buf[520:528]))

	// Field (10) 'ExtraData'
	{
		buf = tail[o10:o13]
		if len(buf) > 32 {
			return ssz.ErrBytesLength
		}
		if cap(e.ExtraData) == 0 {
			e.ExtraData = make([]byte, 0, len(buf))
		}
		e.ExtraData = append(e.ExtraData, buf...)
	}

	// Field (13) 'Transactions'
	{
		buf = tail[o13:o14]
		num, err := ssz.DecodeDynamicLength(buf, 1048576)
		if err != nil {
			return err
		}
		e.Transactions = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 1073741824 {
				return ssz.ErrBytesLength
			}
			if cap(e.Transactions[indx]) == 0 {
				e.Transactions[indx] = make([]byte, 0, len(buf))
			}
			e.Transactions[indx] = append(e.Transactions[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (14) 'Withdrawals'
	{
		buf = tail[o14:]
		num, err := ssz.DivideInt2(len(buf), 44, 16)
		if err != nil {
			return err
		}
		e.Withdrawals = make([]*engineprimitives.Withdrawal, num)
		for ii := 0; ii < num; ii++ {
			if e.Withdrawals[ii] == nil {
				e.Withdrawals[ii] = new(engineprimitives.Withdrawal)
			}
			if err = e.Withdrawals[ii].UnmarshalSSZ(buf[ii*44 : (ii+1)*44]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ExecutableDataDeneb object
func (e *ExecutableDataDeneb) SizeSSZ() (size int) {
	size = 528

	// Field (10) 'ExtraData'
	size += len(e.ExtraData)

	// Field (13) 'Transactions'
	for ii := 0; ii < len(e.Transactions); ii++ {
		size += 4
		size += len(e.Transactions[ii])
	}

	// Field (14) 'Withdrawals'
	size += len(e.Withdrawals) * 44

	return
}

// HashTreeRoot ssz hashes the ExecutableDataDeneb object
func (e *ExecutableDataDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ExecutableDataDeneb object with a hasher
func (e *ExecutableDataDeneb) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ParentHash'
	hh.PutBytes(e.ParentHash[:])

	// Field (1) 'FeeRecipient'
	hh.PutBytes(e.FeeRecipient[:])

	// Field (2) 'StateRoot'
	hh.PutBytes(e.StateRoot[:])

	// Field (3) 'ReceiptsRoot'
	hh.PutBytes(e.ReceiptsRoot[:])

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("ExecutableDataDeneb.LogsBloom", size, 256)
		return
	}
	hh.PutBytes(e.LogsBloom)

	// Field (5) 'Random'
	hh.PutBytes(e.Random[:])

	// Field (6) 'Number'
	hh.PutUint64(uint64(e.Number))

	// Field (7) 'GasLimit'
	hh.PutUint64(uint64(e.GasLimit))

	// Field (8) 'GasUsed'
	hh.PutUint64(uint64(e.GasUsed))

	// Field (9) 'Timestamp'
	hh.PutUint64(uint64(e.Timestamp))

	// Field (10) 'ExtraData'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(e.ExtraData))
		if byteLen > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(e.ExtraData)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (32+31)/32)
	}

	// Field (11) 'BaseFeePerGas'
	hh.PutBytes(e.BaseFeePerGas[:])

	// Field (12) 'BlockHash'
	hh.PutBytes(e.BlockHash[:])

	// Field (13) 'Transactions'
	{
		subIndx := hh.Index()
		num := uint64(len(e.Transactions))
		if num > 1048576 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range e.Transactions {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 1073741824 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (1073741824+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1048576)
	}

	// Field (14) 'Withdrawals'
	{
		subIndx := hh.Index()
		num := uint64(len(e.Withdrawals))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range e.Withdrawals {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 16)
	}

	// Field (15) 'BlobGasUsed'
	hh.PutUint64(uint64(e.BlobGasUsed))

	// Field (16) 'ExcessBlobGas'
	hh.PutUint64(uint64(e.ExcessBlobGas))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ExecutableDataDeneb object
func (e *ExecutableDataDeneb) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(e)
}
