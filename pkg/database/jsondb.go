package database

import (
	"encoding/json"
	"github.com/linqcod/users-service-task/internal/domain/model"
	"github.com/spf13/viper"
	"io/fs"
	"os"
)

type UserStore struct {
	dbFile    string
	Increment int            `json:"increment"`
	List      model.UserList `json:"list"`
}

func InitDB() (*UserStore, error) {
	dbFile := viper.GetString("DB_PATH")

	f, _ := os.ReadFile(dbFile)

	s := &UserStore{
		dbFile: dbFile,
	}

	err := json.Unmarshal(f, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *UserStore) UpdateDB() error {
	b, _ := json.Marshal(s)

	err := os.WriteFile(s.dbFile, b, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
