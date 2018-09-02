
this.onpush = function(event) {
    console.log(event.data.text());

    const title = 'Webpush 101';
    const options = {
      body: event.data.text()
    };
  
    event.waitUntil(self.registration.showNotification(title, options));  
}
