package repository

import (
	"context"

	"blog-backend/internal/domain/entity"
	"github.com/google/uuid"
)

type ArticleRepository interface {
	Create(ctx context.Context, article *entity.Article) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Article, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Article, error)
	GetAll(ctx context.Context, limit, offset int) ([]*entity.Article, error)
	GetPublished(ctx context.Context, limit, offset int) ([]*entity.Article, error)
	GetByCategory(ctx context.Context, category string, limit, offset int) ([]*entity.Article, error)
	GetByTags(ctx context.Context, tags []string, limit, offset int) ([]*entity.Article, error)
	Update(ctx context.Context, article *entity.Article) error
	Delete(ctx context.Context, id uuid.UUID) error
	IncrementViewCount(ctx context.Context, id uuid.UUID) error
	Search(ctx context.Context, keyword string, limit, offset int) ([]*entity.Article, error)
	Count(ctx context.Context) (int64, error)
	CountByCategory(ctx context.Context, category string) (int64, error)
	GetCategories(ctx context.Context) ([]string, error)
	GetTags(ctx context.Context) ([]string, error)
} 