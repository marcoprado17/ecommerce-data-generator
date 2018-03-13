package utils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

const imgsDir = "/home/c041/go/src/github.com/marcoprado17/ecommerce-data-generator/imgs"

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const digits = "0123456789"

const digitsWithoutZero = "123456789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// ValueProbPair represents a value and its probability to be sorted
type ValueProbPair struct {
	V interface{}
	P float64
}

func check(err error) {
	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
		os.Exit(1)
	}
}

// GetRandID returns an id with length between minLength and maxLength
func GetRandID(minLength, maxLength int) string {
	length := seededRand.Intn(maxLength-minLength+1) + minLength
	b := make([]byte, length)
	idCharset := charset + digits
	for i := range b {
		b[i] = idCharset[seededRand.Intn(len(idCharset))]
	}
	return string(b)
}

// GetRandWord return a word with length between minLength and maxLength
func GetRandWord(minLength, maxLength int) string {
	length := seededRand.Intn(maxLength-minLength+1) + minLength
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GetRandGtin returns a random gtin
func GetRandGtin() string {
	b := make([]byte, 13)
	for i := range b {
		if i == 0 {
			b[i] = digitsWithoutZero[seededRand.Intn(len(digitsWithoutZero))]
		} else {
			b[i] = digits[seededRand.Intn(len(digits))]
		}
	}
	return string(b)
}

// GetRandPhrase returns a random phrase
func GetRandPhrase(minWordLength, maxWordLength, nMinWords, nMaxWords int) string {
	nWords := seededRand.Intn(nMaxWords-nMinWords+1) + nMinWords
	s := ""
	for i := 1; i <= nWords; i++ {
		s += GetRandWord(minWordLength, maxWordLength)
		if i < nWords {
			s += " "
		}
	}
	return s
}

// GetImage create a copy of an image in original_images folder, assign a random name and paste it in imgs folder
func GetImage(imageLink string) string {
	files, err := ioutil.ReadDir(imgsDir)
	if err != nil {
		fmt.Println(err)
	}

	fileName := files[seededRand.Intn(len(files))].Name()

	return fmt.Sprintf(imageLink, fileName)
}

// GetValue returns a value from a slice of ValueProbPairs considering the probability of each value
func GetValue(valueProbPairs []ValueProbPair) interface{} {
	valueSumProbPair := make([]ValueProbPair, len(valueProbPairs))
	copy(valueProbPairs, valueProbPairs)
	for idx, valueProbPair := range valueProbPairs {
		sumProb := 0.0
		for i := 0; i <= idx; i++ {
			sumProb += valueProbPairs[i].P
		}
		valueSumProbPair[idx] = ValueProbPair{valueProbPair.V, sumProb}
	}
	totalProb := 0.0
	for _, valueProbPair := range valueProbPairs {
		totalProb += valueProbPair.P
	}
	sortedValue := seededRand.Float64() * totalProb
	for _, valueSumProbPair := range valueSumProbPair {
		if sortedValue <= valueSumProbPair.P {
			return valueSumProbPair.V
		}
	}
	return valueProbPairs[len(valueProbPairs)-1].V
}

// GetRandPrice return a random price according google format
func GetRandPrice(minValue, maxValue float64, currency string) string {
	value := seededRand.Float64()*(maxValue-minValue) + minValue
	return fmt.Sprintf("%.2f %s", value, currency)
}
