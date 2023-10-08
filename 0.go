package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

var line string = `
package main
import (
	"fmt"
)
func main() {
	fmt.Println("начало работы")
}
	`

func createFile(a, b int) {
	for i := a; i <= b; i++ {
		path := "go.dir" + strconv.Itoa(i)
		filename := filepath.Join(path, strconv.Itoa(i)+".go")
		os.Mkdir(path, os.ModeDir|0755)

		file, err := os.Create(filename)
		file.WriteString(line)
		if err != nil {
			panic(err)
		}
		file.Close()

	}
	fmt.Println("Сделано епт")
}

func main() {
	createFile(1, 5)
}
