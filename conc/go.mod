module github.com/andrerrcosta2/gtools/conc

go 1.26.5

require (
	github.com/andrerrcosta2/gtools/conc/actors v0.0.0
	github.com/andrerrcosta2/gtools/conc/core v0.0.0
	github.com/andrerrcosta2/gtools/conc/dispatchers v0.0.0
	github.com/andrerrcosta2/gtools/conc/executors v0.0.0
	github.com/andrerrcosta2/gtools/conc/streams v0.0.0
	github.com/andrerrcosta2/gtools/conc/testing v0.0.0
	github.com/andrerrcosta2/gtools/conc/timing/runtimes v0.0.0
	github.com/andrerrcosta2/gtools/conc/timing/schedulers v0.0.0
)

require (
	github.com/andrerrcosta2/gtools/conc/timing/engines v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/chronos v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/data v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/domain v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/core/format v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/core/util v0.0.0 // indirect
	github.com/andrerrcosta2/gtools/datastr v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/gflux/core v0.0.2-alpha.3 // indirect
	github.com/andrerrcosta2/gtools/gflux/core/domain v0.0.0 // indirect
	golang.org/x/text v0.40.0 // indirect
)

replace (
	github.com/andrerrcosta2/gtools/conc/actors => ./actors

	github.com/andrerrcosta2/gtools/conc/core => ./core
	github.com/andrerrcosta2/gtools/conc/dispatchers => ./dispatchers
	github.com/andrerrcosta2/gtools/conc/executors => ./executors
	github.com/andrerrcosta2/gtools/conc/streams => ./streams
	github.com/andrerrcosta2/gtools/conc/testing => ./testing
	github.com/andrerrcosta2/gtools/conc/timing/runtimes => ./runtimes
	github.com/andrerrcosta2/gtools/conc/timing/engines => ./runtimes/schedulers/engines
	github.com/andrerrcosta2/gtools/conc/timing/schedulers => ./runtimes/schedulers
	github.com/andrerrcosta2/gtools/core/chronos => ../core/chronos
	github.com/andrerrcosta2/gtools/core/data => ../core/data
	github.com/andrerrcosta2/gtools/core/domain => ../core/domain
	github.com/andrerrcosta2/gtools/core/format => ../core/format
	github.com/andrerrcosta2/gtools/core/testlite => ../core/testlite
	github.com/andrerrcosta2/gtools/core/util => ../core/util
	github.com/andrerrcosta2/gtools/datastr => ../datastr
	github.com/andrerrcosta2/gtools/datastr/testing => ../datastr/testing
	github.com/andrerrcosta2/gtools/gflux/core => ../gflux/core
	github.com/andrerrcosta2/gtools/gflux/core/domain => ../gflux/core/domain
)
