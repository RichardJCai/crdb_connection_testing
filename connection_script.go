package main

import (
  "context"
  "os"
  "bufio"
  "github.com/jackc/pgx"
)

// To run:
// go run connection_script.go [pgurls...]
// go run connection_script.go `{roachprod pgurl --external richardc-test3:1-2}`
func main() {
  ctx := context.Background()
  urls := os.Args[1:]

  numConns := 100000
  var conns []*pgx.Conn
  for i := 0; i < numConns; i++ {
    // Alternate between which URL to connect to.
    url := urls[i%len(urls)]
    conn, err := pgx.Connect(ctx, url)
    conns = append(conns, conn)
    if err != nil {
      println(err.Error())
      continue
    }
  }
  println("enter a key to end:")
  input := bufio.NewScanner(os.Stdin)
  input.Scan()

  for _, conn := range conns {
    conn.Close(ctx)
  }
}
