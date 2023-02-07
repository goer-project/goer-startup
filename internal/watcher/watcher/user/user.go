package user

import (
	"context"

	"github.com/go-redsync/redsync/v4"

	"goer-startup/internal/apiserver/store"
	"goer-startup/internal/pkg/log"
	"goer-startup/internal/watcher/watcher"
)

type userWatcher struct {
	ctx   context.Context
	mutex *redsync.Mutex
}

// Run runs the watcher job.
func (w *userWatcher) Run() {
	if err := w.mutex.Lock(); err != nil {
		log.C(w.ctx).Infow("userWatcher already run.")

		return
	}
	defer func() {
		if _, err := w.mutex.Unlock(); err != nil {
			log.C(w.ctx).Errorw("could not release userWatcher lock. err: %v", err)

			return
		}
	}()

	user, err := store.S.Users().Get(w.ctx, "test")
	if err != nil {
		log.Errorw(err.Error())

		return
	}

	log.Infow(user.Email)
}

// Spec is parsed using the time zone of clean Cron instance as the default.
func (w *userWatcher) Spec() string {
	return "@every 5s"
}

// Init initializes the watcher for later execution.
func (w *userWatcher) Init(ctx context.Context, rs *redsync.Mutex, config interface{}) error {
	*w = userWatcher{
		ctx:   ctx,
		mutex: rs,
	}

	return nil
}

func init() {
	watcher.Register("user", &userWatcher{})
}
