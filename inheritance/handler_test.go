package inheritance

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	a := handlerA{}
	w := handlerWrapper{
		handler: a,
	}
	fmt.Println(w.translate("hello"))
}
