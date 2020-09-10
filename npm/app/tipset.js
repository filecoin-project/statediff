'use strict'

const expander = require('./expander');
const renderer = require('./renderer');
const stateroot = require('./stateroot');
const store = require('./store');

class tipset {
    constructor(element, cid) {
        this.element = element;

        cid = cid.trim().replace("{", "").replace("}", "");
        this.cid = cid;
        if (cid.indexOf(",") > 1) {
            cid = cid.split(",");
            // here we know we have a tipset.
            element.innerHTML = "Loading: " + cid[0];
            store(cid[0], "tipset").then((r) => this.onTipset(r));
            return
        }

        element.innerHTML = "Loading: " + cid;
        store(cid, "tipset").then((r) => this.onTipset(r));
    }

    onTipset(resp) {
        if(resp[0]) {
            this.Render(resp[1]);
        } else {
            renderer.FillSlot(this.element, 'stateroot', expander, this.cid, stateroot);
        }
    }

    Render(data) {
        renderer.FillTextSlot(this.element, 'miner', data.Miner);
        renderer.FillTextSlot(this.element, 'height', data.Height);
        renderer.FillTextSlot(this.element, 'date', new Date(data.Timestamp * 1000));
        renderer.FillTextSlot(this.element, 'basefee', data.ParentBaseFee);
        renderer.FillSlot(this.element, 'stateroot', expander, data.ParentStateRoot["/"], stateroot, [data.ParentStateRoot["/"]]);

        // todo: other parents?
        renderer.FillSlot(this.element, 'parents', expander, data.Parents[0]["/"], tipset, [data.Parents[0]["/"]]);
    }

    Close() {
        this.element.innerHTML = "";
    }
}

tipset.Template = `
<div style='display:inline-block;'>
Mined by <slot name='miner'></slot> at height <slot name='height'></slot> at <slot name='date'></slot><br />
Base fee: <slot name='basefee'></slot><br />
State Root: <slot name='stateroot'></slot><br />
Parents: <slot name='parents'></slot>
</div>
`;

module.exports = tipset;
