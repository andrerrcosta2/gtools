module github.com/andrerrcosta2/gtools/pipes

go 1.22.0

require (
	github.com/andrerrcosta2/gtools/core v0.0.0
	github.com/andrerrcosta2/gtools/sorts v0.0.0
)

replace (
	github.com/andrerrcosta2/gtools/core => ../core
	github.com/andrerrcosta2/gtools/sorts => ../sorts
)
