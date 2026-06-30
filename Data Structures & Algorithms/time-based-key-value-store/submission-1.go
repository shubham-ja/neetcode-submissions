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
	if pairs, exists := this.data[key]; exists{
		l, r := 0, len(pairs)-1

   for l <= r {
       mid := (l + r) / 2
       if pairs[mid].time <= timestamp {
           if mid == len(pairs)-1 || pairs[mid+1].time > timestamp {
               return pairs[mid].value
			   }
		   l = mid + 1
		   } else {
			r = mid - 1
       		}
   		}
	}
	return rv
}
