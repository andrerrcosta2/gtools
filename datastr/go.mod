module github.com/andrerrcosta2/gtools/datastr

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha3
	github.com/andrerrcosta2/gtools/core/format v0.0.0
	github.com/andrerrcosta2/gtools/core/seeders v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/gflux/core v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/gtests v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/gtests/testingtools v0.0.0
	github.com/andrerrcosta2/gtools/numbers v0.0.2-alpha.3
	github.com/google/uuid v1.6.0
)

require (
	github.com/andrerrcosta2/gtools/core/util v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/gtests/testingseeds v0.0.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/core => ../core
	github.com/andrerrcosta2/gtools/core/data => ../core/data
	github.com/andrerrcosta2/gtools/core/domain => ../core/domain
	github.com/andrerrcosta2/gtools/core/format => ../core/format
	github.com/andrerrcosta2/gtools/core/seeders => ../core/seeders
	github.com/andrerrcosta2/gtools/core/testlite => ../core/testlite
	github.com/andrerrcosta2/gtools/core/util => ../core/util
	github.com/andrerrcosta2/gtools/gflux/core => ../gflux/core
	github.com/andrerrcosta2/gtools/gtests => ../gtests
	github.com/andrerrcosta2/gtools/gtests/testingseeds => ../gtests/testingseeds
	github.com/andrerrcosta2/gtools/gtests/testingtools => ../gtests/testingtools
	github.com/andrerrcosta2/gtools/numbers => ../numbers
)
