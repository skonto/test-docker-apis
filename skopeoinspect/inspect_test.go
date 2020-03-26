package skopeoinspect

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_inspect(t *testing.T) {

	var ins InspectOptions

	ins.global = &GlobalOptions{}

	ins.image = &ImageOptions{}

	ins.image.DockerImageOptions = DockerImageOptions{global: &GlobalOptions{}, shared: &SharedImageOptions{}}

	elapsed := time.Since(time.Now())
	err := ins.run([]string{"docker://docker.io/lightbend/spark-aggregation:134-d0ec286-dirty"}, os.Stdout)
	fmt.Printf("dockerhub inspect took %s\n", elapsed)
	if err != nil {
		println(err.Error())
	}
	elapsed = time.Since(time.Now())
	err = ins.run([]string{"docker://eu.gcr.io/bubbly-observer-178213/spark-aggregation:134-d0ec286-dirty"}, os.Stdout)
	fmt.Printf("GCloud inspect took %s\n", elapsed)
	if err != nil {
		println(err.Error())
	}
	elapsed = time.Since(time.Now())
	err = ins.run([]string{"docker://405074236871.dkr.ecr.eu-west-1.amazonaws.com/stavros-test/sensor-data-scala:90-d662d87-dirty"}, os.Stdout)
	fmt.Printf("AWS inspect took %s\n", elapsed)
	if err != nil {
		println(err.Error())
	}

	elapsed = time.Since(time.Now())
	err = ins.run([]string{"docker-daemon:lightbend/spark-aggregation:134-d0ec286-dirty"}, os.Stdout)
	fmt.Printf("local daemon took %s\n", elapsed)
	if err != nil {
		println(err.Error())
	}
}
