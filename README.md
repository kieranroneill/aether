<p align="center">
  <a href="https://amphitheatron.kieranoneill.com">
    <img alt="Aether icon - rounded edges" src="assets/icon-rounded@128x128.png" style="padding-top: 15px" height="128" />
  </a>
</p>

<h1 align="center">
   Aether
</h1>

<p align="center">
  A file storage platform that implements a Merkle tree directory to store multiple files.
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

* [1. Development](#-1-development)
  - [1.1. Requirements](#11-requirements)
  - [1.2. Setting up environment variables (optional)](#12-setting-up-environment-variables-optional)
  - [1.3. Running locally](#13-running-locally)
* [2. Appendix](#-2-appendix)
  - [2.1. Useful commands](#21-useful-commands)
  - [Docker Compose service directory](#22-docker-compose-service-directory)
* [3. How To Contribute](#-3-how-to-contribute)
* [4. License](#-4-license)

## 🛠️ 1. Development

### 1.1. Requirements

* Install [Docker][docker]
* Install [Docker Compose v2.5.0+][docker-compose]

<sup>[Back to top ^][table-of-contents]</sup>

### 1.2. Setting up environment variables (optional)

1. Create the `.env.*` files for each application to the `.config/` directory:
```shell script
make install
```

2. Go to the `.config/` directory and edit each `.env.*` file.

<sup>[Back to top ^][table-of-contents]</sup>

### 1.3. Running locally

1. Simply run:
```shell script
make
```

> ⚠️ **NOTE:** The `make` command will re-run `make install`, but will not overwrite any `.env.*` that may have been edited in section [1.2.](#12-setting-up-environment-variables-optional)

2. Navigate to [http://localhost:8080](http://localhost:8080) to access the web portal.

<sup>[Back to top ^][table-of-contents]</sup>

## 📑 2. Appendix

### 2.1. Useful commands

| Command           | Description                                                           |
|-------------------|-----------------------------------------------------------------------|
| `make`            | Setups the basic configuration and runs Docker Compose orchestration. |
| `make install`    | Creates the `.env.*` files to the `.config/` directory.               |
| `make build-core` | Builds the core app into a binary to the `.build/` directory.         |
| `make clean`      | Deletes the build directory                                           |
| `yarn run`        | Checks if the apps are correctly configured and runs Docker Compose.  |

<sup>[Back to top ^][table-of-contents]</sup>

### 2.2. Docker Compose service directory

Here is a list of all the localhost port mappings for each of the apps

| Port   | URL                                            | Docker Compose Service |
|--------|------------------------------------------------|------------------------|
| `3000` | [http://localhost:3000](http://localhost:3000) | `core_app`             |

<sup>[Back to top ^][table-of-contents]</sup>

## 👏 3. How To Contribute

Please read the [**Contributing Guide**][contribute] to learn about the development process.

<sup>[Back to top ^][table-of-contents]</sup>

## 📄 4. License

Please refer to the [COPYING][copying] file.

<sup>[Back to top ^][table-of-contents]</sup>

<!-- Links -->
[contribute]: ./CONTRIBUTING.md
[copying]: ./COPYING
[docker]: https://docs.docker.com/get-docker/
[docker-compose]: https://docs.docker.com/compose/install/
[nodejs]: https://nodejs.org/en/
[table-of-contents]: #table-of-contents
[yarn]: https://yarnpkg.com/
