var express = require('express')

var app = express()

app.use(express.static('.'))

let port = 10000;
app.listen(port)

console.log(`\nNow open the URL ---> http://localhost:${port} <--- in your browser.`);
