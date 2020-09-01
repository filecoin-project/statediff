'use strict'

const json = require('pretty-print-json');
const expander = require('../expander');
const renderer = require('../renderer');
const store = require('../store');
const initActorTable = require('./initActorTable');

class initActor {
    constructor(element, cid) {
        this.element = element;

        store(cid, "initActor").then((r) => this.Render(r));
    }

    Render(data) {
        if (data[0]) {
            renderer.FillHTMLSlot(this.element, 'data', "<pre>" + json.toHtml(data[1], {indent: 2}) + "</pre>");
            renderer.FillSlot(this.element, 'addresses', expander, data[1].AddressMap["/"], initActorTable);
        } else {
            renderer.FillTextSlot(this.element, 'data', data[1]);
        }
    }

    Close() {
        this.element.innerHTML = "";
    }
}

initActor.Name = 'Init Actor';

initActor.Template = `
<div style='display:inline-block;'>
<slot name='data'/></slot>
<slot name='addresses'/></slot>
</div>
`;

module.exports = initActor;
