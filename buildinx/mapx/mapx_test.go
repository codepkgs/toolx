package mapx

import (
	"testing"
)

func TestMapxIn(t *testing.T) {
	m := map[string]string{
		"username": "zhangsan",
		"password": "zhangsan",
	}
	t.Log(In("username", m))
}
