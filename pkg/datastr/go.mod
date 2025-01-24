module github.com/andrerrcosta2/gtools/datastr

go 1.23.1

require (
    github.com/andrerrcosta2/gtools/core v0.0.0
    github.com/andrerrcosta2/gtools/numbers v0.0.0
    github.com/andrerrcosta2/gtools/pipes v0.0.0
)

replace (
	github.com/andrerrcosta2/gtools/numbers => ../numbers
	github.com/andrerrcosta2/gtools/core => ../core
	github.com/andrerrcosta2/gtools/pipes => ../pipes
)
