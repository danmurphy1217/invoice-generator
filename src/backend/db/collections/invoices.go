package collections

import (
	"time"

	"gorm.io/gorm"
)

// Invoice defines our invoice collection,
// which is a postgres table that stores
// data on the invoices we have generated
// Normally, this would look something like:
// type Invoice struct {
//     ID        string   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
//     BucketName      string `gorm:"not null"`
// 	KeyName      string `gorm:"not null"`
//     CreatedAt time.Time
//     UpdatedAt time.Time
// }
// but i did not have time to integrate with S3 for this
type Invoice struct {
    ID        string   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    PDFBytes      []byte `gorm:"not null"`
	Title      string `gorm:"unique;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

// set inserted and updated at before saving
func (i *Invoice) BeforeSave(tx *gorm.DB) (err error) {
    currentTimeUTC := time.Now().UTC()
    if i.CreatedAt.IsZero() {
        i.CreatedAt = currentTimeUTC
    }
    i.UpdatedAt = currentTimeUTC
    return nil
}
