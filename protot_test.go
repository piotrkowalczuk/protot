package protot

import (
	"reflect"
	"testing"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func TestParseString(t *testing.T) {
	cases := map[string]struct {
		given    string
		expected QueryString
	}{
		"exists": {
			given: "ex:",
			expected: QueryString{
				Values:   []string{""},
				Type:     TextQueryType_NOT_A_TEXT,
				Valid:    true,
				Negation: true,
			},
		},
		"not-exists": {
			given: "nex:",
			expected: QueryString{
				Values: []string{""},
				Type:   TextQueryType_NOT_A_TEXT,
				Valid:  true,
			},
		},
		"equal": {
			given: "eq:123",
			expected: QueryString{
				Values: []string{"123"},
				Type:   TextQueryType_EXACT,
				Valid:  true,
			},
		},
		"has-prefix": {
			given: "hp:New",
			expected: QueryString{
				Values: []string{"New"},
				Type:   TextQueryType_HAS_PREFIX,
				Valid:  true,
			},
		},
		"has-suffix": {
			given: "hs:New",
			expected: QueryString{
				Values: []string{"New"},
				Type:   TextQueryType_HAS_SUFFIX,
				Valid:  true,
			},
		},
		"substring": {
			given: "sub:anything",
			expected: QueryString{
				Values: []string{"anything"},
				Type:   TextQueryType_SUBSTRING,
				Valid:  true,
			},
		},
		"pattern": {
			given: "regex:.*",
			expected: QueryString{
				Values: []string{".*"},
				Type:   TextQueryType_PATTERN,
				Valid:  true,
			},
		},
		"max-length": {
			given: "maxl:4",
			expected: QueryString{
				Values: []string{"4"},
				Type:   TextQueryType_MAX_LENGTH,
				Valid:  true,
			},
		},
		"min-length": {
			given: "minl:555",
			expected: QueryString{
				Values: []string{"555"},
				Type:   TextQueryType_MIN_LENGTH,
				Valid:  true,
			},
		},
		"empty": {
			given:    "",
			expected: QueryString{},
		},
		"without-condition": {
			given: "text",
			expected: QueryString{
				Values: []string{"text"},
				Type:   TextQueryType_EXACT,
				Valid:  true,
			},
		},
		"with-condition-but-without-value": {
			given: "neq:",
			expected: QueryString{
				Values:   []string{""},
				Type:     TextQueryType_EXACT,
				Valid:    true,
				Negation: true,
			},
		},
	}

CasesLoop:
	for hint, c := range cases {
		got := ParseString(c.given)

		if got == nil {
			t.Errorf("unexpected nil")
			continue CasesLoop
		}
		if !reflect.DeepEqual(c.expected, *got) {
			t.Errorf("%s: wrong output,\nexpected:\n	%v\nbut got:\n	%v\n", hint, &c.expected, got)
		}
	}
}

func TestExactString(t *testing.T) {
	es := ExactString("John")

	if es.Negation {
		t.Errorf("unexpected negation")
	}
	if es.Value() != "John" {
		t.Errorf("unexpected value")
	}
	if !es.Valid {
		t.Errorf("expected to be valid")
	}
}

func TestBetweenTimestamp(t *testing.T) {
	cases := map[string]struct {
		from     *timestamp.Timestamp
		to       *timestamp.Timestamp
		expected QueryTimestamp
	}{
		"valid": {
			from: &timestamp.Timestamp{Seconds: 0, Nanos: 0},
			to:   &timestamp.Timestamp{Seconds: 0, Nanos: 1},
			expected: QueryTimestamp{
				Valid:    true,
				Negation: false,
				Type:     NumericQueryType_BETWEEN,
				Values: []*timestamp.Timestamp{
					&timestamp.Timestamp{Seconds: 0, Nanos: 0},
					&timestamp.Timestamp{Seconds: 0, Nanos: 1},
				},
			},
		},
		"after-first": {
			from: &timestamp.Timestamp{Seconds: 1, Nanos: 0},
			to:   &timestamp.Timestamp{Seconds: 0, Nanos: 0},
			expected: QueryTimestamp{
				Valid: false,
				Type:  NumericQueryType_BETWEEN,
				Values: []*timestamp.Timestamp{
					&timestamp.Timestamp{Seconds: 1, Nanos: 0},
					&timestamp.Timestamp{Seconds: 0, Nanos: 0},
				},
			},
		},
		"after-first-seconds": {
			from: &timestamp.Timestamp{Seconds: 1, Nanos: 1},
			to:   &timestamp.Timestamp{Seconds: 1, Nanos: 0},
			expected: QueryTimestamp{
				Valid: false,
				Type:  NumericQueryType_BETWEEN,
				Values: []*timestamp.Timestamp{
					&timestamp.Timestamp{Seconds: 1, Nanos: 1},
					&timestamp.Timestamp{Seconds: 1, Nanos: 0},
				},
			},
		},
		"nil-arguments": {
			from:     nil,
			to:       nil,
			expected: QueryTimestamp{},
		},
		"nil-argument-first": {
			from:     nil,
			to:       &timestamp.Timestamp{Seconds: 0, Nanos: 1},
			expected: QueryTimestamp{},
		},
		"nil-argument-second": {
			from:     &timestamp.Timestamp{Seconds: 0, Nanos: 1},
			to:       nil,
			expected: QueryTimestamp{},
		},
	}

	for hint, c := range cases {
		bt := BetweenTimestamp(c.from, c.to)
		if !reflect.DeepEqual(c.expected, *bt) {
			t.Errorf("%s: unexpected output, expected:\n%v\ngot:\n%v\n", hint, c.expected, *bt)
		}
	}
}

func TestQueryInt64_Value(t *testing.T) {
	cases := map[string]struct {
		given    QueryInt64
		expected int64
	}{
		"single": {
			given: QueryInt64{
				Values: []int64{1},
				Valid:  true,
				Type:   NumericQueryType_EQUAL,
			},
			expected: 1,
		},
		"none": {
			given: QueryInt64{
				Valid: true,
				Type:  NumericQueryType_EQUAL,
			},
			expected: 0,
		},
		"multiple": {
			given: QueryInt64{
				Values: []int64{3, 2, 1},
				Valid:  true,
				Type:   NumericQueryType_EQUAL,
			},
			expected: 3,
		},
	}

	for hint, c := range cases {
		if c.given.Value() != c.expected {
			t.Errorf("%s: unexpected value, expected %d but got %d", hint, c.expected, c.given.Value())
		}
	}
}

func TestEqualInt64(t *testing.T) {
	es := EqualInt64(888)

	if es.Negation {
		t.Errorf("unexpected negation")
	}
	if es.Value() != 888 {
		t.Errorf("unexpected value")
	}
	if !es.Valid {
		t.Errorf("expected to be valid")
	}
}

func TestParseInt64(t *testing.T) {
	cases := map[string]struct {
		given    string
		expected QueryInt64
	}{
		"exists": {
			given: "ex:",
			expected: QueryInt64{
				Values:   []int64{},
				Type:     NumericQueryType_NOT_A_NUMBER,
				Valid:    true,
				Negation: true,
			},
		},
		"not-exists": {
			given: "nex:",
			expected: QueryInt64{
				Values: []int64{},
				Type:   NumericQueryType_NOT_A_NUMBER,
				Valid:  true,
			},
		},
		"equal": {
			given: "eq:123",
			expected: QueryInt64{
				Values: []int64{123},
				Type:   NumericQueryType_EQUAL,
				Valid:  true,
			},
		},
		"not-equal": {
			given: "neq:123",
			expected: QueryInt64{
				Values:   []int64{123},
				Type:     NumericQueryType_EQUAL,
				Valid:    true,
				Negation: true,
			},
		},
		"greater": {
			given: "gt:555",
			expected: QueryInt64{
				Values: []int64{555},
				Type:   NumericQueryType_GREATER,
				Valid:  true,
			},
		},
		"greater-equal": {
			given: "gte:666",
			expected: QueryInt64{
				Values: []int64{666},
				Type:   NumericQueryType_GREATER_EQUAL,
				Valid:  true,
			},
		},
		"lesser": {
			given: "lt:777",
			expected: QueryInt64{
				Values: []int64{777},
				Type:   NumericQueryType_LESS,
				Valid:  true,
			},
		},
		"lesser-equal": {
			given: "lte:888",
			expected: QueryInt64{
				Values: []int64{888},
				Type:   NumericQueryType_LESS_EQUAL,
				Valid:  true,
			},
		},
		"between": {
			given: "bw:111,222",
			expected: QueryInt64{
				Values: []int64{111, 222},
				Type:   NumericQueryType_BETWEEN,
				Valid:  true,
			},
		},
		"not-between": {
			given: "nbw:111,222",
			expected: QueryInt64{
				Values:   []int64{111, 222},
				Type:     NumericQueryType_BETWEEN,
				Valid:    true,
				Negation: true,
			},
		},
	}

CasesLoop:
	for hint, c := range cases {
		got, err := ParseInt64(c.given)
		if err != nil {
			t.Errorf("%s: unexpected error: %s", hint, err.Error())
			continue CasesLoop
		}
		if got == nil {
			t.Errorf("%s: unexpected nil", hint)
			continue CasesLoop
		}
		if !reflect.DeepEqual(c.expected, *got) {
			t.Errorf("%s: wrong output,\nexpected:\n	%v\nbut got:\n	%v\n", hint, &c.expected, got)
		}
	}
}

func TestParseInt64_text(t *testing.T) {
	got, err := ParseInt64("ne:long-text")
	if err == nil {
		t.Fatalf("expected error")
	}
	if got != nil {
		t.Fatalf("expected nil")
	}
}
