package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/anandureghu/go-blog/internal/model"
	"github.com/joho/godotenv"
	"github.com/manveru/faker"
)

func init() {
	godotenv.Load()

	connection_string := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
		os.Getenv("POSTGRES_CONNECTION_USER"),
		os.Getenv("POSTGRES_CONNECTION_PASSWORD"),
		os.Getenv("POSTGRES_CONNECTION_HOST"),
		os.Getenv("POSTGRES_CONNECTION_DBNAME"),
	)

	// connection_string := "postgres://postgres:qburst@localhost/goblog?sslmode=disable"
	Connect(connection_string)
}

type BlogRepository struct {
	Db *sql.DB
}

func NewBlogRepository() *BlogRepository {

	conn := GetConnection()
	return &BlogRepository{
		Db: conn,
	}
}

func (b *BlogRepository) GetAllBlogs() ([]model.Blog, error) {
	rows, err := b.Db.Query(`
	SELECT
	id, title, description, cover, name, avatar, created_at, updated_at
	FROM blogs
	`)
	if err != nil {
		log.Println("can't get all blogs")
		return nil, err
	}
	defer rows.Close()

	blogs := []model.Blog{}
	blog := model.Blog{}

	for rows.Next() {
		err := rows.Scan(
			&blog.Id,
			&blog.Title,
			&blog.Description,
			&blog.Cover,
			&blog.Name,
			&blog.Avatar,
			&blog.CreatedAt,
			&blog.UpdatedAt,
		)

		if err != nil {
			log.Println("can't scan blogs")
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	return blogs, rows.Err()
}

func (b *BlogRepository) GetBlog(id int) (model.Blog, error) {
	get_blog := `
		SELECT 
		id, title, description, cover, name, avatar, created_at, updated_at
		FROM blogs
		WHERE id=$1
	`
	blog := model.Blog{}
	row := b.Db.QueryRow(get_blog, id)
	row.Scan(
		&blog.Id,
		&blog.Title,
		&blog.Description,
		&blog.Cover,
		&blog.Name,
		&blog.Avatar,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	)
	return blog, row.Err()
}

func (b *BlogRepository) CreateBlog(blog model.Blog) error {

	fake, _ := faker.New("en")

	create_blog := `
	INSERT INTO blogs
	(title, description, cover, name, avatar, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := b.Db.Exec(create_blog,
		blog.Title,
		blog.Description,
		blog.Cover,
		fake.Name(),
		fake.URL(),
		time.Now(),
		time.Now(),
	)

	if err != nil {
		log.Println("cannot create blog", err)
		return err
	}

	return nil
}

func (b *BlogRepository) UpdateBlog(id int, blog model.Blog) (model.Blog, error) {
	// Checking if blog is in databse
	ub, err := b.GetBlog(id)

	if ub.Id == 0 {
		// Throe not found
		log.Println("No blog with id", id)
		return model.Blog{}, fmt.Errorf("blog with id %v not found", id)
	}
	if err != nil {
		return model.Blog{}, err
	}
	blog.Id = id

	update_blog := `
		UPDATE blogs set
		title=$1,
		description=$2, 
		cover=$3,
		name=$4,
		avatar=$5,
		updated_at=$6
		WHERE id=$7
	`
	_, err = b.Db.Exec(update_blog,
		blog.Title,
		blog.Description,
		blog.Cover,
		blog.Name,
		blog.Avatar,
		time.Now(),
		blog.Id,
	)

	return blog, err
}

func (b *BlogRepository) DeleteBlog(id int) error {

	// Checking if blog is in databse
	ub, err := b.GetBlog(id)

	if err != nil {
		return err
	}
	if ub.Id == 0 {
		// Throe not found
		log.Println("No blog with id", id)
		return fmt.Errorf("blog with id %v not found", id)
	}

	delete_blog := `
		DELETE FROM blogs
		WHERE id=$1
	`
	_, err = b.Db.Exec(delete_blog, id)

	return err
}
