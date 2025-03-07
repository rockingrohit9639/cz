# cz - Commit Management CLI

`cz` is a command-line tool that helps developers craft well-structured commit messages. It provides an interactive experience for selecting commit types, scopes, and writing commit messages. The tool aims to make the process of creating consistent, semantic commit messages easier and more user-friendly.

## Features

- **Interactive Prompts**: Guides the user through a series of prompts to create well-structured commit messages.
- **Commit Types**: Support for standard commit types such as `feat`, `fix`, `docs`, `style`, and more.
- **Scope Selection**: Allows users to specify the scope of the commit, such as a module or feature.
- **Commit Message Generation**: Generates commit messages based on user input using a customizable template.

## Installation

To install `cz`, you can either clone the repository and build it yourself or use pre-built binaries.

### Clone the repository and build

```bash
git clone https://github.com/rockingrohit9639/cz.git
cd cz
go build
```

### Pre-built Binaries

Coming soon!

## Usage

Once installed, you can use cz to create commit messages. Here's an example of how to use the tool:

### Commit Command

To start creating a commit message, run:

```sh
cz
```

The tool will guide you through the following prompts:

Commit Type: Select the type of commit (e.g., feat, fix, docs).
Scope: Optionally specify the scope of the commit.
Message: Write the main commit message.
Body: Add an optional body to provide more details about the commit.

After completing the prompts, you will see a preview of the generated commit message. You can then confirm and proceed with committing the changes.

## Contributing

We welcome contributions! Feel free to open issues, submit pull requests, or suggest improvements.

- Fork the repository
- Create a new branch (git checkout -b feat/some-feature)
- Commit your changes (use our `cz` tool 😄)
- Push to the branch (git push origin feat/some-branch)
- Open a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Future Plans

- [ ] Support for commit templates.
- [ ] Enhancements to the user interface (e.g., improved color schemes).
- [ ] Add flags for type, scope, and message to override interactive prompts.
- [ ] Preview commit message and only commit after confirmation.
- [x] Add retry feature for commit creation.
- [ ] Add undo command to revert the last commit.
- [ ] Warn if no changes are staged
- [ ] Implement `.czrc` file
- [ ] Add command to stage all changes
