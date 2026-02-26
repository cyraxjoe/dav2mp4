package main

import (
	"dav2mp4/pkg/converter"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("DAV to MP4 Converter")
	myWindow.Resize(fyne.NewSize(600, 400))

	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Select input file or directory...")

	outputEntry := widget.NewEntry()
	outputEntry.SetPlaceHolder("Select output directory...")

	// Format selection
	formats := []string{"mp4", "raw", "avi", "asf"}
	formatSelect := widget.NewSelect(formats, nil)
	formatSelect.SetSelected("mp4")

	progressBar := widget.NewProgressBarInfinite()
	progressBar.Hide()

	statusLabel := widget.NewLabel("Ready")

	// Buttons for file dialogs
	inputBtn := widget.NewButton("Browse File", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if reader == nil {
				return
			}
			inputEntry.SetText(reader.URI().Path())
			// Set default output dir if empty
			if outputEntry.Text == "" {
				outputEntry.SetText(filepath.Dir(reader.URI().Path()))
			}
		}, myWindow)
	})

	inputDirBtn := widget.NewButton("Browse Dir", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if list == nil {
				return
			}
			inputEntry.SetText(list.Path())
			// Set default output dir if empty
			if outputEntry.Text == "" {
				outputEntry.SetText(list.Path())
			}
		}, myWindow)
	})

	outputBtn := widget.NewButton("Browse Dir", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if list == nil {
				return
			}
			outputEntry.SetText(list.Path())
		}, myWindow)
	})

	convertBtn := widget.NewButton("Convert", func() {
		inputPath := inputEntry.Text
		outputPath := outputEntry.Text
		formatStr := formatSelect.Selected

		if inputPath == "" || outputPath == "" {
			statusLabel.SetText("Error: Please select input and output paths.")
			return
		}

		var format converter.VideoFormat
		switch formatStr {
		case "mp4":
			format = converter.MP4
		case "raw":
			format = converter.RAW
		case "avi":
			format = converter.AVI
		case "asf":
			format = converter.ASF
		default:
			format = converter.MP4
		}

		progressBar.Show()
		statusLabel.SetText("Converting...")

		go func() {
			defer progressBar.Hide()

			fileInfo, err := os.Stat(inputPath)
			if err != nil {
				statusLabel.SetText(fmt.Sprintf("Error accessing input: %v", err))
				return
			}

			if fileInfo.IsDir() {
				// Batch mode
				files, err := os.ReadDir(inputPath)
				if err != nil {
					statusLabel.SetText(fmt.Sprintf("Error reading directory: %v", err))
					return
				}

				count := 0
				for _, file := range files {
					if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".dav") {
						statusLabel.SetText(fmt.Sprintf("Converting %s...", file.Name()))

						inFile := filepath.Join(inputPath, file.Name())
						outName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + "." + formatStr
						outFile := filepath.Join(outputPath, outName)

						err := converter.Convert(inFile, outFile, format)
						if err != nil {
							fmt.Printf("Error converting %s: %v\n", file.Name(), err)
							statusLabel.SetText(fmt.Sprintf("Error converting %s: %v", file.Name(), err))
							time.Sleep(1 * time.Second)
						} else {
							count++
						}
					}
				}
				statusLabel.SetText(fmt.Sprintf("Done! Converted %d files.", count))
			} else {
				// Single file mode
				filename := filepath.Base(inputPath)
				outName := strings.TrimSuffix(filename, filepath.Ext(filename)) + "." + formatStr
				outFile := filepath.Join(outputPath, outName)

				statusLabel.SetText(fmt.Sprintf("Converting %s...", filename))
				err := converter.Convert(inputPath, outFile, format)
				if err != nil {
					statusLabel.SetText(fmt.Sprintf("Error: %v", err))
				} else {
					statusLabel.SetText("Conversion Complete!")
				}
			}
		}()
	})

	// Layout
	content := container.NewVBox(
		widget.NewLabel("Input (File or Directory):"),
		container.NewBorder(nil, nil, nil, container.NewHBox(inputBtn, inputDirBtn), inputEntry),
		widget.NewLabel("Output Directory:"),
		container.NewBorder(nil, nil, nil, outputBtn, outputEntry),
		widget.NewLabel("Output Format:"),
		formatSelect,
		layout.NewSpacer(),
		convertBtn,
		progressBar,
		statusLabel,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
