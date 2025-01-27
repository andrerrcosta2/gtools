module github.com/andrerrcosta2/gtools/core/seeders

go 1.23.1

require (
    github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.1 // data imports domain
)

replace (
    github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.1 => ../data
)