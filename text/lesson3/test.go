package lesson3

import (
	"fmt"
	"strconv"
)

func TestIntToString() {
	result := strconv.Itoa(1)
	fmt.Printf("TestIntToString - Result: %s; Type: %T\n", result, result)
}

func TestStringToInt() {
	result, err := strconv.Atoi("123")

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("TestStringToInt - Result: %d; Type: %T\n", result, result)
}

func TestStringToFloat() {
	f, err := strconv.ParseFloat("123.12", 32)

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("TestStringToFloat - Result: %f; Type: %T\n", f, f)
}

func TestFloatToString() {
	s := strconv.FormatFloat(12.36, 'g', -1, 64)

	fmt.Printf("TestFloatToString - Result: %s; Type %T\n", s, s)
}

func TestStringToBool() {
	result, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	fmt.Printf("TestStringToBool - Result: %t; Type: %T\n", result, result)
}

func TestBoolToString() {
	b := strconv.FormatBool(true)
	fmt.Printf("TestBoolToString - Result: %s; Type: %T\n", b, b)
}

func TestQuote() {
	s := strconv.Quote("go welcome to the secret shop")
	fmt.Println(s)
}

func TestQuoteASCII() {
	s := strconv.QuoteToASCII("welocme to пикер")
	fmt.Println(s)
}

func TestUnquote() {
	s, err := strconv.Unquote(`"go \u043b\u0443\u0447\u0448\u0435 \u0432\u0441\u0435\u0445"`)
	fmt.Println(s, err)
}

func Test() {
	TestIntToString()
	TestStringToInt()
	TestStringToFloat()
	TestFloatToString()
	TestStringToBool()
	TestBoolToString()
	TestQuote()
	TestQuoteASCII()
	TestUnquote()
}
