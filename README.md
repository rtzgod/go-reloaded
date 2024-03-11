# go-reloaded

01edu project

go-reloaded is a simple text completion/editing/auto-correction tool.

## Installation
1. Install golang 1.20.2

2. Clone repository
```bash
git clone https://github.com/rtzgod/go-reloaded.git
```
3. Create .txt file with your text in /txt folder
```bash
mkdir txt
cd txt/
echo "your text" > yourFileName.txt
```
4. Build code in /go-reloaded dir
```bash
go build -o "correct"
```

## Usage

Program takes two arguments 
First argument is .txt file name with your text
Second argument is txt file name where do you want to save the changes
```bash
./correct 'FileName.txt' 'FileName.txt'
```
