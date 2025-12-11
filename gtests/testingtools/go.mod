module github.com/andrerrcosta2/gtools/gtests/testingtools

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha3
	github.com/andrerrcosta2/gtools/core/format v0.0.0
)

require (
	github.com/andrerrcosta2/gtools/core/util v0.0.0 // indirect
	golang.org/x/text v0.27.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/core/domain => ./../../core/domain
	github.com/andrerrcosta2/gtools/core/format => ./../../core/format
	github.com/andrerrcosta2/gtools/core/util => ./../../core/util
)
