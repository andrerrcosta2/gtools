module github.com/andrerrcosta2/gtools/gtests

go 1.23.1

require (
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha3
	github.com/andrerrcosta2/gtools/core/format v0.0.0
	github.com/andrerrcosta2/gtools/gtests/testingseeds v0.0.0
	github.com/andrerrcosta2/gtools/gtests/testingtools v0.0.0
)

require (
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/seeders v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/util v0.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	golang.org/x/text v0.27.0 // indirect
)

replace (
	// Now we know about one more rule about go modules.
	// replace are not transitive by default. That means it won't propagate
	// the replacements up and down. only for submodules.
	github.com/andrerrcosta2/gtools/core/domain => ./../core/domain
	github.com/andrerrcosta2/gtools/core/format => ./../core/format
	github.com/andrerrcosta2/gtools/core/seeders => ./../core/seeders
	github.com/andrerrcosta2/gtools/core/testlite => ./../core/testlite
	github.com/andrerrcosta2/gtools/core/util => ../core/util
	github.com/andrerrcosta2/gtools/gtests/testingseeds => ./testingseeds
	github.com/andrerrcosta2/gtools/gtests/testingtools => ./testingtools
)
