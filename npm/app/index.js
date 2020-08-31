const StateRootRenderer = require('./stateroot')

let renderedRoot = null;

async function GetCurrentRoot() {
    let data = await fetch("/head").then((r) => {
        return r.text();
    });
    return data;
}

function setup(rootEl, rootcid) {
    renderedRoot = StateRootRenderer(rootEl, rootcid);
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
}

window.addEventListener('load', onLoad);
