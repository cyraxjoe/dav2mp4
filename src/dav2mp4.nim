import strformat, parseopt, tables
import davconverter

const
  progName = "dav2mp4"
  version = "0.1"
  supportedFormats = {
    "mp4": MP4,
    "raw": RAW,
    "avi": AVI,
    "asf": ASF
  }.toTable

type
  Options = object
    inputFile: string
    outputFile: string
    outputFormat: string


# proc `$`(options: Options): string =
#   return fmt"Options: {{inputFile: {options.inputFile}, outputFile: {options.outputFile}, outputFormat: {options.outputFormat}}}"


proc writeVersion =
  echo version
  quit(QuitSuccess)


proc writeHelp(exit = false) =
  var help = &"{prog_name} [options]  <input-dav-file> <output-file>\n"
  help &= "Options:\n"
  help &= "\t -v, --version: Show the version number and exit.\n"
  help &= "\t -h, --help: Show this help message and exit.\n"
  help &= "\t -f, --format: Video output format, supported: mp4 (default), raw, avi and asf.\n"
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
          result.inputFile = key
        elif arg_position == 1:
          result.outputFile = key
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
              showErrorAndExit(fmt"Unsupported format '{val}'")
      of cmdEnd:
        break
  if result.outputFormat == "":
    result.outputFormat = "mp4"
  if result.inputFile == "" or result.outputFile == "":
    showErrorAndExit("Missing required arguments")
        
when isMainModule:
  let o = getOptions()
  convert(o.inputFile, o.outputFile, supportedFormats[o.outputFormat])



