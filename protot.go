package protot

import (
	pbts "github.com/golang/protobuf/ptypes/timestamp"
)

// Value ...
func (qs *QueryString) Value() string {
	if len(qs.Values) == 0 {
		return ""
	}

	return qs.Values[0]
}

// ExactString ...
func ExactString(s string) *QueryString {
	return &QueryString{
		Values: []string{s},
		Valid:  true,
		Type:   TextQueryType_EXACT,
	}
}

// Value ...
func (qi *QueryInt64) Value() int64 {
	if len(qi.Values) == 0 {
		return 0
	}

	return qi.Values[0]
}

// EqualInt64 ...
func EqualInt64(i int64) *QueryInt64 {
	return &QueryInt64{
		Values: []int64{i},
		Valid:  true,
		Type:   NumericQueryType_EQUAL,
	}
}

// Value ...
func (qf *QueryFloat64) Value() float64 {
	if len(qf.Values) == 0 {
		return 0.0
	}

	return qf.Values[0]
}

// Value returns first value or nil if none.
func (qt *QueryTimestamp) Value() *pbts.Timestamp {
	if len(qt.Values) == 0 {
		return nil
	}

	return qt.Values[0]
}

// BetweenTimestamp ...
func BetweenTimestamp(from, to *pbts.Timestamp) *QueryTimestamp {
	if from == nil || to == nil {
		return &QueryTimestamp{}
	}

	v := true
	if to.Seconds < from.Seconds {
		v = false
	}
	if to.Seconds == from.Seconds && to.Nanos < from.Nanos {
		v = false
	}
	return &QueryTimestamp{
		Values: []*pbts.Timestamp{from, to},
		Type:   NumericQueryType_BETWEEN,
		Valid:  v,
	}
}
