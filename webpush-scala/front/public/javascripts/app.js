
var applicationServerKey = 'BMrPL9meMLisGEIFriDzW65WRP8nJULnGbNwfsVvCDJbPGqmz5e9s7yV4tjj7+gSAn7+s0/W9oI/lscHFUZmGCo=';

var serviceWorker = 'serviceworker.js';

var subscribe = {
    id: 'subscribe',
    input: '/subscribe.json',
    method: 'POST'
};

var unsubscribe = {
    id: 'unsubscribe',
    input: '/unsubscribe.json',
    method: 'POST'
};

var testSubscription = {
    id: 'test-subscription',
    input: '/test-subscription.json',
    method: 'POST'
};

function _base64ToArrayBuffer(base64) {
    var binary_string =  window.atob(base64);
    var len = binary_string.length;
    var bytes = new Uint8Array( len );
    for (var i = 0; i < len; i++)        {
        bytes[i] = binary_string.charCodeAt(i);
    }
    return bytes.buffer;
}

function _arrayBufferToBase64(arrayBuffer) {
    return btoa(String.fromCharCode.apply(null, new Uint8Array(arrayBuffer)));
}

function subscriptionJson(subscription) {
    var p256dh = _arrayBufferToBase64(subscription.getKey('p256dh'));
    var auth = _arrayBufferToBase64(subscription.getKey('auth'));
    return {
        endpoint: subscription.endpoint,
        p256dh: p256dh,
        auth: auth
    };
}

function call(api, subscription) {
    var json = subscriptionJson(subscription);
    console.log(api.method, ' ', api.input, ' - ', json);

    fetch(api.input, {
        method: api.method,
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(json)
    }).then(res => res.text()).then(console.log).catch(console.error);
}

function callSubscribe(subscription) {
    call(subscribe, subscription);
}

function callUnsubscribe(subscription) {   
    call(unsubscribe, subscription);
}

function callTestSubscription(subscription) {
    call(testSubscription, subscription);
}

function getPushManager(f) {
    navigator.serviceWorker.register(serviceWorker).then(
        function(serviceWorkerRegistration) {
            f(serviceWorkerRegistration.pushManager)
        },
        function(error) {
            console.log(error);
        }
    );
}

function getSubscription(f) {
    getPushManager(function(pushManager) {
        pushManager.getSubscription().then(
            function(subscription){
                if (subscription) {
                    f(subscription);
                }
            }, 
            function(error){
                console.log(error);
            }
        );
    });
}

window.onload = function() {
    document.getElementById(subscribe.id).onclick = function() {
        getPushManager(
            function(pushManager) {
                pushManager.subscribe({
                    applicationServerKey: _base64ToArrayBuffer(applicationServerKey),
                    userVisibleOnly: true
                }).then(
                    function(pushSubscription) {
                        callSubscribe(pushSubscription);
                    }, 
                    function(error) {
                        console.log(error);
                    }
                );
            }
        );
    }
    
    document.getElementById(unsubscribe.id).onclick = function() {
        getSubscription(
            function(subscription) {
                subscription.unsubscribe().then(
                    function(ok){
                        callUnsubscribe(subscription);
                    },
                    function(error){
                        console.log(error);
                    }
                );
            }
        );
    }
    
    document.getElementById(testSubscription.id).onclick = function() {
        getSubscription(
            function(subscription) {
                callTestSubscription(subscription);
            }
        );
    }
}
