package main

import (
  "bytes"
  "context"
  "flag"
  "fmt"
  "os/exec"

  log "github.com/golang/glog"
)

var (
  inputURI  = flag.String("input_uri", "", "Input file with a number on each line.")
  outputURI = flag.String("output_uri", "", "Output file.")
  runner    = flag.String("runner", "", "Pipeline runner.")

  project         = flag.String("project", "", "Dataflow project.")
  region          = flag.String("region", "", "Dataflow region.")
  tempLocation    = flag.String("temp_location", "", "Dataflow temp location.")
  stagingLocation = flag.String("staging_location", "", "Dataflow staging location.")
  jobName         = flag.String("job_name", "", "Dataflow job name.")
  workerBinary    = flag.String("worker_binary", "/simplesum_main", "Dataflow worker binary.")
)

func main() {
  flag.Parse()

  args := []string{
    "--input_uri=" + *inputURI,
    "--output_uri=" + *outputURI,
    "--runner=" + *runner,
  }

  if *runner == "dataflow" {
    args = append(args,
      "--project="+*project,
      "--region="+*region,
      "--temp_location="+*tempLocation,
      "--staging_location="+*stagingLocation,
      "--worker_binary="+*workerBinary,
    )
    if *jobName != "" {
      args = append(args,
        "--job_name="+*jobName,
      )
    }
  }

  str := *workerBinary
  for _, s := range args {
    str = fmt.Sprintf("%s\n%s", str, s)
  }
  log.Infof("Running command\n%s", str)

  cmd := exec.CommandContext(context.Background(), *workerBinary, args...)
  var out bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &out
  cmd.Stderr = &stderr
  err := cmd.Run()
  if err != nil {
    log.Errorf("%s: %s", err, stderr.String())
  }
  log.Infof("output of cmd: %s", out.String())
}
