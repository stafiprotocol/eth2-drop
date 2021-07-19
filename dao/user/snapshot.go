package dao_user

import (
	"drop/pkg/db"

	"gorm.io/gorm"
)

type Snapshot struct {
	db.BaseModel
	UserAddress string `gorm:"type:varchar(42);not null;default:'';column:user_address;uniqueIndex:uni_idx_user_round"`
	Round       int64  `gorm:"type:bigint;unsigned;not null;column:round;uniqueIndex:uni_idx_user_round"`
	DropAmount  string `gorm:"type:varchar(30);not null;default:'0';column:drop_amount"`
	Claimed     int8   `gorm:"type:tinyint(1);not null;default:0;column:claimed"`
}

func UpOrInSnapshot(db *db.WrapDb, c *Snapshot) error {
	return db.Create(c).Error
}

func GetSnapshotByRoundAndUser(db *db.WrapDb, round int64, user string) (banker *Snapshot, err error) {
	banker = &Snapshot{}
	err = db.Take(banker, "user_address = ? and round = ?", user, round).Error
	if err == gorm.ErrRecordNotFound {
		banker.DropAmount = "0"
	}
	return
}

func GetSnapshotLastRound(db *db.WrapDb) (r int64, err error) {
	type Round struct {
		Round int64
	}
	round := Round{}
	err = db.Raw("select round from snapshots order by round desc limit 1").Scan(&round).Error
	return round.Round, err
}

func GetSnapshotListByRound(db *db.WrapDb, round int64) (claimList []*Snapshot, err error) {
	err = db.Order("id asc").Find(&claimList, "round = ?", round).Error
	return
}

//only update claimd
func UpdateSnapshot(db *db.WrapDb, stat *Snapshot) (err error) {
	return db.Model(stat).Update("claimed", stat.Claimed).Error
}
