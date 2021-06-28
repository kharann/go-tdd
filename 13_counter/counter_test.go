package main

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Increment the coutner 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		expected := 3
		assert.Equal(t, counter.Value(), expected)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assert.Equal(t, counter.Value(), wantedCount)

	})
}
