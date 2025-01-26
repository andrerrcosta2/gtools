module github.com/andrerrcosta2/gtools/datastr

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core v0.0.1-modular
	github.com/andrerrcosta2/gtools/core/data v0.0.1-modular
	github.com/andrerrcosta2/gtools/core/domain v0.0.1-modular
	github.com/andrerrcosta2/gtools/core/seeders v0.0.1-modular
	github.com/andrerrcosta2/gtools/gflux/core v0.0.1-modular
	github.com/andrerrcosta2/gtools/gtests v0.0.1-modular
	github.com/andrerrcosta2/gtools/numbers v0.0.1-modular
	github.com/google/uuid v1.6.0
)

replace (
	github.com/andrerrcosta2/gtools/core => ../core
	github.com/andrerrcosta2/gtools/core/data => ../core/data
	github.com/andrerrcosta2/gtools/core/domain => ../core/domain
	github.com/andrerrcosta2/gtools/core/seeders => ../core/seeders
	github.com/andrerrcosta2/gtools/gflux/core => ../gflux/core
	github.com/andrerrcosta2/gtools/gtests => ../gtests
	github.com/andrerrcosta2/gtools/numbers => ../numbers
)
