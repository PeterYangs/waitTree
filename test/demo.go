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

			defer func() {

				son.Done()

				//从父类中释放
				son.Release()
			}()

			time.Sleep(1 * time.Second)

			fmt.Println("son done")

		}()

	}

	for i := 0; i < 10; i++ {

		son2.Add(1)

		go func() {

			defer func() {

				son2.Done()

				//从父类中释放
				son2.Release()
			}()

			time.Sleep(2 * time.Second)

			fmt.Println("son2 done")

		}()

	}

	father.Wait()

	father.Wait()

	//释放所有子wait
	//father.ReleaseSon()

}
