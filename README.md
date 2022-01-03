```go
package main

import (
	"fmt"
	"time"
	"waitTree"
)

func main() {

	father := waitTree.NewWaitTree(waitTree.Background())

    son:= waitTree.NewWaitTree(father)
    son2:= waitTree.NewWaitTree(father)

	for i := 0; i < 10; i++ {

		son.Add(1)

		go func() {

			defer son.Done()

			time.Sleep(1* time.Second)

			fmt.Println("son done")

		}()

	}

	for i := 0; i < 10; i++ {

		son2.Add(1)

		go func() {

			defer son2.Done()

			time.Sleep(2* time.Second)

			fmt.Println("son2 done")

		}()

	}

	father.Wait()

	//释放子wait
	father.Release()


}
```