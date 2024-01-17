/*
import { createServer } from "http";
import { readFileSync } from "fs";
import { join } from "path";

const server = createServer((req, resp) => {
    readFileSync(join(process.cwd(), "public/index.html"), (err, data) => {
        resp.writeHead(200/!*, { "Content-Type": "text/html" }*!/);
        resp.end(data);
    })
/!*    resp.writeHead(200);
    resp.end("Hello, world!");*!/
});

server.listen(8000);*/

import express from "express";

const app = express();
app.use(express.static("public"));

app.listen(8000);