package models

//User represents users table in database
type Destination struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(50)" json:"title"`
	Description string `gorm:"uniqueIndex;type:varchar(255)" json:"description"`
	Keyword     string `gorm:"type:varchar(25)" json:"keyword"`
	Images      string `gorm:"type:varchar(255)" json:"images"`
	Address     string `gorm:"type:varchar(255)" json:"address"`
	Views       uint64 `gorm:"type:int" json:"views"`

	//jika si user mengambil ke destinasi
	UserID uint64 `gorm:"not null" json:"-"`
	User   User   `gorm:"foreignkey:UserID:constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
