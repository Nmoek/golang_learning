package simpleim

const eventName = "simple_im"

type Event struct {
	Msg      Message
	Receiver int64
}
