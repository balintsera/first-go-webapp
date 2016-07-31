package main

func init() {
  // Read file to a constant
}


/**
func readConfigFromFile() (users []User, err error) {
  file, err := ioutil.ReadFile(configFile)
  if err != nil {
    return users, err 
  }

  // Luca 14:45-kor aludt el
  
  err = json.Unmarshal(file, &users)
  if err != nil {
    fmt.Println("error: ", err)
    return users, err
  }
  
  return users, nil
}
*/