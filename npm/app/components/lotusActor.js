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
    "bafkqadtgnfwc6mjpnv2wy5djonuwo": undefined,
    "bafkqafdgnfwc6mjpobqxs3lfnz2gg2dbnzxgk3a": undefined,
}
class lotusActor {
    constructor(element, state) {
        this.element = element;
        this.Render(state);
    }


    Render(data) {
        let type = codeMap[data.Code["/"]];
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
    }

    Close() {
        this.element.innerHTML = "";
    }
}

lotusActor.Template = `
<div style='display:inline-block;'>
<slot name='type'></slot> {<span style='color:brown'>Nonce</span>: <slot name='nonce'></slot>, <span style='color:brown;'>Balance</span>: <slot name='balance'></slot>}<br />
<slot name='head'></slot>
</div>
`;

module.exports = lotusActor;
