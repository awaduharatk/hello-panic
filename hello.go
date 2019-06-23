package main

import (
	"fmt"
	"net/http"

	"github.com/awaduharatk/hello-panic/errorh"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			if e, ok := err.(*errorh.ErrorMessage); ok {
				fmt.Println("Recover! main()")
				fmt.Println("StatusCd:", e.StatusCd)
			}
		}
	}()

	hello()

	fmt.Println("ここは出力されない")
	defer func() {
		fmt.Println("実行されないDefer")
	}()

}

func hello() {
	defer func() {
		fmt.Println("defer hello()")
		if err := recover(); err != nil {
			// recover後にErrorMessage情報を取得。再度panic呼び出し
			if e, ok := err.(*errorh.ErrorMessage); ok {
				fmt.Println("Recover!")
				fmt.Println("StatusCd:", e.StatusCd)
				panic(err)
			}
		}
	}()
	fmt.Println("hello")
	hellohello()
}

func hellohello() {
	fmt.Println("hellohello")
	helloPanic()
}

func helloPanic() {
	defer fmt.Println("helloPanic")

	// ErrorMessageを生成してpanicに乗せる
	panic(&errorh.ErrorMessage{
		StatusCd: http.StatusInternalServerError,
		Message:  "",
		ErrorCd:  "1005",
		Detail:   nil,
	})

}
