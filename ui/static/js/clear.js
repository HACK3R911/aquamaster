function clearSearch() {
    const searchInput = document.querySelector('.search');
    searchInput.value = '';
}

// Обработчик события для кликов по документу
document.addEventListener('click', function(event) {
    const searchInput = document.querySelector('.search');
    if (!searchInput.contains(event.target)) {
        clearSearch();
    }
});