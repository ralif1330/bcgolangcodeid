package main

import "fmt"

func main() {
          //kondisi1()
          //switch1()
          //switch2()
          //switch3()

}

func kondisi1() {
          currentYear := 2023
          tahunLahir := 2015

          if age := currentYear - tahunLahir; age > 17 {
                    fmt.Println("bisa bikin sim")
          } else {
                    fmt.Println("blm bisa bikin sim")
          }
}

func switch1() {

          score := 8
          switch score {
          case 8:
                    fmt.Println("perfect")
          case 7:
                    fmt.Println("awesome")
          default:
                    fmt.Println("notbad")

          }

}

func switch2() {
          score := 8
          switch {
          case score == 8:
                    fmt.Println("perfect")
          case (score > 3) && (score < 8):
                    fmt.Println("not bad")
          default:
                    {
                              fmt.Println("study harder")
                              fmt.Println("u need to learn more")
                    }
          }
}

func switch3() {
          score := 1

          if score > 7 {
                    switch score {
                    case 10:
                              fmt.Println("bagus")
                    default:
                              fmt.Println("nice")
                    }
          } else {
                    if score <= 5 {
                              fmt.Println("notBad")
                    } else if score == 3 {
                              println("keep trying")
                    } else {
                              if score == 0 {
                                        fmt.Println("u can do it")
                                        fmt.Println("try harder")
                              }
                    }
          }
}
