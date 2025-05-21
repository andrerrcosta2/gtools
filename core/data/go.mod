module github.com/andrerrcosta2/gtools/core/data

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha3
	github.com/andrerrcosta2/gtools/core/format v0.0.0
)

replace (
	github.com/andrerrcosta2/gtools/core/domain => ../domain
	github.com/andrerrcosta2/gtools/core/format => ../format
)
