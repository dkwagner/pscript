package util

import "testing"

func TestIsLetter(t *testing.T) {

	isLetterA := IsLetter('a')
	isLetterZ := IsLetter('Z')
	isLettera := IsLetter('a')
	isLetterz := IsLetter('z')
	isNotLetter := IsLetter('0')

	if !isLetterA {
		t.Error("Test ch A: expected true, got false")
	}

	if !isLetterZ {
		t.Error("Test ch Z: expected true, got false")
	}

	if !isLettera {
		t.Error("Test ch a: expected true, got false")
	}

	if !isLetterz {
		t.Error("Test ch z: expected true, got false")
	}

	if isNotLetter {
		t.Error("Test ch EOF: expected false, got true")
	}
}

func TestIsDigit(t *testing.T) {

	isDigit0 := IsDigit('0')
	isDigit9 := IsDigit('9')
	isNotDigit := IsDigit('/')

	if !isDigit0 {
		t.Error("Test ch 0: expected true, got false")
	}

	if !isDigit9 {
		t.Error("Test ch 9: expected true, got false")
	}

	if isNotDigit {
		t.Error("Test ch /: expected false, got true")
	}
}

func TestIsWhitespace(t *testing.T) {

	tests := []struct {
		ch       byte
		expected bool
	}{
		{ch: ' ', expected: true},
		{ch: '\t', expected: true},
		{ch: '\n', expected: true},
		{ch: '\r', expected: true},
		{ch: '\v', expected: true},
		{ch: '\f', expected: true},
		{ch: 'a', expected: false},
	}

	for _, tt := range tests {
		actual := IsWhitespace(tt.ch)

		if actual != tt.expected {
			t.Errorf("Test '%q' failed: expected %t, got %t",
				tt.ch, tt.expected, actual)
		}
	}
}

func TestContainsString(t *testing.T) {
	testList := []string{"test1", "matches", "test2"}

	tests := []struct {
		testString string
		expected   bool
	}{
		{testString: "matches", expected: true},
		{testString: "no match", expected: false},
	}

	for _, tt := range tests {
		actual := ContainsString(tt.testString, testList)

		if actual != tt.expected {
			t.Errorf("Test [%s] Failed: expected %t, got %t", tt.testString, tt.expected, actual)
		}
	}
}

func TestContainsByte(t *testing.T) {
	testList := []byte{'(', ')', '{', '}', ',', '[', ']'}

	tests := []struct {
		testByte byte
		expected bool
	}{
		{testByte: '[', expected: true},
		{testByte: 'n', expected: false},
	}

	for _, tt := range tests {
		actual := ContainsByte(tt.testByte, testList)

		if actual != tt.expected {
			t.Errorf("Test [%q] Failed: expected %t, got %t", tt.testByte, tt.expected, actual)
		}
	}
}
