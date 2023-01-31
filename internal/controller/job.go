package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/xtabs12/test_dans/internal/dto"
	"github.com/xtabs12/test_dans/internal/ucase"
	"github.com/xtabs12/test_dans/pkg/jwtx"
	"net/http"
	"strconv"
)

func NewJob(jobLogic ucase.JOB, group *echo.Group) {
	kk := jobImpl{
		jobLogic: jobLogic,
	}
	Router(kk, group)
}

type jobImpl struct {
	jobLogic ucase.JOB
}

func Router(c jobImpl, group *echo.Group) {
	group.GET("/job", c.GetJobList)
	group.GET("/job/:id", c.GetJobDetail)
}

func (i *jobImpl) GetJobDetail(e echo.Context) error {
	_, err := jwtx.NewJwtHmac256().ValidateTokenHmac256(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, buildResx(err.Error()))
	}
	id := e.Param("id")
	jobDetails, getJobDetailsErr := i.jobLogic.GetJobDetails(e.Request().Context(),
		dto.GetJobDetailParams{
			ID: id,
		})
	if getJobDetailsErr != nil {
		return e.JSON(http.StatusInternalServerError, buildResErr(getJobDetailsErr.Error()))
	}

	return e.JSON(http.StatusOK, buildResx(jobDetails))
}

func (i *jobImpl) GetJobList(e echo.Context) error {
	_, err := jwtx.NewJwtHmac256().ValidateTokenHmac256(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, buildResx(err.Error()))
	}
	page := e.QueryParam("page")
	description := e.QueryParam("description")
	location := e.QueryParam("location")
	isFullTime := e.QueryParam("full_time")
	pageX, _ := strconv.Atoi(page)
	var isFullTimeX bool
	if isFullTime == "true" {
		isFullTimeX = true
	}
	listJob, getListJobErr := i.jobLogic.GetJobList(e.Request().Context(),
		dto.GetJobListParams{
			Page:               pageX,
			SearchDescription:  description,
			SearchLocation:     location,
			SearchOnlyFullTime: isFullTimeX,
		})
	if getListJobErr != nil {
		return e.JSON(http.StatusInternalServerError, buildResx(getListJobErr.Error()))
	}

	return e.JSON(http.StatusOK, buildResx(listJob))
}
