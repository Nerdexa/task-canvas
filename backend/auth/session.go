package auth

import (
	"task-canvas/config"
	"task-canvas/pkg/auth0"

	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/pkg/errors"
)

type Session struct {
	token       jwt.Token
	tokenString string
	userinfo    *auth0.Userinfo
}

func NewSession(tokenString string, token jwt.Token) Session {
	return Session{
		token:       token,
		tokenString: tokenString,
	}
}

func (s Session) GetUserId() string {
	return s.token.Subject()
}

func (s Session) GetEmail() (string, error) {
	if s.userinfo != nil {
		return s.userinfo.Email, nil
	}
	domain := config.GetConfig().Auth0.Domain

	a := auth0.NewAuth0(domain)
	u, err := a.GetUserInfo(s.tokenString)
	if err != nil {
		err = errors.Wrap(err, "ユーザー情報の取得に失敗しました")
		return "", err
	}

	s.userinfo = u
	return s.userinfo.Email, nil
}

func (s Session) GetSessionTokenString() string {
	return s.tokenString
}
