package main

import("html/template"
        "os"
)
type User struct{
  Name string
  Int int
  Float float64
  Slice []string
  Map map[string]string
}
func main(){
  t, err := template.ParseFiles("hello.gohtml")
  if err != nil{
    panic(err)
  }

  data := User{
    Name: "John Smith",
    Int: 123,
    Float: 3.14,
    Slice: []string{"aa","bb","ccc"},
    Map: map[string]string{
      "key1": "value1",
      "key2": "value2",
    },
  }
// chapter 5,6,7 := 5:23:16
// chapter upto 7] := 8:35:06
//chapter 8 := proper 3 hours

/*so agar chapter 8 complete kar lete h is weekend par
to remaing jo rhega bus
so [9,14] := 14 hrs 2hrs/per day for next week
*/


/*
so by the end of the next week hard course remaining of
8hrs + 40 mins

and bonus 1 of 4 hrs
so by next 2 weeks i.e
by end of 24th april ye wala course complete ho jayega
with 2 hrs on this course per day we need 15-17 more days

rest with 2-3 hrs per day on competitive coding
30 more questions i,e 90 questions or 20percent complete in cp












*/
  err = t.Execute(os.Stdout,data)
  if err != nil{
    panic(err)
  }


}
