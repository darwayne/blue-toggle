package main
import (
  "os/exec"
  "os"
)
func main(){
  cmd := "/usr/local/Cellar/blueutil/2.4.0/bin/blueutil"
  output, _ := exec.Command(cmd, "-p").Output()
  o := string(output)
  if o == "0\n" {
    _, _ = exec.Command(cmd, "-p", "1").Output()
  } else {
   _, _ = exec.Command(cmd, "-p", "0").Output()
  }

  if len(os.Args) == 2 {
   _, _ = exec.Command("pmset", "sleepnow").Output();
  }
}
