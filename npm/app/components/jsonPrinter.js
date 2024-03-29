'use strict'

const renderer = require('../renderer');
const atlas = require('../atlas');
const store = require('../store');

class jsonPrinter {
    constructor(element, cid, at) {
        this.element = element;
        this.at = at;

        if (cid != undefined) {
            let as = atlas.Get(at);
            this.ready = store(cid, as).then((r) => this.Render(r));
        }
    }

    Ready() {
        return this.ready;
    }

    Render(data) {
        if (data[0]) {
            renderer.FillHTMLSlot(this.element, 'data', "<pre>" + this.stringify(data[1], this.at) + "</pre>");
        } else {
            renderer.FillTextSlot(this.element, 'data', data[1]);
        }
    }

    indent(str) {
        return str.split('\n').map((l) => {return '  ' + l}).join('\n');
    }
    stringify(obj, path, full) {
        if (typeof obj == "string") {
            obj = obj.replace(/&/g,   '&amp;')
            .replace(/\\"/g, '&bsol;&quot;')
            .replace(/</g,   '&lt;')
            .replace(/>/g,   '&gt;');
            return `"<span class=json-string>${obj}</span>"`;
        }
        if ([true, false].includes(obj)) {
            return `<span class=json-boolean>${obj}</span>`;
        }
        if (obj == null) {
            return `<span class=json-null>null</span>`;
        }
        if (typeof obj == "number") {
            return `<span class=json-number>${obj}</span>`;
        }
        if (typeof obj == "object") {
            // array.
            let str = "";
            if (obj.__proto__.constructor.name == "Array") {
                if (obj.length == 0) {
                    return '[]';
                }
                // if all entries are numbers, we'll skip returns.
                if (obj.every(n => typeof n == "number")) {
                    str += "[";
                    for (let i = 0; i < obj.length; i++) {
                        if (i > 0) {
                            str += ", ";
                        }
                        str += this.stringify(obj[i], path + "[" + i + "]")
                        if (i > 100 && !full) {
                            str += ` ... and ${obj.length -i} more`;
                            break
                        }
                    }
                    str += "]";
                } else {
                    str += "[\n";
                    for (let i = 0; i < obj.length; i++) {
                        if (i > 0) {
                            str += ",\n";
                        }
                        str += this.indent(this.stringify(obj[i], path + "[" + i + "]"));
                        if (i > 100 && !full) {
                            str += this.indent(` ... and ${obj.length - i} more`);
                            break;
                        }
                    }
                    str += "\n]";
                }
            } else if (Object.keys(obj).length == 1 && typeof obj["/"] == "string") {
                // cid special case.
                str += `<json-cid data-path="${path}">${obj["/"]}</json-cid>`;
            } else if (Object.keys(obj).length == 2 && obj["_type"] == "bitfield") {
                // bitfield special case
                str += `<json-bitfield>${obj["bytes"]}</json-bitfield>`;
            } else if (Object.keys(obj).length == 0) {
                return '{}';
            } else {
                // object.
                str += "{\n";
                let keys = Object.keys(obj);
                for (let i = 0; i < keys.length; i++) {
                    if (i > 0) {
                        str += ",\n";
                    }
                    let subPath = path + "." + keys[i];
                    str += this.indent(`"<span class=json-key><a name="${subPath}" href="#${subPath}">${keys[i]}</a></span>" : ` +
                        this.stringify(obj[keys[i]], subPath))
                    if (i > 100 && !full) {
                        str += this.indent(` ... and ${keys.length - i} more`);
                        break;
                    }    
                }
                str += "\n}";
            }
            return str;
        }
        return obj;
    }

    Close() {
        this.element.innerHTML = "";
    }
}

jsonPrinter.Template = `
<div style='display:inline-block;'>
<slot name='data'/></slot>
</div>
`;

let proto = (new jsonPrinter());
jsonPrinter.Stringify = proto.stringify.bind(proto);

module.exports = jsonPrinter;
