package utils

import (
	"encoding/json"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"github.com/kapcomperu/go-auth0-utils/auth"
	"github.com/kapcomperu/go-common-utils/models/base"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

func GetSortDirection(sortBy string) int {

	if strings.ToLower(sortBy) == "asc" {
		return 1
	} else {
		return -1
	}

}
func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

func ToPageable(paginatedData mongopagination.PaginatedData) *base.Pageable {
	var paging *base.Pageable

	paging.Total = paginatedData.Pagination.Total
	paging.Page = paginatedData.Pagination.Page
	paging.PerPage = paginatedData.Pagination.PerPage
	paging.Prev = paginatedData.Pagination.Prev
	paging.Next = paginatedData.Pagination.Next
	paging.TotalPage = paginatedData.Pagination.TotalPage
	return paging
}

func PrepareResponseList(models []interface{}, pageable *base.Pageable, err error) base.ResponseDataListPageable {
	var response base.ResponseDataListPageable
	if err != nil {
		response = base.ResponseDataListPageable{Status: false, Messages: []string{err.Error()}, StatusCode: http.StatusInternalServerError}
	} else {
		var data []interface{}
		if models != nil {
			for _, element := range models {
				data = append(data, element)
			}
		} else {
			data = make([]interface{}, 0)
		}

		response = base.ResponseDataListPageable{Data: data, Status: true, StatusCode: http.StatusOK, Paging: pageable}
	}
	return response
}

func PrepareResponseOneRecord(model interface{}, err error) base.ResponseDataList {
	var response base.ResponseDataList
	if err != nil {
		response = base.ResponseDataList{Status: false, Messages: []string{err.Error()}, StatusCode: http.StatusInternalServerError}
	} else {
		var data []interface{}
		data = make([]interface{}, 0)
		if model != nil {
			data = append(data, model)
			response = base.ResponseDataList{Data: data, Status: true, StatusCode: http.StatusOK}
		} else {
			response = base.ResponseDataList{Data: data, Status: false, StatusCode: http.StatusNotFound}
		}
	}
	return response
}

func ResponseInsufficientScope(w http.ResponseWriter, statusCode int) {
	message := "Insufficient scope."
	response := auth.Response{Message: message}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

