package main

func Toyek(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}

	return a * b
}
