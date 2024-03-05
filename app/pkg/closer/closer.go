package closer

import (
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"sync"
)

var closer = New()

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

		c.fns = nil
		c.mu.Unlock()

		var err error

		for _, fn := range fns {
			go func(fn func() error) {
				if ferr := fn(); err != nil {
					// Не очень элегантно, возможно. Но предостеречься от гонки стоит
					c.mu.Lock()
					defer c.mu.Unlock()
					err = errors.Join(err, ferr)
				}
			}(fn)
		}

		if err != nil {
			slog.Error("failed to close all functions", slog.String("error", err.Error()))
		}

		return
	})
}
