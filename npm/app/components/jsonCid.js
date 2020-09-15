'use strict'

const changeEvent = require('../event');
const renderer = require('../renderer');
const sparseArray = require('../sparsearray');
const jsonPrinter = require('./jsonPrinter');
const cids = require('cids/src/index');


class jsonCid extends HTMLElement {
    constructor() {
        super();
        const shadowRoot = this.attachShadow({mode: 'open'})
        .appendChild(jsonCid.Template.content.cloneNode(true));

        this.path = this.getAttribute("data-path");
        this.cid = this.innerHTML;
        if (!this.cid) {
            return;
        }
        let c = new cids(this.cid);

        this.open = false;
 
        this.button = this.shadowRoot.children[0].children[0];
        this.deferred = this.shadowRoot.children[0].children[2];
        if (c.codec == "raw") {
            this.button.style.display = 'none';
            this.innerHTML = 'raw:' + new TextDecoder("utf-8").decode(c.bytes);
        } else if (c.codec == "fil-commitment-sealed") {
            this.button.style.display = 'none';
            this.innerHTML = 'Sealed Commitment:' + c.toString();
        } else {
            this.doToggle = this.doToggle.bind(this);
            this.button.addEventListener('click', this.doToggle, false);    
        }
    }

    Render() {
        if (this.open) {
            return renderer.FillSlot(this, 'deferred', jsonPrinter, this.cid, this.path)
        } else {
            return renderer.FillTextSlot(this, 'deferred', '');
        }
    }

    doToggle() {
        this.toggle();
        changeEvent();
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
        return this.Render();
    }
    async GetState() {
        if (!this.open) {
            return 0;
        }

        let s = await Promise.all(this.GetChildren(this).map(async (c) => await c.GetState()));
        return sparseArray.MaybeSparse(s);
    }

    GetChildren(e) {
        let l = [];
        if (!e.children || !e.children.length) {
            return l;
        }
        for (let i = 0; i < e.children.length; i++) {
            if (e.children[i].GetState) {
                l.push(e.children[i]);
            } else if (e.children[i].children && e.children[i].children.length) {
                l = l.concat(this.GetChildren(e.children[i]));
            }
        }
        return l;
    }

    tick() {
        return new Promise(r => setTimeout(r, 0));
    }
    async UpdateState(s) {
        if (s === undefined) {
            s = 0;
        }
        if (s === 0 && this.open) {
            this.toggle();
            return true;
        } else if (s === 0) {
            return true;
        } else if (!this.open) {
            let prom = this.toggle();
            if (prom instanceof Promise) {
                await prom;
            } else if (prom.Ready) {
                await prom.Ready();
            }
            await this.tick();
        }
        // wait a tick for dom to populate
        let children = this.GetChildren(this);
        for (let i = 0; i < children.length; i++) {
            await children[i].UpdateState(s[i] || 0);
        }

        return true;
    }

    Close() {
        this.button.removeEventListener('click', this.doToggle, false);
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
