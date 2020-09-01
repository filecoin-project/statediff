'use strict'

const store = require('./store');
const renderer = require('./renderer');

class expander {
    constructor(element, cid, as) {
        this.element = element;
        this.cid = cid;
        this.as = as;
        this.Render();
        this.open = false;
 
        this.button = this.element.shadowRoot.children[0].children[0];
        this.deferred = this.element.shadowRoot.children[0].children[3];
        this.toggle = this.toggle.bind(this);
        this.button.addEventListener('click', this.toggle, false);
    }

    Render() {
        renderer.FillTextSlot(this.element, 'cid', this.cid);
        if (this.open) {
            renderer.FillSlot(this.element, 'deferred', this.as, this.cid)
        } else {
            renderer.FillTextSlot(this.element, 'deferred', '');
        }
    }

    toggle() {
        this.open = !this.open;
        if (this.open) {
            this.button.innerHTML = "▿";
            this.deferred.style.display = "inline-block";
        } else {
            this.button.innerHTML = "▷";
            this.deferred.style.display = "none";
        }
        this.Render();
    }

    Close() {
        this.button.removeEventListener('click', this.toggle, false);
        this.element.innerHTML = "";
    }
}

expander.Template = `
<div style='display;inline-block;'>
    <button style='border:0;background:transparent;'>▷</button>
    <slot name='cid'></slot>
    <br />
    <div style='display:none; margin-left:2em;'>
        <slot name='deferred'></slot>
    </div>
</div>
`;


module.exports = expander;
