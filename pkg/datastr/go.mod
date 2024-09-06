module github.com/andrerrcosta2/gtools/datastr

go 1.22.0

require (
    github.com/andrerrcosta2/gtools/functions v0.0.0
    github.com/andrerrcosta2/gtools/nums v0.0.0
    github.com/andrerrcosta2/gtools/ranges v0.0.0
	github.com/andrerrcosta2/gtools/constraints v0.0.0
    github.com/andrerrcosta2/gtools/core v0.0.0
    github.com/andrerrcosta2/gtools/testdata v0.0.0
)

replace github.com/andrerrcosta2/gtools/functions => ../core/functions
replace github.com/andrerrcosta2/gtools/nums => ./../numbers
replace github.com/andrerrcosta2/gtools/ranges => ../ranges
replace github.com/andrerrcosta2/gtools/constraints => ./../core/constraints
replace github.com/andrerrcosta2/gtools/core => ../core
replace github.com/andrerrcosta2/gtools/testdata => ../testdata
