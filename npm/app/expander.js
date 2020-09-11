'use strict'

const changeEvent = require('./event');
const renderer = require('./renderer');

class expander {
    constructor(element, label, cls, args) {
        this.element = element;
        this.label = label;
        this.cls = cls;
        this.args = args;
        this.Render();
        this.open = false;
 
        this.button = this.element.shadowRoot.children[0].children[0];
        this.deferred = this.element.shadowRoot.children[0].children[2];
        this.toggle = this.toggle.bind(this);
        this.button.addEventListener('click', this.toggle, false);
    }

    Render() {
        renderer.FillTextSlot(this.element, 'label', this.label);
        if (this.open) {
            this.child = renderer.FillSlotArgs(this.element, 'deferred', this.cls, false, this.args)
        } else {
            renderer.FillTextSlot(this.element, 'deferred', '');
        }
    }

    static async RestoreFromState(element, args, state) {
        let inst = new expander(element, args[0], args[1], args[2]);
        await inst.UpdateState(state);
        return inst;
    }

    async GetState() {
        if(!this.open) {
            return 0;
        } else {
            return [await this.child.GetState()];
        }
    }

    async UpdateState(s) {
        if (s===0) {
            if (this.open) {
              this.doToggle();
            }
            return true;
        } else if (!this.open) {
            this.doToggle();
        }
        if (this.child && await this.child.UpdateState(s[0])) {
            return true;
        }
        return false;
    }

    doToggle() {
        this.open = !this.open;
        if (this.open) {
            this.button.innerHTML = "▿";
            this.deferred.style.display = "block";
        } else {
            this.button.innerHTML = "▷";
            this.deferred.style.display = "none";
        }
        this.Render();
    }
    toggle() {
        this.doToggle();
        changeEvent();
    }

    Close() {
        this.button.removeEventListener('click', this.toggle, false);
        this.element.innerHTML = "";
    }
}

expander.Template = `
<div style='display;inline-block;'>
    <button style='border:0;background:transparent;'>▷</button>
    <slot name='label'></slot>
    <div style='display:none; margin-left:2em;'>
        <slot name='deferred'></slot>
    </div>
</div>
`;


module.exports = expander;
