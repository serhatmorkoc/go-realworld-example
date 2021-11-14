package seed

import (
	"encoding/json"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"os"
)

func Seed(us model.UserStore) error {

	file, err := os.ReadFile("store/shared/db/seed/users.json")
	if err != nil {
		return err
	}

	var users []model.User
	if err := json.Unmarshal(file, &users); err != nil {
		return err
	}

	for _, item := range users {
		_, err := us.Create(&item)
		if err != nil {
			return err
		}
	}

	return nil
}
