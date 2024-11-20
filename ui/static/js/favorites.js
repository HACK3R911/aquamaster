// favorites.js

// Переменная для количества избранных товаров
let favoriteCount = 1; // Устанавливаем значение на 1, чтобы тестировать отображение товаров

// Элементы для отображения статуса и товаров
const favoritesStatus = document.getElementById("favorites-status");
const favoriteItems = document.getElementById("favorite-items");
const favoritesList = document.getElementById("favorites-list");

// Функция для обновления статуса избранного
function updateFavoritesStatus() {
    // Проверяем количество товаров
    if (favoriteCount === 0) {
        favoritesStatus.innerHTML = "<p>Вы никого не избрали :(</p>";
        favoriteItems.style.display = "none"; // Скрываем секцию товаров
    } else {
        favoritesStatus.innerHTML = `<p>У вас ${favoriteCount} товар(ов) в избранном</p>`;
        favoriteItems.style.display = "block"; // Показываем секцию товаров
        displayFavoriteItems(); // Вызываем функцию для отображения товаров
    }
}

// Функция для отображения товаров в избранном
function displayFavoriteItems() {
    // Очищаем список перед добавлением новых элементов
    favoritesList.innerHTML = "";

    // Проверяем, если `favoriteCount` больше 0, добавляем пример товара
    if (favoriteCount > 0) {
        favoritesList.innerHTML += `
            <li>
                <h3>Душевая кабина</h3>
                <p>Цена: 12000 руб.</p>
                <button onclick="removeFavoriteItem()">Удалить</button>
            </li>
        `;
    }
}

// Функция для удаления товара из избранного
function removeFavoriteItem() {
    favoriteCount = 0; // Обнуляем количество избранных товаров
    updateFavoritesStatus(); // Обновляем статус избранного
}

// Инициализация отображения статуса избранного
updateFavoritesStatus();
