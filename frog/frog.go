package frog

// Jump : calculate how many times frog will jump
func Jump(X int, Y int, D int) int {

	l := Y - X

	times := l / D

	if l%D > 0 {
		times++
	}

	return times
}

func main() {

}
