```go
package main

import (
	"fmt"
	"github.com/PeterYangs/waitTree"
	"time"
)

func main() {

	father := waitTree.NewWaitTree(waitTree.Background())

	son := waitTree.NewWaitTree(father)
	son2 := waitTree.NewWaitTree(father)

	for i := 0; i < 10; i++ {

		son.Add(1)

		go func() {

			defer son.Done()

			time.Sleep(1 * time.Second)

			fmt.Println("son done")

		}()

	}

	for i := 0; i < 10; i++ {

		son2.Add(1)

		go func() {

			defer son2.Done()

			time.Sleep(2 * time.Second)

			fmt.Println("son2 done")

		}()

	}

	father.Wait()

	son.Wait()

	//从父wait中删除
	son.Release()

}
```