package fswatcher

import "github.com/fsnotify/fsnotify"

type FSWatcher struct {
	watcher *fsnotify.Watcher
}

func NewFSWatcher(watcher *fsnotify.Watcher) *FSWatcher {
	return &FSWatcher{
		watcher: watcher,
	}
}

func (w FSWatcher) GetEventsChan() <-chan fsnotify.Event {
	return w.watcher.Events
}

func (w FSWatcher) GetErrorsChan() <-chan error {
	return w.watcher.Errors
}
