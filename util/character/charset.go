package character

import (
	"strings"

	"github.com/juju/errors"
	"golang.org/x/text/encoding"
)

//CharSet represent a collection of collations
type CharSet struct {
	//the name of charset
	name string
	//the collations of charset
	collations map[string]Collation
}


