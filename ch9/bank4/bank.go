// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 263.

// Package bank provides a concurrency-safe single-account bank.
package bank

//!+
import (
	"log"
	"sync"
)

var (
	mu      sync.RWMutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

//!-
