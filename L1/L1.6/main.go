package main

done := make(chan struct{})

go func() {
    for {
        select {
        case <-done:
            fmt.Println("Stopped by channel")
            return
        default:
            time.Sleep(200 * time.Millisecond)
        }
    }
}()

time.Sleep(time.Second)
close(done)
