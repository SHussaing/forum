// Function to get a cookie by name
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
    return null;
}

// Function to check if the user has a session_token
function hasSessionToken() {
    return getCookie('session_token') !== null;
}

// Example usage to hide or show elements
document.addEventListener('DOMContentLoaded', () => {
    const elementsToShow = document.querySelectorAll('.show-on-login');
    const elementsToHide = document.querySelectorAll('.hide-on-login');

    if (hasSessionToken()) {
        elementsToShow.forEach(element => element.style.display = 'block');
        elementsToHide.forEach(element => element.style.display = 'none');
    } else {
        elementsToShow.forEach(element => element.style.display = 'none');
        elementsToHide.forEach(element => element.style.display = 'block');
    }
});
