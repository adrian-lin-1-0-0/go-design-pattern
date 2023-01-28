package strategy

import "testing"

func TestCache(t *testing.T) {
	lfu := &Lfu{}
	cache := initCache(lfu)
	got := cache.add("a", "1")
	want := ""
	if got != want {
		t.Errorf("Computer.Print() got = %v, want %v", got, want)
	}
	got = cache.add("b", "2")
	want = "lfu"
	if got != want {
		t.Errorf("Computer.Print() got = %v, want %v", got, want)
	}
	lru := &Lru{}
	cache.setEvictionAlgo(lru)
	got = cache.add("d", "4")
	want = "lru"
	if got != want {
		t.Errorf("Computer.Print() got = %v, want %v", got, want)
	}

}
