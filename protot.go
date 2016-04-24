// Package provides set of types that helps to build complex protobuf messages that can express conditional statements.
package protot

import (
	"fmt"
	"strconv"
	"strings"

	pbts "github.com/golang/protobuf/ptypes/timestamp"
)

const (
	arraySeparator = ","
	// Exists ...
	Exists = "ex"
	// NotExists ...
	NotExists = "nex"
	// Equal ...
	Equal = "eq"
	// NotEqual ...
	NotEqual = "neq"
	// GreaterThan ...
	GreaterThan = "gt"
	// GreaterThanOrEqual ...
	GreaterThanOrEqual = "gte"
	// LessThan ...
	LessThan = "lt"
	// LessThanOrEqual ...
	LessThanOrEqual = "lte"
	// Between ...
	Between = "bw"
	// NotBetween ...
	NotBetween = "nbw"
	// HasPrefix ...
	HasPrefix = "hp"
	// HasSuffix ...
	HasSuffix = "hs"
	// In ...
	In = "in"
	// Substring ...
	Substring = "sub"
	// Pattern ...
	Pattern = "regex"
	// MinLength ...
	MinLength = "minl"
	// MaxLength ...
	MaxLength = "maxl"
)

var (
	prefixes = map[string]string{
		Exists:             Exists + ":",
		NotExists:          NotExists + ":",
		Equal:              Equal + ":",
		NotEqual:           NotEqual + ":",
		GreaterThan:        GreaterThan + ":",
		GreaterThanOrEqual: GreaterThanOrEqual + ":",
		LessThan:           LessThan + ":",
		LessThanOrEqual:    LessThanOrEqual + ":",
		Between:            Between + ":",
		NotBetween:         NotBetween + ":",
		HasPrefix:          HasPrefix + ":",
		HasSuffix:          HasSuffix + ":",
		In:                 In + ":",
		Substring:          Substring + ":",
		Pattern:            Pattern + ":",
		MinLength:          MinLength + ":",
		MaxLength:          MaxLength + ":",
	}
)

// Value returns first value or empty string if none.
func (qs *QueryString) Value() string {
	if len(qs.Values) == 0 {
		return ""
	}

	return qs.Values[0]
}

// ParseString allocates new QueryString object based on given string.
// If string is prefixed with known operator e.g. 'hp:New'
// returned object will get same type.
func ParseString(s string) *QueryString {
	if s == "" {
		return &QueryString{}
	}

	for c, p := range prefixes {
		if strings.HasPrefix(s, p) {
			var (
				t TextQueryType
				n bool
			)
			switch c {
			case Exists:
				t = TextQueryType_NOT_A_TEXT
				n = true
			case NotExists:
				t = TextQueryType_NOT_A_TEXT
			case Equal:
				t = TextQueryType_EXACT
			case NotEqual:
				t = TextQueryType_EXACT
				n = true
			case HasPrefix:
				t = TextQueryType_HAS_PREFIX
			case HasSuffix:
				t = TextQueryType_HAS_SUFFIX
			case Substring:
				t = TextQueryType_SUBSTRING
			case Pattern:
				t = TextQueryType_PATTERN
			case MinLength:
				t = TextQueryType_MIN_LENGTH
			case MaxLength:
				t = TextQueryType_MAX_LENGTH
			}
			return &QueryString{
				Values:   strings.Split(strings.TrimLeft(s, p), arraySeparator),
				Type:     t,
				Negation: n,
				Valid:    true,
			}
		}
	}
	return &QueryString{
		Values: strings.Split(s, arraySeparator),
		Type:   TextQueryType_EXACT,
		Valid:  true,
	}
}

// ExactString ...
func ExactString(s string) *QueryString {
	return &QueryString{
		Values: []string{s},
		Valid:  true,
		Type:   TextQueryType_EXACT,
	}
}

// EqualInt64 allocates valid QueryInt64 object of type equal with given value.
func EqualInt64(i int64) *QueryInt64 {
	return &QueryInt64{
		Values: []int64{i},
		Valid:  true,
		Type:   NumericQueryType_EQUAL,
	}
}

// BetweenInt64 allocates valid QueryInt64 object of type between with given values.
func BetweenInt64(a, b int64) *QueryInt64 {
	return &QueryInt64{
		Values: []int64{a, b},
		Valid:  true,
		Type:   NumericQueryType_BETWEEN,
	}
}

// GreaterInt64 allocates valid QueryInt64 object of type greater with given value.
func GreaterInt64(i int64) *QueryInt64 {
	return &QueryInt64{
		Values: []int64{i},
		Valid:  true,
		Type:   NumericQueryType_GREATER,
	}
}

// LessInt64 allocates valid QueryInt64 object of type less with given value.
func LessInt64(i int64) *QueryInt64 {
	return &QueryInt64{
		Values: []int64{i},
		Valid:  true,
		Type:   NumericQueryType_LESS,
	}
}

// Value ...
func (qi *QueryInt64) Value() int64 {
	if len(qi.Values) == 0 {
		return 0
	}

	return qi.Values[0]
}

func ParseInt64(s string) (*QueryInt64, error) {
	if s == "" {
		return &QueryInt64{}, nil
	}
	var (
		t        NumericQueryType
		n        bool
		incoming []string
	)
	for c, p := range prefixes {
		if strings.HasPrefix(s, p) {
			switch c {
			case Exists:
				t = NumericQueryType_NOT_A_NUMBER
				n = true
			case NotExists:
				t = NumericQueryType_NOT_A_NUMBER
			case Equal:
				t = NumericQueryType_EQUAL
			case NotEqual:
				t = NumericQueryType_EQUAL
				n = true
			case GreaterThan:
				t = NumericQueryType_GREATER
			case GreaterThanOrEqual:
				t = NumericQueryType_GREATER_EQUAL
			case LessThan:
				t = NumericQueryType_LESS
			case LessThanOrEqual:
				t = NumericQueryType_LESS_EQUAL
			case Between:
				t = NumericQueryType_BETWEEN
			case NotBetween:
				t = NumericQueryType_BETWEEN
				n = true
			}

			incoming = strings.Split(strings.TrimLeft(s, p), arraySeparator)

		}
	}
	if len(incoming) == 0 {
		incoming = strings.Split(s, arraySeparator)

	}

	outgoing := make([]int64, 0, len(incoming))
	for i, v := range incoming {
		if v == "" {
			break
		}
		vv, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("protot: query int64 parsing error for valur %d: %s", i, err.Error())
		}
		outgoing = append(outgoing, vv)
	}
	return &QueryInt64{
		Values:   outgoing,
		Type:     t,
		Negation: n,
		Valid:    true,
	}, nil
}

// EqualFloat64 allocates valid QueryFloat64 object of type equal with given value.
func EqualFloat64(i float64) *QueryFloat64 {
	return &QueryFloat64{
		Values: []float64{i},
		Valid:  true,
		Type:   NumericQueryType_EQUAL,
	}
}

// Value returns first available value or 0 if none available.
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

// BetweenTimestamp allocates valid QueryTimestamp object if both timestamps are not nil
// and first is before the second.
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
