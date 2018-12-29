package vendor_dep

//go:generate mockgen -package vendor_dep -destination mock.go golang.google.cn/x/mock/mockgen/tests/vendor_dep VendorsDep
//go:generate mockgen -destination source_mock_package/mock.go -source=vendor_dep.go
