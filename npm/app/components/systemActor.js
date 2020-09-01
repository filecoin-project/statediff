'use strict'

const expander = require('../expander');
const renderer = require('../renderer');

class systemActor {
    constructor(element, cid) {
        this.element = element;
        this.Render(cid);
    }

    Render(data) {

    }

    Close() {
        this.element.innerHTML = "";
    }
}

systemActor.Name = 'System Actor';

systemActor.Template = `
<div style='display:inline-block;'>
The system actor has no state
</div>
`;

module.exports = systemActor;
