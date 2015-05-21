package go_koans

import "bytes"
import (
	"fmt"
	"reflect"
	"io"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")
		fmt.Println("(((((((")
		fmt.Println("%b", in.Bytes())

		out := new(bytes.Buffer)

		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		fmt.Println("++++++++")
		fmt.Println("%s", reflect.TypeOf(in))
//		in.ByteReader()

		fmt.Println(out.Bytes())
		fmt.Println(in.Bytes())
		i, error := out.Read(in.Bytes())
		fmt.Println(out.Bytes())
		fmt.Println(i)
		fmt.Println(error)

		io.Copy(out, in)
		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		io.CopyN(out, in, 5)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
