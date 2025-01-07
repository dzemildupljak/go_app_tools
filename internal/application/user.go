package application

import (
	"context"

	"github.com/dzemildupljak/go_app_tools/internal/entity"
	"github.com/dzemildupljak/go_app_tools/internal/persistence"
	"github.com/dzemildupljak/go_app_tools/utils"
)

func GetUsers(ctx context.Context) ([]entity.User, error) {
	users, err := persistence.GetUsers(ctx)
	if err != nil {
		utils.LogError(ctx, err, "error in user details application func")
		return nil, err
	}

	return users, nil
}
