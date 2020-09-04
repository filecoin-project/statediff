'use strict'

const renderer = require('../renderer');
const jsonPrinter = require('./jsonPrinter');
const cids = require('cids/src/index');

class jsonCid extends HTMLElement {
    constructor() {
        super();
        const shadowRoot = this.attachShadow({mode: 'open'})
        .appendChild(jsonCid.Template.content.cloneNode(true));

        this.path = this.getAttribute("data-path");
        this.cid = this.innerHTML;
        let c = new cids(this.cid);

        this.open = false;
 
        this.button = this.shadowRoot.children[0].children[0];
        this.deferred = this.shadowRoot.children[0].children[2];
        if (c.codec == "raw") {
            this.button.style.display = 'none';
            this.innerHTML = 'raw:' + new TextDecoder("utf-8").decode(c.bytes);
        } else {
            this.toggle = this.toggle.bind(this);
            this.button.addEventListener('click', this.toggle, false);    
        }
    }

    Render() {
        if (this.open) {
            renderer.FillSlot(this, 'deferred', jsonPrinter, this.cid, this.path)
        } else {
            renderer.FillTextSlot(this, 'deferred', '');
        }
    }

    toggle() {
        this.open = !this.open;
        if (this.open) {
            this.button.innerHTML = "▿";
            this.deferred.style.display = "block";
        } else {
            this.button.innerHTML = "▷";
            this.deferred.style.display = "none";
        }
        this.Render();
    }

    Close() {
        this.button.removeEventListener('click', this.toggle, false);
        this.innerHTML = "";
    }
}

jsonCid.Name = 'json-cid';

jsonCid.Template = `
<div style='display:inline-block;white-space:normal;'>
    <button style='border:0;background:transparent;padding:0;margin:0;'>▷</button>
    <div style='display:inline-block;'>
    CID: <slot></slot>
    </div>
    <div style='display:none; margin-left:2em;max-height:100vh;overflow-y:scroll;'>
        <slot name='deferred'></slot>
    </div>
</div>
`;

module.exports = jsonCid;
