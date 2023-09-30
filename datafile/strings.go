package datafile  
  
import (  
 "bufio"  
 "os"  
)  
  
func GetStrings(filename string) ([]string, error) {  
 var lines []string  
 file, err := os.Open(filename)  
 if err != nil {  
 return nil, err  
 }  
 defer file.Close()  
  
 scanner := bufio.NewScanner(file)  
 for scanner.Scan() {  
 line := scanner.Text()  
 lines = append(lines, line)  
 }  
  
 if err := scanner.Err(); err != nil {  
 return nil, err  
 }  
   
 return lines, nil  
}