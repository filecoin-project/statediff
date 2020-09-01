// GetCid loads the JSON for a given cid and type hint.
async function GetCid(cid, as) {
    let data = await fetch("/cid?cid=" + cid + "&as=" + as).then(async (response) => {
        const contentType = response.headers.get('content-type');
        if (!contentType || !contentType.includes('application/json')) {
            return [false, await response.text()];
        }
        return [true, await response.json()];
    });
    return data;
}

module.exports = GetCid;
