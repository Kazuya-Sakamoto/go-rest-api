package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IDiaryController interface {
	GetAllDiaries(c echo.Context) error
	GetDiaryById(c echo.Context) error
	CreateDiary(c echo.Context) error
	UpdateDiary(c echo.Context) error
	DeleteDiary(c echo.Context) error
}

type diaryController struct {
	du usecase.IDiaryUsecase
}

func NewDiaryController(du usecase.IDiaryUsecase) IDiaryController {
	return &diaryController{du}
}

func (dc *diaryController) GetAllDiaries(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	diariesResponse, err := dc.du.GetAllDiaries(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, diariesResponse)
}

func (dc *diaryController) GetDiaryById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("diaryId")
	diaryId, _ := strconv.Atoi(id)
	diaryResponse, err := dc.du.GetDiaryById(uint(userId.(float64)), uint(diaryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, diaryResponse)
}

func (dc *diaryController) CreateDiary(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	diary := model.Diary{}
	if err := c.Bind(&diary); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	diary.UserId = uint(userId.(float64))
	diaryResponse, err := dc.du.CreateDiary(diary)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, diaryResponse)
}

func (dc *diaryController) UpdateDiary(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("diaryId")
	diaryId, _ := strconv.Atoi(id)

	diary := model.Diary{}
	if err := c.Bind(&diary); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	diaryResponse, err := dc.du.UpdateDiary(diary, uint(userId.(float64)), uint(diaryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, diaryResponse)
}

func (dc *diaryController) DeleteDiary(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("diaryId")
	diaryId, _ := strconv.Atoi(id)

	err := dc.du.DeleteDiary(uint(userId.(float64)), uint(diaryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
