package main

type Math interface {
	round(int)
}

func round(number int) int {
	return number - 1
}

func Ten() {
	len()
}
