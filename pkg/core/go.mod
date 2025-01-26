module github.com/andrerrcosta2/gtools/core

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha
	github.com/andrerrcosta2/gtools/core/seeders v0.0.2-alpha
)

require github.com/google/uuid v1.6.0 // indirect

//replace (
//	github.com/andrerrcosta2/gtools/core/data v0.0.1-modular => ./data
//	github.com/andrerrcosta2/gtools/core/domain v0.0.1-modular => ./domain
//	github.com/andrerrcosta2/gtools/core/seeders v0.0.1-modular => ./seeders
//)
