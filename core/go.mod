module github.com/andrerrcosta2/gtools/core

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3
	//github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core/seeders v0.0.2-alpha.3
)

require github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3

require github.com/google/uuid v1.6.0 // indirect

replace (
	github.com/andrerrcosta2/gtools/core/data => ./data
	github.com/andrerrcosta2/gtools/core/domain => ./domain
	github.com/andrerrcosta2/gtools/core/seeders => ./seeders
)
