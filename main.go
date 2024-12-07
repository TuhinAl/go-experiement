package main

import (
	// "fmt"

	// datatypeoperator "github.com/tuhinal/go-experiment/go-basic/data-type-operator"
	// arrayslicemap "github.com/tuhinal/go-experiment/go-basic/arrayslicemap"
	// IOput "github.com/tuhinal/go-experiment/go-basic/core-packages/input-output"
	// structs "github.com/tuhinal/go-experiment/go-basic/structs"
	// datatype "github.com/tuhinal/go-experiment/go-basic/data-type-operator"
	// function "github.com/tuhinal/go-experiment/go-basic/functions"
	methodInterface "github.com/tuhinal/go-experiment/go-basic/struct-method-interface"
)

func main() {
	/* uuid := uuid.NewRandom()
	   fmt.Println(uuid);
	   arrayslicemap.Slice();
	   myEncryptedName := encrypt.Encrypt("Tuhin")
	   myDecryptedName := decrypt.Decrypt(myEncryptedName)
	   fmt.Println(myEncryptedName)
	   fmt.Println(myDecryptedName)
	   myEncryptedUpperName := gochannel.ToLowercase()
	   myDecryptedUpperName := gochannel.ToUppercase()
	   fmt.Println(myEncryptedUpperName)
	   fmt.Println(myDecryptedUpperName) */

	/* start := time.Now()
	for i := 1; i <= 10000; i++ {
		go goroutine.CalculateSquare(i)
	}
	elapsed := time.Since(start)
	time.Sleep(2 * time.Second)
	fmt.Println("Function took: ", elapsed) */

	/* var wg sync.WaitGroup
	start := time.Now()
	wg.Add(10000)
	for i := 1; i <= 10000; i++ {
		go goroutine.CalculateSquareUsingWaitGroup(i, &wg)
	}
	elapsed := time.Since(start)
	wg.Wait()
	fmt.Println("Function toook: ", elapsed) */
	// bufferedchannel.Bufferedchannel()
	// bufferedchannel.ChannelRead()
	// bufferedchannel.ChannelClose()
	// bufferedchannel.Panic()
	// bufferedchannel.ChannelForRange()
	// selectstatement.SelectDemo()
	// selectstatement.DefaultCaseOfSelect()
	// selectstatement.DeadlockAndDefaultCase()
	// selectstatement.NilChannelDefaultCase()
	// selectstatement.RandomSelect()
	// selectstatement.CreatingRaceCondition()
	// concurrencypractice.GoroutineLeak()
	// concurrencypractice.SpawningUpGoroutine()
	// concurrencypractice.Timeout()
	// stringdemo.StringDemo()

	// datatypeoperator.CalculateBoilingPoint()
	// datatypeoperator.PointerPractice()

	/* hasPrefix := datatypeoperator.HasPrefix("tuhin","tuh")
	hasSuffix := datatypeoperator.HasSuffix("Alauddin tuhin","tuhin")
	isSubstring := datatypeoperator.IsSubString("Alauddin tuhin","tuhin")
	fmt.Println(hasPrefix)
	fmt.Println(hasSuffix)
	fmt.Println(isSubstring) */

	// datatypeoperator.UnicodeCodePoint()
	// arrayslicemap.ArrayPractice()
	// arrayslicemap.SliceMethodCall()
	// arrayslicemap.MapMethodCall()
	// structs.EmployeeManipulation()
	// structs.BasicStruct()
	// datatype.InputFromTerminalAndFile()
	// datatype.SortingDemo()
	// arrayslicemap.ReSlicing()
	// IOput.FileIOdemo()
	// IOput.CalculateFileSize()
	// IOput.WordLineAndCharacterCount()
	// function.Fib()
	// function.ClouserWithSlice()
	// function.ClouserDemo()
	// arrayslicemap.SliceInDetails()
	// structs.ComplexDataType()
	// methodInterface.ImplementationTest()
	methodInterface.IntSliceDemo()
}
