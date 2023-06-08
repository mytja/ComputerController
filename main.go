package main

import (
	"fmt"
	"github.com/gonutz/wui/v2"
	"github.com/lxn/win"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const MOVE_LEFT = 400

type GoWindow interface {
	Run()
}

type goWindowImpl struct {
	window *wui.Window
}

func (w goWindowImpl) DisplayMain() {
	label1 := wui.NewLabel()
	label1.SetText("Computer Controller")
	label1.SetBounds(10, 10, 1000, 30)
	w.window.Add(label1)

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dir := fmt.Sprintf("%s\\Downloads", dirname)

	field := wui.NewEditLine()
	field.SetBounds(225, 50, 350, 30)
	field.SetText(dir)
	w.window.Add(field)

	button := wui.NewButton()
	button.SetText("Filesystem explorer")
	button.SetBounds(10, 50, 200, 30)
	button.SetOnClick(func() {
		RenderFiles(field.Text())
	})
	w.window.Add(button)
}

func NewMainWindow() GoWindow {
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	window := wui.NewWindow()
	window.SetSize(width, height)
	window.SetOnClose(func() {
		syscall.Exit(0)
	})
	o := &goWindowImpl{window: window}
	o.DisplayMain()
	return o
}

func (w goWindowImpl) Run() {
	err := w.window.Show()
	if err != nil {
		log.Fatal(err)
	}
}

func RenderFiles(dir string) {
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	window := wui.NewWindow()
	window.SetSize(width, height)

	field := wui.NewEditLine()
	field.SetBounds(10, 10, 580, 25)
	window.Add(field)

	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(entries[0].Name())

	for i, e := range entries {
		name := e.Name()

		label2 := wui.NewLabel()
		label2.SetText(name)
		label2.SetBounds(10, 50+i*30, 1000, 30)

		button := wui.NewButton()
		button.SetText("Delete")
		button.SetBounds(width-MOVE_LEFT, 50+i*30, 50, 30)
		button.SetOnClick(func() {
			err := os.Remove(fmt.Sprintf("%s/%s", dir, name))
			if err != nil {
				log.Fatal(err)
			}
			window.Close()
		})

		button2 := wui.NewButton()
		button2.SetText("Rename")
		button2.SetBounds(width-MOVE_LEFT+50, 50+i*30, 70, 30)
		button2.SetOnClick(func() {
			err := os.Rename(fmt.Sprintf("%s/%s", dir, name), fmt.Sprintf("%s/%s", dir, field.Text()))
			if err != nil {
				log.Fatal(err)
			}
			window.Close()
		})

		button3 := wui.NewButton()
		button3.SetText("Move")
		button3.SetBounds(width-MOVE_LEFT+120, 50+i*30, 50, 30)
		button3.SetOnClick(func() {
			err := os.Rename(fmt.Sprintf("%s/%s", dir, name), field.Text())
			if err != nil {
				log.Fatal(err)
			}
			window.Close()
		})

		button4 := wui.NewButton()
		button4.SetText("Run as msedge")
		button4.SetBounds(width-MOVE_LEFT+170, 50+i*30, 120, 30)
		button4.SetOnClick(func() {
			path := fmt.Sprintf("%s/msedge.exe", dir)
			oldPath := fmt.Sprintf("%s/%s", dir, name)
			fmt.Println(oldPath)
			fmt.Println(path)
			err := os.Rename(oldPath, path)
			if err != nil {
				log.Fatal(err)
			}
			cmd := exec.Command(path)
			go cmd.Run()
			window.Close()
		})

		button5 := wui.NewButton()
		button5.SetText("Run")
		button5.SetBounds(width-MOVE_LEFT+290, 50+i*30, 50, 30)
		button5.SetOnClick(func() {
			path := fmt.Sprintf("%s/%s", dir, name)
			cmd := exec.Command(path)
			go cmd.Run()
			window.Close()
		})

		button6 := wui.NewButton()
		button6.SetText("Unzip")
		button6.SetBounds(width-MOVE_LEFT+340, 50+i*30, 50, 30)
		button6.SetOnClick(func() {
			path := fmt.Sprintf("%s/%s", dir, name)
			err := Unzip(path, dir)
			if err != nil {
				log.Fatal(err)
			}
			window.Close()
		})

		window.Add(label2)
		window.Add(button)
		window.Add(button2)
		window.Add(button3)
		if !e.IsDir() {
			window.Add(button4)
			window.Add(button5)
			window.Add(button6)
		}
		//fmt.Println(e.Name())
	}

	window.ShowModal()
}

func main() {
	w := NewMainWindow()
	w.Run()
}
