export function model(post) {
    const model = document.createElement('div');
    model.className = 'model';
    model.innerHTML =`
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
    </div>
    <div class="comment-area"></div>
    `;
    const comment_section = model.querySelector('.comment-area');
    for (let comment of post.comments) {
        const comment_card = document.createElement('div');
        comment_card.className = 'comment';
        comment_card.innerHTML = `
            <div class="comment-info">
                <span>${comment.author}</span>
                <span>${comment.date}</span>
            </div>
            <div class="comment-content">
                <p>${comment.content}</p>
            </div>
        `;
        comment_section.appendChild(comment_card);
    }
    return model;
}