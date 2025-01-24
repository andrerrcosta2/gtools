// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package contexts

import (
	"context"
	"github.com/andrerrcosta2/gtools/conc/contexts/cancelers"
	"github.com/andrerrcosta2/gtools/core/durations"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"testing"
	"time"
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
	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)

	tt.StackLog("Starting main context")

	tt.RunBlocking(func() {
		tt.StackLog("Creating inner context")
		tt.RegisterCalls(1, "main-flow")

		inner1, exit1 := WithConditionalCancel(main)
		tt.AsyncBefore(func() {
			tt.StackLog("Inside inner1 context")
			time.Sleep(durations.Ms(50))
			// Register call after sleep
			tt.RegisterCalls(1, "main-flow")
			tt.StackLog("Exiting inner1 context")
			exit1.Now()
		}, inner1)

		tt.StackLog("Creating inner2 context")
		inner2, exit2 := WithConditionalCancel(main)
		tt.AsyncBefore(func() {
			tt.StackLog("Inside inner2 context")
			time.Sleep(durations.Ms(80))
			// Register call after sleep
			tt.RegisterCalls(1, "main-flow")
			tt.StackLog("Exiting inner2 context")
			exit2.Now()
		}, inner2)

		exit.When(cancelers.All())
		flowCalls := tt.CallsTo("main-flow")
		tt.Condition(flowCalls == 1, "expected 1 call to 'main-flow' on main context blocking end, got %d", flowCalls)
	}, main)

	tt.StackLog("Main context cancelled")

	flowCalls := tt.CallsTo("main-flow")
	tt.Condition(flowCalls == 3, "expected 3 calls to 'main-flow' after leaving main context, got %d", flowCalls)

	tt.PrintLogStack()
}

func TestContexts_ConditionallyCanceled_Cancellation_MixedContexts(t *testing.T) {
	main, exit := WithConditionalCancel(context.Background())
	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)

	tt.StackLog("Starting main context")

	tt.RunBlocking(func() {
		tt.StackLog("Creating inner context")
		tt.RegisterCalls(1, "main-flow")

		inner1, exit1 := context.WithCancel(main)
		tt.AsyncBefore(func() {
			tt.StackLog("Inside inner1 context")
			time.Sleep(durations.Ms(50))
			// Register call after sleep
			tt.RegisterCalls(1, "main-flow")
			tt.StackLog("Exiting inner1 context")
			exit1()
		}, inner1)

		tt.StackLog("Creating inner2 context")
		inner2, _ := context.WithTimeout(main, durations.Ms(80))
		tt.AsyncBefore(func() {
			tt.StackLog("Inside inner2 context")
			time.Sleep(durations.Ms(20))
			// Register call after sleep
			tt.RegisterCalls(1, "main-flow")
			tt.StackLog("inner2 context bottom")
		}, inner2)

		// Is not possible to leverage the native underlying context
		// to track its children into a custom one.
		// For that reason, pipes like cancelers.All and cancelers.Any
		// will not work with native contexts. (At least i didn't find a way)
		exit.When(cancelers.And(inner1, inner2))
		flowCalls := tt.CallsTo("main-flow")
		tt.Condition(flowCalls == 1, "expected 1 call to 'main-flow' on main context blocking end, got %d", flowCalls)
	}, main)

	tt.StackLog("Main context cancelled")

	flowCalls := tt.CallsTo("main-flow")
	tt.Condition(flowCalls == 3, "expected 3 calls to 'main-flow' after leaving main context, got %d", flowCalls)

	tt.PrintLogStack()
}

