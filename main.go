package main

import (
	c "github.com/jpyeah/goworker/controller"
)

func main () {

	b :=&c.Brand{}

	b.Fetch()

}