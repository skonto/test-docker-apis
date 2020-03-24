package docker

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

var (
	dockerRegistry = "https://registry-1.docker.io"
)

type Token struct {
	Token string `json:"token"`
	Expires_in int `json:"expires_in"`
	Issued_at time.Time `json:"issued_at"`
}

type ImageConfig struct {
	MediaType string `json:"mediaType"`
	Size int `json:"size"`
	Digest string `json:"digest"`
}

func PrintMetadata() {
	httpClient := &http.Client{
		Timeout: 30*time.Second,
	}

	// test dockerhub
	image := "lightbend/spark-aggregation"
	var token = GetToken(*httpClient, image)
	var digest = GetDigest(*httpClient, image, "134-d0ec286-dirty", token, dockerRegistry, true)
	raw := GetImageConfiguration(*httpClient, image, digest, token)


	compressed := base64.NewDecoder(base64.StdEncoding, bytes.NewReader([]byte(raw)))
	reader, err := zlib.NewReader(compressed)
	if err != nil {
		LogAndExit("Failed to decompress the Cloudflow streamlet descriptors label for image %s, %s", image, err.Error())
	}

	// uncompressed data : []byte
	uncompressed, err := ioutil.ReadAll(reader)
	if err != nil {
		LogAndExit("Failed to read the decompressed Cloudflow streamlet descriptors label for image %s, %s", image, err.Error())
	}

	var descriptors Descriptors
	err = json.Unmarshal([]byte(uncompressed), &descriptors)
	fmt.Printf("%v\n", descriptors)

	// test Google image registry

	gImage := "bubbly-observer-178213/spark-aggregation"

	var gkeAuthToken = GetGKERegistryToken()

	var gkeToken = fmt.Sprintf("{\"username\": \"oauth2accesstoken\", \"password\": \"%s\", \"email\": \"stavros@lightbend.gr\", \"serveraddress\": \"ip\"}", gkeAuthToken)

	// https://docs.docker.com/engine/api/v1.39/#section/Authentication
	gkeToken = base64.StdEncoding.EncodeToString([]byte(gkeToken))

	println(GetDigest(*httpClient, gImage, "134-d0ec286-dirty", gkeToken, "https://eu.gcr.io", false))

	// test EKS registry

	// test Openshift

}


func GetGKERegistryToken() string {
	cmd := exec.Command("gcloud", "auth", "print-access-token")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return string(out)
}

// GetToken gets a token to be used with dockerhub.io
func GetToken(client http.Client, image string) string {
	elapsedTotal := time.Since(time.Now())
	req, err := http.NewRequest("GET", fmt.Sprintf("https://auth.docker.io/token?scope=repository:%s:pull&service=registry.docker.io", image), nil)

	elapsed := time.Since(time.Now())
	resp, err := client.Do(req)
	fmt.Printf("getToken raw http request took %s\n", elapsed)
	if err != nil {
		LogAndExit("Could not get a client 2: %s", err.Error())
	}

	if resp == nil || resp.StatusCode != 200 {
		LogAndExit("Could not get a client 3: %d", resp.StatusCode)
	}
	elapsed = time.Since(time.Now())
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Read token http resp took %s\n", elapsed)
	if err != nil {
		LogAndExit("Could not get a client 4: %s", err.Error())
	}
	contents := string(body)
	fmt.Printf("Read token http resp took %s\n", elapsed)
	var tk Token

	elapsed = time.Since(time.Now())
	err = json.Unmarshal([]byte(contents), &tk)
	fmt.Printf("unmarshal token took %s\n", elapsed)

	if err!= nil {
		LogAndExit("Could not get a client 4: %s", err.Error())
	}

	fmt.Printf(" getToken function took %s\n", elapsedTotal)
	return tk.Token
}

// GetDigest gets an image digest to be used for get image configuration data
func GetDigest(client http.Client, image string, tag string, token string, registryUrl string, useBearer bool) string {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/%s/manifests/%s", registryUrl,image, tag), nil)

	if useBearer {
		var tk= "Bearer " + token
		req.Header.Add("Authorization", tk)
	} else {
		req.Header.Add("X-Registry-Auth", token)
	}
	req.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	elapsed := time.Since(time.Now())
	resp, err := client.Do(req)
	fmt.Printf("getDigest raw http request took %s\n", elapsed)
	if err != nil {
		LogAndExit("Could not get a client 2: %s", err.Error())
	}

	if resp == nil || resp.StatusCode != 200 {
		LogAndExit("Could not get a client 3: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogAndExit("Could not get a client 4: %s", err.Error())
	}
	contents := string(body)
	var reqMap map[string]interface{}

	err = json.Unmarshal([]byte(contents), &reqMap)

	configMap := reqMap["config"].(map[string]interface {})

	return configMap["digest"].(string)
}

// GetImageConfiguration gets image configuration based on the image name and digest.
func GetImageConfiguration(client http.Client, image string, digest string, token string) string {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://registry-1.docker.io/v2/%s/blobs/%s",image, digest), nil)
	var tk = "Bearer " + token
	req.Header.Add("Authorization", tk)
	req.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	elapsed := time.Since(time.Now())
	resp, err := client.Do(req)
	fmt.Printf("getImageConfiguration raw http request took %s\n", elapsed)

	if err != nil {
		LogAndExit("Could not get a client 2: %s", err.Error())
	}

	if resp == nil || resp.StatusCode != 200 {
		LogAndExit("Could not get a client 3: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogAndExit("Could not get a client 4: %s", err.Error())
	}
	contents := string(body)

	var reqMap map[string]interface{}

	err = json.Unmarshal([]byte(contents), &reqMap)

	configMap := reqMap["config"].(map[string]interface {})

	labels := configMap["Labels"].(map[string]interface {})

	cloudflowLabel := labels["com.lightbend.cloudflow.streamlet-descriptors"].(string)

	return cloudflowLabel
}
