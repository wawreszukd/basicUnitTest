package main

import "errors"

func main() {

}
func add(a int, b int) (int, error) {
	if a == 0 {
		return 0, errors.New("Chuj")
	}
	return a + b, nil

}
