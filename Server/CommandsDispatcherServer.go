package Server

import (
	"LampIO/Shared/Networking"
	"log"
)

// ICommandsDispatcherServer handles all incoming commands
// dispatches between specific tasks
type ICommandsDispatcherServer interface {
	RegisterForCommands(incomingCommandSignal chan Networking.Command)
	HandleCommand(cmd *Networking.Command)
	Shutdown()
}

type commandsDispatcherServer struct {
	incomingCommandSignal chan Networking.Command
	endSignal             chan bool
}

// ICommandsDispatcherServerFactory creates new factory
func ICommandsDispatcherServerFactory() ICommandsDispatcherServer {
	this := new(commandsDispatcherServer)
	this.endSignal = make(chan bool)
	return this
}

func (instance *commandsDispatcherServer) RegisterForCommands(incomingCommandSignal chan Networking.Command) {
	instance.incomingCommandSignal = incomingCommandSignal
	for {
		select {
		case cmd := <-instance.incomingCommandSignal:
			{
				instance.HandleCommand(&cmd)
				break
			}
		case <-instance.endSignal:
			{
				log.Println("Server Dispatcher Terminated")
				return
			}
		}
	}
}

func (instance *commandsDispatcherServer) HandleCommand(cmd *Networking.Command) {

	// TODO: impl handling
	switch cmd.CommandType {
	case Networking.NewClientCommandRequestType:
		log.Println("Received new client cmd")
	case Networking.NewModuleCommandRequestType:
		log.Println("Received new module cmd")
	case Networking.NewModuleDataRequestType:
		log.Println("Received module data cmd")
	}
}

func (instance *commandsDispatcherServer) Shutdown() {
	instance.endSignal <- true
	close(instance.incomingCommandSignal)
	close(instance.endSignal)
}
