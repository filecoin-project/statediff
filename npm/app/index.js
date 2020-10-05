const expander = require('./expander');
const genericBlock = require('./components/genericBlock');
const jsonPrinter = require('./components/jsonPrinter');
const jsonBitfield = require('./components/jsonBitfield');
const jsonCid = require('./components/jsonCid');
const lotusActor = require('./components/lotusActor');
const msgMeta = require('./msgmeta');
const stateroot = require('./stateroot');
const tipset = require('./tipset');
const changeEvent = require('./event');

const renderer = require('./renderer');

let renderedRoot = null;

async function GetCurrentRoot() {
    let data = await fetch("head").then(async (r) => {
        let cids = await r.json()
        return cids[0]["/"];
    });
    return data;
}

async function GetHeight(h) {
    let data = await fetch("height?h=" + h).then((r) => {
        return r.text();
    });
    return data;
}

function setup(rootEl, rootcid) {
    renderedRoot = renderer.FillSlot(rootEl, 'root', tipset, rootcid);
    changeEvent();
}

async function GetState() {
    if(renderedRoot == null) {
        return "";
    }
    return await renderedRoot.GetState();
}

async function UpdateState(s) {
    let rootEl = document.getElementById('root');
    if(s == null || s == "") {
        renderer.FillTextSlot(rootEl, 'root', '');
        renderedRoot = null;
        return;
    }
    if (renderedRoot == null || !await renderedRoot.UpdateState(s)) {
        renderedRoot = await renderer.RestoreSlot(rootEl, 'root', tipset, s);
    }
}

async function currentClicked() {
    if (renderedRoot != null) {
        renderedRoot.Close();
    }
    rootCid = await GetCurrentRoot()
    setup(document.getElementById('root'), rootCid);
}

function stateRootEnter(ev) {
    if (ev.keyCode != 13) {
        return
    }
    if (renderedRoot != null) {
        renderedRoot.Close();
    }
    let val = document.getElementById('stateroot').value;
    if (val.length < 15 && !isNaN(+val)) {
        GetHeight(val).then((sr) => {
            setup(document.getElementById('root'), sr);
        });
    } else {
        setup(document.getElementById('root'), val);
    }
}

let onLoad = () => {
    document.getElementById('current').addEventListener('click', currentClicked);
    document.getElementById('stateroot').addEventListener('keypress', stateRootEnter);
    renderer.Register(expander);
    renderer.Register(genericBlock);
    renderer.Register(jsonBitfield);
    renderer.Register(jsonCid);
    renderer.Register(jsonPrinter);
    renderer.Register(lotusActor);
    renderer.Register(msgMeta);
    renderer.Register(stateroot);
    renderer.Register(tipset);
    document.addEventListener(changeEvent.Event.type, async () => {
        let newState = await GetState();
        if (newState != history.state) {
            history.pushState(newState, "", "?state=" + btoa(JSON.stringify(newState)) + window.location.hash);
        }
    });
    window.onpopstate = (ev) => {
        if (!ev.state) {
            // this happens e.g. when clicking on the anchor links.
            GetState().then(s => {
                history.replaceState(s, "", "?state=" + btoa(JSON.stringify(s)) + window.location.hash);
            });
        } else {
            UpdateState(ev.state);
        }
    };
    if (window.location.search.indexOf("state=") > -1) {
        let str = decodeURIComponent(window.location.search.split("state=").slice(1).join("state="));
        str = atob(str);
        UpdateState(JSON.parse(str)).then(() => {
            if (window.location.hash != "") {
                window.location.hash = window.location.hash;
            }
        });
    }
}

window.addEventListener('load', onLoad);
