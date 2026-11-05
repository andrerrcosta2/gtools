module github.com/andrerrcosta2/gtools/ast4

go 1.26.5

require github.com/andrerrcosta2/gtools/core/format v0.0.0

require (
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/util v0.0.0 // indirect
	golang.org/x/text v0.40.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/core/domain => ../core/domain
	github.com/andrerrcosta2/gtools/core/format => ../core/format
	github.com/andrerrcosta2/gtools/core/testlite => ../core/testlite
	github.com/andrerrcosta2/gtools/core/util => ../core/util
	github.com/andrerrcosta2/gtools/gflux/core => ../gflux/core
)
