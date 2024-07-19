package realtime

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestPubSub(t *testing.T) {
	is := is.New(t)

	client := NewClient()
	err := client.NewTopic("test")
	is.NoErr(err)

	lch, cancel, err := client.Listen("test")
	is.NoErr(err)

	payload := fmt.Sprint(time.Now().Unix())

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range lch {
			p, ok := data.Data.(string)
			is.True(ok && p == payload)
			println("received:", p)
			return
		}
	}()

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
		defer cancel()

		err = client.Send(ctx, Message{Topic: "test", Strategy: RouteToFirst, Data: payload})
		println("sent:", payload)
		is.NoErr(err)
		wg.Wait()
	}

	{
		cancel()

		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
		defer cancel()

		println("sending payload on 0 listener topic")
		err = client.Send(ctx, Message{Topic: "test", Strategy: RouteToFirst, Data: "bad-payload"})
		is.True(err != nil)
		println("received an expected error:", err.Error())
	}
}
