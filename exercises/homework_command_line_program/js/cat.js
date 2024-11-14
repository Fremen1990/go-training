const fs = require('fs');
const readline = require('readline');

/**
 * Opens a file and scans its lines, handling flags for formatted output.
 * @param {string} fileName - The name of the file to open and scan.
 * @param {string} flag - The flag to determine output format ('', '-n', '-nb').
 */
const openAndScanFile = async (fileName, flag) => {
    try {
        const fileStream = fs.createReadStream(fileName);

        const rl = readline.createInterface({
            input: fileStream,
            crlfDelay: Infinity
        });

        let lineNumber = 1;

        for await (const line of rl) {
            switch (flag) {
                case '':
                    console.log(line);
                    break;
                case '-n':
                    console.log(`${lineNumber}. ${line}`);
                    lineNumber++;
                    break;
                case '-nb':
                    if (line.trim() === '') {
                        console.log('');
                    } else {
                        console.log(`${lineNumber}. ${line}`);
                        lineNumber++;
                    }
                    break;
                default:
                    console.log("Flag not recognized, please use --help to see documentation :D");
                    rl.close();
                    return;
            }
        }
    } catch (error) {
        console.error(`Unable to read file: ${error.message}`);
    }
}


 const cat = async (args) => {
    if (args.length > 0) {
        const fileNameFromCommandLine = args[0];
        if (args.length === 1) {
            await openAndScanFile(fileNameFromCommandLine, '');
        } else if (args.length === 2) {
            await openAndScanFile(fileNameFromCommandLine, args[1]);
        }
    } else {
        console.log("Please provide a file name.");
    }
}

module.exports = cat