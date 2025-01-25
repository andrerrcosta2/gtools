module github.com/andrerrcosta2/gtools/patterns

go 1.23.1

require github.com/andrerrcosta2/gtools/core/domain v0.0.1-modular

replace (
	github.com/andrerrcosta2/gtools/core/domain => ../core/domain
)
