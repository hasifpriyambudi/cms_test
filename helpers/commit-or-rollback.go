package helpers

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	// Cek Jika Error
	if err != nil {
		errRolback := tx.Rollback()
		PanicError(errRolback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicError(errCommit)
	}
}
