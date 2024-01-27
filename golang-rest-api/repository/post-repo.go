package repository

import (
	"context"
	"learn-architecture-golang/golang-rest-api/golang-rest-api/entity"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct {
}

// New Post Repository
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId     string = "golang-architecture"
	colletionName string = "post"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(colletionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {

		log.Fatalf("Failed to adding new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	itr := client.Collection(colletionName).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err != nil {
			if err == iterator.Done {
				// Iterator is done, break out of the loop
				break
			}
			log.Fatalf("Failed to iterate the list of post: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil

}
