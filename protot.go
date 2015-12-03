package protot

import (
	"encoding/json"
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

// CompareString work like NilString but also allow to specify comparison strategy.
type CompareString struct {
	NilString
	Compare string `protobuf:"varint,3,opt,name=compare" json:"compare,omitempty"`
}

// Reset implements proto.Message interface.
func (cs *CompareString) Reset() { *cs = CompareString{} }

// String implements proto.Message interface.
func (cs *CompareString) String() string { return proto.CompactTextString(cs) }

// ProtoMessage implements proto.Message interface.
func (*CompareString) ProtoMessage() {}

// NilInt64 represents a int64 that may be nil.
type NilInt64 struct {
	Int64 int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Valid bool  `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Reset implements proto.Message interface.
func (ni *NilInt64) Reset() { *ni = NilInt64{} }

// String implements proto.Message interface.
func (ni *NilInt64) String() string { return proto.CompactTextString(ni) }

// ProtoMessage implements proto.Message interface.
func (*NilInt64) ProtoMessage() {}

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

// NilFloat64 represents a flaot64 that may be nil.
type NilFloat64 struct {
	Float64 float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	Valid   bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Reset implements proto.Message interface.
func (nf *NilFloat64) Reset() { *nf = NilFloat64{} }

// String implements proto.Message interface.
func (nf *NilFloat64) String() string { return proto.CompactTextString(nf) }

// ProtoMessage implements proto.Message interface.
func (*NilFloat64) ProtoMessage() {}

// Float64Or returns given float64 value if receiver is nil or invalid.
func (nf *NilFloat64) Float64Or(or float64) float64 {
	if nf == nil {
		return or
	}
	if !nf.Valid {
		return or
	}

	return nf.Float64
}

// NilFloat64Range ...
type NilFloat64Range struct {
	From *NilFloat64 `protobuf:"varint,1,opt,name=from" json:"from,omitempty"`
	To   *NilFloat64 `protobuf:"varint,2,opt,name=to" json:"to,omitempty"`
}

// Reset implements proto.Message interface.
func (nfr *NilFloat64Range) Reset() { *nfr = NilFloat64Range{} }

// String implements proto.Message interface.
func (nfr *NilFloat64Range) String() string { return proto.CompactTextString(nfr) }

// ProtoMessage implements proto.Message interface.
func (*NilFloat64Range) ProtoMessage() {}

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
func (t *Timestamp) Less(t2 *Timestamp) bool {
	if t == nil {
		return true
	}
	if t2 == nil {
		return false
	}
	if t.Seconds < t2.Seconds {
		return true
	}
	if t.Seconds > t2.Seconds {
		return false
	}
	return t.Nanos < t2.Nanos
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
