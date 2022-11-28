package myio

import "os"

// @program:     go-basic-exercises
// @file:        myio.go
// @author:      bowen
// @create:      2022-11-28 13:39
// @description:

type MyIO struct {
	data  string
	bytes []byte
	file  *os.File
}

func (myio MyIO) Write(p []byte) (n int, err error) {
	//TODO implement me
	//panic("implement me")
	if len(p) == 0 {
		return 0, nil
	}
	myio.bytes = p
	return myio.file.Write(p)
}

// io reader接口
func (myio MyIO) Read(p []byte) (n int, err error) {
	//TODO implement me
	//panic("implement me")
	if len(p) == 0 {
		return 0, nil
	}
	myio.data = string(p)
	return myio.file.Read(p)
}

func Create(name string) (*MyIO, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	io1 := &MyIO{file: f}
	return io1, nil
}

func (myio *MyIO) Close() error {
	if myio == nil {
		return nil
	}
	return myio.file.Close()
}

func Open(name string) (*MyIO, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	io1 := &MyIO{file: f}
	return io1, nil
}
