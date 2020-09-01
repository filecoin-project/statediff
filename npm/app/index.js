const cronActor = require('./components/cronActor');
const expander = require('./expander');
const genericBlock = require('./components/genericBlock');
const initActor = require('./components/initActor');
const initActorTable = require('./components/initActorTable');
const lotusActor = require('./components/lotusActor');
const rewardActor = require('./components/rewardActor');
const stateroot = require('./stateroot');
const storagePowerActor = require('./components/storagePowerActor');
const storageMarketActor = require('./components/storageMarketActor');
const systemActor = require('./components/systemActor');
const tipset = require('./tipset');

const renderer = require('./renderer');
const verifiedRegistryActor = require('./components/verifiedRegistryActor');

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
    renderer.Register(cronActor);
    renderer.Register(expander);
    renderer.Register(genericBlock);
    renderer.Register(initActor);
    renderer.Register(initActorTable);
    renderer.Register(lotusActor);
    renderer.Register(rewardActor);
    renderer.Register(stateroot);
    renderer.Register(storagePowerActor);
    renderer.Register(storageMarketActor);
    renderer.Register(systemActor);
    renderer.Register(tipset);
    renderer.Register(verifiedRegistryActor);
}

window.addEventListener('load', onLoad);
