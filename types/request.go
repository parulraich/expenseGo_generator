package types

import(
	"net/http"
)
type CreateRequest struct{
      *Expense
}

func (c* CreateRequest) Bind(r *http.Request) error{
   //TODO implement

   return nil

}

type UpdateRequest struct{
    *CreateRequest
}

func (u *UpdateRequest) Bind(r *http.Request) error{
    //TODO implement

      return u.CreateRequest.Bind(r)
}