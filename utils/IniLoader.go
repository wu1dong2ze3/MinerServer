package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Ini struct {
	fileNamePath string
	iMap         *map[string]interface{}
	sort         *[]string
}
type title struct {
}

func New(fileNamePath string) (*Ini, error) {
	f, err := os.OpenFile(fileNamePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0775)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	mm := make(map[string]interface{}, 0)
	sa := make([]string, 0)
	return &Ini{fileNamePath, &mm, &sa}, err
}
func Open(fileNamePath string) (*Ini, error) {
	f, err := os.OpenFile(fileNamePath, os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var imap = make(map[string]interface{}, 0)
	var temp string
	var kv []string
	var sort = make([]string, 0)
	imapPtr := &imap
	sortPtr := &sort
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		temp = scanner.Text()
		kv = strings.Split(temp, "=")
		if kv == nil || len(kv) != 2 || kv[0] == "" || kv[1] == "" {
			log.Println(kv)
			if strings.Contains(temp, "[") && strings.Contains(temp, "]") {
				//title 内容为title
				(*imapPtr)[temp] = &title{}
				*sortPtr = append(*sortPtr, temp)
			}
			continue
		}
		//碰撞冲突
		if (*imapPtr)[kv[0]] != nil {
			v := (*imapPtr)[kv[0]]
			switch v.(type) {
			case string:
				varray := make([]string, 0)
				varrayPtr := &varray
				*varrayPtr = append(*varrayPtr, v.(string))
				*varrayPtr = append(*varrayPtr, kv[1])
				(*imapPtr)[kv[0]] = varrayPtr
				*sortPtr = append(*sortPtr, kv[1])
				break
			case *[]string:
				varrayPtr := v.(*[]string)
				*varrayPtr = append(*varrayPtr, kv[1])
				*sortPtr = append(*sortPtr, kv[1])
				break
			}
		} else {
			(*imapPtr)[kv[0]] = kv[1]
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &Ini{fileNamePath, imapPtr, sortPtr}, nil
}

func (i Ini) Add(key string, value string) error {
	if key == "" || (!strings.Contains(key, "[") && value == "") {
		log.Println("Add", ParamError)
		return ParamError
	}

	if strings.Contains(key, "[") && strings.Contains(key, "]") {
		//title 内容为title
		(*i.iMap)[key] = &title{}
		*i.sort = append(*i.sort, key)
		return nil
	}
	res, b := (*i.iMap)[key]
	if b == false {
		(*i.iMap)[key] = value
		*i.sort = append(*i.sort, key)
	} else {
		v := res
		switch v.(type) {
		case string:
			varray := make([]string, 0)
			ptr := &varray
			*ptr = append(*ptr, v.(string))
			*ptr = append(*ptr, value)
			*i.sort = append(*i.sort, key)
			(*i.iMap)[key] = ptr
			log.Println("Add1", key, v.(string))
			log.Println("Add1", key, value)
			break
		case *[]string:
			ptr := v.(*[]string)
			*ptr = append(*ptr, value)
			*i.sort = append(*i.sort, key)
			log.Println("Add2", key, value)
			break
		}
	}
	return nil
}

func (i Ini) Load(key string, def string) string {
	return i.LoadByIndex(key, 0, def)
}

func (i Ini) LoadByIndex(key string, index int, def string) string {
	if i.iMap == nil || len(*i.iMap) <= 0 {
		return def
	}
	res, b := (*i.iMap)[key]
	if b == false {
		return def
	}
	v := res
	switch v.(type) {
	case string:
		return v.(string)
	case *[]string:
		varray := v.(*[]string)
		if index < len(*varray) {
			return (*varray)[index]
		}
		break
	case title:
		return key
	}
	return def
}

func (i Ini) Save() error {
	f, err := os.OpenFile(i.fileNamePath, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	if i.iMap == nil || len(*i.iMap) == 0 || i.sort == nil || len(*i.sort) == 0 {
		log.Println("save map is nil!", i.sort)
		return nil
	}
	keyDouble := make([]string, 0)
	var ptrDouble = &keyDouble
	for _, k := range *i.sort {
		if k == "" {
			return ParamError
		}
		v := (*i.iMap)[k]
		if v == nil {
			return ParamError
		}
		switch v.(type) {
		case string:
			str := fmt.Sprintf("%s=%s\n", k, v)
			log.Println("Save1", str)
			if _, err := f.WriteString(str); err != nil {
				return err
			}
			break
		case *[]string:
			//判断是否已经写过
			hasKey := false
			for _, dkey := range *ptrDouble {
				if dkey == k {
					//重复
					hasKey = true
					break
				}
			}
			if hasKey {
				break
			} else {
				keyDouble = append(keyDouble, k)
			}
			list := v.(*[]string)
			for _, v2 := range *list {
				str := fmt.Sprintf("%s=%s\n", k, v2)
				log.Println("Save2", str)
				if _, err := f.WriteString(str); err != nil {
					return err
				}
			}
			break
		case *title:
			str := fmt.Sprintf("%s\n", k)
			if _, err := f.WriteString(str); err != nil {
				return err
			}
		}
	}
	return nil
}
