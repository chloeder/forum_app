package memberships

import (
	"context"

	"github.com/chloeder/forum_app/internal/configs"
	"github.com/chloeder/forum_app/internal/models/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, user *memberships.UserModel) error
}


type service struct {
	cfg *configs.Config
	membershipRepo membershipRepository
}

// NewService creates new service instance with repository
func NewService(cfg *configs.Config,membershipRepo membershipRepository) *service {
	return &service{
		cfg: 						cfg,
		membershipRepo: membershipRepo,
	}
}
