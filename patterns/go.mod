module github.com/andrerrcosta2/gtools/patterns

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/conc v0.0.2-alpha.2
	github.com/andrerrcosta2/gtools/core v0.0.2-alpha.2
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.2
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.2
	github.com/andrerrcosta2/gtools/gtests v0.0.2-alpha.2
)

require (
	github.com/andrerrcosta2/gtools/core/seeders v0.0.2-alpha.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/conc => ../conc
	github.com/andrerrcosta2/gtools/core => ../core
	github.com/andrerrcosta2/gtools/core/data => ../core/data
	github.com/andrerrcosta2/gtools/core/domain => ../core/domain
	github.com/andrerrcosta2/gtools/core/seeders => ../core/seeders
	github.com/andrerrcosta2/gtools/gtests => ../gtests
)
