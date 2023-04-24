package main

func main() {

	// Basics
	const a = 42
	var i int = a // 42
	const b float32 = 3
	var f32 float32 = b
	var f64 float64 = float64(b)

	println(i, f32, f64)

	// Expressions
	const c = iota
	println(c)

	const (
		d = 2 * 5
		e        // 2 * 5
		f = iota // 2 related to position
		g
		h = 10 + iota
	)

	println(d, e, f, g, h)

	// Effective
	type ByteSize float64

	const (
		_           = iota
		KB ByteSize = 1 << (10 * iota)
	)
	println(KB)
}
