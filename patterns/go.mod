module github.com/andrerrcosta2/gtools/patterns

go 1.26.5

require (
	github.com/andrerrcosta2/gtools/conc/core v0.0.0
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/io v0.0.0
	github.com/andrerrcosta2/gtools/core/testlite v0.0.2-alpha.3
)

require (
	github.com/andrerrcosta2/gtools/core/format v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/util v0.0.0 // indirect
	golang.org/x/text v0.40.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/conc/core => ../conc/core
	github.com/andrerrcosta2/gtools/core/chronos => ../core/chronos
	github.com/andrerrcosta2/gtools/core/data => ../core/data
	github.com/andrerrcosta2/gtools/core/domain => ../core/domain
	github.com/andrerrcosta2/gtools/core/format => ../core/format
	github.com/andrerrcosta2/gtools/core/io => ../core/io
	github.com/andrerrcosta2/gtools/core/seeders => ../core/seeders
	github.com/andrerrcosta2/gtools/core/testlite => ../core/testlite
	github.com/andrerrcosta2/gtools/core/util => ../core/util
	github.com/andrerrcosta2/gtools/gflux/core => ../gflux/core
	github.com/andrerrcosta2/gtools/gflux/core/domain => ../gflux/core/domain
)
