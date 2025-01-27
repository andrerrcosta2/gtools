module github.com/andrerrcosta2/gtools

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/conc v0.0.2-alpha.1
	github.com/andrerrcosta2/gtools/core v0.0.2-alpha.1
	github.com/andrerrcosta2/gtools/datastr v0.0.2-alpha.1
	github.com/andrerrcosta2/gtools/gflux v0.0.2-alpha.1
	github.com/andrerrcosta2/gtools/gtests v0.0.2-alpha.1
	github.com/andrerrcosta2/gtools/gvalidation v0.0.2-alpha.1
	github.com/andrerrcosta2/gtools/numbers v0.0.2-alpha.1
	github.com/andrerrcosta2/gtools/patterns v0.0.2-alpha.1
	github.com/andrerrcosta2/gtools/reflect4 v0.0.2-alpha.1
)

require (
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.1 // indirect
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.1 // indirect
	github.com/andrerrcosta2/gtools/gflux/core v0.0.2-alpha.1 // indirect
	github.com/andrerrcosta2/gtools/gflux/obs v0.0.2-alpha.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/conc => ./pkg/conc
	github.com/andrerrcosta2/gtools/core => ./pkg/core
	github.com/andrerrcosta2/gtools/core/data => ./pkg/core/data
	github.com/andrerrcosta2/gtools/core/domain => ./pkg/core/domain
	github.com/andrerrcosta2/gtools/core/seeders => ./pkg/core/seeders
	github.com/andrerrcosta2/gtools/datastr => ./pkg/datastr
	github.com/andrerrcosta2/gtools/gflux => ./pkg/gflux
	github.com/andrerrcosta2/gtools/gflux/core => ./pkg/gflux/core
	github.com/andrerrcosta2/gtools/gflux/obs => ./pkg/gflux/obs
	//github.com/andrerrcosta2/gtools/gflux/event => ./pkg/gflux/event
	github.com/andrerrcosta2/gtools/gtests => ./pkg/gtests
	github.com/andrerrcosta2/gtools/gvalidation => ./pkg/gvalidation
	github.com/andrerrcosta2/gtools/numbers => ./pkg/numbers
	github.com/andrerrcosta2/gtools/patterns => ./pkg/patterns
	github.com/andrerrcosta2/gtools/reflect4 => ./pkg/reflect4
)
