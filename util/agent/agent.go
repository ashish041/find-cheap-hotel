package agent

import (
	"log"

	"github.com/google/gops/agent"
)

type Agent struct{}

// Close frees the resources used by the agent.
func (a Agent) Close() {
	agent.Close()
}

// Listen starts the agent and returns a handle for closing the resources used.
func Listen() Agent {
	err := agent.Listen(agent.Options{})
	if err != nil {
		log.Printf("gops/agent.Listen(): %s", err)
	}
	return Agent{}
}
