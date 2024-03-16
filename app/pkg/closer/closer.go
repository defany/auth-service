package closer

import (
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/exp/slices"
)

var closer = New(os.Interrupt, syscall.SIGTERM)

func Add(fns ...func() error) {
	closer.Add(fns...)
}

func Close() {
	closer.Close()

	// Я думаю, что мы можем сделать так, так как Wait неотьемлема от Close
	closer.Wait()
}

type Closer struct {
	mu   sync.Mutex
	once sync.Once
	done chan struct{}
	fns  []func() error
}

func New(sig ...os.Signal) *Closer {
	c := &Closer{
		done: make(chan struct{}),
	}

	if len(sig) == 0 {
		return c
	}

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, sig...)
		<-ch

		signal.Stop(ch)

		c.Close()
	}()

	return c
}

func (c *Closer) Add(fns ...func() error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.fns = append(c.fns, fns...)

	return
}

func (c *Closer) Wait() {
	<-c.done
}

func (c *Closer) Close() {
	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()
		fns := c.fns

		slices.Reverse(fns)

		c.fns = nil
		c.mu.Unlock()

		var err error

		for _, fn := range fns {
			// TODO: if close too long - kill
			if ferr := fn(); err != nil {
				err = errors.Join(err, ferr)
			}
		}

		if err != nil {
			slog.Error("failed to close all functions", slog.String("error", err.Error()))
		}

		return
	})
}
