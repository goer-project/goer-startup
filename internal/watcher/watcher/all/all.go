package all

import (
	"goer-startup/internal/watcher/watcher"
	"goer-startup/internal/watcher/watcher/user"
)

func init() {
	watcher.Register("user", &user.UserWatcher{})
}
