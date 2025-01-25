module github.com/andrerrcosta2/gtools/core/seeders

go 1.23.1

require (
    github.com/andrerrcosta2/gtools/core/data v0.0.1-modular // data imports gtools
    github.com/andrerrcosta2/gtools/core/typers v0.0.1-modular // doesn't import gtools
)

//replace (
//    github.com/andrerrcosta2/gtools/core/data => ../data
//    github.com/andrerrcosta2/gtools/core/typers => ../typers
//)