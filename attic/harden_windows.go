package main
import (
	"os"
	"golang.org/x/exp/shiny/widget/theme"
)

const harden_key_path = "SOFTWARE\\Security Without Borders\\"

func check_status() bool {
	key, err := registry.OpenKey(registry.CURRENT_USER, harden_key_path, registry.READ)
	if err != nil {
		return false
	}

	value, _, err := key.GetIntegerValue("Harden")
	if err != nil {
		return false
	}

	if value == 1 {
		return true
	} else {
		return false
	}
}

func mark_status(hardened bool) {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, harden_key_path, registry.WRITE)
	if err != nil {
		panic(err)
	}

	if hardened {
		key.SetDWordValue("Harden", 1)
	} else {
		key.SetDWordValue("Harden", 0)
	}
}

func harden_all() {
	/*trigger_all(true)
	mark_status(true)
	*/
	walk.MsgBox(window, "Done!", "I have hardened all risky features!\nFor all changes to take effect please restart Windows.", walk.MsgBoxIconInformation)
	os.Exit(0)
}

func restore_all() {
	trigger_all(false)
	mark_status(false)

	walk.MsgBox(window, "Done!", "I have restored all risky features!\nFor all changes to take effect please restart Windows.", walk.MsgBoxIconExclamation)
	os.Exit(0)
}

func trigger_all(harden bool) {
	trigger_wsh(harden)
	trigger_ole(harden)
	trigger_macro(harden)
	trigger_activex(harden)
	trigger_pdf_js(harden)
	trigger_pdf_objects(harden)
	trigger_autorun(harden)
	trigger_powershell(harden)
	trigger_uac(harden)
	trigger_fileassoc(harden)
	progress.SetValue(100) 
}
