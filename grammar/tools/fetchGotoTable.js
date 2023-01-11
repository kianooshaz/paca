const gotoTable = document.querySelector('.goto-table')
const rows = gotoTable.querySelectorAll('tr')

const columns = []
rows[1].querySelectorAll('td').forEach((td, i) => {
    columns.push(td.innerHTML)
})

const obj = {}
rows.forEach((tr, i) => {
    let tds = tr.querySelectorAll('td')
    const rowNum = tds[0].innerHTML
    const tempObj = {}
    tds.forEach((td, i) => {
        const tdValue = td.innerHTML
        const colValue = columns[i]
        tempObj[colValue] = tdValue
    })
    obj[rowNum] = tempObj
})
obj
