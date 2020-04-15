package datastore

import (
	"database/sql"
	"time"
)

// Client is a struct used to represent records in clients table.
type Client struct {
	ID       string `db:"id"`
	Token    string `db:"token"`
	ExpireAt int64  `db:"expire_at"`
}

// TODO: remove this.
// InsertClient create and insert a new client into clients table.
func (pg *Postgres) InsertClient(id string, token string, expireAt int64) (*Client, error) {
	client := Client{ID: id, Token: token, ExpireAt: expireAt}
	stmt, err := pg.PrepareNamed(
		"INSERT INTO clients(id, token, expire_at) VALUES(:id, :token, :expire_at) " +
			"ON CONFLICT (id) DO UPDATE SET token = :token, expire_at = :expire_at " +
			"RETURNING *")
	if err != nil {
		return nil, err
	}

	var savedClient Client
	err = stmt.Get(&savedClient, client)
	return &savedClient, err
}

// InsertClientToken creates and inserts a new client into clients table if not
// exists, and inserts the token into tokens table.
func (pg *Postgres) InsertClientToken(id string, token string, expireAt int64) error {
	tx, err := pg.DB.Beginx()
	if err != nil {
		return err
	}
	defer pg.RollbackTx(tx)

	_, err = tx.Exec(`INSERT INTO clients (id) VALUES ($1) ON CONFLICT DO NOTHING;`, id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO tokens (id, expire_at, client_id) VALUES ($1, $2, $3);`,
		token, expireAt, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

// GetClientID queries the tokens table using token, return the clientID if
// the token is valid, otherwise, return empty string.
func (pg *Postgres) GetClientID(token string) (string, error) {
	var clientID string
	err := pg.Get(&clientID, "SELECT client_id FROM tokens WHERE id = $1 AND expire_at > $2",
		token, time.Now().Unix())
	if err == sql.ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return clientID, nil
}
