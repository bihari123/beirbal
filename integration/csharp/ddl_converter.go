package csharp

import (
	"fmt"
	"os/exec"
)

func DdlConverter(filePath string) {
	// go build -ldflags="-s -w" -buildmode=c-shared -o libgo.dll main.go

	out, err := exec.Command("go", "build", "-ldflags=\"-s -w\"", "-buildmode=c-shared", "-o", "libgo.dll", filePath).
		Output()
	if err != nil {
		fmt.Println("error faced while converting into DLL format: ", err)
		return
	}

	output := string(out[:])
	fmt.Println(output)
}
