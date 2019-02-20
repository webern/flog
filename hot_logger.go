// iqBid Pricing System, Copyright (c) 2019 by Inteliquent, Inc.

package flog

import (
	"fmt"
	"sync"
	"time"
)

var hotLogDispatcher *HotLogDispatcher

// HotLogDispatcher is a buffered logger that defers logging in hot loops.
// Although an allocation is made for each queued message, this should only
// be used for logging "stuff that shouldn't happen in the first place and
// might be spammy", e.g. debugging repetitive hot loop conditions.
type HotLogDispatcher struct {
	TimeUntilDispatch time.Duration
	LastDispatch      time.Time
	MessageQueue      []string
	MessageQueueMutex *sync.Mutex
}

func NewHotLogDispatcher() HotLogDispatcher {
	return HotLogDispatcher{
		TimeUntilDispatch: time.Second,
		MessageQueue:      []string{},
		MessageQueueMutex: &sync.Mutex{},
		LastDispatch:      time.Time{},
	}
}

func (d *HotLogDispatcher) Start() {
	go func() {
		for {
			if time.Since(d.LastDispatch) > d.TimeUntilDispatch {
				d.MessageQueueMutex.Lock()
				if len(d.MessageQueue) > 0 {
					messageCountMap := map[string]int{}
					// count duplicates
					for _, msg := range d.MessageQueue {
						if _, ok := messageCountMap[msg]; ok {
							messageCountMap[msg]++
						} else {
							messageCountMap[msg] = 1
						}
					}
					// print messages with count
					for _, msg := range d.MessageQueue {
						if count, ok := messageCountMap[msg]; ok {
							Print(fmt.Sprintf("[%d TIMES] %s", count, msg))
						}
						delete(messageCountMap, msg)
					}
					d.MessageQueue = nil
					d.LastDispatch = time.Now()
				}
				d.MessageQueueMutex.Unlock()
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
}

func (d *HotLogDispatcher) QueueMessage(str string) {
	d.MessageQueueMutex.Lock()
	d.MessageQueue = append(d.MessageQueue, str)
	d.MessageQueueMutex.Unlock()
}

func HotPrint(str string) {
	if hotLogDispatcher == nil {
		d := NewHotLogDispatcher()
		hotLogDispatcher = &d
		hotLogDispatcher.Start()
	}
	hotLogDispatcher.QueueMessage(str)
}
