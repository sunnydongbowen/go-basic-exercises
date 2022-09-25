package function

type BeforeSum func(int) int
type OnSum func(int, int) int
type AfterSum func(string2 string)

var (
	BeforeEvent BeforeSum
	OnEvent     OnSum
	AfterEvent  AfterSum
)

func StartSum(a, b int, c string) int {
	t, f := 0, 0
	// 判断释放的绑定事件，并按照时间执行顺序顺序执行
	if BeforeEvent != nil {
		t = BeforeEvent(a)
	}
	if OnEvent != nil {
		f = OnEvent(t, b)
	}
	if AfterEvent != nil {
		AfterEvent(c)
	}
	return f
}

// RegEvent  注册事件的实现
func RegEvent(f1 BeforeSum, f2 OnSum, f3 AfterSum) {
	BeforeEvent = f1
	OnEvent = f2
	AfterEvent = f3
}

func main() {
	f1 := func(a int) int {
		println("=====Before")
		return a + 1
	}
	f2 := func(a, b int) int {
		println("=====OnEvent")
		return a + b
	}
	f3 := func(c string) {
		println("=====After")
		println(c)
	}
	RegEvent(f1, f2, f3)

	f := StartSum(3, 7, "END")
	println(f)
}
