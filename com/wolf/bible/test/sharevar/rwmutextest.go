package sharevar

import "sync"

//RWMutex只有当获得锁的大部分goroutine都是读操作，RWMutex才是最能带来好处的。

var mu4 sync.RWMutex // 多读单写锁
var balance4 int

// 读
func Balance4() int {
	mu4.RLock() // readers lock
	defer mu4.RUnlock()
	return balance
}

// 互斥
func Deposit4(amount int) {
	mu4.Lock()
	defer mu4.Unlock()
	deposit(amount)
}
