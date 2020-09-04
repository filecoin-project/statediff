'use strict'

const genericBlock = require('./genericBlock');
const expander = require('../expander');
const renderer = require('../renderer');
const jsonPrinter = require('./jsonPrinter');


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

        this.Render(state);
    }


    Render(data) {
        this.element.innerHTML = '<pre>' + jsonPrinter.Stringify(data, data.Type || "") + '</pre>';
        /*
        let hint = this.element.getAttribute('data-id');
        if (hint) {
            hint = ' @' + hint;
        } else {
            hint = '';
        }
        if (type == undefined) {
            renderer.FillTextSlot(this.element, 'type', data.Code["/"] + hint);
        } else {
            renderer.FillTextSlot(this.element, 'type', type + hint);
        }
        renderer.FillHTMLSlot(this.element, 'head', `<json-cid data-path="${type}">${data.Head["/"]}</json-cid>`);

        renderer.FillTextSlot(this.element, 'nonce', data.Nonce);
        renderer.FillTextSlot(this.element, 'balance', data.Balance);
        */
    }

    Close() {
        this.element.innerHTML = "";
    }
}

lotusActor.Template = `<slot></slot>`;

lotusActor.Codes = codeMap;

module.exports = lotusActor;
