package protot

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
)

// NilString represents a string that may be nil.
type NilString struct {
	String string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Valid  bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

func (m *NilString) Reset() { *m = NilString{} }

//func (m *NilString) String() string { return proto.CompactTextString(m) }
func (*NilString) ProtoMessage() {}

// StringOr returns given string value if receiver is nil or invalid.
func (ns *NilString) StringOr(or string) string {
	if ns == nil {
		return or
	}
	if !ns.Valid {
		return or
	}

	return ns.String
}

// MarshalJSON implements json.Marshaler interface.
func (ns *NilString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return nil, nil
	}

	return json.Marshal(ns.String)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ns *NilString) UnmarshalJSON(data []byte) error {
	if data == nil {
		ns.String, ns.Valid = "", false
		return nil
	}

	ns.Valid = true

	return json.Unmarshal(data, &ns.String)
}

// NilInt64 represents a int64 that may be nil.
type NilInt64 struct {
	Int64 int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Valid bool  `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Int64Or returns given int64 value if receiver is nil or invalid.
func (ni *NilInt64) Int64Or(or int64) int64 {
	if ni == nil {
		return or
	}
	if !ni.Valid {
		return or
	}

	return ni.Int64
}

// NilBool represents a bool that may be nil.
type NilBool struct {
	Bool  bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Valid bool `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Reset implements proto.Message interface.
func (nb *NilBool) Reset() { *nb = NilBool{} }

// String implements proto.Message interface.
func (nb *NilBool) String() string { return proto.CompactTextString(nb) }

// ProtoMessage implements proto.Message interface.
func (*NilBool) ProtoMessage() {}

// BoolOr returns given bool value if receiver is nil or invalid.
func (nb *NilBool) BoolOr(or bool) bool {
	if nb == nil {
		return or
	}
	if !nb.Valid {
		return or
	}

	return nb.Bool
}

// Timestamp ...
type Timestamp struct {
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	Nanos   int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}

// Reset implements proto.Message interface.
func (t *Timestamp) Reset() { *t = Timestamp{} }

// String implements proto.Message interface.
func (t *Timestamp) String() string { return proto.CompactTextString(t) }

// ProtoMessage implements proto.Message interface.
func (*Timestamp) ProtoMessage() {}

// Now returns the current time as a protobuf Timestamp.
func Now() *Timestamp {
	return TimeToTimestamp(time.Now().UTC())
}

// TimeToTimestamp converts a go Time to a protobuf Timestamp.
func TimeToTimestamp(t time.Time) *Timestamp {
	return &Timestamp{
		Seconds: t.UnixNano() / int64(time.Second),
		Nanos:   int32(t.UnixNano() % int64(time.Second)),
	}
}

// Time returns time.Time based on current Timestamp.
func (t *Timestamp) Time() time.Time {
	if t == nil {
		return time.Unix(0, 0).UTC()
	}

	return time.Unix(t.Seconds, int64(t.Nanos)).UTC()
}

// Less returns true if timestamp is before given one.
func (t *Timestamp) Less(i *Timestamp, j *Timestamp) bool {
	if i == nil {
		return true
	}
	if j == nil {
		return false
	}
	if i.Seconds < j.Seconds {
		return true
	}
	if i.Seconds > j.Seconds {
		return false
	}
	return i.Nanos < j.Nanos
}

// TimestampRange ...
type TimestampRange struct {
	From *Timestamp `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To   *Timestamp `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
}

// Reset implements proto.Message interface.
func (tr *TimestampRange) Reset() { *tr = TimestampRange{} }

// String implements proto.Message interface.
func (tr *TimestampRange) String() string { return proto.CompactTextString(tr) }

// ProtoMessage implements proto.Message interface.
func (*TimestampRange) ProtoMessage() {}
