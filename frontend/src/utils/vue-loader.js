import Vue from 'vue'

export function loadComponent(cmp, mixin, parent) {
  if (cmp.default) {
    cmp = cmp.default
  }
  cmp.mixins = [mixin]
  cmp = Vue.extend(cmp)
  const node = document.createElement('div')
  if (!parent.appendChild) {
    parent = document.querySelector(parent)
  }
  parent.appendChild(node)
  new cmp({ // eslint-disable-line
    el: node,
    parent: this
  })
}
