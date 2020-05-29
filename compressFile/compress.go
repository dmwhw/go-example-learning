package main

func main(){

}


func getZip(w io.Writer) error {
    zipW := zip.NewWriter(w)
    defer zipW.Close()

    for i := 0; i < 5; i++ {
        f, err := zipW.Create(strconv.Itoa(i) + ".txt")
        if err != nil {
            return err
        }
        _, err = f.Write([]byte(fmt.Sprintf("Hello file %d", i)))
        if err != nil {
            return err
        }
    }
    return nil
}