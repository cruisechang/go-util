package profile

import (
	"log"
	"os"
	"runtime/pprof"
)

//ProfileCPU sava CPU status into target file.
func ProfileCPU(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
}
