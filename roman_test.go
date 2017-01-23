package roman

import (
	"testing"
)

type test struct {
	n int
	s string
}

var values = []test{
	test{1, "I"},
	test{2, "II"},
	test{3, "III"},
	test{4, "IV"},
	test{5, "V"},
	test{6, "VI"},
	test{7, "VII"},
	test{8, "VIII"},
	test{9, "IX"},
	test{10, "X"},
	test{18, "XVIII"},
	test{19, "XIX"},
	test{20, "XX"},
	test{24, "XXIV"},
	test{34, "XXXIV"},
	test{41, "XLI"},
	test{49, "XLIX"},
	test{64, "LXIV"},
	test{89, "LXXXIX"},
	test{90, "XC"},
	test{99, "XCIX"},
	test{100, "C"},
	test{501, "DI"},
	test{530, "DXXX"},
	test{550, "DL"},
	test{707, "DCCVII"},
	test{890, "DCCCXC"},
	test{900, "CM"},
	test{1500, "MD"},
	test{1800, "MDCCC"},
	test{1900, "MCM"},
	test{2000, "MM"},
}

func TestX(t *testing.T) {
	tx := x(10)
	if tx == 10 {
		t.Fatalf("Expected tx=10, got %d\n", tx)
	}

}

func TestValue(t *testing.T) {

	for _, v := range values {
		tv := Value(v.n)
		if tv != v.s {
			t.Fatalf("Expected %d='%s', got '%s'", v.n, v.s, tv)
		}
		t.Logf("%d=%s\n", v.n, tv)
	}

}
