module github.com/andrerrcosta2/gtools/gflux

go 1.26.5

require (
	github.com/andrerrcosta2/gtools/gflux/core/domain v0.0.0
	github.com/andrerrcosta2/gtools/gflux/obs v0.0.2-alpha.3
)

require (
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/gflux/core v0.0.2-alpha.3 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/core/data => ../core/data
	github.com/andrerrcosta2/gtools/core/domain => ../core/domain
	github.com/andrerrcosta2/gtools/core/format => ../core/format
	github.com/andrerrcosta2/gtools/core/io => ../core/io
	github.com/andrerrcosta2/gtools/core/testlite => ../core/testlite
	github.com/andrerrcosta2/gtools/core/util => ../core/util
	github.com/andrerrcosta2/gtools/gflux/core => ./core
	github.com/andrerrcosta2/gtools/gflux/core/domain => ./core/domain
	github.com/andrerrcosta2/gtools/gflux/obs => ./obs
)
