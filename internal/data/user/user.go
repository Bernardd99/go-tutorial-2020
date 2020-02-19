package user

import (
	"context"
	"log"

	"go-tutorial-2020/pkg/errors"

	userEntity "go-tutorial-2020/internal/entity/user"

	"github.com/jmoiron/sqlx"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

const (
	getAllUsers  = "GetAllUsers"
	qGetAllUsers = "SELECT * FROM user_test"
)

var (
	readStmt = []statement{
		{getAllUsers, qGetAllUsers},
	}
)

// New ...
func New(db *sqlx.DB) Data {
	d := Data{
		db: db,
	}

	d.initStmt()
	return d
}

func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}

// GetAllUsers ...
func (d Data) GetAllUsers(ctx context.Context) ([]userEntity.User, error) {
	var (
		user  userEntity.User
		users []userEntity.User
		err   error
	)

	rows, err := d.stmt[getAllUsers].QueryxContext(ctx)

	for rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return users, errors.Wrap(err, "[DATA][GetAllUsers] ")
		}
		users = append(users, user)
	}
	return users, err
}
