package main

import (
	"GoApi/entity"
	"GoApi/repository"
	"encoding/json"
	"math/rand"
	"net/http"
)

var (
	//posts []entity.Post
	repo repository.PostRepository = repository.NewPostRepository()
)

/*func init() {
	posts = []entity.Post{
		{
			ID:    1,
			Title: "TITLE 1",
			Text:  " text 1",
		},
	}

}*/

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error":"Error getting the posts"}`))
		return

	}
	/*result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error":"Error marshaling"}`))
		return

	}*/
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
	//resp.Write(result)
}
func addPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error":"Error unmarshaling the request"}`))
		return

	}
	post.ID = rand.Int63()
	//posts = append(posts, post)
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	/*result, err := json.Marshal(post)
	resp.Write(result)*/
	json.NewEncoder(resp).Encode(post)

}
