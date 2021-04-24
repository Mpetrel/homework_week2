package main

import (
	"fmt"
	"github.com/pkg/errors"
	"homework_week2/db"
)

/**
 * 选择向上抛出错误，但统一提供一些方法进行判断是否属于“正常”的错误
 * 以便于进行处理，同时避免在多出引入sql包
 */


func main() {
	db.InitDB()
	userService := NewUserService()
	user, err := userService.GetUserInfo(2)

	if err != nil {
		// 检查是否是需要对应处理的错误
		if IsEmptyError(err) {
			fmt.Println("user not exist!")
		} else { // 其他异常进行处理
			fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
			fmt.Printf("stack trace: \n%+v\n", err)
		}
		return
	}
	fmt.Println(user)
}
