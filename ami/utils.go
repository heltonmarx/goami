package ami

import (
	"reflect"
	"strconv"
	"strings"
)

func decode(socket *Socket) (map[string]string, error) {
	message := make(map[string]string)

	for {
		s, err := socket.Recv()
		if err != nil {
			return nil, err
		}
		line := strings.Split(s, "\r\n")
		for i := 0; i < len(line); i++ {
			keys := strings.Split(line[i], ":")
			if len(keys) == 2 {
				action := strings.TrimSpace(keys[0])
				response := strings.TrimSpace(keys[1])
				message[action] = response
			} else if strings.Contains(s, "\r\n\r\n") {
				goto on_exit
			}
		}
	}
on_exit:
	return message, nil
}

func unmarshal(dst interface{}, src map[string]string) {
	s := reflect.ValueOf(dst).Elem()
	t := s.Type()

	for name, value := range src {
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			if f.CanSet() && t.Field(i).Name == name {
				var v interface{}
				switch f.Interface().(type) {
				case bool:
					if value == "yes" {
						v = true
					} else {
						v = false
					}
				case string:
					v = value
				case int, int8, int16, int32, int64:
					v, _ = strconv.Atoi(value)
				}
				f.Set(reflect.ValueOf(v))
				break
			}
		}
	}
}
