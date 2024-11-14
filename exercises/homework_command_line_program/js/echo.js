const echo = (args) => {
    console.log("Echoing...")
    args.forEach(arg => {
        console.log(arg)
    })
}

module.exports = echo