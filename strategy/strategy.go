package strategy

// reference : https://refactoring.guru/design-patterns/strategy/go/example#example-0

type Algo interface {
	evict(c *Cache) string
}

type Cache struct {
	storage     map[string]string
	algo        Algo
	capacity    int
	maxCapacity int
}

func initCache(algo Algo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:     storage,
		algo:        algo,
		capacity:    0,
		maxCapacity: 1,
	}
}

func (c *Cache) setEvictionAlgo(algo Algo) {
	c.algo = algo
}

func (c *Cache) add(key, value string) string {
	algo := ""
	if c.capacity == c.maxCapacity {
		algo = c.evict()
	}
	c.capacity++
	c.storage[key] = value
	return algo
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() string {
	algo := c.algo.evict(c)
	c.capacity--
	return algo
}

type Lru struct {
}

func (l *Lru) evict(c *Cache) string {
	return "lru"
}

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) string {
	return "lfu"
}
