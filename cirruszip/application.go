package cirruszip

import (
	"os"
	"os/exec"
	"path/filepath"
)

func Build() {
    // Create The build/zips path to hold all the intermediate zips
	os.MkdirAll(filepath.Join(".", "build"), os.ModePerm)
	os.MkdirAll(filepath.Join("./build", "zips"), os.ModePerm)

	// Copy over everything needed from the deployer
	cmd := exec.Command("cp", "--recursive", "submodules/cirrus-deployer/deploy-solution.sh", "build/zips/.").CombinedOutput()
	cmd.Run()

	cmd = exec.Command("cp", "--recursive", "submodules/cirrus-deployer/deployer", "build/zips/.").CombinedOutput()
	cmd.Run()

	cmd = exec.Command("cp", "--recursive", "submodules/cirrus-deployer/tools", "build/zips/.").CombinedOutput()
	cmd.Run()

	cmd = exec.Command("cp", "--recursive", "run-manifest.json", "build/zips/.").CombinedOutput()
	cmd.Run()

    // Remove unneeded folders & files from the deployer
	os.RemoveAll("build/zips/tools/examples")
	os.RemoveAll("build/zips/tools/scratch")
	os.RemoveAll("build/zips/tools/html")
	os.RemoveAll("build/zips/tools/install_deps.py")
	os.RemoveAll("build/zips/tools/generate_pdoc.cmd")


    // Create The build/zips/libs path
	os.MkdirAll(filepath.Join("./build/zips", "libs"), os.ModePerm)

	// the CirrusMRM.zip file contains everything that needs to be imported to builder-api
	// namely CirrusMRMApp.zip which contains the ui json files and CirrusMRMLib.zip which
	// contains the library js and json files
	cmd = exec.Command("cp", "--recursive", "libs/dist/CirrusMRM.zip", "build/zips/libs/CirrusMRM.zip").CombinedOutput()
	cmd.Run()

    // Create The build/zips/libs path
	os.MkdirAll(filepath.Join("./build/zips", "cirrus-mrm-app"), os.ModePerm)

	cmd = exec.Command("zip", "-r", "build/temp.zip", "*", "-x 'submodules/*' -x 'saspackage/*' -x '*.git*' -x 'temp/*' -x 'ci/*' -x 'home/*' -x 'dist/*' -x 'libs/*' -x 'build/*' -x 'bd-scan*' -x 'go*' -x 'sage*'").CombinedOutput()
	cmd.Run()

	cmd = exec.Command("unzip", "build/temp.zip", "-d", "build/zips/cirrus-mrm-app").CombinedOutput()
	cmd.Run()


	cmd = exec.Command("zip", "-r", "CirrusMRM.zip", "*").CombinedOutput()
	cmd.Dir = "build/zips"
	cmd.Run()

    // Create The build/zips/libs path
	os.MkdirAll(filepath.Join(".", "ui"), os.ModePerm)
	os.MkdirAll(filepath.Join("./ui", "build"), os.ModePerm)
	
	cmd = exec.Command("ls", "--R", "build/zips").CombinedOutput()
	cmd.Run()

	cmd = exec.Command("cp", "--recursive", "build/zips/CirrusMRM.zip", "ui/build/app-ui.zip").CombinedOutput()
	cmd.Run()
}
