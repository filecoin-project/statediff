'use strict'

const atlas = {
    "initActor.AddressMap": "initActorAddresses",
    "storagePowerActor.CronEventQueue": "storagePowerCronEventQueue",
    "storagePowerActor.Claims": "storagePowerClaims",
};

// Returns type hints for interpretation of CIDs / hamts
function Get(path) {
    if (path.endsWith('Actor')) {
        return path;
    }
    if (atlas[path] != undefined) {
        return atlas[path];
    }
    console.log(path);
    return path;
}

module.exports.Get = Get;
