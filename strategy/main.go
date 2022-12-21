package main

func main() {
	var deleter IAlgo = &Lfu{}
	cache := initCache(deleter)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	deleter = &Lru{}
	cache.setEvictionAlgo(deleter)

	cache.add("d", "4")

	deleter = &Fifo{}
	cache.setEvictionAlgo(deleter)

	cache.add("e", "5")
}
