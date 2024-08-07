function handleImageError(imageId) {
    // Get the image element by ID
    var img = document.getElementById(imageId);
    if (img) {
        // Hide the image element if it fails to load
        img.style.display = 'none';
    }
}