# Interactive Command Shell in Go

This is an interactive command shell written in Go that allows you to execute system commands from the command line.

## Features

- Supports both Windows and Unix-like operating systems (Linux, macOS).
- Executes commands in a new shell session and captures their output.
- Handles basic shell commands and custom commands.
- Displays the current working directory and hostname in the prompt.

## Prerequisites

- Go programming language installed on your system.

## Usage

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/interactive-shell.git
   cd interactive-shell


   ```

2. Compile and run the program:

   ```bash
   go run main.go

   ```

3. The shell will start, displaying the current working directory and hostname in the prompt.

4. Enter a command and press Enter to execute it.

5. To change the current working directory, use the cd command followed by the desired directory:

   ```shell
   cd /path/to/directory
   ```

To exit the shell, press Ctrl+C or enter the exit command.

## Example Usage

1.  ```shell
    /home/user@hostname> ls
    file1.txt  file2.txt  folder1  folder2
    /home/user@hostname> cd folder1
    /home/user/folder1@hostname> pwd
    /home/user/folder1
    /home/user/folder1@hostname> exit
    ```

## Supported Commands

Standard system commands based on the underlying operating system (e.g., ls, dir, echo, pwd, etc.).

Custom commands specified by the user.

## Contributions Welcome

Contributions to this project are welcome and encouraged! If you'd like to contribute:

1. Fork the repo and create a pull request

We appreciate your contributions and look forward to collaborating with you!

## Customization

You can customize this shell by modifying the main.go file. For example, you can add more custom commands or change the appearance of the prompt.

License
This project is licensed under the MIT License - see the LICENSE file for details.
