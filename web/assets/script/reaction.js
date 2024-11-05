export async function like(post, card, likeButton) {
    const icon = likeButton.querySelector('i'); // Get the icon element
    post.liked = !post.liked; // Toggle liked state
    post.likes_count += post.liked ? 1 : -1; // Update likes count

    const action = post.liked ? 'add' : 'remove';
    const resp = await fetch('/api/posts/like', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({post_id: post.id, like_action: action}),})
    if (!resp.ok) {
        window.location.replace("/login");
        return;
    }

    // Update the icon color
    icon.style.color = post.liked ? 'blue' : 'white';

    // Update the likes count in the same div
    const countSpan = card.querySelector('.likes-count'); // Ensure you have a way to select the likes count
    countSpan.textContent = post.likes_count > 0 ? post.likes_count : "";
}

export async function dislike(post, card, dislikeButton) {
    const icon = dislikeButton.querySelector('i'); // Get the icon element
    post.disliked =!post.disliked; // Toggle disliked state
    //send data to the server
    const action = post.disliked ? 'add' : 'remove';
    const resp = await fetch('/api/posts/dislike', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ post_id: post.id, dislike_action: action}),})
    if (!resp.ok) {
        window.location.replace("/login");
        return;
    }
    post.dislike_count += post.disliked ? 1 : -1; // Update dislikes count

    // Update the icon color
    icon.style.color = post.disliked ? 'blue' : 'white';

    // Update the dislikes count in the same div
    const countSpan = card.querySelector('.dislikes-count'); // Ensure you have a way to select the dislikes count
    countSpan.textContent = post.dislike_count > 0 ? post.dislike_count : "";
}