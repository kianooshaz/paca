const mTable = document.querySelector('.action-table')
const trs = mTable.querySelectorAll('tbody tr')
const columns = []
trs[1].querySelectorAll('td').forEach((td, i) => {
    if (i == 0) {
        columns.push(td.innerHTML)
    } else {
        columns.push(td.querySelector('b').innerHTML)
    }
})

const obj = {}
trs.forEach((tr, i) => {
    let tds = tr.querySelectorAll('td')
    const rowNum = tds[0].innerHTML
    const tempObj = {}
    if (i == 1) {
        tds = tr.querySelectorAll('td b')
        tds = [tr.querySelector('td'), ...tds]
    }
    tds.forEach((td, i) => {
        const tdValue = td.innerHTML
        const colValue = columns[i]
        tempObj[colValue] = tdValue
    })
    obj[rowNum] = tempObj
})

obj
