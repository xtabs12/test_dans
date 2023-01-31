package ucase

import (
	"context"
	"github.com/xtabs12/test_dans/internal/dto"
	"github.com/xtabs12/test_dans/internal/service/jobPositionFromDansAPI"

	"log"
)

type jobImpl struct {
}

func NewJob() JOB {
	return &jobImpl{}
}

type JOB interface {
	GetJobDetails(ctx context.Context, params dto.GetJobDetailParams) (*dto.Job, error)
	GetJobList(ctx context.Context, params dto.GetJobListParams) ([]*dto.Job, error)
}

func (i *jobImpl) GetJobDetails(ctx context.Context, params dto.GetJobDetailParams) (*dto.Job, error) {

	jobPosition, getJobPositionErr := jobPositionFromDansAPI.GetDetails(jobPositionFromDansAPI.GetDetailRequestModel{
		ID: params.ID,
	})
	if getJobPositionErr != nil {
		log.Println(getJobPositionErr)
		return nil, getJobPositionErr
	}
	dtoJobPosition := dto.NewJob()
	dtoJobPosition.ID = jobPosition.ID
	dtoJobPosition.Type = jobPosition.Type
	dtoJobPosition.URI = jobPosition.URI
	dtoJobPosition.CreatedAt = jobPosition.CreatedAt
	dtoJobPosition.Company = jobPosition.Company
	dtoJobPosition.CompanyURI = jobPosition.CompanyURI
	dtoJobPosition.Location = jobPosition.Location
	dtoJobPosition.Title = jobPosition.Title
	dtoJobPosition.Description = jobPosition.Description
	dtoJobPosition.HowToApply = jobPosition.HowToApply
	dtoJobPosition.CompanyLogo = jobPosition.CompanyLogo

	return dtoJobPosition, nil
}

func (i *jobImpl) GetJobList(ctx context.Context, params dto.GetJobListParams) ([]*dto.Job, error) {
	listJobPosition, getListJobPositionErr := jobPositionFromDansAPI.GetList(jobPositionFromDansAPI.GetListRequestModel{
		Page:               params.Page,
		SearchDescription:  params.SearchDescription,
		SearchLocation:     params.SearchLocation,
		SearchOnlyFullTime: params.SearchOnlyFullTime,
	})
	if getListJobPositionErr != nil {
		log.Println(getListJobPositionErr)
		return nil, getListJobPositionErr
	}
	var dtoListJobPosition []*dto.Job
	for _, jobPosition := range listJobPosition {
		// when server returned list of null the data will be skipped
		if jobPosition == nil {
			continue
		}
		dtoJobPosition := dto.NewJob()
		dtoJobPosition.ID = jobPosition.ID
		dtoJobPosition.Type = jobPosition.Type
		dtoJobPosition.URI = jobPosition.URI
		dtoJobPosition.CreatedAt = jobPosition.CreatedAt
		dtoJobPosition.Company = jobPosition.Company
		dtoJobPosition.CompanyURI = jobPosition.CompanyURI
		dtoJobPosition.Location = jobPosition.Location
		dtoJobPosition.Title = jobPosition.Title
		dtoJobPosition.Description = jobPosition.Description
		dtoJobPosition.HowToApply = jobPosition.HowToApply
		dtoJobPosition.CompanyLogo = jobPosition.CompanyLogo
		dtoListJobPosition = append(dtoListJobPosition, dtoJobPosition)
	}

	return dtoListJobPosition, nil
}
