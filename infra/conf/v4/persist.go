package v4

import "github.com/v2fly/v2ray-core/v5/app/persist"

type DbCfg struct {
	Type   string `json:"type"`
	Path   string `json:"path"`
	Uname  string `json:"uname"`
	Passwd string `json:"passwd"`
	Dbname string `json:"dbname"`
}

type PersistConfig struct {
	Enable   bool   `json:"enable"`
	Database *DbCfg `json:"database"`
}

func (c *PersistConfig) Build() (*persist.Config, error) {
	p := &persist.Config{}
	p.Enable = c.Enable
	if c.Database != nil {
		pp, err := c.Database.Build()
		if err != nil {
			return nil, err
		}
		p.Database = pp
	}
	return p, nil
}

func (c *DbCfg) Build() (*persist.DBconfig, error) {
	p := &persist.DBconfig{}
	p.Dbname = c.Dbname
	p.Type = c.Type
	p.Path = c.Path
	p.Uname = c.Uname
	p.Passwd = c.Passwd
	return p, nil
}
