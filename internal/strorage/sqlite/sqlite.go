package sqlite

import (
	pb "github.com/DanilaNik/BAUMAN-HACK-IU5/github.com/DanilaNik/BAUMAN-HACK-IU5"
	"github.com/DanilaNik/BAUMAN-HACK-IU5/internal/ds"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db    *sqlx.DB
	Rover *ds.Rover
}

func New(storagePath string) *Storage {
	const op = "storage.sqlite.New"
	db, err := sqlx.Open("sqlite3", storagePath)
	if err != nil {
		return nil
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
		return nil
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil
	}

	storage := &Storage{db: db}
	rover := storage.GetRoverByUUID("00112233-4455-6677-8899-aabbccddeeff")
	storage.Rover = rover
	return storage
}

func (s *Storage) MoveRover(uuid string, req *pb.Request) (*ds.Rover, string) {
	const op = "storage.sqlite.MoveRover"

	rover := s.GetRoverByUUID(uuid)

	var warning string

	// switch move {
	// case "up":
	// 	if rover.Y >= 1 {
	// 		rover.Y -= 1
	// 	}
	// case "down":
	// 	if rover.Y+1 >= 95 {
	// 		warning = fmt.Sprintf("Опасность, достигнута максимальная безопасная глубина %d", rover.X)
	// 		break
	// 	}
	// 	rover.Y += 1
	// case "right":
	// 	if rover.X <= 98 {
	// 		rover.X += 1
	// 	}
	// case "left":
	// 	if rover.X >= 1 {
	// 		rover.X -= 1
	// 	}
	// }

	s.UpdateRover(rover)

	// err = s.AddMovementHistory(rover, move)
	// if err != nil {
	// 	return nil, "", fmt.Errorf("%s: %w", op, err)
	// }

	return rover, warning
}

func (s *Storage) GetRoverByUUID(uuid string) *ds.Rover {
	const op = "storage.sqlite.GetRoverByUUID"

	var rover ds.Rover
	err := s.db.Get(&rover, "SELECT * FROM rover WHERE uuid = ?", uuid)
	if err != nil {
		return nil
	}
	return &rover
}

func (s *Storage) UpdateRover(rover *ds.Rover) {
	const op = "storage.sqlite.UpdateRover"

	_, err := s.db.Exec("UPDATE rover SET x = ?, y = ?, z = ? WHERE id = ?", rover.X, rover.Y, rover.Z, rover.ID)
	if err != nil {
		return
	}

	return
}

func (s *Storage) AddMovementHistory(rover *ds.Rover, stage string) {
	const op = "storage.sqlite.AddMovementHistory"

	_, err := s.db.Exec("INSERT INTO movement_history (rover_id, x, y, z, stage) VALUES (?, ?, ?, ?, ?)", rover.ID, rover.X, rover.Y, rover.Z, stage)
	if err != nil {
		return
	}

	return
}
