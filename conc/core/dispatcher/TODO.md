# Dispatcher

The `dispatcher` package defines the core contracts for asynchronous work dispatching within `gtools`.

A dispatcher is responsible for accepting values and delivering them to a handler for execution. The package intentionally defines only the orchestration contracts; scheduling policies, buffering strategies, retry mechanisms and execution models belong to implementation packages.

This package is dependency-free and serves as the foundation for higher-level concurrency components.

---

## Core Concepts

The minimal abstraction is an executor.

```go
type Executor[T any] interface {
    Dispatch(T) error
}
```

An executor accepts work and decides how and when it will be processed.

The execution may be:

- synchronous;
- asynchronous;
- buffered;
- delayed;
- scheduled;
- distributed.

The interface intentionally makes no assumptions.

---

## Capabilities

Instead of exposing a large hierarchy of dispatcher types, the package models orthogonal capabilities.

### Closeable

```go
type Closeable[T any]
```

Represents executors that can be gracefully shut down.

---

### Buffered

```go
type Buffered[T any]
```

Represents executors exposing queue information.

Implementations are free to decide what constitutes a queue.

---

### Listenable

```go
type Listenable[T any]
```

Allows observing dispatch results.

Listeners are notified after a dispatch operation completes.

---

### Startable

```go
type Startable[T any]
```

Represents executors requiring an explicit start phase.

This is useful for lazy initialization or controlled startup.

---

## Strategies

Strategies allow dispatch behavior to be adapted without modifying the executor itself.

```go
type Strategy[T any]
```

A strategy may:

- reject dispatches
- apply admission control
- implement feedback loops
- collect execution statistics
- throttle producers
- perform adaptive scheduling

Strategies are intended to be composed using decorators rather than inheritance.

---

## Dispatch Records

The package provides a standard execution record.

```go
type Record[T any]
```

A record represents metadata associated with a dispatch operation.

Typical information includes:

- submitted value
- queue position
- timestamps
- execution status
- execution errors
- implementation-specific attributes

The record exists to support instrumentation and adaptive execution policies.

The package deliberately stores facts instead of computed metrics. Higher-level monitoring systems are expected to derive latency, throughput and utilization from recorded timestamps.

---

## Meta Providers

A `MetaProvider` is responsible for creating and updating dispatch records.

Different providers may enrich records with:

- tracing information
- execution context
- custom attributes
- distributed identifiers
- telemetry metadata

This allows implementations to remain independent from any particular monitoring framework.

---

## Design Goals

- Minimal orchestration contracts
- Capability-oriented interfaces
- Composition over inheritance
- Implementation-independent APIs
- No dependency on observability frameworks
- Suitable as a foundation for schedulers, worker pools and pipelines