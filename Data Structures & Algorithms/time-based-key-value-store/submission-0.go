type entry struct {
	time int
	value string
}
type TimeMap struct {
	data map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap {
		data : make(map[string][]entry),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	newEntry := entry{
		time: timestamp,
		value: value,
	}
	if _,exists := this.data[key]; exists{
		this.data[key] = append(this.data[key], newEntry)
	} else {
		this.data[key] = []entry{newEntry}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	rv := ""
	if value, exists := this.data[key]; exists{
		for _, e := range value {
			if e.time <= timestamp {
				rv = e.value
			}
		}
	}
	return rv
}
