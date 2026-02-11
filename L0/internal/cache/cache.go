package cache

import (
	"sync"
	"time"

	"microservice/internal/models"
	"microservice/internal/validation"
)

type CacheItem struct {
	Order     *models.Order
	CreatedAt time.Time
}

type Cache struct {
	data  map[string]*CacheItem
	mutex sync.RWMutex
	ttl   time.Duration
}

// Проверка соответствия интерфейсу
var _ CacheIface = (*Cache)(nil)

// Создание нового кэша с TTL
func NewCache(ttl time.Duration) *Cache {
	c := &Cache{
		data: make(map[string]*CacheItem),
		ttl:  ttl,
	}
	go c.cleanupRoutine()
	return c
}

// Потокобезопасное получение заказа
func (c *Cache) Get(key string) (*models.Order, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, ok := c.data[key]
	if !ok || time.Since(item.CreatedAt) > c.ttl {
		return nil, false
	}
	return item.Order, true
}

// Потокобезопасная запись заказа без проверки
func (c *Cache) Set(key string, order *models.Order) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = &CacheItem{
		Order:     order,
		CreatedAt: time.Now(),
	}
}

// Потокобезопасная запись заказа с валидацией
func (c *Cache) SetValidated(key string, order *models.Order) error {
	if err := validation.ValidateOrder(order); err != nil {
		return err
	}
	c.Set(key, order)
	return nil
}

// Периодическая очистка устаревших элементов
func (c *Cache) cleanupRoutine() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		c.cleanup()
	}
}

// Удаление устаревших элементов
func (c *Cache) cleanup() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	now := time.Now()
	for k, v := range c.data {
		if now.Sub(v.CreatedAt) > c.ttl {
			delete(c.data, k)
		}
	}
}
