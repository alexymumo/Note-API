package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Note struct {
	NoteID      uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"size:255;not null" json:"description"`
	Completed   bool      `gorm:"size:255;not null" json:"completed"`
	CreateAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (n *Note) Validate() error {
	if n.Title == "" {
		return errors.New("title cannot be empty")
	}

	if n.Description == "" {
		return errors.New("description cannot be empty")
	}
	return nil

}

func (n *Note) SaveNote(db *gorm.DB) (*Note, error) {
	err := db.Debug().Model(&Note{}).Create(&n).Error
	if err != nil {
		return &Note{}, err
	}
	return n, nil

}

func (n *Note) DeleteNote(db *gorm.DB, id uint32) (int32, error) {
	db = db.Debug().Model(&Note{}).Where("note_id = ?", id).Take(&Note{}).Delete(&Note{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("not found")
		}
		return 0, db.Error
	}
	return int32(db.RowsAffected), nil

}

func (n *Note) UpdateNote(db *gorm.DB) {

}

func (n *Note) GetAllNotes(db *gorm.DB) (*[]Note, error) {
	notes := []Note{}
	err := db.Debug().Model(&Note{}).Limit(10).Find(&notes).Error
	if err != nil {
		return &[]Note{}, err
	}
	return &notes, nil

}
