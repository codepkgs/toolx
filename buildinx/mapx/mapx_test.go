package mapx

import (
	"testing"
)

func TestMapxHas(t *testing.T) {
	m := map[string]string{
		"username": "zhangsan",
		"password": "zhangsan",
	}
	t.Log(Has(m, "username"))
}

func TestMapxKeys(t *testing.T) {
	m := map[string]string{
		"username": "zhangsan",
		"password": "zhangsan",
	}
	t.Log(Keys(m))
}

func TestMapxValues(t *testing.T) {
	m := map[string]string{
		"username": "zhangsan",
		"password": "zhangsan123",
	}
	t.Log(Values(m))
}
