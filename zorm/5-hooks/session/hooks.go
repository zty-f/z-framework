package session

import (
	"reflect"
	"zorm/log"
)

// Hooks constants
const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

// CallMethodByReflect calls the registered callbacks  反射实现
func (s *Session) CallMethodByReflect(method string, value interface{}) {
	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
	return
}

// 使用接口实现钩子函数调用

type IBeforeQuery interface {
	BeforeQuery(s *Session) error
}

type IAfterQuery interface {
	AfterQuery(s *Session) error
}

type IBeforeInsert interface {
	BeforeInsert(s *Session) error
}

/*
按照上诉格式实现对应方法即可。。。。。
*/

// CallMethodByInterface calls the registered hooks  接口实现
func (s *Session) CallMethodByInterface(method string, value interface{}) {
	param := reflect.ValueOf(value)
	switch method {
	case BeforeQuery:
		if i, ok := param.Interface().(IBeforeQuery); ok {
			i.BeforeQuery(s)
		}
	case AfterQuery:
		if i, ok := param.Interface().(IAfterQuery); ok {
			i.AfterQuery(s)
		}
	case BeforeInsert:
		if i, ok := param.Interface().(IBeforeInsert); ok {
			i.BeforeInsert(s)
		}
	// 继续添加对应类型的钩子即可
	default:
		panic("unsupported hook method")
	}
	return
}
