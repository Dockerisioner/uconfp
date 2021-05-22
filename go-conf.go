package main // import "github.com/dockerisioner/go-conf"

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

var Version = "dev"

func version() string {
	return fmt.Sprintf(`%s (%s on %s/%s; %s)`, Version, runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.Compiler)
}

func usage() string {
	t := template.Must(template.New("usage").Parse(`
Usage: {{ .Self }} go-conf command [args]
   eg: {{ .Self }} get test
{{ .Self }} version: {{ .Version }}
{{ .Self }} license: GPL-3 (full text at https://github.com/dockerisioner/go-conf)
`))
	var b bytes.Buffer
	template.Must(t, t.Execute(&b, struct {
		Self    string
		Version string
	}{
		Self:    filepath.Base(os.Args[0]),
		Version: version(),
	}))
	return strings.TrimSpace(b.String()) + "\n"
}

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "--help", "-h", "-?":
			fmt.Println(usage())
			os.Exit(0)
		case "--version", "-v":
			fmt.Println(version())
			os.Exit(0)
		}
	}
	if len(os.Args) <= 2 {
		fmt.Println(usage())
		os.Exit(1)
	}
}
