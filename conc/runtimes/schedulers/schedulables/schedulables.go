// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package schedulables

import (
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/chrono"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/task"
)

func WrapTask(task task.Runnable, at chrono.Time, next ...func(time chrono.Time) (chrono.Time, bool)) TaskWrapper {
	if len(next) > 0 {
		return TaskWrapper{
			task: task,
			at:   at,
			next: next[0],
		}
	}
	return TaskWrapper{
		task: task,
		at:   at,
		next: func(_ chrono.Time) (chrono.Time, bool) {
			return chrono.Time{}, false
		},
	}
}

type TaskWrapper struct {
	task task.Runnable
	at   chrono.Time
	next func(time chrono.Time) (chrono.Time, bool)
}

func (t TaskWrapper) At() chrono.Time {
	return t.at
}

func (t TaskWrapper) Next(from chrono.Time) (chrono.Time, bool) {
	return t.next(from)
}

func (t TaskWrapper) Run() {
	t.task.Run()
}

func (t TaskWrapper) Unwrap() task.Runnable {
	return t.task
}

var _ task.Schedulable[chrono.Time] = TaskWrapper{}
