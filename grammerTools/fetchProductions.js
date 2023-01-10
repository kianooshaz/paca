const appStates = document.querySelectorAll(
    '#output .debug:not(:nth-child(1), :nth-child(2))'
)

const stateList = {}

appStates.forEach((state) => {
    const rowTitle = state.querySelector('tr .tabtitle')?.innerText
    const rows = state.querySelectorAll('tr')
    const productions = []
    rows.forEach((row) => {
        const tds = row.querySelectorAll('td')
        productions.push({
            Lookahead: '' + tds[0]?.innerText,
            Production: '' + tds[1]?.innerText,
        })
    })
    stateList[rowTitle] = productions
})
