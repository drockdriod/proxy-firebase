# Proxy Firebase

## Objectives
1. Proxy Firestore updates through a MQTT broker that can be retrieved from different types of applications (i.e. Android, IOS, web app)
2. Utilize FCM for push notifications via the proxy server

## Obstacles
- Scalability
  - each client using this service, will have their own set of MQTT connections. Those connections will have to be managed and be persistently online while the client needs it to be. Managing those connections will have to be done in a scalable manner to avoid possible server overload.
  - firebase will also need their connections managed to be sure we are not going to be cut off by Google

### Possible Solutions
- Scalability
  - In Golang, there exists execution threads called goroutines, which allows code to run concurrently and in parallel. They are lightweight by design meaning multiple instances of goroutines would not be a issue from a horizontal scaling perspective. Each time Firebase listens to a snapshot change on a collection, document or collection within a document, it will occur within a separate execution thread via goroutines. Since these lightweight goroutines can handle multiple listeners at once, the responsibility will not have to lie on multiple server instances, which will assist in scalability

## Resources
- https://godoc.org/github.com/eclipse/paho.mqtt.golang
- https://godoc.org/cloud.google.com/go/firestore
- https://www.eclipse.org/paho/clients/js/
