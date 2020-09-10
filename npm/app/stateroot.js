'use strict'

const lotusActor = require('./components/lotusActor');
const expander = require('./expander');
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
        this.pn = 0;

        element.innerHTML = "Loading: " + cid;
        let search = this.element.shadowRoot.children[0].querySelector('input');
        search.addEventListener('keyup', this.onSearch.bind(this));
        let next = this.element.shadowRoot.children[0].querySelector('.next');
        next.addEventListener('click', this.onClickNextButton.bind(this));
        store(cid, "stateRoot").then((r) => this.onStateroot(r));
    }

    onStateroot(resp) {
        if(resp[0]) {
            this.data = resp[1];
            this.Render();
        } else {
            this.element.innerHTML = "Failed to parse: " + resp[1];
        }
    }

    onSearch(ev) {
        if (ev.keyCode == 13) {
            this.filter = this.element.shadowRoot.querySelector('input').value;
            this.Render()
        }
    }

    onClickNextButton() {
        this.pn++;
        this.Render()
    }

    Render() {
        let data = this.data;
        renderer.FillTextSlot(this.element, 'count', Object.keys(data).length);
        Object.keys(ActorAddrs).forEach((k) => {
            renderer.FillSlot(this.element, k, expander, k, lotusActor, [data[ActorAddrs[k]]]);
        });
        let actors = [];
        if (!this.filter) {
            actors = Object.keys(data).filter((k) => {
                return !Object.values(ActorAddrs).includes(k);
            });
        } else {
            let matchedCode = "";
            Object.keys(lotusActor.Codes).forEach((k) => {
                if (lotusActor.Codes[k].indexOf(this.filter) > -1) {
                    matchedCode = k;
                }
            });
            Object.keys(data).forEach((k) => {
                if (k == this.filter ||
                    data[k].Code["/"] == this.filter ||
                    data[k].Code["/"] == matchedCode ||
                    data[k].Head["/"] == this.filter) {
                    actors.push(k);
                    return
                }
            })
        }

        let shownActors = actors.slice(this.pn * 10, (this.pn + 1) * 10);

        let results = this.element.querySelector('[slot="results"]');
        if (!results) {
            results = document.createElement("ul");
            results.slot = 'results';
            this.element.appendChild(results);
        }
        for(let i = 0; i < shownActors.length; i++) {
            if (!results.querySelector(`[data-id="${shownActors[i]}"]`)) {
                let newActor = document.createElement('filecoin-' + lotusActor.name.toLowerCase());
                newActor.style.display = 'list-item';
                newActor.style.verticalAlign = 'top';
                newActor.setAttribute('data-id', shownActors[i]);
                new lotusActor(newActor, data[shownActors[i]]);
                results.appendChild(newActor);
            }
        }
        
        for (let i = 0; i < results.children.length; i++) {
            let ci = results.children[i].getAttribute("data-id");
            if (!shownActors.includes(ci)) {
                results.removeChild(results.children[i]);
                i--;
            }
        }

        this.element.shadowRoot.querySelector('.pn').innerHTML = (this.pn + 1);
        this.element.shadowRoot.querySelector('.pt').innerHTML = Math.ceil(actors.length / 10);
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
<input type='search' />(page <span class='pn'></span> of <span class='pt'></span>)<button class='next'>Next</button><br />
<slot name='results'></slot><br />
</div>
`;

module.exports = stateroot;
