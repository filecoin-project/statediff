const expander = require('./expander');
const genericBlock = require('./components/genericBlock');
const jsonPrinter = require('./components/jsonPrinter');
const jsonCid = require('./components/jsonCid');
const lotusActor = require('./components/lotusActor');
const stateroot = require('./stateroot');
const tipset = require('./tipset');

const renderer = require('./renderer');

let renderedRoot = null;

async function GetCurrentRoot() {
    let data = await fetch("/head").then((r) => {
        return r.text();
    });
    return data;
}

function setup(rootEl, rootcid) {
    renderedRoot = renderer.FillSlot(rootEl, 'root', tipset, rootcid);
}

async function currentClicked() {
    if (renderedRoot != null) {
        renderedRoot.Close();
    }
    rootCid = await GetCurrentRoot()
    renderedRoot = setup(document.getElementById('root'), rootCid);
}

function stateRootEnter() {
    if (renderedRoot != null) {
        renderedRoot.Close();
    }
    renderedRoot = setup(document.getElementById('root'), document.getElementById('stateroot').value);
}

let onLoad = () => {
    document.getElementById('current').addEventListener('click', currentClicked);
    document.getElementById('stateroot').addEventListener('keypress', stateRootEnter);
    renderer.Register(expander);
    renderer.Register(genericBlock);
    renderer.Register(jsonCid);
    renderer.Register(jsonPrinter);
    renderer.Register(lotusActor);
    renderer.Register(stateroot);
    renderer.Register(tipset);
}

window.addEventListener('load', onLoad);
