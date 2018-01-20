package Handler

type IIncomingConnectionsHandler interface {
	AddNewConnection(incomingConn IIncomingConnection) (int, error)
	DeleteConnection(id int) error
}

type incomingConnectionsHandler struct {
	connections map[int]IIncomingConnection
}

func (handler *incomingConnectionsHandler) AddNewConnection(incomingConn IIncomingConnection) (int, error) {
	newID := len(handler.connections) + 1
	handler.connections[newID] = incomingConn
	return newID, nil
}

func (handler *incomingConnectionsHandler) DeleteConnection(id int) error {
	return nil
}
