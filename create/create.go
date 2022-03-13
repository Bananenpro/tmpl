package create

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/Bananenpro/tmpl/templates"
	"github.com/Bananenpro/tmpl/util"
)

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

func Create(name string) {
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("Failed to open the current directory:", err)
		os.Exit(1)
	}
	defer dir.Close()

	_, err = dir.Readdirnames(1)
	if err != io.EOF {
		fmt.Println("The current directory is not empty.")
		os.Exit(1)
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
	default:
		fmt.Println("Unknown template:", name)
		os.Exit(1)
	}

	Git()
	Readme()
	License()
}

func Readme() {
	if !util.YesNo("Create a README file?", true) {
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to create README:", err)
	}
	name := path.Base(dir)

	fileContent := fmt.Sprintf("# %s", name)

	err = os.WriteFile("README.md", []byte(fileContent), 0755)
	if err != nil {
		fmt.Println("Failed to create README:", err)
	}
}

func Git() {
	if util.IsInstalled("git") && util.YesNo("Do you want to initialize git?", true) {
		out, err := util.Execute("git", "init")
		if err != nil {
			fmt.Println("Failed to call 'git init':", out)
		}
	}
}

func License() {
	license := util.Select("Select a license", []string{"None", "MIT", "GPLv3", "Apache 2.0"}, 0)
	var fileContent string
	var readmeLicenseText string
	switch license {
	case "MIT":
		fileContent = fmt.Sprintf(templates.LicenseMIT, time.Now().Year(), util.GetName())
		readmeLicenseText = fmt.Sprintf(templates.LicenseMIT, time.Now().Year(), util.GetName())
	case "GPLv3":
		fileContent = templates.LicenseGPLv3
		readmeLicenseText = fmt.Sprintf(templates.LicenseGPLv3ReadmeText, time.Now().Year(), util.GetName())
	case "Apache 2.0":
		fileContent = templates.LicenseApache2
		readmeLicenseText = fmt.Sprintf(templates.LicenseApache2ReadmeText, time.Now().Year(), util.GetName())
	default:
		return
	}

	if fileContent != "" {
		err := os.WriteFile("LICENSE", []byte(fileContent), 0755)
		if err != nil {
			fmt.Println("Failed to write license text into 'LICENSE':", err)
		}
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