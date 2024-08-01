async function handleLikeDislike(type, id, action) {
    const url = `/LikeDislike?type=${type}&id=${id}&action=${action}`;
    const response = await fetch(url, { method: 'POST' });
    if (response.ok) {
        const data = await response.json();
        if (type === 'post') {
            document.getElementById('post-likes-dislikes').innerText = `Likes: ${data.likes}, Dislikes: ${data.dislikes}`;
        } else if (type === 'comment') {
            document.getElementById(`comment-likes-dislikes-${id}`).innerText = `Likes: ${data.likes}, Dislikes: ${data.dislikes}`;
        }
    } else {
        const errorText = await response.text();
        console.error('Error:', response.status, errorText);
        alert('Failed to update like/dislike');
    }
}
