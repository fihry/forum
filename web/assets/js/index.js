import { like, dislike } from "./reaction.js";
import { checkLogin } from "./auth.js";

// this function fetches all posts from the server
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
    const currentTime = new Date()
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
        return "just now"
    }
}

// this is the post component
const PostComponent = `<div class="post-header">
                            <div class="post-author">Author</div>
                            <div class="creation-date">
                                <span>2 days ago</span>
                            </div>
                        </div>
                        <div class="post-body">
                            <div class="post-title">
                                <h4></h4>
                            </div>
                            <div class="post-content">
                            </div>
                            <div class="post-actions">
                                <span class="like"><i class="fi fi-rr-heart"></i> <span class="count">22</span></span>
                                <span class="dislike"><i class="fi fi-rr-heart-crack"></i><span class="count">12</span></span>
                                <span class="comment"><i class="fi fi-rr-comment"></i>Comment</span>
                            </div>
                        </div>
                    `
// this function displays the posts on the page
async function displayData(posts) {
    const container = document.querySelector(".posts");
    for (let post of posts) {
        post.created_at = formatTime(post.created_at);
        const PostCard = document.createElement("div");
        PostCard.className = "post";
        PostCard.innerHTML = PostComponent;
        const card = PostCard.querySelector(".post-body");
        card.querySelector(".post-title h4").innerText = post.title;
        card.querySelector(".post-content").innerText = post.content;
        card.querySelector(".post-author").innerText = post.author;
        card.querySelector(".creation-date span").innerText = post.created_at;
        card.querySelector(".like .count").innerText = post.likes_count;
        card.querySelector(".dislike .count").innerText = post.dislike_count;
        card.querySelector(".comment").innerText = post.comments ? post.comments.length : "no";
        card.querySelector(".like").style.color = post.liked ? "red" : "white";
        card.querySelector(".dislike").style.color = post.disliked ? "blue" : "white";
        card.querySelector(".like").addEventListener("click", () => {
            if (!post.liked && !post.disliked) {
                like(post, card);
            }
        });
        card.querySelector(".dislike").addEventListener("click", () => {
            if (!post.disliked && !post.liked) {
                dislike(post, card);
            }
        });

        card.querySelector(".comment").addEventListener("click", () => {
            window.location.href = `/post/${post.id}`;
        });

        container.appendChild(PostCard);
    }
}


function updateNavbar(loggedIn) {
    if (loggedIn) {
        const navbarActions = document.getElementById("navbar_actions");
        navbarActions.innerHTML = `
        <a class="newpost">new post</a>
        <a class="logout" >logout</a>
    `;
    } else {
        const navbarActions = document.getElementById("navbar_actions");
        navbarActions.innerHTML = `
        <a class="login" href="/auth">register</a>
    `;
    }
}

export function setupCreatePostModal() {
    const modal = document.getElementById("createPostModal");
    const form = document.getElementById("createPostForm");

    // Add click event to the "new post" button
    document.querySelector(".newpost").addEventListener("click", (e) => {
        e.preventDefault();
        modal.style.display = "flex";
    });

    // Close modal when clicking outside
    window.addEventListener("click", (e) => {
        if (e.target === modal) {
            modal.style.display = "none";
        }
    });

    // Handle form submission
    form.addEventListener("submit", async (e) => {
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
                    window.location.replace("/auth");
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


// Main function to load posts and setup event listeners
async function main() {
    const postsPromise = getAllPosts();
    const loggedIn = await checkLogin();
    updateNavbar(loggedIn);
    const posts = await postsPromise;
    displayData(posts);
    setupCreatePostModal();
}
main();

