import { createServer } from "http";

const server = createServer((_, res) => {
    res.writeHead(200);
    res.end("Hello, world!")
});
//
// server.on("request", (_, res) => {
//     res.write(JSON.stringify({
//         hello: "world"
//     }));
//
//     res.writeHead(200);
//     res.end();
// });

server.listen(8000);

