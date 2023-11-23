package auth

import (
	"context"
	"fmt"
	"task-canvas/config"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/pkg/errors"
)

type key string

const (
	sessionKey key = "sessionKey"
)

func GetSession(ctx context.Context) (*Session, error) {
	v := ctx.Value(sessionKey)

	if v == nil {
		return nil, errors.New("セッション情報の取得に失敗しました")
	}

	token := v.(string)

	tok, err := getJwtToken(token)
	if err != nil {
		return nil, err
	}

	s := NewSession(token, tok)
	return &s, nil
}

func SetSessionKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, sessionKey, key)
}

func getJwtToken(token string) (jwt.Token, error) {
	domain := config.GetConfig().Auth0.Domain
	set, err := jwk.Fetch(context.Background(), fmt.Sprintf("https://%s/.well-known/jwks.json", domain))
	if err != nil {
		return nil, err
	}

	tok, err := jwt.Parse(
		[]byte(token),
		jwt.WithKeySet(set),
		jwt.WithVerify(true),
		jwt.WithValidate(true),
	)

	if err != nil {
		err = errors.Wrap(err, "正しいトークンではありません")
		return nil, err
	}

	return tok, nil
}
