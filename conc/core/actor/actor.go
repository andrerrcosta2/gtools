// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package actor

import "context"

type Actor[M any] interface {
	Receive(ctx context.Context, msg M)
}

type Ref[M any] interface {
	Tell(msg M)
	Ask(ctx context.Context, msg M) error
	Stop()
}

type System[M any] interface {
	Spawn(Actor[M]) Ref[M]
}
