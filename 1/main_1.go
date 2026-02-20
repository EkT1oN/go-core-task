package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func DeterminesVariableTypeAndDisplaysIt(name string, value any) string {
	result := fmt.Sprintf("type of %s: %T", name, value)
	fmt.Println(result)
	return result
}

func ConvertsVariablesToStringAndConcatenatesIntoSingleString(values ...any) string {
	var result strings.Builder
	for _, value := range values {
		fmt.Fprintf(&result, "%v", value)
	}
	return result.String()
}

func ConvertStringToRuneSlice(input string) []rune {
	return []rune(input)
}

func HashSliceRunesWithSHA256WithSaltAndPrint(runeSlice []rune) [32]byte {
	stringFromRune := string(runeSlice)
	stringCenter := len(stringFromRune) / 2
	saltedString := stringFromRune[:stringCenter] + "go-2024" + stringFromRune[stringCenter:]
	hashString := sha256.Sum256([]byte(saltedString))
	fmt.Println("salted string:", saltedString)
	fmt.Printf("SHA256 result: %x\n", hashString)
	return hashString
}

func main() {
	numDecimal := 100
	numOctal := 0144
	numHexadecimal := 0x64
	e := 2.71828
	str := "one hundred"
	isTrue := false
	numComplex := 1 + 1i

	DeterminesVariableTypeAndDisplaysIt("numDecimal", numDecimal)
	DeterminesVariableTypeAndDisplaysIt("numOctal", numOctal)
	DeterminesVariableTypeAndDisplaysIt("numHexadecimal", numHexadecimal)
	DeterminesVariableTypeAndDisplaysIt("e", e)
	DeterminesVariableTypeAndDisplaysIt("str", str)
	DeterminesVariableTypeAndDisplaysIt("isTrue", isTrue)
	DeterminesVariableTypeAndDisplaysIt("numComplex", numComplex)

	unionStr := ConvertsVariablesToStringAndConcatenatesIntoSingleString(numDecimal, numOctal, numHexadecimal, e, str, isTrue, numComplex)

	runeSlice := ConvertStringToRuneSlice(unionStr)

	HashSliceRunesWithSHA256WithSaltAndPrint(runeSlice)

	HashSliceRunesWithSHA256WithSaltAndPrint([]rune("test"))
}
