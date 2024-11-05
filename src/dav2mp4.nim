import strformat, parseopt, tables
import std/[dirs, paths]
import davconverter

const
  progName = "dav2mp4"
  version = "0.2"
  supportedFormats = {
    "mp4": MP4,
    "raw": RAW,
    "avi": AVI,
    "asf": ASF
  }.toTable

type
  Options = object
    inputFileOrDir: Path
    outputFileOrDir: Path
    outputFormat: string
    batchMode = false


# proc `$`(options: Options): string =
#   return fmt"Options: {{inputFile: {options.inputFile}, outputFile: {options.outputFile}, outputFormat: {options.outputFormat}}}"


proc writeVersion =
  echo version
  quit(QuitSuccess)


proc writeHelp(exit = false) =
  var help = &"{prog_name} [options]  <input-dav-file-or-dir> <output-file-or-dir>\n"
  help &= "Options:\n"
  help &= "\t -v, --version: Show the version number and exit.\n"
  help &= "\t -h, --help: Show this help message and exit.\n"
  help &= "\t -f, --format: Video output format, supported: mp4 (default), raw, avi and asf.\n"
  help &= "\t -b, --batch-mode: Use the positional arguments as input and output directories."
  echo help
  if exit:
    quit(QuitSuccess)


proc showErrorAndExit(msg: string) =
  echo "Error: " & msg
  writeHelp()
  quit(QuitFailure)
  

proc getOptions(): Options = 
  var arg_position = 0
  for kind, key, val in getopt():
    case kind:
      of cmdArgument:
        if arg_position == 0:
          result.inputFileOrDir = key.Path
        elif arg_position == 1:
          result.outputFileOrDir = key.Path
        else:
          showErrorAndExit("Invalid number of arguments")
        arg_position += 1
      of cmdShortOption, cmdLongOption:
        case key:
          of "help", "h": writeHelp(exit=true)
          of "version", "v": writeVersion()
          of "format", "f":
            if val in supportedFormats:
              result.outputFormat = val
            else:
              showErrorAndExit(fmt"Unsupported format '{val}' (use ':' to indicate the format '--format:FORMAT')")
          of "batch-mode", "b":
            result.batchMode = true
      of cmdEnd:
        break
  if result.outputFormat == "":
    result.outputFormat = "mp4"
  if result.inputFileOrDir.string == "" or result.outputFileOrDir.string == "":
    showErrorAndExit("Missing required arguments")


proc batchConversion(inputDir, outputDir: Path, videoFormat: string) =
  if not dirExists(outputDir):
    echo fmt"Creating output directory: {outputDir}"
    createDir(outputDir)
  for inputFile in walkDir(inputDir):
    if inputFile.path.splitFile().ext == ".dav":
      let outputFile = outputDir / changeFileExt(extractFilename(inputFile.path), videoFormat)
      echo fmt"Processing file: {inputFile.path} -> {outputFile}"
      convert(inputFile.path.string, outputFile.string, supportedFormats[videoFormat])
  echo "All done."

        
when isMainModule:
  let o = getOptions()
  if o.batchMode:
    batchConversion(o.inputFileOrDir, o.outputFileOrDir, o.outputFormat)
  else:
    convert(o.inputFileorDir.string, o.outputFileorDir.string, supportedFormats[o.outputFormat])



