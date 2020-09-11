function FillTextSlot(el, name, str) {
    let textNode = document.createElement("span");
    textNode.slot = name;
    textNode.innerText = str;
    let existing = el.querySelector('[slot="' + name + '"]');
    if (existing != undefined) {
        el.removeChild(existing);
    }
    el.appendChild(textNode);
}

function FillHTMLSlot(el, name, fragment) {
    let node = document.createElement("span");
    node.slot = name;
    node.innerHTML = fragment;
    let existing = el.querySelector('[slot="' + name + '"]');
    if (existing != undefined) {
        el.removeChild(existing);
    }
    el.appendChild(node);
}

function FillSlot(el, name, cls) /*, args to cls... */ {
    return FillSlotArgs(el, name, cls, false, Array.prototype.slice.call(arguments, 3));
}

async function RestoreSlot(el, name, cls, state, args) {
    return await FillSlotArgs(el, name, cls, state, args);
}

function FillSlotArgs(el, name, cls, restore, args) {
    let node = document.createElement('filecoin-' + cls.name.toLowerCase());
    node.style.display = 'inline-block';
    node.style.verticalAlign = 'top';
    node.slot = name;

    let inst;
    if (restore) {
        inst = cls.RestoreFromState(node, args, restore);
    } else {
        args = [node].concat(args);
        // eww (would be better w/ es6 rest args, per)
        // https://stackoverflow.com/questions/3871731/dynamic-object-construction-in-javascript
        inst = new (Function.prototype.bind.apply(cls, [null].concat(args)))();
    }

    let existing = el.querySelector('[slot="' + name + '"]');
    if (existing != undefined) {
        el.removeChild(existing);
    }
    el.appendChild(node);

    return inst;
}

// Register loads at template for a given class.
function Register(cls) {
    let template = cls.Template;
    let templEl = document.createElement('template');
    templEl.innerHTML = template;

    let name = cls.Name || cls.name.toLowerCase();
    if (name.indexOf('-') == -1) {
        name = 'filecoin-' + name;
    }

    if (cls.prototype instanceof HTMLElement) {
        cls.Template = templEl;
        customElements.define(name, cls);
        return;
    }

    customElements.define(name,
        class extends HTMLElement {
            constructor() {
                super();

                const shadowRoot = this.attachShadow({mode: 'open'})
                .appendChild(templEl.content.cloneNode(true));
            }
        }
    );
}

module.exports = {
    FillTextSlot: FillTextSlot,
    FillHTMLSlot: FillHTMLSlot,
    FillSlot: FillSlot,
    FillSlotArgs: FillSlotArgs,
    RestoreSlot: RestoreSlot,
    Register: Register,
};
