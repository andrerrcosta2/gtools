module github.com/andrerrcosta2/gtools/pipes

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/seeders v0.0.1-modular
	github.com/andrerrcosta2/gtools/gtests v0.0.1-modular
)

replace (
	github.com/andrerrcosta2/gtools/core/seeders => ../core/seeders
	github.com/andrerrcosta2/gtools/gtests => ../gtests
)
