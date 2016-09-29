//fizzBuzz.go

package main

import "fmt"

func main() {

	for i := 0; i <= 1000; i++ {
		fizz := i%3 == 0
		buzz := i%5 == 0
		zero := i == 0

		switch {
		//i = 「3の倍数」かつ「5の倍数でない」かつ「0でない」
		case fizz && !buzz && !zero:
			fmt.Println("Fizz!")
			//i = 「3の倍数でない」かつ「5の倍数」かつ「0でない」
		case !fizz && buzz && !zero:
			fmt.Println("Buzz!")
			//i = 「3の倍数」かつ「5の倍数」（15の倍数）かつ「0でない」
		case fizz && buzz && !zero:
			fmt.Println("FizzBuzz!")
			//i = 「3の倍数でない」かつ「5の倍数でない」かつ「0でない」
		default:
			fmt.Println(i)

		}
	}

}

/*

以下のコードだとi != 0がbool型、i%hoge =0 がint型ということでエラーだった

package main

import "fmt"

func main() {
          for i := 0; i <= 1000; i++ {

            switch i {
            case i != 0 && i % 3 =0:
                fmt.Println("Fizz!")

            case i != 0 && i % 3 =0:
                fmt.Println("Buzz!")

            case i != 0 && i % 3 =0:
                fmt.Println("FizzBuzz!")

            default :
                fmt.Println(i)

            }

          }

*/
