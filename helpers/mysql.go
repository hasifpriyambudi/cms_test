package helpers

import (
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func MysqlError(err error) error {
	if isDuplicateEntry(err) {
		errMsg := extractDuplicateKeyField(err)
		errString := errors.New("duplicate entry for " + errMsg + " field")
		return errString
	}

	PanicError(err)
	return nil
}

func isDuplicateEntry(err error) bool {
	mysqlError, ok := err.(*mysql.MySQLError)
	return ok && mysqlError.Number == 1062
}

func extractDuplicateKeyField(err error) string {
	mysqlError, _ := err.(*mysql.MySQLError)

	start := strings.Index(mysqlError.Message, "'")
	end := strings.LastIndex(mysqlError.Message, "'")
	return mysqlError.Message[start+1 : end]
}
