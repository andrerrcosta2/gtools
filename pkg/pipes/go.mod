module github.com/andrerrcosta2/gtools/pipes

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/seeders v0.0.0
	github.com/andrerrcosta2/gtools/gtests v0.0.0
)

replace (
	github.com/andrerrcosta2/gtools/core/seeders => ../core/seeders
	github.com/andrerrcosta2/gtools/gtests => ../gtests
)
