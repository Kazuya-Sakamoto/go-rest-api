package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IDiaryCommentController interface {
	GetDiaryCommentsByDiaryIDAndUserID(c echo.Context) error
	CreateDiaryComment(c echo.Context) error
	DeleteDiaryComment(c echo.Context) error
}
type diaryCommentController struct {
	dcu usecase.IDiaryCommentUsecase
}

type RequestBodyForGetDiaryCommentsByDiaryIDAndUserID struct {
	DiaryID uint `json:"diary_id"`
}
type RequestBodyDeleteDiaryComment struct {
	ID uint `json:"id" binding:"required"`
}

func NewDiaryCommentController(dcu usecase.IDiaryCommentUsecase) IDiaryCommentController {
	return &diaryCommentController{dcu}
}

func (dcc *diaryCommentController) GetDiaryCommentsByDiaryIDAndUserID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// リクエストボディの JSON データを DiaryIDRequest 構造体にデコード
	var requestBody RequestBodyForGetDiaryCommentsByDiaryIDAndUserID
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	diaryCommentResponse, err := dcc.dcu.GetDiaryCommentsByDiaryIDAndUserID(requestBody.DiaryID, uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, diaryCommentResponse)
}

func (dcc *diaryCommentController) CreateDiaryComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	diaryComment := model.DiaryComment{}
	if err := c.Bind(&diaryComment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// fmt.Println(diaryComment, "controller/diary_comment_controller.go >diaryComment")
	diaryComment.UserId = uint(userId.(float64))
	diaryCommentResponse, err := dcc.dcu.CreateDiaryComment(diaryComment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, diaryCommentResponse)
}

func (dcc *diaryCommentController) DeleteDiaryComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	var request RequestBodyDeleteDiaryComment
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := dcc.dcu.DeleteDiaryComment(uint(userId.(float64)), request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
