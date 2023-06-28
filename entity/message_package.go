package entity

type MessagePackage struct {
	Agent     string      `json:"agent"`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp"`
	AgentId   string      `json:"agent_id"`
	Resend    int16       `json:"resend"`
	Type      string      `json:"type"`
}
