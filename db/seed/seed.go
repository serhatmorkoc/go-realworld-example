package seed

import (
	"encoding/json"
	"fmt"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"io/ioutil"
)

func Seed(us model.UserStore) {

	file, err := ioutil.ReadFile("db/seed/users.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var users []model.User
	if err := json.Unmarshal(file, &users); err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range users {
		_, err := us.Create(&item)
		if err != nil {
			fmt.Println(err)
		}
	}
}
