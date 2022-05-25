package cache

import (
	"errors"
	"fmt"
	"time"
)

var (
	errorWrongTime  = errors.New("wrong time or format")
	errorKeyExpired = errors.New("key deadline expired")
)

type Cache struct {
	data     map[string]string
	deadline map[string]time.Time
}

func NewCache() Cache {
	data := make(map[string]string)
	deadline := make(map[string]time.Time)
	return Cache{data, deadline}
}

func (c *Cache) Get(key string) (string, bool) {
	if c.deadline[key].IsZero() {
		if c.data[key] != "" {
			return c.data[key], true
		} else {
			return "", false
		}
	}

	if c.deadline[key].After(time.Now()) {
		return c.data[key], true
	} else {
		fmt.Println(fmt.Errorf("%w", errorKeyExpired))
		return "", false
	}
}

func (c *Cache) Put(key, value string) {
	c.data[key] = value
}

func (c *Cache) Keys() []string {
	var keys []string
	for k := range c.data {
		if c.deadline[k].After(time.Now()) {
			keys = append(keys, k)
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	if deadline.After(time.Now()) {
		c.Put(key, value)
		c.deadline[key] = deadline
	} else {
		fmt.Println(fmt.Errorf("%w", errorWrongTime))
	}
}
