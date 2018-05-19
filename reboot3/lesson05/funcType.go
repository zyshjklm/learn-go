package main

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	funcMap := map[string]func(int, int) int{
		"+": add,
		"-": sub,
	}

}
