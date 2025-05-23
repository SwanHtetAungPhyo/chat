package auth

import (
	"github.com/SwanHtetAungPhyo/authCognito/internal/model"
	"go.uber.org/fx"
)

var RepoModule = fx.Module("repo_module", fx.Provide(
	NewAuthRepository,
))

type RepoBehaviour interface {
	SignUpSave(req *model.User) (err error)
}
