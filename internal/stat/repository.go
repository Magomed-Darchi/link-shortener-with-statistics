package stat

import (
	"api-main/pkg/db"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.Db.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Data:   currentDate,
		})

	} else {
		stat.Clicks += 1
		repo.Db.Save(&stat)

	}

}

func (repo *StatRepository) GetStat(by string, from, to time.Time) []GetStatResponse {
	var stats []GetStatResponse
	var selectQuery string
	switch by {
	case GroupByDay:
		selectQuery = "to char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case GroupByMonth:
		selectQuery = "to char(date, 'YYYY-MM') as period, sum(clicks)"

	}
	query := repo.DB.Table("stats").
		Select(selectQuery).
		Session(&gorm.Session{})
	if true {
		query.Where("count > 10")
	}

	query.Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)
	return stats
}
