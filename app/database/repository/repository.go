package repository

import (
	"github.com/LOCNNIL/golang-rinha-api/app/database/repository/filters"
	"github.com/LOCNNIL/golang-rinha-api/app/models"
	"gorm.io/gorm"
)

type Repository struct {
	DatabaseConnection *gorm.DB
}

func (repo *Repository) GetPeopleCount() (int64, error) {
	var count int64
	result := repo.DatabaseConnection.Model(&models.People{}).Count(&count)
	return count, result.Error
}

func (repo *Repository) UpsertPerson(person *models.People) error {
	return repo.DatabaseConnection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(person).Error; err != nil {
			return err
		}
		return nil
	})
}

func (repo *Repository) FindPerson(person *models.People) *gorm.DB {
	tx := repo.DatabaseConnection.
		Take(&person)
	return tx
}

func (repo *Repository) FindAllPeople(
	people *[]models.People,
	filter *filters.FindAllPeopleFilter,
) *gorm.DB {

	model := repo.DatabaseConnection.Model(&models.People{})

	if filter.Name != "" {
		model = model.Where("nome LIKE ?", "%"+filter.Name+"%")
	}

	if filter.Nickname != "" {
		model = model.Where("apelido LIKE ?", "%"+filter.Nickname+"%")
	}

	if len(filter.Stack) > 0 {
		for _, v := range filter.Stack {
			model = model.Where("stack LIKE ?", "%"+v+"%")
		}
	}

	tx := model.Find(people)
	return tx
}

func (repo *Repository) FindPeople(
	people *[]models.People,
	search_term string,
) *gorm.DB {

	model := repo.DatabaseConnection.Model(&models.People{})

	model.Where("nickname ILIKE ? OR name ILIKE ? OR stack ILIKE ?", "%"+search_term+"%", "%"+search_term+"%", "%"+search_term+"%")

	tx := model.Find(people)
	return tx
}
