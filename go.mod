module github.com/andrerrcosta2/gtools

go 1.26.5

require (
	github.com/andrerrcosta2/gtools/ast4 v0.0.0-00010101000000-000000000000
	github.com/andrerrcosta2/gtools/conc v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/core v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/datastr v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/gflux v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/gvalidation v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/numbers v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/patterns v0.0.2-alpha.3
	github.com/andrerrcosta2/gtools/reflect4 v0.0.2-alpha.3
)

require (
	github.com/andrerrcosta2/gtools/conc/core v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/chronos v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/testlite v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/util v0.0.0 // indirect
	golang.org/x/text v0.40.0 // indirect
)

require (
	github.com/andrerrcosta2/gtools/conc/actors v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/conc/dispatchers v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/conc/executors v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/conc/streams v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/conc/testing v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/conc/timing/runtimes v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/conc/timing/engines v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/conc/timing/schedulers v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/format v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/io v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/seeders v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/gflux/core v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/gflux/core/domain v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/gflux/obs v0.0.2-alpha.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/ast4 => ./ast4
	github.com/andrerrcosta2/gtools/conc => ./conc
	github.com/andrerrcosta2/gtools/conc/actors => ./conc/actors
	github.com/andrerrcosta2/gtools/conc/core => ./conc/core
	github.com/andrerrcosta2/gtools/conc/dispatchers => ./conc/dispatchers
	github.com/andrerrcosta2/gtools/conc/executors => ./conc/executors
	github.com/andrerrcosta2/gtools/conc/streams => ./conc/streams
	github.com/andrerrcosta2/gtools/conc/testing => ./conc/testing
	github.com/andrerrcosta2/gtools/conc/timing/runtimes => ./conc/runtimes
	github.com/andrerrcosta2/gtools/conc/timing/engines => ./conc/runtimes/schedulers/engines
	github.com/andrerrcosta2/gtools/conc/timing/schedulers => ./conc/runtimes/schedulers
	github.com/andrerrcosta2/gtools/core => ./core
	github.com/andrerrcosta2/gtools/core/chronos => ./core/chronos
	github.com/andrerrcosta2/gtools/core/data => ./core/data
	github.com/andrerrcosta2/gtools/core/domain => ./core/domain
	github.com/andrerrcosta2/gtools/core/format => ./core/format
	github.com/andrerrcosta2/gtools/core/io => ./core/io
	github.com/andrerrcosta2/gtools/core/runtime => ./core/runtime
	github.com/andrerrcosta2/gtools/core/search => ./core/search
	github.com/andrerrcosta2/gtools/core/seeders => ./core/seeders
	github.com/andrerrcosta2/gtools/core/selection => ./core/selection
	github.com/andrerrcosta2/gtools/core/sortables => ./core/sortables
	github.com/andrerrcosta2/gtools/core/testlite => ./core/testlite
	github.com/andrerrcosta2/gtools/core/util => ./core/util
	github.com/andrerrcosta2/gtools/datastr => ./datastr
	github.com/andrerrcosta2/gtools/datastr/testing => ./datastr/testing
	github.com/andrerrcosta2/gtools/gflux => ./gflux
	github.com/andrerrcosta2/gtools/gflux/core => ./gflux/core
	github.com/andrerrcosta2/gtools/gflux/core/domain => ./gflux/core/domain
	github.com/andrerrcosta2/gtools/gflux/events => ./gflux/events
	github.com/andrerrcosta2/gtools/gflux/interceptors => ./gflux/interceptors
	github.com/andrerrcosta2/gtools/gflux/obs => ./gflux/obs
	github.com/andrerrcosta2/gtools/gvalidation => ./gvalidation
	github.com/andrerrcosta2/gtools/numbers => ./numbers
	github.com/andrerrcosta2/gtools/patterns => ./patterns
	github.com/andrerrcosta2/gtools/reflect4 => ./reflect4
)
