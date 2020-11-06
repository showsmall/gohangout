package filter

import "testing"

func TestIncludeKeys(t *testing.T) {
	config := make(map[interface{}]interface{})
	config["field_split"] = " "
	config["value_split"] = "="
	config["src"] = "message"
	config["include_keys"] = []interface{}{"a", "b", "c", "xyz"}
	config["exclude_keys"] = []interface{}{"c"}
	f := methodLibrary.NewKVFilter(config)

	event := make(map[string]interface{})
	event["message"] = "a=aaa b=bbb c=ccc xyz=\txyzxyz\t d=ddd"
	t.Log(event)

	event, ok := f.Filter(event)
	if !ok {
		t.Error("kv failed")
	}
	t.Log(event)

	if event["a"] != "aaa" {
		t.Error("kv failed")
	}
	if event["b"] != "bbb" {
		t.Error("kv failed")
	}
	if _, ok := event["c"]; ok {
		t.Error("c is excluded")
	}
	if event["xyz"] != "\txyzxyz\t" {
		t.Error("kv failed")
	}
	if _, ok := event["d"]; ok {
		t.Error("d is excluded")
	}
}

func TestKVFilter(t *testing.T) {
	config := make(map[interface{}]interface{})
	config["field_split"] = " "
	config["value_split"] = "="
	config["src"] = "message"
	f := methodLibrary.NewKVFilter(config)

	event := make(map[string]interface{})
	event["message"] = "a=aaa b=bbb c=ccc xyz=\txyzxyz\t d=ddd"
	t.Log(event)

	event, ok := f.Filter(event)
	if !ok {
		t.Error("kv failed")
	}
	t.Log(event)

	if event["a"] != "aaa" {
		t.Error("kv failed")
	}
	if event["b"] != "bbb" {
		t.Error("kv failed")
	}
	if event["c"] != "ccc" {
		t.Error("kv failed")
	}
	if event["xyz"] != "\txyzxyz\t" {
		t.Error("kv failed")
	}
	if event["d"] != "ddd" {
		t.Error("kv failed")
	}

	// trim
	config = make(map[interface{}]interface{})
	config["field_split"] = " "
	config["value_split"] = "="
	config["trim"] = "\t \""
	config["trim_key"] = `"`
	config["src"] = "message"
	f = methodLibrary.NewKVFilter(config)

	event = make(map[string]interface{})
	event["message"] = "a=aaa b=bbb xyz=\"\txyzxyz\t\" d=ddd"
	t.Log(event)

	event, ok = f.Filter(event)
	if !ok {
		t.Error("kv failed")
	}
	t.Log(event)

	if event["a"] != "aaa" {
		t.Error("kv failed")
	}
	if event["b"] != "bbb" {
		t.Error("kv failed")
	}
	if event["xyz"] != "xyzxyz" {
		t.Error("kv failed")
	}
	if event["d"] != "ddd" {
		t.Error("kv failed")
	}
}
