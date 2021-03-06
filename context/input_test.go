// Beego (http://beego.me/)
// @description beego is an open-source, high-performance web framework for the Go programming language.
// @link        http://github.com/astaxie/beego for the canonical source repository
// @license     http://github.com/astaxie/beego/blob/master/LICENSE
// @authors     astaxie

package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestParse(t *testing.T) {
	r, _ := http.NewRequest("GET", "/?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie", nil)
	beegoInput := NewInput(r)
	beegoInput.ParseFormOrMulitForm(1 << 20)

	var id int
	err := beegoInput.Bind(&id, "id")
	if id != 123 || err != nil {
		t.Fatal("id should has int value")
	}
	fmt.Println(id)

	var isok bool
	err = beegoInput.Bind(&isok, "isok")
	if !isok || err != nil {
		t.Fatal("isok should be true")
	}
	fmt.Println(isok)

	var float float64
	err = beegoInput.Bind(&float, "ft")
	if float != 1.2 || err != nil {
		t.Fatal("float should be equal to 1.2")
	}
	fmt.Println(float)

	ol := make([]int, 0, 2)
	err = beegoInput.Bind(&ol, "ol")
	if len(ol) != 2 || err != nil || ol[0] != 1 || ol[1] != 2 {
		t.Fatal("ol should has two elements")
	}
	fmt.Println(ol)

	ul := make([]string, 0, 2)
	err = beegoInput.Bind(&ul, "ul")
	if len(ul) != 2 || err != nil || ul[0] != "str" || ul[1] != "array" {
		t.Fatal("ul should has two elements")
	}
	fmt.Println(ul)

	type User struct {
		Name string
	}
	user := User{}
	err = beegoInput.Bind(&user, "user")
	if err != nil || user.Name != "astaxie" {
		t.Fatal("user should has name")
	}
	fmt.Println(user)
}
