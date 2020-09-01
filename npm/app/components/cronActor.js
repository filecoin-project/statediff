'use strict'

const json = require('pretty-print-json');
const renderer = require('../renderer');
const store = require('../store');

class cronActor {
    constructor(element, cid) {
        this.element = element;

        store(cid, "cronActor").then((r) => this.Render(r));
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

cronActor.Name = 'Cron Actor';

cronActor.Template = `
<div style='display:inline-block;'>
<slot name='data'/></slot>
</div>
`;

module.exports = cronActor;
