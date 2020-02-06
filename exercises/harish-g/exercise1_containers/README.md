# This project builds a container image which prints "Hello Openshift"

# To start, run the following commands:

# To create an image using the dockerfile, cd to project location which has the Dockerfile
```sh
$ podman build . -t <image_name>
```

# To run the image in a container: 

```sh
$ podman run <image_name>
```

# To push the container image to quay.io:

```sh
$ podman push <image_id> <quay_repo_path>
```