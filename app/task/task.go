package task

import "eve/internal/crontab"

func Tasks() []crontab.TaskInterface {
	return []crontab.TaskInterface{
		&FooTask{},
	}
}
