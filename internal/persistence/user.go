package persistence

import (
	"context"
	"fmt"
	
	"github.com/dzemildupljak/go_app_tools/internal/entity"
	"github.com/dzemildupljak/go_app_tools/utils"
)

func GetUsers(ctx context.Context) ([]entity.User, error) {
	users := make([]entity.User, 2)
	users[0] = entity.User{
		Email:     "johndoe@example.com",
		Username:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
	}
	users[1] = entity.User{
		Email:     "machelsmith@example.com",
		FirstName: "Machel",
		LastName:  "Smith",
		Username:  "machelsmith",
	}

	err := fmt.Errorf("error in user details persistence func")
	utils.LogError(ctx, err, "error in user details persistence func")

	return nil, err
}
