module github.com/andrerrcosta2/gtools/gflux

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/gflux/core v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/gflux/obs v0.0.2-alpha.3
// github.com/andrerrcosta2/gtools/gflux/event v0.0.0-1-modular
)

require (
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha3 // indirect
	github.com/google/uuid v1.6.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/core/data => ../core/data
	github.com/andrerrcosta2/gtools/core/domain => ../core/domain
	github.com/andrerrcosta2/gtools/gflux/core => ./core
	github.com/andrerrcosta2/gtools/gflux/obs => ./obs
)
