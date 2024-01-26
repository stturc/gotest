package cirruszip

import (
	"os"
	"path/filepath"
)

func Build() {
	newpath := filepath.Join(".", "ui")
	os.MkdirAll(newpath, os.ModePerm)

	newpath1 := filepath.Join("./ui", "build")
	os.MkdirAll(newpath1, os.ModePerm)

	newpath2 := filepath.Join("./ui/build", "content_deployer")
	os.MkdirAll(newpath2, os.ModePerm)

	newpath3 := filepath.Join("./ui/build/content_deployer", "tools")
	os.MkdirAll(newpath3, os.ModePerm)

	newpath4 := filepath.Join("./ui/build", "content")
	os.MkdirAll(newpath4, os.ModePerm)

	newpath5 := filepath.Join("./ui/build/content", "ui_bundles")
	os.MkdirAll(newpath5, os.ModePerm)

	newpath6 := filepath.Join("./ui/build/content", "ui")
	os.MkdirAll(newpath6, os.ModePerm)

	newpath7 := filepath.Join(".", "authorization")
	os.MkdirAll(newpath7, os.ModePerm)                

	newpath8 := filepath.Join(".", "i18n")
	os.MkdirAll(newpath8, os.ModePerm)

	newpath9 := filepath.Join("./i18n", "app_registry")
	os.MkdirAll(newpath9, os.ModePerm)

	newpath10 := filepath.Join(".", "data_loader_files")
	os.MkdirAll(newpath10, os.ModePerm)

	newpath11 := filepath.Join(".", "identities")
	os.MkdirAll(newpath11, os.ModePerm)

	newpath12 := filepath.Join(".", "notifications")
	os.MkdirAll(newpath12, os.ModePerm)

	newpath13 := filepath.Join(".", "reports")
	os.MkdirAll(newpath13, os.ModePerm)

	newpath14 := filepath.Join(".", "spre")
	os.MkdirAll(newpath14, os.ModePerm)

	newpath15 := filepath.Join(".", "workflow")
	os.MkdirAll(newpath15, os.ModePerm)

	newpath16 := filepath.Join(".", "stephen")
	os.MkdirAll(newpath16, os.ModePerm)
}
