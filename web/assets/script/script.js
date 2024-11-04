import {like, dislike} from './reaction.js';

async function getAllPosts() {
    const jsonData = await fetch('/api/posts');
    const posts = await jsonData.json();
    return posts
}

async function loadData(posts) {
    const container = document.getElementById('container')
    for (let post of posts) {
        const card = document.createElement('div');
        card.className = 'post';
        card.innerHTML = `
            <div class="info">
                <span>${post.category}</span>
                <span>${post.author}</span>
                <span>3 min</span>
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
    }
}

async function main() {
    const posts = await getAllPosts();
    loadData(posts)
}
main()
