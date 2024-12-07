const template = document.createElement('template');
template.innerHTML = `
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Aquamaster</title>
    <link rel="stylesheet" href="static/css/main_style.css">
    <script defer src="plumbing.js"></script>
    <link rel="website icon" type="png" href="../img/images/icons/search.png">
`;

// Добавляем содержимое в head
document.head.appendChild(template.content);