package middleware

import (
	"context"
	"fmt"
	"lectronic/src/libs"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func AuthCloudUploadFile() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := context.Background()
			cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))

			r.Body = http.MaxBytesReader(w, r.Body, 2*1024*1024)
			file, fileHeader, err := r.FormFile("image")
			if err != nil {
				if err == http.ErrMissingFile {
					i, _ := cld.Image("v1676028039/gorental/default_image.jpg")
					urlDefault, _ := i.String()
					ctx := context.WithValue(r.Context(), "imageName", urlDefault)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				libs.GetResponse(err.Error(), 401, true).Send(w)
				return
			}

			defer file.Close()

			if fileHeader.Header.Get("Content-Type") != "image/jpeg" && fileHeader.Header.Get("Content-Type") != "image/png" && fileHeader.Header.Get("Content-Type") != "image/jpg" {
				libs.GetResponse("file format is not allowed. Please upload a JPEG, JPG or PNG image", 401, true).Send(w)
				return
			}

			imgName := fmt.Sprintf("%d", time.Now().UnixNano())

			upload, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "lectronic", PublicID: imgName})
			if err != nil {
				libs.GetResponse(err.Error(), 400, true).Send(w)
			}

			cntx := context.WithValue(r.Context(), "imageName", upload.SecureURL)
			next.ServeHTTP(w, r.WithContext(cntx))

		})
	}
}
