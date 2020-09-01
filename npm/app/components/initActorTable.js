'use strict'

const renderer = require('../renderer');
const store = require('../store');
const tableify = require('tableify');

class initActorTable {
    constructor(element, cid) {
        this.element = element;

        store(cid, "initActorAddresses").then((r) => this.Render(r));
    }

    Render(data) {
        if (data[0]) {
            let table = tableify(data[1]);
            renderer.FillHTMLSlot(this.element, 'table', table);
        } else {
            renderer.FillTextSlot(this.element, 'table', data[1]);
        }
    }

    Close() {
        this.element.innerHTML = "";
    }
}

initActorTable.Template = `
<div style='display:inline-block; max-height:15em; overflow:scroll'>
<slot name='table'></slot>
</div>
`;

module.exports = initActorTable;
