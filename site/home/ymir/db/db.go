package db

import "database/sql"

type Setting struct{
	name string
	value string
}

type Store interface {
	AddSetting(s *Setting) error
	GetSettings() ([]*Setting, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) AddSetting(s *Setting) error {
	_, err := store.db.Query("INSERT INTO settings(name, value) VALUES ($1,$2)", s.name, s.value)
	return err
}

func (store *dbStore) GetSettings() ([]*Setting, error) {
	rows, err := store.db.Query("SELECT name, value FROM settings")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var settings []*Setting

	for rows.Next() {
		setting := &Setting{}
		if err := rows.Scan(&setting.name, setting.value); err != nil {
			return nil, err
		}

		settings = append(settings, setting)
	}
	return settings, nil
}

var store Store

func InitStore(s Store){
	store = s
}