package main

import (
	"fmt"
	"github.com/lightbend/cloudflow/test-docker-apis/skopeoinspect"
	_ "github.com/lightbend/cloudflow/test-docker-apis/skopeoinspect"
	"os"
	"time"
)

func main() {
	var ins skopeoinspect.InspectOptions

	ins.Global = &skopeoinspect.GlobalOptions{}

	ins.Image = &skopeoinspect.ImageOptions{}

	ins.Image.DockerImageOptions = skopeoinspect.DockerImageOptions{Global: &skopeoinspect.GlobalOptions{}, Shared: &skopeoinspect.SharedImageOptions{}}

	elapsed := time.Since(time.Now())
	err := ins.Run([]string{"docker://docker.io/lightbend/spark-aggregation:134-d0ec286-dirty"}, os.Stdout)
	fmt.Printf("dockerhub inspect took %s\n", elapsed)
	if err != nil {
		println(err.Error())
	}
	//docker.PrintMetadata()
	elapsed = time.Since(time.Now())
	err = ins.Run([]string{"docker-daemon:lightbend/spark-aggregation:134-d0ec286-dirty"}, os.Stdout)
	fmt.Printf("local daemon took %s\n", elapsed)
	if err != nil {
		println(err.Error())
	}

	elapsed = time.Since(time.Now())
	err = ins.Run([]string{"docker://405074236871.dkr.ecr.eu-west-1.amazonaws.com/stavros-test/sensor-data-scala:90-d662d87-dirty"}, os.Stdout)
	fmt.Printf("AWS inspect took %s\n", elapsed)
	if err != nil {
		println(err.Error())
	}

	elapsed = time.Since(time.Now())
	err = ins.Run([]string{"docker://eu.gcr.io/bubbly-observer-178213/spark-aggregation:134-d0ec286-dirty"}, os.Stdout)
	fmt.Printf("GCloud inspect took %s\n", elapsed)
	if err != nil {
		println(err.Error())
	}
}
