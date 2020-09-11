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
            this.cid = cid[0];
            store(cid[0], "tipset").then((r) => this.onTipset(r));
            return
        }

        element.innerHTML = "Loading: " + cid;
        this.Load();
    }

    Load() {
        this.load = store(this.cid, "tipset");
        this.load.then((r) => this.onTipset(r));
        return this.load;
    }
    Ready() {
        return this.load;
    }

    static async RestoreFromState(element, args, state) {
        let inst = new tipset(element, state[0]);
        await inst.Ready();
        await inst.UpdateState(state);
        return inst;
    }

    async GetState() {
        let state = [this.cid,0,0];
        if (this.stateroot) {
            state[1] = await this.stateroot.GetState();
        }
        if (this.parents) {
            state[2] = await this.parents.GetState();
        }
        return state;
    }

    async UpdateState(s) {
        if (s[0] != this.cid) {
            return false;
        }
        if (!this.stateroot || !this.parents) {
            await this.Load();
        }

        if (!await this.stateroot.UpdateState(s[1])) {
            renderer.RestoreSlot(this.element, 'stateroot', expander, s[1], [
                this.data.ParentStateRoot["/"],
                stateroot,
                [this.data.ParentStateRoot["/"]],
            ]);
        }

        if (!await this.parents.UpdateState(s[2])) {
            renderer.RestoreSlot(this.element, 'parents', expander, s[2], [
                this.data.Parents["/"],
                tipset,
                [this.data.Parents[0]["/"]],
            ]);
        }

        return true;
    }

    onTipset(resp) {
        if(resp[0]) {
            this.data = resp[1];
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
        this.stateroot = renderer.FillSlot(this.element, 'stateroot', expander, data.ParentStateRoot["/"], stateroot, [data.ParentStateRoot["/"]]);

        // todo: other parents?
        this.parents = renderer.FillSlot(this.element, 'parents', expander, data.Parents[0]["/"], tipset, [data.Parents[0]["/"]]);
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
