// Получаем все кнопки для открытия модальных окон
var buttons = document.querySelectorAll(".openModal");
var modals = document.querySelectorAll(".modal");
var spans = document.getElementsByClassName("modal-close");

// Функция для открытия модального окна
buttons.forEach(function(button) {
    button.onclick = function() {
        var modalId = this.getAttribute("data-modal");
        document.getElementById(modalId).style.display = "block";
    }
});

// Функция для закрытия модального окна
for (let i = 0; i < spans.length; i++) {
    spans[i].onclick = function() {
        this.parentElement.parentElement.style.display = "none";
    }
}

// Закрытие модального окна при клике вне его
window.onclick = function(event) {
    modals.forEach(function(modal) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    });
}