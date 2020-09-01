'use strict'

const cronActor = require('./cronActor');
const genericBlock = require('./genericBlock');
const expander = require('../expander');
const initActor = require('./initActor');
const renderer = require('../renderer');
const rewardActor = require('./rewardActor');
const storagePowerActor = require('./storagePowerActor');
const storageMarketActor = require('./storageMarketActor');
const systemActor = require('./systemActor');
const verifiedRegistryActor = require('./verifiedRegistryActor');


// encodings of https://github.com/filecoin-project/specs-actors/blob/master/actors/builtin/codes.go
const codeMap = {
    "bafkqaddgnfwc6mjpon4xg5dfnu": systemActor,
    "bafkqactgnfwc6mjpnfxgs5a": initActor,
    "bafkqaddgnfwc6mjpojsxoylsmq": rewardActor,
    "bafkqactgnfwc6mjpmnzg63q": cronActor,
    "bafkqaetgnfwc6mjpon2g64tbm5sxa33xmvza": storagePowerActor,
    "bafkqae3gnfwc6mjpon2g64tbm5sw2ylsnnsxi": storageMarketActor,
    "bafkqaftgnfwc6mjpozsxe2lgnfswi4tfm5uxg5dspe": verifiedRegistryActor,
    "bafkqafdgnfwc6mjpobqxs3lfnz2gg2dbnzxgk3a": undefined,
}
class lotusActor {
    constructor(element, state) {
        this.element = element;
        this.Render(state);
    }


    Render(data) {
        let type = codeMap[data.Code["/"]];
        if (type == undefined) {
            renderer.FillTextSlot(this.element, 'type', data.Code["/"]);
            renderer.FillSlot(this.element, 'head', expander, data.Head["/"], genericBlock);
        } else {
            renderer.FillTextSlot(this.element, 'type', type.Name);
            renderer.FillSlot(this.element, 'head', expander, data.Head["/"], type);
        }

        renderer.FillTextSlot(this.element, 'nonce', data.Nonce);
        renderer.FillTextSlot(this.element, 'balance', data.Balance);
    }

    Close() {
        this.element.innerHTML = "";
    }
}

lotusActor.Template = `
<div style='display:inline-block;'>
<slot name='type'></slot> Nonce <slot name='nonce' class='json-number'></slot> Balance <slot name='balance'></slot><br />
<slot name='head'></slot>
</div>
`;

module.exports = lotusActor;
