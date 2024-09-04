package cache

import (
	"container/list"
	"reflect"
	"sync"
	"testing"
)

func TestLru_Get(t *testing.T) {
	type fields struct {
		queue    *list.List
		mutex    *sync.RWMutex
		items    map[string]*list.Element
		capacity int
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
		want   interface{}
	}{
		{
			name: "Get success",
			fields: fields{
				queue:    list.New(),
				mutex:    &sync.RWMutex{},
				items:    make(map[string]*list.Element),
				capacity: 3,
			},
			args: []args{
				{
					key:   "testCase1",
					value: "testValue1",
				},
				{
					key:   "testCase2",
					value: "testValue2",
				},
				{
					key:   "testCase3",
					value: "testValue3",
				},
			},
			want: "testValue1",
		},
		{
			name: "Get failed",
			fields: fields{
				queue:    list.New(),
				mutex:    &sync.RWMutex{},
				items:    make(map[string]*list.Element),
				capacity: 3,
			},
			args: []args{
				{
					key:   "failCase",
					value: nil,
				},
				{
					key:   "testCase",
					value: "testValue",
				},
			},
			want: nil,
		},
		{
			name: "Get success",
			fields: fields{
				queue:    list.New(),
				mutex:    &sync.RWMutex{},
				items:    make(map[string]*list.Element),
				capacity: 3,
			},
			args: []args{
				{
					key:   "testCase1",
					value: "testValue1",
				},
				{
					key:   "testCase2",
					value: "testValue2",
				},
				{
					key:   "testCase3",
					value: "testValue3",
				},
				{
					key:   "testCase4",
					value: "testValue4",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Lru{
				queue:    tt.fields.queue,
				mutex:    tt.fields.mutex,
				items:    tt.fields.items,
				capacity: tt.fields.capacity,
			}
			for _, arg := range tt.args {
				c.Set(arg.key, arg.value)
			}
			if got := c.Get(tt.args[0].key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLru_Set(t *testing.T) {
	type fields struct {
		queue    *list.List
		mutex    *sync.RWMutex
		items    map[string]*list.Element
		capacity int
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Set success",
			fields: fields{
				queue:    list.New(),
				mutex:    &sync.RWMutex{},
				items:    make(map[string]*list.Element),
				capacity: 3,
			},
			args: args{
				key:   "testCase",
				value: "testValue",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Lru{
				queue:    tt.fields.queue,
				mutex:    tt.fields.mutex,
				items:    tt.fields.items,
				capacity: tt.fields.capacity,
			}
			if got := c.Set(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