func TestContexts_ConditionallyCanceled_Cancellation_Nested(t *testing.T) {
	main, exit := WithConditionalCancel(context.Background())

	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)
	tt.StackLog("Starting main context")

	tt.RunBlocking(func() {
		tt.StackLog("Creating inner1 context")
		tt.RegisterCalls(1, "main-flow")

		inner1, innerExit1 := WithConditionalCancel(main)
		tt.AsyncBefore(func() {

			tt.StackLog("Inside inner1 context")
			time.Sleep(durations.Ms(50))

			// Register call after sleep
			tt.RegisterCalls(1, "main-flow", "inner1-flow")

			inner1_1, innerExit1_1 := WithConditionalCancel(inner1)

			tt.AsyncBefore(func() {
				tt.StackLog("Inside inner1_1 context")
				time.Sleep(durations.Ms(50))
				// Register call after sleep
				tt.RegisterCalls(1, "main-flow", "inner1-flow")
				tt.StackLog("Exiting inner1_1 context")
				innerExit1_1.Now()
			}, inner1_1)

			inner1_2, innerExit1_2 := WithConditionalCancel(inner1)
			tt.AsyncBefore(func() {
				tt.StackLog("Inside inner1_2 context")
				time.Sleep(durations.Ms(50))
				// Register call after sleep
				tt.RegisterCalls(1, "main-flow", "inner1-flow")
				tt.StackLog("Exiting inner1_2 context")
				innerExit1_2.Now()
			}, inner1_2)

			innerExit1.When(cancelers.Any())

			inner1FlowCalls := tt.CallsTo("inner1-flow")
			tt.Condition(inner1FlowCalls == 1, "expected 1 call to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
		}, inner1)

		tt.StackLog("Creating inner2 context")
		inner2, innerExit2 := WithConditionalCancel(main)

		tt.AsyncBefore(func() {
			tt.StackLog("Inside inner2 context")
			time.Sleep(durations.Ms(150))
			// Register call after sleep
			tt.RegisterCalls(1, "main-flow", "inner2-flow")
			tt.StackLog("inner2 context bottom")
			innerExit2.Now()
		}, inner2)

		exit.When(cancelers.All())

		mainFlowCalls := tt.CallsTo("main-flow")
		inner1FlowCalls := tt.CallsTo("inner1-flow")
		inner2FlowCalls := tt.CallsTo("inner2-flow")

		tt.Condition(mainFlowCalls == 1, "expected 1 call to 'main-flow' on main context blocking end, got %d", mainFlowCalls)
		tt.Condition(inner1FlowCalls == 0, "expected no calls to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
		tt.Condition(inner2FlowCalls == 0, "expected no calls to 'inner2-flow' on main context blocking end, got %d", inner2FlowCalls)

		tt.StackLog("Main context bottom")

	}, main)

	tt.StackLog("Main context cancelled")

	mainFlowCalls := tt.CallsTo("main-flow")
	inner1FlowCalls := tt.CallsTo("inner1-flow")
	inner2FlowCalls := tt.CallsTo("inner2-flow")

	tt.Condition(mainFlowCalls == 5, "expected 5 calls to 'main-flow' on main context blocking end, got %d", mainFlowCalls)
	tt.Condition(inner1FlowCalls == 3, "expected 3 calls to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
	tt.Condition(inner2FlowCalls == 1, "expected 1 call to 'inner2-flow' on main context blocking end, got %d", inner2FlowCalls)

	tt.PrintLogStack()
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
	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)

	tt.StackLog("Starting main context")

	tt.RunBlocking(func() {
		tt.StackLog("Creating inner context")
		tt.RegisterCalls(1, "main-flow")

		inner1, exit1 := Synchronized(main)

		tt.AsyncBefore(func() {
			tt.StackLog("Inside inner1 context")
			time.Sleep(durations.Ms(50))
			// Register call after sleep
			tt.RegisterCalls(1, "main-flow")
			tt.StackLog("Exiting inner1 context")
			exit1.Now()
		}, inner1)

		tt.StackLog("Creating inner2 context")
		inner2, exit2 := Synchronized(main)
		tt.AsyncBefore(func() {
			tt.StackLog("Inside inner2 context")
			time.Sleep(durations.Ms(80))
			// Register call after sleep
			tt.RegisterCalls(1, "main-flow")
			tt.StackLog("Exiting inner2 context")
			exit2.Now()
		}, inner2)

		exit.When(cancelers.All())
		flowCalls := tt.CallsTo("main-flow")
		tt.Condition(flowCalls == 1, "expected 1 call to 'main-flow' on main context blocking end, got %d", flowCalls)
	}, main)

	tt.StackLog("Main context cancelled")

	flowCalls := tt.CallsTo("main-flow")
	tt.Condition(flowCalls == 3, "expected 3 calls to 'main-flow' after leaving main context, got %d", flowCalls)

	tt.PrintLogStack()
}

// TestContexts_Synchronized_Cancellation_Nesting: test a cancellation of all contexts when
// any nested context is cancelled
func TestContexts_Synchronized_Cancellation_Nesting(t *testing.T) {
	main, exit := Synchronized(context.Background())

	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)
	tt.StackLog("Starting main context")

	tt.RunBlocking(func() {
		tt.StackLog("Creating inner1 context")
		tt.RegisterCalls(1, "main-flow")

		inner1, innerExit1 := Synchronized(main)
		tt.AsyncBefore(func() {

			tt.StackLog("Inside inner1 context")
			time.Sleep(durations.Ms(50))

			// Register call after sleep
			tt.RegisterCalls(1, "main-flow", "inner1-flow")

			inner1_1, innerExit1_1 := Synchronized(inner1)

			tt.AsyncBefore(func() {
				tt.StackLog("Inside inner1_1 context")
				time.Sleep(durations.Ms(50))
				// Register call after sleep
				tt.RegisterCalls(1, "main-flow", "inner1-flow")
				tt.StackLog("Exiting inner1_1 context")
				innerExit1_1.Now()
			}, inner1_1)

			inner1_2, innerExit1_2 := Synchronized(inner1)
			tt.AsyncBefore(func() {
				tt.StackLog("Inside inner1_2 context")
				time.Sleep(durations.Ms(50))
				// Register call after sleep
				tt.RegisterCalls(1, "main-flow", "inner1-flow")
				tt.StackLog("Exiting inner1_2 context")
				innerExit1_2.Now()
			}, inner1_2)

			innerExit1.When(cancelers.Any())

			inner1FlowCalls := tt.CallsTo("inner1-flow")
			tt.Condition(inner1FlowCalls == 1, "expected 1 call to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
		}, inner1)

		tt.StackLog("Creating inner2 context")
		inner2, innerExit2 := WithConditionalCancel(main)

		tt.AsyncBefore(func() {
			tt.StackLog("Inside inner2 context")
			time.Sleep(durations.Ms(150))
			// Register call after sleep
			tt.RegisterCalls(1, "main-flow", "inner2-flow")
			tt.StackLog("inner2 context bottom")
			innerExit2.Now()
		}, inner2)

		exit.When(cancelers.All())

		mainFlowCalls := tt.CallsTo("main-flow")
		inner1FlowCalls := tt.CallsTo("inner1-flow")
		inner2FlowCalls := tt.CallsTo("inner2-flow")

		tt.Condition(mainFlowCalls == 1, "expected 1 call to 'main-flow' on main context blocking end, got %d", mainFlowCalls)
		tt.Condition(inner1FlowCalls == 0, "expected no calls to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
		tt.Condition(inner2FlowCalls == 0, "expected no calls to 'inner2-flow' on main context blocking end, got %d", inner2FlowCalls)

		tt.StackLog("Main context bottom")

	}, main)

	tt.StackLog("Main context cancelled")

	mainFlowCalls := tt.CallsTo("main-flow")
	inner1FlowCalls := tt.CallsTo("inner1-flow")
	inner2FlowCalls := tt.CallsTo("inner2-flow")

	tt.Condition(mainFlowCalls == 5, "expected 5 calls to 'main-flow' on main context blocking end, got %d", mainFlowCalls)
	tt.Condition(inner1FlowCalls == 3, "expected 3 calls to 'inner1-flow' on main context blocking end, got %d", inner1FlowCalls)
	tt.Condition(inner2FlowCalls == 1, "expected 1 call to 'inner2-flow' on main context blocking end, got %d", inner2FlowCalls)

	tt.PrintLogStack()
}

func TestContexts_Synchronized_Synchronization(t *testing.T) {
	ctx, _ := Synchronized(context.Background())

	tt := testingtools.AsyncLite(t, gtests.LogOnFailure)
	tt.StackLog("Starting ctx context")

	ctx.Add(501)
	tt.AsyncBefore(func() {
		for i := 0; i < 500; i++ {
			go func() {
				tt.RegisterCalls(1, "ctx-flow")
				ctx.Release()
			}()
		}
		ctx.Release()
		ctx.Done()
	}, ctx)

	mainFlowCalls := tt.CallsTo("ctx-flow")
	tt.Condition(mainFlowCalls == 0, "expected no calls to 'ctx-flow' on before wait got %d", mainFlowCalls)

	ctx.Wait()
	mainFlowCalls = tt.CallsTo("ctx-flow")
	tt.Condition(mainFlowCalls == 500, "expected 500 calls to 'ctx-flow' on after wait got %d", mainFlowCalls)
}
