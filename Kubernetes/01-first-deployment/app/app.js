const express = require('express');

const app = express();

app.get('/', (req, res) => {
  res.send(`
    <h1>Hello from this Node.js app!</h1>
    <p>Try sending a request to /error and see what happens</p>
  `);
});

app.get('/health', (req, res) => {
  res
      .status(200)
      .json("Healthy");
});

app.get('/error', (req, res) => {
  process.exit(1);
});

app.listen(8080);
