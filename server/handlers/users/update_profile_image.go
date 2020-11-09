package users

import (
	"database/sql"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"net/http"
	"os"
	"simple-service/server/errs"
	"simple-service/server/middlewares"
	"strings"

	"github.com/nfnt/resize"

	"github.com/labstack/echo/v4"
)

const profileImage = "profile_image"

// UpdateProfileImage godoc
// @Security bearerAuth
// @Summary Update profile image
// @Tags user
// @Consume multipart/form-data
// @Param profile_image formData file true "select image"
// @Description Update user's profile image
// @Accept  multipart/form-data
// @Produce json
// @Success 200 {} http.StatusOk
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /user/image [post]
func (h *Handler) UpdateProfileImage(c echo.Context) error {

	form, err := c.MultipartForm()
	if err != nil {
		h.log.WithError(err).Error("failed to get data from the multipart form")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	userID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed ot get user ID from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	user, err := h.usersDB.GetByID(userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(errs.UserDoesntExistErr.ToError()).Error("user doesn't exist")
			return c.JSON(http.StatusBadRequest, errs.UserDoesntExistErr)
		default:
			h.log.WithError(err).Error("failed to get user from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	if len(form.File) == 0 {
		h.log.WithError(err).Error("multipart form is empty")
		return c.JSON(http.StatusBadRequest, errs.NoDataInFormErr)
	}

	file := form.File["profile_image"][0]
	photo, err := file.Open()
	defer photo.Close()

	urlByElements := strings.Split(user.ProfileImageUrl, "/")

	fileName := urlByElements[len(urlByElements)-1]

	if user.ProfileImageUrl != "" {
		err = h.s3.DropFile(fileName)
		if err != nil {
			h.log.WithError(err).Error("failed to delete an image from s3")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	im, _, err := image.Decode(photo)
	if err != nil {
		h.log.WithError(err).Error("failed to decode image")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	bound := im.Bounds()

	//Checking the image size and compressing it, if size are not valid
	if bound.Max.X > 160 && bound.Max.Y > 160 {
		compressedImage, err := getCompressedImage(im, user.ID)
		if err != nil {
			h.log.WithError(err).Error("failed to get compressed image")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}

		info, err := compressedImage.Stat()
		if err != nil {
			h.log.WithError(err).Error("failed to get file info")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
		out, err := os.Open(compressedImage.Name())
		if err != nil {
			h.log.WithError(err).Error("failed to open file for uploading")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
		defer out.Close()

		//outputFile, err := os.Create(fmt.Sprintf("%s.jpg", user.ID))
		//defer outputFile.Close()
		//if err != nil {
		//	h.log.WithError(err).Error("failed to create file for uploading image")
		//	return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		//}
		//
		//err = jpeg.Encode(outputFile, newImage, &jpeg.Options{
		//	Quality: 80,
		//})
		//if err != nil {
		//	h.log.WithError(err).Error("failed to encode compressed image")
		//	return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		//}
		//
		//info, err := outputFile.Stat()
		//out, err := os.Open(outputFile.Name())

		user.ProfileImageUrl, err = h.s3.PutObject(out, info.Size(), user.ID)
		if err != nil {
			h.log.WithError(err).Error("failed to put new profile image into bucket")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}

		err = os.Remove(compressedImage.Name())
		if err != nil {
			h.log.WithError(err).Error("failed to remove file after uploading")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	} else {
		user.ProfileImageUrl, err = h.s3.PutObject(photo, file.Size, user.ID)
		if err != nil {
			h.log.WithError(err).Error("failed to put new profile image into bucket")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	err = h.usersDB.Update(user)
	if err != nil {
		h.log.WithError(err).Error("failed to update user's profile photo url in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}

func getCompressedImage(im image.Image, fileName string) (*os.File, error) {
	newImage := resize.Resize(160, 160, im, resize.Lanczos3)
	outputFile, err := os.Create(fmt.Sprintf("%s.jpg", fileName))
	defer outputFile.Close()
	if err != nil {
		return nil, err
	}

	err = jpeg.Encode(outputFile, newImage, &jpeg.Options{
		Quality: 80,
	})

	if err != nil {
		return nil, err
	}

	out, err := os.Open(outputFile.Name())
	if err != nil {
		return nil, err
	}

	return out, nil
}
