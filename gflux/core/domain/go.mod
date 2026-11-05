module github.com/andrerrcosta2/gtools/gflux/core/domain

go 1.26.5

require (
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3
)

replace (
	github.com/andrerrcosta2/gtools/core/data => ../../../core/data
	github.com/andrerrcosta2/gtools/core/domain => ../../../core/domain
)
