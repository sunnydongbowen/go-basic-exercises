package geek_func1

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIterateFuncs(t *testing.T) {
	type args struct {
		val any
	}
	tests := []struct {
		name string // 测试用例名
		args args   // 参数

		want    map[string]*FuncInfo
		wantErr error
	}{
		{
			name:    "nil",
			wantErr: errors.New("输入nil"),
		},
		{
			name: "basic type",
			args: args{
				val: 123,
			},
			wantErr: errors.New("不支持该类型"),
		},
		{
			name: "struct type",
			args: args{
				val: Order{
					buyer:  18,
					seller: 100,
				},
			},
			want: map[string]*FuncInfo{
				"GetBuyer": {
					Name:   "GetBuyer",
					In:     []reflect.Type{reflect.TypeOf(Order{})},
					Out:    []reflect.Type{reflect.TypeOf(int64(0))},
					Result: []any{int64(18)},
				},
				"getSeller": {
					Name:   "getSeller",
					In:     []reflect.Type{reflect.TypeOf(Order{})},
					Out:    []reflect.Type{reflect.TypeOf(int64(0))},
					Result: []any{int64(100)},
				},
			},
		},
	}
	// 循环这个test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateFuncs(tt.args.val)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

type Order struct {
	buyer  int64
	seller int64
}

type OrderPtr struct {
	buyer  int64
	seller int64
}

func (o Order) GetBuyer() int64 {
	return o.buyer
}

func (o Order) getSeller() int64 {
	return o.seller

}
