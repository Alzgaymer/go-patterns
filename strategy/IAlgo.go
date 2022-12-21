package main

type IAlgo interface {
	evict(c *Cache)
}
