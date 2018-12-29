//go:generate mockgen -destination subdir/internal/pkg/reflect_output/mock.go golang.google.cn/x/mock/mockgen/tests/internal_pkg/subdir/internal/pkg Intf
//go:generate mockgen -source subdir/internal/pkg/input.go -destination subdir/internal/pkg/source_output/mock.go
package test
