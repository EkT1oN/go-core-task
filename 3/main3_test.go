package main

import (
	"maps"
	"testing"
)

func TestAdd(t *testing.T) {
	obj := StringIntMap{
		mapValue: make(map[string]int),
	}
	obj.Add("abc", 123)

	expected := map[string]int{
		"abc": 123,
	}
	if !maps.Equal(obj.mapValue, expected) {
		t.Errorf("Expected map %+v, got %+v", expected, obj.mapValue)
	}
}

func TestRemove(t *testing.T) {
	obj := StringIntMap{
		mapValue: map[string]int{
			"a": 0,
			"b": 1,
			"c": 2,
		},
	}
	obj.Remove("a")

	expected := map[string]int{
		"b": 1,
		"c": 2,
	}
	if !maps.Equal(obj.mapValue, expected) {
		t.Errorf("Expected map %+v, got %+v", expected, obj.mapValue)
	}
}

func TestCopy(t *testing.T) {
	obj := StringIntMap{
		mapValue: map[string]int{
			"a": 0,
			"b": 1,
			"c": 2,
		},
	}
	got := obj.Copy()

	if !maps.Equal(got, obj.mapValue) {
		t.Errorf("Expected map %+v, got %+v", obj.mapValue, got)
	}
}

func TestExists(t *testing.T) {
	obj := StringIntMap{
		mapValue: map[string]int{
			"a": 0,
			"b": 1,
			"c": 2,
		},
	}

	ok := obj.Exists("a")
	if !ok {
		t.Errorf("Expected key exists %+v, got %+v", true, ok)
	}

	ok = obj.Exists("abc")
	if ok {
		t.Errorf("Expected key not exists %+v, got %+v", false, ok)
	}
}

func TestGet(t *testing.T) {
	obj := StringIntMap{
		mapValue: map[string]int{
			"a": 0,
			"b": 1,
			"c": 2,
		},
	}
	gotValue, ok := obj.Get("a")
	if !ok {
		t.Errorf("Expected key exists %+v, got %+v", true, ok)
	}
	if gotValue != 0 {
		t.Errorf("Expected value %+v, got %+v", 0, gotValue)
	}

	gotValue, ok = obj.Get("abc")
	if ok {
		t.Errorf("Expected key not exists %+v, got %+v", false, ok)
	}
}
