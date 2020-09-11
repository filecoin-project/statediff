'use strict'

const changeEvent = new Event('stateChange');

let queuedEvent = false;

module.exports = () => {
    if (!queuedEvent) {
        queuedEvent = true;
        setTimeout(() => {
            queuedEvent = false;
            document.dispatchEvent(changeEvent);
        }, 0);
    }
}
module.exports.Event = changeEvent;
