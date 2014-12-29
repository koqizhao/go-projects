package main

const _M3U_TOP_LINE = "#EXTM3U"
const _M3U_INF_LINE_PATTERN = `#EXTINF:(-{0,1}\d+),(.*)`

const _PLS_TOP_LINE = "[playlist]\n"
const _PLS_FILE_LINE_FORMAT = "File%d=%s\n"
const _PLS_TITLE_LINE_FORMAT = "Title%d=%s\n"
const _PLS_LENGTH_LINE_FORMAT = "Length%d=%s\n"
const _PLS_ENTRIES_COUNT_LINE_FORMAT = "NumberOfEntries=%d\n"
const _PLS_VERSION_LINE_FORMAT = "Version=%d\n"

const _M3U_FILE_NAME_FORMAT = "%s.m3u"
const _PLS_FILE_NAME_FORMAT = "%s.pls"

type song struct {
    title string
    length string
    file string
}
