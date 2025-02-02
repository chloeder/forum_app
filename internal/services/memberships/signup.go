package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/chloeder/forum_app/internal/models/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req *memberships.SignUpRequest) error {
	// Check if user already exists
	userExist, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}

	if userExist != nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}


	// Create user
	user := &memberships.UserModel{
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		CreatedAt: time.Time.String(time.Now()),
		UpdatedAt: time.Time.String(time.Now()),
		CreatedBy: req.Username,
		UpdatedBy: req.Username,
	}
	err = s.membershipRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
