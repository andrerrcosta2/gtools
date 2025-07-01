module github.com/andrerrcosta2/gtools/core

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha3
	github.com/andrerrcosta2/gtools/core/format v0.0.0
	github.com/andrerrcosta2/gtools/core/seeders v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/testlite v0.0.0
	github.com/andrerrcosta2/gtools/core/util v0.0.0
)

require github.com/google/uuid v1.6.0 // indirect

replace (
	github.com/andrerrcosta2/gtools/core/data => ./data
	github.com/andrerrcosta2/gtools/core/domain => ./domain
	github.com/andrerrcosta2/gtools/core/format => ./format
	github.com/andrerrcosta2/gtools/core/seeders => ./seeders
	github.com/andrerrcosta2/gtools/core/testlite => ./testlite
	github.com/andrerrcosta2/gtools/core/util => ./util
)
