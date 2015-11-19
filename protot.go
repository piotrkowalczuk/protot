package protot

import "time"

// NilString represents a string that may be nil.
type NilString struct {
	String string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Valid  bool   `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// StringOr returns given string value if receiver is nil or invalid.
func (ns NilString) StringOr(or string) string {
	if !ns.Valid {
		return or
	}

	return ns.String
}

//func (m *NilString) Reset()         { *m = NilString{} }
//func (m *NilString) String() string { return proto.CompactTextString(m) }
//func (*NilString) ProtoMessage()    {}

// NilInt64 represents a int64 that may be nil.
type NilInt64 struct {
	Int64 int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Valid bool  `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// Int64Or returns given int64 value if receiver is nil or invalid.
func (ni NilInt64) Int64Or(or int64) int64 {
	if !ni.Valid {
		return or
	}

	return ni.Int64
}

//func (m *NilInt64) Reset()         { *m = NilInt64{} }
//func (m *NilInt64) String() string { return proto.CompactTextString(m) }
//func (*NilInt64) ProtoMessage()    {}

// NilBool represents a bool that may be nil.
type NilBool struct {
	Bool  bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Valid bool `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
}

// BoolOr returns given bool value if receiver is nil or invalid.
func (nb NilBool) BoolOr(or bool) bool {
	if !nb.Valid {
		return or
	}

	return nb.Bool
}

//func (m *NilBool) Reset()         { *m = NilBool{} }
//func (m *NilBool) String() string { return proto.CompactTextString(m) }
//func (*NilBool) ProtoMessage()    {}

// Timestamp ...
type Timestamp struct {
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	Nanos   int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}

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

//func (m *Timestamp) Reset()         { *m = Timestamp{} }
//func (m *Timestamp) String() string { return proto.CompactTextString(m) }
//func (*Timestamp) ProtoMessage()    {}
