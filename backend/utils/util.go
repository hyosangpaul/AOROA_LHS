package utils

import "backend/models"

// IsValidStatus 함수: 주어진 Status 값이 올바른지 확인
func IsValidStatus(status models.IssueStatus) bool {
	switch status {
	case models.PENDING, models.IN_PROGRESS, models.COMPLETED, models.CANCELLED:
		return true
	default:
		return false
	}
}