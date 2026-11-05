module github.com/andrerrcosta2/gtools/reflect4

go 1.26.5

require (
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/format v0.0.0
	github.com/andrerrcosta2/gtools/core/testlite v0.0.0
)

require (
	github.com/andrerrcosta2/gtools/gflux/core/domain v0.0.0 // indirect
	golang.org/x/text v0.40.0 // indirect
)

require (
	github.com/andrerrcosta2/gtools/core/chronos v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/seeders v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/util v0.0.0
	github.com/google/uuid v1.6.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/core => ../core
	github.com/andrerrcosta2/gtools/core/chronos => ../core/chronos
	github.com/andrerrcosta2/gtools/core/data => ../core/data
	github.com/andrerrcosta2/gtools/core/domain => ../core/domain
	github.com/andrerrcosta2/gtools/core/format => ../core/format
	github.com/andrerrcosta2/gtools/core/seeders => ../core/seeders
	github.com/andrerrcosta2/gtools/core/testlite => ../core/testlite
	github.com/andrerrcosta2/gtools/core/util => ../core/util
	github.com/andrerrcosta2/gtools/gflux/core => ../gflux/core
	github.com/andrerrcosta2/gtools/gflux/core/domain => ../gflux/core/domain
)
