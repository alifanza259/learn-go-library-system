package util

import (
	"strings"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
)

func IsAdminAuthorized(admin db.Admin, permission string) bool {
	adminPermissions := strings.Split(admin.Permission, "|")
	for _, adminPermission := range adminPermissions {
		if adminPermission == permission || adminPermission == "super" {
			return true
		}
	}
	return false
}
