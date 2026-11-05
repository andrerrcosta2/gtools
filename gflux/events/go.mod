module github.com/andrerrcosta2/gtools/gflux/events

go 1.26.5

require github.com/andrerrcosta2/gtools/conc/core v0.0.0

require github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3 // indirect

replace (
	github.com/andrerrcosta2/gtools/conc/core => ../../conc/core
	github.com/andrerrcosta2/gtools/core/domain => ../../core/domain
)
