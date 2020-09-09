'use strict'

const renderer = require('../renderer');

class jsonBitfield extends HTMLElement {
    constructor() {
        super();
        const shadowRoot = this.attachShadow({mode: 'open'})
        .appendChild(jsonBitfield.Template.content.cloneNode(true));

        let data = this.innerHTML.trim().toLowerCase();
        let pairs = data.match(/[\da-f]{2}/gi);
        let dec = pairs.map(s => {
            return parseInt(s, 16);
        });
        this.field = new Uint8Array(dec);

        this.canvas = this.shadowRoot.children[0];
        this.canvas.height = 8;
        this.canvas.width = this.field.byteLength;
        this.Render();
    }

    Render() {
        let ctx = this.canvas.getContext('2d');
        let pixels = ctx.getImageData(0, 0, this.field.byteLength, 8);
        for (let x = 0; x < this.field.byteLength; x++) {
            for (let y = 0; y < 8; y++) {
                let pixel = 4*((this.field.byteLength) * y + x);
                if (this.field[x] & (1 << y)) {
                    pixels.data[pixel] = 0;
                } else {
                    pixels.data[pixel] = 255;
                }
                pixels.data[pixel+1] = pixels.data[pixel];
                pixels.data[pixel+2] = pixels.data[pixel];
                pixels.data[pixel+3] = 255;
            }
        }
        ctx.putImageData(pixels, 0, 0);
    }

    Close() {
        this.innerHTML = "";
    }
}

jsonBitfield.Name = 'json-bitfield';

jsonBitfield.Template = `<canvas style='display:inline-block;'/>`;

module.exports = jsonBitfield;
