    function copyEmailToClipboard(email) {
        const tempInput = document.createElement('input');
        tempInput.value = email;
        document.body.appendChild(tempInput);
        
        tempInput.select();
        tempInput.setSelectionRange(0, 99999);
        
        document.execCommand('copy');
        
        document.body.removeChild(tempInput);

        alert('Email скопирован: ' + email);
    }
