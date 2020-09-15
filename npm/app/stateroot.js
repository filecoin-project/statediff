'use strict'

const lotusActor = require('./components/lotusActor');
const expander = require('./expander');
const renderer = require('./renderer');
const sparseArray = require('./sparsearray');
const store = require('./store');

// from https://github.com/filecoin-project/specs-actors/blob/master/actors/builtin/singletons.go
const ActorAddrs = {
    "System": "t00",
    "Init": "t01",
    "Reward": "t02",
    "Cron": "t03",
    "StoragePower": "t04",
    "StorageMarket": "t05",
    "VerifiedRegistry": "t06"
}

class stateroot {
    constructor(element, cid) {
        this.cid = cid;
        this.element = element;
        this.pn = 0;
        this.children = {};
        this.shownActors = {};

        element.innerHTML = "Loading: " + cid;
        let search = this.element.shadowRoot.children[0].querySelector('input');
        search.addEventListener('keyup', this.onSearch.bind(this));
        let next = this.element.shadowRoot.children[0].querySelector('.next');
        next.addEventListener('click', this.onClickNextButton.bind(this));
        this.Load();
    }

    Load() {
        this.load = store(this.cid, "stateRoot");
        this.load.then((r) => this.onStateroot(r));
        return this.load;
    }
    Ready() {
        return this.load;
    }

    onStateroot(resp) {
        if(resp[0]) {
            this.data = resp[1];
            this.Render();
        } else {
            this.element.innerHTML = "Failed to parse: " + resp[1];
        }
    }

    async onSearch(ev) {
        if (ev === true || ev.keyCode == 13) {
            this.filter = this.element.shadowRoot.querySelector('input').value;
            if (this.filter.indexOf('t') >= 0) {
                let initMap = await this.getInitMap();
                this.Render(initMap);
            } else {
                this.Render({});
            }
        }
    }

    onClickNextButton() {
        this.pn++;
        this.Render()
    }

    async getInitMap() {
        if (this.initMap) {
            return this.initMap;
        }
        let initHeadCid = this.data[ActorAddrs["Init"]]["Head"]["/"];
        let initState = await store(initHeadCid, "initActor");
        let initAddrMap = await store(initState[1]["AddressMap"]["/"], "initActorAddresses");
        this.initMap = initAddrMap;
        return initAddrMap;
    }

    Render(lookupMap) {
        let data = this.data;
        renderer.FillTextSlot(this.element, 'count', Object.keys(data).length);
        Object.keys(ActorAddrs).forEach((k) => {
            this.children[k] = renderer.FillSlot(this.element, k, expander, `${k} @ ${ActorAddrs[k]}`, lotusActor, [data[ActorAddrs[k]]]);
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
            // for explicit addresses.
            if (this.filter[0] == "t" || this.filter[0] == "f") {
                if (this.filter[1] == "0") {
                    let addr = this.filter.substr(2);

                    Object.keys(data).forEach((k) => {
                        if (k == addr) {
                            actors.push(k);
                        }
                    });
                }
            }
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
                this.shownActors[shownActors[i]] = new lotusActor(newActor, data[shownActors[i]]);
                results.appendChild(newActor);
            }
        }
        
        for (let i = 0; i < results.children.length; i++) {
            let ci = results.children[i].getAttribute("data-id");
            if (!shownActors.includes(ci)) {
                results.removeChild(results.children[i]);
                delete this.shownActors[ci];
                i--;
            }
        }

        this.element.shadowRoot.querySelector('.pn').innerHTML = (this.pn + 1);
        this.element.shadowRoot.querySelector('.pt').innerHTML = Math.ceil(actors.length / 10);
    }

    async GetState() {
        await this.Ready();
        let state = [];
        let sys = Object.keys(ActorAddrs).sort();
        for (let i = 0; i < sys.length; i++) {
            state.push(await this.children[sys[i]].GetState());
        }
        state.push(this.filter || '');
        state.push(this.pn || 0);
        let browse = Object.keys(this.shownActors).sort();
        for(let i = 0; i < browse.length; i++) {
            state.push(await this.shownActors[browse[i]].GetState());
        }
        return sparseArray.MaybeSparse(state);
    }

    async UpdateState(s) {
        if (s === 0 || s === undefined) {
            s = {};
        }

        await this.Ready();
        let sys = Object.keys(ActorAddrs).sort();
        let i = 0;
        for (i = 0; i < sys.length; i++) {
            let stateI = s[i] || 0;
            if (!await this.children[sys[i]].UpdateState(stateI)) {
                let addr = ActorAddrs[sys[i]];
                this.children[sys[i]] = renderer.RestoreSlot(this.element, sys[i], expander, stateI, [
                    `${k} @ ${addr}`,
                    lotusActor,
                    [this.data[addr]],
                ]);
            }
        }

        let dirty = false;
        let filter = s[i++] || "";
        if (filter != this.filter) {
            dirty = true;
            this.filter = filter;
            this.element.shadowRoot.querySelector('input').value = filter;
        }
        let pn = s[i++] || 0;
        if (pn != this.pn) {
            this.pn = pn;
            dirty = true;
        }
        if (dirty) {
            await this.onSearch(true);
        }
        let browse = Object.keys(this.shownActors).sort();
        for (let j = 0; j < browse.length; j++) {
            await this.shownActors[browse[j]].UpdateState(s[i + j]);
        }

        return true;
    }

    static async RestoreFromState(element, args, state) {
        let inst = new stateroot(element, args[0]);
        await inst.UpdateState(state);
        return inst;
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
