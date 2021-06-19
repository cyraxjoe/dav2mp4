import os, strformat
import dhplay

const
  PORT: cint = 0
  SOURCEBUF_SIZE: cuint = (500 * 1024)

type
  VideoFormat* = enum
    RAW = DATA_RECORD_ORIGINAL
    AVI = DATA_RECORD_AVI
    ASF = DATA_RECORD_ASF
    MP4 = DATA_RECORD_MP4
  

proc init_stream(output_path: string, output_format: VideoFormat): cint = 
  let port = PORT
  if not bool(playSetStreamOpenMode(port, STREAME_FILE.cuint)):
    raise newException(Exception, "Unable to set stream mode")
  if not bool(playOpenStream(port, nil, 0, SOURCEBUF_SIZE)):
    raise newException(Exception, "Unable to open stream")
  if not bool(playPlay(port, nil)):
    raise newException(Exception, "Unable to play on stream")
  if not bool(playStartDataRecordEx(port, output_path, output_format.cint, nil, nil, nil)):
    raise newException(Exception, "Unable to start data record")
  return port


proc stop_and_close_stream(port: cint) =
  discard playStopDataRecord(port)
  discard playStop(port)
  discard playCloseStream(port)  


proc convert*(input_file_path: string, output_file_path: string, output_format: VideoFormat) =
  if not fileExists(input_file_path):
    raise newException(Exception, fmt"The input file '{input_file_path}' does not exists.")
  if fileExists(output_file_path):
    raise newException(Exception, fmt"The output file '{output_file_path}' already exists.")
  let
    port = init_stream(output_file_path, output_format)
    inputFile = open(input_file_path)
  const readLen = 8 * 1024
  var
    readBuff: array[readLen, cuchar]
    bytesRead: int = 0
  bytesRead = readBuffer(inputFile, readBuff.addr, readLen)
  while bytesRead > 0:
    while not playInputData(port, readBuff[0].addr, bytesRead.cuint).bool:
      sleep(10)
    bytesRead = readBuffer(inputFile, readBuff.addr, readLen)
  stop_and_close_stream(port)

