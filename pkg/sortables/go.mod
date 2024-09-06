module github.com/andrerrcosta2/gtools/sortables

go 1.22.0

require (
	github.com/andrerrcosta2/gtools/core v0.0.0
	github.com/andrerrcosta2/gtools/comparables v0.0.0
	github.com/andrerrcosta2/gtools/sorts v0.0.0
    github.com/andrerrcosta2/gtools/typers v0.0.0
)

replace (
   github.com/andrerrcosta2/gtools/core => ../core
   github.com/andrerrcosta2/gtools/comparables => ../comparables
   github.com/andrerrcosta2/gtools/sorts => ../sorts
   github.com/andrerrcosta2/gtools/typers => ../typers
)
