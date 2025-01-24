module github.com/andrerrcosta2/gtools/gtests

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/gtools v0.0.0
	github.com/andrerrcosta2/gtools/core/format v0.0.0
)
replace (
	github.com/andrerrcosta2/gtools/core/gtools => ./../core/gtools
	github.com/andrerrcosta2/gtools/core/format => ./../core/format
)
