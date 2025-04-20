package llm

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
)

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}
