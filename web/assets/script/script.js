import {like, dislike} from './reaction.js';

async function getAllPosts() {
    const jsonData = await fetch('/api/posts');
    const posts = await jsonData.json();
    return posts
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
    }else {
        return "just now"
    }
}

async function loadData(posts) {
    const container = document.getElementById('container')
    for (let post of posts) {
        post.created_at = formatTime(post.created_at);
        const card = document.createElement('div');
        card.className = 'post';
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
                        <span class="likes-count">${post.likes_count > 0 ? post.likes_count : ""}</span>
                        <button class="like"><i style="color: ${post.liked ? 'blue' : 'white'};" class="fi fi-rr-social-network"></i></button>
                    </div>
                    <div class="dislike-area">
                        <span class="dislikes-count">${post.dislike_count > 0 ? post.dislike_count : ""}</span>
                        <button class="dislike"><i style="color: ${post.disliked ? 'blue' : 'white'};" class="fi fi-rr-hand"></i></button>
                    </div>
                </div>
                <button class="comment"><span class="comment-count">${post.comments ?post.comments.length: 'no'}</span> comments</button>
            </div>
        `;
        //add event listeners
        const likeButton = card.querySelector('.like')
        const dislikeButton = card.querySelector('.dislike');

        likeButton.addEventListener('click', () => {
            if (!post.liked && post.disliked) dislike(post, card, dislikeButton)
            like(post, card, likeButton)
        });

        dislikeButton.addEventListener('click', () => {
            if (!post.disliked && post.liked) like(post, card, likeButton)
            dislike(post, card, dislikeButton)
        } );
        container.appendChild(card)
        console.log(post)
    }
}



async function main() {
    const posts = await getAllPosts();
    loadData(posts)
}
main()
