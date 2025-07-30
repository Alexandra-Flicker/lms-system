package entity

import "lms_system/internal/domain/common"

type UserContext struct {
	UserID   string          `json:"UserID"`
	Role     common.UserRole `json:"role"`
	Username string          `json:"username"`
}
