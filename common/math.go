package common

func Add(a, b int) int {
	return a + b
}

/*func addInt(a, b int) int {
	return a + b
}

func addFloat(a, b float64) float64 {
	return a + b
}*/

type number interface {
	int | float64 | uint
}

/*func add[T int | float64 | uint](a, b T) T {
	return a + b
}*/

func add[T number](a, b T) T {
	return a + b
}

func testAdd() {
	var a int = 10
	var b int = 20
	add(a, b)
}
