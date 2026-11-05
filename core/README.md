## gtools/core
___
Collection of tools for Go.

```shell
  go get github.com/andrerrcosta2/gtools/core
```

## Modules
___
- [core](README.md)
- [data](data/README.md)
- [domain](domain/README.md)
- [runtime](runtime/README.md)
- [seeders](seeders/README.md)
- [testlite](testlite/README.md)

## Architecture Notes
___
- `core/domain`: contracts, constraints, and domain-level interfaces.
- `core/runtime`: runtime infrastructure (execution, chrono, and other implementation-oriented utilities).
- `core/seeders`: deterministic and pseudo-random seed/value generation utilities.
- `core/testlite`: lightweight testing assertions and test helpers.

## Migration
___
- `core/domain/execution` -> `core/runtime/execution` (compat wrapper kept).
- Clock abstractions moved to `core/domain/contracts/chrono`.
- System clock implementation is `core/runtime/chronos.SystemClock()`.
- Detailed mapping and examples: `core/MIGRATION.md`.
