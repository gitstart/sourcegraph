import { describe, beforeEach, test, expect } from 'vitest';

import * as dom from './highlightNode'

describe('util/dom', () => {
    describe('highlightNode', () => {
        const cellInnerHTML = `<div><span class="hl-source hl-tsx"><span class="hl-meta hl-function hl-tsx"><span class="hl-keyword hl-control hl-export hl-tsx">export</span> <span class="hl-storage hl-type hl-function hl-tsx">function</span> <span class="hl-meta hl-definition hl-function hl-tsx"><span class="hl-entity hl-name hl-function hl-tsx">fetchSavedSearch</span></span><span class="hl-meta hl-parameters hl-tsx"><span class="hl-punctuation hl-definition hl-parameters hl-begin hl-tsx">(</span><span class="hl-variable hl-parameter hl-tsx">id</span><span class="hl-meta hl-type hl-annotation hl-tsx"><span class="hl-keyword hl-operator hl-type hl-annotation hl-tsx">:</span> <span class="hl-entity hl-name hl-type hl-tsx">Scalars</span><span class="hl-meta hl-type hl-tuple hl-tsx"><span class="hl-meta hl-brace hl-square hl-tsx">[</span><span class="hl-string hl-quoted hl-single hl-tsx"><span class="hl-punctuation hl-definition hl-string hl-begin hl-tsx">'</span>ID<span class="hl-punctuation hl-definition hl-string hl-end hl-tsx">'</span></span><span class="hl-meta hl-brace hl-square hl-tsx">]</span></span></span><span class="hl-punctuation hl-definition hl-parameters hl-end hl-tsx">)</span></span><span class="hl-meta hl-return hl-type hl-tsx"><span class="hl-keyword hl-operator hl-type hl-annotation hl-tsx">:</span> <span class="hl-entity hl-name hl-type hl-tsx">Observable</span><span class="hl-meta hl-type hl-parameters hl-tsx"><span class="hl-punctuation hl-definition hl-typeparameters hl-begin hl-tsx">&lt;</span></span><span class="hl-meta hl-type hl-parameters hl-tsx"><span class="hl-entity hl-name hl-type hl-module hl-tsx">GQL</span><span class="hl-punctuation hl-accessor hl-tsx">.</span><span class="hl-entity hl-name hl-type hl-tsx">ISavedSearch</span></span><span class="hl-meta hl-type hl-parameters hl-tsx"><span class="hl-punctuation hl-definition hl-typeparameters hl-end hl-tsx">&gt;</span></span> </span><span class="hl-meta hl-block hl-tsx"><span class="hl-punctuation hl-definition hl-block hl-tsx">{</span>
</span></span></span></div>`
        const cellWithEmojiInnerHTML = `<div><span class="hl-text hl-html hl-basic"><span class="hl-meta hl-tag hl-custom hl-html"><span class="hl-punctuation hl-definition hl-tag hl-begin hl-html">&lt;</span></span><span class="hl-meta hl-tag hl-custom hl-html"><span class="hl-entity hl-name hl-tag hl-custom hl-html">g-emoji</span> <span class="hl-meta hl-attribute-with-value hl-html"><span class="hl-entity hl-other hl-attribute-name hl-html">alias</span><span class="hl-punctuation hl-separator hl-key-value hl-html">=</span><span class="hl-string hl-quoted hl-double hl-html"><span class="hl-punctuation hl-definition hl-string hl-begin hl-html">"</span>+1<span class="hl-punctuation hl-definition hl-string hl-end hl-html">"</span></span></span> <span class="hl-meta hl-attribute-with-value hl-html"><span class="hl-entity hl-other hl-attribute-name hl-html">fallback-src</span><span class="hl-punctuation hl-separator hl-key-value hl-html">=</span><span class="hl-string hl-quoted hl-double hl-html"><span class="hl-punctuation hl-definition hl-string hl-begin hl-html">"</span>http://ghe.sgdev.org/images/icons/emoji/unicode/1f44d.png<span class="hl-punctuation hl-definition hl-string hl-end hl-html">"</span></span></span> <span class="hl-meta hl-attribute-with-value hl-class hl-html"><span class="hl-entity hl-other hl-attribute-name hl-class hl-html">class</span><span class="hl-punctuation hl-separator hl-key-value hl-html">=</span><span class="hl-string hl-quoted hl-double hl-html"><span class="hl-punctuation hl-definition hl-string hl-begin hl-html">"</span></span><span class="hl-string hl-quoted hl-double hl-html"><span class="hl-meta hl-class-name hl-html">emoji</span><span class="hl-punctuation hl-definition hl-string hl-end hl-html">"</span></span></span><span class="hl-punctuation hl-definition hl-tag hl-end hl-html">&gt;</span></span>👍<span class="hl-meta hl-tag hl-custom hl-html"><span class="hl-punctuation hl-definition hl-tag hl-begin hl-html">&lt;/</span></span><span class="hl-meta hl-tag hl-custom hl-html"><span class="hl-entity hl-name hl-tag hl-custom hl-html">g-emoji</span><span class="hl-punctuation hl-definition hl-tag hl-end hl-html">&gt;</span></span>
</span></div>`
        let cell: HTMLTableCellElement
        let cellWithEmoji: HTMLTableCellElement

        beforeEach(() => {
            document.body.innerHTML = `<table><tbody><td id="cell">${cellInnerHTML}</td><td id="cellWithEmoji">${cellWithEmojiInnerHTML}</td></tbody></table>`
            cell = window.document.querySelector<HTMLTableCellElement>('#cell')!
            cellWithEmoji = window.document.querySelector<HTMLTableCellElement>('#cellWithEmoji')!
        })

        test('highlights no characters', () => {
            dom.highlightNode(cell, 0, 0)
            expect(cell.innerHTML).toBe(cellInnerHTML) // no changes

            dom.highlightNode(cellWithEmoji, 0, 0)
            expect(cellWithEmoji.innerHTML).toBe(cellWithEmojiInnerHTML) // no changes
        })

        test('handles invalid start position', () => {
            dom.highlightNode(cell, -1, 3)
            expect(cell.innerHTML).toBe(cellInnerHTML) // no changes
            dom.highlightNode(cell, cell.textContent!.length, 3)
            expect(cell.innerHTML).toBe(cellInnerHTML) // no changes

            dom.highlightNode(cellWithEmoji, -1, 3)
            expect(cellWithEmoji.innerHTML).toBe(cellWithEmojiInnerHTML) // no changes
            dom.highlightNode(cellWithEmoji, cellWithEmoji.textContent!.length, 3)
            expect(cellWithEmoji.innerHTML).toBe(cellWithEmojiInnerHTML) // no changes
        })

        test('handles invalid length', () => {
            dom.highlightNode(cell, 0, 2000) // length longer than cell.innerText
            expect(cell.innerHTML).toBe(cellInnerHTML) // no changes
            dom.highlightNode(cell, 22, 80) // length longer than characters between start and end
            expect(cell.innerHTML).toBe(cellInnerHTML) // no changes

            dom.highlightNode(cellWithEmoji, 0, 2000) // length longer than cellWithEmoji.innerText
            expect(cellWithEmoji.innerHTML).toBe(cellWithEmojiInnerHTML) // no changes
            dom.highlightNode(cellWithEmoji, 22, 100) // length longer than characters between start and end
            expect(cellWithEmoji.innerHTML).toBe(cellWithEmojiInnerHTML) // no changes
        })

        test('highlights a single node', () => {
            dom.highlightNode(cell, 0, 1)
            expect(cell.innerHTML).toMatchInlineSnapshot(`
                "<div><span class=\\"hl-source hl-tsx\\"><span class=\\"hl-meta hl-function hl-tsx\\"><span class=\\"hl-keyword hl-control hl-export hl-tsx\\"><span class=\\"a11y-ignore\\"><span class=\\"match-highlight a11y-ignore\\">e</span>xport</span></span> <span class=\\"hl-storage hl-type hl-function hl-tsx\\">function</span> <span class=\\"hl-meta hl-definition hl-function hl-tsx\\"><span class=\\"hl-entity hl-name hl-function hl-tsx\\">fetchSavedSearch</span></span><span class=\\"hl-meta hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-parameters hl-begin hl-tsx\\">(</span><span class=\\"hl-variable hl-parameter hl-tsx\\">id</span><span class=\\"hl-meta hl-type hl-annotation hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\">:</span> <span class=\\"hl-entity hl-name hl-type hl-tsx\\">Scalars</span><span class=\\"hl-meta hl-type hl-tuple hl-tsx\\"><span class=\\"hl-meta hl-brace hl-square hl-tsx\\">[</span><span class=\\"hl-string hl-quoted hl-single hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-tsx\\">'</span>ID<span class=\\"hl-punctuation hl-definition hl-string hl-end hl-tsx\\">'</span></span><span class=\\"hl-meta hl-brace hl-square hl-tsx\\">]</span></span></span><span class=\\"hl-punctuation hl-definition hl-parameters hl-end hl-tsx\\">)</span></span><span class=\\"hl-meta hl-return hl-type hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\">:</span> <span class=\\"hl-entity hl-name hl-type hl-tsx\\">Observable</span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-begin hl-tsx\\">&lt;</span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-entity hl-name hl-type hl-module hl-tsx\\">GQL</span><span class=\\"hl-punctuation hl-accessor hl-tsx\\">.</span><span class=\\"hl-entity hl-name hl-type hl-tsx\\">ISavedSearch</span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-end hl-tsx\\">&gt;</span></span> </span><span class=\\"hl-meta hl-block hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-block hl-tsx\\">{</span>
                </span></span></span></div>"
            `)
        })

        test('highlights multiple nodes', () => {
            dom.highlightNode(cell, 2, 2)
            dom.highlightNode(cell, 23, 2)
            expect(cell.innerHTML).toMatchInlineSnapshot(`
                "<div><span class=\\"hl-source hl-tsx\\"><span class=\\"hl-meta hl-function hl-tsx\\"><span class=\\"hl-keyword hl-control hl-export hl-tsx\\"><span class=\\"a11y-ignore\\">ex<span class=\\"match-highlight a11y-ignore\\">po</span>rt</span></span> <span class=\\"hl-storage hl-type hl-function hl-tsx\\">function</span> <span class=\\"hl-meta hl-definition hl-function hl-tsx\\"><span class=\\"hl-entity hl-name hl-function hl-tsx\\"><span class=\\"a11y-ignore\\">fetchSa<span class=\\"match-highlight a11y-ignore\\">ve</span>dSearch</span></span></span><span class=\\"hl-meta hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-parameters hl-begin hl-tsx\\">(</span><span class=\\"hl-variable hl-parameter hl-tsx\\">id</span><span class=\\"hl-meta hl-type hl-annotation hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\">:</span> <span class=\\"hl-entity hl-name hl-type hl-tsx\\">Scalars</span><span class=\\"hl-meta hl-type hl-tuple hl-tsx\\"><span class=\\"hl-meta hl-brace hl-square hl-tsx\\">[</span><span class=\\"hl-string hl-quoted hl-single hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-tsx\\">'</span>ID<span class=\\"hl-punctuation hl-definition hl-string hl-end hl-tsx\\">'</span></span><span class=\\"hl-meta hl-brace hl-square hl-tsx\\">]</span></span></span><span class=\\"hl-punctuation hl-definition hl-parameters hl-end hl-tsx\\">)</span></span><span class=\\"hl-meta hl-return hl-type hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\">:</span> <span class=\\"hl-entity hl-name hl-type hl-tsx\\">Observable</span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-begin hl-tsx\\">&lt;</span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-entity hl-name hl-type hl-module hl-tsx\\">GQL</span><span class=\\"hl-punctuation hl-accessor hl-tsx\\">.</span><span class=\\"hl-entity hl-name hl-type hl-tsx\\">ISavedSearch</span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-end hl-tsx\\">&gt;</span></span> </span><span class=\\"hl-meta hl-block hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-block hl-tsx\\">{</span>
                </span></span></span></div>"
            `)
        })

        test('does not repeatedly highlight multiple nodes', () => {
            dom.highlightNode(cell, 0, 11)
            dom.highlightNode(cell, 0, 23)
            expect(cell.innerHTML).toMatchInlineSnapshot(`
                "<div><span class=\\"hl-source hl-tsx\\"><span class=\\"hl-meta hl-function hl-tsx\\"><span class=\\"hl-keyword hl-control hl-export hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">export</span></span><span class=\\"match-highlight a11y-ignore\\"> </span><span class=\\"hl-storage hl-type hl-function hl-tsx\\"><span class=\\"a11y-ignore\\"><span class=\\"match-highlight a11y-ignore\\">func</span><span class=\\"match-highlight a11y-ignore\\">tion</span></span></span><span class=\\"match-highlight a11y-ignore\\"> </span><span class=\\"hl-meta hl-definition hl-function hl-tsx\\"><span class=\\"hl-entity hl-name hl-function hl-tsx\\"><span class=\\"a11y-ignore\\"><span class=\\"match-highlight a11y-ignore\\">fetchSa</span>vedSearch</span></span></span><span class=\\"hl-meta hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-parameters hl-begin hl-tsx\\">(</span><span class=\\"hl-variable hl-parameter hl-tsx\\">id</span><span class=\\"hl-meta hl-type hl-annotation hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\">:</span> <span class=\\"hl-entity hl-name hl-type hl-tsx\\">Scalars</span><span class=\\"hl-meta hl-type hl-tuple hl-tsx\\"><span class=\\"hl-meta hl-brace hl-square hl-tsx\\">[</span><span class=\\"hl-string hl-quoted hl-single hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-tsx\\">'</span>ID<span class=\\"hl-punctuation hl-definition hl-string hl-end hl-tsx\\">'</span></span><span class=\\"hl-meta hl-brace hl-square hl-tsx\\">]</span></span></span><span class=\\"hl-punctuation hl-definition hl-parameters hl-end hl-tsx\\">)</span></span><span class=\\"hl-meta hl-return hl-type hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\">:</span> <span class=\\"hl-entity hl-name hl-type hl-tsx\\">Observable</span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-begin hl-tsx\\">&lt;</span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-entity hl-name hl-type hl-module hl-tsx\\">GQL</span><span class=\\"hl-punctuation hl-accessor hl-tsx\\">.</span><span class=\\"hl-entity hl-name hl-type hl-tsx\\">ISavedSearch</span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-end hl-tsx\\">&gt;</span></span> </span><span class=\\"hl-meta hl-block hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-block hl-tsx\\">{</span>
                </span></span></span></div>"
            `)
        })

        test('highlights after offset', () => {
            dom.highlightNode(cell, 2, 3)
            expect(cell.innerHTML).toMatchInlineSnapshot(`
                "<div><span class=\\"hl-source hl-tsx\\"><span class=\\"hl-meta hl-function hl-tsx\\"><span class=\\"hl-keyword hl-control hl-export hl-tsx\\"><span class=\\"a11y-ignore\\">ex<span class=\\"match-highlight a11y-ignore\\">por</span>t</span></span> <span class=\\"hl-storage hl-type hl-function hl-tsx\\">function</span> <span class=\\"hl-meta hl-definition hl-function hl-tsx\\"><span class=\\"hl-entity hl-name hl-function hl-tsx\\">fetchSavedSearch</span></span><span class=\\"hl-meta hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-parameters hl-begin hl-tsx\\">(</span><span class=\\"hl-variable hl-parameter hl-tsx\\">id</span><span class=\\"hl-meta hl-type hl-annotation hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\">:</span> <span class=\\"hl-entity hl-name hl-type hl-tsx\\">Scalars</span><span class=\\"hl-meta hl-type hl-tuple hl-tsx\\"><span class=\\"hl-meta hl-brace hl-square hl-tsx\\">[</span><span class=\\"hl-string hl-quoted hl-single hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-tsx\\">'</span>ID<span class=\\"hl-punctuation hl-definition hl-string hl-end hl-tsx\\">'</span></span><span class=\\"hl-meta hl-brace hl-square hl-tsx\\">]</span></span></span><span class=\\"hl-punctuation hl-definition hl-parameters hl-end hl-tsx\\">)</span></span><span class=\\"hl-meta hl-return hl-type hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\">:</span> <span class=\\"hl-entity hl-name hl-type hl-tsx\\">Observable</span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-begin hl-tsx\\">&lt;</span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-entity hl-name hl-type hl-module hl-tsx\\">GQL</span><span class=\\"hl-punctuation hl-accessor hl-tsx\\">.</span><span class=\\"hl-entity hl-name hl-type hl-tsx\\">ISavedSearch</span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-end hl-tsx\\">&gt;</span></span> </span><span class=\\"hl-meta hl-block hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-block hl-tsx\\">{</span>
                </span></span></span></div>"
            `)
        })

        test('highlights entire cell', () => {
            dom.highlightNode(cell, 0, cell.textContent!.length)
            expect(cell.innerHTML).toMatchInlineSnapshot(`
                "<div><span class=\\"hl-source hl-tsx\\"><span class=\\"hl-meta hl-function hl-tsx\\"><span class=\\"hl-keyword hl-control hl-export hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">export</span></span><span class=\\"match-highlight a11y-ignore\\"> </span><span class=\\"hl-storage hl-type hl-function hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">function</span></span><span class=\\"match-highlight a11y-ignore\\"> </span><span class=\\"hl-meta hl-definition hl-function hl-tsx\\"><span class=\\"hl-entity hl-name hl-function hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">fetchSavedSearch</span></span></span><span class=\\"hl-meta hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-parameters hl-begin hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">(</span></span><span class=\\"hl-variable hl-parameter hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">id</span></span><span class=\\"hl-meta hl-type hl-annotation hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">:</span></span><span class=\\"match-highlight a11y-ignore\\"> </span><span class=\\"hl-entity hl-name hl-type hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">Scalars</span></span><span class=\\"hl-meta hl-type hl-tuple hl-tsx\\"><span class=\\"hl-meta hl-brace hl-square hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">[</span></span><span class=\\"hl-string hl-quoted hl-single hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">'</span></span><span class=\\"match-highlight a11y-ignore\\">ID</span><span class=\\"hl-punctuation hl-definition hl-string hl-end hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">'</span></span></span><span class=\\"hl-meta hl-brace hl-square hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">]</span></span></span></span><span class=\\"hl-punctuation hl-definition hl-parameters hl-end hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">)</span></span></span><span class=\\"hl-meta hl-return hl-type hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">:</span></span><span class=\\"match-highlight a11y-ignore\\"> </span><span class=\\"hl-entity hl-name hl-type hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">Observable</span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-begin hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">&lt;</span></span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-entity hl-name hl-type hl-module hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">GQL</span></span><span class=\\"hl-punctuation hl-accessor hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">.</span></span><span class=\\"hl-entity hl-name hl-type hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">ISavedSearch</span></span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-end hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">&gt;</span></span></span><span class=\\"match-highlight a11y-ignore\\"> </span></span><span class=\\"hl-meta hl-block hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-block hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">{</span></span><span class=\\"match-highlight a11y-ignore\\">
                </span></span></span></span></div>"
            `)
        })

        // https://github.com/sourcegraph/sourcegraph/issues/20510
        test('highlights deeply nested spans correctly', () => {
            dom.highlightNode(cell, 26, 57)
            expect(cell.innerHTML).toMatchInlineSnapshot(`
                "<div><span class=\\"hl-source hl-tsx\\"><span class=\\"hl-meta hl-function hl-tsx\\"><span class=\\"hl-keyword hl-control hl-export hl-tsx\\">export</span> <span class=\\"hl-storage hl-type hl-function hl-tsx\\">function</span> <span class=\\"hl-meta hl-definition hl-function hl-tsx\\"><span class=\\"hl-entity hl-name hl-function hl-tsx\\"><span class=\\"a11y-ignore\\">fetchSaved<span class=\\"match-highlight a11y-ignore\\">Search</span></span></span></span><span class=\\"hl-meta hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-parameters hl-begin hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">(</span></span><span class=\\"hl-variable hl-parameter hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">id</span></span><span class=\\"hl-meta hl-type hl-annotation hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">:</span></span><span class=\\"match-highlight a11y-ignore\\"> </span><span class=\\"hl-entity hl-name hl-type hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">Scalars</span></span><span class=\\"hl-meta hl-type hl-tuple hl-tsx\\"><span class=\\"hl-meta hl-brace hl-square hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">[</span></span><span class=\\"hl-string hl-quoted hl-single hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">'</span></span><span class=\\"match-highlight a11y-ignore\\">ID</span><span class=\\"hl-punctuation hl-definition hl-string hl-end hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">'</span></span></span><span class=\\"hl-meta hl-brace hl-square hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">]</span></span></span></span><span class=\\"hl-punctuation hl-definition hl-parameters hl-end hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">)</span></span></span><span class=\\"hl-meta hl-return hl-type hl-tsx\\"><span class=\\"hl-keyword hl-operator hl-type hl-annotation hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">:</span></span><span class=\\"match-highlight a11y-ignore\\"> </span><span class=\\"hl-entity hl-name hl-type hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">Observable</span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-begin hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">&lt;</span></span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-entity hl-name hl-type hl-module hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">GQL</span></span><span class=\\"hl-punctuation hl-accessor hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">.</span></span><span class=\\"hl-entity hl-name hl-type hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">ISavedSearch</span></span></span><span class=\\"hl-meta hl-type hl-parameters hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-typeparameters hl-end hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">&gt;</span></span></span><span class=\\"match-highlight a11y-ignore\\"> </span></span><span class=\\"hl-meta hl-block hl-tsx\\"><span class=\\"hl-punctuation hl-definition hl-block hl-tsx\\"><span class=\\"match-highlight a11y-ignore\\">{</span></span>
                </span></span></span></div>"
            `)
        })

        test('highlights range containing unicode character correctly', () => {
            dom.highlightNode(cellWithEmoji, 106, 3)
            expect(cellWithEmoji.innerHTML).toMatchInlineSnapshot(`
                "<div><span class=\\"hl-text hl-html hl-basic\\"><span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-punctuation hl-definition hl-tag hl-begin hl-html\\">&lt;</span></span><span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-entity hl-name hl-tag hl-custom hl-html\\">g-emoji</span> <span class=\\"hl-meta hl-attribute-with-value hl-html\\"><span class=\\"hl-entity hl-other hl-attribute-name hl-html\\">alias</span><span class=\\"hl-punctuation hl-separator hl-key-value hl-html\\">=</span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-html\\">\\"</span>+1<span class=\\"hl-punctuation hl-definition hl-string hl-end hl-html\\">\\"</span></span></span> <span class=\\"hl-meta hl-attribute-with-value hl-html\\"><span class=\\"hl-entity hl-other hl-attribute-name hl-html\\">fallback-src</span><span class=\\"hl-punctuation hl-separator hl-key-value hl-html\\">=</span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-html\\">\\"</span>http://ghe.sgdev.org/images/icons/emoji/unicode/1f44d.png<span class=\\"hl-punctuation hl-definition hl-string hl-end hl-html\\">\\"</span></span></span> <span class=\\"hl-meta hl-attribute-with-value hl-class hl-html\\"><span class=\\"hl-entity hl-other hl-attribute-name hl-class hl-html\\">class</span><span class=\\"hl-punctuation hl-separator hl-key-value hl-html\\">=</span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-html\\">\\"</span></span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-meta hl-class-name hl-html\\">emoji</span><span class=\\"hl-punctuation hl-definition hl-string hl-end hl-html\\">\\"</span></span></span><span class=\\"hl-punctuation hl-definition hl-tag hl-end hl-html\\"><span class=\\"match-highlight a11y-ignore\\">&gt;</span></span></span><span class=\\"match-highlight a11y-ignore\\">👍</span><span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-punctuation hl-definition hl-tag hl-begin hl-html\\"><span class=\\"a11y-ignore\\"><span class=\\"match-highlight a11y-ignore\\">&lt;</span>/</span></span></span><span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-entity hl-name hl-tag hl-custom hl-html\\">g-emoji</span><span class=\\"hl-punctuation hl-definition hl-tag hl-end hl-html\\">&gt;</span></span>
                </span></div>"
            `)
        })

        test('highlights ranges before unicode character correctly', () => {
            dom.highlightNode(cellWithEmoji, 104, 3)
            expect(cellWithEmoji.innerHTML).toMatchInlineSnapshot(`
                "<div><span class=\\"hl-text hl-html hl-basic\\"><span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-punctuation hl-definition hl-tag hl-begin hl-html\\">&lt;</span></span><span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-entity hl-name hl-tag hl-custom hl-html\\">g-emoji</span> <span class=\\"hl-meta hl-attribute-with-value hl-html\\"><span class=\\"hl-entity hl-other hl-attribute-name hl-html\\">alias</span><span class=\\"hl-punctuation hl-separator hl-key-value hl-html\\">=</span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-html\\">\\"</span>+1<span class=\\"hl-punctuation hl-definition hl-string hl-end hl-html\\">\\"</span></span></span> <span class=\\"hl-meta hl-attribute-with-value hl-html\\"><span class=\\"hl-entity hl-other hl-attribute-name hl-html\\">fallback-src</span><span class=\\"hl-punctuation hl-separator hl-key-value hl-html\\">=</span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-html\\">\\"</span>http://ghe.sgdev.org/images/icons/emoji/unicode/1f44d.png<span class=\\"hl-punctuation hl-definition hl-string hl-end hl-html\\">\\"</span></span></span> <span class=\\"hl-meta hl-attribute-with-value hl-class hl-html\\"><span class=\\"hl-entity hl-other hl-attribute-name hl-class hl-html\\">class</span><span class=\\"hl-punctuation hl-separator hl-key-value hl-html\\">=</span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-html\\">\\"</span></span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-meta hl-class-name hl-html\\"><span class=\\"a11y-ignore\\">emoj<span class=\\"match-highlight a11y-ignore\\">i</span></span></span><span class=\\"hl-punctuation hl-definition hl-string hl-end hl-html\\"><span class=\\"match-highlight a11y-ignore\\">\\"</span></span></span></span><span class=\\"hl-punctuation hl-definition hl-tag hl-end hl-html\\"><span class=\\"match-highlight a11y-ignore\\">&gt;</span></span></span>👍<span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-punctuation hl-definition hl-tag hl-begin hl-html\\">&lt;/</span></span><span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-entity hl-name hl-tag hl-custom hl-html\\">g-emoji</span><span class=\\"hl-punctuation hl-definition hl-tag hl-end hl-html\\">&gt;</span></span>
                </span></div>"
            `)
        })

        test('highlights range after unicode character correctly', () => {
            dom.highlightNode(cellWithEmoji, 109, 9)
            expect(cellWithEmoji.innerHTML).toMatchInlineSnapshot(`
                "<div><span class=\\"hl-text hl-html hl-basic\\"><span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-punctuation hl-definition hl-tag hl-begin hl-html\\">&lt;</span></span><span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-entity hl-name hl-tag hl-custom hl-html\\">g-emoji</span> <span class=\\"hl-meta hl-attribute-with-value hl-html\\"><span class=\\"hl-entity hl-other hl-attribute-name hl-html\\">alias</span><span class=\\"hl-punctuation hl-separator hl-key-value hl-html\\">=</span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-html\\">\\"</span>+1<span class=\\"hl-punctuation hl-definition hl-string hl-end hl-html\\">\\"</span></span></span> <span class=\\"hl-meta hl-attribute-with-value hl-html\\"><span class=\\"hl-entity hl-other hl-attribute-name hl-html\\">fallback-src</span><span class=\\"hl-punctuation hl-separator hl-key-value hl-html\\">=</span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-html\\">\\"</span>http://ghe.sgdev.org/images/icons/emoji/unicode/1f44d.png<span class=\\"hl-punctuation hl-definition hl-string hl-end hl-html\\">\\"</span></span></span> <span class=\\"hl-meta hl-attribute-with-value hl-class hl-html\\"><span class=\\"hl-entity hl-other hl-attribute-name hl-class hl-html\\">class</span><span class=\\"hl-punctuation hl-separator hl-key-value hl-html\\">=</span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-punctuation hl-definition hl-string hl-begin hl-html\\">\\"</span></span><span class=\\"hl-string hl-quoted hl-double hl-html\\"><span class=\\"hl-meta hl-class-name hl-html\\">emoji</span><span class=\\"hl-punctuation hl-definition hl-string hl-end hl-html\\">\\"</span></span></span><span class=\\"hl-punctuation hl-definition hl-tag hl-end hl-html\\">&gt;</span></span>👍<span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-punctuation hl-definition hl-tag hl-begin hl-html\\"><span class=\\"a11y-ignore\\">&lt;<span class=\\"match-highlight a11y-ignore\\">/</span></span></span></span><span class=\\"hl-meta hl-tag hl-custom hl-html\\"><span class=\\"hl-entity hl-name hl-tag hl-custom hl-html\\"><span class=\\"match-highlight a11y-ignore\\">g-emoji</span></span><span class=\\"hl-punctuation hl-definition hl-tag hl-end hl-html\\"><span class=\\"match-highlight a11y-ignore\\">&gt;</span></span></span>
                </span></div>"
            `)
        })
    })
})
