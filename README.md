# FXBE

Http api to access a filesystem and do operation on it


## Usage with curl

```bash

# Get authentification cookie
curl -c cookies.txt "http://localhost:8080/auth/?username=&password="

# Get request to get content of directory of your root folder
curl -b cookies.txt "http://localhost:8080/files/"

# Download a file
curl -b cookies.txt "http://localhost:8080/files/dir/mycoolfile.txt" --output file.txt

# Post a file
curl -b -F 'file=@/full/path/to/file' "http://localhost:8080/files/mynamehere"

# Remove a file or folder

# Copy a file or folder

# Move a file or folder

# Rename a file or folder

# Create a new empty file

# Download multiple directory or file in an archive format

# Create an archive with directory or file


```

