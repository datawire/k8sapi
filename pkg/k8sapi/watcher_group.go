package k8sapi

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
)

type WatcherGroup[T runtime.Object] map[string]*Watcher[T]

func (w WatcherGroup[T]) AddWatcher(watcher *Watcher[T]) error {
	_, ok := w[watcher.namespace]
	if ok {
		return fmt.Errorf("Watcher for namespace %s already exists", watcher.namespace)
	}
	w[watcher.namespace] = watcher
	return nil
}

func NewWatcherGroup[T runtime.Object]() WatcherGroup[T] {
	return make(WatcherGroup[T])
}

func (w WatcherGroup[T]) EnsureStarted(ctx context.Context, cb func(bool)) {
	for _, v := range w {
		v.EnsureStarted(ctx, cb)
	}
}

func (w WatcherGroup[T]) Cancel() {
	for _, v := range w {
		v.Cancel()
	}
}

func (w WatcherGroup[T]) Get(ctx context.Context, namespace string, obj T) (T, bool, error) {
	watcher, ok := w[namespace]
	if !ok {
		var zeroValue T
		return zeroValue, false, fmt.Errorf("Watcher for namespace %s does not exist", namespace)
	}
	return watcher.Get(ctx, obj)
}

func (w WatcherGroup[T]) List(ctx context.Context, namespace string, obj T) []T {
	a := make([]T, 0)
	for _, v := range w {
		a = append(a, v.List(ctx)...)
	}
	return a
}
