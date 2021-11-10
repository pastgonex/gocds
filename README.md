# gocds
A go language implementation of a CLI, based on input directory files, using tree output of all file names, based on input database links, dynamic reading of database table information, based on input swagger files to automate the generation of the RESTFUL API.

## Quick Start

### Build and run
```shell
# cd ./gocds
go build -o cds
export PATH=`pwd`:$PATH
```

### Example
```sh
cds dirtree [Path] # Path to the directory file
cds db [user]:[passwd]@tcp(url:port)/[databaseName]" [package name]  # [Package name] is the directory where the generated go files will be placed
cds s2r [Path] # Path to swagger file 
```

## Features
- Tree output of all files and folders based on the input path.
- Based on the database link, dynamically read the database table information and generate the corresponding structure.
- Automated generation of restful api based on input swagger file.

## License
MIT License

Copyright (c) 2021 Eric

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

