package create

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/Bananenpro/tmpl/external"
	"github.com/Bananenpro/tmpl/input"
	"github.com/Bananenpro/tmpl/templates"
)

func Create(name string) {
	dir, err := os.Open(".")
	if err != nil {
		abort(fmt.Sprint("Failed to open the current directory: ", err))
	}
	defer dir.Close()

	_, err = dir.Readdirnames(1)
	if err != io.EOF {
		abort("The current directory is not empty.")
	}

	switch name {
	case "go", "golang":
		CreateGo()
	case "python":
		CreatePython()
	case "c#", "csharp", "cs":
		CreateCSharp()
	case "rust":
		CreateRust()
	case "c":
		CreateC()
	default:
		abort(fmt.Sprint("Unknown template: ", name))
	}

	Git()
	Readme()
	License()
}

func Git() {
	if external.IsInstalled("git") && input.YesNo("Do you want to initialize git?", true) {
		out, err := external.Execute("git", "init")
		if err != nil {
			fmt.Println("Failed to call 'git init':", out)
		}
	}
}

func Readme() {
	if !input.YesNo("Create a README file?", true) {
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to create README:", err)
	}
	name := path.Base(dir)

	fileContent := fmt.Sprintf("# %s", name)

	CreateFile("README.md", fileContent)
}

func License() {
	license := input.Select("Select a license", []string{"None", "MIT", "GPLv3", "Apache 2.0"}, 0)
	var fileContent string
	var readmeLicenseText string
	switch license {
	case "MIT":
		fileContent = fmt.Sprintf(templates.LicenseMIT, time.Now().Year(), external.GetUsername())
		readmeLicenseText = fmt.Sprintf(templates.LicenseMIT, time.Now().Year(), external.GetUsername())
	case "GPLv3":
		fileContent = templates.LicenseGPLv3
		readmeLicenseText = fmt.Sprintf(templates.LicenseGPLv3ReadmeText, time.Now().Year(), external.GetUsername())
	case "Apache 2.0":
		fileContent = templates.LicenseApache2
		readmeLicenseText = fmt.Sprintf(templates.LicenseApache2ReadmeText, time.Now().Year(), external.GetUsername())
	default:
		return
	}

	if fileContent != "" {
		CreateFile("LICENSE", fileContent)
	}

	if _, err := os.Stat("README.md"); err == nil {
		readme, err := os.OpenFile("README.md", os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			fmt.Println("Failed to append license text to README:", err)
		}
		defer readme.Close()

		text := "\n\n## License\n\n%s"
		readme.WriteString(fmt.Sprintf(text, readmeLicenseText))
	}
}

func CreateFile(path, content string) {
	err := os.WriteFile(path, []byte(content), 0755)
	if err != nil {
		abort(fmt.Sprintf("Error while creating file '%s': %s", path, err))
	}
}

func Clean() error {
	files, err := os.ReadDir(".")
	if err != nil {
		return err
	}
	for _, f := range files {
		err = os.RemoveAll(f.Name())
		if err != nil {
			fmt.Println("Failed to remove file:", f.Name())
		}

	}
	return nil
}

func abort(msg string) {
	Clean()
	fmt.Println(msg)
	os.Exit(1)
}
