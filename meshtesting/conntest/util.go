package main

//type stringset map[string]struct{}
//
//func (ss stringset) Set(value string) error {
//	ss[value] = struct{}{}
//	return nil
//}
//
//func (ss stringset) String() string {
//	return strings.Join(ss.slice(), ",")
//}
//
//func (ss stringset) slice() []string {
//	slice := make([]string, 0, len(ss))
//	for k := range ss {
//		slice = append(slice, k)
//	}
//	sort.Strings(slice)
//	return slice
//}
//
//func MustHardwareAddr() string {
//	ifaces, err := net.Interfaces()
//	if err != nil {
//		panic(err)
//	}
//	for _, iface := range ifaces {
//		if s := iface.HardwareAddr.String(); s != "" {
//			return s
//		}
//	}
//	panic("no valid network interfaces")
//}
//
//func MustHostname() string {
//	hostname, err := os.Hostname()
//	if err != nil {
//		panic(err)
//	}
//	return hostname
//}
