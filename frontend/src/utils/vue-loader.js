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

export function initMeta(meta) {
  if (meta.init) {
    return
  }
  const dict = {}
  for (let i = 0; i < meta.model.attributes.length; i++) {
    const item = meta.model.attributes[i]
    dict[item.name] = item
  }
  for (let i = 0; i < meta.table.actions.length; i++) {
    const item = meta.table.actions[i]
    if (item.type === 'dialog') {
      for (let i = 0; i < item.dialog.fields.length; i++) {
        item.dialog.fields[i] = dict[item.dialog.fields[i].name]
      }
    }
  }
  meta.init = true
}
