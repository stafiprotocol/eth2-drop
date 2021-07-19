package dao_user

import (
	"drop/pkg/db"
)

func AutoMigrate(db *db.WrapDb) error {
	return db.Set("gorm:table_options", "ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8").
		AutoMigrate(Snapshot{}, DropFlow{}, DropLedger{}, MetaData{})
}
