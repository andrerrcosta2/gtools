module github.com/andrerrcosta2/gtools/datastr

go 1.23.1

require (
    github.com/andrerrcosta2/gtools/core v0.0.1-modular
    github.com/andrerrcosta2/gtools/numbers v0.0.1-modular
    github.com/andrerrcosta2/gtools/pipes v0.0.1-modular
    github.com/andrerrcosta2/gtools/gtests v0.0.1-modular
)

replace (
	github.com/andrerrcosta2/gtools/numbers => ../numbers
	github.com/andrerrcosta2/gtools/core => ../core
	github.com/andrerrcosta2/gtools/pipes => ../pipes
	github.com/andrerrcosta2/gtools/gtests => ../gtests
)
