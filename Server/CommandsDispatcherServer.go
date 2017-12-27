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
}

// ICommandsDispatcherServerFactory creates new factory
func ICommandsDispatcherServerFactory() ICommandsDispatcherServer {
	return &commandsDispatcherServer{}
}

func (instance *commandsDispatcherServer) RegisterForCommands(incomingCommandSignal chan Networking.Command) {

	instance.incomingCommandSignal = incomingCommandSignal
	select {
	case cmd := <-instance.incomingCommandSignal:
		{
			instance.HandleCommand(&cmd)
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
	close(instance.incomingCommandSignal)
}
