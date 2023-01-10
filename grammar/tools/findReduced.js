const path = require('path')
const fs = require('fs')
const file = fs.readFileSync(path.join(__dirname, '../data/actionTable.json'))

const obj = JSON.parse(file)
Object.keys(obj).map((key) => {
    const rules = []
    Object.values(obj[key]).forEach((v) => {
        if (new RegExp('r').test(v)) {
          rules.push(v)
        }
    })
    console.log(key, ':', rules.join(','))
})
