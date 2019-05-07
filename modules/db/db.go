package db

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" // postgresql driver
)

// ORM control database transaction
type ORM struct {
	orm.Ormer
}

// NewOrm create new orm instance
func NewOrm() ORM {
	return ORM{
		Ormer: orm.NewOrm(),
	}
}

// RegisterModel register struct to database table
func RegisterModel(models ...interface{}) {
	orm.RegisterModel(models...)
}

// StartDB start database connection
func StartDB(dbAlias, dbDriver, dbString string, dbForce, dbVerbose bool) {
	orm.RegisterDataBase(dbAlias, dbDriver, dbString)

	// Database alias.
	name := "default"

	// Error.
	err := orm.RunSyncdb(name, dbForce, dbVerbose)
	if err != nil {
		logs.Error("DB Error: ", err)
	}
}
