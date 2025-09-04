# Zip to PDF Converter

A simple command-line interface (CLI) tool written in Go to convert the contents of a `.zip` archive into a single, well-formatted PDF document.

This tool is ideal for archiving source code, sharing projects for review, or feeding an entire codebase into a Large Language Model (LLM) for context.

## Features

-   **ZIP to PDF**: Converts all files within a zip archive into one PDF.
-   **Code-Friendly Formatting**: Uses a monospaced font, perfect for reading code.
-   **File Headers**: Each file's content in the PDF is preceded by a header indicating its original path.
-   **Layered Architecture**: Clean and maintainable code structure.
-   **Robust Logging**: Provides clear feedback on the conversion process.

## Good Practices Applied

-   **Idiomatic Go**: The code follows standard Go conventions.
-   **Go Modules**: Dependencies are managed using Go Modules.
-   **Layered Architecture**: Separation of concerns (CLI, service, handlers) makes the code testable and scalable.
-   **Robust Error Handling**: Errors are handled gracefully and provide meaningful messages.
-   **Unit Tests**: The project includes unit tests to ensure reliability.

## How to Compile

To compile the binary, you need to have Go installed (version 1.18 or newer).

1.  **Clone the repository (or create the files as described)**:
    ```bash
    git clone [https://github.com/lucasmasceno96/code/zip2pdf.git](https://github.com/lucasmasceno96/code/zip2pdf.git)
    cd zip2pdf
    ```

2.  **Tidy dependencies**:
    ```bash
    go mod tidy
    ```

3.  **Build the binary**:
    ```bash
    go build -o zip2pdf ./cmd/main.go
    ```
    This command will create an executable file named `zip2pdf` in the current directory.

## How to Use

Run the compiled binary from your terminal, providing the path to the `.zip` file as an argument.

### Syntax

```bash
./zip2pdf <path_to_your_zip_file>