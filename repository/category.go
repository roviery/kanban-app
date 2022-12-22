package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	results := []entity.Category{}
	rows, err := r.db.Table("categories").Where("user_id = ?", id).Select("*").Rows()

	if err != nil {
		return []entity.Category{}, err
	}

	defer rows.Close()

	for rows.Next() {
		r.db.ScanRows(rows, &results)
	}

	return results, nil
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	result := r.db.Create(&category)

	if result.Error != nil {
		return 0, result.Error
	}

	return category.ID, nil
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	result := r.db.Create(&categories)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	result := entity.Category{}
	response := r.db.Where("id = ?", id).Model(&entity.Category{}).First(&result)

	if response.Error != nil {
		return entity.Category{}, response.Error
	}

	if response.RowsAffected == 0 {
		return entity.Category{}, nil
	}

	return result, nil
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	result := r.db.Model(&category).Updates(category)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	result := r.db.Delete(&entity.Category{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
