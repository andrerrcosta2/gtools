module github.com/andrerrcosta2/gtools/core

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/gtools v0.0.0 // No dependencies
	github.com/andrerrcosta2/gtools/core/seeders v0.0.0
)

require (
	github.com/andrerrcosta2/gtools/core/data v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/typers v0.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/core/data => ./data
	github.com/andrerrcosta2/gtools/core/format => ./format
	github.com/andrerrcosta2/gtools/core/gtools => ../core/gtools
	github.com/andrerrcosta2/gtools/core/seeders => ./seeders
	github.com/andrerrcosta2/gtools/core/typers => ./typers
)
