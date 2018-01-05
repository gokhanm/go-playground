// AbstractFactory is a interface which provides a method CreateMyLove().
// This a basic thing we need for Abstract Factory.
// CreateMyLove() is a factory method.
// https://golangvedu.wordpress.com/2017/01/31/golang-design-pattern-abstract-factory-and-factory-method/

package main

import "fmt"

// this is my concrete girl friend
type GirlFriend struct {
       nationality string
       eyesColor string
       language string
}

// abstract factory for creating girlfriend
type AbstractFactory interface {
        CreateMyLove() GirlFriend
}


// my indian girlfriend
type IndianGirlFriendFactory struct {

}

// mu korean girlfirend
type KoreanGirlFriendFactory struct {

}

// concrete implementation
func (a IndianGirlFriendFactory) CreateMyLove() GirlFriend {
       return GirlFriend{"Indian" , "Black" , "Hindi"}
}


// concrete implementation
func (a KoreanGirlFriendFactory) CreateMyLove() GirlFriend {
       return GirlFriend{"Korean" , "Brown" , "Korean"}
}


// main factory method
func getGirlFriend(typeGf string) GirlFriend {

       var gffact AbstractFactory
       switch typeGf {
       case "Indian":
              gffact = IndianGirlFriendFactory{}
              return gffact.CreateMyLove()
       case "Korean":
              gffact = KoreanGirlFriendFactory{}
              return gffact.CreateMyLove()
       }
       return GirlFriend{}
}

func main(){

       var typeGf string
       fmt.Scanf("%s", &typeGf)
       a := getGirlFriend(typeGf)

       fmt.Println(a.eyesColor)

}
