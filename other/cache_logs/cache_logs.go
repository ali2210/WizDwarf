package cache_logs

import (
	"errors"
	"reflect"
	"time"

	"github.com/allegro/bigcache"
)

type Bigcached struct {
	client *bigcache.BigCache
}

func (b *Bigcached) Set_Key(key, value string) error {

	if reflect.DeepEqual(b.client, Bigcached{}) {
		panic(errors.New("Bigcached object is not initialized"))
	}
	b.client.Set(key, []byte(value))
	return nil
}

func (b *Bigcached) Get_Key(key string) (string, error) {

	value, err := b.client.Get(key)
	if err != nil {
		panic(err.Error())
	}
	return string(value), nil
}

func New(config bigcache.Config) *Bigcached {

	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		panic(err.Error())
	}

	return &Bigcached{client: cache}
}

func GetBigcached_config() bigcache.Config {

	return bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Second,
		MaxEntriesInWindow: 1024,
		MaxEntrySize:       2 * 1024,
		Verbose:            true,
		HardMaxCacheSize:   8192,
		CleanWindow:        10 * time.Second,
	}
}
