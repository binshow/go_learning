package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


// unbuffered channel_demo 演示案例一
// 演示两个 goroutine 相互打网球比赛
// 使用一个 unbuffered channel_demo 来保证求同时被两边击中或miss

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// use unbuffered channel_demo
	court := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// launch two players
	go func() {
		player("Honanh" , court)
		wg.Done()
	}()

	go func() {
		player("Andrew" , court)
		wg.Done()
	}()

	// 开始比赛
	court <- 1

	wg.Wait() // Wait for the game to finish.

}



// player simulates a person playing the game of tennis.
// We are asking for a channel_demo value using value semantic.
func player(name string, court chan int) {
	for {
		// Wait for the ball to be hit back to us.
		// Notice that this is another form of receive. Instead of getting just the value, we can
		// get a flag indicating how the receive is returned. If the signal happens because of the
		// data, ok will be true. If the signal happens without data, in other word, the channel_demo is
		// closed, ok will be false. In this case, we are gonna use that to determine who won.
		ball, ok := <-court
		if !ok {
			// 当前 goroutine 在接球，如果没有另外一个goroutine没有往channel中成功发送数据，则说明当前goroutine赢了
			// If the channel_demo was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball (or we lose).
		// If we lose the game, we are gonna close the channel_demo. It then causes the other player to
		// know that he is receiving the signal but without data. The channel_demo is closed so he won.
		// They both return.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel_demo to signal we lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		// If the 2 cases above doesn't happen, we still have the ball. Increase the value of the
		// ball by one and perform a send. We know that the other player is still in receive mode,
		// therefore, the send and receive will eventually come together.
		// Again, in an unbuffered channel_demo, the receive happens first because it gives us the
		// guarantee.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}