# Async Dispatchers

The `dispatchers` package provides asynchronous implementations of the core dispatcher contracts.

All implementations are built through composition and decorators, allowing behaviors to be combined without creating large inheritance hierarchies.

---

# Hot Dispatcher

```go
Hot(...)
```

Creates an always-running dispatcher.

Worker goroutines are started immediately and remain available until the dispatcher is closed.

Characteristics:

- asynchronous execution
- configurable worker count
- graceful shutdown
- panic recovery
- optional listeners

---

# Buffered Dispatcher

```go
WithBackpressure(...)
```

Wraps the hot dispatcher with a buffered channel.

The queue introduces natural producer backpressure whenever the buffer becomes full.

This implementation guarantees that producers are slowed rather than losing submitted work.

---

# Retry Decorator

```go
WithRetry(...)
```

Adds retry behavior to any executor.

Retries are fully configurable through:

- retry predicate
- maximum retry count
- backoff policy

This decorator is implementation-independent and may wrap any executor.

---

# Strategy Decorator

```go
WithStrategy(...)
```

Adds adaptive dispatch capabilities using dispatch records.

A strategy may:

- reject work before dispatch
- collect execution information
- implement adaptive admission control
- observe queue behavior
- perform feedback-driven scheduling

Strategies receive metadata through a configurable `MetaProvider`.

---

# Listener Support

Dispatchers may expose listeners.

Listeners receive the dispatched value together with the execution error (if any).

Listeners are intended for lightweight notifications and should avoid blocking execution whenever possible.

---

# Panic Safety

Worker panics are recovered internally.

Recovered panics are converted into regular errors and propagated through listeners.

The dispatcher itself remains operational after handler failures.

---

# Shutdown

Closeable dispatchers perform graceful shutdown.

Shutdown semantics are:

- reject future dispatches
- finish active workers
- drain queued values when appropriate
- notify listeners of rejected work

---

# Design Principles

The implementations favor:

- composition
- decorators
- explicit concurrency
- deterministic shutdown
- implementation reuse

Behavioral extensions are expected to be implemented as wrappers rather than specialized dispatcher types.

---

# Future Ideas

The current implementation intentionally remains small. Possible future extensions include:

## Dispatch Policies

- Drop Oldest
- Drop Newest
- Overwrite
- Ring-buffer dispatchers
- Priority queues
- Deadline-aware dispatching

---

## Scheduling Strategies

- Fair scheduling
- Weighted dispatch
- Priority dispatch
- Work stealing
- Adaptive worker allocation

---

## Reliability

- Circuit breakers
- Bulkheads
- Dead-letter queues
- Failure classification
- Retry budgets

---

## Observability

- OpenTelemetry exporters
- Structured execution records
- Tracing propagation
- Event streams
- Custom metadata providers

---

## Performance

- Lock-free queues
- MPSC/SPMC dispatchers
- Cache-friendly ring buffers
- Adaptive batching
- Zero-allocation dispatch paths

---

## Distributed Dispatch

- Network-backed executors
- Remote worker pools
- Cluster-aware dispatchers
- Distributed queues
- Persistent dispatch logs

---

The goal of this package is not to provide every dispatcher variant, but to offer a small set of composable building blocks from which more specialized execution models can be constructed.