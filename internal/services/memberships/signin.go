package memberships

import (
	"context"
	"errors"

	"github.com/chloeder/forum_app/internal/models/memberships"
	"github.com/chloeder/forum_app/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignIn(ctx context.Context, req *memberships.SignInRequest) (string, error) {
	// Check if user already exists
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Password)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	// Check if password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Generate token
	token, err := jwt.CreateToken(user.ID, user.Username, user.Email, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
