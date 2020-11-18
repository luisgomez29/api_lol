package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/api_lol/models"
	"github.com/luisgomez29/api_lol/repositories"
	"net/http"
	"strconv"
)

type CharacterController interface {
	GetAll(echo.Context) error
	FindById(echo.Context) error
	Create(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type characterController struct {
	characterRepository repositories.CharacterRepository
}

func NewCharacterController(characterRepository repositories.CharacterRepository) CharacterController {
	return &characterController{characterRepository}
}

func (ctl characterController) GetAll(c echo.Context) error {
	//s := auth.UserIDFromToken(c)
	//fmt.Printf("USR => %v\n", s)
	ch, err := ctl.characterRepository.All()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ch)
}

func (ctl *characterController) FindById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest
	}

	ch, err := ctl.characterRepository.FindById(uint32(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ch)
}

func (ctl *characterController) Create(c echo.Context) error {
	ch := new(models.Character)
	if err := c.Bind(ch); err != nil {
		return err
	}
	ch, err := ctl.characterRepository.Create(ch)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, ch)
}

func (ctl *characterController) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest
	}

	ch := new(models.Character)
	if err := c.Bind(ch); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	ch, err = ctl.characterRepository.Update(uint32(id), ch)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ch)
}

func (ctl *characterController) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.ErrBadRequest
	}

	rowsAffected, err := ctl.characterRepository.Delete(uint32(id))

	if err != nil {
		return err
	}
	return c.JSON(http.StatusNoContent, rowsAffected)
}
