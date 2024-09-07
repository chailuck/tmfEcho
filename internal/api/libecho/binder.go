package libecho

/*
import (
	"reflect"

	"github.com/labstack/echo"
)

type OMBinder struct{}

func (b *OMBinder) Bind(i interface{}, c echo.Context) (err error) {
	db := new(echo.DefaultBinder)
	if err := db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return err
	}

	// Bind the request to the struct as usual
	if err := db.BindPathParams(c, i); err != nil {
		return err
	}
	if err := db.BindQueryParams(c, i); err != nil {
		return err
	}
	if err := db.BindHeaders(c, i); err != nil {
		return err
	}

	// If the passed interface is a pointer, get the value it points to
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	return
}
*/
/*
type Binder struct {
	binder echo.DefaultBinder
	c      echo.Context
	err    error
}

func NewBinder(c echo.Context) *Binder {
	return &newBinder{
		baseBinder: new(echo.DefaultBinder),
	}
}

func (b *Binder) Error() error {
	return b.err
}





func (b *Binder) Bind(i interface{}) *Binder {
	if b.err == nil {
		b.err = b.binder.Bind(i, b.c)
	}
	return b
}

func (b *Binder) BindPathParams(i interface{}) *Binder {
	if b.err == nil {
		//form+body}
		b.err = b.binder.BindPathParames(i, b.c)
	}
	return b
}

func (b *Binder) BindBody(i interface{}) *Binder {
	if b.err == nil {
		// fixme: there should be check if we have not already read the request body (by double calling bindBody or binding
		b.err = b.binder.BindBody(b.c, i)
	}
	return b
}
*/
