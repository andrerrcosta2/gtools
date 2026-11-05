module github.com/andrerrcosta2/gtools/conc/executors

go 1.26.5

replace (
	github.com/andrerrcosta2/gtools/conc/core => ../core
	github.com/andrerrcosta2/gtools/core/domain => ../../core/domain
)

require (
	github.com/andrerrcosta2/gtools/conc/core v0.0.0-00010101000000-000000000000
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3
)
