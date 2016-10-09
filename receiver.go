package amqp

type FnReceiveMessage func()
type Receiver struct {
}

func (r *Receiver) Receive() {

}

func (r *Receiver) OnReceive(fn FnReceiveMessage) {

}
