module github.com/andrerrcosta2/gtools/core/format

go 1.26.5

require (
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/testlite v0.0.0
	github.com/andrerrcosta2/gtools/core/util v0.0.0
)

require golang.org/x/text v0.40.0 // indirect

replace (
	github.com/andrerrcosta2/gtools/core/domain => ../domain
	github.com/andrerrcosta2/gtools/core/testlite => ../testlite
	github.com/andrerrcosta2/gtools/core/util => ../util
)
