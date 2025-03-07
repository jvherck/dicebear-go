# Dicebear-go
`dicebear-go` is a package that wraps around [DiceBear](https://dicebear.com)'s API. 
This allows you to easily generate tons of unique avatars in Go. 
The API can generate more than 1 sexdecillion (that's 17 zeroes!) unique avatars!

![dicebear-go version](https://img.shields.io/github/v/tag/jvherck/dicebear-go?sort=semver&label=version)
![Go Version](https://img.shields.io/github/go-mod/go-version/jvherck/dicebear-go)
![License](https://img.shields.io/github/license/jvherck/dicebear-go)
![Code of Conduct](https://img.shields.io/badge/code%20of%20conduct-contributor%20covenant-green.svg)

---

## Other languages
If you want to use Dicebear avatars but don't use Golang, you can use the 
[dicebear Python package](https://github.com/jvherck/dicebear) or the 
[official JS/TS package](https://github.com/dicebear/dicebear).

---

## Table of contents
<!-- TOC -->
* [Dicebear-go](#dicebear-go)
  * [Other languages](#other-languages)
  * [Table of contents](#table-of-contents)
  * [Useful links](#useful-links)
  * [Installation](#installation)
  * [Basic usage](#basic-usage)
  * [Contributing](#contributing)
  * [Credits](#credits)
  * [Licenses and privacy policy](#licenses-and-privacy-policy)
<!-- TOC -->

---

## Useful links
* Official Dicebear: https://dicebear.com/
* Go package: https://pkg.go.dev/github.com/jvherck/dicebear-go
* Docs: https://dicebear-go.vhjan.me/#section-documentation
* GitHub: https://github.com/jvherck/dicebear-go
* Python package: https://github.com/jvherck/dicebear
* JS/TS library: https://github.com/dicebear/dicebear

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

We welcome contributions from the community! Whether you're fixing a bug, adding a feature, or improving documentation, your help is appreciated. Find out in our [Contributing Guide](https://github.com/jvherck/dicebear-go/blob/main/CONTRIBUTING.md) how to contribute to this project.

--- 

## Credits
Special thanks to [Dicebear](https://dicebear.com) ([Florian Körner](https://github.com/FlorianKoerner)) 
for making this amazing API and to [all artists](https://dicebear.com/licenses) 
for creating all the awesome avatars and artwork!

Disclaimer: this repository and its owner are not affiliated with DiceBear.

---

## Licenses and privacy policy
* Dicebear **Licenses**: [https://dicebear.com/licenses](https://dicebear.com/licenses)
* Dicebear **Privacy Policy**: [https://dicebear.com/legal/privacy-policy](https://dicebear.com/legal/privacy-policy)
* Dicebear Go package (this project): [https://dicebear-go.vhjan.me/license](https://dicebear-go.vhjan.me/license)
