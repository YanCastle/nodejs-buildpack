//go:generate mockgen -destination mock/concurrent_mock.go golang.google.cn/x/mock/sample/concurrent Math

// Package concurrent demonstrates how to use gomock with goroutines.
package concurrent

type Math interface {
	Sum(a, b int) int
}
