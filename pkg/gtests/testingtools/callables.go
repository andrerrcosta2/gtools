// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"encoding/json"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"strings"
	"sync"
)

func caller(value int, name string, callback functions.Runnable) *call {
	return &call{value: value, name: name, callback: callback}
}

type call struct {
	mtx      sync.RWMutex
	name     string
	callOn   int
	value    int
	callback functions.Runnable
}

func (c *call) add(value int) int {
	c.mtx.Lock()

	defer func() {
		if c.value == c.callOn {
			c.callback()
		}
		c.mtx.Unlock()
	}()

	c.value += value
	return c.value
}

func (c *call) setCallback(callback functions.Runnable, on int) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.callOn = on
	c.callback = callback
}

func (c *call) Name() string {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.name
}

func (c *call) load() int {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.value
}

func (c *call) String() string {
	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")

	// Marshal a struct copy that omits mutex and function fields
	err := enc.Encode(struct {
		Name   string `json:"name"`
		CallOn int    `json:"call_on"`
		Value  int    `json:"value"`
	}{
		Name:   c.name,
		CallOn: c.callOn,
		Value:  c.value,
	})

	if err != nil {
		return fmt.Sprintf("Error marshaling to JSON: %v\n", err)
	}
	return buf.String()
}
