package roman

import (
	"math"
)

func Value(v int) string {
	var s string

	mf := m(v)
	if mf > 0 {
		s = s + factor(mf, "M", "")
		v = v - mf*1000
	}

	mc := c(v)
	if mc > 0 {
		s = s + factormiddle(mc, "C", "D", "M")
		v = v - mc*100
	}

	mx := x(v)
	if mx > 0 {
		s = s + factormiddle(mx, "X", "L", "C")
		v = v - mx*10
	}

	if v < 10 {
		s = s + unit(v)
	}

	return s
}

// How many times 1000 ?
func m(v int) int {
	return v / 1000
}

// how many times 100 ?
func c(v int) int {
	return v / 100
}

// how many times 10 ?
func x(v int) int {
	return int(math.Floor(float64(v / 10)))
}

func factormiddle(v int, s string, mid string, sup string) string {
	var r string

	switch v {
	case 1, 2, 3:
		for i := 0; i < v; i++ {
			r = r + s
		}
		return r
	case 4:
		return s + mid
	case 5:
		return mid
	case 6, 7, 8:
		r = mid
		for i := 0; i < v-5; i++ {
			r = r + s
		}
	case 9:
		return s + sup
	}

	return r
}

func factor(v int, s string, sup string) string {
	var r string

	switch v {
	case 1, 2, 3:
		for i := 0; i < v; i++ {
			r = r + s
		}
		return r
	case 4:
		return s + sup
	case 5:
		return sup
	case 6, 7, 8:
		r = sup
		for i := 0; i < v-5; i++ {
			r = r + s
		}
	case 9:
		return s + sup
	}

	return r
}

// unit
func unit(v int) string {
	switch v {
	default:
		return ""
	case 1, 2, 3:
		var s string
		for i := 0; i < v; i++ {
			s = s + "I"
		}
		return s
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6, 7, 8:
		s := "V"
		for i := 0; i < v-5; i++ {
			s = s + "I"
		}
		return s
	case 9:
		return "IX"
	case 10:
		return "X"
	}

}
