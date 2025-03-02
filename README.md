# Dicebear-go
`dicebear-go` is a package that wraps around [dicebear](https://dicebear.com)'s API. 
This allows you to easily generate tons of unique avatars in Go. 
The API can generate more than 1 sexdecillion (that's 17 zeroes!) unique avatars!

![Code of Conduct](https://img.shields.io/badge/code%20of%20conduct-contributor%20covenant-green.svg)
![Go Version](https://img.shields.io/github/go-mod/go-version/yourusername/dicebear)
![License](https://img.shields.io/github/license/yourusername/dicebear)
![Coverage](https://img.shields.io/codecov/c/github/yourusername/dicebear)

---

## Table of contents
<!-- TOC -->
* [Dicebear-go](#dicebear-go)
  * [Table of contents](#table-of-contents)
  * [Installation](#installation)
  * [Basic usage](#basic-usage)
  * [Contributing](#contributing)
    * [1. Setting Up the Project](#1-setting-up-the-project)
    * [2. Making Changes](#2-making-changes)
    * [3. Submitting a Pull Request](#3-submitting-a-pull-request)
    * [4. Code Style and Guidelines](#4-code-style-and-guidelines)
    * [5. Reporting Issues](#5-reporting-issues)
    * [6. Code of Conduct](#6-code-of-conduct)
    * [7. Thank You!](#7-thank-you)
  * [Credits](#credits)
  * [Licenses and privacy policy](#licenses-and-privacy-policy)
<!-- TOC -->

---

## Installation
Install the `dicebear-go` package using the following command:
```shell
go get -u github.com/jvherck/dicebear-go
```

---

## Basic usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/jvherck/dicebear-go"
)

func main() {
	// Create a new avatar with the "adventurer" style and a random seed
	avatar, err := dicebear.NewAvatar(dicebear.Adventurer, "", nil, nil)
	if err != nil {
		log.Fatalf("Failed to create avatar: %v", err)
	}

	// Get the avatar URL
	fmt.Println("Avatar URL:", avatar.URL())

	// Save the avatar as an SVG file
	outputPath := "avatar.svg"
	_, err = avatar.Save(dicebear.SVG, outputPath, false)
	if err != nil {
		log.Fatalf("Failed to save avatar: %v", err)
	}

	fmt.Println("Avatar saved to:", outputPath)
}

/*
Output:
Avatar URL: https://api.dicebear.com/9.x/adventurer/svg?seed=random-seed
Avatar saved to: avatar.svg
*/
```

---

## Contributing

We welcome contributions from the community! Whether you're fixing a bug, adding a feature, or improving documentation, your help is appreciated. Here's how you can get started:

### 1. Setting Up the Project

1. **Fork the Repository**: Click the "Fork" button on the top right of the repository page to create your own copy.

2. **Clone the Repository**:
   ```bash
   git clone https://github.com/jvherck/dicebear-go.git
   cd dicebear-go
   ```

3. **Install Dependencies**:
   Make sure you have Go installed (version 1.21 or higher). Then, install the project dependencies:
   ```bash
   go mod download
   ```

4. **Run the Tests**:
   Ensure all tests pass before making changes:
   ```bash
   go test -v ./...
   ```

---

### 2. Making Changes

1. **Create a New Branch**:
   Create a branch for your changes:
   ```bash
   git checkout -b my-feature-branch
   ```

2. **Make Your Changes**:
  - Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) for code style.
  - Write unit tests for new functionality.
  - Update the documentation if necessary.

3. **Commit Your Changes**:
   Write clear and concise commit messages.
   ```bash
   git commit -m "Add feature X"
   ```

4. **Push Your Changes**:
   Push your branch to your forked repository:
   ```bash
   git push origin my-feature-branch
   ```

---

### 3. Submitting a Pull Request

1. **Open a Pull Request**:
   Go to the original repository and click "New Pull Request". Select your branch and provide a detailed description of your changes.

2. **Wait for Review**:
   Your pull request will be reviewed by the maintainers. Be prepared to make additional changes if requested.

---

### 4. Code Style and Guidelines

- **Formatting**: Use `gofmt` or `goimports` to format your code.
- **Testing**: Write unit tests for new functionality. Use table-driven tests where applicable.
- **Documentation**: Add GoDoc comments for all exported types, functions, and methods.

---

### 5. Reporting Issues

If you find a bug or have a feature request, please open an issue on GitHub. Include the following information:
- A clear description of the issue.
- Steps to reproduce the issue.
- Expected and actual behavior.
- Screenshots or logs (if applicable).

---

### 6. Code of Conduct

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project, you agree to abide by its terms.

---

### 7. Thank You!

Thank you for contributing to `dicebear-go`! Your efforts help make this project better for everyone. 🎉

--- 

## Credits
Special thanks to [Dicebear](https://dicebear.com) ([Florian Körner](https://github.com/FlorianKoerner)) 
for making this amazing API and to [all artists](https://dicebear.com/licenses) 
for creating all the awesome avatars and artwork!

---

## Licenses and privacy policy
* Dicebear **Licenses**: [https://dicebear.com/licenses](https://dicebear.com/licenses)
* Dicebear **Privacy Policy**: [https://dicebear.com/legal/privacy-policy](https://dicebear.com/legal/privacy-policy)
* Dicebear Go package (this project): [https://dicebear-go.vhjan.me/license](https://dicebear-go.vhjan.me/license)