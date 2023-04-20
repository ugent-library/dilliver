
import htmx from 'htmx.org'
import bsn from 'bootstrap.native/dist/bootstrap-native-v4'
import bsnPopper from './bsn_popper.js'
import toast from './toast.js'
import formSubmit from './form_submit.js'
import formUploadProgress from './form_upload_progress.js'
import clipboard from './clipboard.js'
import selectValue from './select_value.js'

window.htmx = htmx

// load htmx extensions
require('htmx.org/dist/ext/ws.js');

htmx.config.defaultFocusScroll = true

htmx.onLoad(function(el) {
    bsn.initCallback(el)
    bsnPopper(el)
    toast(el)
    formSubmit(el)
    formUploadProgress(el)
    clipboard(el)
    selectValue(el)
})

htmx.on('htmx:config-request', (evt) => {
    evt.detail.headers['X-CSRF-Token'] = document.querySelector('meta[name="csrf-token"]').content
});

htmx.on('htmx:confirm', (evt) => {
    let el = evt.detail.elt

    if (el.dataset.confirm) {
        evt.preventDefault()

        let modalEl = document.getElementById('modal-confirm').content.firstElementChild.cloneNode(true)
        document.body.appendChild(modalEl)

        if (el.dataset.confirmHeader) {
            modalEl.querySelector('.confirm-header').innerHTML = el.dataset.confirmHeader
        }
        if (el.dataset.confirmProceed) {
            modalEl.querySelector('.confirm-proceed').innerHTML = el.dataset.confirmProceed
        }

        modalEl.querySelector('.confirm-proceed').addEventListener('click', () => {
            evt.detail.issueRequest()
        }, false)

        modalEl.addEventListener('hidden.bs.modal', () => {
            modalEl.remove()
        }, false);

        new bsn.Modal(modalEl).show()
    }
});
