package message

type Message struct {
	ObjReference string
	MethodID     string
	Args         []byte
	T            int
	ID           int
	StatusCode   int
}

func NewMessage(objRef, methodID string, args []byte, t, statusCode int, idGen *IDGenerator) *Message {
	return &Message{
		ObjReference: objRef,
		MethodID:     methodID,
		Args:         args,
		T:            t,
		ID:           idGen.GetNextID(),
		StatusCode:   statusCode,
	}
}
