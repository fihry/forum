import { like, dislike } from "./reaction.js";

let loading = false; // Flag to prevent multiple requests

async function getAllPosts() {
  const jsonData = await fetch("/api/posts");
  if (jsonData.ok) {
    const posts = await jsonData.json();
    return posts;
  } else {
    console.error("Failed to fetch posts");
  }
}
// THIS FUNCTION FORMATS THE DATE 20-2024-11-05T19:30:00Z TO "2 years ago"
function formatTime(time) {
  const itemTime = new Date(time).getTime();
  const currentTime = new Date();
  const timeDiff = currentTime - itemTime;
  const days = Math.floor(timeDiff / (1000 * 60 * 60 * 24));
  const hours = Math.floor(timeDiff / (1000 * 60 * 60));
  const minutes = Math.floor(timeDiff / (1000 * 60));
  const seconds = Math.floor(timeDiff / 1000);
  if (days > 0) {
    return `${days} days ago`;
  } else if (hours > 0) {
    return `${hours} hours ago`;
  } else if (minutes > 0) {
    return `${minutes} minutes ago`;
  } else if (seconds > 0) {
    return `${seconds} seconds ago`;
  } else {
    return "just now";
  }
}

async function loadData(posts) {
  const container = document.getElementById("container");
  for (let post of posts) {
    post.created_at = formatTime(post.created_at);
    const card = document.createElement("div");
    card.className = "post";
    card.innerHTML = `
            <div class="info">
                <span>${post.category}</span>
                <span>${post.author}</span>
                <span>${post.created_at ? post.created_at : ""}</span>
            </div>
            <div class="content">
                <h2>${post.title}</h2>
                <p>${post.content}</p>
            </div>
            <div class="reaction">
                <div class="likes">
                    <div class="like-area">
                        <span class="likes-count">${
                          post.likes_count > 0 ? post.likes_count : ""
                        }</span>
                        <button class="like"><i style="color: ${
                          post.liked ? "blue" : "white"
                        };" class="fi fi-rr-social-network"></i></button>
                    </div>
                    <div class="dislike-area">
                        <span class="dislikes-count">${
                          post.dislike_count > 0 ? post.dislike_count : ""
                        }</span>
                        <button class="dislike"><i style="color: ${
                          post.disliked ? "blue" : "white"
                        };" class="fi fi-rr-hand"></i></button>
                    </div>
                </div>
                <button class="comment"><span class="comment-count">${
                  post.comments ? post.comments.length : "no"
                }</span> comments</button>
            </div>
        `;
    //add event listeners
    const likeButton = card.querySelector(".like");
    const dislikeButton = card.querySelector(".dislike");

    likeButton.addEventListener("click", () => {
      if (!post.liked && post.disliked) dislike(post, card, dislikeButton);
      like(post, card, likeButton);
    });

    dislikeButton.addEventListener("click", () => {
      if (!post.disliked && post.liked) like(post, card, likeButton);
      dislike(post, card, dislikeButton);
    });
    container.appendChild(card);
  }
}

async function isLoggedIn() {
  try {
    const response = await fetch("/api/session_check");
    return response.ok;
  } catch (error) {
    console.error("Error checking login status:", error);
    return false;
  }
}

function updateNavbar(loggedIn) {
  if (loggedIn) {
    const navbarAcions = document.getElementById("navbar_actions");
    navbarAcions.innerHTML = `
        <a class="newpost">new post</a>
    `;
  }
}

async function main() {
  const postsPromise = getAllPosts();
  const loggedIn = await isLoggedIn();
  updateNavbar(loggedIn);
  const posts = await postsPromise;
  loadData(posts);
  setupCreatePostModal();
}
main();

export function setupCreatePostModal() {
  const modal = document.getElementById("createPostModal");
  const form = document.getElementById("createPostForm");

  // Add click event to the "new post" button
  document.querySelector(".newpost").addEventListener("click", function (e) {
    e.preventDefault();
    modal.style.display = "flex";
  });

  // Close modal when clicking outside
  window.addEventListener("click", function (event) {
    if (event.target === modal) {
      modal.style.display = "none";
    }
  });

  // Handle form submission
  form.addEventListener("submit", async function (e) {
    e.preventDefault();

    const formData = new FormData(form);
    const postData = {
      title: formData.get("title"),
      category: formData.get("category"),
      content: formData.get("content"),
    };

    try {
      const response = await fetch("/api/createPost", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(postData),
      });

      if (response.ok) {
        // Clear form and close modal
        form.reset();
        modal.style.display = "none";
        // Refresh the page or update posts list
        window.location.reload();
      } else {
        // If user is not logged in, redirect to login page
        if (response.status === 401) {
          window.location.replace("/login");
          return;
        }
        // Display error message
        const errMsg = await response.text();
        alert(errMsg);
      }
    } catch (error) {
      console.error("Error:", error);
    }
  });
}
