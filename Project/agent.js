// Создание контейнера для чата
document.addEventListener('DOMContentLoaded', () => {
    const chatContainer = document.createElement('div');
    chatContainer.setAttribute('id', 'chatBubbleRoot');
    document.body.appendChild(chatContainer);

    // Настройка токена агента
    window.agx = '674c390983db69d68e1818febxvDyQkbBy+1JCaeZ4Xavg==|MzX1101DlUteaWpvXFyNANoeHUf/Z8WEtKiQ8v440EQ=';

    // Подключение скрипта чата
    const chatScript = document.createElement('script');
    chatScript.src = 'https://storage.googleapis.com/agentx-cdn-01/agentx-chat.js';
    chatScript.defer = true;
    document.head.appendChild(chatScript);

    // Проверка статуса загрузки чата
    setTimeout(() => {
        if (chatContainer.childElementCount > 0) {
            console.log('Chat loaded successfully!');
        } else {
            console.error('Failed to load chat. Please check your configuration.');
        }
    }, 3000); // Проверить через 3 секунды
});
