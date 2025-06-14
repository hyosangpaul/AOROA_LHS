package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"backend/models"
	"backend/utils"
)

// 데이터 저장 (임시 DB)
var issues = []models.Issue{}

// POST
func CreateIssue(w http.ResponseWriter, r *http.Request) {
	var newIssue models.Issue
	err := json.NewDecoder(r.Body).Decode(&newIssue)
	if err != nil || newIssue.Title == "" || newIssue.Description == "" {
		http.Error(w, `{"error": "필수값이 없습니다.", "code": 400}`, http.StatusBadRequest)
		return
	}

	// 담당자 검증
	if newIssue.User != nil {
		if !utils.IsValidUser(newIssue.User.ID, users) {
			http.Error(w, `{"error": "유저를 찾을 수 없습니다.", "code": 404}`, http.StatusNotFound)
			return
		}
		newIssue.Status = models.IN_PROGRESS
	} else {
		newIssue.Status = models.PENDING
	}

	// ID 및 시간 설정
	newIssue.ID = uint(len(issues) + 1)
	newIssue.CreatedAt = time.Now()
	newIssue.UpdatedAt = time.Now()

	issues = append(issues, newIssue)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newIssue)
}

// GET
func GetIssues(w http.ResponseWriter, r *http.Request) {
	statusFilter := r.URL.Query().Get("status")
	var filteredIssues []models.Issue

	for _, issue := range issues {
		if statusFilter == "" || string(issue.Status) == statusFilter {
			filteredIssues = append(filteredIssues, issue)
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"issues": filteredIssues})
}

// GET (DETAIL)
func GetIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, issue := range issues {
		if string(issue.ID) == params["id"] {
			json.NewEncoder(w).Encode(issue)
			return
		}
	}
	http.Error(w, `{"error": "이슈를 찾을 수 없습니다.", "code": 404}`, http.StatusNotFound)
}

// PATCH (UPDATE)
func UpdateIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedData models.Issue
	json.NewDecoder(r.Body).Decode(&updatedData)

	for i, issue := range issues {
		if string(issue.ID) == params["id"] {
			// COMPLETED 또는 CANCELLED 상태에서는 수정 불가
			if issue.Status == models.COMPLETED || issue.Status == models.CANCELLED {
				http.Error(w, `{"error": "COMPLETED 또는 CANCELLED 상태에서는 수정이 불가합니다.", "code": 403}`, http.StatusForbidden)
				return
			}

			// 부분 업데이트 처리 (명시되지 않은 필드는 유지)
			if updatedData.Title != "" {
				issue.Title = updatedData.Title
			}
			if updatedData.Description != "" {
				issue.Description = updatedData.Description
			}
			if updatedData.User != nil {
				if !utils.IsValidUser(updatedData.User.ID, users) {
					http.Error(w, `{"error": "유저를 찾을 수 없습니다.", "code": 404}`, http.StatusNotFound)
					return
				}
				issue.User = updatedData.User
			}

			// 상태 변경 규칙 적용
			if updatedData.Status != "" {
				// 담당자가 없는 상태에서 IN_PROGRESS 또는 COMPLETED 변경 불가
				if (updatedData.Status == models.IN_PROGRESS || updatedData.Status == models.COMPLETED) && issue.User == nil {
					http.Error(w, `{"error": "담당자가 없는 상태에서 IN_PROGRESS 또는 COMPLETED 변경이 불가합니다.", "code": 400}`, http.StatusBadRequest)
					return
				}
				issue.Status = updatedData.Status
			} else {
				// 상태 미지정 시 기본값 적용
				issue.Status = models.IN_PROGRESS
			}

			// 담당자 제거 시 상태 변경
			if updatedData.User == nil {
				issue.Status = models.PENDING
			}

			issue.UpdatedAt = time.Now()
			issues[i] = issue
			json.NewEncoder(w).Encode(issue)
			return
		}
	}
	http.Error(w, `{"error": "이슈를 찾을 수 없습니다.", "code": 404}`, http.StatusNotFound)
}

// DELETE
func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, issue := range issues {
		if string(issue.ID) == params["id"] {
			issues = append(issues[:i], issues[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, `{"error": "이슈를 찾을 수 없습니다.", "code": 404}`, http.StatusNotFound)
}