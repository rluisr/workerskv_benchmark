package main

import (
	"os"

	"github.com/axiomzen/memcache"
)

type Memcache struct {
	MC *memcache.Client
}

func NewMem() *Memcache {
	mc, err := memcache.New(os.Getenv("MEM_USER"), os.Getenv("MEM_PASS"), os.Getenv("MEM_SERVER"))
	if err != nil {
		panic(err)
	}

	return &Memcache{
		MC: mc,
	}
}

func (m *Memcache) set(key, value string) error {
	return m.MC.Set(&memcache.Item{Key: key, Value: []byte(value)})
}

func (m *Memcache) get(key string) (*memcache.Item, error) {
	return m.MC.Get(key)
}
