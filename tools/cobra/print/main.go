package main

func main() {

}

func OutSide(a int) func(y int) int {
	return func(y int) int {
		return a * y
	}
}
