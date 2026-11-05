// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package actors

type localActor[M any] struct {
	mailbox chan M
	stop    chan struct{}
}

func (a *localActor[M]) loop(handler func(M)) {
	for {
		select {
		case msg := <-a.mailbox:
			handler(msg)

		case <-a.stop:
			return
		}
	}
}
