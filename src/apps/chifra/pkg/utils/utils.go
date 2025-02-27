// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"golang.org/x/crypto/ssh/terminal"
)

// IsTestModeServer return true if we are running from the testing harness
func IsTestModeServer(r *http.Request) bool {
	return r.Header.Get("User-Agent") == "testRunner"
}

func IsApiMode() bool {
	return os.Getenv("API_MODE") == "true"
}

func IsTerminal() bool {
	return terminal.IsTerminal(int(os.Stdout.Fd()))
}

func AsciiFileToString(fileName string) string {
	if !file.FileExists(fileName) {
		return ""
	}

	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(contents)
}

func AsciiFileToLines(fileName string) []string {
	contents := AsciiFileToString(fileName)
	return strings.Split(contents, "\n")
}

func OpenBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func PadLeft(str string, totalLen int) string {
	if len(str) >= totalLen {
		return str
	}
	zeros := ""
	for i := 0; i < totalLen-len(str); i++ {
		zeros += "0"
	}
	return zeros + str
}

func ToCamelCase(in string) string {
	if len(in) == 0 {
		return in
	}

	var arr []string
	fields := strings.Fields(in)
	for _, field := range fields {
		arr = append(arr, strings.Title(field))
	}
	arr[0] = strings.ToLower(arr[0])
	return strings.Join(arr, "")
}

// maximum uint64
const NOPOS = ^uint64(0)

// Min calculates the minimum between two unsigned integers (golang has no such function)
func Min(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}

// Max calculates the max between two unsigned integers (golang has no such function)
func Max(x, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}

func MakeFirstLowerCase(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}
	bts := []byte(s)
	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]
	return string(bytes.Join([][]byte{lc, rest}, nil))
}

func MakeFirstUpperCase(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}
	bts := []byte(s)
	lc := bytes.ToUpper([]byte{bts[0]})
	rest := bts[1:]
	return string(bytes.Join([][]byte{lc, rest}, nil))
}

// TODO: Multi-chain this is the earliest timestamp on mainnet. It's larger than any block on any EVM chain by an order of maginitude
// TODO: but is not a legit way to distinguish between blocks and timestamps. Fix this when it becomes a problem. All other EVM based
// TODO: chains have later first timestamps, so this works for other chains as well
const EarliestTs = 1438269975
const BlockOneTs = 1438269988
