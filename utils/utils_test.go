package utils

import (
	"fmt"
	"testing"
)

func TestGetRandID(t *testing.T) {
	fmt.Println("--------------------")
	fmt.Println(GetRandID(2, 5))
	fmt.Println(GetRandID(7, 10))
	fmt.Println(GetRandID(7, 10))
	fmt.Println(GetRandID(7, 10))
}

func TestGetRandWord(t *testing.T) {
	fmt.Println("--------------------")
	fmt.Println(GetRandWord(2, 5))
	fmt.Println(GetRandWord(2, 5))
	fmt.Println(GetRandWord(2, 5))
	fmt.Println(GetRandWord(1, 2))
	fmt.Println(GetRandWord(1, 2))
	fmt.Println(GetRandWord(1, 2))
	fmt.Println(GetRandWord(1, 2))
	fmt.Println(GetRandWord(1, 1))
	fmt.Println(GetRandWord(2, 2))
	fmt.Println(GetRandWord(0, 0))
}

func TestGetRandPhrase(t *testing.T) {
	fmt.Println("--------------------")
	fmt.Println(GetRandPhrase(2, 5, 3, 10))
	fmt.Println(GetRandPhrase(3, 7, 1, 1))
	fmt.Println(GetRandPhrase(3, 7, 1, 2))
	fmt.Println(GetRandPhrase(3, 7, 1, 3))
	fmt.Println(GetRandPhrase(3, 7, 3, 4))
	fmt.Println(GetRandPhrase(3, 7, 4, 4))
}

func TestGetImage(t *testing.T) {
	fmt.Println(GetImage("http://www.loja/produto/asd2432/img/%s"))
}

func TestGetValue(t *testing.T) {
	valueProbPairs := []ValueProbPair{
		{"a", 0.4},
		{"b", 0.4},
		{"c", 0.1},
	}
	fmt.Println(GetValue(valueProbPairs))
	fmt.Println(GetValue(valueProbPairs))
	fmt.Println(GetValue(valueProbPairs))
	fmt.Println(GetValue(valueProbPairs))
	fmt.Println(GetValue(valueProbPairs))
	fmt.Println(GetValue(valueProbPairs))
	fmt.Println(GetValue(valueProbPairs))
	fmt.Println(GetValue(valueProbPairs))
	fmt.Println(GetValue(valueProbPairs))
	fmt.Println(GetValue(valueProbPairs))
	fmt.Println(GetValue(valueProbPairs))
	fmt.Println(GetValue(valueProbPairs))
}

func TestGetRandPrice(t *testing.T) {
	fmt.Println(GetRandPrice(1.0, 50.0, "BRL"))
	fmt.Println(GetRandPrice(1.0, 50.0, "BRL"))
	fmt.Println(GetRandPrice(1.0, 50.0, "BRL"))
	fmt.Println(GetRandPrice(1.0, 50.0, "BRL"))
	fmt.Println(GetRandPrice(1.0, 50.0, "BRL"))
	fmt.Println(GetRandPrice(1.0, 50.0, "BRL"))
	fmt.Println(GetRandPrice(1.0, 50.0, "BRL"))
	fmt.Println(GetRandPrice(1.0, 50.0, "BRL"))
	fmt.Println(GetRandPrice(1.0, 50.0, "BRL"))
	fmt.Println(GetRandPrice(1.0, 50.0, "BRL"))
}

func TestGetRandGtin(t *testing.T) {
	fmt.Println(GetRandGtin())
	fmt.Println(GetRandGtin())
	fmt.Println(GetRandGtin())
	fmt.Println(GetRandGtin())
	fmt.Println(GetRandGtin())
	fmt.Println(GetRandGtin())
	fmt.Println(GetRandGtin())
	fmt.Println(GetRandGtin())
	fmt.Println(GetRandGtin())
	fmt.Println(GetRandGtin())
}
