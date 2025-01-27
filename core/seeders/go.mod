module github.com/andrerrcosta2/gtools/core/seeders

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.2 // data imports domain
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.2
	github.com/google/uuid v1.6.0
)

replace (
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.2 => ../data
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.2 => ../domain
)
