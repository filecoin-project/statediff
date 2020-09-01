'use strict'

const store = require('../store');
const renderer = require('../renderer');

class genericBlock {
    constructor(element, cid) {
        this.element = element;

        store(cid, "unknown").then((r) => this.Render(r));
    }


    Render(data) {
        renderer.FillTextSlot(this.element, 'data', JSON.stringify(data[1]));
    }

    Close() {
        this.element.innerHTML = "";
    }
}

genericBlock.Template = `
<div style='display:inline-block;'>
<slot name='data'></slot>
</div>
`;

module.exports = genericBlock;
