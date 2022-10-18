package singleton

import (
	"reflect"
	"sync"

	"github.com/MouseChannel/MoCheng_Go_Server/face"
)

var cache sync.Map

// Singleton returns a singleton of T.
func Singleton[T any]() (t *T) {
	hash := reflect.TypeOf(t)
	v, ok := cache.Load(hash)

	if ok {
		return v.(*T)
	}

	v = new(T)

	v, _ = cache.LoadOrStore(hash, v)

	v.(face.ISingletonItem).Init()

	return v.(*T)
}
