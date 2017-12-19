package controller

import (
	"fmt"
)


type Brand struct {

	BrandName string
	BrandId   int
	BrandUrl  string

}

func (B *Brand)Fetch(){
	fmt.Println("this is fetch")
}

func (B *Brand)Handle(){
	fmt.Println("this is handle")
}