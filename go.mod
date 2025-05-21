module github.com/andrerrcosta2/gtools

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/conc v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/datastr v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/gflux v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/gtests v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/gvalidation v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/numbers v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/patterns v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/reflect4 v0.0.2-alpha.3
)

require (
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/format v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/seeders v0.0.2-alpha.3 //
	github.com/andrerrcosta2/gtools/gflux/core v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/gflux/obs v0.0.2-alpha.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/conc => ./conc
	github.com/andrerrcosta2/gtools/core => ./core
	github.com/andrerrcosta2/gtools/core/data => ./core/data
	github.com/andrerrcosta2/gtools/core/domain => ./core/domain
	github.com/andrerrcosta2/gtools/core/format => ./core/format
	github.com/andrerrcosta2/gtools/core/seeders => ./core/seeders
	github.com/andrerrcosta2/gtools/datastr => ./datastr
	github.com/andrerrcosta2/gtools/gflux => ./gflux
	github.com/andrerrcosta2/gtools/gflux/core => ./gflux/core
	github.com/andrerrcosta2/gtools/gflux/obs => ./gflux/obs
	//github.com/andrerrcosta2/gtools/gflux/event => ./gflux/event
	github.com/andrerrcosta2/gtools/gtests => ./gtests
	github.com/andrerrcosta2/gtools/gvalidation => ./gvalidation
	github.com/andrerrcosta2/gtools/numbers => ./numbers
	github.com/andrerrcosta2/gtools/patterns => ./patterns
	github.com/andrerrcosta2/gtools/reflect4 => ./reflect4
)
