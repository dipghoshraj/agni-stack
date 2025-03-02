package directive

import (
	"app-gateway/utils"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
)

var jwtSecret = []byte("secret")

func AuthDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	// Get the token from the context
	token, ok := ctx.Value("Authorization").(string)

	if !ok || token == "" {
		return nil, errors.New("missing authorization header")
	}

	// Verify the token
	userid, err := utils.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	// Inject user ID into context for resolvers
	ctx = context.WithValue(ctx, "user_id", userid)

	// Call the next resolver
	return next(ctx)
}
