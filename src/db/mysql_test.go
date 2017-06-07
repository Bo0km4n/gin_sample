package db

import (
	"github.com/coopernurse/gorp"
    "testing"
)

var dsn = "root:@/test"

func TestMySQL(t *testing.T) {
	dbmap := NewDbMap(dsn)
    CreateUserTable(dbmap)
    defer dropAndClose(dbmap)
    rows, _ := dbmap.Select(User{}, "SELECT * FROM users")
    if len(rows) != 0 {
        t.Errorf("Expected 0 invoice rows, got %d", len(rows))
    }
}

func dropAndClose(dbmap *gorp.DbMap) {
	dbmap.DropTablesIfExists()
    dbmap.Db.Close()
}
