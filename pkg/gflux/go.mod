module github.com/andrerrcosta2/gtools/gflux

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/gflux/obs v0.0.0
    github.com/andrerrcosta2/gtools/gflux/retry v0.0.0
    github.com/andrerrcosta2/gtools/gflux/core v0.0.0
)

replace (
	github.com/andrerrcosta2/gtools/gflux/obs => ./obs
	github.com/andrerrcosta2/gtools/gflux/retry => ./retry
	github.com/andrerrcosta2/gtools/gflux/core => ./core
)