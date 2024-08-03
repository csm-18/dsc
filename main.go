package main

import (
	"fmt"
	"os"
	"os/user"
)

const VERSION = "v1.0.0"

func main() {
	var command string
	var args = os.Args

	if len(args) == 1 {
		command = "about"
	} else if len(args) == 2 {
		command = args[1]
	} else {
		fmt.Println("dsc: invalid command!")
		fmt.Println("On how to use dsc, run:\n dsc help")
		os.Exit(0)
	}

	userPrompt(command)
}

func userPrompt(command string) {
	if command == "about" {
		fmt.Println("dsc: Desktop Shortcut Creator")
		fmt.Println("version:", VERSION)
		fmt.Println("\nSupported app formats:")
		fmt.Println(" AppImages,")
		fmt.Println(" binary (gui) apps,")
		fmt.Println(" jar format (gui) apps")
		fmt.Println("\nFor more info, run:")
		fmt.Println(" dsc help")
	} else if command == "version" {
		fmt.Println("dsc", VERSION)
		fmt.Println("\nFor more info, run:")
		fmt.Println(" dsc help")
	} else if command == "help" {
		fmt.Println("dsc commands list:")
		fmt.Println(" dsc          --> prints about dsc")
		fmt.Println(" dsc version  --> prints dsc version")
		fmt.Println(" dsc help     --> prints this commands list")
		fmt.Println(" dsc new      --> to create a new shortcut")
		fmt.Println(" dsc delete   --> to delete a shortcut")
	} else if command == "new" {
		var shortcut Shortcut
		println("dsc", VERSION)
	Main:
		for {
			fmt.Println("name(shortcut name):")
			fmt.Scanln(&shortcut.appName)
			if shortcut.appName == "" {
				fmt.Println("Enter all fields!")
				continue Main
			}

			fmt.Println("Enter app type (1,2,3):")
			fmt.Println("  1. AppImage")
			fmt.Println("  2. Binary")
			fmt.Println("  3. jar (java jar)")
			fmt.Scanln(&shortcut.appType)
			switch shortcut.appType {
			case "1":
				shortcut.appType = "bin"
			case "2":
				shortcut.appType = "bin"
			case "3":
				shortcut.appType = "jar"
			default:
				fmt.Println("Enter valid option!")
				continue Main
			}

			fmt.Println("Enter full path to app (should not have spaces in it):")
			fmt.Scanln(&shortcut.appPath)
			if shortcut.appPath == "" {
				fmt.Println("Enter all fields!")
				continue Main
			}

			fmt.Println("Enter full path to icon (optional):")
			fmt.Scanln(&shortcut.appIconPath)
			if shortcut.appIconPath == "" {
				shortcut.appIconPath = "utilities-terminal"
			}

			shortcut.create()
			break Main
		}

	} else if command == "delete" {
		var shortcutName string
		var deleteShortcut string
		var desktopShortcutPath string
		var current_user string
		fmt.Println("dsc", VERSION)
		for {
			fmt.Println("Enter name of the shortcut:")
			fmt.Scanln(&shortcutName)
			if shortcutName == "" {
				fmt.Println("Enter all fields!")
				continue
			}

			fmt.Println("Do you really want to delete the shortcut?(y/n)")
			fmt.Scanln(&deleteShortcut)
			if deleteShortcut == "" {
				fmt.Println("Enter all fields!")
				continue
			}

			if deleteShortcut == "y" || deleteShortcut == "Y" {

				usr, error := user.Current()
				if error != nil {
					fmt.Println("Error while deleting shortcut!")
					os.Exit(0)
				} else {
					current_user = usr.Username
				}

				desktopShortcutPath = "/home/" + current_user + "/Desktop/" + shortcutName + ".desktop"
				_, err := os.Stat(desktopShortcutPath)
				if os.IsNotExist(err) {
					fmt.Println("\nNo shortcut exists with the specified name!")
					fmt.Println("Enter correct name!")
					os.Exit(0)
				}

				fmt.Println("\nFor safety reasons, only the 'desktop shortcut' is deleted,\n not the 'desktop entry' of the application.")
				err = os.Remove(desktopShortcutPath)
				if err != nil {
					fmt.Println("Error while deleting shortcut!")
					os.Exit(0)
				} else {
					fmt.Println("\nSuccessfully deleted shortcut:")
					fmt.Println(" ", desktopShortcutPath)
				}

			} else if deleteShortcut == "n" || deleteShortcut == "N" {
				os.Exit(0)
			} else {
				fmt.Println("Enter valid option!")
				continue
			}
			break
		}
	}
}

type Shortcut struct {
	appName             string
	appType             string
	appPath             string
	appIconPath         string
	desktopEntryPath    string
	desktopShortcutPath string
}

func (s Shortcut) create() {
	var content string
	var current_user string
	usr, error := user.Current()
	if error != nil {
		fmt.Println("Error while creating shortcut!")
		os.Exit(0)
	} else {
		current_user = usr.Username
	}

	if s.appType == "jar" {
		s.appPath = "java -jar " + s.appPath
	}
	content += "[Desktop Entry]\n"
	content += "Type=Application\n"
	content += fmt.Sprintf("Name=%s\n", s.appName)
	content += fmt.Sprintf("Comment=%s\n", s.appName)
	content += "Terminal=false\n"
	content += fmt.Sprintf("Exec=%s\n", s.appPath)
	content += fmt.Sprintf("Icon=%s\n", s.appIconPath)

	s.desktopEntryPath = "/home/" + current_user + "/.local/share/applications/" + s.appName + ".desktop"

	_, err := os.Stat(s.desktopEntryPath)
	if os.IsNotExist(err) {
		data := []byte(content)
		err := os.WriteFile(s.desktopEntryPath, data, 0755)
		if err != nil {
			fmt.Println("Error while creating shortcut!")
			os.Exit(0)
		}

		s.desktopShortcutPath = "/home/" + current_user + "/Desktop/" + s.appName + ".desktop"
		err = os.WriteFile(s.desktopShortcutPath, data, 0755)
		if err != nil {
			fmt.Println("Error while creating shortcut!")
			os.Exit(0)
		}

		fmt.Println("Successfully created shorcut!")
		fmt.Println("Now, go to desktop and test your shorcut!")

	} else {
		fmt.Println("A shorcut already exists with the")
		fmt.Println(" same name!")
		fmt.Println("Please try another name for shortcut!")
	}
}
