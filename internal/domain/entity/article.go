package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Slug        string         `gorm:"uniqueIndex;not null" json:"slug"`
	Content     string         `gorm:"type:text;not null" json:"content"`
	Summary     string         `gorm:"type:text" json:"summary"`
	Category    string         `gorm:"size:100" json:"category"`
	Thumbnail   string         `gorm:"size:500" json:"thumbnail"`
	Tags        []string       `gorm:"type:text[]" json:"tags"`
	ViewCount   int64          `gorm:"default:0" json:"view_count"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	PublishedAt *time.Time     `json:"published_at,omitempty"`
	IsPublished bool           `gorm:"default:false" json:"is_published"`
}

func (Article) TableName() string {
	return "articles"
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}

func (a *Article) Publish() {
	a.IsPublished = true
	now := time.Now()
	a.PublishedAt = &now
}

func (a *Article) Unpublish() {
	a.IsPublished = false
	a.PublishedAt = nil
}

func (a *Article) IncrementViewCount() {
	a.ViewCount++
} 