var express = require('express')

var app = express()

app.use(express.static('.'))

let port = 9090;
app.listen(port)

console.log(`===> Now open the URL http://localhost:${port} in your browser.`);
