package sqlite

import (
	"fmt"

	"github.com/DanilaNik/BAUMAN-HACK-IU5/internal/ds"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"
	db, err := sqlx.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS rover (
		id INTEGER PRIMARY KEY,
		uuid TEXT UNIQUE NOT NULL,
		name TEXT UNIQUE NOT NULL,
		x INTEGER NOT NULL CHECK (x >= 0 AND x <= 99),
		y INTEGER NOT NULL CHECK (y >= 0 AND y <= 99),
		z INTEGER NOT NULL CHECK (y >= 0 AND y <= 99),
		angle INTEGER NOT NULL,
		charge INTEGER NOT NULL
	  );
	  
	  CREATE TABLE IF NOT EXISTS movement_history (
		id INTEGER PRIMARY KEY,
		rover_id INTEGER NOT NULL,
		x INTEGER NOT NULL,
		y INTEGER NOT NULL,
		z INTEGER NOT NULL,
		stage TEXT NOT NULL,
		FOREIGN KEY (rover_id) REFERENCES rover(id)
	  );
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) MoveRover(uuid string, move string) (*ds.Rover, error) {
	const op = "storage.sqlite.MoveRover"

	rover, err := s.GetRoverByUUID(uuid)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	switch move {
	case "up":
		rover.Y -= 1
	case "down":
		rover.Y += 1
	case "right":
		rover.X += 1
	case "left":
		rover.X -= 1
	}

	err = s.UpdateRover(rover)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = s.AddMovementHistory(rover, move)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return rover, nil
}

func (s *Storage) GetRoverByUUID(uuid string) (*ds.Rover, error) {
	const op = "storage.sqlite.GetRoverByUUID"

	var rover ds.Rover
	err := s.db.Get(&rover, "SELECT * FROM rover WHERE uuid = ?", uuid)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &rover, nil
}

func (s *Storage) UpdateRover(rover *ds.Rover) error {
	const op = "storage.sqlite.UpdateRover"

	_, err := s.db.Exec("UPDATE rover SET x = ?, y = ?, z = ? WHERE id = ?", rover.X, rover.Y, rover.Z, rover.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) AddMovementHistory(rover *ds.Rover, stage string) error {
	const op = "storage.sqlite.AddMovementHistory"

	_, err := s.db.Exec("INSERT INTO movement_history (rover_id, x, y, z, stage) VALUES (?, ?, ?, ?, ?)", rover.ID, rover.X, rover.Y, rover.Z, stage)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
