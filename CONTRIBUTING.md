# Contributing

We welcome contributions from the community! Whether you're fixing a bug, adding a feature, or improving documentation, your help is appreciated. Here's how you can get started:

## 1. Setting Up the Project

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

## 2. Making Changes

1. **Create a New Branch**:
   Create a branch for your changes:
   ```bash
   git checkout -b feat/new-feature
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
   git push origin feat/new-feature
   ```

---

## 3. Submitting a Pull Request

1. **Open a Pull Request**:
   Go to the original repository and click "New Pull Request". Select your branch and provide a detailed description of your changes.

2. **Wait for Review**:
   Your pull request will be reviewed by the maintainers. Be prepared to make additional changes if requested.

---

## 4. Code Style and Guidelines

- **Formatting**: Use `gofmt` or `goimports` to format your code.
- **Testing**: Write unit tests for new functionality. Use table-driven tests where applicable.
- **Documentation**: Add GoDoc comments for all exported types, functions, and methods.

---

## 5. Reporting Issues

If you find a bug or have a feature request, please open an issue on GitHub. Include the following information:
- A clear description of the issue.
- Steps to reproduce the issue.
- Expected and actual behavior.
- Screenshots or logs (if applicable).

---

## 6. Code of Conduct

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project, you agree to abide by its terms.

---

## 7. Thank You!

Thank you for contributing to `dicebear-go`! Your efforts help make this project better for everyone. 🎉
