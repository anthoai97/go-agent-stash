package serializer

// WriteProtobufToJSONFile writes protocol buffer message to JSON file
func WriteArrayStringToByte(array []string) []byte {
	data := ""
	for i := 0; i < len(array); i++ {
		str := array[i]
		if i != len(array)-1 {
			str += "\n"
		}
		data += str
	}
	return []byte(data)
}
