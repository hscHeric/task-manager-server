package db

import "database/sql"

type DatabaseService struct {
	Db *sql.DB
}

func NewDatabaseService(db *sql.DB) *DatabaseService {
	return &DatabaseService{
		Db: db,
	}
}

func (db *DatabaseService) InsertTask(v any) error {
	return nil
}

func (db *DatabaseService) GetTaskByID(id string) (any, error) {
	return nil, nil
}

func (db *DatabaseService) DeleteTask(id string) error {
	return nil
}
