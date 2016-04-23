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
		"equal": {
			given: "eq:123",
			expected: QueryString{
				Values: []string{"123"},
				Type:   TextQueryType_EXACT,
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
