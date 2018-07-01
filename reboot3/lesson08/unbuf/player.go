package unbuf

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Player info
type Player struct {
	User1 string
	User2 string
}

// NewPlayer to create a Player object
func NewPlayer(u1, u2 string) *Player {
	return &Player{
		User1: u1,
		User2: u2,
	}
}

// StartPlayer to start play
func (p *Player) StartPlayer() {
	ch := make(chan int, 0)
	var wg sync.WaitGroup

	fmt.Println("start playing!!")
	wg.Add(2)
	go player(p.User1, ch, &wg)
	go player(p.User2, ch, &wg)

	ch <- 1
	wg.Wait()
}

// player for two players
func player(name string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s starting!\n", name)
	var n int
	for {
		ball, ok := <-ch
		if !ok {
			// 对方失败，关了通道
			fmt.Printf("%s won!!!\n", name)
			break
		}
		n = rand.Intn(100)
		if n%19 == 0 {
			// 自己失败
			fmt.Printf("%s miss, the number is %d\n", name, n)
			close(ch)
			break
		}
		ball++
		fmt.Printf("Player %s hit ball %d with rand %d\n", name, ball, n)
		ch <- ball
	}
}
