Creating a **Docker automation tool** in **Go (Golang)** involves using Go's standard libraries along with the **Docker Remote API** or the **Docker Go SDK** (`github.com/docker/docker/client`). This allows you to programmatically manage Docker containers, images, and volumes.

Here’s a **step-by-step guide** to building a basic Docker automation tool in Go:

---

## 🚀 **Step 1: Prerequisites**

- Install **Go** (https://golang.org/doc/install).
- Install **Docker** and ensure the Docker daemon is running.
- Basic knowledge of **Go programming** and **Docker concepts**.

---

## 📦 **Step 2: Set Up Your Go Project**

```bash
mkdir docker-automation
cd docker-automation
go mod init docker-automation
go get github.com/docker/docker/client
go get github.com/docker/docker/api/types
go get github.com/docker/docker/api/types/container
```

---

## ⚙️ **Step 3: Create the Main Go File**

### 📋 **`main.go`**

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Function to build a Docker image
func buildImage(ctx context.Context, cli *client.Client, imageName string) {
	buildContext := "./" // Assuming a Dockerfile in the current directory

	// Build the image
	buildResponse, err := cli.ImageBuild(ctx, nil, types.ImageBuildOptions{
		Tags:       []string{imageName},
		Remove:     true,
		NoCache:    false,
	})
	if err != nil {
		log.Fatalf("Error building image: %v", err)
	}
	defer buildResponse.Body.Close()

	fmt.Println("Image built successfully:", imageName)
}

// Function to run a Docker container
func runContainer(ctx context.Context, cli *client.Client, imageName string) {
	resp, err := cli.ContainerRun(ctx, imageName, []string{}, nil, container.HostConfig{
		PortBindings: map[string][]container.PortBinding{
			"80/tcp": {{HostPort: "8080"}},
		},
		PublishAllPorts: true,
	}, nil, nil)
	if err != nil {
		log.Fatalf("Error running container: %v", err)
	}
	fmt.Println("Container started:", resp.ID)
}

// Function to list running containers
func listContainers(ctx context.Context, cli *client.Client) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		log.Fatalf("Error listing containers: %v", err)
	}
	for _, container := range containers {
		fmt.Printf("ID: %s, Image: %s, Status: %s\n", container.ID[:10], container.Image, container.Status)
	}
}

func main() {
	// Create Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}
	cli.NegotiateAPIVersion(context.Background())

	// Command-line interface
	for {
		fmt.Println("\n--- Docker Automation Menu ---")
		fmt.Println("1. Build Image")
		fmt.Println("2. Run Container")
		fmt.Println("3. List Containers")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		ctx := context.Background()

		switch choice {
		case 1:
			fmt.Print("Enter image name: ")
			var imageName string
			fmt.Scan(&imageName)
			buildImage(ctx, cli, imageName)
		case 2:
			fmt.Print("Enter image name to run: ")
			var imageName string
			fmt.Scan(&imageName)
			runContainer(ctx, cli, imageName)
		case 3:
			listContainers(ctx, cli)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
```

---

## 📊 **Step 4: Run the Automation Tool**

1. **Build the Go application:**

```bash
go build -o docker-automation
```

2. **Run the application:**

```bash
./docker-automation
```

---

## ⚡ **Step 5: Test the Features**

- **Build Docker Image:**
  - Ensure you have a `Dockerfile` in the same directory.
  - Example Dockerfile:

    ```dockerfile
    FROM nginx:alpine
    CMD ["nginx", "-g", "daemon off;"]
    ```

- **Run a Container:**
  - Map container port `80` to `8080` on your local machine.

- **List Containers:**
  - Displays all running and stopped containers.

---

## 🗒️ **Step 6: Advanced Features (Optional)**

1. **Stop and Remove Containers:**

```go
func stopContainer(ctx context.Context, cli *client.Client, containerID string) {
	err := cli.ContainerStop(ctx, containerID, nil)
	if err != nil {
		log.Fatalf("Error stopping container: %v", err)
	}
	fmt.Println("Container stopped:", containerID)
}
```

2. **Handle Logs:**

```go
func containerLogs(ctx context.Context, cli *client.Client, containerID string) {
	logs, err := cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Tail: "10"})
	if err != nil {
		log.Fatalf("Error fetching logs: %v", err)
	}
	defer logs.Close()

	// Print logs
	io.Copy(os.Stdout, logs)
}
```

3. **Image Cleanup:**

```go
func cleanupImages(ctx context.Context, cli *client.Client) {
	images, _ := cli.ImageList(ctx, types.ImageListOptions{All: true})
	for _, image := range images {
		err := cli.ImageRemove(ctx, image.ID, types.ImageRemoveOptions{Force: true, PruneChildren: true})
		if err != nil {
			log.Printf("Failed to remove image %s: %v", image.ID, err)
		} else {
			fmt.Printf("Removed image: %s\n", image.ID)
		}
	}
}
```

---

## 🎯 **Step 7: Final Thoughts**

- This tool can be extended to support **Docker Compose** or integrated with **CI/CD pipelines**.
- For production-grade tools, consider adding **error handling**, **logging**, and **configuration management**.

---

Would you like to explore more advanced automation, like integrating this with Kubernetes or CI/CD pipelines? 🚀