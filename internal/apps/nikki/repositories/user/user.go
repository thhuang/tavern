package user

import (
	"github.com/google/uuid"
	"github.com/thhuang/go-server/internal/apps/nikki/models/user"
)

type UserRepository interface {
	Get(uuid.UUID) (user.User, error)
}
