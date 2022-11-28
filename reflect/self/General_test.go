package self

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// @program:     go-basic-exercises
// @file:        General_test.go.go
// @author:      bowen
// @create:      2022-11-12 20:17
// @description:

func Test_reflectapp(t *testing.T) {
	tests := []struct {
		name    string
		entity  any
		wantErr error
	}{
		{
			name:    "nil",
			entity:  nil,
			wantErr: errors.New("传参不能为空"),
		},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := reflectapp(tt.entity)
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr != nil {
				//t.Errorf("reflectapp() error = %v, wantErr %v", err, tt.wantErr)
				// 预期会有错误返回，就不需要进一步校验其它两个返回值了
				return
			}
		})
	}
}
