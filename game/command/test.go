package command

import (
	"server/leaf/gate"
	"server/msg/protocol"
)

func HandleTestEcho(i interface{}, agent gate.Agent) {
	m := i.(*protocol.TestEcho)
	//LogDebug("TestEcho: %d", m.Int1)
	agent.WriteMsg(m)
}
