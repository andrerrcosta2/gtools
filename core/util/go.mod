module github.com/andrerrcosta2/gtools/core/util

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha3
	github.com/andrerrcosta2/gtools/core/testlite v0.0.0
	golang.org/x/text v0.27.0
)

replace (
	github.com/andrerrcosta2/gtools/core/domain => ../domain
	github.com/andrerrcosta2/gtools/core/testlite => ../testlite
)
