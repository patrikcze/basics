# Concurrent Copy Program

This is just a simple program for concurrent copy.

## Short description

In this `program`, we first define `two slices`: one for the `filenames` of the files we want to copy,and one for the names of the `folders` we want to copy. We then use a `sync.WaitGroup` to ensure that all goroutines finish before the program exits.

We use two for loops to iterate over both slices. For each file, we use a `goroutine` to copy the file to a destination folder, and for each folder, we use a goroutine to copy the entire folder and its contents to a destination folder.

Each `goroutine` is defined as `an anonymous function` that takes the name of the file or folder as an argument, allowing us to use the go keyword to create a `separate concurrent thread` of execution for each copy operation.

After all the `goroutines` have finished, we print a message to the console indicating that all files and folders have been copied.

The `copyFolder` function is used to `recursively` copy a folder and its contents. We use the os. Stat function to get information about the source folder, and the os.MkdirAll function to create the destination folder with the same permissions as the source folder.

We use the `os.ReadDir` function to get a list of all the entries in the source folder, and then iterate over them using a for loop. For each entry, we check whether it is a file or directory using the `entry.IsDir` method. If it is a `directory`, we `recursively call copyFolder` with the source and destination paths. If it is a file, we use the `io.Copy` function to copy the file to the destination path.

## Usage

```bash
conc_copy-<platform>-<architecture> <sourcefolder> <destinationfolder>
```

### Example

```bash
conc_copy-lin-amd64 ./source /tmp/destination
```

### Project 

```bash
.
├── README.md
├── bin
│   ├── conc_copy-lin-amd64
│   ├── conc_copy-lin-arm64
│   ├── conc_copy-mac-amd64
│   ├── conc_copy-mac-arm64
│   └── conc_copy-win-amd64.exe
├── go.mod
└── main.go

2 directories, 8 files
```