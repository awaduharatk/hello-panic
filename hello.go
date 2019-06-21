package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer")
		err := recover()
		if err != nil {
			fmt.Println("Recover!:", err)
		}
	}()

	hello()
	// hello()
	fmt.Println("ここは出力されない")

}

func hello() {
	//止めたいところでdefer書いとけば止めれる
	// mainに戻ってPrintして処理狩猟ルートに落ちる

	// defer func() {
	// 	fmt.Println("defer hello()")
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println("Recover!:", err)
	// 	}
	// }()
	fmt.Println("hello")
	hellohello()
}

func hellohello() {
	fmt.Println("hellohello")
	helloPanic()
}

func helloPanic() {
	defer fmt.Println("helloPanic")
	panic("helloPanic!!")
}
