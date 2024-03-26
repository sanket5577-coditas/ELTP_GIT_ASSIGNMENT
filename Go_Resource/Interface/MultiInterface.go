package main

import (
	"fmt"
)

func (e email) cost() int {
	if e.isSubscribed {
		return 0.05 * float(len(body))
	}

	return 0.01 * float(len(body))
}

func (e email) format() string {
	fmt.Println(body)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
