package stat

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Stat struct {
	gorm.Model
	LinkId uint           `json:"links_id"`
	Clicks int            `json:"clicks"`
	Data   datatypes.Date `json:"date"`
}
