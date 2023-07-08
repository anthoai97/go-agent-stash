package entity

type FileExecuteStatus struct {
	Success   bool   `json:"success"`
	AgentID   string `json:"agent_id"`
	MessageID string `json:"message_id"`
	Path      string `json:"path"`
	Error     error  `json:"error"`
}
