package middleware

import (
	"errors"

	"github.com/anandureghu/go-blog/internal/model"
)

func ValidateBlog(blog model.Blog) error {

	if blog.Id < 0 {
		return errors.New("invalid id")
	}

	if blog.Description == "" || len(blog.Description) <= 0 {
		return errors.New("invalid description")
	}

	if blog.Title == "" || len(blog.Title) <= 0 {
		return errors.New("invalid title")
	}

	if blog.Cover == "" || len(blog.Cover) <= 0 {
		return errors.New("invalid cover image url")
	}

	return nil
}
