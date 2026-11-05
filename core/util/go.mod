module github.com/andrerrcosta2/gtools/core/util

go 1.26.5

require (
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/testlite v0.0.0
	golang.org/x/text v0.40.0
)

require (
	github.com/andrerrcosta2/gtools/core/chronos v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/format v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/seeders v0.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/core/chronos => ../chronos
	github.com/andrerrcosta2/gtools/core/data => ../data
	github.com/andrerrcosta2/gtools/core/domain => ../domain
	github.com/andrerrcosta2/gtools/core/format => ../format
	github.com/andrerrcosta2/gtools/core/seeders => ../seeders
	github.com/andrerrcosta2/gtools/core/testlite => ../testlite
)
