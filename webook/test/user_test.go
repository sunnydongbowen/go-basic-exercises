package test

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      user_test.go
// @author:    bowen
// @time:      2023-09-22 14:58
// @description:

func TestEncrypt(t *testing.T) {
	passwd := "123499"
	encrypted, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword(encrypted, []byte(passwd))
	require.NoError(t, err)
}
