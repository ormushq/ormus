package inmemorytaskrepo

import (
	"sync"

	"github.com/ormushq/ormus/destination/entity/taskentity"
)

type DB struct {
	tasks map[string]*taskentity.Task
	mutex *sync.Mutex
}

func New() DB {
	db := DB{
		tasks: make(map[string]*taskentity.Task),
		mutex: &sync.Mutex{},
	}

	return db
}
