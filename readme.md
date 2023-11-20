<div align="center">
<h1 align="center">
<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" />
<br>NASTAPP-API</h1>
<h3>API for Nastapp, an app that keeps track of the prices of gas for several gas stations across a region. </h3>


<p align="center">
<img src="https://img.shields.io/badge/Files-4285F4.svg?style=for-the-badge&logo=Files&logoColor=white" alt="Files" />
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=for-the-badge&logo=Go&logoColor=white" alt="Go" />
</p>
<img src="https://img.shields.io/github/license/juampiq6/nastapp-api?style=for-the-badge&color=5D6D7E" alt="GitHub license" />
<img src="https://img.shields.io/github/last-commit/juampiq6/nastapp-api?style=for-the-badge&color=5D6D7E" alt="git-last-commit" />
<img src="https://img.shields.io/github/commit-activity/m/juampiq6/nastapp-api?style=for-the-badge&color=5D6D7E" alt="GitHub commit activity" />
<img src="https://img.shields.io/github/languages/top/juampiq6/nastapp-api?style=for-the-badge&color=5D6D7E" alt="GitHub top language" />
</div>

---

## 📖 Table of Contents
- [📖 Table of Contents](#-table-of-contents)
- [📍 Overview](#-overview)
- [📦 Features](#-features)
- [📂 Repository Structure](#-repository-structure)
- [⚙️ Modules](#️-modules)
- [🚀 Getting Started](#-getting-started)
  - [🔧 Installation](#-installation)
  - [🤖 Running nastapp-api](#-running-nastapp-api)
  - [🧪 Tests](#-tests)
- [🛣 Project Roadmap](#-project-roadmap)
- [🤝 Contributing](#-contributing)
    - [*Contributing Guidelines*](#contributing-guidelines)
- [📄 License](#-license)
- [👏 Acknowledgments](#-acknowledgments)

---


## 📍 Overview

► [WIP]

---

## 📦 Features

► Get gas stations near a point at certain maximum distance

---


## 📂 Repository Structure

```sh
└── nastapp-api/
    ├── .tool-versions
    ├── TODOS
    ├── db/
    │   ├── db_repository.go
    │   ├── lat_long.go
    │   ├── mongo_client.go
    │   └── places_result.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── router/
        ├── api_error.go
        ├── handlers.go
        ├── router_setup.go
        └── swagger_setup.go

```

---


## ⚙️ Modules


<details closed><summary>Db</summary>

| File                                                                                      | Summary       |
| ----------------------------------------------------------------------------------------- | ------------- |
| [mongo_client.go](https://github.com/juampiq6/nastapp-api/blob/main/db/mongo_client.go)   | ► INSERT-TEXT |
| [db_repository.go](https://github.com/juampiq6/nastapp-api/blob/main/db/db_repository.go) | ► INSERT-TEXT |

</details>

<details closed><summary>Router</summary>

| File                                                                                          | Summary       |
| --------------------------------------------------------------------------------------------- | ------------- |
| [handlers.go](https://github.com/juampiq6/nastapp-api/blob/main/router/handlers.go)           | ► INSERT-TEXT |
| [router_setup.go](https://github.com/juampiq6/nastapp-api/blob/main/router/router_setup.go)   | ► INSERT-TEXT |
| [swagger_setup.go](https://github.com/juampiq6/nastapp-api/blob/main/router/swagger_setup.go) | ► INSERT-TEXT |

</details>

---

## 🚀 Getting Started

***Dependencies***

`- ℹ️ Go 1.21.3`

### 🔧 Installation

1. Clone the nastapp-api repository:
```sh
git clone https://github.com/juampiq6/nastapp-api
```

2. Change to the project directory:
```sh
cd nastapp-api
```

3. Install the dependencies:
```sh
go get
```

### 🤖 Running nastapp-api

```sh
go run *.go
```

### 🧪 Tests
```sh
go test
```

---


## 🛣 Project Roadmap

See TODOS file

> - [X] `ℹ️  Task 1: Implement X`
> - [ ] `ℹ️  Task 2: Implement Y`
> - [ ] `ℹ️ ...`


---

## 🤝 Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Submit Pull Requests](https://github.com/juampiq6/nastapp-api/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/juampiq6/nastapp-api/discussions)**: Share your insights, provide feedback, or ask questions.
- **[Report Issues](https://github.com/juampiq6/nastapp-api/issues)**: Submit bugs found or log feature requests for JUAMPIQ6.

#### *Contributing Guidelines*

<details closed>
<summary>Click to expand</summary>

1. **Fork the Repository**: Start by forking the project repository to your GitHub account.
2. **Clone Locally**: Clone the forked repository to your local machine using a Git client.
   ```sh
   git clone <your-forked-repo-url>
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear and concise message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to GitHub**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.

Once your PR is reviewed and approved, it will be merged into the main branch.

</details>

---

## 📄 License


This project is protected under the [SELECT-A-LICENSE](https://choosealicense.com/licenses) License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/) file.

---

## 👏 Acknowledgments

- List any resources, contributors, inspiration, etc. here.

[**Return to top**](#Top)

---

