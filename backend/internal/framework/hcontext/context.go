package hcontext

type ContextKey string

const (
	UserID    ContextKey = "schema_creator_user_id"
	UserAgent ContextKey = "schema_creator_user_agent"
)

func (c ContextKey) String() string {
	return string(c)
}
