# Godemon

### What do you have to install before using `godemon`:
1. Go compiler,
2. Git,
3. Wget ( if you're using linux )

### How to use it?
First of all you need to build your version:
`cd ./godemon/`
`go build`

### Using default `./godemon`
Now you can move file `godemon` to folder with your go project and 
start it using command: `./godemon cnf <command>` or `./godemon cmd <fileOrDir> <fileOrModule>`.
In first argument you're giving godemon the path to project, and in second
argument you're choosing is it simple go file or go module - if module pass
argument - `mod` , if file pass argument - `file`

### Installing `godemon` with `installer.sh`
1. `wget https://github.com/nProgrammer/godemon/releases/download/1.0.0/installer.sh`
2. `sudo chmod 777 ./installer.sh`
3. `./installer.sh`

### Installing `godemon` from source code
1. Download source code from releases or use `git clone https://github.com/nProgrammer/godemon`
2. `cd ./godemon`
3. `go build`
4. `sudo cp godemon /bin/godemon`

### Installing `godemon` from `godemon.exe`
1. Download `godemon.exe`
2. Add path to `godemon.exe` to system PATH

### How to work with configuration json file
1. Create in project file `godemon-cnf.json`
2. Create configuration of commands - sample code: 

```
{
    "commands": [
        {
            "name": "run",
            "path": "/home/nwagner/Desktop/godemon/api/",
            "file": "mod"
        }
    ]
}
```
3. Now use command: `godemon cnf <command-name>`