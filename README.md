# concurrency-in-go-course
udemy course working with concurrency in go: https://www.udemy.com/course/working-with-concurrency-in-go-golang/

<br >

# Dining philosophers problem
https://en.wikipedia.org/wiki/Dining_philosophers_problem

solved dining philosophers problem using wait group and mutex locks.

<br >

# Producer-consumer problem
https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem

simulated pizza producer and consumer using golang channels 
<br >

# Sleeping barber problem
https://en.wikipedia.org/wiki/Sleeping_barber_problem

solved the sleeping barber problem with golang channels

<br >


# Subscription service
we have to register a customer for some kind of subscription service, and take care of invoicing, registration, and all the things necessary to get a new customer up and running. We'll do so, naturally, as quickly as we can by dividing the necessary tasks up into smaller tasks (sending invoice email, generate a user manual pdf and sending it via email), and having them run concurrently.

The highlight of the app is SubscribeToPlan in cmd/web/handlers.go