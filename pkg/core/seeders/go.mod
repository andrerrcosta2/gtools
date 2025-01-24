module github.com/andrerrcosta2/gtools/core/seeders

go 1.23.1

require (
    github.com/andrerrcosta2/gtools/core/data v0.0.0 // data imports gtools
    github.com/andrerrcosta2/gtools/core/typers v0.0.0 // doesn't import gtools
)

replace (
    github.com/andrerrcosta2/gtools/core/data => ../data
    github.com/andrerrcosta2/gtools/core/typers => ../typers
)