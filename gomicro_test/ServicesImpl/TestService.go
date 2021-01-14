package ServicesImpl

import (
	"context"
	"myProject/gomicro_test/Services"
	"strconv"
)

type TestService struct {

}

func newStudent(id int32, name, pwd string) *Services.TestModel{
	return &Services.TestModel{UserId: id,UserName: name,UserPwd: pwd}
}

func (this *TestService) TestReg(ctx context.Context, in *Services.TestRequest, out *Services.TestReponse) error {

	models := make([]*Services.TestModel,0)
	var i int32
	for i = 0; i < in.Size; i++ {
		models = append(models, newStudent( i, "luo"+strconv.Itoa(int(i)),"luo123456") )
	}
	out.Data = models
	out.Message = "Hello, " + in.Name
	out.Status = "Success, Service is Running"
	return nil
}
