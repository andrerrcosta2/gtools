// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package stepper

import (
	"errors"
)

var (
	ErrStepNotFound = errors.New("step not found")
)

// Stepper is an executor of position-ordered broadcasters in a pull-driven manner.
// Conceptually, this is a minimal discrete-event simulator kernel:
//   - Steps are ordered by a logical coordinate T (e.g., time, sequence number)
//
// This abstraction is useful for:
//   - Deterministic testing of time-dependent systems
//   - Incremental computation with explicit progression control
//   - Protocol simulation with causal ordering
//
// See also: Discrete Event Simulation (DES) kernels, small-step operational semantics.
type Stepper interface {
	// Step executes the next due step (small-step semantics)
	// 	Returns:
	//   - executed: whether a step was executed
	//   - err: error returned by the step (if executed)
	Step() (executed bool, errs error)
}

// Advanceable extends Stepper with target-based progression and position tracking.
// Conceptually, this is a minimal discrete-event simulator kernel:
//   - Steps are ordered by a logical coordinate T (e.g., time, sequence number)
//   - Step() executes the next due step (small-step semantics)
//   - StepTo(target) advances the logical clock, executing all due broadcasters
//   - Now() reports the current logical position
//
// This abstraction is useful for:
//   - Deterministic testing of time-dependent systems
//   - Incremental computation with explicit progression control
//   - Protocol simulation with causal ordering
//
// See also: Discrete Event Simulation (DES) kernels, small-step operational semantics.
type Advanceable[T any] interface {
	Stepper

	// StepTo advances until no pending step is due at or before target.
	// It returns true when at least one step was executed, and all
	// step errors
	StepTo(target T) (executed bool, errs []error)

	// Now returns the current logical position of the stepper.
	Now() T
}
