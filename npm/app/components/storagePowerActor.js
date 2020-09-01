'use strict'

const json = require('pretty-print-json');
const renderer = require('../renderer');
const store = require('../store');

class storagePowerActor {
    constructor(element, cid) {
        this.element = element;

        store(cid, "storagePowerActor").then((r) => this.Render(r));
    }

    Render(data) {
        if (data[0]) {
            renderer.FillHTMLSlot(this.element, 'data', "<pre>" + json.toHtml(data[1], {indent: 2}) + "</pre>");
        } else {
            renderer.FillTextSlot(this.element, 'data', data[1]);
        }
    }

    Close() {
        this.element.innerHTML = "";
    }
}

storagePowerActor.Name = 'Storage Power Actor';

storagePowerActor.Template = `
<div style='display:inline-block;'>
<slot name='data'/></slot>
</div>
`;

module.exports = storagePowerActor;
