package main

import (
	"dav2mp4/pkg/converter"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	version      = "0.2"
	outputFormat string
	showVersion  bool
	showHelp     bool
	batchMode    bool
)

func main() {
	flag.StringVar(&outputFormat, "format", "mp4", "Video output format")
	flag.StringVar(&outputFormat, "f", "mp4", "Video output format (short)")

	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.BoolVar(&showVersion, "v", false, "Show version (short)")

	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.BoolVar(&showHelp, "h", false, "Show help (short)")

	flag.BoolVar(&batchMode, "batch-mode", false, "Batch mode")
	flag.BoolVar(&batchMode, "b", false, "Batch mode (short)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "dav2mp4 [options] <input-dav-file-or-dir> <output-file-or-dir>\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "  -f, --format <format>     Video output format (mp4, raw, avi, asf) [default: mp4]\n")
		fmt.Fprintf(os.Stderr, "  -b, --batch-mode          Use positional arguments as input and output directories\n")
		fmt.Fprintf(os.Stderr, "  -v, --version             Show version number and exit\n")
		fmt.Fprintf(os.Stderr, "  -h, --help                Show this help message and exit\n")
	}

	flag.Parse()

	if showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Error: Missing required arguments")
		flag.Usage()
		os.Exit(1)
	}

	inputPath := args[0]
	outputPath := args[1]

	var fmtEnum converter.VideoFormat
	switch outputFormat {
	case "mp4":
		fmtEnum = converter.MP4
	case "raw":
		fmtEnum = converter.RAW
	case "avi":
		fmtEnum = converter.AVI
	case "asf":
		fmtEnum = converter.ASF
	default:
		fmt.Printf("Error: Unsupported format '%s' (use ':' to indicate the format '--format:FORMAT')\n", outputFormat)
		os.Exit(1)
	}

	if batchMode {
		processBatch(inputPath, outputPath, outputFormat, fmtEnum)
	} else {
		processSingle(inputPath, outputPath, fmtEnum)
	}
}

func processSingle(input, output string, format converter.VideoFormat) {
	// check if input is directory? The original code didn't check explicitly in non-batch mode but passed it to convert which checks file existence.
	// But let's follow the original flow.
	err := converter.Convert(input, output, format)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func processBatch(inputDir, outputDir, formatStr string, format converter.VideoFormat) {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		fmt.Printf("Creating output directory: %s\n", outputDir)
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			os.Exit(1)
		}
	}

	entries, err := os.ReadDir(inputDir)
	if err != nil {
		fmt.Printf("Error reading input directory: %v\n", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".dav") {
			inFile := filepath.Join(inputDir, entry.Name())
			// Original code: outputDir / changeFileExt(extractFilename(inputFile.path), videoFormat)
			// changeFileExt replaces extension.
			outName := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())) + "." + formatStr
			outFile := filepath.Join(outputDir, outName)

			fmt.Printf("Processing file: %s -> %s\n", inFile, outFile)
			if err := converter.Convert(inFile, outFile, format); err != nil {
				fmt.Printf("Error converting %s: %v\n", inFile, err)
			}
		}
	}
	fmt.Println("All done.")
}
