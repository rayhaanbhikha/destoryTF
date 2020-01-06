const fs = require('fs')
const path = require('path')
console.log("hello running a javascript action")

console.log(fs.readdirSync(path.join(__dirname, "../../..")))