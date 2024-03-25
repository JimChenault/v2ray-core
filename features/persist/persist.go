package persist

import (
	"github.com/v2fly/v2ray-core/v5/features"
)

type Manager interface {
	features.Feature

	//NewRecorder(config interface{}) (r interface{}, err error)
	//Connect(db interface{}) (err error)
	//Close() (err error)
	//Save(usr memory.User) error
	//Get(uname string) (*memory.User, error)
	//Update(usr memory.User) error
}

// NoopManager is an implementation of Manager, which doesn't has actual functionalities.
type NoopManager struct{}

func ManagerType() interface{} {
	return (*Manager)(nil)
}

// Type implements common.HasType.
func (NoopManager) Type() interface{} {
	return ManagerType()
}

// Start implements common.Runnable.
func (NoopManager) Start() error { return nil }

// Close implements common.Closable.
func (NoopManager) Close() error { return nil }
