package external

import "mime/multipart"

type External interface {
	UploadAttachment(file *multipart.FileHeader) (string, error)
}
