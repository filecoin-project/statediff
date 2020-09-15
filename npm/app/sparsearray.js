'use strict'

// encode an array as an object when smaller.
function MaybeSparse(arr) {
    let obj = {};
    for (let x = 0; x < arr.length; x++) {
        if(arr[x] !== 0) {
            obj[x] = arr[x];
        }
    }
    if (Object.keys(obj).length === 0) {
        return {};
    }
    if (JSON.stringify(obj).length < JSON.stringify(arr).length) {
        return obj;
    }
    return arr;
}

module.exports.MaybeSparse = MaybeSparse;
