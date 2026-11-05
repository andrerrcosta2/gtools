## gtools/conc
___
Collection of concurrency tools for Go.

```shell
  go get github.com/andrerrcosta2/gtools/conc
```

## Packages
___
- `conc/core`: concurrency contracts and compatibility aliases (clock contracts are now sourced from `core/domain/contracts/chrono`).
- `conc/schedulers`: unified scheduler engine built on top of `conc/core` (`Scheduler` + `Task`).
- `conc/executor`: reusable execution runtime (`QueueExecutor`, worker pool, retry helpers, shared supplier/consumer state).
- `conc/async`, `conc/channels`, `conc/streams`, `conc/syncs`: execution/runtime primitives.

## Migration Guide
___
- `conc.NewWorkerPool` -> `executor.NewPool`
- `conc.NewScheduler` -> `schedulers.New`
- `conc.NewExecutorPool` -> `executor.NewPool`
- `conc/core.Clock` -> `core/domain/contracts/chrono.Clock` (preferred)
- `conc/core.SystemClock()` -> `core/runtime/chronos.SystemClock()` (preferred)
- `conc` supplier/consumer helpers -> prefer composition with `core` contracts + `executor` runtime + `schedulers`
- Detailed examples: `conc/MIGRATION.md`
- Release notes draft: `conc/RELEASE_NOTES_NEXT.md`

## Contributing
___
