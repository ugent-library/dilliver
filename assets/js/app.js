import BSN from "bootstrap.native/dist/bootstrap-native-v4";
import bootstrapPopper from './bootstrap_popper.js'
import htmx from 'htmx.org'
import bootstrap from './bootstrap.js'
import toast from './toast.js'
import csrf from './csrf.js'
import formSubmit from './form_submit.js'
import formUploadProgress from './form_upload_progress.js'
import modalClose from './modal_close.js'

// configure htmx
htmx.config.defaultFocusScroll = true
// load htmx extensions
window.htmx = htmx

// initialize everyting
document.addEventListener('DOMContentLoaded', function () {
    BSN.initCallback()
    bootstrapPopper()
    bootstrap()
    toast()
    csrf()
    formSubmit()
    formUploadProgress()
    modalClose()
});
