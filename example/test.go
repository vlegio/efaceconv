//go:generate efaceconv -bench
//ec::[]byte:SByte
//ec::string:String
//ec::[]int:SInt

package main

func main() {
	str, ok := Eface2String("sd")
	if ok {
		println(*str)
	}
}
