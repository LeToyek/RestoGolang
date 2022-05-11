package repository

import (
	"database/sql"
	"resto/entities"
)

type FoodsRepository interface {
	AddFood(food entities.Food) error
	GetAllfoods() ([]entities.Food, error)
	GetFoodsByCategory(category string) ([]entities.Food, error)
}

func (r *Repository) AddFood(food entities.Food) error {
	_, err := r.DB.Exec(
		queryInputFood,
		food.Food_id,
		food.Name,
		food.Price,
		food.Image,
		food.Category,
		food.Created_at,
		food.Updated_at,
	)
	return err
}

func (r *Repository) GetAllfoods() ([]entities.Food, error) {
	var foods []entities.Food

	rows, err := r.DB.Query(queryGetAllFoods)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	for rows.Next() {
		var food = entities.Food{}
		err := rows.Scan(
			&food.Food_id,
			&food.Name,
			&food.Price,
			&food.Image,
			&food.Category,
			&food.Created_at,
			&food.Updated_at,
		)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}
	return foods, err
}
func (r *Repository) GetFoodsByCategory(category string) ([]entities.Food, error) {
	var foods []entities.Food

	rows, err := r.DB.Query(queryGetFoodsByCategory, category)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	for rows.Next() {
		var food = entities.Food{}
		err := rows.Scan(
			&food.Food_id,
			&food.Name,
			&food.Price,
			&food.Image,
			&food.Category,
			&food.Created_at,
			&food.Updated_at,
		)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}
	return foods, err
}
