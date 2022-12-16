package setx

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	s1 := NewSet()
	t.Log(s1)
	s2 := NewSet([]string{"a", "b", "c"}...)
	t.Log(s2)
}

func TestSetAdd(t *testing.T) {
	s1 := NewSet()
	s1.Add("1", "2", "3")
	t.Log(s1)
}

func TestSetDel(t *testing.T) {
	s2 := NewSet([]string{"a", "b", "c"}...)
	t.Log(s2)

	//删除一个不存在的元素
	s2.Del("d")
	t.Log(s2)

	// 删除一个存在的元素
	s2.Del("a")
	t.Log(s2)
}

func TestSetIntersection(t *testing.T) {
	s1 := NewSet("1", "2", "3")
	s2 := NewSet("2", "3", "4")
	s3 := NewSet("3", "4", "5")
	//[2 3]
	t.Log(s1.Intersection(s2))

	//[3]
	t.Log(s1.Intersection(s3))

	//[3 4]
	t.Log(s2.Intersection(s3))
}

func TestSetUnion(t *testing.T) {
	s1 := NewSet("1", "2", "3")
	s2 := NewSet("2", "3", "4")
	s3 := NewSet("3", "4", "5")
	//[1 2 3 4]
	t.Log(s1.Union(s2))

	//[1 2 3 4 5]
	t.Log(s1.Union(s3))

	//[2 3 4 5]
	t.Log(s2.Union(s3))
}

func TestSetHas(t *testing.T) {
	s1 := NewSet("1", "2", "3")
	t.Log(s1.Has("1"))
	t.Log(s1.Has("4"))
}
func TestSetKeys(t *testing.T) {
	s1 := NewSet("1", "2", "3")
	t.Log(s1.Keys())
}
