<p align="center">
  <a href="https://amphitheatron.kieranoneill.com">
    <img alt="Aether icon - rounded edges" src="assets/icon-rounded@128x128.png" style="padding-top: 15px" height="128" />
  </a>
</p>

<h1 align="center">
   aether
</h1>

<p align="center">
  Journey into the Aether where storage defies conventional bounds.
</p>

<p align="center">
  Aether is a modern file storage platform that implements a state-of-the-art Merkle tree structure to store multiple files.
</p>

<p align="center">
  <a href="https://github.com/kieranroneill/aether/releases/latest">
    <img alt="GitHub release" src="https://img.shields.io/github/v/release/kieranroneill/aether?&logo=github">
  </a>
  <a href="https://github.com/kieranroneill/amphitheatron/releases/latest">
    <img alt="GitHub release date - published at" src="https://img.shields.io/github/release-date/kieranroneill/aether?logo=github">
  </a>
</p>

<p align="center">
  <a href="https://github.com/kieranroneill/aether/blob/main/COPYING">
    <img alt="GitHub license" src="https://img.shields.io/github/license/kieranroneill/aether">
  </a>
</p>

#### Table of contents

* [1. Overview](#-1-overview)
  - [1.1. Project structure](#11-project-structure)
* [2. Development](#-2-development)
  - [2.1. Requirements](#21-requirements)
  - [2.2. Setting up environment variables (optional)](#22-setting-up-environment-variables-optional)
  - [2.3. Running locally](#23-running-locally)
* [3. Appendix](#-3-appendix)
  - [3.1. Useful commands](#31-useful-commands)
  - [3.2. Docker Compose service directory](#32-docker-compose-service-directory)
* [4. How To Contribute](#-4-how-to-contribute)
* [5. License](#-5-license)

## 🗂️ 1. Overview

### 1.1. Project structure

The project structure is loosely based on the layout outlined in [golang-standards/project-layout](https://github.com/golang-standards/project-layout). The exceptions are:

* The `web` directory has been renamed to `app` to conform to the [Next.js project structure][nextjs-project-structure].
* The `public` directory contains static assets for the web app to conform to [Next.js project structure][nextjs-project-structure].

<sup>[Back to top ^][table-of-contents]</sup>

## 🛠️ 2. Development

### 2.1. Requirements

* [Docker][docker]
* [Docker Compose v2.5.0+][docker-compose]
* [Make][make]
* [Node v20.9.0+][node]
* [Yarn v1.22.5+][yarn]

<sup>[Back to top ^][table-of-contents]</sup>

### 2.2. Setting up environment variables (optional)

1. Create the `.env.*` files for each application to the `.config/` directory:
```shell script
make setup
```

2. Go to the `.config/` directory and edit each `.env.*` file.

<sup>[Back to top ^][table-of-contents]</sup>

### 2.3. Running locally

1. Simply run:
```shell script
make
```

> ⚠️ **NOTE:** The `make` command will run/re-run `make setup`, but will not overwrite any `.env.*` that may have been edited in section [1.2.](#22-setting-up-environment-variables-optional)

2. Navigate to [http://localhost:8080](http://localhost:8080) to access the web portal.

<sup>[Back to top ^][table-of-contents]</sup>

## 📑 3. Appendix

### 3.1. Useful commands

| Command           | Description                                                                                                                              |
|-------------------|------------------------------------------------------------------------------------------------------------------------------------------|
| `make`            | Installs dependencies, setups the basic configuration and runs the Docker Compose configuration. Intended for development purposes only. |
| `make build-core` | Builds the core app into a binary to the `.build/` directory.                                                                            |
| `make build-web`  | Builds the web app to the `.next/` directory.                                                                                            |
| `make dev-core`   | Runs the core app using `go run`. Intended for development purposes only.                                                                |
| `make dev-web`    | Runs the web app using `next dev`. Intended for development purposes only.                                                               |
| `make clean`      | Deletes the build directory.                                                                                                             |
| `make install`    | Installs the yarn and golang dependencies.                                                                                               |
| `make run`        | Checks if the apps are correctly configured and runs Docker Compose.  Intended for development purposes only.                            |
| `make setup`      | Creates `.env.*` files in the `.config/` directory.                                                                                      |

<sup>[Back to top ^][table-of-contents]</sup>

### 3.2. Docker Compose service directory

Here is a list of all the localhost port mappings for each of the apps

| Port   | URL                                            | Docker Compose Service |
|--------|------------------------------------------------|------------------------|
| `3000` | [http://localhost:3000](http://localhost:3000) | `core_app`             |
| `8080` | [http://localhost:8080](http://localhost:8080) | `web_app`              |

<sup>[Back to top ^][table-of-contents]</sup>

## 👏 4. How To Contribute

Please read the [**Contributing Guide**][contribute] to learn about the development process.

<sup>[Back to top ^][table-of-contents]</sup>

## 📄 5. License

Please refer to the [COPYING][copying] file.

<sup>[Back to top ^][table-of-contents]</sup>

<!-- Links -->
[contribute]: ./CONTRIBUTING.md
[copying]: ./COPYING
[docker]: https://docs.docker.com/get-docker/
[docker-compose]: https://docs.docker.com/compose/install/
[make]: https://www.gnu.org/software/make/
[nextjs-project-structure]: https://nextjs.org/docs/getting-started/project-structure
[node]: https://nodejs.org/en/
[table-of-contents]: #table-of-contents
[yarn]: https://yarnpkg.com/
