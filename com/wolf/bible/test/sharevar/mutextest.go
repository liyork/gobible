package sharevar

import "sync"

var (
	// 使用chan模拟锁
	sema     = make(chan struct{}, 1) // a binary semaphore guarding balance
	mu       sync.Mutex               // guards balance
	balance2 int                      // 惯例来说，被mutex所保护的变量是在mutex变量声明之后立刻声明的。
	balance3 int
)

func Deposit2(amount int) {
	//sema <- struct{}{} // acquire token
	mu.Lock() // 若已被获取则阻塞直到其他goroutine调用了unlock
	defer mu.Unlock()
	balance2 = balance2 + amount
	//<-sema // release token
}

func Balance2() int {
	//sema <- struct{}{} // acquire token
	mu.Lock()
	defer mu.Unlock() // return语句执行完之后执行
	b := balance2
	//<-sema // release token
	return b
}

// NOTE: not atomic!，包含三个步骤，每个步骤是原子的，但是整体不是
func withdraw(amount int) bool {
	Deposit2(-amount)
	if Balance2() < 0 {
		Deposit2(amount)
		return false // insufficient funds
	}
	return true
}

// 错误,go里没有重入锁，关于重入锁的概念
func withdraw2(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit2(-amount)
	if Balance2() < 0 {
		Deposit2(amount)
		return false // insufficient funds
	}
	return true
}

// 正确方式锁全局
func withdraw3(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance3 < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}

// 对外提供加锁，小写deposit对内不加锁
func Deposit3(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance3() int {
	mu.Lock()
	defer mu.Unlock()
	return balance3
}

// This function requres that the lock be held.
func deposit(amount int) {
	balance3 += amount
}
