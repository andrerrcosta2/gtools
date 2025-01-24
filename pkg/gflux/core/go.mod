module github.com/andrerrcosta2/gtools/gflux/core

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/gtools v0.0.0
	github.com/andrerrcosta2/gtools/core/data v0.0.0
)

replace (
	github.com/andrerrcosta2/gtools/core/gtools => ./../../core/gtools
	github.com/andrerrcosta2/gtools/core/data => ./../../core/data
)