package main

import (
	// "fmt"
	// "io"
	"os"
)

func main() {
	// f, err := os.Open("24_files/example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// FileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name : ",FileInfo.Name())
	// fmt.Println("file size : ",FileInfo.Size())
	// fmt.Println("file ModTime : ",FileInfo.ModTime())
	// fmt.Println("file Mode : ",FileInfo.Mode())


	// Read operation 
	// buffer := make([]byte,FileInfo.Size())
	
	// noOfBytes,err := f.Read(buffer)
	// if err!=nil{
	// 	panic(err)
	// }

	// fmt.Printf("Data in files :  <%s> of size = %d \n",buffer,noOfBytes)

	// data, err := os.ReadFile("24_files/example.txt")
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// Read directory
	// dir,err := os.Open(".")
	// if err != nil{
	// 	panic(err)
	// }
	// DirInfo, err := dir.ReadDir(0)
	// if err != nil{
	// 	panic(err)
	// }

	// for d,fi := range DirInfo{
	// 	fmt.Println(fi,d)
	// }


	// // Create a file 
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("Hello Go Learner")

	// // we can write using slice
	// bytes := []byte("\nHello Golang")
	// f.Write(bytes)

	// read and write to another file (streaming fashion)
	// sourceFile, err := os.Open("example2.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example4.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for{
	// 	b, err := reader.ReadByte()
	// 	if err!= nil{
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}

	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e!= nil{
	// 		panic(e)
	// 	}	
	// }

	// writer.Flush()


	// noOfBytes,err := io.Copy(destFile, sourceFile)

	// fmt.Println("wrote ",noOfBytes ," into", destFile.Name()," into sucessfully")


	err := os.Remove("example4.txt")
	if err!= nil{
		panic(err)
	}

}
