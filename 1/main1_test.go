package main

import (
	"fmt"
	"testing"
)

func TestDeterminesVariableTypeAndDisplaysIt(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{"testInt", 42, "type of testInt: int"},
		{"testString", "hello", "type of testString: string"},
		{"testBool", true, "type of testBool: bool"},
		{"testFloat", 3.14, "type of testFloat: float64"},
		{"testComplex", 1 + 2i, "type of testComplex: complex128"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeterminesVariableTypeAndDisplaysIt(tt.name, tt.input)
			if got != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestConvertsVariablesToStringAndConcatenatesIntoSingleString(t *testing.T) {
	tests := []struct {
		name     string
		inputs   []interface{}
		expected string
	}{
		{
			"convert test",
			[]interface{}{20, 0xFFFF, "hello", true, 1 + 1i},
			"2065535hellotrue(1+1i)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertsVariablesToStringAndConcatenatesIntoSingleString(tt.inputs...)
			if got != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestConvertStringToRuneSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []rune
	}{
		{
			"string",
			"hello",
			[]rune{'h', 'e', 'l', 'l', 'o'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertStringToRuneSlice(tt.input)

			if len(got) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(got))
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("At index %d: expected %c (%U), got %c (%U)",
						i, tt.expected[i], tt.expected[i], got[i], got[i])
				}
			}
		})
	}
}

func TestHashSliceRunesWithSHA256WithSaltAndPrint(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		expected string
	}{
		{
			"simple input",
			[]rune("test"),
			"6ba9907f83b2a0d28d591f27544e2954d87f67bbec98809a6c484313c501affc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HashSliceRunesWithSHA256WithSaltAndPrint(tt.input)
			gotHex := fmt.Sprintf("%x", got)

			if gotHex != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, gotHex)
			}
		})
	}
}

func TestMainFunctionality(t *testing.T) {
	numDecimal := 100
	numOctal := 0144
	numHexadecimal := 0x64
	e := 2.71828
	str := "one hundred"
	isTrue := false
	numComplex := 1 + 1i

	expectedTypeResult := "type of numDecimal: int"
	gotTypeResult := DeterminesVariableTypeAndDisplaysIt("numDecimal", numDecimal)
	if gotTypeResult != expectedTypeResult {
		t.Errorf("Unexpected type result: %s", gotTypeResult)
	}

	expectedUnionStr := "1001001002.71828one hundredfalse(1+1i)"
	gotUnionStr := ConvertsVariablesToStringAndConcatenatesIntoSingleString(numDecimal, numOctal, numHexadecimal, e, str, isTrue, numComplex)
	if gotUnionStr != expectedUnionStr {
		t.Errorf("Expected union string %q, got %q", expectedUnionStr, gotUnionStr)
	}

	runeSlice := ConvertStringToRuneSlice(expectedUnionStr)
	if len(runeSlice) != len([]rune(expectedUnionStr)) {
		t.Errorf("Expected rune slice length %d, got %d", len([]rune(expectedUnionStr)), len(runeSlice))
	}

	expectedHashString := "90bbb9ec6cac9f8b24c468d88224c70b6fcf30e4f5effc5787fc423f1a07e05d"
	gotHash := HashSliceRunesWithSHA256WithSaltAndPrint(runeSlice)
	if gotHash == [32]byte{} {
		t.Error("Hash should not be zero")
	}
	gotHashHex := fmt.Sprintf("%x", gotHash)
	if expectedHashString != gotHashHex {
		t.Errorf("Hash mismatch. Expected %x, got %x", expectedHashString, gotHashHex)
	}
}
