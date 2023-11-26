package external

import "mime/multipart"

// Implement upload attachment with GCS
type External interface {
	UploadAttachment(file *multipart.FileHeader) (string, error)
}
