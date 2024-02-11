const path = require("path");
const Dotenv = require('dotenv-webpack');

module.exports = {
    context: __dirname,
    entry: "./src/main.js",
    output: {
        filename: "main.js",
        path: path.resolve(__dirname, "dist"),
    },
    mode: "none",
    plugins: [
        new Dotenv(),
    ]
}
