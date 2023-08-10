package timers

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"github.com/algorand/msgp/msgp"
)

// The following msgp objects are implemented in this file:
// TimeoutType
//      |-----> MarshalMsg
//      |-----> CanMarshalMsg
//      |-----> (*) UnmarshalMsg
//      |-----> (*) CanUnmarshalMsg
//      |-----> Msgsize
//      |-----> MsgIsZero
//      |-----> TimeoutTypeMaxSize()
//

// MarshalMsg implements msgp.Marshaler
func (z TimeoutType) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt8(o, int8(z))
	return
}

func (_ TimeoutType) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(TimeoutType)
	if !ok {
		_, ok = (z).(*TimeoutType)
	}
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TimeoutType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int8
		zb0001, bts, err = msgp.ReadInt8Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = TimeoutType(zb0001)
	}
	o = bts
	return
}

func (_ *TimeoutType) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*TimeoutType)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TimeoutType) Msgsize() (s int) {
	s = msgp.Int8Size
	return
}

// MsgIsZero returns whether this is a zero value
func (z TimeoutType) MsgIsZero() bool {
	return z == 0
}

// MaxSize returns a maximum valid message size for this message type
func TimeoutTypeMaxSize() (s int) {
	s = msgp.Int8Size
	return
}