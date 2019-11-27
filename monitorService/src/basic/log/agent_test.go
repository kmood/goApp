package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBatchSendSwitch(t *testing.T) {
	agent := NewAgent(nil)
	assert.Equal(t, true, agent.batchSend)

	agent = NewAgent(&AgentConfig{TaskID: "000161", Proto: "unixpacket", Addr: "/var/run/lancer/collector_tcp.sock"})
	assert.Equal(t, true, agent.batchSend)

	agent = NewAgent(&AgentConfig{TaskID: "000161", Proto: "unixgram", Addr: "/var/run/lancer/collector.sock"})
	assert.Equal(t, false, agent.batchSend)
}
