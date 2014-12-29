package main

import (
    "fmt"
    "os"
    "strings"
    "path/filepath"
    "io/ioutil"
    "bufio"
    "log"
    "regexp"
    )

func main() {
    if l := len(os.Args); l < 2 {
        fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(os.Args[0]))
        os.Exit(1)
    }

    var fileName string
    fmt.Sscanf(os.Args[1], _M3U_FILE_NAME_FORMAT, &fileName)
    if fileName == "" {
        fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(os.Args[0]))
        os.Exit(1)
    }

    if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
        log.Fatal(err)
    } else {
        songs := loadSongsFromM3uFile(string(rawBytes))
        writeSongsToPlsFile(songs, fileName)
    }
}

func loadSongsFromM3uFile(text string) []song {
    lines := strings.Split(text, "\n")
    songs := []song {}
    rx := regexp.MustCompile(_M3U_INF_LINE_PATTERN)
    for i, l := 1, len(lines); i < l && i + 1 < l; i += 2 {
        matches := rx.FindStringSubmatch(lines[i])
        songs = append(songs, song { matches[2], matches[1], lines[i + 1] })
    }
    return songs
}

func writeSongsToPlsFile(songs []song, fileName string) (err error) {
    var file *os.File
    filePath := fmt.Sprintf(_PLS_FILE_NAME_FORMAT, fileName)
    if file, err = os.Create(filePath); err != nil {
        log.Fatalf("Cannot create output file %s", filePath)
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    defer func() {
        if err == nil {
            err = writer.Flush()
        }
    }()

    writer.WriteString(_PLS_TOP_LINE)
    l := len(songs)
    for i := 0; i < l; i++ {
        fmt.Fprintf(writer, _PLS_FILE_LINE_FORMAT, i, songs[i].file)
        fmt.Fprintf(writer, _PLS_TITLE_LINE_FORMAT, i, songs[i].title)
        fmt.Fprintf(writer, _PLS_LENGTH_LINE_FORMAT, i, songs[i].length)
    }
    fmt.Fprintf(writer, _PLS_ENTRIES_COUNT_LINE_FORMAT, l)
    fmt.Fprintf(writer, _PLS_VERSION_LINE_FORMAT, 2)
    return err
}
