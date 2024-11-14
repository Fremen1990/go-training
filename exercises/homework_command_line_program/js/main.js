const echo = require("./echo")
const cat = require('./cat')
const find = require("./find")

const args = process.argv.slice(2)

const command = args[0]

const commandArgs = args.slice(1)


switch (command) {
    case 'echo':
        echo(commandArgs);
        break;
    case 'cat':
        cat(commandArgs);
        break;
    case 'find':
        find(commandArgs);
        break;
    default:
        console.log(`Unknown command: ${command}`);
        console.log("Available commands: echo, anotherTool, someOtherTool");
}