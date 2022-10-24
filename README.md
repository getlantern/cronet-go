# lantern cronet-go fork

This is a fork of https://github.com/SagerNet/cronet-go to allow it to be build as a part of Lantern.

# cronet-go

Cronet is the Chromium network stack made available to Android apps as a library. Cronet takes advantage of multiple
technologies that reduce the latency and increase the throughput of the network requests that your app needs to work.

The Cronet Library handles the requests of apps used by millions of people on a daily basis, such as YouTube, Google
App, Google Photos, and Maps - Navigation & Transit.

This experimental project ported Cronet to golang with NaiveProxy support. To learn how to use the Cronet Library in
your app, see the [transport](./transport_test.go) and [naive-go](./naive/main.go) example.