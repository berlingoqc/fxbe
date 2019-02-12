# FXBE

Http api to access a filesystem and do operation on it

Run with 

```bash
fxbe --config ./myconfig.json
```

Example config file

```json
{
    "root": "/my/root/dir",
    "bind": "0.0.0.0:8080"
}
```


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
curl -b cookies.txt --header "Content-Type: application/json" \
  --request POST \
  --data '{"filename":"lol/","recursive":"true"}' \
  http://localhost:8080/op/rm

# Copy a file or folder
curl -b cookies.txt --header "Content-Type: application/json" \
  --request POST \
  --data '{"filename":"lol.txt","destination":"Project/"}' \
  http://localhost:8080/op/cp

# Move a file or folder
curl -b cookies.txt --header "Content-Type: application/json" \
  --request POST \
  --data '{"origin":"lol/","destination:"Project/lol/"}' \
  http://localhost:8080/op/mv

# Download multiple directory or file in an archive format

# Create an archive with directory or file


```

