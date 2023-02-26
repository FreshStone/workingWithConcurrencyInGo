package main
import (
	"testing"
	"io"
	"strings"
	"os"
)

func Test_printMessage(t *testing.T){
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w 
	msg = "hey there"
	printMessage()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdout
	w.Close()
	if !strings.Contains(output, "hey there"){
		t.Errorf("error in printMessage")
	}
}

func Test_updateMessage(t *testing.T){
	wg.Add(1) //wg & msgfrom main.go can be accessed from here bcoz its a package level variable
	go updateMessage("gloria", &wg)
	wg.Wait()
	if msg != "gloria"{
		t.Error("error in updateMessage")
	}
}
	

func Test_Main(t *testing.T){
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdout

	if !strings.Contains(output, "hello, universe!"){
		t.Errorf("error in main")
	}
	if !strings.Contains(output, "hello, cosmos!"){
		t.Errorf("error in main")
	}
	if !strings.Contains(output, "hello, world!"){
		t.Errorf("error in main")
	}

}
	
