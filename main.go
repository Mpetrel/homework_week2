package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"homework_week2/db"
)




func main() {
	db.InitDB()
	userService := NewUserService()
	user, err := userService.GetUserInfo(2)

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			fmt.Println("user not exist!")
		} else {
			fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
			fmt.Printf("stack trace: \n%+v\n", err)
		}
		return
	}
	fmt.Println(user)
}
