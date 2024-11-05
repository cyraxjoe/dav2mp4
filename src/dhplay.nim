
const
  libName = when defined(linux):
              "libdhplay.so"
            elif defined(windows):
              "play.dll"
            else:
              ""
  STREAME_REALTIME*: cint = 0
  STREAME_FILE*: cint = 1


if libName == "":
  raise newException(ValueError, "Unsupported platform")

type
  DATA_RECORD_TYPE* = enum
    DATA_RECORD_ORIGINAL,         # Raw video stream
    DATA_RECORD_AVI,              # AVI
    DATA_RECORD_ASF,              # ASF   
    DATA_RECORD_ORIGINAL_SEGMENT, # Segment Raw video stream
    DATA_RECORD_RESIZE_AVI,       # Resolution Changed AVI, target width and height is needed. use PLAY_ResolutionScale.
    DATA_RECORD_MP4,              # MP4 
    DATA_RECORD_RESIZE_MP4,       # Resolution Changed MP4, target width and height is needed. use PLAY_ResolutionScale.
    DATA_RECORD_MP4_NOSEEK,       # No Seek MP4    
    DATA_RECORD_RESIZE_MP4_NOSEEK,# Resolution Changed No Seek MP4, target width and height is needed. use PLAY_ResolutionScale.
    DATA_RECORD_TS,
    DATA_RECORD_PS,               # PS
    DATA_RECORD_RESIZE_DAV,
    DATA_RECORD_DAV,
    DATA_RECORD_AAC,              # AAC(Raw Audio)
    DATA_RECORD_WAV,              # WAV(Raw Audio)
    # if need to add enumeration, add beore it
    DATA_RECORD_COUNT             # record type count 
  


{.push dynlib: libName.}
# BOOL PLAY_InputData(LONG nPort,PBYTE pBuf,DWORD nSize);
proc playInputData*(nPort: cint, pBuf: ptr cuchar, nSize: cuint): cint {.importc: "PLAY_InputData".}

# unsigned int PLAY_GetBufferValue(LONG nPort,DWORD nBufType);
proc playGetBufferValue*(nPort: cint, nBufType: cuint): uint {.importc: "PLAY_GetBufferValue".}

#unsigned int PLAY_GetSourceBufferRemain(LONG nPort);
proc playGetSourceBufferRemains*(nPort: cint): uint {.importc: "PLAY_GetSourceBufferRemain".}

#BOOL PLAY_SetStreamOpenMode(LONG nPort,DWORD nMode);
proc playSetStreamOpenMode*(nPort: cint, nMode: cuint): cint {.importc: "PLAY_SetStreamOpenMode".}

#BOOL PLAY_OpenStream(LONG nPort,PBYTE pFileHeadBuf,DWORD nSize,DWORD nBufPoolSize);
proc playOpenStream*(nPort: cint, pFileHeadBuf: ptr cuchar, nSize, nBufPoolSize: cuint): cint {.importc: "PLAY_OpenStream".}

#BOOL PLAY_CloseStream(LONG nPort);
proc playCloseStream*(nPort: cint): cint {.importc: "PLAY_CloseStream".}

#BOOL PLAY_Play(LONG nPort, HWND hWnd);
proc playPlay*(nPort: cint, hWnd: pointer): cint {.importc: "PLAY_Play".}

#BOOL PLAY_Stop(LONG nPort);
proc playStop*(nPort: cint): cint {.importc: "PLAY_Stop".}


#PLAYSDK_API BOOL CALLMETHOD PLAY_StartDataRecord(LONG nPort, char *sFileName, int idataType);
proc playStartDataRecord*(nPort: cint, sfileName: cstring, idataType: cint): cint {.importc: "PLAY_StartDataRecord".}

#BOOL PLAY_StartDataRecordEx(LONG nPort, char *sFileName, int idataType, fRecordDataCBFun fListenter, void* pUserData);
proc playStartDataRecordEx*(nPort: cint, sFileName: cstring, idataType: cint, fRecordDataCBFun, fListener, pUserData: pointer): cint {.importc: "PLAY_StartDataRecordEx".}

#BOOL PLAY_StopDataRecord(LONG nPort);
proc playStopDataRecord*(nPort: cint): cint {.importc: "PLAY_StopDataRecord".}

