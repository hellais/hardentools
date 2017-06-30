import regedit from 'regedit'

const adobeVersions = [
	"DC", // Acrobat Reader DC
	"XI", // Acrobat Reader XI - To test
]

export default class Adobe {

	constructor() {
		this.onEvent = this.onEvent.bind(this)
		this.onError = this.onError.bind(this)
	}

	onEvent(text) {
		// XXX move this into the base class
		console.log(text)
	}

	onError(context, err) {
		// XXX move this into the base class
		console.log(context, err)
	}

	triggerPdfObjects(harden, cb) {
    /*
    bAllowOpenFile set to 0 and
    bSecureOpenFile set to 1 to disable
    the opening of non-PDF documents
    */

    /*func trigger_pdf_objects(harden bool) {
	var allow_value uint32
	var secure_value uint32

	if harden==false {
		events.AppendText("Restoring default by enabling embedded objects in PDFs\n")
		allow_value = 1
		secure_value = 0
	} else {
		events.AppendText("Hardening by disabling embedded objects in PDFs\n")
		allow_value = 0
		secure_value = 1
	}

	for _, adobe_version := range adobe_versions {
		path := fmt.Sprintf("SOFTWARE\\Adobe\\Acrobat Reader\\%s\\Originals", adobe_version)
		key, _, _ := registry.CreateKey(registry.CURRENT_USER, path, registry.WRITE)
		
		key.SetDWordValue("bAllowOpenFile", allow_value)
		key.SetDWordValue("bSecureOpenFile", secure_value)
		key.Close()
	}
  }
  */

  }

	triggerPdfJs(harden, cb) {
		/*
		bEnableJS possible values:
		0 - Disable AcroJS
		1 - Enable AcroJS
		*/
		let keyValue

		if (harden === false) {
			onEvent("Restoring default by enabling Acrobat Reader JavaScript")
			keyValue = 1
		} else {
			onEvent("Hardening by disabling Acrobat Reader JavaScript")
			keyValue = 0
		}
		
		adobeVersions.forEach((adobeVersion) => {
			let path = `SOFTWARE\Adobe\Acrobat Reader\${adobeVersion}\JSPrefs`
      regedit.createKey(path, (err) => {
        if (err) {
          onError({'module': 'adobe', 'what': 'createKey'}, err)
        } else {
          let value = {}
          value[path] = {
            'bEnabledJS': {
              value: keyValue,
              type: 'REG_DWORD'
            }
          }
          regedit.putValue(value, (err) => {
            if (err) {
              onError({'module': 'adobe', 'what': 'putValue'}, err)
            } else {
              cb("Hardened keys")
            }
          })
        }
      })
		}
	}

	run(cb) {
    return triggerPdfJs(true, cb)
	}

	revert() {
    return triggerPdfJs(false, cb)
	}
}

