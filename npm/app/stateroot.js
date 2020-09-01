'use strict'

const lotusActor = require('./components/lotusActor');
const renderer = require('./renderer');
const store = require('./store');

// from https://github.com/filecoin-project/specs-actors/blob/master/actors/builtin/singletons.go
const ActorAddrs = {
    "System": "0000",
    "Init": "0001",
    "Reward": "0002",
    "Cron": "0003",
    "StoragePower": "0004",
    "StorageMarket": "0005",
    "VerifiedRegistry": "0006"
}

class stateroot {
    constructor(element, cid) {
        this.element = element;

        element.innerHTML = "Loading: " + cid;
        store(cid, "stateRoot").then((r) => this.onStateroot(r));
    }

    onStateroot(resp) {
        if(resp[0]) {
            this.Render(resp[1]);
        } else {
            this.element.innerHTML = "Failed to parse: " + resp[1];
        }
    }

    Render(data) {
        renderer.FillTextSlot(this.element, 'count', Object.keys(data).length);
        Object.keys(ActorAddrs).forEach((k) => {
            renderer.FillSlot(this.element, k, lotusActor, data[ActorAddrs[k]]);
        });
    }

    Close() {
        this.element.innerHTML = "";
    }
}

stateroot.Template = `
<div style='display:inline-block;'>
System Actors:<br />
<slot name='System'></slot><br />
<slot name='Init'></slot><br />
<slot name='Reward'></slot><br />
<slot name='Cron'></slot><br />
<slot name='StoragePower'></slot><br />
<slot name='StorageMarket'></slot><br />
<slot name='VerifiedRegistry'></slot><br />
<slot name='BurntFunds'></slot><br />
(<slot name='count'></slot> total)<br />
<input type='search' /><br />
<slot name='results'></slot><br />
</div>
`;

module.exports = stateroot;
