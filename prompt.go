package main

import (
	"github.com/Songmu/prompter"
	"github.com/gen2brain/dlgs"
)

func promptPassphrase(title string, message string) (value string, ok bool, err error) {
	// first try with GUI prompt
	value, ok, err = dlgs.Password(title, message)
	if err == nil {
		return
	}

	// if GUI is not available, try CUI one
	value = prompter.Password(message)
	ok = value != ""
	err = nil
	return
}
