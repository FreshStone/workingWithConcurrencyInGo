package main

import (
	_"fmt"
	"os"
	"sync"
	"io"
	"strings"
	"testing"
)

func Test_print(t *testing.T){
	stdOut := os.Stdout       // capturing
	r, w, _ := os.Pipe()      // output from
	os.Stdout = w             // console
	var wg sync.WaitGroup
	wg.Add(1)
	go print("alpha", &wg)
	wg.Wait()
	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if !strings.Contains(output, "alha"){
		t.Errorf("error in print routine")
	}
}