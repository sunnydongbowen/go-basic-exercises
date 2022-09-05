package structdemo

// for 结构体工厂
// 私有
type file struct {
	fd   int
	name string
}

type matrix struct {
	name string
}

// 公有
func NewFile(fd int, name string) *file {
	// 一通操作
	return &file{fd, name}
}

func NewMatrix(name string) *matrix {
	// 要记住new初始化实例的方式
	m := new(matrix)
	*m = matrix{name: name}
	//m1:=matrix{name: name}
	return m
}
