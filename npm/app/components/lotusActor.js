'use strict'

const jsonPrinter = require('./jsonPrinter');
const jsonCid = require('./jsonCid');


// encodings of https://github.com/filecoin-project/specs-actors/blob/master/actors/builtin/codes.go
const codeMap = {
    "bafkqaddgnfwc6mjpon4xg5dfnu": "systemActor",
    "bafkqactgnfwc6mjpnfxgs5a": "initActor",
    "bafkqaddgnfwc6mjpojsxoylsmq": "rewardActor",
    "bafkqactgnfwc6mjpmnzg63q": "cronActor",
    "bafkqaetgnfwc6mjpon2g64tbm5sxa33xmvza": "storagePowerActor",
    "bafkqae3gnfwc6mjpon2g64tbm5sw2ylsnnsxi": "storageMarketActor",
    "bafkqaftgnfwc6mjpozsxe2lgnfswi4tfm5uxg5dspe": "verifiedRegistryActor",
    "bafkqadlgnfwc6mjpmfrwg33vnz2a": "accountActor",
    "bafkqadtgnfwc6mjpnv2wy5djonuwo": "multisigActor",
    "bafkqafdgnfwc6mjpobqxs3lfnz2gg2dbnzxgk3a": "paymentChannelActor",
    "bafkqaetgnfwc6mjpon2g64tbm5sw22lomvza": "storageMinerActor",
}
class lotusActor {
    constructor(element, state) {
        this.element = element;
        let type = codeMap[state.Code["/"]];
        if (type != undefined) {
            state.Type = type;
        }
        let hint = this.element.getAttribute('data-id');
        if (hint != undefined && hint != null) {
            state.Address = hint;
        }
        let jp = new jsonCid();
        this.GetChildren = jp.GetChildren.bind(this);

        this.Render(state);
    }


    Render(data) {
        this.element.innerHTML = '<pre>' + jsonPrinter.Stringify(data, data.Type || "") + '</pre>';
    }

    Close() {
        this.element.innerHTML = "";
    }

    async GetState() {
        let s = await Promise.all(this.GetChildren(this.element).map(async (c) => await c.GetState()));
        return s;
    }

    async UpdateState(s) {
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
        let inst = new lotusActor(element, args[0]);
        await inst.UpdateState(state);
        return inst;
    }
}

lotusActor.Template = `<slot></slot>`;

lotusActor.Codes = codeMap;

module.exports = lotusActor;
