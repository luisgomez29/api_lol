package repositories

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/api_lol/models"
	"github.com/luisgomez29/api_lol/utils"
	"gorm.io/gorm"
)

type CharacterRepository interface {
	All() ([]*models.Character, error)
	FindById(uint32) (*models.Character, error)
	Create(*models.Character) (*models.Character, error)
	Update(uint32, *models.Character) (*models.Character, error)
	Delete(uint32) (int64, error)
}

type CharacterDB struct {
	conn *gorm.DB
}

func NewCharacterRepository(db *gorm.DB) CharacterRepository {
	return &CharacterDB{db}
}

func (db *CharacterDB) All() ([]*models.Character, error) {
	var c []*models.Character
	db.conn.Select(utils.Fields(&models.Character{})).Limit(100).Find(&c)
	return c, nil
}

func (db *CharacterDB) FindById(id uint32) (*models.Character, error) {
	c := new(models.Character)
	fields := utils.Fields(c)
	if err := db.conn.Select(fields).Take(c, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, echo.ErrNotFound
	}
	return c, nil
}

func (db *CharacterDB) Create(ch *models.Character) (*models.Character, error) {
	err := db.conn.Create(ch).Error
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	return ch, nil
}

func (db *CharacterDB) Update(id uint32, ch *models.Character) (*models.Character, error) {
	c, err := db.FindById(id)
	if err != nil {
		return nil, err
	}
	c.Name = ch.Name
	c.Description = ch.Description
	c.Position = ch.Position

	if err := db.conn.Save(c).Error; err != nil {
		return nil, echo.ErrInternalServerError
	}
	return c, nil
}

func (db *CharacterDB) Delete(id uint32) (int64, error) {
	rs := db.conn.Select("id").Take(&models.Character{}, id).Delete(&models.Character{})
	if errors.Is(rs.Error, gorm.ErrRecordNotFound) {
		return 0, echo.ErrNotFound
	}
	return rs.RowsAffected, nil
}
