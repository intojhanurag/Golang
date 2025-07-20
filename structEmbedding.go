package main

import "fmt"

type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
    base //container automatically gets all the exported methods and fields of base
    str string

	//Container is embedding the base struct
}

func main() {

    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }

    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

    fmt.Println("also num:", co.base.num)

    fmt.Println("describe:", co.describe()) //calling describe through container
	// then you are assigning container to an interface describer
	//co.describe()=co.base.describe()

    type describer interface {
        describe() string
    }
	//defined an interface describer with just one method
	//container has a method describe() (inherited from the base)
	// so it implements the describer interface
	// so we can assign co to variable d of type describer

    var d describer = co
    fmt.Println("describer:", d.describe())
}