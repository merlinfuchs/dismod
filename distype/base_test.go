package distype

import (
	"encoding/json"
	"testing"
)

func TestNullable(t *testing.T) {
	type testStruct struct {
		Val Nullable[int] `json:"val,omitempty"`
	}

	tests := []struct {
		name     string
		input    testStruct
		expected string
	}{
		{
			name:     "emptry",
			input:    testStruct{},
			expected: `{"val":null}`,
		},
		{
			name: "null",
			input: testStruct{Val: Nullable[int]{
				Valid: false,
			}},
			expected: `{"val":null}`,
		},
		{
			name: "zero",
			input: testStruct{Val: Nullable[int]{
				Valid: true,
				Value: 0,
			}},
			expected: `{"val":0}`,
		},
		{
			name: "non-zero",
			input: testStruct{Val: Nullable[int]{
				Valid: true,
				Value: 1,
			}},
			expected: `{"val":1}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, data)
			}
		})
	}
}

func TestOptional(t *testing.T) {
	type testStruct struct {
		Val Optional[int] `json:"val,omitempty"`
	}

	tests := []struct {
		name     string
		input    testStruct
		expected string
	}{
		{
			name:     "emptry",
			input:    testStruct{},
			expected: `{}`,
		},
		{
			name:     "nil",
			input:    testStruct{Val: nil},
			expected: `{}`,
		},
		{
			name:     "zero",
			input:    testStruct{Val: new(int)},
			expected: `{"val":0}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, data)
			}
		})
	}
}

func TestNullish(t *testing.T) {
	type testStruct struct {
		Val Optional[Nullable[int]] `json:"val,omitempty"`
	}

	tests := []struct {
		name     string
		input    testStruct
		expected string
	}{
		{
			name:     "emptry",
			input:    testStruct{},
			expected: `{}`,
		},
		{
			name: "null",
			input: testStruct{Val: &Nullable[int]{
				Valid: false,
			}},
			expected: `{"val":null}`,
		},
		{
			name: "zero",
			input: testStruct{Val: &Nullable[int]{
				Valid: true,
				Value: 0,
			}},
			expected: `{"val":0}`,
		},
		{
			name: "non-zero",
			input: testStruct{Val: &Nullable[int]{
				Valid: true,
				Value: 1,
			}},
			expected: `{"val":1}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, data)
			}
		})
	}
}

func TestIntOrString(t *testing.T) {
	type testStruct struct {
		Val IntOrString `json:"val,omitempty"`
	}

	tests := []struct {
		name     string
		input    testStruct
		expected string
	}{
		{
			name:     "emptry",
			input:    testStruct{},
			expected: `{"val":""}`,
		},
		{
			name: "string",
			input: testStruct{Val: IntOrString{
				String: "test",
			}},
			expected: `{"val":"test"}`,
		},
		{
			name: "int",
			input: testStruct{Val: IntOrString{
				Int: 1,
			}},
			expected: `{"val":1}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, data)
			}

			var output testStruct
			err = json.Unmarshal(data, &output)
			if err != nil {
				t.Fatal(err)
			}

			if output.Val != test.input.Val {
				t.Fatalf("expected %v, got %v", test.input.Val, output.Val)
			}
		})
	}
}
