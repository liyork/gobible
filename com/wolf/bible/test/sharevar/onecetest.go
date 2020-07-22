package sharevar

import "sync"

var rw sync.RWMutex
var iconsRw map[string]string

// Concurrency-safe.读锁先获取，没有再用写锁加载
func IconRW(name string) string {
	rw.RLock()
	if iconsRw != nil {
		icon := iconsRw[name]
		rw.RUnlock()
		return icon
	}
	rw.RUnlock()

	// acquire an exclusive lock
	mu.Lock()
	if iconsRw == nil { // must recheck for nil
		loadIcons()
	}
	icon := icons[name]
	mu.Unlock()
	return icon
}

func loadIcons() {

}

var loadIconsOnce sync.Once

// Concurrency-safe.一次性初始化
func IconOnce(name string) string {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
