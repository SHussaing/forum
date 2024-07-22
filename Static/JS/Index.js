function logout() {
    fetch('/Logout', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: 'include' // to include cookies in the request
    })
    .then(response => {
        if (response.ok) {
            window.location.href = '/';
        } else {
            alert('Logout failed.');
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Logout failed.');
    });
}
