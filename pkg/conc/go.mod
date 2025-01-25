module github.com/andrerrcosta2/gtools/conc

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core v0.0.1-modular
	github.com/andrerrcosta2/gtools/gtests v0.0.1-modular
	github.com/andrerrcosta2/gtools/gflux/core v0.0.1-modular
)

replace (
	github.com/andrerrcosta2/gtools/core => ../core
	github.com/andrerrcosta2/gtools/gtests => ../gtests
	github.com/andrerrcosta2/gtools/gflux/core => ../gflux/core
)