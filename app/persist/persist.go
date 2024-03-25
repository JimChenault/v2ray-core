package persist

import (
	"context"

	"github.com/v2fly/v2ray-core/v5/common"
	"github.com/v2fly/v2ray-core/v5/features/persist"
)

type Recorder struct {
	cfg *Config
	ctx context.Context
}

func NewRecorder(ctx context.Context, config *Config) (*Recorder, error) {
	m := &Recorder{cfg: config, ctx: ctx}
	newError(">>>>persist database: ", config.Database).AtDebug().WriteToLog()
	return m, nil
}

// Type implements common.HasType.
func (*Recorder) Type() interface{} {
	return persist.ManagerType()
}

// Start implements common.Runnable.Start().
func (m *Recorder) Start() error {
	newError(">>>>persist start: ").AtDebug().WriteToLog()
	return nil
}

// Close implements common.Closable.Close().
func (m *Recorder) Close() error {
	return nil
}

func init() {
	//newError(">>>>persist init").AtDebug().WriteToLog()
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return NewRecorder(ctx, config.(*Config))
	}))
}
