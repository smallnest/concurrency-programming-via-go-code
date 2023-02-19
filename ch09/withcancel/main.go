package main

import (
	"context"
	"crypto/sha256"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
)

func main() {
	targetBits, _ := strconv.Atoi(os.Args[1])
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	pow := func(ctx context.Context, targetBits int, ch chan string) {
		target := big.NewInt(1)
		target.Lsh(target, uint(256-targetBits)) // 除去前targetBits位，其余位都是1

		var hashInt big.Int
		var hash [32]byte
		nonce := 0 // 随机数

		// 寻找一个数
		log.Println("开始寻找一个数，使得hash值小于目标值")
		for {
			select {
			case <-ctx.Done():
				log.Println("context is canceled")
				ch <- ""
				return
			default:
				data := "hello world " + strconv.Itoa(nonce)
				hash = sha256.Sum256([]byte(data)) // 计算hash值
				hashInt.SetBytes(hash[:])          // 将hash值转换为big.Int

				if hashInt.Cmp(target) < 1 { // hashInt < target, 找到一个超过目标数值的数，也就是至少前targetBits位为0
					ch <- data
					return
				} else { // 没找到，继续找
					nonce++
				}
			}

		}
	}

	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan string, 1)
	go pow(ctx, targetBits, ch)

	time.Sleep(time.Second)

	defer cancel()

	select {
	case result := <-ch:
		log.Println("找到一个比目标值小的数：", result)
		return
	default:
		log.Println("没有找到比目标值小的数:", ctx.Err())
	}
}
