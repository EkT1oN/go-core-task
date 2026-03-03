package main

import "maps"

type StringIntMap struct {
	mapValue map[string]int
}

func (s *StringIntMap) Add(key string, value int) {
	s.mapValue[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.mapValue, key)
}

func (s *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	maps.Copy(newMap, s.mapValue)

	return newMap
}

func (s *StringIntMap) Exists(key string) bool {
	_, ok := s.mapValue[key]
	return ok
}

func (s *StringIntMap) Get(key string) (int, bool) {
	value, ok := s.mapValue[key]
	return value, ok
}

func main() {

}
