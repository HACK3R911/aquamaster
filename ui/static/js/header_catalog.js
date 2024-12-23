const template = document.createElement('template');
template.innerHTML = `
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Aquamaster</title>
    <link rel="stylesheet" href="../../static/css/main_style.css">
    <link rel="website icon" type="png" href="../img/images/icons/search.png">
`;

document.head.appendChild(template.content);

document.addEventListener('DOMContentLoaded', () => {
    const navbarHTML = `
    <div class="navbar">
        <div class="navbar-item">
            <input class="search">
        </div>
        <div class="navbar-item">
            <a href="../../favorites.html">Избранное</a>
        </div>
        <div class="navbar-item">
            <a href="../../cart.html">Корзина</a>
        </div>
                <div class="navbar-item">
            <a href="../../plumbing.html">Сантехника</a>
        </div>
        <div class="navbar-item">
            <a href="../../interior.html">Интерьер</a>
        </div>
        <div class="navbar-item">
            <a href="../../accessories.html">Аксессуары</a>
        </div>
        <div class="navbar-item">
            <a href="../../MainMenu.html">Вернуться к истокам</a>
        </div>
    </div>
`;
    document.body.insertAdjacentHTML('afterbegin', navbarHTML);
});

const modal = document.createElement('script');
modal.src = "../../static/css/js/modal.js";
modal.defer = true;
document.head.appendChild(modal);