async function addComment(event) {
    event.preventDefault();
    if (!validateForm()) return;

    const form = document.getElementById('comment-form');
    const formData = new FormData(form);
    const postID = formData.get('post_id');
    const content = formData.get('content');

    const response = await fetch('/AddComment', {
        method: 'POST',
        body: new URLSearchParams(formData)
    });

    if (response.ok) {
        const data = await response.json();
        const commentsSection = document.getElementById('comments-section');
        const newComment = document.createElement('div');
        newComment.className = 'comment';
        newComment.id = `comment-${data.commentID}`;
        newComment.innerHTML = `
            <p><span class="username">${data.username}</span>:</p>
            <p>${data.content}</p>
            <div class="actions">
                <p id="comment-likes-dislikes-${data.commentID}">Likes: 0, Dislikes: 0</p>
                <button class="like-button show-on-login" onclick="handleLikeDislike('comment', ${data.commentID}, 'like')">Like</button>
                <button class="dislike-button show-on-login" onclick="handleLikeDislike('comment', ${data.commentID}, 'dislike')">Dislike</button>
            </div>
        `;
        commentsSection.appendChild(newComment);
        form.reset();
    } else {
        const errorText = await response.text();
        console.error('Error:', response.status, errorText);
        alert('Failed to add comment');
    }
}

function validateForm() {
    const content = document.getElementById('content').value.trim();
    if (content === '') {
        alert('Comment cannot be empty');
        return false;
    }
    return true;
}