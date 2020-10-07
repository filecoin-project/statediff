'use strict'

const jsonPrinter = require('./components/jsonPrinter');
const jsonCid = require('./components/jsonCid');
const sparseArray = require('./sparsearray');

class metaMsg {
    constructor(element, state) {
        this.element = element;
        let hint = this.element.getAttribute('data-id');
        if (hint != undefined && hint != null) {
            state.Address = hint;
        }
        let jp = new jsonCid();
        this.GetChildren = jp.GetChildren.bind(this);

        this.Render(state);
    }


    Render(data) {
        this.element.innerHTML = '<pre><json-cid data-path="msgMeta">' + data + '</json-cid></pre>';
    }

    Close() {
        this.element.innerHTML = "";
    }

    async GetState() {
        let s = await Promise.all(this.GetChildren(this.element).map(async (c) => await c.GetState()));
        return sparseArray.MaybeSparse(s);
    }

    async UpdateState(s) {
        if (s === 0 || s === undefined) {
            s = {};
        }
        await this.tick();
        let children = this.GetChildren(this.element);
        for (let i = 0; i < children.length; i++) {
            await children[i].UpdateState(s[i]);
        }
        return true;
    }

    tick() {
        return new Promise(r => setTimeout(r, 0));
    }

    static async RestoreFromState(element, args, state) {
        let inst = new metaMsg(element, args[0]);
        await inst.UpdateState(state);
        return inst;
    }
}

metaMsg.Template = `<slot></slot>`;

module.exports = metaMsg;
