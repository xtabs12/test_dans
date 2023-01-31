package jobPositionFromDansAPI

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type GetListResponseModel struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	URI         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	Company     string `json:"company"`
	CompanyURI  string `json:"company_url"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	CompanyLogo string `json:"company_logo"`
}

type GetListRequestModel struct {
	Page               int    `json:"page"`
	SearchDescription  string `json:"search_description"`
	SearchLocation     string `json:"search_location"`
	SearchOnlyFullTime bool   `json:"search_only_full_time"`
}

func GetList(r GetListRequestModel) ([]*GetListResponseModel, error) {
	endpoint := "http://dev3.dansmultipro.co.id/api/recruitment/positions.json"
	req, generateRequestErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if generateRequestErr != nil {
		return nil, generateRequestErr
	}
	// setup query filtering
	urlQuery := req.URL.Query()
	if r.Page != 0 {
		urlQuery.Add("page", strconv.Itoa(r.Page))
	}
	if r.SearchDescription != "" {
		urlQuery.Add("description", r.SearchDescription)
	}
	if r.SearchLocation != "" {
		urlQuery.Add("location", r.SearchLocation)
	}
	if r.SearchOnlyFullTime == true {
		urlQuery.Add("full_time", "true")
	}
	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{
		Timeout: 3600 * time.Second,
	}
	res, clientErr := client.Do(req)
	if clientErr != nil {
		return nil, clientErr
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("failed to close http body: ", err.Error())
		}
	}(res.Body)
	if res.StatusCode > 299 {
		return nil, errors.New(http.StatusText(res.StatusCode))
	}

	buf := new(strings.Builder)
	_, bufferCopyErr := io.Copy(buf, res.Body)
	if bufferCopyErr != nil {
		return nil, clientErr
	}

	var responses []*GetListResponseModel
	if err := json.Unmarshal([]byte(buf.String()), &responses); err != nil {
		return nil, err
	}

	return responses, nil
}

type GetDetailResponseModel struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	URI         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	Company     string `json:"company"`
	CompanyURI  string `json:"company_url"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	CompanyLogo string `json:"company_logo"`
}

type GetDetailRequestModel struct {
	ID string `json:"id"`
}

func GetDetails(r GetDetailRequestModel) (*GetDetailResponseModel, error) {
	endpoint := "http://dev3.dansmultipro.co.id/api/recruitment/positions/" + r.ID
	req, generateRequestErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if generateRequestErr != nil {
		return nil, generateRequestErr
	}
	client := http.Client{
		Timeout: 3600 * time.Second,
	}
	res, clientErr := client.Do(req)
	if clientErr != nil {
		return nil, clientErr
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("failed to close http body: ", err.Error())
		}
	}(res.Body)
	if res.StatusCode > 299 {
		return nil, errors.New(http.StatusText(res.StatusCode))
	}

	buf := new(strings.Builder)
	_, bufferCopyErr := io.Copy(buf, res.Body)
	if bufferCopyErr != nil {
		return nil, clientErr
	}

	var responses GetDetailResponseModel
	if err := json.Unmarshal([]byte(buf.String()), &responses); err != nil {
		return nil, err
	}

	return &responses, nil
}
