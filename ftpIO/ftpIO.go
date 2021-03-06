package ftpIO

import (
    "fmt"
    "net"
    "net/http"
    "io"
)

func Close(conn net.Conn) {
    fmt.Printf("Closing connection\n")
    conn.Close()
}

func WriteRaw(conn net.Conn, rawText string) {
    conn.Write([]byte(rawText))
}

func Write(conn net.Conn, status int, text string) {
    // Must format string according to FTP spec, see :
    // https://github.com/dagwieers/vsftpd/blob/master/ftpcmdio.c
    response := fmt.Sprintf("%d %s\r\n", status, text)
    WriteRaw(conn, response)
}

/*
func CheckUrl(httpIp string, filePath string) (bool) {
    url := fmt.Sprintf("http://%s%s", httpIp, filePath)
    fmt.Printf("Checking file: %s\n", url)

    // Check if file is accessible
    resp, err := http.Head(url)
    if err != nil || resp.StatusCode != 200 {
        fmt.Printf("Error trying to HEAD url: %s\n", url)
        return false
    }
    return true
}
*/

func OpenUrl(httpIp string, filePath string, resp **http.Response) (bool) {
    if (resp == nil) {
        return false
    }

    url := fmt.Sprintf("http://%s%s", httpIp, filePath)
    fmt.Printf("Opening url: %s\n", url)

    var err error
    *resp, err = http.Get(url)
    if err != nil || resp == nil || (*resp).StatusCode != 200 {
        fmt.Printf("Error trying to GET url: %s\n", url)
        CloseUrl(*resp)
        return false
    }
    /* do not forget to CloseUrl() from here */
    return true
}

func SendUrl(conn net.Conn, resp *http.Response) (bool) {
    if (resp == nil) {
        return false
    }

    _, err := io.Copy(conn, resp.Body)
    if err != nil {
        fmt.Println("Error copying file (for RETR):", err.Error())
        return false
    }
    return true
}

func CloseUrl(resp *http.Response) (bool) {
    if resp == nil || resp.Body == nil {
        return false
    }

    resp.Body.Close()
    return true
}
