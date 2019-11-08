package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var arch string
var supportedArch = []string{"amd64", "arm32v7", "arm64v8"}
var defaultArch = "amd64"

func init() {
	DockerBuildCmd.PersistentFlags().StringVar(&arch, "arch", defaultArch, "target architecture among: "+strings.Join(supportedArch, ", "))
	DockerPushCmd.PersistentFlags().StringVar(&arch, "arch", defaultArch, "target architecture among: "+strings.Join(supportedArch, ", "))

}

func checkArchIsSupported(arch string) {
	for _, a := range supportedArch {
		if arch == a {
			return
		}
	}
	log.Fatal("Architecture is not supported. Please select one of " + strings.Join(supportedArch, ", ") + ".")
}

func dockerBuildOfficialImage(arch string) error {
	docker := &Docker{}
	// Set default Architecture Dockerfile to amd64
	dockerfile := "Dockerfile"

	// If not the default value
	if arch != defaultArch {
		dockerfile = fmt.Sprintf("%s.%s", dockerfile, arch)
	}

	if arch == "arm32v7" {
		err := CommandWithStdout("docker", "run", "--rm", "--privileged", "multiarch/qemu-user-static", "--reset", "-p", "yes").Run()

		if err != nil {
			panic(err)
		}

		err = CommandWithStdout("bash", "-c", "wget https://github.com/multiarch/qemu-user-static/releases/download/v4.1.0-1/qemu-arm-static -O ./qemu-arm-static && chmod +x ./qemu-arm-static").Run()

		if err != nil {
			panic(err)
		}
	} else if arch == "arm64v8" {
		err := CommandWithStdout("docker", "run", "--rm", "--privileged", "multiarch/qemu-user-static", "--reset", "-p", "yes").Run()

		if err != nil {
			panic(err)
		}

		err = CommandWithStdout("bash", "-c", "wget https://github.com/multiarch/qemu-user-static/releases/download/v4.1.0-1/qemu-aarch64-static -O ./qemu-aarch64-static && chmod +x ./qemu-aarch64-static").Run()

		if err != nil {
			panic(err)
		}
	}

	return docker.Build(IntermediateDockerImageName, dockerfile, ".")
}

// DockerBuildCmd Command for building docker image of Authelia.
var DockerBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the docker image of Authelia",
	Run: func(cmd *cobra.Command, args []string) {
		checkArchIsSupported(arch)
		err := dockerBuildOfficialImage(arch)

		if err != nil {
			log.Fatal(err)
		}

		docker := &Docker{}
		err = docker.Tag(IntermediateDockerImageName, DockerImageName)

		if err != nil {
			panic(err)
		}
	},
}

// DockerPushCmd Command for pushing Authelia docker image to Dockerhub
var DockerPushCmd = &cobra.Command{
	Use:   "push-image",
	Short: "Publish Authelia docker image to Dockerhub",
	Run: func(cmd *cobra.Command, args []string) {
		checkArchIsSupported(arch)
		publishDockerImage(arch)
	},
}

// DockerManifestCmd Command for pushing Authelia docker manifest to Dockerhub
var DockerManifestCmd = &cobra.Command{
	Use:   "push-manifest",
	Short: "Publish Authelia docker manifest to Dockerhub",
	Run: func(cmd *cobra.Command, args []string) {
		publishDockerManifest()
	},
}

func login(docker *Docker) {
	username := os.Getenv("DOCKER_USERNAME")
	password := os.Getenv("DOCKER_PASSWORD")

	if username == "" {
		panic(errors.New("DOCKER_USERNAME is empty"))
	}

	if password == "" {
		panic(errors.New("DOCKER_PASSWORD is empty"))
	}

	fmt.Println("Login to dockerhub as " + username)
	err := docker.Login(username, password)

	if err != nil {
		fmt.Println("Login to dockerhub failed")
		panic(err)
	}
}

func deploy(docker *Docker, tag string) {
	imageWithTag := DockerImageName + ":" + tag
	fmt.Println("===================================================")
	fmt.Println("Docker image " + imageWithTag + " will be deployed on Dockerhub.")
	fmt.Println("===================================================")

	err := docker.Tag(DockerImageName, imageWithTag)

	if err != nil {
		panic(err)
	}

	err = docker.Push(imageWithTag)

	if err != nil {
		panic(err)
	}
}

func deployManifest(docker *Docker, tag string, amd64tag string, arm32v7tag string, arm64v8tag string) {
	dockerImagePrefix := DockerImageName + ":"
	fmt.Println("===================================================")
	fmt.Println("Docker manifest " + dockerImagePrefix + tag + " will be deployed on Dockerhub.")
	fmt.Println("===================================================")

	err := docker.Manifest(dockerImagePrefix+tag, dockerImagePrefix+amd64tag, dockerImagePrefix+arm32v7tag, dockerImagePrefix+arm64v8tag)

	if err != nil {
		panic(err)
	}

	tags := []string{amd64tag, arm32v7tag, arm64v8tag}
	for _, t := range tags {
		username := os.Getenv("DOCKER_USERNAME")
		password := os.Getenv("DOCKER_PASSWORD")

		fmt.Println("===================================================")
		fmt.Println("Docker removing tag for " + dockerImagePrefix + t + " on Dockerhub.")
		fmt.Println("===================================================")

		err = docker.CleanTag(username, password, t)

		if err != nil {
			panic(err)
		}
	}
}

func publishDockerImage(arch string) {
	docker := &Docker{}

	travisBranch := os.Getenv("TRAVIS_BRANCH")
	travisPullRequest := os.Getenv("TRAVIS_PULL_REQUEST")
	travisTag := os.Getenv("TRAVIS_TAG")

	if travisBranch == "master" && travisPullRequest == "false" {
		login(docker)
		deploy(docker, "master-"+arch)
	} else if travisTag != "" {
		login(docker)
		deploy(docker, travisTag+"-"+arch)
		deploy(docker, "latest-"+arch)
	} else {
		fmt.Println("Docker image will not be published")
	}
}

func publishDockerManifest() {
	docker := &Docker{}

	travisBranch := os.Getenv("TRAVIS_BRANCH")
	travisPullRequest := os.Getenv("TRAVIS_PULL_REQUEST")
	travisTag := os.Getenv("TRAVIS_TAG")

	if travisBranch == "master" && travisPullRequest == "false" {
		login(docker)
		deployManifest(docker, "master", "master-amd64", "master-arm32v7", "master-arm64v8")
	} else if travisTag != "" {
		login(docker)
		deployManifest(docker, travisTag, travisTag+"-amd64", travisTag+"-arm32v7", travisTag+"-arm64v8")
		deployManifest(docker, "latest", "latest-amd64", "latest-arm32v7", "latest-arm64v8")
	} else {
		fmt.Println("Docker manifest will not be published")
	}
}
