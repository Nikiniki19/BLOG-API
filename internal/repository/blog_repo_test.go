package repository_test

import (
	"blog-api/internal/model"
	"blog-api/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "postgresql://neondb_owner:npg_FYHLDX4Qf3ns@ep-twilight-leaf-a4sw3hym.us-east-1.aws.neon.tech/test_db?sslmode=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	if err := db.AutoMigrate(&model.BlogPost{}); err != nil {
		t.Fatalf("Failed to migrate test DB: %v", err)
	}

	t.Cleanup(func() {
		if err := db.Exec("TRUNCATE TABLE blog_posts RESTART IDENTITY CASCADE").Error; err != nil {
			t.Fatalf("Failed to clean up test DB: %v", err)
		}
	})

	return db
}

func TestCreatePost(t *testing.T) {
	db := setupTestDB(t)
	repo, err := repository.NewRepository(db)
	if err != nil {
		t.Fatalf("failed to create repository: %v", err)
	}

	post := &model.BlogPost{
		Title:       "Integration Title",
		Description: "Integration Desc",
		Body:        "Integration Body",
	}

	msg, err := repo.Create(post)
	assert.NoError(t, err)
	assert.Equal(t, "Blog post created successfully", msg)
	assert.NotZero(t, post.ID)
}

func TestGetAllPosts(t *testing.T) {
	db := setupTestDB(t)
	repo, err := repository.NewRepository(db)
	assert.NoError(t, err)

	db.Exec("DELETE FROM blog_posts")

	post := &model.BlogPost{Title: "Title", Description: "Desc", Body: "Body"}
	_, err = repo.Create(post)
	assert.NoError(t, err)

	posts, err := repo.GetAllPosts()
	assert.NoError(t, err)
	assert.Len(t, posts, 1)
	assert.Equal(t, "Title", posts[0].Title)
}

func TestGetPostByID(t *testing.T) {
	db := setupTestDB(t)
	repo, err := repository.NewRepository(db)
	assert.NoError(t, err)

	post := &model.BlogPost{Title: "Title", Description: "Desc", Body: "Body"}
	_, err = repo.Create(post)
	assert.NoError(t, err)

	gotPost, err := repo.GetPostByID(post.ID)
	assert.NoError(t, err)
	assert.Equal(t, post.Title, gotPost.Title)

	_, err = repo.GetPostByID(9999)
	assert.Error(t, err)
}

func TestUpdatePostByID(t *testing.T) {

	db := setupTestDB(t)
	db = db.Debug()
	repo, err := repository.NewRepository(db)
	assert.NoError(t, err)

	post := &model.BlogPost{
		Title:       "Old Title",
		Description: "Old Description",
		Body:        "Old Body",
	}
	msg, err := repo.Create(post)
	require.NoError(t, err)
	assert.Equal(t, "Blog post created successfully", msg)

	t.Log("Updating post")
	updatedPost := &model.BlogPost{
		ID:          post.ID, 
		Title:       "New Title",
		Description: "New Description",
		Body:        "New Body",
	}

	err = repo.UpdatePostByID(updatedPost)
	require.NoError(t, err)
}

func TestDeletePostByID(t *testing.T) {
	db := setupTestDB(t)
	repo, err := repository.NewRepository(db)
	assert.NoError(t, err)

	post := &model.BlogPost{
		Title:       "Title",
		Description: "Desc",
		Body:        "Body",
	}
	_, err = repo.Create(post)
	assert.NoError(t, err)

	err = repo.DeletePostByID(post.ID)
	assert.NoError(t, err)
	
	err = repo.DeletePostByID(post.ID)
	assert.Error(t, err)
}
