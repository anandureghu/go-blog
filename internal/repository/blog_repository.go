package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/anandureghu/go-blog/internal/model"
	"github.com/joho/godotenv"
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

func (b *BlogRepository) GetAllBlogs() []model.Blog {
	rows, err := b.Db.Query(`
	SELECT
	id, title, description, cover, name, avatar, created_at, updated_at
	FROM blogs
	`)
	if err != nil {
		log.Fatalln("can't get all blogs")
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
			log.Println("can't iterate over blogs")
		}

		blogs = append(blogs, blog)
	}

	return blogs
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

func (b *BlogRepository) CreateBlog(blog model.Blog) {
	create_blog := `
	INSERT INTO blogs
	(title, description, cover, name, avatar, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := b.Db.Exec(create_blog,
		blog.Title,
		blog.Description,
		blog.Cover,
		blog.Name,
		blog.Avatar,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		log.Println("Cannot create blog", err)
	}
}

func (b *BlogRepository) UpdateBlog(id int, blog model.Blog) model.Blog {
	return blog
}

func (b *BlogRepository) DeleteBlog(id int) {
	delete_blog := `
		DELETE FROM blogs
		WHERE id=$1
	`
	b.Db.Exec(delete_blog, id)
}
