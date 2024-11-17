// cart.js

// Переменная для количества товаров в корзине
let itemCount = 1; // Можно изменить значение на 0, чтобы протестировать "Корзина пуста"

// Элемент для отображения статуса корзины
const cartStatus = document.getElementById("cart-status");
const cartItems = document.getElementById("cart-items");
const itemsList = document.getElementById("items-list");

// Функция для обновления статуса корзины
function updateCartStatus() {
    if (itemCount === 0) {
        cartStatus.innerHTML = "<p>Ваша корзина пуста, милорд</p>";
        cartItems.style.display = "none"; // Скрываем секцию товаров
    } else {
        cartStatus.innerHTML = `<p>У вас ${itemCount} товар(ов) в корзине</p>`;
        cartItems.style.display = "block"; // Показываем секцию товаров
        displayCartItems();
    }
}

// Функция для отображения товаров в корзине
function displayCartItems() {
    itemsList.innerHTML = ""; // Очищаем список перед добавлением
    // Добавляем пример товара
    itemsList.innerHTML += `
        <li>
            <h3>Смеситель для ванной</h3>
            <p>Цена: 3500 руб.</p>
            <button onclick="removeItem()">Удалить</button>
        </li>
    `;
}

// Функция для удаления товара из корзины
function removeItem() {
    itemCount = 0;
    updateCartStatus();
}

// Инициализация отображения статуса корзины
updateCartStatus();
