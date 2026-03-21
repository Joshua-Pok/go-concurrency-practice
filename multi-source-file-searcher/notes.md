<!--markdownlint-disable-->

# Step 1: Implement Function that takes a file path, reads the file line by line and prints any line containing search term to the console


**os.Open**: Used to get a file handle. Musat remember to close file to avoid resource leaks


Returns 2 values: 
1) pointer to the file *os.File
2) error



**bufio.Scanner**: Memory efficient way to read file line by line rather than loading the whole thing at once into RAM
we need to pass a os.File 

returns a scanner to read the iterate the file


use for loop with scanner.Scan(), .Scan() returns true as long as there is another token to be read


inside the loop we use scanner.text() to get the current line as a string


scanner.Err() returns an error or nil if scanner reached end of input successfully


**strings.Contains** check if substring exists within a larger string


## Sketch

func sketch(searchterm string, filepath string)(err){


file, err := os.Open(filepath)
if err != nil{
// return error
}


defer file.Close()


scanner := bufio.NewScanner(file)
for scanner.Scan(){
line := scanner.Text()

if strings.Contains(line, searchterm){
fmt.Println(line)
}
}


if scanner.Err(){
return err
}


}


# Step 2: Recursive Directory walk

**os.ReadDir**  function that returns a list of directory entries for a specific path


**filepath.Join(root, entry.Name())** insetead of manual string concatenation to ensure it works on both windows and unix systems



**info.IsDir()** tells u if directory entry is a directory or file



## Goal: Implement a function in search/engine.go that takes a root path and a search term, finds every file in that directory and calls ProcessFile for each one sequentially



function should be recursive
Sketch:
func Search(rootPath string, searchTerm string) error{


files, err := os.Readdir(rootPath)
if err != nil{
return err
}



for _, file := range(files){
if file.isDir(){
full_path := filepath.Join(rootPath, file.Name())
Search(full_path, searchTerm)
}else{

full_path := filepath.Join(rootPath, file.Name())

err := Processfile(full_path, searchTerm)
if err != nil{
return err
}




}

}





}





# Step 3: The Concurrent Engine


**go** keyword is used to prefix a function to start it in a goroutine. Meaning it does not wait for function to return and moves on immediately


**sync.WaitGroup** is a counter that tracks how many goroutines are running

- **Add(n)** increases the counter by > n
- **Done()** decreases the counter by 1. Usually called inside the Goroutine
- **Wait()** Blocks current execution until counter hits 0


## Closures

When we spanw a go routine using an annonymous function, it "captures" variables from surrounding scope. Meaning it does not maintain its own scope but shares it with the parent function


## Step Goal: Refactor Search Function to spawn a new Goroutine for every new file it finds. Use sync.WaitGroup to ensure program stays alive until every file has been processed
