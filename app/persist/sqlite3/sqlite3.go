package sqlite3

import (
	"database/sql"

	"github.com/v2fly/v2ray-core/v5/app/persist/memory"
	//_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Path string
}

type SqliteMng struct {
	Config
	db *sql.DB
}

func NewRecorder(cfg Config) (db *SqliteMng, err error) {
	sm := &SqliteMng{}
	sm.Path = cfg.Path
	return sm, nil
}

func (s *SqliteMng) Connect(db interface{}) (err error) {
	d, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return err
	}
	s.db = d
	return nil
}

func (s SqliteMng) Close(db interface{}) {
	s.db.Close()
}

func (s *SqliteMng) Save(usr memory.User) error {
	return nil
}

func (s SqliteMng) Get(uname string) (*memory.User, error) {

	return nil, nil
}
func (s *SqliteMng) Update(usr memory.User) error {
	return nil
}
