package main

import (
	"sync"
)

// SyncMapWrapper обертка над sync.Map с типизированными методами
type SyncMapWrapper struct {
	m sync.Map
}

// Set устанавливает значение по ключу
func (s *SyncMapWrapper) Set(key string, value interface{}) {
	s.m.Store(key, value)
}

// Get получает значение по ключу
func (s *SyncMapWrapper) Get(key string) (interface{}, bool) {
	return s.m.Load(key)
}

// Delete удаляет значение по ключу
func (s *SyncMapWrapper) Delete(key string) {
	s.m.Delete(key)
}

// Range применяет функцию ко всем парам ключ-значение
func (s *SyncMapWrapper) Range(f func(key string, value interface{}) bool) {
	s.m.Range(func(key, value interface{}) bool {
		return f(key.(string), value)
	})
}

// Len возвращает количество элементов (приблизительное)
func (s *SyncMapWrapper) Len() int {
	count := 0
	s.m.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	return count
}
