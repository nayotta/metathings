package encode_decode

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"

	pb "github.com/nayotta/metathings/pkg/proto/identityd"
)

type _User struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Enabled          bool   `json:"enabled"`
	DomainId         string `json:"domain_id"`
	DefaultProjectId string `json:"default_project_id"`
	Nickname         string `json:"nickname,ommitempty"`
	Email            string `json:"email,ommitempty"`
	Phone            string `json:"phone,ommitempty"`
}

type userResponseBody struct {
	User _User `json:"user"`
}

type usersResponseBody struct {
	Users []_User `json:"users"`
}

func copyUser(usr _User) *pb.User {
	pb_usr := &pb.User{
		Id:               usr.Id,
		Name:             usr.Name,
		DefaultProjectId: usr.DefaultProjectId,
		DomainId:         usr.DomainId,
		Enabled:          usr.Enabled,
		Extra:            map[string]string{},
	}

	if usr.Nickname != "" {
		pb_usr.Extra["nickname"] = usr.Nickname
	}
	if usr.Email != "" {
		pb_usr.Extra["email"] = usr.Email
	}
	if usr.Phone != "" {
		pb_usr.Extra["phone"] = usr.Phone
	}

	return pb_usr
}

func copyUsers(usrs []_User) []*pb.User {
	pb_usrs := []*pb.User{}

	for _, usr := range usrs {
		pb_usrs = append(pb_usrs, copyUser(usr))
	}

	return pb_usrs
}

func decodeUser(res gorequest.Response, body string) (*pb.User, error) {
	b := userResponseBody{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return nil, err
	}

	user := copyUser(b.User)

	return user, nil
}

func decodeUsers(res gorequest.Response, body string) ([]*pb.User, error) {
	b := usersResponseBody{}
	err := json.Unmarshal([]byte(body), &b)
	if err != nil {
		return nil, err
	}

	uesrs := copyUsers(b.Users)

	return uesrs, nil
}
