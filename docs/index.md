<p style="text-align: center; margin-left: 1.6rem;">
  <img alt="Logotype depicting a lighthouse" src="./images/logo-450px.png" width="450" />
</p>
<h1 align="center">
  Watchtower
</h1>

<p align="center">
  A container-based solution for automating Docker container base image updates.
  <br/><br/>
  <a href="https://godoc.org/github.com/timschumi/watchtower">
    <img alt="GoDoc" src="https://godoc.org/github.com/timschumi/watchtower?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/timschumi/watchtower">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/timschumi/watchtower" />
  </a>
  <a href="https://github.com/timschumi/watchtower/releases">
    <img alt="latest version" src="https://img.shields.io/github/tag/timschumi/watchtower.svg" />
  </a>
  <a href="https://www.apache.org/licenses/LICENSE-2.0">
    <img alt="Apache-2.0 License" src="https://img.shields.io/github/license/timschumi/watchtower.svg" />
  </a>
  <a href="https://github.com/timschumi/watchtower/#contributors">
    <img alt="All Contributors" src="https://img.shields.io/github/all-contributors/timschumi/watchtower" />
  </a>
  <a href="https://hub.docker.com/r/timschumi/watchtower">
    <img alt="Pulls from DockerHub" src="https://img.shields.io/docker/pulls/timschumi/watchtower.svg" />
  </a>
</p>

## Quick Start

With watchtower you can update the running version of your containerized app simply by pushing a new image to the Docker
Hub or your own image registry. Watchtower will pull down your new image, gracefully shut down your existing container
and restart it with the same options that were used when it was deployed initially. Run the watchtower container with
the following command:

=== "docker run"

    ```bash
    $ docker run -d \
    --name watchtower \
    -v /var/run/docker.sock:/var/run/docker.sock \
    timschumi/watchtower
    ```

=== "docker-compose.yml"

    ```yaml
    version: "3"
    services:
      watchtower:
        image: timschumi/watchtower
        volumes:
          - /var/run/docker.sock:/var/run/docker.sock
    ```
