package main

import ( // "fmt"
	// datatypeoperator "github.com/tuhinal/go-experiment/go-basic/data-type-operator"
	// arrayslicemap "github.com/tuhinal/go-experiment/go-basic/arrayslicemap"
	// IOput "github.com/tuhinal/go-experiment/go-basic/core-packages/input-output"
	// structs "github.com/tuhinal/go-experiment/go-basic/structs"
	// datatype "github.com/tuhinal/go-experiment/go-basic/data-type-operator"
	// function "github.com/tuhinal/go-experiment/go-basic/functions"
	// methodInterface "github.com/tuhinal/go-experiment/go-basic/struct-method-interface"
	//  ref "github.com/tuhinal/go-experiment/go-basic/reflection"
	errrorbasic "github.com/tuhinal/go-experiment/go-basic/error"
	//"log"
	"net/http"
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
	// methodInterface.IntSliceDemo()
	// datatype.BlockAndScoping()
	// datatype.SingleInput()
	// function.CalculateCircle()
	// methodInterface.InterfaceExplore()
	// methodInterface.InterfaceComposition()
	// methodInterface.ShortableInterface()
	// errordemo.ErrorInGolang()
	// errordemo.PanicInGolang()
	// errordemo.ReflectionInGolang()

	// ====================== Web App Basic ======================

	/*
		1. initiate a servemux router
		2. register initHandler for "/" URL pattern
	*/
	/* mux := http.NewServeMux()
	mux.HandleFunc("/", initHandler)
	mux.HandleFunc("/student-save", saveStudent)

	log.Print("Server started at localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err) */
	// ref.Reflection()
	errrorbasic.ErrorBasic()
}

func initHandler(write http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(write, req)
		return
	}
	write.Write([]byte("This is your landing page response"))
}

func saveStudent(write http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		write.WriteHeader(http.StatusMethodNotAllowed)
		write.Write([]byte("Method not allowed"))
		return
	}
	// req.URL.
	write.Header().Set("Content-Type", "application/json")
	write.Header().Set("Allow", "POST")
	write.Header().Set("data", "Tuhin")
	// write.Write([]byte("Student created successfully"))
	write.Write([]byte(`{"name": "Alauddin Tuhin", "age": 25, "email": "alu@example.com"}`))
}
