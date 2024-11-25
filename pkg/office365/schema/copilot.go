package schema

// CopilotBase represents the base schema for copilot records.
type CopilotBase struct {
	AuditRecord
	CopilotEventData struct {
		ThreadID string `json:"ThreadId,omitempty"`
	} `json:"CopilotEventData,omitempty"`
}
