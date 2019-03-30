# Proxy Firebase

## Objectives
1. Proxy Firestore updates through a MQTT broker that can be retrieved from different types of applications (i.e. Android, IOS, web app)
2. Utilize FCM for push notifications via the proxy server

## Obstacles
- Scalability
  - each client that will use this service, will have their own set of MQTT connections. Those connections will have to be managed and be persistently online while the client needs it to be. Managing those connections will have to be done in a scalable manner to avoid possible server overload.
  - firebase will also need their connections managed to be sure we are not going to be cut off by Google
  - 
