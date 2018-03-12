package product

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestNewSimpleValid(t *testing.T) {
	fmt.Println("--------------------")
	product := NewSimpleValid(
		"http://www.oi.com.br/produto/%s",
		"http://www.oi.com.br/produto/img/%s",
	)
	spew.Dump(*product)
}

func TestNewInvalidByAusentTitle(t *testing.T) {
	fmt.Println("--------------------")
	product := NewInvalidByAusentTitle(
		"http://www.oi.com.br/produto/%s",
		"http://www.oi.com.br/produto/img/%s",
	)
	spew.Dump(*product)
}

func TestNewInvalidByAusentImage(t *testing.T) {
	fmt.Println("--------------------")
	product := NewInvalidByAusentImage(
		"http://www.oi.com.br/produto/%s",
		"http://www.oi.com.br/produto/img/%s",
	)
	spew.Dump(*product)
}
