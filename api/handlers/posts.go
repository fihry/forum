package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"forum/api/controllers"
	"forum/api/models"
	"forum/utils"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	session, err := r.Cookie("session")
	if err != nil {
		// Session cookie is not present; return all posts without engagement info
		posts, err := controllers.GetAllPosts()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
		return
	}

	// User is authenticated; get user information based on session
	user, err := controllers.GetUserBySession(session.Value)
	log.Println("username by session", user.Username)
	if err != nil {
		// Invalid session; return all posts without engagement info
		posts, err := controllers.GetAllPosts()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
		return
	}

	// User is authenticated; get posts with engagement info
	posts, err := controllers.GetAllPostsByUser(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Respond with posts including engagement info
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Poste
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Check if user is authenticated
	session, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	user, err := controllers.GetUserBySession(session.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	post.Author = user.Username
	// Validate post data
	ok, err := utils.CheckDataForPost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	// Create new post here and return the created post
	id, err := controllers.CreatePoste(post)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	post.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
	w.WriteHeader(http.StatusCreated)
}

// Reaction handlers =================================
func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user, err := controllers.GetUserBySession(session.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	engagement := models.Engagement{}
	err = json.NewDecoder(r.Body).Decode(&engagement)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update post likes and dislikes
	var n int64
	post := models.Poste{}
	if engagement.LikeAction == "add" {
		controllers.GetPostByEngagement(user.ID, post)
		if *post.Liked {
			w.WriteHeader(http.StatusConflict)
			return
		}
		controllers.AddLikeToEngament(engagement.PosteID, 5)
		n, err = controllers.LikePost(engagement.PosteID)
	} else if engagement.LikeAction == "remove" {
		if !*post.Liked {
			w.WriteHeader(http.StatusConflict)
			return
		}
		controllers.RemLikeFromEngagement(engagement.PosteID, 5)
		n, err = controllers.RemoveLike(engagement.PosteID)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil || n == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}


func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session")
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
    user, err := controllers.GetUserBySession(session.Value)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
    engagement := models.Engagement{}
    err = json.NewDecoder(r.Body).Decode(&engagement)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Update post likes and dislikes
    var n int64
    post := models.Poste{}
	if engagement.DislikeAction == "add" {
		controllers.GetPostByEngagement(user.ID, post)
        if *post.Disliked {
            w.WriteHeader(http.StatusConflict)
            return
        }
        controllers.AddDislikeToEngament(engagement.PosteID, 5)
        n, err = controllers.LikePost(engagement.PosteID)
    } else if engagement.DislikeAction == "remove" {
		if !*post.Disliked {
            w.WriteHeader(http.StatusConflict)
            return
        }
		controllers.RemDislikeFromEngagement(engagement.PosteID, 5)
		n, err = controllers.RemoveDislike(engagement.PosteID)
		} else {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
	if err != nil || n == 0 {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
	w.WriteHeader(http.StatusOK)
}