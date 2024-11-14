const fs = require('fs');
const path = require('path');

/**
 * Checks if a file or directory matches a specified type.
 * @param {fs.Stats} fileInfo - File information.
 * @param {string} fileType - Expected type ('file', 'directory', or 'link').
 * @returns {boolean} - True if the file matches the type, false otherwise.
 */
function checkType(fileInfo, fileType) {
    switch (fileType) {
        case 'file':
            return fileInfo.isFile();
        case 'directory':
            return fileInfo.isDirectory();
        case 'link':
            return fileInfo.isSymbolicLink();
        default:
            return false;
    }
}

/**
 * Walks through the directory and finds files or directories matching a pattern and type.
 * @param {string} startPath - The starting path for the search.
 * @param {string} pattern - Pattern to match file names.
 * @param {string} fileType - Type of file to match ('file', 'directory', or 'link').
 */
function findLogic(startPath, pattern, fileType) {
    function walk(currentPath) {
        fs.readdir(currentPath, { withFileTypes: true }, (err, entries) => {
            if (err) {
                console.error(`Error reading directory: ${err.message}`);
                return;
            }

            entries.forEach(entry => {
                const fullPath = path.join(currentPath, entry.name);

                fs.stat(fullPath, (err, stats) => {
                    if (err) {
                        console.error(`Error reading file stats: ${err.message}`);
                        return;
                    }

                    const matchesPattern = new RegExp(pattern).test(entry.name);
                    const matchesType = checkType(stats, fileType);

                    if (matchesPattern && matchesType) {
                        console.log(fullPath);
                    }

                    if (stats.isDirectory()) {
                        walk(fullPath); // Recurse into subdirectories
                    }
                });
            });
        });
    }

    walk(startPath);
}

/**
 * Main function to initiate the find operation with arguments.
 * @param {string[]} args - Command-line arguments for path, pattern, and type.
 */
function find(args) {
    if (args.length < 3) {
        console.error("Usage: find <path> <pattern> <type> (type should be 'file', 'directory', or 'link')");
        return;
    }

    const startPath = args[0];
    const pattern = args[1];
    const fileType = args[2];

    if (!['file', 'directory', 'link'].includes(fileType)) {
        console.error(`Invalid type: ${fileType}. Allowed types are 'file', 'directory', or 'link'.`);
        return;
    }

    findLogic(startPath, pattern, fileType);
}

module.exports =  find;
