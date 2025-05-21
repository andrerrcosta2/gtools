module github.com/andrerrcosta2/gtools/gtests

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha3
	github.com/andrerrcosta2/gtools/core/format v0.0.0
	github.com/andrerrcosta2/gtools/gtests/testingseeds v0.0.0
	github.com/andrerrcosta2/gtools/gtests/testingtools v0.0.0
)

replace (
	github.com/andrerrcosta2/gtools/core/domain => ./../core/domain
	github.com/andrerrcosta2/gtools/core/format => ./../core/format
	github.com/andrerrcosta2/gtools/gtests/testingseeds => ./testingseeds
	github.com/andrerrcosta2/gtools/gtests/testingtools => ./testingtools
)
