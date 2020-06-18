package repository

import (
	"GoApi/entity"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
)

const (
	projectId      string = "goproject-25d97"
	collectionName string = "posts"
)

type repo struct {
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile("C:/Users/Lenovo/go/src/GoApi/cred.json"))
	if err != nil {
		log.Fatalf("failed to create firestore client: %v ", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("failed to add post: %v ", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile("C:/Users/Lenovo/go/src/GoApi/cred.json"))
	if err != nil {
		log.Fatalf("failed to create firestore client: %v ", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("failed to retrieve the list of posts: %v ", err)
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

//newPostRepository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}
