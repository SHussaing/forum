document.addEventListener('DOMContentLoaded', () => {
    function truncateText(selector, maxLength) {
        const elements = document.querySelectorAll(selector);
        elements.forEach(element => {
            if (element.textContent.length > maxLength) {
                const truncated = element.textContent.slice(0, maxLength) + '...';
                element.textContent = truncated;
            }
        });
    }

    truncateText('.post-title', 50);  // Limit for the title
    truncateText('.post-content', 150);  // Limit for the content
});