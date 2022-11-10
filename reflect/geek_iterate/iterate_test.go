package geek_iterate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterate(t *testing.T) {
	tests := []struct {
		name  string
		input any

		wantRes []any
		wantErr error
	}{
		{
			name:    "slice",
			input:   []int{1, 2, 3},
			wantRes: []any{1, 2, 3},
		},
		{
			name:    "array",
			input:   [5]int{1, 2, 3, 4, 5},
			wantRes: []any{1, 2, 3, 4, 5},
		},
		//	{ 不要写这种操作，没意义！
		//		name:    "string",
		//		input:   "hello",
		//		wantRes: []any{"hello"},
		//	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Iterate(tt.input)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantRes, res)

		})
	}
}

//func TestIterateMap(t *testing.T) {
//
//	tests := []struct {
//		name       string
//		input      any
//		wantKeys   []any
//		wantValues []any
//		wantErr    error
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			keys, values, err := IterateMapV2(tt.input)
//
//			assert.Equalf(t, tt.want, got, "IterateMapV2(%v)", tt.args.input)
//			assert.Equalf(t, tt.want1, got1, "IterateMapV2(%v)", tt.args.input)
//		})
//	}
//}
