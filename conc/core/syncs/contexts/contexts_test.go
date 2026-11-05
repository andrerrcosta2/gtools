// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package contexts

import (
	"context"
	"testing"

	"github.com/andrerrcosta2/gtools/conc/core/async"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/contexts/cancelers"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono/durations"
	"github.com/andrerrcosta2/gtools/core/domain/function/runnables"
	"github.com/andrerrcosta2/gtools/core/testlite"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

func TestContexts_ConditionallyCanceled_Create(t *testing.T) {
	main, exit := WithConditionalCancel(context.Background())

	if main == nil {
		t.Fatal("main is nil")
	}
	if exit == nil {
		t.Fatal("exit is nil")
	}
}

func TestContexts_ConditionallyCanceled_Cancellation_All(t *testing.T) {
	main, exit := WithConditionalCancel(context.Background())
	counter := testlite.Map(t)

	runnables.ContextRelease(func() {
		counter.Add(1, "main-flow")

		inner1, exit1 := WithConditionalCancel(main)
		async.ContextReleaseRun(func() {
			chrono.Sleep(durations.Ms(50))
			// Register call after sleep
			counter.Add(1, "main-flow")
			exit1.Now()
		}, inner1)

		inner2, exit2 := WithConditionalCancel(main)
		async.ContextReleaseRun(func() {
			chrono.Sleep(durations.Ms(80))
			// Register call after sleep
			counter.Add(1, "main-flow")
			exit2.Now()
		}, inner2)

		exit.When(cancelers.All())
		flowCalls := counter.Get("main-flow")
		assertlite.True(t, flowCalls == 1, "expected 1 call to 'main-flow' on main context blocking end, got %d", flowCalls)
	}, main)

	flowCalls := counter.Get("main-flow")
	assertlite.True(t, flowCalls == 3, "expected 3 calls to 'main-flow' after leaving main context, got %d", flowCalls)
}

func TestContexts_ConditionallyCanceled_Cancellation_MixedContexts(t *testing.T) {
	main, exit := WithConditionalCancel(context.Background())
	counter := testlite.Map(t)

	runnables.ContextRelease(func() {
		counter.Add(1, "main-flow")

		inner1, exit1 := context.WithCancel(main)
		async.ContextReleaseRun(func() {
			chrono.Sleep(durations.Ms(50))
			// Register call after sleep
			counter.Add(1, "main-flow")
			exit1()
		}, inner1)

		inner2, _ := context.WithTimeout(main, durations.Ms(80).ToStd())
		async.ContextReleaseRun(func() {
			chrono.Sleep(durations.Ms(20))
			// Register call after sleep
			counter.Add(1, "main-flow")
		}, inner2)

		// Is not possible to leverage the native underlying context
		// to track its children into a custom one.
		// For that reason, pipes like cancelers.All and cancelers.Any
		// will not work with native contexts. (At least i didn't find a way)
		exit.When(cancelers.And(inner1, inner2))
		flowCalls := counter.Get("main-flow")
		assertlite.True(t, flowCalls == 1, "expected 1 call to 'main-flow' on main context blocking end, got %d", flowCalls)
	}, main)

	flowCalls := counter.Get("main-flow")
	assertlite.True(t, flowCalls == 3, "expected 3 calls to 'main-flow' after leaving main context, got %d", flowCalls)

}

func TestContexts_ConditionallyCanceled_Cancellation_Nested(t *testing.T) {
	main, exit := WithConditionalCancel(context.Background())
	counter := testlite.Map(t)

	runnables.ContextRelease(func() {
		counter.Add(1, "main-flow")

		inner1, innerExit1 := WithConditionalCancel(main)
		async.ContextReleaseRun(func() {

			chrono.Sleep(durations.Ms(50))

			// Register call after sleep
			counter.Add(1, "main-flow", "inner1-flow")

			inner1_1, innerExit1_1 := WithConditionalCancel(inner1)

			async.ContextReleaseRun(func() {
				chrono.Sleep(durations.Ms(50))
				// Register call after sleep
				counter.Add(1, "main-flow", "inner1-flow")
				innerExit1_1.Now()
			}, inner1_1)

			inner1_2, innerExit1_2 := WithConditionalCancel(inner1)
			async.ContextReleaseRun(func() {
				chrono.Sleep(durations.Ms(50))
				// Register call after sleep
				counter.Add(1, "main-flow", "inner1-flow")
				innerExit1_2.Now()
			}, inner1_2)

			innerExit1.When(cancelers.Any())

			inner1FlowCalls := counter.Get("inner1-flow")
			assertlite.True(t, inner1FlowCalls == 1, "expected 1 call to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
		}, inner1)

		inner2, innerExit2 := WithConditionalCancel(main)

		async.ContextReleaseRun(func() {
			chrono.Sleep(durations.Ms(150))
			// Register call after sleep
			counter.Add(1, "main-flow", "inner2-flow")
			innerExit2.Now()
		}, inner2)

		exit.When(cancelers.All())

		mainFlowCalls := counter.Get("main-flow")
		inner1FlowCalls := counter.Get("inner1-flow")
		inner2FlowCalls := counter.Get("inner2-flow")

		assertlite.True(t, mainFlowCalls == 1, "expected 1 call to 'main-flow' on main context blocking end, got %d", mainFlowCalls)
		assertlite.True(t, inner1FlowCalls == 0, "expected no calls to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
		assertlite.True(t, inner2FlowCalls == 0, "expected no calls to 'inner2-flow' on main context blocking end, got %d", inner2FlowCalls)

	}, main)

	mainFlowCalls := counter.Get("main-flow")
	inner1FlowCalls := counter.Get("inner1-flow")
	inner2FlowCalls := counter.Get("inner2-flow")

	assertlite.True(t, mainFlowCalls == 5, "expected 5 calls to 'main-flow' on main context blocking end, got %d", mainFlowCalls)
	assertlite.True(t, inner1FlowCalls == 3, "expected 3 calls to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
	assertlite.True(t, inner2FlowCalls == 1, "expected 1 call to 'inner2-flow' on main context blocking end, got %d", inner2FlowCalls)

}

func TestContexts_Synchronized_Create(t *testing.T) {
	main, exit := Synchronized(context.Background())

	if main == nil {
		t.Fatal("main is nil")
	}
	if exit == nil {
		t.Fatal("exit is nil")
	}
}

// TestContexts_Synchronized_Cancellation_All: (5,0)-(5,1)
func TestContexts_Synchronized_Cancellation_All(t *testing.T) {
	main, exit := Synchronized(context.Background())
	counter := testlite.Map(t)

	runnables.ContextRelease(func() {
		counter.Add(1, "main-flow")

		inner1, exit1 := Synchronized(main)

		async.ContextReleaseRun(func() {
			chrono.Sleep(durations.Ms(50))
			// Register call after sleep
			counter.Add(1, "main-flow")
			exit1.Now()
		}, inner1)

		inner2, exit2 := Synchronized(main)
		async.ContextReleaseRun(func() {
			chrono.Sleep(durations.Ms(80))
			// Register call after sleep
			counter.Add(1, "main-flow")
			exit2.Now()
		}, inner2)

		exit.When(cancelers.All())
		flowCalls := counter.Get("main-flow")
		assertlite.True(t, flowCalls == 1, "expected 1 call to 'main-flow' on main context blocking end, got %d", flowCalls)
	}, main)

	flowCalls := counter.Get("main-flow")
	assertlite.True(t, flowCalls == 3, "expected 3 calls to 'main-flow' after leaving main context, got %d", flowCalls)

}

// TestContexts_Synchronized_Cancellation_Nesting: test a cancellation of all contexts when
// any nested context is cancelled
func TestContexts_Synchronized_Cancellation_Nesting(t *testing.T) {
	main, exit := Synchronized(context.Background())
	counter := testlite.Map(t)

	runnables.ContextRelease(func() {
		counter.Add(1, "main-flow")

		inner1, innerExit1 := Synchronized(main)
		async.ContextReleaseRun(func() {

			chrono.Sleep(durations.Ms(50))

			// Register call after sleep
			counter.Add(1, "main-flow", "inner1-flow")

			inner1_1, innerExit1_1 := Synchronized(inner1)

			async.ContextReleaseRun(func() {
				chrono.Sleep(durations.Ms(50))
				// Register call after sleep
				counter.Add(1, "main-flow", "inner1-flow")
				innerExit1_1.Now()
			}, inner1_1)

			inner1_2, innerExit1_2 := Synchronized(inner1)
			async.ContextReleaseRun(func() {
				chrono.Sleep(durations.Ms(50))
				// Register call after sleep
				counter.Add(1, "main-flow", "inner1-flow")
				innerExit1_2.Now()
			}, inner1_2)

			innerExit1.When(cancelers.Any())

			inner1FlowCalls := counter.Get("inner1-flow")
			assertlite.True(t, inner1FlowCalls == 1, "expected 1 call to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
		}, inner1)

		inner2, innerExit2 := WithConditionalCancel(main)

		async.ContextReleaseRun(func() {
			chrono.Sleep(durations.Ms(150))
			// Register call after sleep
			counter.Add(1, "main-flow", "inner2-flow")
			innerExit2.Now()
		}, inner2)

		exit.When(cancelers.All())

		mainFlowCalls := counter.Get("main-flow")
		inner1FlowCalls := counter.Get("inner1-flow")
		inner2FlowCalls := counter.Get("inner2-flow")

		assertlite.True(t, mainFlowCalls == 1, "expected 1 call to 'main-flow' on main context blocking end, got %d", mainFlowCalls)
		assertlite.True(t, inner1FlowCalls == 0, "expected no calls to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
		assertlite.True(t, inner2FlowCalls == 0, "expected no calls to 'inner2-flow' on main context blocking end, got %d", inner2FlowCalls)

	}, main)

	mainFlowCalls := counter.Get("main-flow")
	inner1FlowCalls := counter.Get("inner1-flow")
	inner2FlowCalls := counter.Get("inner2-flow")

	assertlite.True(t, mainFlowCalls == 5, "expected 5 calls to 'main-flow' on main context blocking end, got %d", mainFlowCalls)
	assertlite.True(t, inner1FlowCalls == 3, "expected 3 calls to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
	assertlite.True(t, inner2FlowCalls == 1, "expected 1 call to 'inner2-flow' on main context blocking end, got %d", inner2FlowCalls)

}

func TestContexts_Synchronized_Synchronization(t *testing.T) {
	ctx, _ := Synchronized(context.Background())
	counter := testlite.Map(t)

	ctx.Add(501)
	async.ContextReleaseRun(func() {
		for i := 0; i < 500; i++ {
			go func() {
				counter.Add(1, "ctx-flow")
				ctx.Release()
			}()
		}
		ctx.Release()
		ctx.Done()
	}, ctx)

	mainFlowCalls := counter.Get("ctx-flow")
	assertlite.True(t, mainFlowCalls == 0, "expected no calls to 'ctx-flow' on before wait got %d", mainFlowCalls)

	ctx.Wait()
	mainFlowCalls = counter.Get("ctx-flow")
	assertlite.True(t, mainFlowCalls == 500, "expected 500 calls to 'ctx-flow' on after wait got %d", mainFlowCalls)
}
