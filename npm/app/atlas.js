'use strict'

// Returns type hints for interpretation of CIDs / hamts
function Get(path) {
    if (path.endsWith('Actor')) {
        return path;
    }
    console.log(path);
    return "unknown";
}

module.exports.Get = Get;
