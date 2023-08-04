package test

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"github.com/algorand/msgp/msgp"
)

// The following msgp objects are implemented in this file:
// testSlice
//     |-----> MarshalMsg
//     |-----> CanMarshalMsg
//     |-----> (*) UnmarshalMsg
//     |-----> (*) CanUnmarshalMsg
//     |-----> Msgsize
//     |-----> MsgIsZero
//     |-----> TestSliceMaxSize()
//

// MarshalMsg implements msgp.Marshaler
func (z testSlice) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	if z == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendArrayHeader(o, uint32(len(z)))
	}
	for za0001 := range z {
		o = msgp.AppendUint64(o, z[za0001])
	}
	return
}

func (_ testSlice) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(testSlice)
	if !ok {
		_, ok = (z).(*testSlice)
	}
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *testSlice) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 int
	var zb0003 bool
	zb0002, zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0002 > 16 {
		err = msgp.ErrOverflow(uint64(zb0002), uint64(16))
		err = msgp.WrapError(err)
		return
	}
	if zb0003 {
		(*z) = nil
	} else if (*z) != nil && cap((*z)) >= zb0002 {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(testSlice, zb0002)
	}
	for zb0001 := range *z {
		(*z)[zb0001], bts, err = msgp.ReadUint64Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	o = bts
	return
}

func (_ *testSlice) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*testSlice)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z testSlice) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (len(z) * (msgp.Uint64Size))
	return
}

// MsgIsZero returns whether this is a zero value
func (z testSlice) MsgIsZero() bool {
	return len(z) == 0
}

// MaxSize returns a maximum valid message size for this message type
func TestSliceMaxSize() (s int) {
	// Calculating size of slice: z
	s += msgp.ArrayHeaderSize + ((16) * (msgp.Uint64Size))
	return
}
