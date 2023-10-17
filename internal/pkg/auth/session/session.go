package session

import (
	"context"

	ory "github.com/ory/client-go"
)

const (
	// contextSessionKey is the key used to store the session in the context.
	contextSessionKey = "session"
)

func WithSession(ctx context.Context, session *ory.Session) context.Context {
	return context.WithValue(ctx, contextSessionKey, session)
}

func GetSession(ctx context.Context) *ory.Session {
	sess := ctx.Value(contextSessionKey)
	if sess == nil {
		return nil
	}

	return sess.(*ory.Session)
}
