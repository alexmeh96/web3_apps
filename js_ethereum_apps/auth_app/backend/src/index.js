const express = require('express')
const cors = require('cors')
const router = require('./route')
const Session = require('express-session')

const app = express()
app.use(express.json())

app.use(cors({
        origin: 'http://localhost:3000',
        credentials: true,
}))

app.use(
    Session({
        name: 'siwe-quickstart',
        secret: "siwe-quickstart-secret",
        resave: true,
        saveUninitialized: true,
        cookie: {secure: false, sameSite: true}
    })
)
app.use(router)

console.log("server started on port 8085")
app.listen(8085)
